package config

import (
	"github.com/calmato/gran-book/api/server/book/pkg/database"
	"github.com/calmato/gran-book/api/server/book/registry"
)

// Execute - gRPC Serverの起動
func Execute() error {
	env, err := loadEnvironment()
	if err != nil {
		return err
	}

	// MySQL Clientの設定
	dbp := &database.Params{
		Socket:   env.DBSocket,
		Host:     env.DBHost,
		Port:     env.DBPort,
		Database: env.DBDatabase,
		Username: env.DBUsername,
		Password: env.DBPassword,
	}
	db, err := database.NewClient(dbp)
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
