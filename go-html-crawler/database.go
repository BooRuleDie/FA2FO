package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func initDatabase() error {
	dbPath := os.Getenv("SQLITE_DB_PATH")
	if dbPath == "" {
		dbPath = "exams.db"
	}

	// Open database connection
	var err error
	db, err = sql.Open("sqlite3", dbPath)
	if err != nil {
		return fmt.Errorf("failed to open database: %v", err)
	}

	// Test the connection
	if err = db.Ping(); err != nil {
		return fmt.Errorf("failed to ping database: %v", err)
	}

	// Create tables if they don't exist
	if err = createTablesIfNotExist(); err != nil {
		return fmt.Errorf("failed to create tables: %v", err)
	}

	log.Printf("Database initialized successfully at %s", dbPath)
	return nil
}

func createTablesIfNotExist() error {
	createExamsTable := `
		CREATE TABLE IF NOT EXISTS exams (
			id TEXT PRIMARY KEY,
			type TEXT NOT NULL,
			date TEXT NOT NULL,
			process_dates TEXT,
			status TEXT NOT NULL
		);`

	_, err := db.Exec(createExamsTable)
	if err != nil {
		return fmt.Errorf("failed to create exams table: %v", err)
	}

	return nil
}

func getExamByID(id string) (*Exam, error) {
	query := `SELECT id, type, date, process_dates, status FROM exams WHERE id = ?`

	row := db.QueryRow(query, id)

	var exam Exam
	err := row.Scan(&exam.ID, &exam.Type, &exam.Date, &exam.ProcessDates, &exam.Status)
	if err != nil {
		return nil, err
	}

	return &exam, nil
}

func updateExam(exam *Exam) error {
	query := `
		UPDATE exams
		SET type = ?, date = ?, process_dates = ?, status = ?
		WHERE id = ?
	`

	_, err := db.Exec(query, exam.Type, exam.Date, exam.ProcessDates, exam.Status, exam.ID)
	if err != nil {
		return fmt.Errorf("failed to update exam: %v", err)
	}

	return nil
}

func createExam(exam *Exam) error {
	query := `
		INSERT INTO exams (id, type, date, process_dates, status)
		VALUES (?, ?, ?, ?, ?)
	`

	_, err := db.Exec(query, exam.ID, exam.Type, exam.Date, exam.ProcessDates, exam.Status)
	if err != nil {
		return fmt.Errorf("failed to create exam: %v", err)
	}

	return nil
}
