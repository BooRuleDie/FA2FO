package generator

import (
	"fmt"
	"time"
)

func gen(numbers ...int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, n := range numbers {
			// simulate processing input
			time.Sleep(500 * time.Millisecond)
			out <- n
		}
	}()

	return out
}

func Run() {
	in := gen(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	for n := range in {
		fmt.Println(n)
	}
}
