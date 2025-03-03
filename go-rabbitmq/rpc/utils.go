package rpc

import (
	"log"
	"os"
	"strconv"
)

var memo = make(map[int]int)

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%v: %s\n", err, msg)
	}
}

func fib(n int) int {
	if n <= 1 {
		return n
	}
	if val, exists := memo[n]; exists {
		return val
	}
	memo[n] = fib(n-1) + fib(n-2)
	return memo[n]
}

func getIntegerArgs() []int {
	var values []int
	for _, arg := range os.Args[2:] {
		val, err := strconv.Atoi(arg)
		if err != nil {
			log.Fatalln("all values after the 'rpc-client' must be an integer")
		}
		values = append(values, val)
	}
	return values
}

func generateCorrelationId() string {
 return strconv.FormatInt(int64(os.Getpid()), 10) + 
  "-" + 
  strconv.FormatInt(int64(os.Getpid()), 16)
}
