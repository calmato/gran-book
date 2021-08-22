package server

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type HTTPServer struct {
	ctx    context.Context
	server *http.Server
}

type MetricsServer struct {
	ctx    context.Context
	server *http.Server
}

func NewHTTPServer(ctx context.Context, port string, r *gin.Engine) *HTTPServer {
	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: r,
	}

	return &HTTPServer{
		ctx:    ctx,
		server: server,
	}
}

func (s *HTTPServer) Serve() error {
	return s.server.ListenAndServe()
}

func (s *HTTPServer) Stop() error {
	return s.server.Shutdown(s.ctx)
}

func NewMetricsServer(ctx context.Context, port string) *MetricsServer {
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
