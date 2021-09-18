package v2

import (
	"github.com/calmato/gran-book/api/gateway/native/proto/service/book"
	"github.com/calmato/gran-book/api/gateway/native/proto/service/user"
	"github.com/gin-gonic/gin"
)

/**
 * ###############################################
 * handler
 * ###############################################
 */
type ApiV2Handler interface {
	Routes(rg *gin.RouterGroup)
}

type apiV2Handler struct {
	Params
}

type Params struct {
	Auth user.AuthServiceClient
	Book book.BookServiceClient
	User user.UserServiceClient
}

func NewApiV2Handler(params *Params) ApiV2Handler {
	return &apiV2Handler{
		Params: *params,
	}
}

/**
 * ###############################################
 * routes
 * ###############################################
 */
func (h *apiV2Handler) Routes(rg *gin.RouterGroup) {
	apiV2 := rg.Group("/v2")
	h.v2BookRoutes(apiV2.Group("/books"))
	h.v2BookshelfRoutes(apiV2.Group("/users/:userID/books"))
}

// Book Service
func (h *apiV2Handler) v2BookRoutes(rg *gin.RouterGroup) {
	rg.GET("/:bookID", h.getBook)
}

// Bookshelf Service
func (h *apiV2Handler) v2BookshelfRoutes(rg *gin.RouterGroup) {
	rg.GET("", h.listBookshelf)
	rg.GET("/:bookID", h.getBookshelf)
}
