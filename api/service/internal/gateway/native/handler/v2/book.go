package v2

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	gentity "github.com/calmato/gran-book/api/service/internal/gateway/entity"
	response "github.com/calmato/gran-book/api/service/internal/gateway/native/response/v2"
	"github.com/calmato/gran-book/api/service/internal/gateway/util"
	"github.com/calmato/gran-book/api/service/pkg/exception"
	"github.com/calmato/gran-book/api/service/proto/book"
	"github.com/calmato/gran-book/api/service/proto/user"
	"github.com/gin-gonic/gin"
)

// getBook - 書籍情報取得
func (h *apiV2Handler) getBook(ctx *gin.Context) {
	c := util.SetMetadata(ctx)
	bookID := ctx.Param("bookID")
	key := ctx.DefaultQuery("key", "id")

	b, err := h.getBookByKey(c, bookID, key)
	if err != nil {
		util.ErrorHandling(ctx, err)
		return
	}

	reviewsInput := &book.ListBookReviewRequest{
		BookId: b.Id,
		Limit:  20,
		Offset: 0,
	}
	reviewsOutput, err := h.Book.ListBookReview(c, reviewsInput)
	if err != nil {
		util.ErrorHandling(ctx, err)
		return
	}

	rs := gentity.NewReviews(reviewsOutput.Reviews)

	usersInput := &user.MultiGetUserRequest{
		UserIds: rs.UserIDs(),
	}
	usersOutput, err := h.User.MultiGetUser(c, usersInput)
	if err != nil {
		util.ErrorHandling(ctx, err)
		return
	}

	us := gentity.NewUsers(usersOutput.Users)
	res := response.NewBookResponse(b, rs, us.Map(), reviewsOutput.Limit, reviewsOutput.Offset, reviewsOutput.Total)
	ctx.JSON(http.StatusOK, res)
}

func (h *apiV2Handler) getBookByKey(ctx context.Context, bookID, key string) (*gentity.Book, error) {
	switch key {
	case "id":
		bookID, err := strconv.ParseInt(bookID, 10, 64)
		if err != nil {
			return nil, exception.ErrInvalidArgument.New(err)
		}
		out, err := h.Book.GetBook(ctx, &book.GetBookRequest{BookId: bookID})
		if err != nil {
			return nil, err
		}
		return gentity.NewBook(out.Book), nil
	case "isbn":
		out, err := h.Book.GetBookByIsbn(ctx, &book.GetBookByIsbnRequest{Isbn: bookID})
		if err != nil {
			return nil, err
		}
		return gentity.NewBook(out.Book), nil
	default:
		err := fmt.Errorf("this key is invalid argument")
		return nil, exception.ErrInvalidArgument.New(err)
	}
}
