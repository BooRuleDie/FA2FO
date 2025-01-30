package main

import (
	"context"
	"log"
	"net/http"
	"time"

	_ "github.com/joho/godotenv/autoload"

	common "github.com/BooRuleDie/Microservice-in-Go/common"
	"github.com/BooRuleDie/Microservice-in-Go/common/discovery"
	"github.com/BooRuleDie/Microservice-in-Go/common/discovery/consul"
	"github.com/BooRuleDie/Microservice-in-Go/gateway/gateway"
)

var (
	consulAddr  = common.EnvString("CONSUL_ADDR", "localhost:8500")
	jaegerAddr  = common.EnvString("JAEGER_ADDR", "localhost:4318")
	httpAddr    = common.EnvString("HTTP_ADDR", ":8888")
	serviceName = "gateway"
)

// Hardcoded values are currently used for demonstration purposes
// to understand how gRPC communication works among microservices.
// In the future, these values will be replaced with dynamic service
// discovery mechanisms (e.g., Consul, etcd, or Kubernetes DNS)
// to avoid dependency on hardcoded addresses and improve scalability.
// var orderServiceAddr = "localhost:3000"

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
	if err := registry.Register(ctx, instanceID, serviceName, httpAddr); err != nil {
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

	ordersGateway := gateway.NewGateway(registry)
	handler := NewHandler(ordersGateway)

	mux := http.NewServeMux()
	handler.registerRoutes(mux)

	log.Printf("server started listening on localhost%s", httpAddr)

	if err := http.ListenAndServe(httpAddr, mux); err != nil {
		log.Fatal("failed to start the server")
	}

}
