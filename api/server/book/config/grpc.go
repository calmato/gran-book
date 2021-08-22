package config

import (
	"context"
	"encoding/json"
	"fmt"
	"net"
	"strings"

	"github.com/calmato/gran-book/api/server/book/internal/interface/server"
	pb "github.com/calmato/gran-book/api/server/book/proto"
	"github.com/calmato/gran-book/api/server/book/registry"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_ctxzap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap/ctxzap"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"go.uber.org/zap"
	"golang.org/x/xerrors"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
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
	pb.RegisterBookServiceServer(s,
		server.NewBookServer(reg.BookRequestValidation, reg.BookApplication),
	)

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
		clientIP := ""
		if p, ok := peer.FromContext(ctx); ok {
			clientIP = p.Addr.String()
		}

		requestID := ""
		if md, ok := metadata.FromIncomingContext(ctx); ok {
			if id, ok := md["x-request-id"]; ok {
				requestID = strings.Join(id, ",")
			}
		}

		userAgent := ""
		if md, ok := metadata.FromIncomingContext(ctx); ok {
			if u, ok := md["user-agent"]; ok {
				userAgent = strings.Join(u, ",")
			}
		}

		res, err := handler(ctx, req)

		// Request / Response Messages
		reqParams := map[string]interface{}{}
		if p, ok := req.(proto.Message); ok {
			reqParams, _ = filterParams(p)
		}

		resParams := map[string]interface{}{}
		if p, ok := res.(proto.Message); ok {
			resParams, _ = filterParams(p)
		}

		ds := getErrorDetails(err)

		grpc_ctxzap.AddFields(
			ctx,
			zap.String("request.client_ip", clientIP),
			zap.String("request.request_id", requestID),
			zap.String("request.user_agent", userAgent),
			zap.Reflect("request.content", reqParams),
			zap.Reflect("response.content", resParams),
			zap.Reflect("error.details", ds),
		)

		return res, err
	}
}

func filterParams(pb proto.Message) (map[string]interface{}, error) {
	bs, err := protojson.Marshal(pb)
	if err != nil {
		return nil, fmt.Errorf("jsonpb serializer failed: %v", err)
	}

	bj := make(map[string]interface{})
	_ = json.Unmarshal(bs, &bj) // ignore error here.

	var toFilter []string
	for k := range bj {
		if strings.Contains(strings.ToLower(k), "password") {
			toFilter = append(toFilter, k)
		}
	}

	for _, k := range toFilter {
		bj[k] = "<FILTERED>"
	}

	return bj, nil
}

func getErrorDetails(err error) interface{} {
	if err == nil {
		return ""
	}

	s, ok := status.FromError(err)
	if !ok {
		return ""
	}

	// TODO: 配列に1つしか値入れないようにしてるからいったんこれで
	for _, detail := range s.Details() {
		switch v := detail.(type) {
		case *errdetails.LocalizedMessage:
			return v.Message
		case *errdetails.BadRequest:
			return v.FieldViolations
		}
	}

	return ""
}
