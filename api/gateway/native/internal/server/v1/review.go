package v1

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/calmato/gran-book/api/gateway/native/internal/entity"
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

	c := util.SetMetadata(ctx)
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

	c := util.SetMetadata(ctx)
	reviewsOutput, err := h.bookClient.ListUserReview(c, reviewsInput)
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

	booksOutput, err := h.bookClient.MultiGetBooks(c, booksInput)
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

	c := util.SetMetadata(ctx)
	_, err = h.bookClient.GetBook(c, bookInput)
	if err != nil {
		util.ErrorHandling(ctx, err)
		return
	}

	reviewInput := &pb.GetReviewRequest{
		ReviewId: reviewID,
	}

	reviewOutput, err := h.bookClient.GetReview(c, reviewInput)
	if err != nil {
		util.ErrorHandling(ctx, err)
		return
	}

	userInput := &pb.GetUserRequest{
		UserId: reviewOutput.GetUserId(),
	}

	userOutput, err := h.userClient.GetUser(c, userInput)
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

	c := util.SetMetadata(ctx)
	reviewOutput, err := h.bookClient.GetReview(c, reviewInput)
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

	bookOutput, err := h.bookClient.GetBook(c, bookInput)
	if err != nil {
		util.ErrorHandling(ctx, err)
		return
	}

	res := h.getUserReviewResponse(reviewOutput, bookOutput)
	ctx.JSON(http.StatusOK, res)
}

func (h *reviewHandler) getBookReviewResponse(
	reviewOutput *pb.ReviewResponse, userOutput *pb.UserResponse,
) *pb.BookReviewV1Response {
	user := &pb.BookReviewV1Response_User{
		Id:       reviewOutput.GetUserId(),
		Username: "unknown",
	}

	if userOutput != nil {
		user.Username = userOutput.GetUsername()
		user.ThumbnailUrl = userOutput.GetThumbnailUrl()
	}

	return &pb.BookReviewV1Response{
		Id:         reviewOutput.GetId(),
		Impression: reviewOutput.GetImpression(),
		CreatedAt:  reviewOutput.GetCreatedAt(),
		UpdatedAt:  reviewOutput.GetUpdatedAt(),
		User:       user,
	}
}

func (h *reviewHandler) getUserReviewResponse(
	reviewOutput *pb.ReviewResponse, bookOutput *pb.BookResponse,
) *pb.UserReviewV1Response {
	book := &pb.UserReviewV1Response_Book{
		Id:           bookOutput.GetId(),
		Title:        bookOutput.GetTitle(),
		ThumbnailUrl: bookOutput.GetThumbnailUrl(),
	}

	return &pb.UserReviewV1Response{
		Id:         reviewOutput.GetId(),
		Impression: reviewOutput.GetImpression(),
		CreatedAt:  reviewOutput.GetCreatedAt(),
		UpdatedAt:  reviewOutput.GetUpdatedAt(),
		Book:       book,
	}
}

func (h *reviewHandler) getBookReviewListResponse(
	reviewsOutput *pb.ReviewListResponse, usersOutput *pb.UserMapResponse,
) *pb.BookReviewListV1Response {
	reviews := make([]*pb.BookReviewListV1Response_Review, len(reviewsOutput.GetReviews()))
	for i, r := range reviewsOutput.GetReviews() {
		user := &pb.BookReviewListV1Response_User{
			Id:       r.GetUserId(),
			Username: "unknown",
		}

		if usersOutput.GetUsers()[r.GetUserId()] != nil {
			user.Username = usersOutput.GetUsers()[r.GetUserId()].GetUsername()
			user.ThumbnailUrl = usersOutput.GetUsers()[r.GetUserId()].GetThumbnailUrl()
		}

		review := &pb.BookReviewListV1Response_Review{
			Id:         r.GetId(),
			Impression: r.GetImpression(),
			CreatedAt:  r.GetCreatedAt(),
			UpdatedAt:  r.GetUpdatedAt(),
			User:       user,
		}

		reviews[i] = review
	}

	return &pb.BookReviewListV1Response{
		Reviews: reviews,
		Limit:   reviewsOutput.GetLimit(),
		Offset:  reviewsOutput.GetOffset(),
		Total:   reviewsOutput.GetTotal(),
	}
}

func (h *reviewHandler) getUserReviewListResponse(
	reviewsOutput *pb.ReviewListResponse, booksOutput *pb.BookMapResponse,
) *pb.UserReviewListV1Response {
	reviews := make([]*pb.UserReviewListV1Response_Review, len(reviewsOutput.GetReviews()))
	for i, r := range reviewsOutput.GetReviews() {
		book := &pb.UserReviewListV1Response_Book{
			Id:           booksOutput.GetBooks()[r.GetBookId()].GetId(),
			Title:        booksOutput.GetBooks()[r.GetBookId()].GetTitle(),
			ThumbnailUrl: booksOutput.GetBooks()[r.GetBookId()].GetThumbnailUrl(),
		}

		review := &pb.UserReviewListV1Response_Review{
			Id:         r.GetId(),
			Impression: r.GetImpression(),
			CreatedAt:  r.GetCreatedAt(),
			UpdatedAt:  r.GetUpdatedAt(),
			Book:       book,
		}

		reviews[i] = review
	}

	return &pb.UserReviewListV1Response{
		Reviews: reviews,
		Limit:   reviewsOutput.GetLimit(),
		Offset:  reviewsOutput.GetOffset(),
		Total:   reviewsOutput.GetTotal(),
	}
}
