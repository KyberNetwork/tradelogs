package rpcnode

import (
	"context"

	"github.com/KyberNetwork/tradelogs/pkg/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"go.uber.org/zap"
)

type Client struct {
	l         *zap.SugaredLogger
	ethClient []*ethclient.Client
}

func NewClient(l *zap.SugaredLogger, ethClient ...*ethclient.Client) *Client {
	return &Client{
		l:         l,
		ethClient: ethClient,
	}
}

func (c *Client) FetchTraceCall(ctx context.Context, txHash string) (types.CallFrame, error) {
	var (
		result types.CallFrame
		err    error
	)
	for i, ethClient := range c.ethClient {
		err = ethClient.Client().CallContext(ctx, &result, "debug_traceTransaction", txHash, map[string]interface{}{
			"tracer": "callTracer",
			"tracerConfig": map[string]interface{}{
				"withLog": true,
			},
		})
		if err != nil {
			c.l.Errorw("fetch trace call failed", "error", err, "txHash", txHash, "clientID", i)
			continue
		}
		return result, nil
	}
	return result, err
}
