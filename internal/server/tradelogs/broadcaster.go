package server

import (
	"strings"
	"sync"
	"time"

	"github.com/KyberNetwork/tradelogs/pkg/storage"
	"github.com/gorilla/websocket"
	"github.com/rs/xid"
	"go.uber.org/zap"
)

type Con struct {
	id    string
	ws    *websocket.Conn
	param RegisterRequest
}

type Broadcaster struct {
	mu           sync.Mutex
	l            *zap.SugaredLogger
	clients      map[RegisterRequest]map[string]Con
	tradeLogChan chan storage.TradeLog
}

type RegisterRequest struct {
	EventHash       string `form:"event_hash"`
	Maker           string `form:"maker"`
	ContractAddress string `form:"contract_address"`
}

func NewBroadcaster(tradeChan chan storage.TradeLog) *Broadcaster {
	return &Broadcaster{
		l:            zap.S(),
		clients:      make(map[RegisterRequest]map[string]Con),
		tradeLogChan: tradeChan,
	}
}

func (b *Broadcaster) BroadcastLog() {
	for log := range b.tradeLogChan {
		b.writeEvent(log)
	}
}

func (b *Broadcaster) newConn(param RegisterRequest, conn *websocket.Conn) {
	id := xid.New().String()
	b.l.Infow("connected socket", "id", id)
	go func() {
		msgType, msg, err := conn.ReadMessage()
		b.l.Infow("read msg result", "id", id, "msgType", msgType, "msg", msg, "err", err)
		if err != nil {
			b.removeConn(conn, id, param)
		}
	}()
	b.addConn(conn, id, param)
}

func (b *Broadcaster) Test() {
	for range time.NewTicker(time.Second * 5).C {
		b.tradeLogChan <- storage.TradeLog{
			EventHash: "0xac75f773e3a92f1a02b12134d65e1f47f8a14eabe4eaf1e24624918e6a8b269f",
			Maker:     "0x807cF9A772d5a3f9CeFBc1192e939D62f0D9bD38",
		}
	}
}

func (b *Broadcaster) removeConn(conn *websocket.Conn, id string, param RegisterRequest) {
	b.mu.Lock()
	defer b.mu.Unlock()
	param.Maker = strings.ToLower(param.Maker)
	param.ContractAddress = strings.ToLower(param.ContractAddress)
	e, ok := b.clients[param]
	if !ok {
		return
	}
	conn.Close()
	delete(e, id)
}

func (b *Broadcaster) addConn(conn *websocket.Conn, id string, param RegisterRequest) {
	b.mu.Lock()
	defer b.mu.Unlock()
	param.Maker = strings.ToLower(param.Maker)
	param.ContractAddress = strings.ToLower(param.ContractAddress)
	cons, ok := b.clients[param]
	if !ok {
		cons = map[string]Con{}
	}
	cons[id] = Con{
		id:    id,
		ws:    conn,
		param: param,
	}
	b.clients[param] = cons
}

func (b *Broadcaster) writeEvent(log storage.TradeLog) {
	b.mu.Lock()
	defer b.mu.Unlock()
	maker := strings.ToLower(log.Maker)
	contractAddress := strings.ToLower(log.ContractAddress)

	cons := b.clients[RegisterRequest{EventHash: log.EventHash, Maker: maker}]
	for _, c := range cons {
		if err := c.ws.WriteJSON(log); err != nil {
			b.l.Errorw("error when send msg", "err", err)
		}
	}

	if len(contractAddress) == 0 {
		return
	}

	cons = b.clients[RegisterRequest{EventHash: log.EventHash, Maker: maker, ContractAddress: contractAddress}]
	for _, c := range cons {
		if err := c.ws.WriteJSON(log); err != nil {
			b.l.Errorw("error when send msg", "err", err)
		}
	}
}
