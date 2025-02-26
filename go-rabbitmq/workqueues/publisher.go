package workqueues

import (
	"context"
	"log"
	"time"
	
	amqp "github.com/rabbitmq/amqp091-go"
)

func publish()  {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "failed to connect to the RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "failed to open a channel")
	defer ch.Close()

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

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	body := "Some Task"
	err = ch.PublishWithContext(
		ctx,
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType: "text/plain",
			Body:        []byte(body),
		},
	)
	failOnError(err, "failed to publish a message")
	log.Printf("message sent. message: %s\n", body)
}
