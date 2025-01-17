package main

import (
	"context"
	"log"

	pb "github.com/BooRuleDie/Microservice-in-Go/common/api"
	"google.golang.org/grpc"
)

type grpcHandler struct {
	pb.UnimplementedOrderServiceServer
}

func NewGrpcHandler(grpcServer *grpc.Server) {
	handler := &grpcHandler{}
	pb.RegisterOrderServiceServer(grpcServer, handler)
}

func (gh *grpcHandler) CreateOrder(ctx context.Context, p *pb.CreateOrderRequest) (*pb.Order, error) {
	log.Printf("Order received! Order: %v", p)
	o := &pb.Order{
		ID: "42",
	}

	return o, nil
}

