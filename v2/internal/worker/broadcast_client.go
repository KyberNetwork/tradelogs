package worker

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/IBM/sarama"
	"github.com/KyberNetwork/tradelogs/v2/pkg/kafka"
	"github.com/KyberNetwork/tradelogs/v2/pkg/storage/tradelogs/types"
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
	ackChan := make(chan bool)

	go func() {
		err = consumer.Consume(ctx, c.l, topic, tradeLogsChan, ackChan)
		if err != nil {
			panic(fmt.Errorf("failed to consume trade logs: %w", err))
		}
	}()

	retry := 3
	delay := time.Second * 5
	i := 0

	for msg := range tradeLogsChan {
		for i = 0; i < retry; i++ {
			err = c.broadcast(msg)
			if err != nil {
				time.Sleep(delay)
				continue
			}
			ackChan <- true
			break
		}
		if i == retry {
			ackChan <- false
			break
		}
	}
	return fmt.Errorf("broadcast trade logs to client %s failed", c.id)
}

func (c *Client) broadcast(msg *sarama.ConsumerMessage) error {
	var newMsg kafka.Message
	err := json.Unmarshal(msg.Value, &newMsg)
	if err != nil {
		c.l.Errorw("error when unmarshal message", "err", err, "msg", string(msg.Value))
		return err
	}
	dataBytes, err := json.Marshal(newMsg.Data)
	if err != nil {
		c.l.Errorw("error when marshal message data", "err", err, "data", newMsg.Data)
		return err
	}

	switch newMsg.Type {
	case kafka.MessageTypeRevert:
		var blocks []uint64
		err = json.Unmarshal(dataBytes, &blocks)
		if err != nil {
			c.l.Errorw("error when unmarshal reverted blocks", "err", err, "data", string(dataBytes))
			return err
		}
		newMsg.Data = blocks
		if err = c.ws.WriteJSON(newMsg); err != nil {
			c.l.Errorw("error when send msg", "err", err)
			return err
		}
		c.l.Infow("broadcast revert message", "message", newMsg)

	case kafka.MessageTypeTradeLog:
		var tradeLog types.TradeLog
		err = json.Unmarshal(dataBytes, &tradeLog)
		if err != nil {
			c.l.Errorw("error when unmarshal trade log", "err", err, "data", string(dataBytes))
			return err
		}
		newMsg.Data = tradeLog
		if c.match(tradeLog) {
			if err = c.ws.WriteJSON(newMsg); err != nil {
				c.l.Errorw("error when send msg", "err", err)
				return err
			}
			c.l.Infow("broadcast trade log message", "message", newMsg)
		}
	}
	return nil
}

func (c *Client) match(log types.TradeLog) bool {
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
