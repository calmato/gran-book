package config

import (
	"net"

	v1 "github.com/calmato/gran-book/api/hello/internal/interface/grpc/v1"
	pb "github.com/calmato/gran-book/api/hello/proto"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"golang.org/x/xerrors"
	"google.golang.org/grpc"
)

type grpcServer struct {
	s   *grpc.Server
	lis net.Listener
}

// NewGRPCServer - gRPC Serverの生成
func NewGRPCServer(port string, logPath string) (*grpcServer, error) {
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		err = xerrors.Errorf("Faled to listen: %w", err)
		return nil, err
	}

	logger, err := newLogger(logPath)
	if err != nil {
		return nil, err
	}

	opts := []grpc.ServerOption{
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			grpc_prometheus.StreamServerInterceptor,
			grpc_zap.StreamServerInterceptor(logger),
		)),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_prometheus.UnaryServerInterceptor,
			grpc_zap.UnaryServerInterceptor(logger),
		)),
	}

	s := grpc.NewServer(opts...)

	pb.RegisterGreeterServer(s, &v1.HelloServer{})

	gs := &grpcServer{
		s:   s,
		lis: lis,
	}

	return gs, nil
}

func (s *grpcServer) Serve() error {
	return s.s.Serve(s.lis)
}

func (s *grpcServer) Stop() {
	s.s.GracefulStop()
}
