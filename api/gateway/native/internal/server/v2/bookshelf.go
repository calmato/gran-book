package v2

import (
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

type BookshelfHandler interface {
	List(ctx *gin.Context)
	Get(ctx *gin.Context)
}

type bookshelfHandler struct {
	bookClient pb.BookServiceClient
	userClient pb.UserServiceClient
}

func NewBookshelfHandler(bookConn *grpc.ClientConn, userConn *grpc.ClientConn) BookshelfHandler {
	bc := pb.NewBookServiceClient(bookConn)
	uc := pb.NewUserServiceClient(userConn)

	return &bookshelfHandler{
		bookClient: bc,
		userClient: uc,
	}
}

// List - 本棚の書籍一覧取得
func (h *bookshelfHandler) List(ctx *gin.Context) {
	limit := ctx.GetInt64(ctx.DefaultQuery("limit", entity.ListLimitDefault))
	offset := ctx.GetInt64(ctx.DefaultQuery("offset", entity.ListOffsetDefault))
	userID := ctx.Param("userID")

	in := &pb.ListBookshelfRequest{
		UserId: userID,
		Limit:  limit,
		Offset: offset,
	}

	out, err := h.bookClient.ListBookshelf(ctx, in)
	if err != nil {
		util.ErrorHandling(ctx, err)
		return
	}

	res := h.getBookshelfListResponse(out)
	ctx.JSON(http.StatusOK, res)
}

// Get - 本棚の書籍情報取得
func (h *bookshelfHandler) Get(ctx *gin.Context) {
	userID := ctx.Param("userID")
	bookID, err := strconv.ParseInt(ctx.Param("bookID"), 10, 64)
	if err != nil {
		util.ErrorHandling(ctx, entity.ErrBadRequest.New(err))
		return
	}

	bookshelfInput := &pb.GetBookshelfRequest{
		UserId: userID,
		BookId: bookID,
	}

	bookshelfOutput, err := h.bookClient.GetBookshelf(ctx, bookshelfInput)
	if err != nil {
		util.ErrorHandling(ctx, err)
		return
	}

	reviewsInput := &pb.ListBookReviewRequest{
		BookId: bookshelfOutput.GetBookId(),
		Limit:  20,
		Offset: 0,
	}

	reviewsOutput, err := h.bookClient.ListBookReview(ctx, reviewsInput)
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

	usersOutput, err := h.userClient.MultiGetUser(ctx, usersInput)
	if err != nil {
		util.ErrorHandling(ctx, err)
		return
	}

	res := h.getBookshelfResponse(bookshelfOutput, reviewsOutput, usersOutput)
	ctx.JSON(http.StatusOK, res)
}

func (h *bookshelfHandler) getBookshelfResponse(
	bookshelfOutput *pb.BookshelfResponse, reviewsOutput *pb.ReviewListResponse, usersOutput *pb.UserMapResponse,
) *response.BookshelfResponse {
	bookshelf := &response.BookshelfBookshelf{
		ReviewID:  bookshelfOutput.GetReviewId(),
		Status:    entity.BookshelfStatus(bookshelfOutput.GetStatus()).Name(),
		ReadOn:    bookshelfOutput.GetReadOn(),
		CreatedAt: bookshelfOutput.GetCreatedAt(),
		UpdatedAt: bookshelfOutput.GetUpdatedAt(),
	}

	reviews := make([]*response.BookshelfReview, len(reviewsOutput.GetReviews()))
	for i, r := range reviewsOutput.GetReviews() {
		user := &response.BookshelfUser{
			ID:       r.GetUserId(),
			Username: "unknown",
		}

		if usersOutput.GetUsers()[r.GetUserId()] != nil {
			user.Username = usersOutput.GetUsers()[r.GetUserId()].GetUsername()
			user.ThumbnailURL = usersOutput.GetUsers()[r.GetUserId()].GetThumbnailUrl()
		}

		review := &response.BookshelfReview{
			ID:         r.GetId(),
			Impression: r.GetImpression(),
			CreatedAt:  r.GetCreatedAt(),
			UpdatedAt:  r.GetUpdatedAt(),
			User:       user,
		}

		reviews[i] = review
	}

	authorNames := make([]string, len(bookshelfOutput.GetBook().GetAuthors()))
	authorNameKanas := make([]string, len(bookshelfOutput.GetBook().GetAuthors()))
	for i, a := range bookshelfOutput.GetBook().GetAuthors() {
		authorNames[i] = a.GetName()
		authorNameKanas[i] = a.GetNameKana()
	}

	return &response.BookshelfResponse{
		ID:           bookshelfOutput.GetBook().GetId(),
		Title:        bookshelfOutput.GetBook().GetTitle(),
		TitleKana:    bookshelfOutput.GetBook().GetTitleKana(),
		Description:  bookshelfOutput.GetBook().GetDescription(),
		Isbn:         bookshelfOutput.GetBook().GetIsbn(),
		Publisher:    bookshelfOutput.GetBook().GetPublisher(),
		PublishedOn:  bookshelfOutput.GetBook().GetPublishedOn(),
		ThumbnailURL: bookshelfOutput.GetBook().GetThumbnailUrl(),
		RakutenURL:   bookshelfOutput.GetBook().GetRakutenUrl(),
		Size:         bookshelfOutput.GetBook().GetRakutenSize(),
		Author:       strings.Join(authorNames, "/"),
		AuthorKana:   strings.Join(authorNameKanas, "/"),
		CreatedAt:    bookshelfOutput.GetBook().GetCreatedAt(),
		UpdatedAt:    bookshelfOutput.GetBook().GetUpdatedAt(),
		Bookshelf:    bookshelf,
		Reviews:      reviews,
		ReviewLimit:  reviewsOutput.GetLimit(),
		ReviewOffset: reviewsOutput.GetOffset(),
		ReviewTotal:  reviewsOutput.GetTotal(),
	}
}

func (h *bookshelfHandler) getBookshelfListResponse(out *pb.BookshelfListResponse) *response.BookshelfListResponse {
	books := make([]*response.BookshelfListBook, len(out.GetBookshelves()))
	for i, b := range out.GetBookshelves() {
		bookshelf := &response.BookshelfListBookshelf{
			Status:    entity.BookshelfStatus(b.GetStatus()).Name(),
			ReadOn:    b.GetReadOn(),
			ReviewID:  b.GetReviewId(),
			CreatedAt: b.GetCreatedAt(),
			UpdatedAt: b.GetUpdatedAt(),
		}

		authorNames := make([]string, len(b.GetBook().GetAuthors()))
		authorNameKanas := make([]string, len(b.GetBook().GetAuthors()))
		for i, a := range b.GetBook().GetAuthors() {
			authorNames[i] = a.GetName()
			authorNameKanas[i] = a.GetNameKana()
		}

		book := &response.BookshelfListBook{
			ID:           b.GetBook().GetId(),
			Title:        b.GetBook().GetTitle(),
			TitleKana:    b.GetBook().GetTitleKana(),
			Description:  b.GetBook().GetDescription(),
			Isbn:         b.GetBook().GetIsbn(),
			Publisher:    b.GetBook().GetPublisher(),
			PublishedOn:  b.GetBook().GetPublishedOn(),
			ThumbnailURL: b.GetBook().GetThumbnailUrl(),
			RakutenURL:   b.GetBook().GetRakutenUrl(),
			Size:         b.GetBook().GetRakutenSize(),
			Author:       strings.Join(authorNames, "/"),
			AuthorKana:   strings.Join(authorNameKanas, "/"),
			CreatedAt:    b.GetBook().GetCreatedAt(),
			UpdatedAt:    b.GetBook().GetUpdatedAt(),
			Bookshelf:    bookshelf,
		}

		books[i] = book
	}

	return &response.BookshelfListResponse{
		Books:  books,
		Limit:  out.GetLimit(),
		Offset: out.GetOffset(),
		Total:  out.GetTotal(),
	}
}
