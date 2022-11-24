package main

import (
	pb "grpc-go/calculator/proto"
	"log"
	"net"

	"google.golang.org/grpc"
)

type Server struct {
	pb.CalculatorServiceServer
}

var (
	addr string = "0.0.0.0:50051"
)

func main() {
	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("Failed to listen on: %v\n", err)
	}

	log.Printf("Listening on %s\n", addr)

	s := grpc.NewServer()
	pb.RegisterCalculatorServiceServer(s, &Server{})

	if err = s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}
}
