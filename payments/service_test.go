package main

import (
	"context"
	"testing"

	pb "github.com/BooRuleDie/Microservice-in-Go/common/api"
	inmemRegistry "github.com/BooRuleDie/Microservice-in-Go/common/discovery/inmem"
	"github.com/BooRuleDie/Microservice-in-Go/payments/processor/inmem"
	"github.com/BooRuleDie/Microservice-in-Go/payments/gateway"
)

func TestService(t *testing.T) {
	processor := inmem.NewInmem()
	registry := inmemRegistry.NewRegistry()
	gateway := gateway.NewGateway(registry)

	srv := NewService(processor, gateway)

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
