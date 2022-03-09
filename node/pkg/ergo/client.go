package ergo

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/SuSy-One/susy-v2/node/pkg/common"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"strings"
	"time"
)

type ErgoAddress [36]byte

type ErgDoer interface {
	Do(req *http.Request) (*http.Response, error)
}

type ErgOptions struct {
	BaseUrl string
	Doer    ErgDoer
	ApiKey  string
}

var DefaultOptions = ErgOptions{
	BaseUrl: "",
	Doer:    &http.Client{Timeout: 30 * time.Second},
}

type ErgClient struct {
	Options ErgOptions
}

type Response struct {
	*http.Response
}

type RequestError struct {
	Err  error
	Body string
}

func (a *RequestError) Error() string {
	if a.Body != "" {
		return errors.Wrap(a.Err, a.Body).Error()
	}
	return a.Err.Error()
}

type ParseError struct {
	Err error
}

type Message struct {
	TxId             string `json:"txId"`
	Timestamp        uint64 `json:"timestamp"`
	Nonce            uint32 `json:"nonce"`
	Sequence         uint64 `json:"sequence"`
	ConsistencyLevel uint8  `json:"consistencyLevel"`
	EmitterAddress   string `json:"emitterAddress"`
	Payload          []byte `json:"payload"`
	Height           uint64 `json:"height"`
}

func (a ParseError) Error() string {
	return a.Err.Error()
}

// Creates new client instance
// If no options provided will use default
func NewClient(options ...ErgOptions) (*ErgClient, error) {
	if len(options) > 1 {
		return nil, errors.New("too many options provided. Expects no or just one item")
	}

	opts := DefaultOptions

	if len(options) == 1 {
		option := options[0]

		if option.BaseUrl != "" {
			opts.BaseUrl = option.BaseUrl
		}

		if option.Doer != nil {
			opts.Doer = option.Doer
		}

		if option.ApiKey != "" {
			opts.ApiKey = option.ApiKey
		}
	}

	c := &ErgClient{
		Options: opts,
	}

	return c, nil
}

func (client ErgClient) GetOptions() ErgOptions {
	return client.Options
}

func newResponse(response *http.Response) *Response {
	return &Response{
		Response: response,
	}
}

func (client ErgClient) getObservations(ctx context.Context, offsetHeight int64, limitHeight int64) ([]Message, error) {

	type Result struct {
		Success      bool      `json:"success"`
		Observations []Message `json:"VaaData"`
	}

	route := fmt.Sprintf("scanner/observations/%d/%d", offsetHeight, limitHeight)

	requestUrl, err := JoinUrl(client.Options.BaseUrl, route)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequestWithContext(ctx, "GET", requestUrl.String(), nil)
	if err != nil {
		return nil, err
	}
	var out Result
	_, err = client.Do(req, &out)
	if err != nil {
		return nil, err
	}
	if !out.Success {
		return nil, errors.New("can not get data from RPC")
	}

	return out.Observations, nil
}

func (client ErgClient) getlastHeight(ctx context.Context) (int64, error) {
	type Result struct {
		Success bool  `json:"success"`
		Height  int64 `json:"height"`
	}

	requestUrl, err := JoinUrl(client.Options.BaseUrl, "height")
	if err != nil {
		return 0, err
	}
	req, err := http.NewRequestWithContext(ctx, "GET", requestUrl.String(), nil)
	if err != nil {
		return 0, err
	}
	var out Result
	_, err = client.Do(req, &out)
	if err != nil {
		return 0, err
	}
	if !out.Success {
		return 0, errors.New("can not get data from RPC")
	}
	return out.Height, nil
}

// GetCurrentGuardianSet TODO: must be developed
func (client ErgClient) GetCurrentGuardianSet(ctx context.Context) (common.ErgoGuardianSet, error) {

	return common.ErgoGuardianSet{}, nil
}

func (client *ErgClient) Do(req *http.Request, v interface{}) (*Response, error) {
	return doHttp(client.Options, req, v)
}
func doHttp(options ErgOptions, req *http.Request, v interface{}) (*Response, error) {
	if req.Header.Get("Accept") == "" {
		req.Header.Set("Accept", "application/json")
	}
	req.Header.Set("Content-Type", "application/json")

	req.Close = true
	resp, err := options.Doer.Do(req)
	if err != nil {
		return nil, &RequestError{Err: err}
	}
	defer resp.Body.Close()

	response := newResponse(resp)
	body, _ := ioutil.ReadAll(response.Body)
	urlSlice := strings.Split(req.URL.String(), "/")
	zap.L().Sugar().Debugf("\n%s, response: %v\n", urlSlice[len(urlSlice)-1], string(body))

	if response.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(response.Body)
		return response, &RequestError{
			Err:  errors.Errorf("Invalid status code: expect 200 got %d", response.StatusCode),
			Body: string(body),
		}
	}

	if v != nil {
		if err = json.Unmarshal(body, v); err != nil {
			zap.L().Sugar().Debugf("json parse error")
			return response, &ParseError{Err: err}
		}
	}
	return response, err
}

func JoinUrl(baseRaw string, pathRaw string) (*url.URL, error) {
	baseUrl, err := url.Parse(baseRaw)
	if err != nil {
		return nil, err
	}

	pathUrl, err := url.Parse(pathRaw)
	if err != nil {
		return nil, err
	}

	baseUrl.Path = path.Join(baseUrl.Path, pathUrl.Path)

	query := baseUrl.Query()
	for k := range pathUrl.Query() {
		query.Set(k, pathUrl.Query().Get(k))
	}
	baseUrl.RawQuery = query.Encode()

	return baseUrl, nil
}
