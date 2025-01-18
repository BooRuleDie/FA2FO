package main

import (
	"context"
	"log"

	pb "github.com/BooRuleDie/Microservice-in-Go/common/api"
	"google.golang.org/grpc"
)

type grpcHandler struct {
	pb.UnimplementedOrderServiceServer
	service OrderService
}

func NewGrpcHandler(grpcServer *grpc.Server, orderService OrderService) {
	handler := &grpcHandler{service: orderService}
	pb.RegisterOrderServiceServer(grpcServer, handler)
}

func (gh *grpcHandler) CreateOrder(ctx context.Context, p *pb.CreateOrderRequest) (*pb.Order, error) {
	err := gh.service.ValidateItems(ctx, p)
	if err != nil {
		return nil, err
	}

	log.Printf("Order received! Order: %v", p)
	o := &pb.Order{
		ID: "42",
	}

	return o, nil
}
