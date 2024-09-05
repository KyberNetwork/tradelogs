package rpcnode

import (
	"context"

	"github.com/KyberNetwork/tradelogs/pkg/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Client struct {
	ethClient         *ethclient.Client
	fallbackETHClient *ethclient.Client
}

func NewClient(ethClient, fallbackClient *ethclient.Client) *Client {
	return &Client{
		ethClient:         ethClient,
		fallbackETHClient: fallbackClient,
	}
}

func (c *Client) FetchTraceCall(ctx context.Context, txHash string) (types.CallFrame, error) {
	var result types.CallFrame
	err := c.ethClient.Client().CallContext(ctx, &result, "debug_traceTransaction", txHash, map[string]interface{}{
		"tracer": "callTracer",
		"tracerConfig": map[string]interface{}{
			"withLog": true,
		},
	})
	if err != nil && c.fallbackETHClient != nil {
		err = c.fallbackETHClient.Client().CallContext(ctx, &result, "debug_traceTransaction", txHash, map[string]interface{}{
			"tracer": "callTracer",
			"tracerConfig": map[string]interface{}{
				"withLog": true,
			},
		})
	}
	return result, err
}
