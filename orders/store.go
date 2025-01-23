package main

import (
	"context"
	"errors"

	pb "github.com/BooRuleDie/Microservice-in-Go/common/api"
)

var inmemStore = make([]*pb.Order, 0)

type store struct {
	// mongoDB store here
}

func NewStore() *store {
	return &store{}
}

func (s *store) Create(ctx context.Context, p *pb.CreateOrderRequest, items []*pb.Items) (string, error) {
	id := "42"
	order := &pb.Order{
		ID:         id,
		CustomerID: p.CustomerID,
		Status:     "pending",
		Items:      items,
	}
	inmemStore = append(inmemStore, order)
	// fmt.Println("inmemStore: ", inmemStore)

	return id, nil
}

func (s *store) Get(ctx context.Context, p *pb.GetOrderRequest) (*pb.Order, error) {
	for _, o := range inmemStore {
		if o.ID == p.OrderID && o.CustomerID == p.CustomerID {
			return o, nil
		}
	}

	// fmt.Println("inmemStore: ", inmemStore)

	return nil, errors.New("order not found")
}
