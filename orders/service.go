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

func (s *service) GetOrder(ctx context.Context, p *pb.GetOrderRequest) (*pb.Order, error) {
	o, err := s.store.Get(ctx, p)
	if err != nil {
		return nil, err
	}
	return o, nil
}

func (s *service) CreateOrder(ctx context.Context, p *pb.CreateOrderRequest, items []*pb.Item) (*pb.Order, error) {
	id, err := s.store.Create(ctx, p, items)
	if err != nil {
		return nil, err
	}

	o := &pb.Order{
		ID: id,
		Items: items,
		Status: "pending",
		CustomerID: p.CustomerID,
		PaymentLink: "",
	}

	return o, nil
}

func (s *service) UpdateOrder(ctx context.Context, o *pb.Order) (*pb.Order, error) {
	err := s.store.Update(ctx, o.ID, o)
	if err != nil {
		return nil, err
	}
	return o, nil
}


func (s *service) ValidateItems(ctx context.Context, p *pb.CreateOrderRequest) ([]*pb.Item, error) {
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
	var itemsWithPrice []*pb.Item
	for _, i := range uniqueItems {
		itemsWithPrice = append(itemsWithPrice, &pb.Item{
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
