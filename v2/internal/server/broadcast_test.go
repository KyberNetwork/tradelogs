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
	t.Skip("set env to run test")
	wsURL := os.Getenv("URL") // Replace with actual WebSocket URL
	key := os.Getenv("KEY")
	secret := os.Getenv("SECRET")

	request, err := http.NewRequest(http.MethodGet, wsURL+"?id=cscv9ubrk77vgbjftu5g", nil)
	if err != nil {
		t.Fatal(err)
	}
	_, err = httpsign.Sign(request, key, []byte(secret))
	if err != nil {
		t.Fatal(err)
	}

	// local test
	// request, err := http.NewRequest("GET", "ws://localhost:8082/eventlogws", nil)
	// wsURL := "ws://localhost:8082/eventlogws"

	// Establish a connection with the WebSocket server
	conn, _, err := websocket.DefaultDialer.Dial(wsURL+"?id=cscv9ubrk77vgbjftu5g", request.Header)
	if err != nil {
		t.Fatal("Error connecting to WebSocket server:", err)
	}
	defer conn.Close()

	go func() {
		ticker := time.NewTicker(time.Minute)
		defer ticker.Stop()
		for ; ; <-ticker.C {
			if err = conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				log.Println("Ping error:", err)
				return
			}
		}
	}()

	for {
		// cscv9ubrk77vgbjftu5g
		_, message, err := conn.ReadMessage()
		if err != nil {
			t.Fatal("Error reading from websocket:", err)
		}
		log.Printf("Received: %s", message)
	}
}
