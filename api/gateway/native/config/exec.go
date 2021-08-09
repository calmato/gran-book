package config

import (
	"context"

	"github.com/calmato/gran-book/api/gateway/native/internal/server"
	"github.com/calmato/gran-book/api/gateway/native/pkg/cors"
	"github.com/calmato/gran-book/api/gateway/native/pkg/firebase"
	"github.com/calmato/gran-book/api/gateway/native/pkg/firebase/authentication"
	"github.com/calmato/gran-book/api/gateway/native/pkg/logger"
	"github.com/calmato/gran-book/api/gateway/native/registry"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/option"
)

// Execute - HTTP Serverの起動
func Execute() error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	env, err := loadEnvironment()
	if err != nil {
		return err
	}

	// Firebaseの設定
	opt := option.WithCredentialsJSON([]byte(env.GCPServiceKeyJSON))

	fb, err := firebase.InitializeApp(ctx, nil, opt)
	if err != nil {
		return err
	}

	// Firebase Authenticationの設定
	fa, err := authentication.NewClient(ctx, fb.App)
	if err != nil {
		return err
	}

	opts := []gin.HandlerFunc{gin.Recovery()}

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
	reg, err := registry.NewRegistry(fa, env.AuthServiceURL, env.UserServiceURL)
	if err != nil {
		return err
	}

	// メトリクス用のHTTP Serverの起動
	ms := server.NewMetricsServer(env.MetricsPort)
	go func() {
		if err := ms.Serve(); err != nil {
			panic(err)
		}
	}()

	// HTTP Serverの起動
	r := server.Router(reg, opts...)
	hs := server.NewHTTPServer(r, env.Port)
	if err := hs.Serve(); err != nil {
		panic(err)
	}

	return nil
}
