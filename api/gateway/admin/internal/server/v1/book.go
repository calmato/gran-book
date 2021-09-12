package v1

import (
	"net/http"
	"strconv"

	"github.com/calmato/gran-book/api/gateway/admin/internal/entity"
	"github.com/calmato/gran-book/api/gateway/admin/internal/server/util"
	"github.com/calmato/gran-book/api/gateway/admin/proto/service/book"
	"github.com/gin-gonic/gin"
)

type BookHandler interface {
	Delete(ctx *gin.Context)
}

type bookHandler struct {
	bookClient book.BookServiceClient
}

func NewBookHandler(bc book.BookServiceClient) BookHandler {
	return &bookHandler{
		bookClient: bc,
	}
}

func (h *bookHandler) Delete(ctx *gin.Context) {
	bookID, err := strconv.ParseInt(ctx.Param("bookID"), 10, 64)
	if err != nil {
		util.ErrorHandling(ctx, entity.ErrBadRequest.New(err))
		return
	}

	in := &book.DeleteBookRequest{
		BookId: bookID,
	}

	c := util.SetMetadata(ctx)
	_, err = h.bookClient.DeleteBook(c, in)
	if err != nil {
		util.ErrorHandling(ctx, err)
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}
