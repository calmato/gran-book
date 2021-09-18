package server

import (
	"fmt"

	"github.com/calmato/gran-book/api/gateway/native/internal/entity"
	"github.com/calmato/gran-book/api/gateway/native/internal/server/util"
	"github.com/gin-gonic/gin"
)

// Router - ルーティングの定義
func Router(reg *Registry, opts ...gin.HandlerFunc) *gin.Engine {
	r := gin.Default()
	r.Use(opts...)

	authRequiredGroup := r.Group("")
	authRequiredGroup.Use(reg.HTTP.SetRequestID())
	authRequiredGroup.Use(reg.Authenticator.Authentication())

	// auth required routes
	reg.V1Api.Routes(authRequiredGroup)
	reg.V2Api.Routes(authRequiredGroup)

	// non auth required routes
	reg.V1Api.Routes(r.Group(""))
	r.GET("/health", reg.HTTP.HealthCheck)

	// other routes
	r.NoRoute(func(c *gin.Context) {
		err := fmt.Errorf("not found")
		util.ErrorHandling(c, entity.ErrNotFound.New(err))
	})

	return r
}
