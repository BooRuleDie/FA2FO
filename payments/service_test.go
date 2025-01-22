package main

import (
	"context"
	"testing"

	pb "github.com/BooRuleDie/Microservice-in-Go/common/api"
	"github.com/BooRuleDie/Microservice-in-Go/payments/processor/inmem"
)

func TestService(t *testing.T) {
	processor := inmem.NewInmem()
	srv := NewService(processor)

	t.Run("should create a payment link", func(t *testing.T) {
		link, err := srv.CreatePayment(context.Background(), &pb.Order{})
		if err != nil {
			t.Errorf("err is not nil. err: %v", err)
		}

		if link == "" {
			t.Error("link is empty")
		}
	})
}
