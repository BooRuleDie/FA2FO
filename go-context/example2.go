package main

import (
	"context"
	"fmt"
	"time"
)

func example2() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	result := performTask(ctx, 3)

	start := time.Now()
	select {
	case <-ctx.Done():
		fmt.Println("Task timed out")
	case <-result:
		fmt.Println("Task completed successfully")
	}
	took := time.Since(start)
	fmt.Printf("it took %s seconds", took)
}

func performTask(ctx context.Context, processTime time.Duration) chan struct{} {
	result := make(chan struct{})
	go func() {
		select {
		case <-time.After(processTime * time.Second):
			result <- struct{}{}
		}
	}()
	return result
}
