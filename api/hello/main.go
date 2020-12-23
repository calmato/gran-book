package main

import (
	"context"
	"log"
	"net"

	pb "github.com/calmato/gran-book/api/hello/proto"
	"google.golang.org/grpc"
)

// Here is the gRPC server that GRPCClient talks to.
type server struct {
	pb.UnimplementedGreeterServer // embedding
}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: in.Name + " world"}, nil
}

func main() {
	// Create a listener on TCP port
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	// Create a gRPC server object
	s := grpc.NewServer()

	// Attach the Greeter service to the server
	pb.RegisterGreeterServer(s, &server{})

	// Serve gRPC server
	if err := s.Serve(lis); err != nil {
		log.Fatalln("Failed to dial server:", err)
	}
}
