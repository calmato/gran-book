package config

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// Execute - gRPC Serverの起動
func Execute() error {
	env, err := loadEnvironment()
	if err != nil {
		return err
	}

	// gRPC Serverの起動
	gs, err := newGRPCServer(env.Port, env.LogPath, env.LogLevel)
	if err != nil {
		return err
	}

	if err := gs.Serve(); err != nil {
		panic(err)
	}

	// メトリクス用のHTTP Serverの起動
	mux := http.NewServeMux()
	mux.Handle("/metrics", promhttp.Handler())

	go func() {
		if http.ListenAndServe(env.MetricsPort, mux); err != nil {
			panic(err)
		}
	}()

	return nil
}
