package native

import (
	"context"
	"fmt"
	"net/http"

	"github.com/calmato/gran-book/api/internal/gateway/util"
	"github.com/calmato/gran-book/api/pkg/exception"
	"github.com/gin-gonic/gin"
)

type HTTPServer struct {
	ctx    context.Context
	server *http.Server
}

func newHTTPServer(ctx context.Context, port string, r *gin.Engine) *HTTPServer {
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

func newRouter(reg *registry, opts ...gin.HandlerFunc) *gin.Engine {
	rt := gin.Default()
	rt.Use(opts...)

	authRequiredGroup := rt.Group("")
	authRequiredGroup.Use(util.SetRequestID())
	authRequiredGroup.Use(reg.Authenticator.Authentication())

	// auth required routes
	reg.V1Api.Routes(authRequiredGroup)
	reg.V2Api.Routes(authRequiredGroup)

	// non auth required routes
	reg.V1Api.NonAuthRoutes(rt.Group(""))
	rt.GET("/health", util.HealthCheck)

	// other routes
	rt.NoRoute(func(c *gin.Context) {
		err := fmt.Errorf("not found")
		util.ErrorHandling(c, exception.ErrNotFound.New(err))
	})

	return rt
}
