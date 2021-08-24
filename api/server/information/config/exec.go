package config

import (
	"context"

	"github.com/calmato/gran-book/api/server/information/pkg/database"
	"github.com/calmato/gran-book/api/server/information/pkg/firebase"
	"github.com/calmato/gran-book/api/server/information/pkg/firebase/storage"
	"github.com/calmato/gran-book/api/server/information/registry"
	"google.golang.org/api/option"
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

	// Firebaseの設定
	opt := option.WithCredentialsJSON([]byte(env.GCPServiceKeyJSON))

	fb, err := firebase.InitializeApp(ctx, nil, opt)
	if err != nil {
		return err
	}

	// Cloud Storageの設定
	gcs, err := storage.NewClient(ctx, fb.App, env.GCPStorageBucketName)
	if err != nil {
		return err
	}

	reg := registry.NewRegistry(db, gcs)

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
