package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func ServerSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Print("wsHandler error")
	}
	defer conn.Close()

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			fmt.Print("message error: ", err)
			break
		}
		fmt.Printf("Received: %s\\n", message)

		//return the message
		if err := conn.WriteMessage(websocket.TextMessage, message); err != nil {
			fmt.Print("message write error: ", err)
			break
		}
	}
}

func ClientSocket(option string) {
	u := url.URL{Scheme: "ws", Host: "localhost:8080", Path: "/ws"}
	fmt.Print("connected to", u)

	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("connection error: ", err)
	}
	defer conn.Close()

	err = conn.WriteMessage(websocket.TextMessage, []byte(option))
	if err != nil {
		log.Fatal("write message error: ", err)
	}

	_, message, err := conn.ReadMessage()
	if err != nil {
		log.Fatal("connection message received error: ", err)
	}

	fmt.Print("message: ", message)
}
