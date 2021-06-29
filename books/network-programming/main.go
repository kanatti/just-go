package main

import (
	"fmt"
	"os"

	"github.com/kanatti/just-go/books/network-programming/tcp"
)

func main() {

	if len(os.Args) <= 2 {
		panic("Please provide a command and port")
	}

	command := os.Args[1]
	port := os.Args[2]
	if command == "server" {
		server := tcp.Server{Port: port}
		server.Run()
	} else if command == "client" {
		client := tcp.Client{Port: port}
		client.Run()
	} else {
		fmt.Printf("Unknown command %s", command)
	}
}
