package main

import (
	"context"
	"log"
	"net"
	"time"

	_ "github.com/joho/godotenv/autoload"
	"github.com/BooRuleDie/Microservice-in-Go/common"
	"github.com/BooRuleDie/Microservice-in-Go/common/discovery"
	"github.com/BooRuleDie/Microservice-in-Go/common/discovery/consul"
	"google.golang.org/grpc"
)

var (
	consulAddr  string = common.EnvString("CONSUL_ADDR", "localhost:8500")
	grpcAddr    string = common.EnvString("GRPC_ADDR", "localhost:3000")
	serviceName string = "orders"
)

func main() {
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
				log.Fatalf("failed to health check. error: %v", err)
			}
			time.Sleep(time.Second * 1)
		}
	}()
	defer registry.Deregister(ctx, instanceID, serviceName)

	// create the grpc server
	grpcServer := grpc.NewServer()
	l, err := net.Listen("tcp", grpcAddr)
	if err != nil {
		log.Fatalf("failed to start grpc listener. error: %v", err)
	}
	defer l.Close()

	store := NewStore()
	service := NewService(store)
	NewGrpcHandler(grpcServer, service)

	service.CreateOrder(context.Background())

	log.Println("GRPC server started listening on", grpcAddr)

	if err := grpcServer.Serve(l); err != nil {
		log.Fatalf("grpc server failed to serve. error: %v", err)
	}
}
