package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {

	var cmd string

	flag.StringVar(&cmd, "cmd", "", "Valid options: current, play, skip, previous, volume")
	flag.Parse()

	if cmd == "" {
		fmt.Print("Insert a valid option! -cmd=[option]")
		os.Exit(1)
	}

	if err := Listener(cmd); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	os.Exit(1)

}
