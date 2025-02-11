package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

func example5() {
	var timeout int = 3 // in seconds
	var contextTimeout time.Duration = 2 * time.Second

	// create the context
	ctx, cancel := context.WithTimeout(context.Background(), contextTimeout)
	defer cancel()

	// create the request struct
	url := fmt.Sprintf("https://httpbin.org/delay/%d", timeout)
	log.Printf("URL: %s\n", url)
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		log.Printf("failed to create the new request, err: %s", err)
		return
	}

	// make the request
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Printf("failed to make the request, err: %s", err)
		return
	}
	defer res.Body.Close()

	// processing the response ...
}
