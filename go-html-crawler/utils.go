package main

import "fmt"

func formatExam(exam *Exam) string {
	return fmt.Sprintf(
		"ID: %s\nType: %s\nDate: %s\nProcessDates: %s\nStatus: %s\n",
		exam.ID, exam.Type, exam.Date, exam.ProcessDates, exam.Status,
	)
}
