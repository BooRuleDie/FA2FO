package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	pb "github.com/BooRuleDie/Microservice-in-Go/common/api"
	"github.com/BooRuleDie/Microservice-in-Go/common/broker"
	amqp "github.com/rabbitmq/amqp091-go"
	"go.opentelemetry.io/otel"
)

type consumer struct {
	service OrderService
}

func NewConsumer(service OrderService) *consumer {
	return &consumer{service: service}
}

func (c *consumer) Listen(ch *amqp.Channel) {
	// declare the queue
	q, err := ch.QueueDeclare("", true, false, true, false, nil)
	if err != nil {
		log.Fatal(err)
	}

	err = ch.QueueBind(q.Name, "", broker.OrderPaidEvent, false, nil)
	if err != nil {
		log.Fatal(err)
	}

	// autoAck -> false
	msgChan, err := ch.Consume(q.Name, "", false, false, false, false, nil)
	if err != nil {
		log.Fatal(err)
	}

	go func(msgChan <-chan amqp.Delivery) {
		for msg := range msgChan {
			log.Printf("Message recevied: %s", msg.Body)

			o := &pb.Order{}
			if err := json.Unmarshal(msg.Body, o); err != nil {
				// don't put it back to queue again
				// because it it gets and json unmarshaling error
				// it'll always get that error put it into DLQ
				msg.Nack(false, false)
				log.Printf("failed to unmarshal order: %v", err)
				continue
			}

			// extract amqp headers
			ctx := broker.ExtractAMQPHeaders(context.Background(), msg.Headers)
			tr := otel.Tracer("amqp")
			spanName := fmt.Sprintf("AMQP - consume - %s", q.Name)
			_, messageSpan := tr.Start(ctx, spanName)

			_, err := c.service.UpdateOrder(
				context.Background(),
				o,
			)
			if err != nil {
				log.Printf("failed to update order: %v", err)

				// retry with the broker.HandleRetry
				if err := broker.HandleRetry(ch, &msg); err != nil {
					log.Printf("failed to handle retry in orders: %v", err)
				}

				continue
			}

			messageSpan.AddEvent("order.updated")
			messageSpan.End()

			log.Println("order has been updated from AMQP")

			msg.Ack(false)
		}
	}(msgChan)

	forever := make(chan struct{})
	<-forever
}
