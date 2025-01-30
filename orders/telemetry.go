package main

import (
	"context"
	"fmt"

	pb "github.com/BooRuleDie/Microservice-in-Go/common/api"
	"go.opentelemetry.io/otel/trace"
)

type serviceWithTelemetry struct {
	// it's just named that way because
	// of the naming conventions in middlewares
	next OrderService
}

func NewServiceWithTelemetry(service OrderService) *serviceWithTelemetry {
	return &serviceWithTelemetry{
		next: service,
	}
}

func (s *serviceWithTelemetry) GetOrder(ctx context.Context, p *pb.GetOrderRequest) (*pb.Order, error) {
	span := trace.SpanFromContext(ctx)
	eventValue := fmt.Sprintf("GetOrder: %v", p)
	span.AddEvent(eventValue)
	return s.next.GetOrder(ctx, p)
}

func (s *serviceWithTelemetry) CreateOrder(ctx context.Context, p *pb.CreateOrderRequest, items []*pb.Items) (*pb.Order, error) {
	span := trace.SpanFromContext(ctx)
	eventValue := fmt.Sprintf("CreateOrder: %v, items: %v", p, items)
	span.AddEvent(eventValue)
	return s.next.CreateOrder(ctx, p, items)
}

func (s *serviceWithTelemetry) UpdateOrder(ctx context.Context, o *pb.Order) (*pb.Order, error) {
	span := trace.SpanFromContext(ctx)
	eventValue := fmt.Sprintf("UpdateOrder: %v", o)
	span.AddEvent(eventValue)
	return s.next.UpdateOrder(ctx, o)
}

func (s *serviceWithTelemetry) ValidateItems(ctx context.Context, p *pb.CreateOrderRequest) ([]*pb.Items, error) {
	span := trace.SpanFromContext(ctx)
	eventValue := fmt.Sprintf("ValidateItems: %v", p)
	span.AddEvent(eventValue)
	return s.next.ValidateItems(ctx, p)
}
