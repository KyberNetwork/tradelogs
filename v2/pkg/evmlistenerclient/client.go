package evmlistenerclient

import (
	"context"
	"fmt"
	"github.com/KyberNetwork/evmlistener/pkg/block"
	"github.com/KyberNetwork/evmlistener/pkg/evmclient"
	"github.com/KyberNetwork/evmlistener/pkg/listener"
	"github.com/KyberNetwork/evmlistener/pkg/pubsub"
	"github.com/KyberNetwork/evmlistener/pkg/redis"
	"go.uber.org/zap"
	"net/http"
	"time"
)

func SubscribeEvent(l *zap.SugaredLogger, wsRPC, httpRPC string, maxTrackingBlock int, blockExpiration time.Duration,
	redisCfg redis.Config, publisher pubsub.Publisher) error {
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

	redisClient, err := redis.New(redisCfg)
	if err != nil {
		return fmt.Errorf("failed to connect to redis: %w", err)
	}
	blockKeeper := block.NewRedisBlockKeeper(l, redisClient, maxTrackingBlock, blockExpiration)

	handler := listener.NewHandler(l, "block-topic", restClient, blockKeeper, publisher)
	evl := listener.New(l, wsClient, restClient, handler, nil, 0) //TODO: add event filter

	return evl.Run(context.Background())
}
