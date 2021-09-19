package v2

import (
	"time"

	"github.com/calmato/gran-book/api/gateway/native/proto/service/book"
	"github.com/calmato/gran-book/api/gateway/native/proto/service/user"
	"github.com/gin-gonic/gin"
)

/**
 * ###############################################
 * handler
 * ###############################################
 */
type APIV2Handler interface {
	Routes(rg *gin.RouterGroup)
}

type apiV2Handler struct {
	Params
	now func() time.Time
}

type Params struct {
	Auth user.AuthServiceClient
	Book book.BookServiceClient
	User user.UserServiceClient
}

func NewAPIV2Handler(params *Params, now func() time.Time) APIV2Handler {
	return &apiV2Handler{
		Params: *params,
		now:    now,
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
