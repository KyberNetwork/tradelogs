package tracecall

import (
	"context"

	"github.com/KyberNetwork/tradelogs/internal/types"
	"github.com/KyberNetwork/tradelogs/pkg/rpcnode"
	"github.com/ethereum/go-ethereum/common/lru"
)

type Cache struct {
	rpcClient       rpcnode.Client
	latestTraceCall lru.BasicLRU[string, types.CallFrame]
}

func NewCache(client *rpcnode.Client) *Cache {
	return &Cache{
		rpcClient:       *client,
		latestTraceCall: lru.NewBasicLRU[string, types.CallFrame](100),
	}
}

func (c *Cache) GetTraceCall(tx string) (types.CallFrame, error) {
	data, ok := c.latestTraceCall.Get(tx)
	if ok {
		return data, nil
	}
	data, err := c.rpcClient.FetchTraceCall(context.Background(), tx)
	if err != nil {
		return types.CallFrame{}, err
	}
	c.latestTraceCall.Add(tx, data)
	return data, nil
}
