package main

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	pb "github.com/tlkhan/hello-micro/server/proto/hello"
)

const (
	address         = "localhost:50051"
	defaultFilename = "consignment.json"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewHelloServiceClient(conn)
	name := &pb.Name{Name: "josh"}
	r, err := client.SayHello(context.Background(), name)
	if err != nil {
		log.Fatalf("Could not greet: %v", err)
	}
	log.Printf("Created: %s", r.GreetingWords)
}