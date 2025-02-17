package errgroup

import (
	"fmt"
	"math/rand/v2"

	"golang.org/x/sync/errgroup"
)

func Run() {
	var eg errgroup.Group

	for i := 0; i < 5; i++ {
		eg.Go(func() error {
			if rand.Float64() < 0.2 {
				return fmt.Errorf("oops, something went wrong")
			}
			fmt.Printf("Executing %d...\n", i)
			return nil
		})
	}

	if err := eg.Wait(); err != nil {
		fmt.Printf("eg.Wait error: %v", err)
	} else {
		fmt.Println("Everything is alright")
	}
}
