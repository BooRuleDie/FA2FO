package main

import (
	"context"
	"fmt"
	"sync"
)


func example3() {
	// userID -> userName
	ctx := context.WithValue(context.Background(), 123, "Foo Bar")
	
	var wg sync.WaitGroup
	wg.Add(1)
	go getUserIdFromContext(ctx, &wg)
	wg.Wait()
}	

func getUserIdFromContext(ctx context.Context, wg *sync.WaitGroup) {
	username := ctx.Value(123)
	fmt.Printf("Username of the user %d is %s", 123, username)
	wg.Done()	
}