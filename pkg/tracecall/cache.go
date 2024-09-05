package tracecall

import (
	"context"
	"fmt"
	"github.com/KyberNetwork/tradelogs/pkg/rpcnode"
	"github.com/KyberNetwork/tradelogs/pkg/types"
	"github.com/ethereum/go-ethereum/common/lru"
)

type Cache struct {
	rpcClient       rpcnode.Client
	fallbackClient  rpcnode.Client
	latestTraceCall lru.BasicLRU[string, types.CallFrame]
}

func NewCache(client, fallbackClient *rpcnode.Client) *Cache {
	return &Cache{
		rpcClient:       *client,
		fallbackClient:  *fallbackClient,
		latestTraceCall: lru.NewBasicLRU[string, types.CallFrame](100),
	}
}

func (c *Cache) GetTraceCall(tx string) (types.CallFrame, error) {
	data, ok := c.latestTraceCall.Get(tx)
	if ok {
		return data, nil
	}

	var err error
	data, err = c.rpcClient.FetchTraceCall(context.Background(), tx)
	if err != nil {
		data, err = c.fallbackClient.FetchTraceCall(context.Background(), tx)
		if err != nil {
			return types.CallFrame{}, err
		}
	}
	if data.From == "" {
		return types.CallFrame{}, fmt.Errorf("trace call not found")
	}
	c.latestTraceCall.Add(tx, data)
	return data, nil
}
