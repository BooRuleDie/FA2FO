package main

import "log"

func main() {
	output, err := colorAI("#112233 aqua")
	if err != nil {
		log.Printf("colorAI err: %s", err.Error())
		return
	} else {
		log.Printf("colorAI: %s", output)
	}

	output, err = sizeAI("StanDar D")
	if err != nil {
		log.Printf("sizeAI err: %s", err.Error())
		return
	} else {
		log.Printf("sizeAI: %s", output)
	}

}
