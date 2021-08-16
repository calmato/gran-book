package server

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type HTTPServer struct {
	router *gin.Engine
	port   string
}

type MetricsServer struct {
	serve *http.ServeMux
	port  string
}

func NewHTTPServer(r *gin.Engine, port string) *HTTPServer {
	s := &HTTPServer{
		router: r,
		port:   fmt.Sprintf(":%s", port),
	}

	return s
}

func (s *HTTPServer) Serve() error {
	return http.ListenAndServe(s.port, s.router)
}

func NewMetricsServer(port string) *MetricsServer {
	mux := http.NewServeMux()
	mux.Handle("/metrics", promhttp.Handler())

	s := &MetricsServer{
		serve: mux,
		port:  fmt.Sprintf(":%s", port),
	}

	return s
}

func (s *MetricsServer) Serve() error {
	return http.ListenAndServe(s.port, s.serve)
}
