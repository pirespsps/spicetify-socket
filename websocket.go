package main

import (
	"net/http"

	"github.com/gorilla/websocket"
)

func Listener(cmd string) error {

	upgrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		conn, _ := upgrader.Upgrade(w, r, nil)
		defer conn.Close()

		conn.WriteJSON(map[string]string{"action": cmd})
	})

	return http.ListenAndServe(":8080", nil)
}
