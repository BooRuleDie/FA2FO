package main

import (
	"context"
	"log"

	"github.com/BooRuleDie/Microservice-in-Go/common"
	pb "github.com/BooRuleDie/Microservice-in-Go/common/api"
)

type service struct {
	store OrderStore
}

func NewService(store OrderStore) *service {
	return &service{
		store: store,
	}
}

func (s *service) CreateOrder(ctx context.Context) error {
	return nil
}

func (s *service) ValidateItems(ctx context.Context, p *pb.CreateOrderRequest) error {
	for _, i := range p.Items {
		if i.ID == "" {
			return common.ErrNoID
		}

		if i.Quantity <= 0 {
			return common.ErrInvalidQuantity
		}
	}

	uniqueItems := getUniqueItems(p.Items)
	log.Println("Unique items:", uniqueItems)

	if len(uniqueItems) == 0 {
		return common.ErrNoItems
	}

	// validate stock service

	return nil
}

func getUniqueItems(items []*pb.ItemsWithQuantity) []*pb.ItemsWithQuantity {
	uniqueItems := make(map[string]*pb.ItemsWithQuantity, len(items))
	for _, item := range items {
		uniqueItems[item.ID] = item
	}

	result := make([]*pb.ItemsWithQuantity, 0, len(uniqueItems))
	for _, item := range uniqueItems {
		result = append(result, item)
	}

	return result
}
