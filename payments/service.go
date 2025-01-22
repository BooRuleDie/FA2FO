package main

import (
	"context"

	pb "github.com/BooRuleDie/Microservice-in-Go/common/api"
	stripeProcesser "github.com/BooRuleDie/Microservice-in-Go/payments/processor"
)

type service struct {
	processor stripeProcesser.PaymentProcessor
}

func NewService(processor stripeProcesser.PaymentProcessor) *service {
	return &service{processor: processor}
}

func (s *service) CreatePayment(ctx context.Context, p *pb.Order) (string, error) {
	link, err := s.processor.CreatePaymentLink(p)
	if err != nil {
		return "", nil
	}
	return link, err
}
