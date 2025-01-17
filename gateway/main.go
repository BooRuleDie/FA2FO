package main

import (
	"log"
	"net/http"

	_ "github.com/joho/godotenv/autoload"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	common "github.com/BooRuleDie/Microservice-in-Go/common"
	pb "github.com/BooRuleDie/Microservice-in-Go/common/api"
)

var addr string = common.EnvString("HTTP_ADDR", ":8888")

// Hardcoded values are currently used for demonstration purposes
// to understand how gRPC communication works among microservices.
// In the future, these values will be replaced with dynamic service
// discovery mechanisms (e.g., Consul, etcd, or Kubernetes DNS)
// to avoid dependency on hardcoded addresses and improve scalability.
var orderServiceAddr = "localhost:3000"

func main() {
	conn, error := grpc.NewClient(orderServiceAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if error != nil {
		log.Fatalf("failed to create grpc client. error: %v", error)
	}
	defer conn.Close()
	log.Printf("grpc client started listening on %v", orderServiceAddr)

	c := pb.NewOrderServiceClient(conn)

	mux := http.NewServeMux()
	handler := NewHandler(c)
	handler.registerRoutes(mux)

	log.Printf("server started listening on localhost%s", addr)

	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatal("failed to start the server")
	}

}
