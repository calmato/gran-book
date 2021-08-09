package server

import (
	"fmt"

	"github.com/calmato/gran-book/api/gateway/native/internal/entity"
	"github.com/calmato/gran-book/api/gateway/native/internal/server/util"
	"github.com/calmato/gran-book/api/gateway/native/registry"
	"github.com/gin-gonic/gin"
)

// Router - ルーティングの定義
func Router(reg *registry.Registry, opts ...gin.HandlerFunc) *gin.Engine {
	r := gin.Default()
	r.Use(opts...)

	authRequiredGroup := r.Group("")
	authRequiredGroup.Use(reg.Authenticator.Authentication())

	// non auth required routes
	r.GET("/health", reg.Health.Check)
	r.POST("/v1/auth", reg.V1Auth.Create)

	// API v1 routes
	apiV1 := authRequiredGroup.Group("/v1")
	{
		// Auth Service
		apiV1Auth := apiV1.Group("/auth")
		{
			apiV1Auth.GET("/", reg.V1Auth.Get)
			apiV1Auth.DELETE("/", reg.V1Auth.Delete)
			apiV1Auth.PATCH("/address", reg.V1Auth.UpdateAddress)
			apiV1Auth.POST("/device", reg.V1Auth.RegisterDevice)
			apiV1Auth.PATCH("/email", reg.V1Auth.UpdateEmail)
			apiV1Auth.PATCH("/password", reg.V1Auth.UpdatePassword)
			apiV1Auth.PATCH("/profile", reg.V1Auth.UpdateProfile)
			apiV1Auth.POST("/thumbnail", reg.V1Auth.UploadThumbnail)
		}
	}

	r.NoRoute(func(c *gin.Context) {
		err := fmt.Errorf("not found")
		util.ErrorHandling(c, entity.ErrNotFound.New(err))
	})

	return r
}
