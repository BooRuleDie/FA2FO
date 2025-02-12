package pipeline

import (
	"fmt"
	"math"
)

// stage 1: generate numbers
func Gen(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, n := range nums {
			out <- n
		}
	}()

	return out
}

// stage 2: filter out even numbers, keep only odd numbers
func Filter(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range in {
			if n % 2 == 0 {
				out <- n
			}
		}
	}()

	return out
}

// stage 3: calculate cube root of remaining numbers
func Calc(in <-chan int) <-chan float64 {
	out := make(chan float64)
	go func() {
		defer close(out)
		for n := range in {
			result := math.Cbrt(float64(n))
            rounded := math.Round(result*100) / 100
            out <- rounded
		}
	}()
	return out
}

func Run() {
	c1 := Gen(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	c2 := Filter(c1)
	c3 := Calc(c2)

	for result := range c3 {
		fmt.Printf("Cube root: %.2f\n", result)
	}
}
