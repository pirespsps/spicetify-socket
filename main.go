package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

func main() {

	clientSocket()
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

func clientSocket() {

	//option := flag.String("option", "current", "select an option from [skip,previous,play,current]")

	flag.Parse()

	var opt = "current"

	//data no meio da execucao.....

	FirstConnection()
	ClientSocket(opt)

}
