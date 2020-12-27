package config

import (
	"context"
	"flag"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
)

var (
	// command-line options:
	// gRPC server endpoint
	helloAPIEndpoint = flag.String("HELLO_API", "hello_api:8080", "gRPC server endpoint")
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

	mux := runtime.NewServeMux()

	if err := registerServiceHandlers(ctx, mux, env.LogPath, env.LogLevel); err != nil {
		return err
	}

	handler := setCors(mux)

	if err := http.ListenAndServe(":"+env.Port, handler); err != nil {
		return err
	}

	return nil
}
