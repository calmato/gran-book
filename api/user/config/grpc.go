package config

import (
	"context"
	"net"
	"strings"

	v1 "github.com/calmato/gran-book/api/user/internal/interface/grpc/v1"
	pb "github.com/calmato/gran-book/api/user/proto"
	"github.com/calmato/gran-book/api/user/registry"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_ctxzap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap/ctxzap"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"go.uber.org/zap"
	"golang.org/x/xerrors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
)

type grpcServer struct {
	s   *grpc.Server
	lis net.Listener
}

func newGRPCServer(port, logPath, logLevel string, reg *registry.Registry) (*grpcServer, error) {
	opts, err := grpcServerOptions(logPath, logLevel)
	if err != nil {
		return nil, err
	}

	s := grpc.NewServer(opts...)
	pb.RegisterUserServiceServer(s, &v1.UserServer{UserApplication: reg.UserApplication})

	grpc_prometheus.Register(s)
	grpc_prometheus.EnableHandlingTimeHistogram()

	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		err = xerrors.Errorf("Faled to listen: %w", err)
		return nil, err
	}

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

/*
 * ServerOptions
 */
func grpcServerOptions(logPath, logLevel string) ([]grpc.ServerOption, error) {
	streamInterceptors, err := grpcStreamServerInterceptors()
	if err != nil {
		return nil, err
	}

	unaryInterceptors, err := grpcUnaryServerInterceptors(logPath, logLevel)
	if err != nil {
		return nil, err
	}

	opts := []grpc.ServerOption{
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(streamInterceptors...)),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(unaryInterceptors...)),
	}

	return opts, nil
}

/*
 * ServerOptions - StremServerInterceptor
 */
func grpcStreamServerInterceptors() ([]grpc.StreamServerInterceptor, error) {
	interceptors := []grpc.StreamServerInterceptor{}

	return interceptors, nil
}

/*
 * ServerOptions - UnaryServerInterceptor
 */
func grpcUnaryServerInterceptors(logPath, logLevel string) ([]grpc.UnaryServerInterceptor, error) {
	logger, err := newLogger(logPath, logLevel)
	if err != nil {
		return nil, err
	}

	opts := []grpc_zap.Option{}

	interceptors := []grpc.UnaryServerInterceptor{
		grpc_ctxtags.UnaryServerInterceptor(),
		grpc_prometheus.UnaryServerInterceptor,
		grpc_zap.UnaryServerInterceptor(logger, opts...),
		accessLogUnaryServerInterceptor(),
		grpc_recovery.UnaryServerInterceptor(),
	}

	return interceptors, nil
}

func accessLogUnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler,
	) (interface{}, error) {
		clientIP := "unknown"
		if p, ok := peer.FromContext(ctx); ok {
			clientIP = p.Addr.String()
		}

		userAgent := "unknown"
		if md, ok := metadata.FromIncomingContext(ctx); ok {
			if u, ok := md["user-agent"]; ok {
				userAgent = strings.Join(u, ",")
			}
		}

		res, err := handler(ctx, req)
		if err != nil {
			return nil, err
		}

		grpc_ctxzap.AddFields(
			ctx,
			zap.String("request.client_ip", clientIP),
			zap.String("request.user_agent", userAgent),
			zap.Reflect("request.body", req),
			zap.Reflect("response.body", res),
		)

		return res, nil
	}
}
