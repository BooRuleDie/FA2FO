package main

import (
	"go-backend/internal/env"
	"go-backend/internal/store"
	"log"
)

func main() {
	cfg := config{addr: env.GetString("ADDR", ":8080")}

	store := store.NewPostgreSQLStorage(nil)

	app := &application{
		config: cfg,
		store:  store,
	}

	log.Fatal(app.run())
}
