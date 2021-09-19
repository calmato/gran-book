package v1

import (
	"context"
	"time"

	"github.com/calmato/gran-book/api/gateway/native/internal/entity"
	"github.com/calmato/gran-book/api/gateway/native/proto/service/book"
	"github.com/calmato/gran-book/api/gateway/native/proto/service/chat"
	"github.com/calmato/gran-book/api/gateway/native/proto/service/user"
	"github.com/gin-gonic/gin"
)

/**
 * ###############################################
 * handler
 * ###############################################
 */
type ApiV1Handler interface {
	Routes(rg *gin.RouterGroup)
	NonAuthRoutes(rg *gin.RouterGroup)
}

type apiV1Handler struct {
	Params
	now func() time.Time
}

type Params struct {
	Auth user.AuthServiceClient
	Book book.BookServiceClient
	Chat chat.ChatServiceClient
	User user.UserServiceClient
}

func NewApiV1Handler(params *Params, now func() time.Time) ApiV1Handler {
	return &apiV1Handler{
		Params: *params,
		now:    now,
	}
}

/**
 * ###############################################
 * routes
 * ###############################################
 */
func (h *apiV1Handler) Routes(rg *gin.RouterGroup) {
	apiV1 := rg.Group("/v1")
	h.v1ReviewRoutes(apiV1.Group(""))
	h.v1AuthRoutes(apiV1.Group("/auth"))
	h.v1BookRoutes(apiV1.Group("/books"))
	h.v1TopRoutes(apiV1.Group("/top"))
	h.v1UserRoutes(apiV1.Group("/users"))
	h.v1ChatRoutes(apiV1.Group("/users/:userID/chat"))
	h.v1BookshelfRoutes(apiV1.Group("/users/:userID/books"))
}

func (h *apiV1Handler) NonAuthRoutes(rg *gin.RouterGroup) {
	rg.POST("/v1/auth", h.createAuth)
}

// Top Service
func (h *apiV1Handler) v1TopRoutes(rg *gin.RouterGroup) {
	rg.GET("/user", h.getTopUser)
}

// Auth Service
func (h *apiV1Handler) v1AuthRoutes(rg *gin.RouterGroup) {
	rg.GET("", h.getAuth)
	rg.DELETE("", h.deleteAuth)
	rg.POST("/device", h.registerAuthDevice)
	rg.POST("/thumbnail", h.uploadAuthThumbnail)
	rg.PATCH("/address", h.updateAuthAddress)
	rg.PATCH("/email", h.updateAuthEmail)
	rg.PATCH("/password", h.updateAuthPassword)
	rg.PATCH("/profile", h.updateAuthProfile)
}

// User Service
func (h *apiV1Handler) v1UserRoutes(rg *gin.RouterGroup) {
	rg.GET("/:userID/follows", h.listUserFollow)
	rg.GET("/:userID/followers", h.listUserFollower)
	rg.GET("/:userID/profile", h.getUserProfile)
	rg.POST("/:userID/follow/:followerID", h.userFollow)
	rg.DELETE("/:userID/follow/:followerID", h.userUnfollow)
}

// Chat Service
func (h *apiV1Handler) v1ChatRoutes(rg *gin.RouterGroup) {
	rg.GET("", h.listChatRoom)
	rg.POST("", h.createChatRoom)
	rg.POST("/:roomID/messages/text", h.createChatTextMessage)
	rg.POST("/:roomID/messages/image", h.createChatImageMessage)
}

// Book Service
func (h *apiV1Handler) v1BookRoutes(rg *gin.RouterGroup) {
	rg.POST("", h.createBook)
	rg.PATCH("", h.updateBook)
	rg.GET("/:bookID", h.getBook)
}

// Bookshelf Service
func (h *apiV1Handler) v1BookshelfRoutes(rg *gin.RouterGroup) {
	rg.GET("", h.listBookshelf)
	rg.GET("/:bookID", h.getBookshelf)
	rg.POST("/:bookID/read", h.readBookshelf)
	rg.POST("/:bookID/reading", h.readingBookshelf)
	rg.POST("/:bookID/stack", h.stackedBookshelf)
	rg.POST("/:bookID/want", h.wantBookshelf)
	rg.POST("/:bookID/release", h.releaseBookshelf)
	rg.DELETE("/:bookID", h.deleteBookshelf)
}

func (h *apiV1Handler) v1ReviewRoutes(rg *gin.RouterGroup) {
	rg.GET("/books/:bookID/reviews", h.listReviewByBook)
	rg.GET("/books/:bookID/reviews/:reviewID", h.getBookReview)
	rg.GET("/users/:userID/reviews", h.listReviewByUser)
	rg.GET("/users/:userID/reviews/:reviewID", h.getUserReview)
}

/**
 * ###############################################
 * private methods
 * ###############################################
 */
// currentUser - UserIDを基に現在のログインユーザ情報を取得
func (h *apiV1Handler) currentUser(ctx context.Context, userID string) (*entity.Auth, error) {
	out, err := h.Auth.GetAuth(ctx, &user.Empty{})
	if err != nil {
		return nil, err
	}

	a := entity.NewAuth(out.Auth)

	if a.Id != userID {
		return nil, entity.ErrForbidden.New(err)
	}

	return a, nil
}
