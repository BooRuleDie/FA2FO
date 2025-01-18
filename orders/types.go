package main

import (
	"context"

	pb "github.com/BooRuleDie/Microservice-in-Go/common/api"
)

type OrderService interface {
	CreateOrder(context.Context) error
	ValidateItems(context.Context, *pb.CreateOrderRequest) error
}

type OrderStore interface {
	Create(context.Context) error
}
