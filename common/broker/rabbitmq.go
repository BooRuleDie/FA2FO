package broker

import (
	"context"
	"fmt"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

const (
	retryHeaderName string = "x-retry-count"
	maxRetryCount   int64  = 3
	DLQ             string = "dql_main"
)

func Connect(user, pass, host, port string) (*amqp.Channel, func() error) {
	url := fmt.Sprintf("amqp://%s:%s@%s:%s", user, pass, host, port)

	conn, err := amqp.Dial(url)
	if err != nil {
		log.Fatal(err)
	}

	ch, err := conn.Channel()
	if err != nil {
		log.Fatal(err)
	}

	err = ch.ExchangeDeclare(
		OrderCreatedEvent,
		"direct",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}

	err = ch.ExchangeDeclare(
		OrderPaidEvent,
		"fanout",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}

	err = createDLQAndDLX(ch)
	if err != nil {
		log.Fatal(err)
	}


	return ch, conn.Close
}

func HandleRetry(ch *amqp.Channel, d *amqp.Delivery) error {
	if d.Headers == nil {
		d.Headers = amqp.Table{}
	}

	retryCount, ok := d.Headers[retryHeaderName].(int64)
	if !ok {
		retryCount = 0
	}
	retryCount++
	d.Headers[retryHeaderName] = retryCount

	log.Printf("retrying message %s, retry count: %d", d.Body, retryCount)

	if retryCount > maxRetryCount {
		// DLQ
		log.Printf("moving message to DLQ %s", DLQ)

		return ch.PublishWithContext(context.Background(), "", DLQ, false, false, amqp.Publishing{
			ContentType:  "application/json",
			Headers:      d.Headers,
			Body:         d.Body,
			DeliveryMode: amqp.Persistent,
		})
	}

	// sleep before resending the message
	time.Sleep(time.Second * time.Duration(retryCount))
	return ch.PublishWithContext(context.Background(), d.Exchange, d.RoutingKey, false, false, amqp.Publishing{
		ContentType:  "application/json",
		Headers:      d.Headers,
		Body:         d.Body,
		DeliveryMode: amqp.Persistent,
	})
}

func createDLQAndDLX(ch *amqp.Channel) error {
	q, err := ch.QueueDeclare(
		"main_queue",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	dlx := "dlx_main"
	err = ch.ExchangeDeclare(
		dlx,
		"fanout",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	err = ch.QueueBind(
		q.Name,
		"",
		dlx,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	// declare DLQ
	_, err = ch.QueueDeclare(
		DLQ,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	return err
}
