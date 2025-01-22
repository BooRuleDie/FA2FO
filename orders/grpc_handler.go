package main

import (
	"context"
	"encoding/json"
	"log"

	pb "github.com/BooRuleDie/Microservice-in-Go/common/api"
	"github.com/BooRuleDie/Microservice-in-Go/common/broker"
	amqp "github.com/rabbitmq/amqp091-go"

	"google.golang.org/grpc"
)

type grpcHandler struct {
	pb.UnimplementedOrderServiceServer
	service OrderService
	channel *amqp.Channel
}

func NewGrpcHandler(grpcServer *grpc.Server, orderService OrderService, channel *amqp.Channel) {
	handler := &grpcHandler{service: orderService, channel: channel}
	pb.RegisterOrderServiceServer(grpcServer, handler)
}

func (gh *grpcHandler) CreateOrder(ctx context.Context, p *pb.CreateOrderRequest) (*pb.Order, error) {
	o, err := gh.service.CreateOrder(ctx, p)
	if err != nil {
		return nil, err
	}

	marshalledOrder, err := json.Marshal(o)
	if err != nil {
		return nil, err
	}

	// declare the queue
	q, err := gh.channel.QueueDeclare(broker.OrderCreatedEvent, true, false, false, false, nil)
	if err != nil {
		log.Fatal(err)
	}

	// publish the message
	gh.channel.PublishWithContext(ctx, "", q.Name, false, false, amqp.Publishing{
		ContentType: "application/json",
		Body:        marshalledOrder,
		DeliveryMode: amqp.Persistent,
	})

	return o, nil
}
