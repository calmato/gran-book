package v2

import (
	"net/http"
	"strconv"

	"github.com/calmato/gran-book/api/gateway/native/internal/entity"
	response "github.com/calmato/gran-book/api/gateway/native/internal/response/v2"
	"github.com/calmato/gran-book/api/gateway/native/internal/server/util"
	"github.com/calmato/gran-book/api/gateway/native/proto/service/book"
	"github.com/calmato/gran-book/api/gateway/native/proto/service/user"
	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
)

// listBookshelf - 本棚の書籍一覧取得
func (h *apiV2Handler) listBookshelf(ctx *gin.Context) {
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

	bookshelvesInput := &book.ListBookshelfRequest{
		UserId: userID,
		Limit:  limit,
		Offset: offset,
	}
	bookshelvesOutput, err := h.Book.ListBookshelf(c, bookshelvesInput)
	if err != nil {
		util.ErrorHandling(ctx, err)
		return
	}

	bss := entity.NewBookshelves(bookshelvesOutput.Bookshelves)

	booksInput := &book.MultiGetBooksRequest{
		BookIds: bss.BookIDs(),
	}
	booksOutput, err := h.Book.MultiGetBooks(c, booksInput)
	if err != nil {
		util.ErrorHandling(ctx, err)
		return
	}

	bs := entity.NewBooks(booksOutput.Books)
	res := response.NewBookshelfListResponse(
		bss, bs.Map(), bookshelvesOutput.Limit, bookshelvesOutput.Offset, bookshelvesOutput.Total,
	)
	ctx.JSON(http.StatusOK, res)
}

// getBookshelf - 本棚の書籍情報取得
func (h *apiV2Handler) getBookshelf(ctx *gin.Context) {
	c := util.SetMetadata(ctx)
	userID := ctx.Param("userID")
	bookID, err := strconv.ParseInt(ctx.Param("bookID"), 10, 64)
	if err != nil {
		util.ErrorHandling(ctx, entity.ErrBadRequest.New(err))
		return
	}

	eg, ectx := errgroup.WithContext(c)

	var b *entity.Book
	eg.Go(func() error {
		in := &book.GetBookRequest{
			BookId: bookID,
		}
		out, err := h.Book.GetBook(ectx, in)
		if err != nil {
			return err
		}
		b = entity.NewBook(out.Book)
		return nil
	})

	var bs *entity.Bookshelf
	eg.Go(func() error {
		in := &book.GetBookshelfRequest{
			UserId: userID,
			BookId: bookID,
		}
		out, err := h.Book.GetBookshelf(ectx, in)
		if err != nil {
			return err
		}
		bs = entity.NewBookshelf(out.Bookshelf)
		return nil
	})

	var rs entity.Reviews
	var limit, offset, total int64
	eg.Go(func() error {
		in := &book.ListBookReviewRequest{
			BookId: bookID,
			Limit:  20,
			Offset: 0,
		}
		out, err := h.Book.ListBookReview(ectx, in)
		if err != nil {
			return err
		}
		limit = out.Limit
		offset = out.Offset
		total = out.Total
		rs = entity.NewReviews(out.Reviews)
		return nil
	})

	if err := eg.Wait(); err != nil {
		util.ErrorHandling(ctx, err)
		return
	}

	usersInput := &user.MultiGetUserRequest{
		UserIds: rs.UserIDs(),
	}

	usersOutput, err := h.User.MultiGetUser(c, usersInput)
	if err != nil {
		util.ErrorHandling(ctx, err)
		return
	}

	us := entity.NewUsers(usersOutput.Users)
	res := response.NewBookshelfResponse(bs, b, rs, us.Map(), limit, offset, total)
	ctx.JSON(http.StatusOK, res)
}
