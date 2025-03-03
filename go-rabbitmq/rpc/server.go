package rpc

import (
	"context"
	"log"
	"strconv"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

const rpcServerQueue = "rpc-queue"

func startServer() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "failed to connect to the RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		rpcServerQueue, // name
		false,          // durable
		false,          // delete when unused
		false,          // exclusive
		false,          // no-wait
		nil,            // arguments
	)
	failOnError(err, "failed to declare a queue")

	err = ch.Qos(
		1,     // prefetch count
		0,     // prefetch size
		false, // global
	)
	failOnError(err, "failed to set QoS")

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
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		
		log.Println("grpc server started!")
		for msg := range msgs {
			// process data
			msn, err := strconv.Atoi(string(msg.Body))
			failOnError(err, "Failed to convert body to integer")
			log.Printf("processing, %d\n", msn)
			result := fib(msn)

			err = ch.PublishWithContext(
				ctx,
				"",
				msg.ReplyTo,
				false,
				false,
				amqp.Publishing{
					ContentType:   "text/plain",
					CorrelationId: msg.CorrelationId,
					Body:          []byte(strconv.Itoa(result)),
				},
			)
			if err != nil {
				log.Fatalf("failed to publish message, error: %v\n", err)
			}
			msg.Ack(false)
		}
	}()

	<-forever
}
