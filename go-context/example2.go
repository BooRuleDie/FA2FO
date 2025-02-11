package main

import (
    "context"
    "fmt"
    "time"
)

func example2() {
    ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
    defer cancel()

    result := performTask(ctx, 2)

    start := time.Now()
    select {
    case <-ctx.Done():
        fmt.Println("Task timed out")
    case _, ok := <-result:
        if ok {
            fmt.Println("Task completed successfully")
        }
    }
    took := time.Since(start)
    fmt.Printf("it took %s seconds", took)
}

func performTask(ctx context.Context, processTime time.Duration) chan struct{} {
    result := make(chan struct{}, 1)
    
    go func() {
        defer close(result)
        
        select {
        case <-time.After(processTime * time.Second):
            result <- struct{}{}
        case <-ctx.Done():
            return
        }
    }()
    
    return result
}