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
		// Book Service
		apiV1Book := apiV1.Group("/books")
		{
			apiV1Book.POST("", reg.V1Book.Create)
			apiV1Book.PATCH("", reg.V1Book.Update)
			apiV1Book.GET("/:isbn", reg.V1Book.Get)
		}
		// Bookshelf Service
		apiV1Bookshelf := apiV1.Group("/users/:userID/books")
		{
			apiV1Bookshelf.GET("", reg.V1Bookshelf.List)
			apiV1Bookshelf.GET("/:bookID", reg.V1Bookshelf.Get)
			apiV1Bookshelf.POST("/:bookID/read", reg.V1Bookshelf.Read)
			apiV1Bookshelf.POST("/:bookID/reading", reg.V1Bookshelf.Reading)
			apiV1Bookshelf.POST("/:bookID/stack", reg.V1Bookshelf.Stacked)
			apiV1Bookshelf.POST("/:bookID/want", reg.V1Bookshelf.Want)
			apiV1Bookshelf.POST("/:bookID/release", reg.V1Bookshelf.Release)
			apiV1Bookshelf.DELETE("/:bookID", reg.V1Bookshelf.Delete)
		}
		// Other
		apiV1.GET("/books/:bookID/reviews", reg.V1Review.ListByBook)
		apiV1.GET("/books/:bookID/reviews/:reviewID", reg.V1Review.GetByBook)
		apiV1.GET("/users/:userID/reviews", reg.V1Review.ListByUser)
		apiV1.GET("/users/:userID/reviews/:reviewID", reg.V1Review.GetByUser)
	}
	// API v2 routes
	apiV2 := authRequiredGroup.Group("/v2")
	{
		// Book Service
		apiV2Book := apiV2.Group("/books")
		{
			apiV2Book.GET("/:bookID", reg.V2Book.Get)
		}
		// Bookshelf Service
		apiV2Bookshelf := apiV2.Group("/users/:userID/books")
		{
			apiV2Bookshelf.GET("", reg.V2Bookshelf.List)
			apiV2Bookshelf.GET("/:bookID", reg.V2Bookshelf.Get)
		}
	}

	r.NoRoute(func(c *gin.Context) {
		err := fmt.Errorf("not found")
		util.ErrorHandling(c, entity.ErrNotFound.New(err))
	})

	return r
}
