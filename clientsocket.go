package main

import (
	"fmt"
	"log"
	"net/url"

	"github.com/gorilla/websocket"
)

func FirstConnection() {
	u := url.URL{Scheme: "ws", Host: "localhost:8080", Path: "/ws"}
	fmt.Printf("connected to %v \n", u)

	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("connection error: ", err)
	}
	defer conn.Close()

	message := Message{
		Sender:  "qsbar",
		Message: "",
	}

	err = conn.WriteJSON(message)
	if err != nil {
		log.Fatal("write message error: ", err)
	}
}

func ClientSocket(option string) {
	u := url.URL{Scheme: "ws", Host: "localhost:8080", Path: "/ws"}

	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("connection error: ", err)
	}
	defer conn.Close()

	message := Message{
		Sender:  "qsbar",
		Message: option,
	}

	err = conn.WriteJSON(message)
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
			fmt.Printf("message received: %v \n", string(message))
		}

	}

}
