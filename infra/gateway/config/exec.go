package config

import (
	"context"
	"net/http"

	"github.com/calmato/gran-book/infra/gateway/internal/handler"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
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

	err = registerServiceHandlers(
		ctx, mux, env.LogPath, env.LogLevel, env.UserAPIURL, env.SSLValidation, env.SSLVerify,
	)
	if err != nil {
		return err
	}

	c := setCors()

	if err = http.ListenAndServe(":"+env.Port, c.Handler(mux)); err != nil {
		return err
	}

	return nil
}
