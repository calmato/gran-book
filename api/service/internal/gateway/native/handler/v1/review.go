package v1

import (
	"fmt"
	"net/http"
	"strconv"

	gentity "github.com/calmato/gran-book/api/service/internal/gateway/entity"
	"github.com/calmato/gran-book/api/service/internal/gateway/native/entity"
	response "github.com/calmato/gran-book/api/service/internal/gateway/native/response/v1"
	"github.com/calmato/gran-book/api/service/internal/gateway/util"
	"github.com/calmato/gran-book/api/service/pkg/exception"
	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
)

// listReviewByBook - 書籍のレビュー一覧取得
func (h *apiV1Handler) listReviewByBook(ctx *gin.Context) {
	c := util.SetMetadata(ctx)

	limit, err := strconv.ParseInt(ctx.DefaultQuery("limit", entity.ListLimitDefault), 10, 64)
	if err != nil {
		util.ErrorHandling(ctx, exception.ErrInvalidArgument.New(err))
		return
	}
	offset, err := strconv.ParseInt(ctx.DefaultQuery("offset", entity.ListOffsetDefault), 10, 64)
	if err != nil {
		util.ErrorHandling(ctx, exception.ErrInvalidArgument.New(err))
		return
	}
	bookID, err := strconv.ParseInt(ctx.Param("bookID"), 10, 64)
	if err != nil {
		util.ErrorHandling(ctx, exception.ErrInvalidArgument.New(err))
		return
	}

	rs, total, err := h.bookListBookReview(c, bookID, limit, offset)
	if err != nil {
		util.ErrorHandling(ctx, err)
		return
	}

	us, err := h.userMultiGetUser(c, rs.UserIDs())
	if err != nil {
		util.ErrorHandling(ctx, err)
		return
	}

	res := &response.BookReviewListResponse{
		Reviews: entity.NewBookReviews(rs, us.Map()),
		Limit:   limit,
		Offset:  offset,
		Total:   total,
	}
	ctx.JSON(http.StatusOK, res)
}

// listReviewByUser - ユーザーのレビュー一覧取得
func (h *apiV1Handler) listReviewByUser(ctx *gin.Context) {
	c := util.SetMetadata(ctx)

	userID := ctx.Param("userID")
	limit, err := strconv.ParseInt(ctx.DefaultQuery("limit", entity.ListLimitDefault), 10, 64)
	if err != nil {
		util.ErrorHandling(ctx, exception.ErrInvalidArgument.New(err))
		return
	}
	offset, err := strconv.ParseInt(ctx.DefaultQuery("offset", entity.ListOffsetDefault), 10, 64)
	if err != nil {
		util.ErrorHandling(ctx, exception.ErrInvalidArgument.New(err))
		return
	}

	rs, total, err := h.bookListUserReview(c, userID, limit, offset)
	if err != nil {
		util.ErrorHandling(ctx, err)
		return
	}

	bs, err := h.bookMultiGetBooks(c, rs.BookIDs())
	if err != nil {
		util.ErrorHandling(ctx, err)
		return
	}

	res := &response.UserReviewListResponse{
		Reviews: entity.NewUserReviews(rs, bs.Map()),
		Limit:   limit,
		Offset:  offset,
		Total:   total,
	}
	ctx.JSON(http.StatusOK, res)
}

// getBookReview - 書籍のレビュー情報取得
func (h *apiV1Handler) getBookReview(ctx *gin.Context) {
	c := util.SetMetadata(ctx)

	bookID, err := strconv.ParseInt(ctx.Param("bookID"), 10, 64)
	if err != nil {
		util.ErrorHandling(ctx, exception.ErrInvalidArgument.New(err))
		return
	}
	reviewID, err := strconv.ParseInt(ctx.Param("reviewID"), 10, 64)
	if err != nil {
		util.ErrorHandling(ctx, exception.ErrInvalidArgument.New(err))
		return
	}

	eg, ectx := errgroup.WithContext(c)
	eg.Go(func() error {
		_, err = h.bookGetBook(ectx, bookID)
		return err
	})
	var r *gentity.Review
	eg.Go(func() error {
		r, err = h.bookGetReview(ectx, reviewID)
		return err
	})
	if err := eg.Wait(); err != nil {
		util.ErrorHandling(ctx, err)
		return
	}

	u, err := h.userGetUser(c, r.UserId)
	if err != nil {
		util.ErrorHandling(ctx, err)
		return
	}

	res := &response.BookReviewResponse{
		BookReview: entity.NewBookReview(r, u),
	}
	ctx.JSON(http.StatusOK, res)
}

// getUserReview - ユーザーのレビュー情報取得
func (h *apiV1Handler) getUserReview(ctx *gin.Context) {
	c := util.SetMetadata(ctx)

	userID := ctx.Param("userID")
	reviewID, err := strconv.ParseInt(ctx.Param("reviewID"), 10, 64)
	if err != nil {
		util.ErrorHandling(ctx, exception.ErrInvalidArgument.New(err))
		return
	}

	r, err := h.bookGetReview(c, reviewID)
	if err != nil {
		util.ErrorHandling(ctx, err)
		return
	}
	if r.UserId != userID {
		err := fmt.Errorf("user id is invalid: %s", userID)
		util.ErrorHandling(ctx, exception.ErrInvalidArgument.New(err))
		return
	}

	b, err := h.bookGetBook(c, r.BookId)
	if err != nil {
		util.ErrorHandling(ctx, err)
		return
	}

	res := &response.UserReviewResponse{
		UserReview: entity.NewUserReview(r, b),
	}
	ctx.JSON(http.StatusOK, res)
}
