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

		var msg Message
		//_, message, err := conn.ReadMessage()
		err := conn.ReadJSON(&msg)
		if err != nil {
			fmt.Print("error in server json")
		}

		if msg.Sender != "" {
			fmt.Printf("sender: %v -- message: %v\n", msg.Sender, msg.Message)
		}
	}

}
