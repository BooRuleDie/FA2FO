package main

import (
	"context"
	"time"

	pb "github.com/BooRuleDie/Microservice-in-Go/common/api"
	"go.uber.org/zap"
)

type serviceWithLogging struct {
	// it's just named that way because
	// of the naming conventions in middlewares
	next   PaymentService
	logger *zap.Logger
}

func NewServiceWithLogging(service PaymentService, logger *zap.Logger) *serviceWithLogging {
	return &serviceWithLogging{
		next:   service,
		logger: logger,
	}
}

func (s *serviceWithLogging) CreatePayment(ctx context.Context, o *pb.Order) (string, error) {
	start := time.Now()
	defer func() {
		s.logger.Info("CreatePayment", zap.Duration("took", time.Since(start)))
	}()
	return s.next.CreatePayment(ctx, o)
}
