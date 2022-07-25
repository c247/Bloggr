package main

import (
	"io"
	"log"

	pb "github.com/Clement-Jean/grpc-go-course/calculator/proto"
)

func (s *Server) Avg(stream pb.CalculatorService_AvgServer) error {
	log.Println("Avg function was invoked")
	res := []int32{}
	var sum int32
	sum = 0
	for {
		req, err := stream.Recv()

		if err == io.EOF {

			for _, val := range res {
				sum += val
			}
			return stream.SendAndClose(&pb.AvgResponse{
				Avg: sum / int32(len(res)),
			})
		}
		if err != nil {
			log.Fatalf("Error while reading client stream %v\n", err)
		}

		log.Printf("Recieving: %v\n", req)
		res = append(res, req.Val)
	}
}
