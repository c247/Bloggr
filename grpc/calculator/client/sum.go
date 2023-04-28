package main

import (
	"context"
	"log"

	pb "github.com/Clement-Jean/grpc-go-course/calculator/proto"
)

func doSum(c pb.CalculatorServiceClient) {
	log.Println("doSum was invoked")
	res, err := c.Sum(context.Background(), &pb.SumRequest{
		A: 1,
		B: 2,
	})

	if err != nil {
		log.Fatalf("Could nott sum: %v\n", err)
	}
	log.Printf("Sum: %d\n", res.Sum)
}
