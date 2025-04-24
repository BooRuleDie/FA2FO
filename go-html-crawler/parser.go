package main

import (
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func parseExamTable(html string) ([]Exam, error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return nil, err
	}

	var exams []Exam
	doc.Find("#tbl_surec tbody tr").Each(func(i int, s *goquery.Selection) {
		exam := Exam{
			ID:           s.AttrOr("data-surecid", ""),
			Type:         s.Find("td.surec-td").Text(),
			Date:         s.Find("td:nth-child(3)").Text(),
			ProcessDates: strings.TrimSpace(s.Find("td:nth-child(4)").Text()),
			Status:       strings.TrimSpace(s.Find("td:nth-child(5)").Text()),
		}
		exam.Type = strings.TrimSpace(exam.Type)
		exam.Date = strings.TrimSpace(strings.ReplaceAll(exam.Date, "\n", ""))
		exams = append(exams, exam)
	})

	if len(exams) == 0 {
		return nil, fmt.Errorf("no exam data found in table")
	}

	return exams, nil
}
