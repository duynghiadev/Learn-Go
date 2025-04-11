package main

import (
	"fmt"
	"os"

	"github.com/duynghiadev/go-chat-tcp/chat"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please specify 'server' or 'client' as argument")
		fmt.Println("Usage: go run main.go [server|client]")
		os.Exit(1)
	}

	mode := os.Args[1]

	switch mode {
	case "server":
		chat.RunServer()
	case "client":
		chat.RunClient()
	default:
		fmt.Println("Invalid argument. Use 'server' or 'client'")
		os.Exit(1)
	}
}
