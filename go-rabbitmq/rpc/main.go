package rpc

import (
	"log"
	"os"
)

func Run() {
	command := os.Args[1]
	switch command {
	case "rpc-client":
		startClient()
	case "rpc-server":
		startServer()
	default:
		log.Fatal("invalid argument")
	}
}
