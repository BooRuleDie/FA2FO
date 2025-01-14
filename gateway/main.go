package main

import (
	"log"
	"net/http"

	_ "github.com/joho/godotenv/autoload"
	common "github.com/BooRuleDie/Microservice-in-Go/common"
)

var addr string = common.EnvString("HTTP_ADDR", ":8888")

func main()  {
	mux := http.NewServeMux()
	handler := NewHandler()
	handler.registerRoutes(mux)

	log.Printf("server started listening on localhost%s", addr)

	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatal("failed to start the server")
	}

}