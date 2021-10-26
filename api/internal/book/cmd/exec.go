package cmd

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	config "github.com/calmato/gran-book/api/config/book"
	"github.com/calmato/gran-book/api/pkg/database"
	"github.com/calmato/gran-book/api/pkg/log"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
)

func Exec() error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	env, err := config.LoadEnvironment()
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

	// Loggerの設定
	logger, err := log.NewLogger(env.LogPath, env.LogLevel)
	if err != nil {
		return err
	}

	reg := newRegistry(db)

	// gRPC Serverの設定取得
	gs, err := newGRPCServer(env.Port, logger, reg)
	if err != nil {
		return err
	}

	// gRPC Serverの起動
	eg, ectx := errgroup.WithContext(ctx)
	eg.Go(func() (err error) {
		err = gs.Serve()
		if err != nil {
			logger.Error("Failed to run server", zap.Error(err))
		}
		return
	})
	logger.Info("Started server", zap.String("port", env.Port))

	// シグナル検知設定
	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, syscall.SIGTERM, syscall.SIGINT)
	select {
	case <-ectx.Done():
		logger.Error("Done context", zap.Error(ectx.Err()))
	case <-signalCh:
		logger.Info("Received signal")
		delay := 20 * time.Second
		logger.Info("Pre-shutdown", zap.String("delay", delay.String()))
		time.Sleep(delay)
	}

	logger.Info("Shutdown...")
	gs.Stop()
	cancel()
	return eg.Wait()
}
