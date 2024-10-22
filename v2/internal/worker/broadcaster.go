package worker

import (
	"encoding/json"
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
	clients map[string]Conn
	channel chan *sarama.ConsumerMessage
}

type (
	Conn struct {
		id     string
		ws     *websocket.Conn
		params RegisterRequest
	}
	RegisterRequest struct {
		EventHash  string `form:"event_hash"`
		Maker      string `form:"maker"`
		Taker      string `form:"taker"`
		TakerToken string `form:"taker_token"`
		MakerToken string `form:"maker_token"`
	}
)

func NewBroadcaster(logger *zap.SugaredLogger, ch chan *sarama.ConsumerMessage) *Broadcaster {
	return &Broadcaster{
		l:       logger,
		channel: ch,
		clients: make(map[string]Conn),
	}
}

func (b *Broadcaster) Run() {
	for msg := range b.channel {
		b.broadcast(msg)
	}
}

func (b *Broadcaster) broadcast(msg *sarama.ConsumerMessage) {
	var newMsg kafka.Message
	err := json.Unmarshal(msg.Value, &newMsg)
	if err != nil {
		b.l.Errorw("error when unmarshal message", "err", err, "msg", string(msg.Value))
		return
	}
	dataBytes, err := json.Marshal(newMsg.Data)
	if err != nil {
		b.l.Errorw("error when marshal message data", "err", err, "data", newMsg.Data)
		return
	}

	switch newMsg.Type {
	case kafka.MessageTypeRevert:
		var blocks []uint64
		err = json.Unmarshal(dataBytes, &blocks)
		if err != nil {
			b.l.Errorw("error when unmarshal reverted blocks", "err", err, "data", string(dataBytes))
			return
		}
		b.writeRevert(blocks)
	case kafka.MessageTypeTradeLog:
		var tradelog storage.TradeLog
		err = json.Unmarshal(dataBytes, &tradelog)
		if err != nil {
			b.l.Errorw("error when unmarshal trade log", "err", err, "data", string(dataBytes))
			return
		}
		b.writeTradeLog(tradelog)
	}
}

func (b *Broadcaster) NewConn(req RegisterRequest, conn *websocket.Conn) {
	id := xid.New().String()
	b.l.Infow("connected socket", "id", id)
	go func() {
		msgType, msg, err := conn.ReadMessage()
		b.l.Infow("read msg result", "id", id, "msgType", msgType, "msg", msg, "err", err)
		if err != nil {
			b.removeConn(conn, id)
		}
	}()
	b.addConn(conn, id, req)
}

func (b *Broadcaster) removeConn(conn *websocket.Conn, id string) {
	b.mu.Lock()
	defer b.mu.Unlock()
	_ = conn.Close()
	delete(b.clients, id)
}

func (b *Broadcaster) addConn(conn *websocket.Conn, id string, params RegisterRequest) {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.clients[id] = Conn{
		id:     id,
		ws:     conn,
		params: params,
	}
}

func (b *Broadcaster) writeTradeLog(log storage.TradeLog) {
	b.mu.Lock()
	defer b.mu.Unlock()
	msg := kafka.Message{
		Type: kafka.MessageTypeTradeLog,
		Data: log,
	}
	var failCount int
	for _, conn := range b.clients {
		if match(log, conn.params) {
			if err := conn.ws.WriteJSON(msg); err != nil {
				b.l.Errorw("error when send msg", "err", err, "id", conn.id)
				failCount++
			}
		}
	}
	b.l.Infow("broadcast trade log message", "total", len(b.clients), "fail", failCount, "message", msg)
}

func match(log storage.TradeLog, params RegisterRequest) bool {
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

func (b *Broadcaster) writeRevert(blocks []uint64) {
	b.mu.Lock()
	defer b.mu.Unlock()
	msg := kafka.Message{
		Type: kafka.MessageTypeRevert,
		Data: blocks,
	}
	var failCount int
	for _, conn := range b.clients {
		if err := conn.ws.WriteJSON(msg); err != nil {
			b.l.Errorw("error when send msg", "err", err, "id", conn.id)
			failCount++
		}
	}
	b.l.Infow("broadcast revert message", "total", len(b.clients), "fail", failCount, "message", msg)
}
