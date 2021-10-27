package admin

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	config "github.com/calmato/gran-book/api/config/gateway/admin"
	"github.com/calmato/gran-book/api/pkg/cors"
	"github.com/calmato/gran-book/api/pkg/firebase"
	"github.com/calmato/gran-book/api/pkg/firebase/authentication"
	"github.com/calmato/gran-book/api/pkg/log"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
	"google.golang.org/api/option"
)

//nolint:funlen
func Exec() error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	env, err := config.LoadEnvironment()
	if err != nil {
		return err
	}

	// Loggerの設定
	logger, err := log.NewLogger(env.LogPath, env.LogLevel)
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
	lm, err := log.NewGinMiddleware(env.LogPath, env.LogLevel)
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
	p := &params{
		FirebaseAuth:    fa,
		AdminServiceURL: env.AdminServiceURL,
		AuthServiceURL:  env.AuthServiceURL,
		BookServiceURL:  env.BookServiceURL,
		UserServiceURL:  env.UserServiceURL,
	}
	reg, err := newRegistry(p)
	if err != nil {
		return err
	}

	// Metrics Serverの設定取得
	ms := newMetricsServer(ctx, env.MetricsPort)

	// HTTP Serverの設定取得
	rt := newRouter(reg, opts...)
	hs := newHTTPServer(ctx, env.Port, rt)

	// HTTP Server, Metrics Serverの起動
	eg, ectx := errgroup.WithContext(ctx)
	eg.Go(func() (err error) {
		err = ms.Serve()
		if err != nil {
			logger.Error("Failed to run metrics server", zap.Error(err))
		}
		return
	})
	eg.Go(func() (err error) {
		err = hs.Serve()
		if err != nil {
			logger.Error("Failed to run http server", zap.Error(err))
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
	err = hs.Stop()
	if err != nil {
		return err
	}
	err = ms.Stop()
	if err != nil {
		return err
	}
	return eg.Wait()
}
