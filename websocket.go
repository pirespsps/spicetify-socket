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
		log.Fatal("wsHandler error: ", err)
	}
	defer conn.Close()

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			fmt.Printf("message error: %v", err)
			break
		}
		fmt.Printf("Received: %s \n", message)

		switch string(message) {

		case "current":
			fmt.Print("current \n")
		case "previous":
			fmt.Print("previous \n")
		case "skip":
			fmt.Print("skip \n")
		case "play":
			fmt.Print("play \n")
		default:
			fmt.Print("invalid option \n")
		}

		//return the message
		if err := conn.WriteMessage(websocket.TextMessage, message); err != nil {
			fmt.Printf("message write error: %v \n", err)
		}
	}
}

func ClientSocket(option string) {
	u := url.URL{Scheme: "ws", Host: "localhost:8080", Path: "/ws"}
	fmt.Printf("connected to %v \n", u)

	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("connection error: ", err)
	}
	defer conn.Close()

	err = conn.WriteMessage(websocket.TextMessage, []byte(option))
	if err != nil {
		log.Fatal("write message error: ", err)
	}

	for {

		_, message, err := conn.ReadMessage()

		if err != nil {
			fmt.Printf("connection message received error: %v \n", err)
			break
		}

		if message != nil {
			//do something
			fmt.Printf("message received: %v \n", message)
			break
		}

	}

}
