package admin

import (
	"context"
	"fmt"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type MetricsServer struct {
	ctx    context.Context
	server *http.Server
}

func newMetricsServer(ctx context.Context, port string) *MetricsServer {
	mux := http.NewServeMux()
	mux.Handle("/metrics", promhttp.Handler())

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: mux,
	}

	return &MetricsServer{
		ctx:    ctx,
		server: server,
	}
}

func (s *MetricsServer) Serve() error {
	return s.server.ListenAndServe()
}

func (s *MetricsServer) Stop() error {
	return s.server.Shutdown(s.ctx)
}
