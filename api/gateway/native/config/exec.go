package config

import (
	"context"

	"github.com/calmato/gran-book/api/gateway/native/internal/server"
	"github.com/gin-gonic/gin"
)

// Execute - HTTP Serverの起動
func Execute() error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	env, err := loadEnvironment()
	if err != nil {
		return err
	}

	opts := []gin.HandlerFunc{}

	// Logger設定
	logger, err := newLogger(env.LogPath, env.LogLevel)
	if err != nil {
		return err
	}
	opts = append(opts, logger)

	// CORS設定
	cors, err := newCors()
	if err != nil {
		return err
	}
	opts = append(opts, cors)

	// 依存関係の解決
	reg := registry.NewRegistry()

	// メトリクス用のHTTP Serverの起動
	ms := server.NewMetricsServer(env.MetricsPort)
	go func() {
		if err := ms.Serve(); err != nil {
			panic(err)
		}
	}()

	// HTTP Serverの起動
	r := server.Router(reg, opts)
	hs := server.NewHTTPServer(r, env.Port)
	if err := hs.Serve(); err != nil {
		panic(err)
	}

	return nil
}
