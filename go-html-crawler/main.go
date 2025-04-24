package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var (
	BOT_TOKEN      string
	CHAT_ID        string
	SLEEP_DURATION time.Duration = 5 * time.Second
)

func init() {
	if err := initDatabase(); err != nil {
		log.Fatalf("failed to init db, err: %v", err)
	}

	BOT_TOKEN = os.Getenv("TELEGRAM_BOT_TOKEN")
	CHAT_ID = os.Getenv("TELEGRAM_CHAT_ID")

	if BOT_TOKEN == "" {
		log.Fatalf("TELEGRAM_BOT_TOKEN environment variable not set")
	}
	if CHAT_ID == "" {
		log.Fatalf("TELEGRAM_CHAT_ID environment variable not set")
	}
}

func runTask() error {
	html, err := getOSYMPage()
	if err != nil {
		return fmt.Errorf("failed to fetch html from osym, err: %v", err)
	}

	exams, err := parseExamTable(html)
	if err != nil {
		return fmt.Errorf("failed to parse exam html, err: %v", err)
	}

	for _, exam := range exams {
		result, err := getExamByID(exam.ID)
		// if exam is new
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				createExam(&exam)

				notificationTitle := "New Exam Detected"
				formattedExam := formatExam(&exam)
				message := notificationTitle + "\n\n" + formattedExam
				fmt.Println(message)
				SendTelegramNotification(message)
				continue
			} else {
				return fmt.Errorf("unknown error received from getExamByID, err: %v", err)
			}
		}

		// if result has been updated
		if exam.ID != result.ID ||
			exam.Date != result.Date ||
			exam.Type != result.Type ||
			exam.Status != result.Status ||
			exam.ProcessDates != result.ProcessDates {

			updateExam(&exam)

			notificationTitle := "Exam Updated"
			formattedExam := formatExam(&exam)
			message := notificationTitle + "\n\n" + formattedExam
			fmt.Println(message)
			SendTelegramNotification(message)
			continue
		}

		// if exam is not new and it's not updated
		// not we have already been stored that exam
		// no need to do anything
	}
	
	// sleep for 5 seconds
	fmt.Printf("Sleeping for %v seconds...\n", SLEEP_DURATION)
	time.Sleep(SLEEP_DURATION)

	return nil

}

func main() {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)

	for {
		select {
		case <-signalChan:
			fmt.Println("\n🛑 CTRL+C received. Shutting down.")
			return
		default:
			fmt.Println("Running the task...")
			if err := runTask(); err != nil {
				fmt.Println("Error running task:", err)
			}
		}
	}
}
