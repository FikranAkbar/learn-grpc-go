package main

import (
	"context"
	pb "grpc-go/calculator/proto"
	"log"
)

func (s *Server) Sum(ctx context.Context, in *pb.SumRequest) (*pb.SumResponse, error) {
	log.Printf("Sum function was invoked with %v\n", in)
	return &pb.SumResponse{
		ResultNumber: in.Number_1 + in.Number_2,
	}, nil
}
