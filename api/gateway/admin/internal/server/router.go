package server

import (
	"fmt"

	"github.com/calmato/gran-book/api/gateway/admin/internal/entity"
	"github.com/calmato/gran-book/api/gateway/admin/internal/server/util"
	"github.com/gin-gonic/gin"
)

// Router - ルーティングの定義
func Router(reg *Registry, opts ...gin.HandlerFunc) *gin.Engine {
	r := gin.Default()
	r.Use(opts...)

	// non auth required routes
	r.GET("/health", reg.Health.Check)

	r.NoRoute(func(c *gin.Context) {
		err := fmt.Errorf("not found")
		util.ErrorHandling(c, entity.ErrNotFound.New(err))
	})

	return r
}
