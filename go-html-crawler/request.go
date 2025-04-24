package main

import (
	"io"
	"net/http"
)

var (
	testURL     = "http://localhost:8080"
	originalURL = "https://ais.osym.gov.tr/"
)

func getOSYMPage() (string, error) {
	resp, err := http.Get(originalURL)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
