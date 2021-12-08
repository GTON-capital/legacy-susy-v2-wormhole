package guardiand

import (
	"fmt"
	"io/ioutil"
	"os"

	p2pcrypto "github.com/libp2p/go-libp2p-core/crypto"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var NodeKeygenCmd = &cobra.Command{
	Use:   "node-keygen [KEYFILE]",
	Short: "Create node key at the specified path",
	Run:   runNodeKeygen,
	Args:  cobra.ExactArgs(1),
}

func runNodeKeygen(cmd *cobra.Command, args []string) {
	key, err := getOrCreateNodeKey(zap.L(), args[0])
	if err != nil {
		panic(err)
	}
	peerID, err := peer.IDFromPrivateKey(key)
	if err != nil {
		panic(err)
	}
	fmt.Printf("PeerID: %s\n", peerID.String())

}
func getOrCreateNodeKey(logger *zap.Logger, path string) (p2pcrypto.PrivKey, error) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			logger.Info("No node key found, generating a new one...", zap.String("path", path))

			priv, _, err := p2pcrypto.GenerateKeyPair(p2pcrypto.Ed25519, -1)
			if err != nil {
				panic(err)
			}

			s, err := p2pcrypto.MarshalPrivateKey(priv)
			if err != nil {
				panic(err)
			}

			err = ioutil.WriteFile(path, s, 0600)
			if err != nil {
				return nil, fmt.Errorf("failed to write node key: %w", err)
			}

			return priv, nil
		} else {
			return nil, fmt.Errorf("failed to read node key: %w", err)
		}
	}

	priv, err := p2pcrypto.UnmarshalPrivateKey(b)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal node key: %w", err)
	}

	peerID, err := peer.IDFromPrivateKey(priv)
	if err != nil {
		panic(err)
	}

	logger.Info("Found existing node key",
		zap.String("path", path),
		zap.Stringer("peerID", peerID))

	return priv, nil
}
