package main

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
	pb "simplecomputation-service/proto"
)

const (
	address = "localhost:50052"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewComputeClient(conn)
	inputs := &pb.Input{Num: []int64{1, 2, 3, 4, 5}}
	ctx := context.Background()
	resp, err := client.Add(ctx, inputs)
	if err != nil {
		fmt.Println(err)
		fmt.Println(resp)
	}

	fmt.Println(err)
	fmt.Println(resp)

	results, err := client.GetResults(ctx, &pb.Option{})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(results)

}
