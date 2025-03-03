package rpc

import (
	"context"
	"log"
	"strconv"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func startClient() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "failed to connect to the RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "failed to open a channel")
	defer ch.Close()

	// Declaring a queue is idempotent which means it won't be created if it's already created.
	q, err := ch.QueueDeclare(
		"",    // name
		false, // durable
		false, // delete when unused
		true,  // exclusive
		false, // no-wait
		nil,   // arguments
	)
	failOnError(err, "failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "failed to register a consumer")

	args := getIntegerArgs()
	corrID := generateCorrelationId()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	go func() {
		for _, arg := range args {
			err := ch.PublishWithContext(
				ctx,
				"",
				rpcServerQueue,
				false,
				false,
				amqp.Publishing{
					ContentType:   "text/plain",
					CorrelationId: corrID,
					ReplyTo:       q.Name,
					Body:          []byte(strconv.Itoa(arg)),
				},
			)
			if err != nil {
				log.Fatalf("failed to publish message, error: %v\n", err)
			}
		}
	}()

	forever := make(chan struct{})
	go func() {
		log.Println("grpc client started!!")
		for d := range msgs {
			if corrID == d.CorrelationId {
				res, err := strconv.Atoi(string(d.Body))
                failOnError(err, "Failed to convert body to integer")
                log.Printf("Result received, %d\n", res)
			}
		}

	}()

	<-forever
}
