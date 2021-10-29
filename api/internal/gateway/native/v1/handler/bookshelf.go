package handler

import (
	"fmt"
	"net/http"
	"strconv"

	gentity "github.com/calmato/gran-book/api/internal/gateway/entity"
	"github.com/calmato/gran-book/api/internal/gateway/native/v1/entity"
	request "github.com/calmato/gran-book/api/internal/gateway/native/v1/request"
	response "github.com/calmato/gran-book/api/internal/gateway/native/v1/response"
	"github.com/calmato/gran-book/api/internal/gateway/util"
	"github.com/calmato/gran-book/api/pkg/exception"
	"github.com/calmato/gran-book/api/proto/book"
	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
)

// listBookshelf - 本棚の書籍一覧取得 ※廃止予定
// Deprecated: use v2.listBookshelf
func (h *apiV1Handler) listBookshelf(ctx *gin.Context) {
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

	bookshelfInput := &book.ListBookshelfRequest{
		UserId: userID,
		Limit:  limit,
		Offset: offset,
	}
	bookshelfOutput, err := h.Book.ListBookshelf(c, bookshelfInput)
	if err != nil {
		util.ErrorHandling(ctx, err)
		return
	}

	bss := gentity.NewBookshelves(bookshelfOutput.Bookshelves)

	booksInput := &book.MultiGetBooksRequest{
		BookIds: bss.BookIDs(),
	}
	booksOutput, err := h.Book.MultiGetBooks(c, booksInput)
	if err != nil {
		util.ErrorHandling(ctx, err)
		return
	}

	bs := gentity.NewBooks(booksOutput.Books)
	res := response.NewBookshelfListResponse(
		bss, bs.Map(), bookshelfOutput.Limit, bookshelfOutput.Offset, bookshelfOutput.Total,
	)
	ctx.JSON(http.StatusOK, res)
}

// getBookshelf - 本棚の書籍情報取得 ※廃止予定
// Deprecated: use v2.getBookshelf
func (h *apiV1Handler) getBookshelf(ctx *gin.Context) {
	c := util.SetMetadata(ctx)
	userID := ctx.Param("userID")
	bookID, err := strconv.ParseInt(ctx.Param("bookID"), 10, 64)
	if err != nil {
		util.ErrorHandling(ctx, exception.ErrInvalidArgument.New(err))
		return
	}

	eg, ectx := errgroup.WithContext(c)

	var b *gentity.Book
	eg.Go(func() error {
		in := &book.GetBookRequest{
			BookId: bookID,
		}
		out, err := h.Book.GetBook(ectx, in)
		if err != nil {
			return err
		}
		b = gentity.NewBook(out.Book)
		return nil
	})

	var bs *gentity.Bookshelf
	eg.Go(func() error {
		in := &book.GetBookshelfRequest{
			UserId: userID,
			BookId: bookID,
		}
		out, err := h.Book.GetBookshelf(ectx, in)
		if err != nil {
			return err
		}
		bs = gentity.NewBookshelf(out.Bookshelf)
		return nil
	})

	if err := eg.Wait(); err != nil {
		util.ErrorHandling(ctx, err)
		return
	}

	res := response.NewBookshelfResponse(bs, b)
	ctx.JSON(http.StatusOK, res)
}

// readBookshelf - 読んだ本の登録
func (h *apiV1Handler) readBookshelf(ctx *gin.Context) {
	c := util.SetMetadata(ctx)

	userID := ctx.Param("userID")
	bookID, err := strconv.ParseInt(ctx.Param("bookID"), 10, 64)
	if err != nil {
		util.ErrorHandling(ctx, exception.ErrInvalidArgument.New(err))
		return
	}

	req := &request.ReadBookshelfRequest{}
	if err = ctx.BindJSON(req); err != nil {
		util.ErrorHandling(ctx, exception.ErrInvalidArgument.New(err))
		return
	}

	eg, ectx := errgroup.WithContext(c)
	eg.Go(func() error {
		ok, err := h.correctUser(ectx, userID)
		if err != nil || !ok {
			return fmt.Errorf("v1: user id is not correct: %w", err)
		}
		return nil
	})
	var b *gentity.Book
	eg.Go(func() error {
		b, err = h.bookGetBook(ectx, bookID)
		return err
	})
	if err := eg.Wait(); err != nil {
		util.ErrorHandling(ctx, err)
		return
	}

	in := &book.ReadBookshelfRequest{
		UserId:     userID,
		BookId:     bookID,
		Impression: req.Impression,
		ReadOn:     req.ReadOn,
	}
	out, err := h.Book.ReadBookshelf(ectx, in)
	if err != nil {
		util.ErrorHandling(ctx, err)
		return
	}
	bs := gentity.NewBookshelf(out.Bookshelf)

	res := response.NewBookshelfResponse(bs, b)
	ctx.JSON(http.StatusOK, res)
}

// readingBookshelf - 読んでいる本の登録
func (h *apiV1Handler) readingBookshelf(ctx *gin.Context) {
	c := util.SetMetadata(ctx)

	userID := ctx.Param("userID")
	bookID, err := strconv.ParseInt(ctx.Param("bookID"), 10, 64)
	if err != nil {
		util.ErrorHandling(ctx, exception.ErrInvalidArgument.New(err))
		return
	}

	eg, ectx := errgroup.WithContext(c)
	eg.Go(func() error {
		ok, err := h.correctUser(ectx, userID)
		if err != nil || !ok {
			return fmt.Errorf("v1: user id is not correct: %w", err)
		}
		return nil
	})
	var b *gentity.Book
	eg.Go(func() error {
		b, err = h.bookGetBook(ectx, bookID)
		return err
	})
	if err := eg.Wait(); err != nil {
		util.ErrorHandling(ctx, err)
		return
	}

	in := &book.ReadingBookshelfRequest{
		UserId: userID,
		BookId: bookID,
	}
	out, err := h.Book.ReadingBookshelf(ectx, in)
	if err != nil {
		util.ErrorHandling(ctx, err)
		return
	}
	bs := gentity.NewBookshelf(out.Bookshelf)

	res := response.NewBookshelfResponse(bs, b)
	ctx.JSON(http.StatusOK, res)
}

// stackedBookshelf - 積読本の登録
func (h *apiV1Handler) stackedBookshelf(ctx *gin.Context) {
	c := util.SetMetadata(ctx)

	userID := ctx.Param("userID")
	bookID, err := strconv.ParseInt(ctx.Param("bookID"), 10, 64)
	if err != nil {
		util.ErrorHandling(ctx, exception.ErrInvalidArgument.New(err))
		return
	}

	eg, ectx := errgroup.WithContext(c)
	eg.Go(func() error {
		ok, err := h.correctUser(ectx, userID)
		if err != nil || !ok {
			return fmt.Errorf("v1: user id is not correct: %w", err)
		}
		return nil
	})
	var b *gentity.Book
	eg.Go(func() error {
		b, err = h.bookGetBook(ectx, bookID)
		return err
	})
	if err := eg.Wait(); err != nil {
		util.ErrorHandling(ctx, err)
		return
	}

	in := &book.StackedBookshelfRequest{
		UserId: userID,
		BookId: bookID,
	}
	out, err := h.Book.StackedBookshelf(ectx, in)
	if err != nil {
		util.ErrorHandling(ctx, err)
		return
	}
	bs := gentity.NewBookshelf(out.Bookshelf)

	res := response.NewBookshelfResponse(bs, b)
	ctx.JSON(http.StatusOK, res)
}

// wantBookshelf - 欲しい本の登録
func (h *apiV1Handler) wantBookshelf(ctx *gin.Context) {
	c := util.SetMetadata(ctx)

	userID := ctx.Param("userID")
	bookID, err := strconv.ParseInt(ctx.Param("bookID"), 10, 64)
	if err != nil {
		util.ErrorHandling(ctx, exception.ErrInvalidArgument.New(err))
		return
	}

	eg, ectx := errgroup.WithContext(c)
	eg.Go(func() error {
		ok, err := h.correctUser(ectx, userID)
		if err != nil || !ok {
			return fmt.Errorf("v1: user id is not correct: %w", err)
		}
		return nil
	})
	var b *gentity.Book
	eg.Go(func() error {
		b, err = h.bookGetBook(ectx, bookID)
		return err
	})
	if err := eg.Wait(); err != nil {
		util.ErrorHandling(ctx, err)
		return
	}

	in := &book.WantBookshelfRequest{
		UserId: userID,
		BookId: bookID,
	}
	out, err := h.Book.WantBookshelf(ectx, in)
	if err != nil {
		util.ErrorHandling(ctx, err)
		return
	}
	bs := gentity.NewBookshelf(out.Bookshelf)

	res := response.NewBookshelfResponse(bs, b)
	ctx.JSON(http.StatusOK, res)
}

// releaseBookshelf - 手放したい本の登録
func (h *apiV1Handler) releaseBookshelf(ctx *gin.Context) {
	c := util.SetMetadata(ctx)

	userID := ctx.Param("userID")
	bookID, err := strconv.ParseInt(ctx.Param("bookID"), 10, 64)
	if err != nil {
		util.ErrorHandling(ctx, exception.ErrInvalidArgument.New(err))
		return
	}

	eg, ectx := errgroup.WithContext(c)
	eg.Go(func() error {
		ok, err := h.correctUser(ectx, userID)
		if err != nil || !ok {
			return fmt.Errorf("v1: user id is not correct: %w", err)
		}
		return nil
	})
	var b *gentity.Book
	eg.Go(func() error {
		b, err = h.bookGetBook(ectx, bookID)
		return err
	})
	if err := eg.Wait(); err != nil {
		util.ErrorHandling(ctx, err)
		return
	}

	in := &book.ReleaseBookshelfRequest{
		UserId: userID,
		BookId: bookID,
	}
	out, err := h.Book.ReleaseBookshelf(ectx, in)
	if err != nil {
		util.ErrorHandling(ctx, err)
		return
	}
	bs := gentity.NewBookshelf(out.Bookshelf)

	res := response.NewBookshelfResponse(bs, b)
	ctx.JSON(http.StatusOK, res)
}

// deleteBookshelf - 本棚から書籍の削除
func (h *apiV1Handler) deleteBookshelf(ctx *gin.Context) {
	c := util.SetMetadata(ctx)

	userID := ctx.Param("userID")
	bookID, err := strconv.ParseInt(ctx.Param("bookID"), 10, 64)
	if err != nil {
		util.ErrorHandling(ctx, exception.ErrInvalidArgument.New(err))
		return
	}

	eg, ectx := errgroup.WithContext(c)
	eg.Go(func() error {
		ok, err := h.correctUser(ectx, userID)
		if err != nil || !ok {
			return fmt.Errorf("v1: user id is not correct: %w", err)
		}
		return nil
	})
	eg.Go(func() error {
		_, err = h.bookGetBook(ectx, bookID)
		return err
	})
	if err := eg.Wait(); err != nil {
		util.ErrorHandling(ctx, err)
		return
	}

	in := &book.DeleteBookshelfRequest{
		UserId: userID,
		BookId: bookID,
	}
	_, err = h.Book.DeleteBookshelf(c, in)
	if err != nil {
		util.ErrorHandling(ctx, err)
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}
