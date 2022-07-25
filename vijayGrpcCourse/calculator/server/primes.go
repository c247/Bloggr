package main

import (
	"log"

	pb "github.com/Clement-Jean/grpc-go-course/calculator/proto"
)

func (s *Server) Primes(in *pb.PrimesRequest, stream pb.CalculatorService_PrimesServer) error {

	log.Printf("Primes function was invoked with %v\n", in)
	n := in.A
	k := int32(2)

	for n > 1 {
		if n%k == 0 {
			// return k
			stream.Send(&pb.PrimesResponse{
				Res: k,
			})
			n = n / k
		} else {
			k++
		}

	}
	return nil

}
