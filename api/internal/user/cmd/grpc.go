package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"net"
	"strings"

	"github.com/calmato/gran-book/api/proto/chat"
	"github.com/calmato/gran-book/api/proto/user"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_ctxzap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap/ctxzap"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware/v2"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/recovery"
	"go.uber.org/zap"
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

func newGRPCServer(port string, logger *zap.Logger, reg *registry) (*grpcServer, error) {
	opts := grpcServerOptions(logger)

	s := grpc.NewServer(opts...)
	user.RegisterAdminServiceServer(s, reg.admin)
	user.RegisterAuthServiceServer(s, reg.auth)
	chat.RegisterChatServiceServer(s, reg.chat)
	user.RegisterUserServiceServer(s, reg.user)

	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return nil, fmt.Errorf("config: failed to listen: %w", err)
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
func grpcServerOptions(logger *zap.Logger) []grpc.ServerOption {
	streamInterceptors := grpcStreamServerInterceptors(logger)
	unaryInterceptors := grpcUnaryServerInterceptors(logger)

	opts := []grpc.ServerOption{
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(streamInterceptors...)),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(unaryInterceptors...)),
	}

	return opts
}

/*
 * ServerOptions - StremServerInterceptor
 */
func grpcStreamServerInterceptors(logger *zap.Logger) []grpc.StreamServerInterceptor {
	opts := []grpc_zap.Option{}

	interceptors := []grpc.StreamServerInterceptor{
		grpc_ctxtags.StreamServerInterceptor(),
		grpc_zap.StreamServerInterceptor(logger, opts...),
		grpc_recovery.StreamServerInterceptor(),
	}

	return interceptors
}

/*
 * ServerOptions - UnaryServerInterceptor
 */
func grpcUnaryServerInterceptors(logger *zap.Logger) []grpc.UnaryServerInterceptor {
	opts := []grpc_zap.Option{}

	interceptors := []grpc.UnaryServerInterceptor{
		grpc_ctxtags.UnaryServerInterceptor(),
		grpc_zap.UnaryServerInterceptor(logger, opts...),
		accessLogUnaryServerInterceptor(),
		grpc_recovery.UnaryServerInterceptor(),
	}

	return interceptors
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
