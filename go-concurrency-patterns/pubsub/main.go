package pubsub

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func Run() {
	publisher := newPublisher()

	// Example with timeout context
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	// Create subscribers with unique IDs
	sub1, err := publisher.subscribe(ctx, "subscriber-1")
	if err != nil {
		fmt.Printf("Failed to subscribe: %v\n", err)
		return
	}

	// Start subscriber goroutine
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		for msg := range sub1 {
			fmt.Printf("[subscriber-1] Topic: %s, Content: %s\n", msg.topic, msg.content)
		}
	}()

	// Publish some messages
	publisher.publish(Message{
		topic:   "greetings",
		content: "Hello!",
	})
	time.Sleep(200 * time.Millisecond)
	publisher.publish(Message{
		topic:   "info",
		content: "Hey!",
	})
	time.Sleep(2000 * time.Millisecond) // This will trigger the timeout

	// Wait for subscriber goroutine to finish
	wg.Wait()
	publisher.close()
}
