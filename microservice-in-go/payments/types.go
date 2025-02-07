package main

import (
	"context"

	pb "github.com/BooRuleDie/Microservice-in-Go/common/api"
)

type PaymentService interface {
	CreatePayment(context.Context, *pb.Order) (string, error)
}
