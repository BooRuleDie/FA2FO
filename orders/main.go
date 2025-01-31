package main

import (
	"context"
	"net"
	"time"

	"github.com/BooRuleDie/Microservice-in-Go/common"
	"github.com/BooRuleDie/Microservice-in-Go/common/broker"
	"github.com/BooRuleDie/Microservice-in-Go/common/discovery"
	"github.com/BooRuleDie/Microservice-in-Go/common/discovery/consul"
	_ "github.com/joho/godotenv/autoload"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

var (
	consulAddr  = common.EnvString("CONSUL_ADDR", "localhost:8500")
	grpcAddr    = common.EnvString("GRPC_ADDR", "localhost:3000")
	serviceName = "orders"
	amqpUser    = common.EnvString("AMQP_USER", "guest")
	amqpPass    = common.EnvString("AMQP_PASS", "guest")
	amqpHost    = common.EnvString("AMQP_HOST", "localhost")
	amqpPort    = common.EnvString("AMQP_PORT", "5672")
	jaegerAddr  = common.EnvString("JAEGER_ADDR", "localhost:4318")
)

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync() // flushes buffer, if any

	// set global tracer
	if err := common.SetGlobalTracer(context.TODO(), serviceName, jaegerAddr); err != nil {
		logger.Fatal("failed to start global tracer", zap.Error(err))
	}

	registry, err := consul.NewRegistry(consulAddr, serviceName)
	if err != nil {
		panic(err)
	}
	ctx := context.Background()
	instanceID := discovery.GenerateInstanceID(serviceName)
	if err := registry.Register(ctx, instanceID, serviceName, grpcAddr); err != nil {
		panic(err)
	}
	go func() {
		for {
			if err := registry.HealthCheck(instanceID, serviceName); err != nil {
				logger.Fatal("failed to health check", zap.Error(err))
			}
			time.Sleep(time.Second * 1)
		}
	}()
	defer registry.Deregister(ctx, instanceID, serviceName)

	// message broker setup
	ch, close := broker.Connect(amqpUser, amqpPass, amqpHost, amqpPort)
	defer func() {
		close()
		ch.Close()
	}()

	// create the grpc server
	grpcServer := grpc.NewServer()
	l, err := net.Listen("tcp", grpcAddr)
	if err != nil {
		logger.Fatal("failed to start grpc listener", zap.Error(err))
	}
	defer l.Close()

	store := NewStore()
	service := NewService(store)
	// middlewares
	serviceWithTelemetry := NewServiceWithTelemetry(service)
	serviceWithLogger := NewServiceWithLogging(serviceWithTelemetry, logger)

	NewGrpcHandler(grpcServer, serviceWithLogger, ch)

	// start up rabbitmq consumer
	consumer := NewConsumer(service)
	go consumer.Listen(ch)

	logger.Info("GRPC server started listening on", zap.String("port", grpcAddr))

	if err := grpcServer.Serve(l); err != nil {
		logger.Fatal("GRPC server failed to serve", zap.Error(err))
	}
}
