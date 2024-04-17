package server

import (
	"fmt"
	"sync"
	"time"

	"github.com/KyberNetwork/tradelogs/pkg/storage"
	"github.com/gorilla/websocket"
	"github.com/rs/xid"
	"go.uber.org/zap"
)

type Con struct {
	id        string
	ws        *websocket.Conn
	eventHash string
	maker     string
}

type Broadcaster struct {
	mu           sync.Mutex
	l            *zap.SugaredLogger
	clients      map[string]map[string]Con
	tradeLogChan chan storage.TradeLog
}

func NewBroadcaster(tradeChan chan storage.TradeLog) *Broadcaster {
	return &Broadcaster{
		l:            zap.S(),
		clients:      make(map[string]map[string]Con),
		tradeLogChan: tradeChan,
	}
}

func (b *Broadcaster) BroadcastLog() {
	for log := range b.tradeLogChan {
		b.mu.Lock()
		cons := b.clients[combine(log.EventHash, log.Maker)]
		for _, c := range cons {
			if err := c.ws.WriteJSON(log); err != nil {
				b.l.Errorw("error when send msg", "err", err)
			}
		}
		b.mu.Unlock()
	}
}

func (b *Broadcaster) addConn(event, maker string, conn *websocket.Conn) {
	id := xid.New().String()
	b.l.Infow("connected socket", "id", id)
	go func() {
		msgType, msg, err := conn.ReadMessage()
		b.l.Infow("read msg result", "id", id, "msgType", msgType, "msg", msg, "err", err)
		if err != nil {
			b.mu.Lock()
			e, ok := b.clients[fmt.Sprintf("%s-%s", event, maker)]
			if !ok {
				return
			}
			conn.Close()
			delete(e, id)
			b.mu.Unlock()
		}
	}()
	b.mu.Lock()
	cons, ok := b.clients[combine(event, maker)]
	if !ok {
		cons = map[string]Con{}
	}
	cons[id] = Con{
		id:        id,
		ws:        conn,
		maker:     maker,
		eventHash: event,
	}
	b.clients[combine(event, maker)] = cons
	b.mu.Unlock()
}

func combine(event, maker string) string {
	return fmt.Sprintf("%s-%s", event, maker)
}

func (b *Broadcaster) Test() {
	for range time.NewTicker(time.Second * 5).C {
		b.tradeLogChan <- storage.TradeLog{
			EventHash: "0xac75f773e3a92f1a02b12134d65e1f47f8a14eabe4eaf1e24624918e6a8b269f",
			Maker:     "0x807cF9A772d5a3f9CeFBc1192e939D62f0D9bD38",
		}
	}
}
