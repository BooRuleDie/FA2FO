package helloworld

import "log"

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%v: %s\n", err, msg)
	}
}
