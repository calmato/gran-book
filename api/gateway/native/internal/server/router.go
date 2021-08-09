package server

import (
	"errors"

	"github.com/calmato/gran-book/api/gateway/native/internal/entity"
	"github.com/calmato/gran-book/api/gateway/native/internal/server/util"
	"github.com/calmato/gran-book/api/gateway/native/registry"
	"github.com/gin-gonic/gin"
)

// Router - ルーティングの定義
func Router(reg *registry.Registry, opts ...gin.HandlerFunc) *gin.Engine {
	r := gin.Default()

	for _, opt := range opts {
		r.Use(opt)
	}

	r.NoRoute(func(c *gin.Context) {
		err := errors.New("not found")
		util.ErrorHandling(c, entity.ErrNotFound.New(err))
	})

	// API v1 routes
	apiV1 := r.Group("/v1")
	{
		// Auth Service
		apiV1.GET("/auth", reg.V1Auth.Get)
	}

	return r
}
