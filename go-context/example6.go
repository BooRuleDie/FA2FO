package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"
	
	_ "modernc.org/sqlite"
)

type User struct {
	UserID    int
	Firstname string
	Lastname  string
}

func example6() {
	// create the timeout context
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	// open the db connection
	db, err := sql.Open("sqlite", ":memory:")
	if err != nil {
		log.Fatalf("failed to open the database connection, err: %v", err)
	}
	defer db.Close()

	// verify database connection
	if err := db.PingContext(ctx); err != nil {
		log.Fatalf("failed to ping database: %v", err)
	}

	// create the table
	_, err = db.ExecContext(
		ctx,
		`CREATE TABLE IF NOT EXISTS users (
		    user_id INTEGER PRIMARY KEY AUTOINCREMENT,
		    firstname TEXT NOT NULL,
		    lastname TEXT NOT NULL
		);`,
	)
	if err != nil {
		log.Fatalf("failed to create users table, err: %v", err)
	}

	// insert the sample users
	_, err = db.ExecContext(
		ctx,
		`INSERT INTO users (firstname, lastname) VALUES
		    ('John', 'Doe'),
		    ('Jane', 'Smith'),
		    ('Michael', 'Johnson');`,
	)
	if err != nil {
		log.Fatalf("failed to insert sample data, err: %v", err)
	}

	// get the users
	rows, err := db.QueryContext(
		ctx,
		`SELECT
			user_id,
			firstname,
			lastname
		FROM users;`,
	)
	if err != nil {
		log.Fatalf("failed to fetch users, err: %v", err)
	}
	defer rows.Close()

	// iterate through the rows
	var users []User
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.UserID, &user.Firstname, &user.Lastname); err != nil {
			log.Fatalf("failed to scan row, err: %v", err)
		}
		users = append(users, user)
	}

	// check for errors from iterating over rows
	if err := rows.Err(); err != nil {
		log.Fatalf("error iterating rows, err: %v", err)
	}

	// print the results
	for _, user := range users {
		fmt.Printf("User ID: %d, Name: %s %s\n", user.UserID, user.Firstname, user.Lastname)
	}
}
