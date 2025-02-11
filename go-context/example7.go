package main

import (
	"context"
	"errors"
	"log"
	"time"
)

func example7(){
	ctx, cancel := context.WithTimeoutCause(
		context.Background(), 
		2 * time.Second, 
		errors.New("some verbose custom timeout error becuase default error sucks"),
	)
	defer cancel()
	
	select {
		case <-ctx.Done():
			log.Printf(
				"context is cancelled\nerror: %v\ncause: %v", 
				ctx.Err(), 
				context.Cause(ctx),
			)
		case <-time.After(3 * time.Second):
			log.Println("task completed!")
	}
}