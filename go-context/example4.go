package main

import (
	"context"
	"fmt"
	"time"
)

func example4() {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(2*time.Second))
	defer cancel()

	go processData(ctx)

	time.Sleep(3 * time.Second)
}

func processData(ctx context.Context) {
	select {
	case <-ctx.Done():
		// context cancelled, err: context deadline exceeded
		fmt.Printf("context cancelled, err: %v", ctx.Err())
	}
}
