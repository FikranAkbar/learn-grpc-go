package main

import (
	"context"
	pb "grpc-go/calculator/proto"
	"log"
)

func sumNumber(c pb.CalculatorServiceClient) {
	log.Println("sumNumber was invoked")
	res, err := c.Sum(context.Background(), &pb.SumRequest{
		Number_1: 10,
		Number_2: 3,
	})

	if err != nil {
		log.Fatalf("Could not sum: %v\n", err)
	}

	log.Printf("Sum: %s\n", res.ResultNumber)
}
