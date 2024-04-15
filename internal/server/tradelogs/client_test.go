package server

import (
	"encoding/json"
	"log"
	"testing"
	"time"

	"github.com/KyberNetwork/tradelogs/pkg/storage"
	"github.com/gorilla/websocket"
	"github.com/stretchr/testify/require"
)

const (
	url1 = "wss://tradelogs.kyberengineering.io/eventlogws?maker=0x807cF9A772d5a3f9CeFBc1192e939D62f0D9bD38&event_hash=0xac75f773e3a92f1a02b12134d65e1f47f8a14eabe4eaf1e24624918e6a8b269f"
	url2 = "ws://localhost:8080/eventlogws?maker=0x807cF9A772d5a3f9CeFBc1192e939D62f0D9bD38&event_hash=0xac75f773e3a92f1a02b12134d65e1f47f8a14eabe4eaf1e24624918e6a8b269f"
)

func TestWSClient(t *testing.T) {
	t.Skip()
	for {
		conn, _, err := websocket.DefaultDialer.Dial(url2, nil)
		if err != nil {
			log.Fatal("dial:", err)
		}
		for range time.NewTicker(time.Second).C {
			if _, data, err := conn.ReadMessage(); err == nil {
				var l storage.TradeLog
				err := json.Unmarshal(data, &l)
				require.NoError(t, err)
				log.Println("receive:", l)
			} else {
				conn.Close()
				break
			}
		}
	}
}
