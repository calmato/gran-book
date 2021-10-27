package v2

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	gentity "github.com/calmato/gran-book/api/internal/gateway/entity"
	"github.com/calmato/gran-book/api/internal/gateway/native/entity"
	response "github.com/calmato/gran-book/api/internal/gateway/native/response/v2"
	"github.com/calmato/gran-book/api/internal/gateway/util"
	"github.com/calmato/gran-book/api/pkg/exception"
	"github.com/calmato/gran-book/api/proto/book"
	"github.com/gin-gonic/gin"
)

// getBook - 書籍情報取得
func (h *apiV2Handler) getBook(ctx *gin.Context) {
	c := util.SetMetadata(ctx)

	limit, _ := strconv.ParseInt(entity.ListLimitDefault, 10, 64)
	offset, _ := strconv.ParseInt(entity.ListOffsetDefault, 10, 64)

	bookID := ctx.Param("bookID")
	key := ctx.DefaultQuery("key", "id")

	b, err := h.bookGetBookByKey(c, bookID, key)
	if err != nil {
		util.ErrorHandling(ctx, err)
		return
	}

	rs, total, err := h.bookListBookReview(c, b.Id, limit, offset)
	if err != nil {
		util.ErrorHandling(ctx, err)
		return
	}

	us, err := h.userMultiGetUser(c, rs.UserIDs())
	if err != nil {
		util.ErrorHandling(ctx, err)
		return
	}

	reviews := entity.NewBookReviews(rs, us.Map())
	book := entity.NewBook(b, reviews)
	book.Fill(limit, offset, total)
	res := &response.BookResponse{
		Book: book,
	}
	ctx.JSON(http.StatusOK, res)
}

func (h *apiV2Handler) bookGetBookByKey(ctx context.Context, bookID, key string) (*gentity.Book, error) {
	switch key {
	case "id":
		bookID, err := strconv.ParseInt(bookID, 10, 64)
		if err != nil {
			return nil, exception.ErrInvalidArgument.New(err)
		}
		return h.bookGetBook(ctx, bookID)
	case "isbn":
		out, err := h.Book.GetBookByIsbn(ctx, &book.GetBookByIsbnRequest{Isbn: bookID})
		if err != nil {
			return nil, err
		}
		if out.Book == nil {
			err := fmt.Errorf("book is not found: %s", bookID)
			return nil, exception.ErrNotFound.New(err)
		}
		return gentity.NewBook(out.Book), nil
	default:
		err := fmt.Errorf("this key is invalid argument")
		return nil, exception.ErrInvalidArgument.New(err)
	}
}

func (h *apiV2Handler) bookMultiGetBooks(ctx context.Context, bookIDs []int64) (gentity.Books, error) {
	in := &book.MultiGetBooksRequest{
		BookIds: bookIDs,
	}
	out, err := h.Book.MultiGetBooks(ctx, in)
	if err != nil {
		return nil, err
	}

	return gentity.NewBooks(out.Books), nil
}

func (h *apiV2Handler) bookGetBook(ctx context.Context, bookID int64) (*gentity.Book, error) {
	in := &book.GetBookRequest{
		BookId: bookID,
	}
	out, err := h.Book.GetBook(ctx, in)
	if err != nil {
		return nil, err
	}
	if out.Book == nil {
		err := fmt.Errorf("book is not found: %d", bookID)
		return nil, exception.ErrNotFound.New(err)
	}

	return gentity.NewBook(out.Book), nil
}

func (h *apiV2Handler) bookListBookReview(
	ctx context.Context, bookID, limit, offset int64,
) (gentity.Reviews, int64, error) {
	in := &book.ListBookReviewRequest{
		BookId: bookID,
		Limit:  limit,
		Offset: offset,
	}
	out, err := h.Book.ListBookReview(ctx, in)
	if err != nil {
		return nil, 0, err
	}

	return gentity.NewReviews(out.Reviews), out.Total, nil
}
