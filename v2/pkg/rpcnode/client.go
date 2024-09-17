package rpcnode

import (
	"context"
	"fmt"
	"math/big"

	"github.com/KyberNetwork/tradelogs/v2/pkg/types"
	ethereumTypes "github.com/ethereum/go-ethereum/core/types"
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

func (c *Client) FetchTraceCalls(ctx context.Context, blockHash string) ([]types.TransactionCallFrame, error) {
	var (
		result []types.TransactionCallFrame
		err    error
	)
	for i, client := range c.ethClient {
		err = client.Client().CallContext(ctx, &result, "debug_traceBlockByHash", blockHash, map[string]interface{}{
			"tracer": "callTracer",
			"tracerConfig": map[string]interface{}{
				"withLog": true,
			},
		})
		if err != nil {
			c.l.Errorw("fetch trace call failed", "error", err, "blockHash", blockHash, "clientID", i)
			continue
		}
		return result, nil
	}

	return result, err
}

func (c *Client) FetchLogs(ctx context.Context, fromBlock, toBlock *big.Int, topics []string) ([]ethereumTypes.Log, error) {
	var (
		result []ethereumTypes.Log
		err    error
	)
	for i, client := range c.ethClient {
		err = client.Client().CallContext(ctx, &result, "eth_getLogs", map[string]interface{}{
			"fromBlock": fromBlock,
			"toBlock":   toBlock,
			"topics":    topics,
		})
		if err != nil {
			c.l.Errorw("fetch logs failed", "error", err, "clientID", i)
			continue
		}
		return result, nil
	}

	return result, err
}

func (c *Client) FetchLogsByBlockHash(ctx context.Context, blockHash string) ([]ethereumTypes.Log, error) {
	var (
		result []ethereumTypes.Log
		err    error
	)
	for i, client := range c.ethClient {
		err = client.Client().CallContext(ctx, &result, "eth_getLogs", map[string]interface{}{
			"blockHash": blockHash,
		})
		if err != nil {
			c.l.Errorw("fetch logs by block hash failed", "error", err, "blockHash", blockHash, "clientID", i)
			continue
		}
		return result, nil
	}

	return result, err
}

func (c *Client) GetBlockNumber(ctx context.Context) (uint64, error) {
	var (
		blockNumber uint64
		err         error
	)
	for i, client := range c.ethClient {
		blockNumber, err = client.BlockNumber(ctx)
		if err != nil {
			c.l.Errorw("get block number failed", "error", err, "clientID", i)
			continue
		}
		return blockNumber, nil
	}
	return 0, fmt.Errorf("block number not found: %w", err)
}
