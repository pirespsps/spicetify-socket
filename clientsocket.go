package main

import (
	"fmt"
	"log"
	"net/url"

	"github.com/gorilla/websocket"
)

func ClientSocket() {
	u := url.URL{Scheme: "ws", Host: "localhost:8080", Path: "/ws"}

	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("connection error: ", err)
	}
	defer conn.Close()

	initialMessage := Message{
		Sender:  "qsbar",
		Message: "",
	}

	err = conn.WriteJSON(initialMessage)
	if err != nil {
		log.Fatal("write message error: ", err)
	} else {
		fmt.Printf("connected to %v \n\n", u)
	}

	var option string

	go func() {
		for {
			_, err := fmt.Scan(&option)

			if err != nil {
				fmt.Printf("error in scanning command: %v\n", err)
			}

			if option != "" {
				sendCommand(option, conn)
			}
		}
	}()

	for {

		var msg Message
		err := conn.ReadJSON(&msg)
		if err != nil {
			fmt.Print("error in server json")
		}

		if msg.Sender != "" {
			fmt.Print(msg.Message)
		}
	}

}

func sendCommand(opt string, c *websocket.Conn) {

	err := c.WriteJSON(Message{
		Sender:  "qsbar",
		Message: opt,
	})

	if err != nil {
		log.Fatal("error in sending option: ", err)
	}
}
