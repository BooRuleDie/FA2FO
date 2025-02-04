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
	next OrderService
	logger *zap.Logger
}

func NewServiceWithLogging(service OrderService, logger *zap.Logger) *serviceWithLogging {
	return &serviceWithLogging{
		next: service,
		logger: logger,
	}
}

func (s *serviceWithLogging) GetOrder(ctx context.Context, p *pb.GetOrderRequest) (*pb.Order, error) {
	start := time.Now()
	defer func(){
		s.logger.Info("GetOrder", zap.Duration("took", time.Since(start)))
	}()
	return s.next.GetOrder(ctx, p)
}

func (s *serviceWithLogging) CreateOrder(ctx context.Context, p *pb.CreateOrderRequest, items []*pb.Item) (*pb.Order, error) {
	start := time.Now()
	defer func(){
		s.logger.Info("CreateOrder", zap.Duration("took", time.Since(start)))
	}()
	return s.next.CreateOrder(ctx, p, items)
}

func (s *serviceWithLogging) UpdateOrder(ctx context.Context, o *pb.Order) (*pb.Order, error) {
	start := time.Now()
	defer func(){
		s.logger.Info("UpdateOrder", zap.Duration("took", time.Since(start)))
	}()
	return s.next.UpdateOrder(ctx, o)
}

func (s *serviceWithLogging) ValidateItems(ctx context.Context, p *pb.CreateOrderRequest) ([]*pb.Item, error) {
	start := time.Now()
	defer func(){
		s.logger.Info("ValidateItems", zap.Duration("took", time.Since(start)))
	}()
	return s.next.ValidateItems(ctx, p)
}
