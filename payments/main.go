package main

import (
	"context"
	"log"
	"net"
	"time"
	
	stripeProcesser "github.com/BooRuleDie/Microservice-in-Go/payments/processor/stripe"
	"github.com/BooRuleDie/Microservice-in-Go/common"
	"github.com/BooRuleDie/Microservice-in-Go/common/broker"
	"github.com/BooRuleDie/Microservice-in-Go/common/discovery"
	"github.com/BooRuleDie/Microservice-in-Go/common/discovery/consul"
	_ "github.com/joho/godotenv/autoload"
	"google.golang.org/grpc"
	"github.com/stripe/stripe-go/v81"
)

var (
	consulAddr  string = common.EnvString("CONSUL_ADDR", "localhost:8500")
	grpcAddr    string = common.EnvString("GRPC_ADDR", "localhost:3002")
	serviceName string = "payment"
	amqpUser    string = common.EnvString("AMQP_USER", "guest")
	amqpPass    string = common.EnvString("AMQP_PASS", "guest")
	amqpHost    string = common.EnvString("AMQP_HOST", "localhost")
	amqpPort    string = common.EnvString("AMQP_PORT", "5672")
	stripeKey   string = common.EnvString("STRIPE_KEY", "")
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

	// setup stripe
	stripe.Key = stripeKey

	// message broker setup
	ch, close := broker.Connect(amqpUser, amqpPass, amqpHost, amqpPort)
	defer func() {
		close()
		ch.Close()
	}()

	// start message broker listener
	processor := stripeProcesser.NewStripe()
	srv := NewService(processor)
	consumer := NewConsumer(srv)
	go consumer.Listen(ch)

	// create the grpc server
	grpcServer := grpc.NewServer()
	l, err := net.Listen("tcp", grpcAddr)
	if err != nil {
		log.Fatalf("failed to start grpc listener. error: %v", err)
	}
	defer l.Close()

	log.Println("GRPC server started listening on", grpcAddr)

	if err := grpcServer.Serve(l); err != nil {
		log.Fatalf("grpc server failed to serve. error: %v", err)
	}
}
