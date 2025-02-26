package workqueues

import (
	"log"
	"math/rand"
	"time"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%v: %s\n", err, msg)
	}
}

func processTask() {
	delay := 500 + rand.Intn(1501) // Random delay between 500-2000ms
	time.Sleep(time.Duration(delay) * time.Millisecond)
}
