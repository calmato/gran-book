package v1

import (
	"context"

	pb "github.com/calmato/gran-book/api/hello/proto"
)

type HelloServer struct {
	pb.UnimplementedGreeterServer
}

func (s *HelloServer) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	res := &pb.HelloReply{
		Message: req.Name + " world!",
	}

	return res, nil
}
