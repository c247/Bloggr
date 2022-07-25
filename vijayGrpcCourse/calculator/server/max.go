package main

import (
	"io"
	"log"

	pb "github.com/Clement-Jean/grpc-go-course/calculator/proto"
)

func (s *Server) Max(stream pb.CalculatorService_MaxServer) error {
	log.Println("Max function was invoked")
	m := int32(-100)
	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatalf("Error while reading client stream %v\n", err)
		}

		if req.Num > m {
			m = req.Num
			res := m
			err = stream.Send(&pb.MaxResponse{
				Max: res,
			})
		}

		if err != nil {
			log.Fatalf("Error while sending data to client %v\n", err)
		}
	}
}
