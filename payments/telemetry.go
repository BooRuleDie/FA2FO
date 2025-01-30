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
	next PaymentService
}

func NewServiceWithTelemetry(service PaymentService) *serviceWithTelemetry {
	return &serviceWithTelemetry{
		next: service,
	}
}

func (s *serviceWithTelemetry) CreatePayment(ctx context.Context, o *pb.Order) (string, error) {
	span := trace.SpanFromContext(ctx)
	eventValue := fmt.Sprintf("CreatePayment: %v", o)
	span.AddEvent(eventValue)
	return s.next.CreatePayment(ctx, o)
}