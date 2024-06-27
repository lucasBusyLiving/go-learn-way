package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gorilla/websocket"
)

func main() {
	addr := "localhost:8080"
	url := "ws://" + addr + "/ws"
	conn, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer conn.Close()

	for i := 0; i < 10; i++ {
		message := fmt.Sprintf("Hello %d", i)
		err := conn.WriteMessage(websocket.TextMessage, []byte(message))
		if err != nil {
			log.Println("write:", err)
			return
		}

		_, reply, err := conn.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			return
		}
		log.Printf("Received: %s", reply)
		time.Sleep(1 * time.Second)
	}
}
