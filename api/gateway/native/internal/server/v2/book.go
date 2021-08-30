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
	pb "github.com/calmato/gran-book/api/gateway/native/proto"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

type BookHandler interface {
	Get(ctx *gin.Context)
}

type bookHandler struct {
	bookClient pb.BookServiceClient
	authClient pb.AuthServiceClient
	userClient pb.UserServiceClient
}

func NewBookHandler(bookConn *grpc.ClientConn, authConn *grpc.ClientConn, userConn *grpc.ClientConn) BookHandler {
	bc := pb.NewBookServiceClient(bookConn)
	ac := pb.NewAuthServiceClient(authConn)
	uc := pb.NewUserServiceClient(userConn)

	return &bookHandler{
		bookClient: bc,
		authClient: ac,
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

	reviewsInput := &pb.ListBookReviewRequest{
		BookId: bookOutput.GetId(),
		Limit:  20,
		Offset: 0,
	}

	reviewsOutput, err := h.bookClient.ListBookReview(c, reviewsInput)
	if err != nil {
		util.ErrorHandling(ctx, err)
		return
	}

	userIDs := make([]string, len(reviewsOutput.GetReviews()))
	for i, r := range reviewsOutput.GetReviews() {
		userIDs[i] = r.GetUserId()
	}

	usersInput := &pb.MultiGetUserRequest{
		UserIds: userIDs,
	}

	usersOutput, err := h.userClient.MultiGetUser(c, usersInput)
	if err != nil {
		util.ErrorHandling(ctx, err)
		return
	}

	us := entity.NewUsers(usersOutput.Users)

	res := h.getBookResponse(bookOutput, reviewsOutput, us.Map())
	ctx.JSON(http.StatusOK, res)
}

func (h *bookHandler) getBook(ctx context.Context, bookID, key string) (*pb.BookResponse, error) {
	switch key {
	case "id":
		bookID, err := strconv.ParseInt(bookID, 10, 64)
		if err != nil {
			return nil, entity.ErrBadRequest.New(err)
		}

		in := &pb.GetBookRequest{
			BookId: bookID,
		}

		return h.bookClient.GetBook(ctx, in)
	case "isbn":
		in := &pb.GetBookByIsbnRequest{
			Isbn: bookID,
		}

		return h.bookClient.GetBookByIsbn(ctx, in)
	default:
		err := fmt.Errorf("this key is invalid argument")
		return nil, entity.ErrBadRequest.New(err)
	}
}

func (h *bookHandler) getBookResponse(
	bookOutput *pb.BookResponse,
	reviewsOutput *pb.ReviewListResponse,
	us map[string]*entity.User,
) *response.BookResponse {
	authorNames := make([]string, len(bookOutput.GetAuthors()))
	authorNameKanas := make([]string, len(bookOutput.GetAuthors()))
	for i, a := range bookOutput.GetAuthors() {
		authorNames[i] = a.GetName()
		authorNameKanas[i] = a.GetNameKana()
	}

	reviews := make([]*response.BookResponse_Review, len(reviewsOutput.GetReviews()))
	for i, r := range reviewsOutput.GetReviews() {
		user := &response.BookResponse_User{
			ID:       r.GetUserId(),
			Username: "unknown",
		}

		if us[r.GetUserId()] != nil {
			user.Username = us[r.GetUserId()].Username
			user.ThumbnailURL = us[r.GetUserId()].ThumbnailUrl
		}

		review := &response.BookResponse_Review{
			ID:         r.GetId(),
			Impression: r.GetImpression(),
			CreatedAt:  r.GetCreatedAt(),
			UpdatedAt:  r.GetUpdatedAt(),
			User:       user,
		}

		reviews[i] = review
	}

	return &response.BookResponse{
		ID:           bookOutput.GetId(),
		Title:        bookOutput.GetTitle(),
		TitleKana:    bookOutput.GetTitleKana(),
		Description:  bookOutput.GetDescription(),
		Isbn:         bookOutput.GetIsbn(),
		Publisher:    bookOutput.GetPublisher(),
		PublishedOn:  bookOutput.GetPublishedOn(),
		ThumbnailURL: bookOutput.GetThumbnailUrl(),
		RakutenURL:   bookOutput.GetRakutenUrl(),
		Size:         bookOutput.GetRakutenSize(),
		Author:       strings.Join(authorNames, "/"),
		AuthorKana:   strings.Join(authorNameKanas, "/"),
		CreatedAt:    bookOutput.GetCreatedAt(),
		UpdatedAt:    bookOutput.GetUpdatedAt(),
		Reviews:      reviews,
		ReviewLimit:  reviewsOutput.GetLimit(),
		ReviewOffset: reviewsOutput.GetOffset(),
		ReviewTotal:  reviewsOutput.GetTotal(),
	}
}
