package v1

import (
	"net/http"
	"strconv"

	"github.com/calmato/gran-book/api/service/internal/gateway/util"
	"github.com/calmato/gran-book/api/service/pkg/exception"
	"github.com/calmato/gran-book/api/service/proto/book"
	"github.com/gin-gonic/gin"
)

func (h *apiV1Handler) deleteBook(ctx *gin.Context) {
	c := util.SetMetadata(ctx)

	bookID, err := strconv.ParseInt(ctx.Param("bookID"), 10, 64)
	if err != nil {
		util.ErrorHandling(ctx, exception.ErrInvalidArgument.New(err))
		return
	}

	in := &book.DeleteBookRequest{
		BookId: bookID,
	}
	_, err = h.Book.DeleteBook(c, in)
	if err != nil {
		util.ErrorHandling(ctx, err)
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}
