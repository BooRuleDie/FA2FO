package db

import (
	"context"
	"database/sql"
	"time"
)

func New(addr, maxIdleTime string, maxOpenConns, maxIdleConns int) (*sql.DB, error) {
	db, err := sql.Open("postgres", addr)
	if err != nil {
		return nil, err
	}

	// parse maxIdleTime string into time.Duration
	parsedTime, err := time.ParseDuration(maxIdleTime)
	if err != nil {
		return nil, err
	}

	// update the database config
	db.SetMaxOpenConns(maxOpenConns)
	db.SetMaxIdleConns(maxIdleConns)
	db.SetConnMaxIdleTime(parsedTime)

	// if it takes more than 5 seconds to ping the
	// database, cancel the context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	return db, nil
}
