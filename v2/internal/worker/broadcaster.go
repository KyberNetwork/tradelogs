package worker

import (
	"context"
	"fmt"
	"sync"

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

type RegisterRequest struct {
	ID         string `form:"id"`
	EventHash  string `form:"event_hash"`
	Maker      string `form:"maker"`
	Taker      string `form:"taker"`
	TakerToken string `form:"taker_token"`
	MakerToken string `form:"maker_token"`
}

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
