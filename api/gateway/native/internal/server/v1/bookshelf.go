package v1

import (
	"net/http"
	"strconv"

	"github.com/calmato/gran-book/api/gateway/native/internal/entity"
	request "github.com/calmato/gran-book/api/gateway/native/internal/request/v1"
	response "github.com/calmato/gran-book/api/gateway/native/internal/response/v1"
	"github.com/calmato/gran-book/api/gateway/native/internal/server/util"
	"github.com/calmato/gran-book/api/gateway/native/proto/service/book"
	"github.com/calmato/gran-book/api/gateway/native/proto/service/user"
	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
)

type BookshelfHandler interface {
	List(ctx *gin.Context)
	Get(ctx *gin.Context)
	Read(ctx *gin.Context)
	Reading(ctx *gin.Context)
	Stacked(ctx *gin.Context)
	Want(ctx *gin.Context)
	Release(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type bookshelfHandler struct {
	authClient user.AuthServiceClient
	bookClient book.BookServiceClient
}

func NewBookshelfHandler(ac user.AuthServiceClient, bc book.BookServiceClient) BookshelfHandler {
	return &bookshelfHandler{
		authClient: ac,
		bookClient: bc,
	}
}

// List - 本棚の書籍一覧取得 ※廃止予定
func (h *bookshelfHandler) List(ctx *gin.Context) {
	c := util.SetMetadata(ctx)
	limit := ctx.GetInt64(ctx.DefaultQuery("limit", entity.ListLimitDefault))
	offset := ctx.GetInt64(ctx.DefaultQuery("offset", entity.ListOffsetDefault))
	userID := ctx.Param("userID")

	bookshelfInput := &book.ListBookshelfRequest{
		UserId: userID,
		Limit:  limit,
		Offset: offset,
	}
	bookshelfOutput, err := h.bookClient.ListBookshelf(c, bookshelfInput)
	if err != nil {
		util.ErrorHandling(ctx, err)
		return
	}

	bss := entity.NewBookshelves(bookshelfOutput.Bookshelves)

	booksInput := &book.MultiGetBooksRequest{
		BookIds: bss.BookIDs(),
	}
	booksOutput, err := h.bookClient.MultiGetBooks(c, booksInput)
	if err != nil {
		util.ErrorHandling(ctx, err)
		return
	}

	bs := entity.NewBooks(booksOutput.Books)
	res := response.NewBookshelfListResponse(
		bss, bs.Map(), bookshelfOutput.Limit, bookshelfOutput.Offset, bookshelfOutput.Total,
	)
	ctx.JSON(http.StatusOK, res)
}

// Get - 本棚の書籍情報取得 ※廃止予定
func (h *bookshelfHandler) Get(ctx *gin.Context) {
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
		out, err := h.bookClient.GetBook(ectx, in)
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
		out, err := h.bookClient.GetBookshelf(ectx, in)
		if err != nil {
			return err
		}
		bs = entity.NewBookshelf(out.Bookshelf)
		return nil
	})

	if err := eg.Wait(); err != nil {
		util.ErrorHandling(ctx, err)
		return
	}

	res := response.NewBookshelfResponse(bs, b)
	ctx.JSON(http.StatusOK, res)
}

// Read - 読んだ本の登録
func (h *bookshelfHandler) Read(ctx *gin.Context) {
	c := util.SetMetadata(ctx)
	userID := ctx.Param("userID")
	bookID, err := strconv.ParseInt(ctx.Param("bookID"), 10, 64)
	if err != nil {
		util.ErrorHandling(ctx, entity.ErrBadRequest.New(err))
		return
	}

	req := &request.ReadBookshelfRequest{}
	if err = ctx.BindJSON(req); err != nil {
		util.ErrorHandling(ctx, entity.ErrBadRequest.New(err))
		return
	}

	eg, ectx := errgroup.WithContext(c)

	var b *entity.Book
	eg.Go(func() error {
		in := &book.GetBookRequest{
			BookId: bookID,
		}
		out, err := h.bookClient.GetBook(ectx, in)
		if err != nil {
			return err
		}
		b = entity.NewBook(out.Book)
		return nil
	})

	var bs *entity.Bookshelf
	eg.Go(func() error {
		in := &book.ReadBookshelfRequest{
			UserId:     userID,
			BookId:     bookID,
			Impression: req.Impression,
			ReadOn:     req.ReadOn,
		}
		out, err := h.bookClient.ReadBookshelf(ectx, in)
		if err != nil {
			return err
		}
		bs = entity.NewBookshelf(out.Bookshelf)
		return nil
	})

	if err := eg.Wait(); err != nil {
		util.ErrorHandling(ctx, err)
		return
	}

	res := response.NewBookshelfResponse(bs, b)
	ctx.JSON(http.StatusOK, res)
}

// Reading - 読んでいる本の登録
func (h *bookshelfHandler) Reading(ctx *gin.Context) {
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
		out, err := h.bookClient.GetBook(ectx, in)
		if err != nil {
			return err
		}
		b = entity.NewBook(out.Book)
		return nil
	})

	var bs *entity.Bookshelf
	eg.Go(func() error {
		in := &book.ReadingBookshelfRequest{
			UserId: userID,
			BookId: bookID,
		}
		out, err := h.bookClient.ReadingBookshelf(ectx, in)
		if err != nil {
			return err
		}
		bs = entity.NewBookshelf(out.Bookshelf)
		return nil
	})

	if err := eg.Wait(); err != nil {
		util.ErrorHandling(ctx, err)
		return
	}

	res := response.NewBookshelfResponse(bs, b)
	ctx.JSON(http.StatusOK, res)
}

// Stacked - 積読本の登録
func (h *bookshelfHandler) Stacked(ctx *gin.Context) {
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
		out, err := h.bookClient.GetBook(ectx, in)
		if err != nil {
			return err
		}
		b = entity.NewBook(out.Book)
		return nil
	})

	var bs *entity.Bookshelf
	eg.Go(func() error {
		in := &book.StackedBookshelfRequest{
			UserId: userID,
			BookId: bookID,
		}
		out, err := h.bookClient.StackedBookshelf(ectx, in)
		if err != nil {
			return err
		}
		bs = entity.NewBookshelf(out.Bookshelf)
		return nil
	})

	if err := eg.Wait(); err != nil {
		util.ErrorHandling(ctx, err)
		return
	}

	res := response.NewBookshelfResponse(bs, b)
	ctx.JSON(http.StatusOK, res)
}

// Want - 欲しい本の登録
func (h *bookshelfHandler) Want(ctx *gin.Context) {
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
		out, err := h.bookClient.GetBook(ectx, in)
		if err != nil {
			return err
		}
		b = entity.NewBook(out.Book)
		return nil
	})

	var bs *entity.Bookshelf
	eg.Go(func() error {
		in := &book.WantBookshelfRequest{
			UserId: userID,
			BookId: bookID,
		}
		out, err := h.bookClient.WantBookshelf(ectx, in)
		if err != nil {
			return err
		}
		bs = entity.NewBookshelf(out.Bookshelf)
		return nil
	})

	if err := eg.Wait(); err != nil {
		util.ErrorHandling(ctx, err)
		return
	}

	res := response.NewBookshelfResponse(bs, b)
	ctx.JSON(http.StatusOK, res)
}

// Release - 手放したい本の登録
func (h *bookshelfHandler) Release(ctx *gin.Context) {
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
		out, err := h.bookClient.GetBook(ectx, in)
		if err != nil {
			return err
		}
		b = entity.NewBook(out.Book)
		return nil
	})

	var bs *entity.Bookshelf
	eg.Go(func() error {
		in := &book.ReleaseBookshelfRequest{
			UserId: userID,
			BookId: bookID,
		}
		out, err := h.bookClient.ReleaseBookshelf(ectx, in)
		if err != nil {
			return err
		}
		bs = entity.NewBookshelf(out.Bookshelf)
		return nil
	})

	if err := eg.Wait(); err != nil {
		util.ErrorHandling(ctx, err)
		return
	}

	res := response.NewBookshelfResponse(bs, b)
	ctx.JSON(http.StatusOK, res)
}

// Delete - 本棚から書籍の削除
func (h *bookshelfHandler) Delete(ctx *gin.Context) {
	c := util.SetMetadata(ctx)
	userID := ctx.Param("userID")
	bookID, err := strconv.ParseInt(ctx.Param("bookID"), 10, 64)
	if err != nil {
		util.ErrorHandling(ctx, entity.ErrBadRequest.New(err))
		return
	}

	in := &book.DeleteBookshelfRequest{
		UserId: userID,
		BookId: bookID,
	}
	_, err = h.bookClient.DeleteBookshelf(c, in)
	if err != nil {
		util.ErrorHandling(ctx, err)
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}
