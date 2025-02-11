package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func example1() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	urls := []string{
		"https://dummyjson.com/products/1",
		"https://dummyjson.com/products/2",
		"https://dummyjson.com/products/3",
	}

	results := make(chan string)

	for index, url := range urls {
		fmt.Printf("[%d] goroutine started!\n", index)
		go fetchAPI(ctx, url, results)
	}

	for range urls {
		fmt.Println(<-results)
	}
	close(results)
}

func fetchAPI(ctx context.Context, url string, results chan<- string) {
	// time.Sleep(3 * time.Second)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		results <- fmt.Sprintf("Error creating request for %s: %s", url, err.Error())
		return
	}

	// this code below returns the context error if the timeout occurs. 
	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		results <- fmt.Sprintf("Error making request to %s: %s", url, err.Error())
		return
	}
	defer resp.Body.Close()

	results <- fmt.Sprintf("Response from %s: %d", url, resp.StatusCode)
}
