package main

import (
	"fmt"
	"log"
	"net"

	"github.com/PhilippWegner/hochschule-bochum-masterarbeit/microservices-with-grpc/grpc-api/model"
	"google.golang.org/grpc"
)

const (
	grpcPort = "8080"
)

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", grpcPort))
	if err != nil {
		log.Fatalf("Failed to listen for gRPC: %v", err)
	}

	s := grpc.NewServer()

	model.RegisterModelServiceServer(s, &ModelServiceServer{})

	log.Printf("gRPC Server started on port %s", grpcPort)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to listen for gRPC: %v", err)
	}
}
