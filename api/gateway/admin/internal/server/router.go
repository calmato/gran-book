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

// apiV1Router - API v1 routes
func apiV1Router(reg *Registry, rg *gin.RouterGroup) {
	apiV1 := rg.Group("/v1")
	{
		// Auth Service
		apiV1Auth := apiV1.Group("/auth")
		{
			apiV1Auth.GET("", reg.V1Auth.Get)
			apiV1Auth.PATCH("/email", reg.V1Auth.UpdateEmail)
			apiV1Auth.PATCH("/password", reg.V1Auth.UpdatePassword)
			apiV1Auth.PATCH("/profile", reg.V1Auth.UpdateProfile)
			apiV1Auth.POST("/thumbnail", reg.V1Auth.UploadThumbnail)
		}
	}
}
