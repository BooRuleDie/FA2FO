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

func (s *service) CreateOrder(ctx context.Context, p *pb.CreateOrderRequest) (*pb.Order, error) {
	items, err := s.ValidateItems(ctx, p)
	if err != nil {
		return nil, err
	}

	o := &pb.Order{
		ID:         "42",
		Items:      items,
		CustomerID: p.CustomerID,
		Status:     "pending",
	}

	return o, nil
}

func (s *service) ValidateItems(ctx context.Context, p *pb.CreateOrderRequest) ([]*pb.Items, error) {
	for _, i := range p.Items {
		if i.ID == "" {
			return nil, common.ErrNoID
		}

		if i.Quantity <= 0 {
			return nil, common.ErrInvalidQuantity
		}
	}

	uniqueItems := getUniqueItems(p.Items)
	log.Println("Unique items:", uniqueItems)

	if len(uniqueItems) == 0 {
		return nil, common.ErrNoItems
	}

	// validate stock service

	// temp solution
	var itemsWithPrice []*pb.Items
	for _, i := range uniqueItems {
		itemsWithPrice = append(itemsWithPrice, &pb.Items{
			ID:       i.ID,
			PriceID:  "price_1QjkX803Z3EsCQWvk1md8nT5",
			Quantity: i.Quantity,
		})
	}

	return itemsWithPrice, nil
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
