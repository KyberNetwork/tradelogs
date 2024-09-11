package evmlistenerclient

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/KyberNetwork/evmlistener/pkg/block"
	"github.com/KyberNetwork/evmlistener/pkg/evmclient"
	"github.com/KyberNetwork/evmlistener/pkg/listener"
	"github.com/KyberNetwork/evmlistener/pkg/pubsub"
	"go.uber.org/zap"
)

func SubscribeEvent(l *zap.SugaredLogger, wsRPC, httpRPC string, maxTrackingBlock int,
	publisher pubsub.Publisher, processedBlock uint64) error {
	httpClient := &http.Client{
		Timeout:   time.Second * 15,
		Transport: &http.Transport{},
	}

	restClient, err := evmclient.Dial(httpRPC, httpClient)
	if err != nil {
		return fmt.Errorf("dial main node error: %w", err)
	}

	wsClient, err := evmclient.Dial(wsRPC, httpClient)
	if err != nil {
		return fmt.Errorf("dial ws node error: %w", err)
	}

	blockKeeper := block.NewBaseBlockKeeper(maxTrackingBlock)

	l.Infof("start subscribe from block %v", processedBlock)

	// pre-load 4 recent blocks
	const initBlockRange = 4
	blocks, err := listener.GetBlocks(context.Background(), restClient, processedBlock-initBlockRange,
		processedBlock, false, nil, nil)
	if err != nil {
		return fmt.Errorf("get recent block error: %w", err)
	}
	for _, v := range blocks {
		err = blockKeeper.Add(v)
		if err != nil {
			return fmt.Errorf("add block error: %w", err)
		}
	}

	handler := listener.NewHandler(l, "block-topic", restClient, blockKeeper, publisher)
	evl := listener.New(l, wsClient, restClient, handler, nil, 0) //TODO: add event filter

	return evl.Run(context.Background())
}
