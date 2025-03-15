package main

import (
	"go-backend/internal/db"
	"go-backend/internal/env"
	"go-backend/internal/store"
)

func main() {

	// setup the database
	database, err := db.New(
		env.MustGetString("DB_ADDR"),
		env.MustGetString("DB_MAX_IDLE_TIME"),
		env.MustGetInt("DB_MAX_OPEN_CONNS"),
		env.MustGetInt("DB_MAX_IDLE_CONNS"),
	)
	if err != nil {
		panic(err)
	}
	defer database.Close()

	// create the store
	str := store.NewPostgreSQLStorage(database)

	// start the database seed
	err = db.Seed(str, database)
	if err != nil {
		panic(err)
	}
}
