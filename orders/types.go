package main

import (
	"context"

	pb "github.com/BooRuleDie/Microservice-in-Go/common/api"
)

type OrderService interface {
	CreateOrder(context.Context, *pb.CreateOrderRequest, []*pb.Items) (*pb.Order, error)
	ValidateItems(context.Context, *pb.CreateOrderRequest) ([]*pb.Items, error)
	GetOrder(context.Context, *pb.GetOrderRequest) (*pb.Order, error)
}

type OrderStore interface {
	Create(context.Context, *pb.CreateOrderRequest, []*pb.Items) (string, error)
	Get(context.Context, *pb.GetOrderRequest) (*pb.Order, error)
}
