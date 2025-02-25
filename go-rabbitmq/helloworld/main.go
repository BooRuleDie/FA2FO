package helloworld

import (
	"log"
	"os"
)

func Run() {
	command := os.Args[1]
	switch command {
	case "publish":
		publish()
	case "consume":
		consume()
	default:
		log.Println("invalid argument")
	}
}
