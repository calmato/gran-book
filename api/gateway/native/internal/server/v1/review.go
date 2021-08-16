package v1

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/calmato/gran-book/api/gateway/native/internal/entity"
	response "github.com/calmato/gran-book/api/gateway/native/internal/response/v1"
	"github.com/calmato/gran-book/api/gateway/native/internal/server/util"
	pb "github.com/calmato/gran-book/api/gateway/native/proto"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

type ReviewHandler interface {
	ListByBook(ctx *gin.Context)
	ListByUser(ctx *gin.Context)
	GetByBook(ctx *gin.Context)
	GetByUser(ctx *gin.Context)
}

type reviewHandler struct {
	bookClient pb.BookServiceClient
	authClient pb.AuthServiceClient
	userClient pb.UserServiceClient
}

func NewReviewHandler(bookConn *grpc.ClientConn, authConn *grpc.ClientConn, userConn *grpc.ClientConn) ReviewHandler {
	bc := pb.NewBookServiceClient(bookConn)
	ac := pb.NewAuthServiceClient(authConn)
	uc := pb.NewUserServiceClient(authConn)

	return &reviewHandler{
		bookClient: bc,
		authClient: ac,
		userClient: uc,
	}
}

// ListByBook - 書籍のレビュー一覧取得
func (h *reviewHandler) ListByBook(ctx *gin.Context) {
	limit := ctx.GetInt64(ctx.DefaultQuery("limit", entity.ListLimitDefault))
	offset := ctx.GetInt64(ctx.DefaultQuery("offset", entity.ListOffsetDefault))
	bookID, err := strconv.ParseInt(ctx.Param("bookID"), 10, 64)
	if err != nil {
		util.ErrorHandling(ctx, entity.ErrBadRequest.New(err))
		return
	}

	reviewsInput := &pb.ListBookReviewRequest{
		BookId: bookID,
		Limit:  limit,
		Offset: offset,
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

	res := h.getBookReviewListResponse(reviewsOutput, usersOutput)
	ctx.JSON(http.StatusOK, res)
}

// ListByUser - ユーザーのレビュー一覧取得
func (h *reviewHandler) ListByUser(ctx *gin.Context) {
	userID := ctx.Param("userID")
	limit := ctx.GetInt64(ctx.DefaultQuery("limit", entity.ListLimitDefault))
	offset := ctx.GetInt64(ctx.DefaultQuery("offset", entity.ListOffsetDefault))

	reviewsInput := &pb.ListUserReviewRequest{
		UserId: userID,
		Limit:  limit,
		Offset: offset,
	}

	reviewsOutput, err := h.bookClient.ListUserReview(ctx, reviewsInput)
	if err != nil {
		util.ErrorHandling(ctx, err)
		return
	}

	bookIDs := make([]int64, len(reviewsOutput.GetReviews()))
	for i, r := range reviewsOutput.GetReviews() {
		bookIDs[i] = r.GetBookId()
	}

	booksInput := &pb.MultiGetBooksRequest{
		BookIds: bookIDs,
	}

	booksOutput, err := h.bookClient.MultiGetBooks(ctx, booksInput)
	if err != nil {
		util.ErrorHandling(ctx, err)
		return
	}

	res := h.getUserReviewListResponse(reviewsOutput, booksOutput)
	ctx.JSON(http.StatusOK, res)
}

// GetByBook - 書籍のレビュー情報取得
func (h *reviewHandler) GetByBook(ctx *gin.Context) {
	bookID, err := strconv.ParseInt(ctx.Param("bookID"), 10, 64)
	if err != nil {
		util.ErrorHandling(ctx, entity.ErrBadRequest.New(err))
		return
	}

	reviewID, err := strconv.ParseInt(ctx.Param("reviewID"), 10, 64)
	if err != nil {
		util.ErrorHandling(ctx, entity.ErrBadRequest.New(err))
		return
	}

	bookInput := &pb.GetBookRequest{
		BookId: bookID,
	}

	_, err = h.bookClient.GetBook(ctx, bookInput)
	if err != nil {
		util.ErrorHandling(ctx, err)
		return
	}

	reviewInput := &pb.GetReviewRequest{
		ReviewId: reviewID,
	}

	reviewOutput, err := h.bookClient.GetReview(ctx, reviewInput)
	if err != nil {
		util.ErrorHandling(ctx, err)
		return
	}

	userInput := &pb.GetUserRequest{
		UserId: reviewOutput.GetUserId(),
	}

	userOutput, err := h.userClient.GetUser(ctx, userInput)
	if err != nil && !util.IsNotFound(err) {
		util.ErrorHandling(ctx, err)
		return
	}

	res := h.getBookReviewResponse(reviewOutput, userOutput)
	ctx.JSON(http.StatusOK, res)
}

// GetByUser - ユーザーのレビュー情報取得
func (h *reviewHandler) GetByUser(ctx *gin.Context) {
	userID := ctx.Param("userID")
	reviewID, err := strconv.ParseInt(ctx.Param("reviewID"), 10, 64)
	if err != nil {
		util.ErrorHandling(ctx, entity.ErrBadRequest.New(err))
		return
	}

	reviewInput := &pb.GetReviewRequest{
		ReviewId: reviewID,
	}

	reviewOutput, err := h.bookClient.GetReview(ctx, reviewInput)
	if err != nil {
		util.ErrorHandling(ctx, err)
		return
	}

	if reviewOutput.GetUserId() != userID {
		err := fmt.Errorf("user id is invalid")
		util.ErrorHandling(ctx, entity.ErrBadRequest.New(err))
		return
	}

	bookInput := &pb.GetBookRequest{
		BookId: reviewOutput.GetBookId(),
	}

	bookOutput, err := h.bookClient.GetBook(ctx, bookInput)
	if err != nil {
		util.ErrorHandling(ctx, err)
		return
	}

	res := h.getUserReviewResponse(reviewOutput, bookOutput)
	ctx.JSON(http.StatusOK, res)
}

func (h *reviewHandler) getBookReviewResponse(
	reviewOutput *pb.ReviewResponse, userOutput *pb.UserResponse,
) *response.BookReviewResponse {
	user := &response.BookReviewUser{
		ID:       reviewOutput.GetUserId(),
		Username: "unknown",
	}

	if userOutput != nil {
		user.Username = userOutput.GetUsername()
		user.ThumbnailURL = userOutput.GetThumbnailUrl()
	}

	return &response.BookReviewResponse{
		ID:         reviewOutput.GetId(),
		Impression: reviewOutput.GetImpression(),
		CreatedAt:  reviewOutput.GetCreatedAt(),
		UpdatedAt:  reviewOutput.GetUpdatedAt(),
		User:       user,
	}
}

func (h *reviewHandler) getUserReviewResponse(
	reviewOutput *pb.ReviewResponse, bookOutput *pb.BookResponse,
) *response.UserReviewResponse {
	book := &response.UserReviewBook{
		ID:           bookOutput.GetId(),
		Title:        bookOutput.GetTitle(),
		ThumbnailURL: bookOutput.GetThumbnailUrl(),
	}

	return &response.UserReviewResponse{
		ID:         reviewOutput.GetId(),
		Impression: reviewOutput.GetImpression(),
		CreatedAt:  reviewOutput.GetCreatedAt(),
		UpdatedAt:  reviewOutput.GetUpdatedAt(),
		Book:       book,
	}
}

func (h *reviewHandler) getBookReviewListResponse(
	reviewsOutput *pb.ReviewListResponse, usersOutput *pb.UserMapResponse,
) *response.BookReviewListResponse {
	reviews := make([]*response.BookReviewListReview, len(reviewsOutput.GetReviews()))
	for i, r := range reviewsOutput.GetReviews() {
		user := &response.BookReviewListUser{
			ID:       r.GetUserId(),
			Username: "unknown",
		}

		if usersOutput.GetUsers()[r.GetUserId()] != nil {
			user.Username = usersOutput.GetUsers()[r.GetUserId()].GetUsername()
			user.ThumbnailURL = usersOutput.GetUsers()[r.GetUserId()].GetThumbnailUrl()
		}

		review := &response.BookReviewListReview{
			ID:         r.GetId(),
			Impression: r.GetImpression(),
			CreatedAt:  r.GetCreatedAt(),
			UpdatedAt:  r.GetUpdatedAt(),
			User:       user,
		}

		reviews[i] = review
	}

	return &response.BookReviewListResponse{
		Reviews: reviews,
		Limit:   reviewsOutput.GetLimit(),
		Offset:  reviewsOutput.GetOffset(),
		Total:   reviewsOutput.GetTotal(),
	}
}

func (h *reviewHandler) getUserReviewListResponse(
	reviewsOutput *pb.ReviewListResponse, booksOutput *pb.BookMapResponse,
) *response.UserReviewListResponse {
	reviews := make([]*response.UserReviewListReview, len(reviewsOutput.GetReviews()))
	for i, r := range reviewsOutput.GetReviews() {
		book := &response.UserReviewListBook{
			ID:           booksOutput.GetBooks()[r.GetBookId()].GetId(),
			Title:        booksOutput.GetBooks()[r.GetBookId()].GetTitle(),
			ThumbnailURL: booksOutput.GetBooks()[r.GetBookId()].GetThumbnailUrl(),
		}

		review := &response.UserReviewListReview{
			ID:         r.GetId(),
			Impression: r.GetImpression(),
			CreatedAt:  r.GetCreatedAt(),
			UpdatedAt:  r.GetUpdatedAt(),
			Book:       book,
		}

		reviews[i] = review
	}

	return &response.UserReviewListResponse{
		Reviews: reviews,
		Limit:   reviewsOutput.GetLimit(),
		Offset:  reviewsOutput.GetOffset(),
		Total:   reviewsOutput.GetTotal(),
	}
}
