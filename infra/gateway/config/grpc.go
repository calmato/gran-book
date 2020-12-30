package config

import (
	"context"
	"encoding/json"
	"fmt"
	"path"
	"strings"
	"time"

	gw "github.com/calmato/gran-book/infra/gateway/proto"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func registerServiceHandlers(ctx context.Context, mux *runtime.ServeMux, logPath, logLevel string) error {
	opts := grpcDialOptions(logPath, logLevel)

	err := gw.RegisterGreeterHandlerFromEndpoint(ctx, mux, *helloAPIEndpoint, opts)
	if err != nil {
		return err
	}

	grpc_prometheus.EnableClientHandlingTimeHistogram()

	return nil
}

/*
 * DialOptions
 */
func grpcDialOptions(logPath, logLevel string) []grpc.DialOption {
	unaryInterceptors, _ := grpcUnaryClientInterceptors(logPath, logLevel)

	opts := []grpc.DialOption{
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(grpc_middleware.ChainUnaryClient(unaryInterceptors...)),
	}

	return opts
}

/*
 * DialOptions - UnaryClientInterceptor
 */
func grpcUnaryClientInterceptors(logPath, logLevel string) ([]grpc.UnaryClientInterceptor, error) {
	logger, err := newLogger(logPath, logLevel)
	if err != nil {
		return nil, err
	}

	opts := []grpc_zap.Option{}

	interceptors := []grpc.UnaryClientInterceptor{
		grpc_zap.UnaryClientInterceptor(logger, opts...),
		grpc_prometheus.UnaryClientInterceptor,
		accessLogUnaryClientInterceptor(logger),
	}

	return interceptors, nil
}

func accessLogUnaryClientInterceptor(logger *zap.Logger) grpc.UnaryClientInterceptor {
	return func(
		ctx context.Context, method string, req, reply interface{},
		cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption,
	) error {
		startTime := time.Now()

		err := invoker(ctx, method, req, reply, cc, opts...)

		duration := time.Since(startTime)
		fields := newClientLoggerFields(ctx, method, req, reply, duration, err)
		logger.Check(zap.InfoLevel, "client request/response payload logged").Write(fields...)

		return err
	}
}

/*
 * ログの整形用
 */
func newClientLoggerFields(
	ctx context.Context, fullMethodString string, reqPbMsg, resPbMsg interface{}, duration time.Duration, err error,
) []zapcore.Field {
	service := path.Dir(fullMethodString)[1:]
	method := path.Base(fullMethodString)

	fields := []zapcore.Field{
		zap.String("grpc.service", service),
		zap.String("grpc.method", method),
		zap.Duration("grpc.duration", duration),
	}

	if md, ok := metadata.FromOutgoingContext(ctx); ok {
		if values := md.Get("grpcgateway-user-agent"); len(values) > 0 {
			fields = append(fields, zap.String("grpc.user_agent", values[0]))
		}
		if values := md.Get("x-forwarded-for"); len(values) > 0 {
			fields = append(fields, zap.String("gprc.remote_ip", values[0]))
		}
	}

	if p, ok := reqPbMsg.(proto.Message); ok {
		req, _ := filterParams(p)
		fields = append(fields, zap.Reflect("grpc.request.content", req))
	}

	if p, ok := resPbMsg.(proto.Message); ok {
		res, _ := filterParams(p)
		fields = append(fields, zap.Reflect("grpc.request.content", res))
	}

	if err != nil {
		fields = append(fields, zap.String("grpc.err", err.Error()))
	}

	return fields
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
