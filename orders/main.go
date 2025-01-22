package main

import (
	"context"
	"log"
	"net"
	"time"

	"github.com/BooRuleDie/Microservice-in-Go/common"
	"github.com/BooRuleDie/Microservice-in-Go/common/broker"
	"github.com/BooRuleDie/Microservice-in-Go/common/discovery"
	"github.com/BooRuleDie/Microservice-in-Go/common/discovery/consul"
	_ "github.com/joho/godotenv/autoload"
	"google.golang.org/grpc"
)

var (
	consulAddr  string = common.EnvString("CONSUL_ADDR", "localhost:8500")
	grpcAddr    string = common.EnvString("GRPC_ADDR", "localhost:3000")
	serviceName string = "orders"
	amqpUser    string = common.EnvString("AMQP_USER", "guest")
	amqpPass    string = common.EnvString("AMQP_PASS", "guest")
	amqpHost    string = common.EnvString("AMQP_HOST", "localhost")
	amqpPort    string = common.EnvString("AMQP_PORT", "5672")

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

	// message broker setup
	ch, close := broker.Connect(amqpUser, amqpPass, amqpHost, amqpPort)
	defer func(){
		close()
		ch.Close()
	}()


	// create the grpc server
	grpcServer := grpc.NewServer()
	l, err := net.Listen("tcp", grpcAddr)
	if err != nil {
		log.Fatalf("failed to start grpc listener. error: %v", err)
	}
	defer l.Close()

	store := NewStore()
	service := NewService(store)
	NewGrpcHandler(grpcServer, service, ch)


	log.Println("GRPC server started listening on", grpcAddr)

	if err := grpcServer.Serve(l); err != nil {
		log.Fatalf("grpc server failed to serve. error: %v", err)
	}
}
