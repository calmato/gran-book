package config

import (
	"context"

	"github.com/calmato/gran-book/api/server/user/internal/infrastructure/repository"
	"github.com/calmato/gran-book/api/server/user/lib/firebase"
	"github.com/calmato/gran-book/api/server/user/lib/firebase/authentication"
	"github.com/calmato/gran-book/api/server/user/lib/firebase/firestore"
	"github.com/calmato/gran-book/api/server/user/lib/firebase/storage"
	"github.com/calmato/gran-book/api/server/user/registry"
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
	db, err := repository.NewDBClient(
		env.DBSocket, env.DBHost, env.DBPort, env.DBDatabase, env.DBUsername, env.DBPassword,
	)
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

	// Firestoreの設定
	fs, err := firestore.NewClient(ctx, fb.App)
	if err != nil {
		panic(err)
	}
	defer fs.Close()

	// Cloud Storageの設定
	gcs, err := storage.NewClient(ctx, fb.App, env.GCPStorageBucketName)
	if err != nil {
		return err
	}

	reg := registry.NewRegistry(db, fa, fs, gcs)

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
