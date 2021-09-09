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
	"golang.org/x/sync/errgroup"
)

type ReviewHandler interface {
	ListByBook(ctx *gin.Context)
	ListByUser(ctx *gin.Context)
	GetByBook(ctx *gin.Context)
	GetByUser(ctx *gin.Context)
}

type reviewHandler struct {
	authClient pb.AuthServiceClient
	bookClient pb.BookServiceClient
	userClient pb.UserServiceClient
}

func NewReviewHandler(ac pb.AuthServiceClient, bc pb.BookServiceClient, uc pb.UserServiceClient) ReviewHandler {
	return &reviewHandler{
		authClient: ac,
		bookClient: bc,
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

	rs := entity.NewReviews(reviewsOutput.Reviews)

	usersInput := &pb.MultiGetUserRequest{
		UserIds: rs.UserIDs(),
	}

	usersOutput, err := h.userClient.MultiGetUser(c, usersInput)
	if err != nil {
		util.ErrorHandling(ctx, err)
		return
	}

	us := entity.NewUsers(usersOutput.Users)
	res := h.getBookReviewListResponse(rs, us.Map(), reviewsOutput.Limit, reviewsOutput.Offset, reviewsOutput.Total)
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

	rs := entity.NewReviews(reviewsOutput.Reviews)

	booksInput := &pb.MultiGetBooksRequest{
		BookIds: rs.BookIDs(),
	}

	booksOutput, err := h.bookClient.MultiGetBooks(c, booksInput)
	if err != nil {
		util.ErrorHandling(ctx, err)
		return
	}

	bs := entity.NewBooks(booksOutput.Books)
	res := h.getUserReviewListResponse(rs, bs.Map(), reviewsOutput.Limit, reviewsOutput.Offset, reviewsOutput.Total)
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

	c := util.SetMetadata(ctx)
	eg, ectx := errgroup.WithContext(c)

	eg.Go(func() error {
		bookInput := &pb.GetBookRequest{
			BookId: bookID,
		}
		_, err = h.bookClient.GetBook(ectx, bookInput)
		return err
	})

	var r *entity.Review
	eg.Go(func() error {
		reviewInput := &pb.GetReviewRequest{
			ReviewId: reviewID,
		}
		reviewOutput, err := h.bookClient.GetReview(ectx, reviewInput)
		if err != nil {
			return err
		}
		r = entity.NewReview(reviewOutput.Review)
		return nil
	})

	if err := eg.Wait(); err != nil {
		util.ErrorHandling(ctx, err)
		return
	}

	userInput := &pb.GetUserRequest{
		UserId: r.UserId,
	}

	userOutput, err := h.userClient.GetUser(c, userInput)
	if err != nil && !util.IsNotFound(err) {
		util.ErrorHandling(ctx, err)
		return
	}

	u := entity.NewUser(userOutput.User)
	res := h.getBookReviewResponse(r, u)
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

	r := entity.NewReview(reviewOutput.Review)
	if r.UserId != userID {
		err := fmt.Errorf("user id is invalid")
		util.ErrorHandling(ctx, entity.ErrBadRequest.New(err))
		return
	}

	bookInput := &pb.GetBookRequest{
		BookId: r.BookId,
	}

	bookOutput, err := h.bookClient.GetBook(c, bookInput)
	if err != nil {
		util.ErrorHandling(ctx, err)
		return
	}

	b := entity.NewBook(bookOutput.Book)
	res := h.getUserReviewResponse(r, b)
	ctx.JSON(http.StatusOK, res)
}

func (h *reviewHandler) getBookReviewResponse(r *entity.Review, u *entity.User) *response.BookReviewResponse {
	user := &response.BookReviewUser{
		ID:           r.UserId,
		Username:     "unknown",
		ThumbnailURL: "",
	}

	if u != nil {
		user.Username = u.Username
		user.ThumbnailURL = u.ThumbnailUrl
	}

	return &response.BookReviewResponse{
		ID:         r.Id,
		Impression: r.Impression,
		CreatedAt:  r.CreatedAt,
		UpdatedAt:  r.UpdatedAt,
		User:       user,
	}
}

func (h *reviewHandler) getUserReviewResponse(r *entity.Review, b *entity.Book) *response.UserReviewResponse {
	book := &response.UserReviewBook{
		ID:           b.Id,
		Title:        b.Title,
		ThumbnailURL: b.ThumbnailUrl,
	}

	return &response.UserReviewResponse{
		ID:         r.Id,
		Impression: r.Impression,
		CreatedAt:  r.CreatedAt,
		UpdatedAt:  r.UpdatedAt,
		Book:       book,
	}
}

func (h *reviewHandler) getBookReviewListResponse(
	rs entity.Reviews, us map[string]*entity.User, limit, offset, total int64,
) *response.BookReviewListResponse {
	reviews := make([]*response.BookReviewListReview, len(rs))
	for i, r := range rs {
		user := &response.BookReviewListUser{
			ID:           r.UserId,
			Username:     "unknown",
			ThumbnailURL: "",
		}

		if us[r.UserId] != nil {
			user.Username = us[r.UserId].Username
			user.ThumbnailURL = us[r.UserId].ThumbnailUrl
		}

		review := &response.BookReviewListReview{
			ID:         r.Id,
			Impression: r.Impression,
			CreatedAt:  r.CreatedAt,
			UpdatedAt:  r.UpdatedAt,
			User:       user,
		}

		reviews[i] = review
	}

	return &response.BookReviewListResponse{
		Reviews: reviews,
		Limit:   limit,
		Offset:  offset,
		Total:   total,
	}
}

func (h *reviewHandler) getUserReviewListResponse(
	rs entity.Reviews, bs map[int64]*entity.Book, limit, offset, total int64,
) *response.UserReviewListResponse {
	reviews := make([]*response.UserReviewListReview, 0, len(rs))
	for i, r := range rs {
		b, ok := bs[r.BookId]
		if !ok {
			continue
		}

		book := &response.UserReviewListBook{
			ID:           b.Id,
			Title:        b.Title,
			ThumbnailURL: b.ThumbnailUrl,
		}

		review := &response.UserReviewListReview{
			ID:         r.Id,
			Impression: r.Impression,
			CreatedAt:  r.CreatedAt,
			UpdatedAt:  r.UpdatedAt,
			Book:       book,
		}

		reviews[i] = review
	}

	return &response.UserReviewListResponse{
		Reviews: reviews,
		Limit:   limit,
		Offset:  offset,
		Total:   total,
	}
}
