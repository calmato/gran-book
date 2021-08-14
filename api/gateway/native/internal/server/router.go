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
			apiV1Auth.GET("", reg.V1Auth.Get)
			apiV1Auth.DELETE("", reg.V1Auth.Delete)
			apiV1Auth.PATCH("/address", reg.V1Auth.UpdateAddress)
			apiV1Auth.POST("/device", reg.V1Auth.RegisterDevice)
			apiV1Auth.PATCH("/email", reg.V1Auth.UpdateEmail)
			apiV1Auth.PATCH("/password", reg.V1Auth.UpdatePassword)
			apiV1Auth.PATCH("/profile", reg.V1Auth.UpdateProfile)
			apiV1Auth.POST("/thumbnail", reg.V1Auth.UploadThumbnail)
		}
		// User Service
		apiV1User := apiV1.Group("/users")
		{
			apiV1User.GET("/:userID/follows", reg.V1User.ListFollow)
			apiV1User.GET("/:userID/followers", reg.V1User.ListFollower)
			apiV1User.GET("/:userID/profile", reg.V1User.GetProfile)
			apiV1User.POST("/:userID/follow/:followerID", reg.V1User.Follow)
			apiV1User.DELETE("/:userID/follow/:followerID", reg.V1User.Unfollow)
		}
		// Chat Service
		apiV1Chat := apiV1.Group("/users/:userID/chat")
		{
			apiV1Chat.GET("", reg.V1Chat.ListRoom)
			apiV1Chat.POST("/:roomID", reg.V1Chat.CreateRoom)
			apiV1Chat.POST("/:roomID/messages/text", reg.V1Chat.CreateTextMessage)
			apiV1Chat.POST("/:roomID/messages/image", reg.V1Chat.CreateImageMessage)
		}
		// Book Servicea
		apiV1Book := apiV1.Group("/books")
		{
			apiV1Book.POST("", reg.V1Book.Create)
			apiV1Book.PATCH("", reg.V1Book.Update)
			apiV1Book.GET("/:isbn", reg.V1Book.Get)
			apiV1Book.GET("/:bookID/reviews", reg.V1Book.ListReview)
			apiV1Book.GET("/:bookID/reviews/:reviewID", reg.V1Book.GetReview)
		}
	}

	r.NoRoute(func(c *gin.Context) {
		err := fmt.Errorf("not found")
		util.ErrorHandling(c, entity.ErrNotFound.New(err))
	})

	return r
}
