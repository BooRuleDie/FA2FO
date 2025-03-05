package main

import (
	"go-backend/internal/env"
	"log"
)

func main() {
	cfg := config{
		addr: env.GetString("ADDR", ":8080"),
	}
	app := &application{
		config: cfg,
	}

	log.Fatal(app.run())
}
