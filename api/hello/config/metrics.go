package config

import (
	"fmt"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type httpServer struct {
	s    *http.ServeMux
	port string
}

func newHTTPServer(port string) *httpServer {
	mux := http.NewServeMux()
	mux.Handle("/metrics", promhttp.Handler())

	hs := &httpServer{
		s:    mux,
		port: fmt.Sprintf(":%s", port),
	}

	return hs
}

func (s *httpServer) Serve() error {
	return http.ListenAndServe(s.port, s.s)
}
