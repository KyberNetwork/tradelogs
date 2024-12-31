package rpcnode

import (
	"context"
	"fmt"

	"github.com/KyberNetwork/tradelogs/pkg/types"
	"github.com/ethereum/go-ethereum/common"
	ethereum_types "github.com/ethereum/go-ethereum/core/types"
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

func (c *Client) GetTxOriginByTxHash(ctx context.Context, txHash string) (common.Address, error) {
	for i, ethClient := range c.ethClient {
		tx, _, err := ethClient.TransactionByHash(ctx, common.HexToHash(txHash))
		if err != nil {
			c.l.Errorw("get transaction by hash failed", "error", err, "txHash", txHash, "clientID", i)
			continue
		}
		sender, err := ethereum_types.Sender(ethereum_types.NewCancunSigner(tx.ChainId()), tx)
		if err != nil {
			c.l.Errorw("get sender failed", "error", err, "txHash", txHash, "clientID", i)
			continue
		}
		if sender != (common.Address{}) {
			return sender, nil
		}
	}
	return common.Address{}, fmt.Errorf("failed to get sender")
}
