package main

import (
	"context"
	"fmt"
	"log"
	"net"

	// Import the generated protobuf code
	pb "github.com/tlkhan/hello-micro/server/proto/hello"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50051"
)

type service struct {
}

// CreateConsignment - we created just one method on our service,
// which is a create method, which takes a context and a request as an
// argument, these are handled by the gRPC server.
func (s *service) SayHello(ctx context.Context, name *pb.Name) (*pb.Response, error) {
	fmt.Printf("hello %s", name.Name)
	return &pb.Response{GreetingWords: name.Name}, nil
}

func main() {
	// Set-up our gRPC server.
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	// Register our service with the gRPC server, this will tie our
	// implementation into the auto-generated interface code for our
	// protobuf definition.
	pb.RegisterHelloServiceServer(s, &service{})

	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
