package server

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type httpServer struct {
	router *gin.Engine
	port   string
}

type metricsServer struct {
	serve *http.ServeMux
	port  string
}

func NewHTTPServer(r *gin.Engine, port string) *httpServer {
	s := &httpServer{
		router: r,
		port:   fmt.Sprint(":%s", port),
	}

	return s
}

func (s *httpServer) Serve() error {
	return http.ListenAndServe(s.port, s.router)
}

func NewMetricsServer(port string) *metricsServer {
	mux := http.NewServeMux()
	mux.Handle("/metrics", promhttp.Handler())

	s := &metricsServer{
		serve: mux,
		port:  fmt.Sprint(":%s", port),
	}

	return s
}

func (s *metricsServer) Serve() error {
	return http.ListenAndServe(s.port, s.serve)
}
