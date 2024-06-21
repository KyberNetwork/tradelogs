package rpcnode

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/KyberNetwork/tradelogs/pkg/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Client struct {
	httpClient *http.Client
	rpcUrl     string
	ethClient  *ethclient.Client
}

func NewClient(httpClient *http.Client, rpcUrl string) (*Client, error) {
	ethClient, err := ethclient.Dial(rpcUrl)
	if err != nil {
		return nil, err
	}
	return &Client{
		httpClient: httpClient,
		rpcUrl:     rpcUrl,
		ethClient:  ethClient,
	}, nil
}

func (c *Client) FetchTraceCall(ctx context.Context, txHash string) (types.CallFrame, error) {
	payload := map[string]interface{}{
		"method":  "debug_traceTransaction",
		"id":      1,
		"jsonrpc": "2.0",
		"params": []interface{}{
			txHash,
			map[string]interface{}{
				"tracer": "callTracer",
				"tracerConfig": map[string]interface{}{
					"withLog": true,
				},
			},
		},
	}

	b, err := json.Marshal(payload)
	if err != nil {
		return types.CallFrame{}, err
	}

	req, err := http.NewRequest("POST", c.rpcUrl, bytes.NewBuffer(b))

	if err != nil {
		return types.CallFrame{}, err
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := c.httpClient.Do(req)
	if err != nil {
		return types.CallFrame{}, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return types.CallFrame{}, fmt.Errorf("error code %d", res.StatusCode)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return types.CallFrame{}, err
	}

	var rpcResponse types.TraceCallResponse
	if err := json.Unmarshal(body, &rpcResponse); err != nil {
		return types.CallFrame{}, err
	}

	return rpcResponse.Result, err
}
