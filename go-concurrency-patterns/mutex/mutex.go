package mutex

import (
	"fmt"
	"sync"
	"time"
)

type sharedData struct {
	mu      sync.Mutex   // Regular mutex - only one goroutine can access at a time
	rmu     sync.RWMutex // Read-write mutex - allows multiple readers but only one writer
	counter int
}

func (sd *sharedData) increase() {
	sd.mu.Lock() // Acquire exclusive lock - blocks all other accesses
	defer sd.mu.Unlock()
	sd.counter++
}

func (sd *sharedData) read() int {
	sd.mu.Lock() // Regular mutex blocks all access during read
	defer sd.mu.Unlock()
	return sd.counter
}

func (sd *sharedData) rIncrease() {
	sd.rmu.Lock() // Acquire write lock - blocks all other accesses
	defer sd.rmu.Unlock()
	sd.counter++
}

func (sd *sharedData) rRead() int {
	sd.rmu.RLock() // Read lock allows other readers but blocks writers
	defer sd.rmu.RUnlock()
	return sd.counter
}

func testMutex(useRWMutex bool, sd *sharedData) time.Duration {
	var wg sync.WaitGroup
	start := time.Now()
	wg.Add(100)

	// Launch 10 writers
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			if useRWMutex {
				sd.rIncrease()
			} else {
				sd.increase()
			}
		}()
	}

	// Launch 90 readers
	for i := 0; i < 90; i++ {
		go func() {
			defer wg.Done()
			if useRWMutex {
				_ = sd.rRead()
			} else {
				_ = sd.read()
			}
		}()
	}

	wg.Wait()
	return time.Since(start)
}

func Run() {
	sd := sharedData{}

	// Test regular mutex
	duration := testMutex(false, &sd)
	fmt.Printf("Regular Mutex took: %v\nFinal counter value: %d\n\n", duration, sd.counter)

	// Test RWMutex
	duration = testMutex(true, &sd)
	fmt.Printf("RWMutex took: %v\nFinal counter value: %d\n", duration, sd.counter)
}
