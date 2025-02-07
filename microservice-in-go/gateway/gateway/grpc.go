package gateway

import (
	"context"
	"log"

	pb "github.com/BooRuleDie/Microservice-in-Go/common/api"
	"github.com/BooRuleDie/Microservice-in-Go/common/discovery"
)

type gateway struct {
	registry discovery.Registry
}

func NewGateway(registry discovery.Registry) *gateway {
	return &gateway{
		registry: registry,
	}
}

func (g *gateway) CreateOrder(ctx context.Context, p *pb.CreateOrderRequest) (*pb.Order, error) {
	conn, err := discovery.ServiceConnection(ctx, "orders", g.registry)
	if err != nil {
		log.Fatalf("failed to create the order. error: %v", err)
	}

	c := pb.NewOrderServiceClient(conn)

	o, err := c.CreateOrder(ctx, p)
	return o, err
}

func (g *gateway) GetOrder(ctx context.Context, p *pb.GetOrderRequest) (*pb.Order, error) {
	conn, err := discovery.ServiceConnection(ctx, "orders", g.registry)
	if err != nil {
		log.Fatalf("failed to create the order. error: %v", err)
	}

	c := pb.NewOrderServiceClient(conn)

	o, err := c.GetOrder(ctx, p)
	return o, err
}
