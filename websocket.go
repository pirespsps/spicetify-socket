package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func WsHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Print("wsHandler error")
	}
	defer conn.Close()

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			fmt.Print("message error")
			break
		}
		fmt.Printf("Received: %s\\n", message)

		if err := conn.WriteMessage(websocket.TextMessage, message); err != nil {
			fmt.Print("message write error")
			break
		}
	}
}
