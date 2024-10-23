package worker

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/IBM/sarama"
	"github.com/KyberNetwork/tradelogs/pkg/storage"
	"github.com/KyberNetwork/tradelogs/v2/pkg/kafka"
	"github.com/gorilla/websocket"
	"go.uber.org/zap"
)

type Client struct {
	l      *zap.SugaredLogger
	id     string
	ws     *websocket.Conn
	params RegisterRequest
}

func (c *Client) run(ctx context.Context, cfg *kafka.Config, topic string) error {
	// kafka consumer for broadcasting trade logs
	consumer, err := kafka.NewConsumer(cfg, c.id)
	if err != nil {
		return fmt.Errorf("failed to create kafka consumer: %w", err)
	}

	tradeLogsChan := make(chan *sarama.ConsumerMessage, 100)

	go func() {
		err = consumer.Consume(ctx, c.l, topic, tradeLogsChan)
		if err != nil {
			panic(fmt.Errorf("failed to consume trade logs: %w", err))
		}
	}()

	go func() {
		for msg := range tradeLogsChan {
			c.broadcast(msg)
		}
	}()

	return nil
}

func (c *Client) broadcast(msg *sarama.ConsumerMessage) {
	var newMsg kafka.Message
	err := json.Unmarshal(msg.Value, &newMsg)
	if err != nil {
		c.l.Errorw("error when unmarshal message", "err", err, "msg", string(msg.Value))
		return
	}
	dataBytes, err := json.Marshal(newMsg.Data)
	if err != nil {
		c.l.Errorw("error when marshal message data", "err", err, "data", newMsg.Data)
		return
	}

	switch newMsg.Type {
	case kafka.MessageTypeRevert:
		var blocks []uint64
		err = json.Unmarshal(dataBytes, &blocks)
		if err != nil {
			c.l.Errorw("error when unmarshal reverted blocks", "err", err, "data", string(dataBytes))
			return
		}
		newMsg.Data = blocks
		if err = c.ws.WriteJSON(newMsg); err != nil {
			c.l.Errorw("error when send msg", "err", err)
		}
		c.l.Infow("broadcast revert message", "message", newMsg)

	case kafka.MessageTypeTradeLog:
		var tradelog storage.TradeLog
		err = json.Unmarshal(dataBytes, &tradelog)
		if err != nil {
			c.l.Errorw("error when unmarshal trade log", "err", err, "data", string(dataBytes))
			return
		}
		newMsg.Data = tradelog
		if c.match(tradelog) {
			if err = c.ws.WriteJSON(newMsg); err != nil {
				c.l.Errorw("error when send msg", "err", err)
			}
		}
		c.l.Infow("broadcast trade log message", "message", newMsg)
	}
}

func (c *Client) match(log storage.TradeLog) bool {
	params := c.params
	if len(params.EventHash) != 0 && !strings.EqualFold(params.EventHash, log.EventHash) {
		return false
	}
	if len(params.Maker) != 0 && !strings.EqualFold(params.Maker, log.Maker) {
		return false
	}
	if len(params.Taker) != 0 && !strings.EqualFold(params.Taker, log.Taker) {
		return false
	}
	if len(params.MakerToken) != 0 && !strings.EqualFold(params.MakerToken, log.MakerToken) {
		return false
	}
	if len(params.TakerToken) != 0 && !strings.EqualFold(params.TakerToken, log.TakerToken) {
		return false
	}
	return true
}
