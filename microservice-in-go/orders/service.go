package main

import (
	"context"

	"github.com/BooRuleDie/Microservice-in-Go/common"
	pb "github.com/BooRuleDie/Microservice-in-Go/common/api"
	"github.com/BooRuleDie/Microservice-in-Go/orders/gateway"
)

type service struct {
	store   OrderStore
	gateway gateway.StockGateway
}

func NewService(store OrderStore, gateway gateway.StockGateway) *service {
	return &service{
		store:   store,
		gateway: gateway,
	}
}

func (s *service) GetOrder(ctx context.Context, p *pb.GetOrderRequest) (*pb.Order, error) {
	o, err := s.store.Get(ctx, p.OrderID, p.CustomerID)
	if err != nil {
		return nil, err
	}
	
	return o.toProto(), nil
}

func (s *service) CreateOrder(ctx context.Context, p *pb.CreateOrderRequest, items []*pb.Item) (*pb.Order, error) {
	id, err := s.store.Create(ctx, Order{
		Items:       items,
		Status:      "pending",
		CustomerID:  p.CustomerID,
		PaymentLink: "",
	})
	if err != nil {
		return nil, err
	}

	// TODO: store.CreateOrder can return the pb.Order 
	// creating almost same data twice in the service
	// doesn't make sense
	o := &pb.Order{
		ID:          id.Hex(),
		Items:       items,
		Status:      "pending",
		CustomerID:  p.CustomerID,
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
	// log.Println("Unique items:", uniqueItems)

	if len(uniqueItems) == 0 {
		return nil, common.ErrNoItems
	}

	// validate stock service
	inStock, items, err := s.gateway.CheckIfItemIsInStock(ctx, p.CustomerID, uniqueItems)
	if err != nil {
		return nil, err
	}
	if !inStock {
		return items, common.ErrOutOfStock
	}

	return items, nil
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
