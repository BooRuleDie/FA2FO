package main

import (
	"context"
	"encoding/json"
	"log"

	pb "github.com/BooRuleDie/Microservice-in-Go/common/api"
	"github.com/BooRuleDie/Microservice-in-Go/common/broker"
	amqp "github.com/rabbitmq/amqp091-go"
)

type consumer struct {
	service PaymentService
}

func NewConsumer(service PaymentService) *consumer {
	return &consumer{service: service}
}

func (c *consumer) Listen(ch *amqp.Channel) {
	// declare the queue
	q, err := ch.QueueDeclare(broker.OrderCreatedEvent, true, false, false, false, nil)
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

			paymentLink, err := c.service.CreatePayment(
				context.Background(),
				o,
			)
			if err != nil {
				log.Printf("failed to create payment: %v", err)

				// retry with the broker.HandleRetry
				if err := broker.HandleRetry(ch, &msg); err != nil {
					log.Printf("failed to handle retry in orders: %v", err)
				}

				msg.Nack(false, false)

				continue
			}

			log.Printf("payment link: %v", paymentLink)
			msg.Ack(false)
		}
	}(msgChan)

	forever := make(chan struct{})
	<-forever
}
