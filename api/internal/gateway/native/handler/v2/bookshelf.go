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
	"golang.org/x/sync/errgroup"
)

// listBookshelf - 本棚の書籍一覧取得
func (h *apiV2Handler) listBookshelf(ctx *gin.Context) {
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

	bss, total, err := h.bookListBookshelf(c, userID, limit, offset)
	if err != nil {
		util.ErrorHandling(ctx, err)
		return
	}

	bs, err := h.bookMultiGetBooks(c, bss.BookIDs())
	if err != nil {
		util.ErrorHandling(ctx, err)
		return
	}

	res := &response.BookshelfListResponse{
		Books:  entity.NewBooksOnBookshelf(bs.Map(), bss),
		Limit:  limit,
		Offset: offset,
		Total:  total,
	}
	ctx.JSON(http.StatusOK, res)
}

// getBookshelf - 本棚の書籍情報取得
func (h *apiV2Handler) getBookshelf(ctx *gin.Context) {
	c := util.SetMetadata(ctx)

	limit, _ := strconv.ParseInt(entity.ListLimitDefault, 10, 64)
	offset, _ := strconv.ParseInt(entity.ListOffsetDefault, 10, 64)

	userID := ctx.Param("userID")
	bookID, err := strconv.ParseInt(ctx.Param("bookID"), 10, 64)
	if err != nil {
		util.ErrorHandling(ctx, exception.ErrInvalidArgument.New(err))
		return
	}

	eg, ectx := errgroup.WithContext(c)
	var b *gentity.Book
	eg.Go(func() error {
		b, err = h.bookGetBook(ectx, bookID)
		return err
	})
	var bs *gentity.Bookshelf
	eg.Go(func() error {
		bs, err = h.bookGetBookshelf(ectx, userID, bookID)
		return err
	})
	var rs gentity.Reviews
	var total int64
	eg.Go(func() error {
		rs, total, err = h.bookListBookReview(ectx, bookID, limit, offset)
		return err
	})
	if err := eg.Wait(); err != nil {
		util.ErrorHandling(ctx, err)
		return
	}

	us, err := h.userMultiGetUser(c, rs.UserIDs())
	if err != nil {
		util.ErrorHandling(ctx, err)
		return
	}

	reviews := entity.NewBookReviews(rs, us.Map())
	bookshelf := entity.NewBookOnBookshelf(b, bs)
	bookshelf.Fill(reviews, limit, offset, total)
	res := &response.BookshelfResponse{
		BookOnBookshelf: bookshelf,
	}
	ctx.JSON(http.StatusOK, res)
}

func (h *apiV2Handler) bookListBookshelf(
	ctx context.Context, userID string, limit, offset int64,
) (gentity.Bookshelves, int64, error) {
	in := &book.ListBookshelfRequest{
		UserId: userID,
		Limit:  limit,
		Offset: offset,
	}
	out, err := h.Book.ListBookshelf(ctx, in)
	if err != nil {
		return nil, 0, err
	}

	return gentity.NewBookshelves(out.Bookshelves), out.Total, nil
}

func (h *apiV2Handler) bookGetBookshelf(ctx context.Context, userID string, bookID int64) (*gentity.Bookshelf, error) {
	in := &book.GetBookshelfRequest{
		UserId: userID,
		BookId: bookID,
	}
	out, err := h.Book.GetBookshelf(ctx, in)
	if err != nil {
		return nil, err
	}
	if out.Bookshelf == nil {
		err := fmt.Errorf("bookshelf is not found: %s, %d", userID, bookID)
		return nil, exception.ErrNotFound.New(err)
	}

	return gentity.NewBookshelf(out.Bookshelf), nil
}
