package ergo

import (
	"context"
	"encoding/hex"
	"fmt"
	"github.com/SuSy-One/susy-v2/node/pkg/readiness"
	"github.com/SuSy-One/susy-v2/node/pkg/vaa"
	"github.com/mr-tron/base58"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"time"

	"github.com/SuSy-One/susy-v2/node/pkg/p2p"
	gossipv1 "github.com/SuSy-One/susy-v2/node/pkg/proto/gossip/v1"
	"go.uber.org/zap"

	"github.com/SuSy-One/susy-v2/node/pkg/common"
	"github.com/SuSy-One/susy-v2/node/pkg/supervisor"
	eth_common "github.com/ethereum/go-ethereum/common"
)

var (
	ergoConnectionErrors = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "wormhole_ergo_connection_errors_total",
			Help: "Total number of Ergo connection errors",
		}, []string{"ergo_network", "reason"})

	currentErgoHeight = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "wormhole_ergo_current_height",
			Help: "Current Ergo block height",
		}, []string{"operation"})
	queryLatency = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name: "wormhole_ergo_query_latency",
			Help: "Latency histogram for Ergo calls",
		}, []string{"operation"})
)

type (
	ErgoWatcher struct {
		// ergo watcher url
		url string

		// Address of the ergo contract address
		contract ErgoAddress

		// Channel to send new messages to.
		msgChan chan *common.MessagePublication

		// Channel to send guardian set changes to.
		// setChan can be set to nil if no guardian set changes are needed.

		setChan chan *common.ErgoGuardianSet

		// 0 is a valid guardian set, so we need a nil value here
		currentGuardianSet *uint32
	}

	PendingMessage struct {
		Message common.MessagePublication
		Height  uint64
	}
)

func NewErgoWatcher(
	url string,
	messageEvents chan *common.MessagePublication,
	setEvents chan *common.ErgoGuardianSet) *ErgoWatcher {
	return &ErgoWatcher{
		url:     url,
		msgChan: messageEvents,
		setChan: setEvents,
	}
}

func (e *ErgoWatcher) Run(ctx context.Context) error {
	logger := supervisor.Logger(ctx)
	var lastHeight int64
	errC := make(chan error)

	// Initialize gossip metrics (we want to broadcast the address even if we're not yet syncing)
	contractAddr := base58.Encode(e.contract[:])
	p2p.DefaultRegistry.SetNetworkStats(vaa.ChainIDErgo, &gossipv1.Heartbeat_Network{
		ContractAddress: contractAddr,
	})

	client, err := NewClient(ErgOptions{ApiKey: "", BaseUrl: e.url})
	if err != nil {
		p2p.DefaultRegistry.AddErrorCount(vaa.ChainIDErgo, 1)
		return fmt.Errorf("dialing ergo client failed: %w", err)
	}

	// Fetch initial guardian set
	if err := e.fetchAndUpdateGuardianSet(logger, ctx, client); err != nil {
		return fmt.Errorf("failed to request guardian set: %v", err)
	}

	// Poll for guardian set.
	go func() {
		t := time.NewTicker(15 * time.Second)
		defer t.Stop()
		for {
			select {
			case <-ctx.Done():
				return
			case <-t.C:
				if err := e.fetchAndUpdateGuardianSet(logger, ctx, client); err != nil {
					logger.Error("failed updating guardian set")
				}
			}
		}
	}()

	go func() {
		timer := time.NewTicker(time.Second * 1)
		defer timer.Stop()

		for {
			select {
			case <-ctx.Done():
				return
			case <-timer.C:

				start := time.Now()

				// Watch headers
				height, err := client.getlastHeight(ctx)
				if err != nil {
					ergoConnectionErrors.WithLabelValues("get_last_height_error").Inc()
					p2p.DefaultRegistry.AddErrorCount(vaa.ChainIDErgo, 1)
					logger.Error("failed to get last height events: %w", zap.Error(err))
					errC <- err
					return
				}

				if lastHeight == 0 {
					lastHeight = height - 1
				}
				currentErgoHeight.WithLabelValues("get_last_height").Set(float64(height))
				readiness.SetReady(common.ReadinessSolanaSyncing)
				p2p.DefaultRegistry.SetNetworkStats(vaa.ChainIDSolana, &gossipv1.Heartbeat_Network{
					Height:          height,
					ContractAddress: contractAddr,
				})

				rangeStart := lastHeight + 1
				rangeEnd := height

				logger.Info("fetching slots in range",
					zap.Uint64("from", uint64(rangeStart)), zap.Uint64("to", uint64(rangeEnd)),
					zap.Duration("took", time.Since(start)))
				// Requesting each slot
				if rangeStart <= rangeEnd {
					go e.retryGetVAAData(ctx, client, logger, rangeStart, rangeEnd, errC)
				}
				lastHeight = height
			}
		}
	}()

	select {
	case <-ctx.Done():
		return ctx.Err()
	case err := <-errC:
		return err
	}
}

func (e *ErgoWatcher) fetchAndUpdateGuardianSet(
	logger *zap.Logger,
	ctx context.Context,
	client *ErgClient,
) error {
	msm := time.Now()
	logger.Info("fetching guardian set")
	timeout, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()
	gs, err := fetchCurrentGuardianSet(timeout, client)
	if err != nil {
		p2p.DefaultRegistry.AddErrorCount(vaa.ChainIDErgo, 1)
		return err
	}

	queryLatency.WithLabelValues("get_guardian_set").Observe(time.Since(msm).Seconds())

	if e.currentGuardianSet != nil && *(e.currentGuardianSet) == gs.Index {
		return nil
	}

	logger.Info("updated guardian set found",
		zap.Any("value", gs), zap.Uint32("index", gs.Index))

	e.currentGuardianSet = &gs.Index

	if e.setChan != nil {
		e.setChan <- &common.ErgoGuardianSet{
			ErgoKeys:     gs.ErgoKeys,
			WormholeKeys: gs.WormholeKeys,
			Index:        gs.Index,
		}
	}

	return nil
}

// Fetch the current guardian set ID and guardian set from the chain.
func fetchCurrentGuardianSet(ctx context.Context, client *ErgClient) (*common.ErgoGuardianSet, error) {

	gs, err := client.GetCurrentGuardianSet(ctx)
	if err != nil {
		return nil, fmt.Errorf("error requesting current guardian set value: %w", err)
	}

	return &gs, nil
}

func (e *ErgoWatcher) retryGetVAAData(ctx context.Context, client *ErgClient, logger *zap.Logger, offsetHeight int64, limitHeight int64, errC chan error) {

	messages, err := client.getObservations(ctx, offsetHeight, limitHeight)
	if err != nil {
		p2p.DefaultRegistry.AddErrorCount(vaa.ChainIDErgo, 1)
		logger.Error("failed to get message publication events: %w", zap.Error(err))
		errC <- err
		return
	}
	// Request timestamp for block
	for _, msg := range messages {
		txHash, _ := hex.DecodeString(msg.TxId)
		emitterAddressBytes, _ := hex.DecodeString(msg.EmitterAddress)
		addr := vaa.Address{}
		copy(addr[:], emitterAddressBytes)

		observation := &common.MessagePublication{
			TxHash:           eth_common.BytesToHash(txHash),
			Timestamp:        time.Unix(int64(msg.Timestamp), 0),
			Nonce:            msg.Nonce,
			Sequence:         msg.Sequence,
			EmitterChain:     vaa.ChainIDErgo,
			EmitterAddress:   addr,
			Payload:          msg.Payload,
			ConsistencyLevel: msg.ConsistencyLevel,
		}

		logger.Info("message observed",
			zap.Uint64("height", msg.Height),
			zap.Time("timestamp", observation.Timestamp),
			zap.Uint32("nonce", observation.Nonce),
			zap.Uint64("sequence", observation.Sequence),
			zap.Stringer("emitter_chain", observation.EmitterChain),
			zap.Stringer("emitter_address", observation.EmitterAddress),
			zap.Binary("payload", observation.Payload),
			zap.Uint8("consistency_level", observation.ConsistencyLevel),
		)
		logger.Info("found new message publication transaction", zap.String("tx", msg.TxId),
			zap.Uint64("height", msg.Height))

		e.msgChan <- observation
	}
}
