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

	authRequiredGroup := r.Group("")
	authRequiredGroup.Use(reg.Authenticator.Authentication())

	// non auth required routes
	r.GET("/health", reg.Health.Check)

	// auth required routes
	apiV1Router(reg, authRequiredGroup)

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
		// Admin Service
		apiV1Admin := apiV1.Group("/admin")
		{
			apiV1AdminAllRole := apiV1Admin.Group("", reg.Authenticator.Authentication())
			{
				apiV1AdminAllRole.GET("", reg.V1Admin.List)
				apiV1AdminAllRole.GET("/:userID", reg.V1Admin.Get)
			}
			apiV1AdminAdminRole := apiV1Admin
			{
				apiV1AdminAdminRole.POST("/:userID", reg.V1Admin.Create, reg.Authenticator.HasAdminRole())
				apiV1AdminAdminRole.POST("/:userID/thumbnail", reg.V1Admin.UploadThumbnail, reg.Authenticator.HasAdminRole())
				apiV1AdminAdminRole.DELETE("/:userID", reg.V1Admin.Delete, reg.Authenticator.HasAdminRole())
				apiV1AdminAdminRole.PATCH("/:userID/contact", reg.V1Admin.UpdateContact, reg.Authenticator.HasAdminRole())
				apiV1AdminAdminRole.PATCH("/:userID/profile", reg.V1Admin.UpdateProfile, reg.Authenticator.HasAdminRole())
				apiV1AdminAdminRole.PATCH("/:userID/password", reg.V1Admin.UpdatePassword, reg.Authenticator.HasAdminRole())
			}
		}
		// User Service
		apiV1User := apiV1.Group("/users")
		{
			apiV1UserAllRole := apiV1User.Group("", reg.Authenticator.Authorization())
			{
				apiV1UserAllRole.GET("", reg.V1User.List)
				apiV1UserAllRole.GET("/:userID", reg.V1User.Get)
			}
		}
		// Book Service
		apiV1Book := apiV1.Group("/books")
		{
			apiV1BookAllRole := apiV1Book.Group("", reg.Authenticator.Authorization())
			{
				apiV1BookAllRole.DELETE("/:bookID", reg.V1Book.Delete)
			}
		}
	}
}
