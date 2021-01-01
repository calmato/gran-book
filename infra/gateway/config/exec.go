package config

import (
	"context"
	"flag"
	"net/http"

	"github.com/calmato/gran-book/infra/gateway/internal/handler"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
)

var (
	// command-line options:
	// gRPC server endpoint
	helloAPIEndpoint = flag.String("HELLO_API", "hello_api:8080", "gRPC server endpoint")
	userAPIEndpoint  = flag.String("USER_API", "user_api:8080", "gRPC server endpoint")
)

// Execute - gRPC Gatewayの起動
func Execute() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	env, err := loadEnvironment()
	if err != nil {
		return err
	}

	// メトリクス用のHTTP Serverの起動
	hs := newHTTPServer(env.MetricsPort)

	go func() {
		if err := hs.Serve(); err != nil {
			panic(err)
		}
	}()

	// Client側のHTTP Serverの起動
	mux := runtime.NewServeMux(
		runtime.WithErrorHandler(handler.CustomHTTPError),
	)
	if err := registerServiceHandlers(ctx, mux, env.LogPath, env.LogLevel); err != nil {
		return err
	}

	c := setCors()

	if err := http.ListenAndServe(":"+env.Port, c.Handler(mux)); err != nil {
		return err
	}

	return nil
}
