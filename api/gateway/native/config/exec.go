package config

import (
	"fmt"

	"github.com/calmato/gran-book/api/gateway/native/internal/server"
	"github.com/calmato/gran-book/api/gateway/native/pkg/cors"
	"github.com/calmato/gran-book/api/gateway/native/pkg/logger"
	"github.com/calmato/gran-book/api/gateway/native/registry"
	"github.com/gin-gonic/gin"
)

// Execute - HTTP Serverの起動
func Execute() error {
	env, err := loadEnvironment()
	if err != nil {
		return err
	}

	opts := []gin.HandlerFunc{}

	// Logger設定
	lm, err := logger.NewGinMiddleware(env.LogPath, env.LogLevel)
	if err != nil {
		return err
	}
	opts = append(opts, lm)

	// CORS設定
	cm, err := cors.NewGinMiddleware()
	if err != nil {
		return err
	}
	opts = append(opts, cm)

	// 依存関係の解決
	reg, err := registry.NewRegistry(env.AuthServiceURL)
	if err != nil {
		return err
	}

	// メトリクス用のHTTP Serverの起動
	ms := server.NewMetricsServer(env.MetricsPort)
	go func() {
		if err := ms.Serve(); err != nil {
			panic(err)
		}
		fmt.Println("metrics server is runnning...")
	}()

	// HTTP Serverの起動
	r := server.Router(reg, opts...)
	hs := server.NewHTTPServer(r, env.Port)
	if err := hs.Serve(); err != nil {
		panic(err)
	}
	fmt.Println("http server is runnning...")

	return nil
}
