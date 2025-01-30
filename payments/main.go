package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/BooRuleDie/Microservice-in-Go/common"
	"github.com/BooRuleDie/Microservice-in-Go/common/broker"
	"github.com/BooRuleDie/Microservice-in-Go/common/discovery"
	"github.com/BooRuleDie/Microservice-in-Go/common/discovery/consul"
	"github.com/BooRuleDie/Microservice-in-Go/payments/gateway"
	stripeProcesser "github.com/BooRuleDie/Microservice-in-Go/payments/processor/stripe"
	_ "github.com/joho/godotenv/autoload"
	"github.com/stripe/stripe-go/v81"
	"google.golang.org/grpc"
)

var (
	consulAddr           = common.EnvString("CONSUL_ADDR", "localhost:8500")
	grpcAddr             = common.EnvString("GRPC_ADDR", "localhost:3002")
	serviceName          = "payment"
	amqpUser             = common.EnvString("AMQP_USER", "guest")
	amqpPass             = common.EnvString("AMQP_PASS", "guest")
	amqpHost             = common.EnvString("AMQP_HOST", "localhost")
	amqpPort             = common.EnvString("AMQP_PORT", "5672")
	stripeKey            = common.EnvString("STRIPE_KEY", "")
	httpAddr             = common.EnvString("HTTP_ADDR", "localhost:8082")
	endpointStripeSecret = common.EnvString("ENDPOINT_STRIPE_SECRET", "")
	jaegerAddr = common.EnvString("JAEGER_ADDR", "localhost:4318")
)

func main() {
	// set global tracer
	if err := common.SetGlobalTracer(context.TODO(), serviceName, jaegerAddr); err != nil {
		log.Fatalf("failed to start global tracer: %v", err)
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
	gateway := gateway.NewGateway(registry)
	srv := NewService(processor, gateway)
	serviceWithTelemetry := NewServiceWithTelemetry(srv)
	consumer := NewConsumer(serviceWithTelemetry)
	go consumer.Listen(ch)

	// set up the http server for webhook
	mux := http.NewServeMux()
	httpServer := NewPaymentHTTPHandler(ch)
	httpServer.registerRoutes(mux)

	go func() {
		log.Printf("starting http server at %v", httpAddr)
		if err := http.ListenAndServe(httpAddr, mux); err != nil {
			log.Fatalf("failed to start payment http server. err: %v", err)
		}
	}()

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
