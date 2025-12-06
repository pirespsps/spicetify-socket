package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

var conns sync.Map

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type Message struct {
	Sender  string `json:"sender"`
	Message string `json:"message"`
}

func ServerSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Printf("wsHandler error: %v", err)
		return
	}
	defer conn.Close()

	for {
		_, messageJson, err := conn.ReadMessage()
		if err != nil {
			fmt.Printf("message error: %v\n", err)
			return
		}

		var message Message
		err = json.Unmarshal(messageJson, &message)
		if err != nil {
			fmt.Printf("message unmarshal error: %v\n", err)
			return
		}

		if _, ok := conns.Load(message.Sender); !ok {
			fmt.Printf("\n==================connection added: %v==================\n", message.Sender)
			conns.Store(message.Sender, conn)
			continue
		}

		redirectMsg := Message{
			Sender:  "server",
			Message: message.Message,
		}

		fmt.Printf("\nsender: %v  -- msg:%v\n", message.Sender, message.Message)

		if message.Sender == "spicetify" {

			sendMessage("qsbar", redirectMsg)

		} else {

			sendMessage("spicetify", redirectMsg)
		}
	}
}

func sendMessage(conn string, data any) {
	c, ok := conns.Load(conn)
	if !ok {
		fmt.Printf("\nsender not found: %v\n", conn)
		return
	}

	ws, ok := c.(*websocket.Conn)
	if !ok {
		fmt.Printf("type error in conn convertion: %v\n", conn)
	}

	ws.WriteJSON(data)
}
