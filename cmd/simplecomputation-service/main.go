package main

import (
	"log"
	"net"

	pb "github.com/rahulPalliyalil/simplecompuation-service/proto"
	"google.golang.org/grpc"
)

const (
	port = ":50052"
)

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterComputeServer(s, nil)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
