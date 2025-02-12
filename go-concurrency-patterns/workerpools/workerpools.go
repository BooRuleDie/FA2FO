package workerpools

import (
	"fmt"
	"sync"
	"time"
)

type job struct {
	ID       int
	Input    int
	Result   int
	WorkerID int
}

func worker(ID int, jobs <-chan job, results chan<- job, wg *sync.WaitGroup) {
	for j := range jobs {
		// simulate waiting for process to complete
		time.Sleep(100 * time.Millisecond)

		j.Result = j.Input * j.Input
		j.WorkerID = ID
		results <- j

		wg.Done()
	}
}

func Run() {
	numJobs := 10
	numWorkers := 3

	// create the waitgroup
	var wg sync.WaitGroup
	wg.Add(numJobs)

	// jobs & results channel
	jc := make(chan job, numJobs)
	rc := make(chan job, numJobs)

	// deploy workers
	for i := 0; i < numWorkers; i++ {
		go worker(i, jc, rc, &wg)
	}

	// send jobs
	go func() {
		for i := 0; i < numJobs; i++ {
			job := job{ID: i, Input: i}
			jc <- job
		}
		// closing the channel here make sense
		// because closing a channel doesn't make
		// a channel non-consumable, it just signals
		// that no more value will be sent to the channel
		close(jc)
	}()

	// close the result channel
	go func() {
		wg.Wait()
		close(rc)
	}()

	// consume results
	for r := range rc {
		fmt.Printf(
			"Job ID: %d, Input: %d, Result: %d, Worker ID: %d\n",
			r.ID, r.Input, r.Result, r.WorkerID,
		)
	}
}
