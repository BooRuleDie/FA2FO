package main

import (
	"context"

	pb "github.com/BooRuleDie/Microservice-in-Go/common/api"
	"github.com/BooRuleDie/Microservice-in-Go/payments/gateway"
	stripeProcesser "github.com/BooRuleDie/Microservice-in-Go/payments/processor"
)

type service struct {
	processor stripeProcesser.PaymentProcessor
	gateway   gateway.OrdersGateway
}

func NewService(processor stripeProcesser.PaymentProcessor, gateway gateway.OrdersGateway) *service {
	return &service{processor: processor, gateway: gateway}
}

func (s *service) CreatePayment(ctx context.Context, p *pb.Order) (string, error) {
	link, err := s.processor.CreatePaymentLink(p)
	if err != nil {
		return "", nil
	}

	// update order with the link
	err = s.gateway.UpdateOrderAfterPaymentLink(ctx, p.ID, link)
	if err != nil {
		return "", err
	}

	return link, err
}
