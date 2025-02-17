package faninfanout

import (
	"fmt"
	"sync"
	"time"
)

func gen(numbers ...int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, n := range numbers {
			out <- n
		}
	}()
	return out
}

func sq(input <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range input {
			// simulate process waiting
			time.Sleep(500 * time.Millisecond)
			out <- n * n
		}
	}()
	return out
}

func merge(cs ...<-chan int) <-chan int {
	out := make(chan int)
	var wg sync.WaitGroup
	wg.Add(len(cs))

	handleChannel := func(ch <-chan int) {
		defer wg.Done()
		for n := range ch {
			out <- n
		}
	}

	for _, c := range cs {
		go handleChannel(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func Run() {
	in := gen(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)

	// Distribute the sq work across two goroutines that both read from in.
	c1 := sq(in)
	c2 := sq(in)

	// Consume the merged output from c1 and c2.
	for n := range merge(c1, c2) {
		fmt.Println(n) // 1, 4, 9 ... 100
	}
}
