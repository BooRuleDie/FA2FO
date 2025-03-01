package topic

import (
	"log"
	"os"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%v: %s\n", err, msg)
	}
}

func getLogDetails() (string, string) {
	if len(os.Args) < 4 {
		log.Fatal("not enough arguments, expected at least 4")
	}
	if os.Args[1] != "publish" {
		log.Fatal("invalid command, expected 'publish' as second argument")
	}

	logType := os.Args[2]
	logMessage := os.Args[3]

	return logType, logMessage
}

func getBindingKeys() []string {
	if len(os.Args) < 3 {
		log.Fatal("not enough arguments, expected at least 3")
	}
	if os.Args[1] != "consume" {
		log.Fatal("invalid command, expected 'consume' as second argument")
	}
	return os.Args[2:]
}
