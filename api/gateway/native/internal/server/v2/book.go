package v2

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/calmato/gran-book/api/gateway/native/internal/entity"
	response "github.com/calmato/gran-book/api/gateway/native/internal/response/v2"
	"github.com/calmato/gran-book/api/gateway/native/internal/server/util"
	"github.com/calmato/gran-book/api/gateway/native/proto/service/book"
	"github.com/calmato/gran-book/api/gateway/native/proto/service/user"
	"github.com/gin-gonic/gin"
)

type BookHandler interface {
	Get(ctx *gin.Context)
}

type bookHandler struct {
	authClient user.AuthServiceClient
	bookClient book.BookServiceClient
	userClient user.UserServiceClient
}

func NewBookHandler(ac user.AuthServiceClient, bc book.BookServiceClient, uc user.UserServiceClient) BookHandler {
	return &bookHandler{
		authClient: ac,
		bookClient: bc,
		userClient: uc,
	}
}

// Get - 書籍情報取得
func (h *bookHandler) Get(ctx *gin.Context) {
	bookID := ctx.Param("bookID")
	key := ctx.DefaultQuery("key", "id")

	c := util.SetMetadata(ctx)
	bookOutput, err := h.getBook(c, bookID, key)
	if err != nil {
		util.ErrorHandling(ctx, err)
		return
	}

	b := entity.NewBook(bookOutput.Book)

	reviewsInput := &book.ListBookReviewRequest{
		BookId: b.Id,
		Limit:  20,
		Offset: 0,
	}

	reviewsOutput, err := h.bookClient.ListBookReview(c, reviewsInput)
	if err != nil {
		util.ErrorHandling(ctx, err)
		return
	}

	rs := entity.NewReviews(reviewsOutput.Reviews)

	usersInput := &user.MultiGetUserRequest{
		UserIds: rs.UserIDs(),
	}

	usersOutput, err := h.userClient.MultiGetUser(c, usersInput)
	if err != nil {
		util.ErrorHandling(ctx, err)
		return
	}

	us := entity.NewUsers(usersOutput.Users)
	res := h.getBookResponse(b, rs, us.Map(), reviewsOutput.Limit, reviewsOutput.Offset, reviewsOutput.Total)
	ctx.JSON(http.StatusOK, res)
}

func (h *bookHandler) getBook(ctx context.Context, bookID, key string) (*book.BookResponse, error) {
	switch key {
	case "id":
		bookID, err := strconv.ParseInt(bookID, 10, 64)
		if err != nil {
			return nil, entity.ErrBadRequest.New(err)
		}
		return h.bookClient.GetBook(ctx, &book.GetBookRequest{BookId: bookID})
	case "isbn":
		return h.bookClient.GetBookByIsbn(ctx, &book.GetBookByIsbnRequest{Isbn: bookID})
	default:
		err := fmt.Errorf("this key is invalid argument")
		return nil, entity.ErrBadRequest.New(err)
	}
}

func (h *bookHandler) getBookResponse(
	b *entity.Book,
	rs entity.Reviews,
	us map[string]*entity.User,
	reviewLimit int64,
	reviewOffset int64,
	reviewTotal int64,
) *response.BookResponse {
	reviews := make([]*response.BookReview, len(rs))
	for i, r := range rs {
		user := &response.BookUser{
			ID:           r.UserId,
			Username:     "unknown",
			ThumbnailURL: "",
		}

		if us[r.UserId] != nil {
			user.Username = us[r.UserId].Username
			user.ThumbnailURL = us[r.UserId].ThumbnailUrl
		}

		review := &response.BookReview{
			ID:         r.Id,
			Impression: r.Impression,
			CreatedAt:  r.CreatedAt,
			UpdatedAt:  r.UpdatedAt,
			User:       user,
		}

		reviews[i] = review
	}

	return &response.BookResponse{
		ID:           b.Id,
		Title:        b.Title,
		TitleKana:    b.TitleKana,
		Description:  b.Description,
		Isbn:         b.Isbn,
		Publisher:    b.Publisher,
		PublishedOn:  b.PublishedOn,
		ThumbnailURL: b.ThumbnailUrl,
		RakutenURL:   b.RakutenUrl,
		Size:         b.RakutenSize,
		Author:       strings.Join(b.AuthorNames(), "/"),
		AuthorKana:   strings.Join(b.AuthorNameKanas(), "/"),
		CreatedAt:    b.CreatedAt,
		UpdatedAt:    b.UpdatedAt,
		Reviews:      reviews,
		ReviewLimit:  reviewLimit,
		ReviewOffset: reviewOffset,
		ReviewTotal:  reviewTotal,
	}
}
