package gateway

import (
	"context"

	pb "github.com/BooRuleDie/Microservice-in-Go/common/api"
)

type OrdersGateway interface {
	CreateOrder(context.Context, *pb.CreateOrderRequest) (*pb.Order, error)
}
