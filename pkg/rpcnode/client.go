package rpcnode

import (
	"context"
	"encoding/json"
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
	var result json.RawMessage
	err := c.ethClient.Client().CallContext(ctx, &result, "debug_traceTransaction", txHash, map[string]interface{}{
		"tracer": "callTracer",
		"tracerConfig": map[string]interface{}{
			"withLog": true,
		},
	})
	if err != nil {
		return types.CallFrame{}, err
	}
	var callFrame types.CallFrame
	err = json.Unmarshal(result, &callFrame)
	if err != nil {
		return types.CallFrame{}, err
	}
	return callFrame, nil
}
