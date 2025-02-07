package main

import (
	"context"

	pb "github.com/BooRuleDie/Microservice-in-Go/common/api"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type OrderService interface {
	UpdateOrder(context.Context, *pb.Order) (*pb.Order, error)
	GetOrder(context.Context, *pb.GetOrderRequest) (*pb.Order, error)
	ValidateItems(context.Context, *pb.CreateOrderRequest) ([]*pb.Item, error)
	CreateOrder(context.Context, *pb.CreateOrderRequest, []*pb.Item) (*pb.Order, error)
}

type OrderStore interface {
	Update(ctx context.Context, ID string, o *pb.Order) error
	Create(context.Context, Order) (primitive.ObjectID, error)
	Get(ctx context.Context, id, customerID string) (*Order, error)
}

// mongoDB model
type Order struct {
	// omitempty -> any value with zero values will be omitted when marshalling
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Items       []*pb.Item         `bson:"items,omitempty"`
	Status      string             `bson:"status,omitempty"`
	CustomerID  string             `bson:"customerID,omitempty"`
	PaymentLink string             `bson:"paymentLink,omitempty"`
}

func (o *Order) toProto() *pb.Order {
	return &pb.Order{
		ID:          o.ID.Hex(),
		Status:      o.Status,
		CustomerID:  o.CustomerID,
		PaymentLink: o.PaymentLink,
	}
}
