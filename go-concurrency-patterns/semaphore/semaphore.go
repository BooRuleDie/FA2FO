package semaphore

import (
	"fmt"
	"sync"
	"time"
)

// sc represents a semaphore channel used to limit concurrent workers
func worker(id int, sc chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("worker %d waiting...\n", id)

	// Block if buffer is full
	sc <- struct{}{}
	fmt.Printf("worker %d started processing...\n", id)

	// Simulate work by sleeping
	time.Sleep(1500 * time.Millisecond)

	// Release semaphore slot by receiving from channel,
	// allowing another worker to start
	<-sc
	fmt.Printf("worker %d is done\n", id)
}

func Run() {
	const maxConcurrentWorkers = 3
	const totalTasks = 10

	var wg sync.WaitGroup // Initialize WaitGroup
	sc := make(chan struct{}, maxConcurrentWorkers)

	for i := 0; i < totalTasks; i++ {
		wg.Add(1)
		go worker(i, sc, &wg)
	}

	wg.Wait()
	fmt.Println("All processes are completed!")
}
