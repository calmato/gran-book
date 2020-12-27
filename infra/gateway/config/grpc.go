package config

import (
	"bytes"
	"context"
	"fmt"
	"path"

	gw "github.com/calmato/gran-book/infra/gateway/proto"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_logging "github.com/grpc-ecosystem/go-grpc-middleware/logging"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc"
)

var (
	JsonPbMarshaller grpc_logging.JsonPbMarshaler = &jsonpb.Marshaler{}
)

func registerServiceHandlers(ctx context.Context, mux *runtime.ServeMux, logPath, logLevel string) error {
	opts := grpcDialOptions(logPath, logLevel)

	err := gw.RegisterGreeterHandlerFromEndpoint(ctx, mux, *helloAPIEndpoint, opts)
	if err != nil {
		return err
	}

	return nil
}

func grpcDialOptions(logPath, logLevel string) []grpc.DialOption {
	unaryInterceptors, _ := grpcUnaryClientInterceptors(logPath, logLevel)

	opts := []grpc.DialOption{
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(grpc_middleware.ChainUnaryClient(unaryInterceptors...)),
	}

	return opts
}

func grpcUnaryClientInterceptors(logPath, logLevel string) ([]grpc.UnaryClientInterceptor, error) {
	logger, err := newLogger(logPath, logLevel)
	if err != nil {
		return nil, err
	}

	opts := []grpc_zap.Option{}

	interceptors := []grpc.UnaryClientInterceptor{
		grpc_zap.UnaryClientInterceptor(logger, opts...),
		accessLogUnaryClientInterceptor(logger),
	}

	return interceptors, nil
}

func accessLogUnaryClientInterceptor(logger *zap.Logger) grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		logEntry := logger.With(newClientLoggerFields(ctx, method, req)...)

		err := invoker(ctx, method, req, reply, cc, opts...)
		if err == nil {
			logProtoMessageAsJSON(logEntry, reply, "grpc.response.content", "client access log")
		}

		return err
	}
}

func newClientLoggerFields(ctx context.Context, fullMethodString string, req interface{}) []zapcore.Field {
	service := path.Dir(fullMethodString)[1:]
	method := path.Base(fullMethodString)

	return []zapcore.Field{
		zap.String("grpc.service", service),
		zap.String("grpc.method", method),
		zap.Reflect("grpc.request.content", req),
	}
}

func logProtoMessageAsJSON(logger *zap.Logger, pbMsg interface{}, key string, msg string) {
	if p, ok := pbMsg.(proto.Message); ok {
		logger.Check(zapcore.InfoLevel, msg).Write(zap.Object(key, &jsonpbObjectMarshaler{pb: p}))
	}
}

type jsonpbObjectMarshaler struct {
	pb proto.Message
}

func (j *jsonpbObjectMarshaler) MarshalLogObject(e zapcore.ObjectEncoder) error {
	return e.AddReflected("msg", j)
}

func (j *jsonpbObjectMarshaler) MarshalJSON() ([]byte, error) {
	b := &bytes.Buffer{}
	if err := JsonPbMarshaller.Marshal(b, j.pb); err != nil {
		return nil, fmt.Errorf("jsonpb serializer failed: %v", err)
	}
	return b.Bytes(), nil
}
