package v1

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/calmato/gran-book/api/gateway/native/internal/entity"
	response "github.com/calmato/gran-book/api/gateway/native/internal/response/v1"
	"github.com/calmato/gran-book/api/gateway/native/internal/server/util"
	"github.com/calmato/gran-book/api/gateway/native/proto/service/book"
	"github.com/calmato/gran-book/api/gateway/native/proto/service/user"
	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
)

// listReviewByBook - 書籍のレビュー一覧取得
func (h *apiV1Handler) listReviewByBook(ctx *gin.Context) {
	c := util.SetMetadata(ctx)
	limit, err := strconv.ParseInt(ctx.DefaultQuery("limit", entity.ListLimitDefault), 10, 64)
	if err != nil {
		util.ErrorHandling(ctx, entity.ErrBadRequest.New(err))
		return
	}
	offset, err := strconv.ParseInt(ctx.DefaultQuery("offset", entity.ListOffsetDefault), 10, 64)
	if err != nil {
		util.ErrorHandling(ctx, entity.ErrBadRequest.New(err))
		return
	}
	bookID, err := strconv.ParseInt(ctx.Param("bookID"), 10, 64)
	if err != nil {
		util.ErrorHandling(ctx, entity.ErrBadRequest.New(err))
		return
	}

	reviewsInput := &book.ListBookReviewRequest{
		BookId: bookID,
		Limit:  limit,
		Offset: offset,
	}
	reviewsOutput, err := h.Book.ListBookReview(c, reviewsInput)
	if err != nil {
		util.ErrorHandling(ctx, err)
		return
	}

	rs := entity.NewReviews(reviewsOutput.Reviews)

	usersInput := &user.MultiGetUserRequest{
		UserIds: rs.UserIDs(),
	}
	usersOutput, err := h.User.MultiGetUser(c, usersInput)
	if err != nil {
		util.ErrorHandling(ctx, err)
		return
	}

	us := entity.NewUsers(usersOutput.Users)
	res := response.NewBookReviewListResponse(
		rs, us.Map(), reviewsOutput.Limit, reviewsOutput.Offset, reviewsOutput.Total,
	)
	ctx.JSON(http.StatusOK, res)
}

// listReviewByUser - ユーザーのレビュー一覧取得
func (h *apiV1Handler) listReviewByUser(ctx *gin.Context) {
	c := util.SetMetadata(ctx)
	userID := ctx.Param("userID")
	limit, err := strconv.ParseInt(ctx.DefaultQuery("limit", entity.ListLimitDefault), 10, 64)
	if err != nil {
		util.ErrorHandling(ctx, entity.ErrBadRequest.New(err))
		return
	}
	offset, err := strconv.ParseInt(ctx.DefaultQuery("offset", entity.ListOffsetDefault), 10, 64)
	if err != nil {
		util.ErrorHandling(ctx, entity.ErrBadRequest.New(err))
		return
	}

	reviewsInput := &book.ListUserReviewRequest{
		UserId: userID,
		Limit:  limit,
		Offset: offset,
	}
	reviewsOutput, err := h.Book.ListUserReview(c, reviewsInput)
	if err != nil {
		util.ErrorHandling(ctx, err)
		return
	}

	rs := entity.NewReviews(reviewsOutput.Reviews)

	booksInput := &book.MultiGetBooksRequest{
		BookIds: rs.BookIDs(),
	}
	booksOutput, err := h.Book.MultiGetBooks(c, booksInput)
	if err != nil {
		util.ErrorHandling(ctx, err)
		return
	}

	bs := entity.NewBooks(booksOutput.Books)
	res := response.NewUserReviewListResponse(
		rs, bs.Map(), reviewsOutput.Limit, reviewsOutput.Offset, reviewsOutput.Total,
	)
	ctx.JSON(http.StatusOK, res)
}

// getBookReview - 書籍のレビュー情報取得
func (h *apiV1Handler) getBookReview(ctx *gin.Context) {
	c := util.SetMetadata(ctx)
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

	eg, ectx := errgroup.WithContext(c)

	eg.Go(func() error {
		in := &book.GetBookRequest{
			BookId: bookID,
		}
		_, err = h.Book.GetBook(ectx, in)
		return err
	})

	var r *entity.Review
	eg.Go(func() error {
		in := &book.GetReviewRequest{
			ReviewId: reviewID,
		}
		out, err := h.Book.GetReview(ectx, in)
		if err != nil {
			return err
		}
		r = entity.NewReview(out.Review)
		return nil
	})

	if err := eg.Wait(); err != nil {
		util.ErrorHandling(ctx, err)
		return
	}

	in := &user.GetUserRequest{
		UserId: r.UserId,
	}
	out, err := h.User.GetUser(c, in)
	if err != nil && !util.IsNotFound(err) {
		util.ErrorHandling(ctx, err)
		return
	}

	u := entity.NewUser(out.User)
	res := response.NewBookReviewResponse(r, u)
	ctx.JSON(http.StatusOK, res)
}

// getUserReview - ユーザーのレビュー情報取得
func (h *apiV1Handler) getUserReview(ctx *gin.Context) {
	c := util.SetMetadata(ctx)
	userID := ctx.Param("userID")
	reviewID, err := strconv.ParseInt(ctx.Param("reviewID"), 10, 64)
	if err != nil {
		util.ErrorHandling(ctx, entity.ErrBadRequest.New(err))
		return
	}

	reviewInput := &book.GetReviewRequest{
		ReviewId: reviewID,
	}
	reviewOutput, err := h.Book.GetReview(c, reviewInput)
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

	bookInput := &book.GetBookRequest{
		BookId: r.BookId,
	}
	bookOutput, err := h.Book.GetBook(c, bookInput)
	if err != nil {
		util.ErrorHandling(ctx, err)
		return
	}

	b := entity.NewBook(bookOutput.Book)
	res := response.NewUserReviewResponse(r, b)
	ctx.JSON(http.StatusOK, res)
}
