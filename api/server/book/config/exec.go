package config

import (
	"context"

	"github.com/calmato/gran-book/api/server/book/pkg/database"
	"github.com/calmato/gran-book/api/server/book/registry"
)

// Execute - gRPC Serverの起動
func Execute() error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	env, err := loadEnvironment()
	if err != nil {
		return err
	}

	// MySQL Clientの設定
	db, err := database.NewClient(
		env.DBSocket, env.DBHost, env.DBPort, env.DBDatabase, env.DBUsername, env.DBPassword,
	)
	if err != nil {
		return err
	}

	reg := registry.NewRegistry(db)

	// gRPC Serverの設定取得
	gs, err := newGRPCServer(env.Port, env.LogPath, env.LogLevel, reg)
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

	// gRPC Serverの起動
	if err := gs.Serve(); err != nil {
		panic(err)
	}

	return nil
}
