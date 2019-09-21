package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
	pb "simplecomputation-service/proto"
	"simplecomputation-service/service/adapters"
	"simplecomputation-service/service/infrastructure"
	"simplecomputation-service/service/usecases"
)

const (
	port = ":50052"
)

func main() {
	repo := &infrastructure.Repo{}
	interactor := usecases.NewInteractor(repo)
	processor := adapters.NewProcessor(interactor)
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterComputeServer(s, processor)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
