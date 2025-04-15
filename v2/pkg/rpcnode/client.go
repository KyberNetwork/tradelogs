package rpcnode

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/KyberNetwork/tradelogs/v2/pkg/types"
	"github.com/KyberNetwork/tradelogs/v2/pkg/util"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	ethereumTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"go.uber.org/zap"
)

type IClient interface {
	FetchTraceCalls(ctx context.Context, blockHash string) ([]types.TransactionCallFrame, error)
	FetchLogsByBlockHash(ctx context.Context, blockHash string) ([]ethereumTypes.Log, error)
	GetBlockNumber(ctx context.Context) (uint64, error)
	BlockByNumber(ctx context.Context, blockNumber uint64) (*ethereumTypes.Block, error)
	FetchLogs(ctx context.Context, from, to uint64, address string, topics []string) ([]ethereumTypes.Log, error)
}

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
		result  []types.TransactionCallFrame
		err     error
		retries = 5
		delay   = 5 * time.Second
	)
	for attempt := 0; attempt < retries; attempt++ {
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
			// success case
			logIndexStart := 0
			for _, call := range result {
				logIndexStart = util.AssignLogIndexes(&call.CallFrame, logIndexStart)
			}
			return result, nil
		}
		// retry if having error
		if attempt < retries-1 {
			c.l.Warnw("retrying fetching trace call", "attempt", attempt+1, "blockHash", blockHash)
			time.Sleep(delay)
		}
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

func (c *Client) BlockByNumber(ctx context.Context, blockNumber uint64) (*ethereumTypes.Block, error) {
	var (
		block *ethereumTypes.Block
		err   error
	)
	for i, client := range c.ethClient {
		block, err = client.BlockByNumber(ctx, new(big.Int).SetUint64(blockNumber))
		if err != nil {
			c.l.Errorw("get block failed", "error", err, "clientID", i, "blockNumber", blockNumber)
			continue
		}
		return block, nil
	}
	return nil, fmt.Errorf("block with number %d not found: %w", blockNumber, err)
}

func (c *Client) FetchLogs(ctx context.Context, from, to uint64, address string, topics []string) ([]ethereumTypes.Log, error) {
	var (
		logs []ethereumTypes.Log
		err  error
	)
	var addresses []common.Address
	if len(address) > 0 {
		addresses = append(addresses, common.HexToAddress(address))
	}
	filter := ethereum.FilterQuery{
		FromBlock: new(big.Int).SetUint64(from),
		ToBlock:   new(big.Int).SetUint64(to),
		Addresses: addresses,
	}
	newTopics := make([]common.Hash, len(topics))
	if len(topics) > 0 {
		for i, topic := range topics {
			newTopics[i] = common.HexToHash(topic)
		}
		filter.Topics = [][]common.Hash{newTopics}
	}
	for i, client := range c.ethClient {
		logs, err = client.FilterLogs(ctx, filter)
		if err != nil {
			c.l.Errorw("get logs failed", "error", err, "clientID", i, "from", from, "to", to, "topics", topics)
			continue
		}
		return logs, nil
	}
	return nil, fmt.Errorf("error when get logs: %w", err)
}
