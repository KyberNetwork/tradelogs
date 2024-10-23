package worker

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"sync"

	"github.com/IBM/sarama"
	"github.com/KyberNetwork/tradelogs/pkg/storage"
	"github.com/KyberNetwork/tradelogs/v2/pkg/kafka"
	"github.com/gorilla/websocket"
	"github.com/rs/xid"
	"go.uber.org/zap"
)

type Broadcaster struct {
	l       *zap.SugaredLogger
	mu      sync.Mutex
	clients map[string]*Client
	config  *kafka.Config
	topic   string
}

type (
	Client struct {
		l      *zap.SugaredLogger
		id     string
		ws     *websocket.Conn
		params RegisterRequest
	}
	RegisterRequest struct {
		ID         string `form:"id"`
		EventHash  string `form:"event_hash"`
		Maker      string `form:"maker"`
		Taker      string `form:"taker"`
		TakerToken string `form:"taker_token"`
		MakerToken string `form:"maker_token"`
	}
)

func NewBroadcaster(logger *zap.SugaredLogger, cfg *kafka.Config, topic string) *Broadcaster {
	return &Broadcaster{
		l:       logger,
		config:  cfg,
		topic:   topic,
		clients: make(map[string]*Client),
	}
}

func (b *Broadcaster) NewConn(req RegisterRequest, conn *websocket.Conn) error {
	ctx, cancel := context.WithCancel(context.Background())

	// new connection
	if len(req.ID) == 0 {
		// create id for the connection
		req.ID = xid.New().String()

		// write the id
		err := conn.WriteJSON(map[string]interface{}{"id": req.ID})
		if err != nil {
			b.removeClient(cancel, conn, req.ID)
			return fmt.Errorf("write conn id err: %v", err)
		}
	}
	b.l.Infow("connected socket", "id", req.ID)

	go func() {
		msgType, msg, err := conn.ReadMessage()
		b.l.Infow("read msg result", "id", req.ID, "msgType", msgType, "msg", msg, "err", err)
		if err != nil {
			b.l.Errorw("error when read ws connection", "id", req.ID, "err", err)
			b.removeClient(cancel, conn, req.ID)
		}
	}()

	b.addClient(conn, req.ID, req)

	err := b.clients[req.ID].run(ctx, b.config, b.topic)
	if err != nil {
		b.removeClient(cancel, conn, req.ID)
		return fmt.Errorf("cannot run client: %w", err)
	}
	return nil
}

func (b *Broadcaster) removeClient(cancel context.CancelFunc, conn *websocket.Conn, id string) {
	b.mu.Lock()
	defer b.mu.Unlock()
	cancel()
	_ = conn.Close()
	delete(b.clients, id)
}

func (b *Broadcaster) addClient(conn *websocket.Conn, id string, params RegisterRequest) {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.clients[id] = &Client{
		l:      b.l.With("id", id),
		id:     id,
		ws:     conn,
		params: params,
	}
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
