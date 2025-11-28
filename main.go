package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/ws", WsHandler)
	fmt.Printf("Websocket server started on :8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}

}
