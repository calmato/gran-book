package config

import (
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
)

func setCors(mux *runtime.ServeMux) http.Handler {
	handler := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{http.MethodPost, http.MethodGet, http.MethodPut, http.MethodDelete}),
		handlers.AllowedHeaders([]string{"Authorization", "Content-Type", "Accept-Encoding", "Accept"}),
	)(mux)

	return handler
}
