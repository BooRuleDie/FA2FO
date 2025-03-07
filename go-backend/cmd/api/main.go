package main

import (
	"go-backend/internal/db"
	"go-backend/internal/env"
	"go-backend/internal/store"
	"log"
)

func main() {
	cfg := config{
		addr: env.MustGetString("ADDR"),
		db: dbConfig{
			addr:         env.MustGetString("DB_ADDR"),
			maxOpenConns: env.MustGetInt("DB_MAX_OPEN_CONNS"),
			maxIdleConns: env.MustGetInt("DB_MAX_IDLE_CONNS"),
			maxIdleTime:  env.MustGetString("DB_MAX_IDLE_TIME"),
		},
	}
	
	// setup the database
	db, err := db.New(
		cfg.db.addr, 
		cfg.db.maxIdleTime, 
		cfg.db.maxOpenConns, 
		cfg.db.maxIdleConns,
	)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	store := store.NewPostgreSQLStorage(db)

	app := &application{
		config: cfg,
		store:  store,
	}

	log.Fatal(app.run())
}
