package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var conns = make(map[string]*websocket.Conn)

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
		log.Fatal("wsHandler error: ", err)
	}
	defer conn.Close()

	for {
		_, messageJson, err := conn.ReadMessage()
		if err != nil {
			fmt.Printf("message error: %v", err)
			break
		}

		var message Message
		json.Unmarshal(messageJson, &message)

		conns[message.Sender] = conn

		if message.Sender != "spicetify" {

			redirectMsg := Message{
				Sender:  "server",
				Message: message.Message,
			}

			sendMessage("spicetify", redirectMsg)

		} else {
			if err := conn.WriteJSON(message); err != nil {
				fmt.Printf("message returning error: %v\n", err)
			}
		}

		//return the message
		//if err := conn.WriteJSON(message); err != nil {
		//	fmt.Printf("message write error: %v \n", err)
		//}
	}
}

func sendMessage(conn string, data any) {
	c := conns[conn]
	if c == nil {
		fmt.Print("sender not found: ", conn)
		return
	}
	c.WriteJSON(data)
}
