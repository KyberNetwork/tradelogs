package server

import (
	"log"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/KyberNetwork/tradinglib/pkg/httpsign"
	"github.com/gorilla/websocket"
)

func TestSubscribe(t *testing.T) {
	wsURL := os.Getenv("URL") // Replace with actual WebSocket URL
	key := os.Getenv("KEY")
	secret := os.Getenv("SECRET")

	request, err := http.NewRequest(http.MethodGet, wsURL+"?id=cscv9ubrk77vgbjftu5g", nil)
	_, err = httpsign.Sign(request, key, []byte(secret))
	if err != nil {
		t.Fatal(err)
	}

	// Establish a connection with the WebSocket server
	conn, _, err := websocket.DefaultDialer.Dial(wsURL+"?id=cscv9ubrk77vgbjftu5g", request.Header)
	if err != nil {
		log.Fatal("Error connecting to WebSocket server:", err)
	}
	defer conn.Close()

	go func() {
		ticker := time.NewTicker(time.Minute)
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				if err := conn.WriteMessage(websocket.PingMessage, nil); err != nil {
					log.Println("Ping error:", err)
					return
				}
			}
		}
	}()

	for {
		// cscv9ubrk77vgbjftu5g
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("Error reading message:", err)
			return
		}
		log.Printf("Received: %s", message)
	}
}
