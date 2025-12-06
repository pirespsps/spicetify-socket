package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	ClientSocket()
	//startServer()

}

func startServer() {
	http.HandleFunc("/ws", ServerSocket)
	fmt.Printf("Websocket server started on :8080 \n")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
