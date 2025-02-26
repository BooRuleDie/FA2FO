package workqueues

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

func consume() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "failed to connect to the RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "failed to open a channel")
	defer ch.Close()

	// Quality of Service
	ch.Qos(
		1,     // prefetchCount
		0,     // prefetchSize
		false, // global
	)

	// Declaring a queue is idempotent which means it won't be created if it's already created.
	q, err := ch.QueueDeclare(
		"tasks", // name
		true,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		false,  // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "failed to register a consumer")

	forever := make(chan struct{})
	go func() {
		log.Println("consumer has been started!")
		for d := range msgs {
			log.Printf("received a message: %s, Processing...\n", d.Body)
			processTask() // sleeps between 500 - 2000 ms
			d.Ack(false)
			log.Println("processing completed, message ack sent!")
		}

	}()

	<-forever
}
