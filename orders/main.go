package main

import (
	"context"
	"log"
	"net"

	"github.com/BooRuleDie/Microservice-in-Go/common"
	"google.golang.org/grpc"
)

var grpcAddr = common.EnvString("GRPC_ADDR", "localhost:3000")

func main() {
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
