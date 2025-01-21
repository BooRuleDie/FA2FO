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

	msgChan, err := ch.Consume(q.Name, "", true, false, false, false, nil)
	if err != nil {
		log.Fatal(err)
	}

	go func(msgChan <-chan amqp.Delivery) {
		for msg := range msgChan {
			log.Printf("Message recevied: %s", msg.Body)

			o := &pb.Order{}
			if err := json.Unmarshal(msg.Body, o); err != nil {
				log.Printf("failed to unmarshal order: %v", err)
			}

			paymentLink, err := c.service.CreatePayment(
				context.Background(),
				o,
			)
			if err != nil {
				log.Printf("failed to create payment: %v", err)
			}

			log.Printf("payment link: %v", paymentLink)
		}
	}(msgChan)

	forever := make(chan struct{})
	<-forever
}
