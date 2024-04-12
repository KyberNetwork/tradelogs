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
	ID := xid.New().String()
	b.l.Infow("connected socket", "id", ID)
	b.mu.Lock()
	cons, ok := b.clients[combine(event, maker)]
	if !ok {
		cons = map[string]Con{}
	}
	cons[ID] = Con{
		id:        ID,
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

func (b *Broadcaster) CheckDisconnect() {
	for {
		b.mu.Lock()
		for _, cons := range b.clients {
			for id, c := range cons {
				if _, _, err := c.ws.ReadMessage(); err != nil {
					if websocket.IsCloseError(err, websocket.CloseNormalClosure, websocket.CloseGoingAway, websocket.CloseNoStatusReceived) {
						b.l.Infow("socket is closed", "id", id)
						delete(cons, id)
					}
				}
			}
		}
		b.mu.Unlock()
		time.Sleep(time.Minute)
	}
}
