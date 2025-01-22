package main

import (
	"context"

	pb "github.com/BooRuleDie/Microservice-in-Go/common/api"
)

type OrderService interface {
	CreateOrder(context.Context, *pb.CreateOrderRequest) (*pb.Order, error)
	ValidateItems(context.Context, *pb.CreateOrderRequest) ([]*pb.Items, error)
}

type OrderStore interface {
	Create(context.Context) error
}
