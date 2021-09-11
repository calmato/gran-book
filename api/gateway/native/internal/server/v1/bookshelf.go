package v1

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/calmato/gran-book/api/gateway/native/internal/entity"
	request "github.com/calmato/gran-book/api/gateway/native/internal/request/v1"
	response "github.com/calmato/gran-book/api/gateway/native/internal/response/v1"
	"github.com/calmato/gran-book/api/gateway/native/internal/server/util"
	"github.com/calmato/gran-book/api/gateway/native/proto/book"
	"github.com/calmato/gran-book/api/gateway/native/proto/user"
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
	limit := ctx.GetInt64(ctx.DefaultQuery("limit", entity.ListLimitDefault))
	offset := ctx.GetInt64(ctx.DefaultQuery("offset", entity.ListOffsetDefault))
	userID := ctx.Param("userID")

	bookshelfInput := &book.ListBookshelfRequest{
		UserId: userID,
		Limit:  limit,
		Offset: offset,
	}

	c := util.SetMetadata(ctx)
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
	res := h.getBookshelfListResponse(
		bss, bs.Map(), bookshelfOutput.Limit, bookshelfOutput.Offset, bookshelfOutput.Total,
	)
	ctx.JSON(http.StatusOK, res)
}

// Get - 本棚の書籍情報取得 ※廃止予定
func (h *bookshelfHandler) Get(ctx *gin.Context) {
	userID := ctx.Param("userID")
	bookID, err := strconv.ParseInt(ctx.Param("bookID"), 10, 64)
	if err != nil {
		util.ErrorHandling(ctx, entity.ErrBadRequest.New(err))
		return
	}

	c := util.SetMetadata(ctx)
	eg, ectx := errgroup.WithContext(c)

	var b *entity.Book
	eg.Go(func() error {
		bookInput := &book.GetBookRequest{
			BookId: bookID,
		}
		bookOutput, err := h.bookClient.GetBook(ectx, bookInput)
		if err != nil {
			return err
		}
		b = entity.NewBook(bookOutput.Book)
		return nil
	})

	var bs *entity.Bookshelf
	eg.Go(func() error {
		bookshelfInput := &book.GetBookshelfRequest{
			UserId: userID,
			BookId: bookID,
		}
		bookshelfOutput, err := h.bookClient.GetBookshelf(ectx, bookshelfInput)
		if err != nil {
			return err
		}
		bs = entity.NewBookshelf(bookshelfOutput.Bookshelf)
		return nil
	})

	if err := eg.Wait(); err != nil {
		util.ErrorHandling(ctx, err)
		return
	}

	res := h.getBookshelfResponse(bs, b)
	ctx.JSON(http.StatusOK, res)
}

// Read - 読んだ本の登録
func (h *bookshelfHandler) Read(ctx *gin.Context) {
	userID := ctx.Param("userID")
	bookID, err := strconv.ParseInt(ctx.Param("bookID"), 10, 64)
	if err != nil {
		util.ErrorHandling(ctx, entity.ErrBadRequest.New(err))
		return
	}

	req := &request.ReadBookshelfRequest{}
	err = ctx.BindJSON(req)
	if err != nil {
		util.ErrorHandling(ctx, entity.ErrBadRequest.New(err))
		return
	}

	c := util.SetMetadata(ctx)
	eg, ectx := errgroup.WithContext(c)

	var b *entity.Book
	eg.Go(func() error {
		bookInput := &book.GetBookRequest{
			BookId: bookID,
		}
		bookOutput, err := h.bookClient.GetBook(ectx, bookInput)
		if err != nil {
			return err
		}
		b = entity.NewBook(bookOutput.Book)
		return nil
	})

	var bs *entity.Bookshelf
	eg.Go(func() error {
		bookshelfInput := &book.ReadBookshelfRequest{
			UserId:     userID,
			BookId:     bookID,
			Impression: req.Impression,
			ReadOn:     req.ReadOn,
		}
		bookshelfOutput, err := h.bookClient.ReadBookshelf(ectx, bookshelfInput)
		if err != nil {
			return err
		}
		bs = entity.NewBookshelf(bookshelfOutput.Bookshelf)
		return nil
	})

	if err := eg.Wait(); err != nil {
		util.ErrorHandling(ctx, err)
		return
	}

	res := h.getBookshelfResponse(bs, b)
	ctx.JSON(http.StatusOK, res)
}

// Reading - 読んでいる本の登録
func (h *bookshelfHandler) Reading(ctx *gin.Context) {
	userID := ctx.Param("userID")
	bookID, err := strconv.ParseInt(ctx.Param("bookID"), 10, 64)
	if err != nil {
		util.ErrorHandling(ctx, entity.ErrBadRequest.New(err))
		return
	}

	c := util.SetMetadata(ctx)
	eg, ectx := errgroup.WithContext(c)

	var b *entity.Book
	eg.Go(func() error {
		bookInput := &book.GetBookRequest{
			BookId: bookID,
		}
		bookOutput, err := h.bookClient.GetBook(ectx, bookInput)
		if err != nil {
			return err
		}
		b = entity.NewBook(bookOutput.Book)
		return nil
	})

	var bs *entity.Bookshelf
	eg.Go(func() error {
		bookshelfInput := &book.ReadingBookshelfRequest{
			UserId: userID,
			BookId: bookID,
		}
		bookshelfOutput, err := h.bookClient.ReadingBookshelf(ectx, bookshelfInput)
		if err != nil {
			return err
		}
		bs = entity.NewBookshelf(bookshelfOutput.Bookshelf)
		return nil
	})

	if err := eg.Wait(); err != nil {
		util.ErrorHandling(ctx, err)
		return
	}

	res := h.getBookshelfResponse(bs, b)
	ctx.JSON(http.StatusOK, res)
}

// Stacked - 積読本の登録
func (h *bookshelfHandler) Stacked(ctx *gin.Context) {
	userID := ctx.Param("userID")
	bookID, err := strconv.ParseInt(ctx.Param("bookID"), 10, 64)
	if err != nil {
		util.ErrorHandling(ctx, entity.ErrBadRequest.New(err))
		return
	}

	c := util.SetMetadata(ctx)
	eg, ectx := errgroup.WithContext(c)

	var b *entity.Book
	eg.Go(func() error {
		bookInput := &book.GetBookRequest{
			BookId: bookID,
		}
		bookOutput, err := h.bookClient.GetBook(ectx, bookInput)
		if err != nil {
			return err
		}
		b = entity.NewBook(bookOutput.Book)
		return nil
	})

	var bs *entity.Bookshelf
	eg.Go(func() error {
		bookshelfInput := &book.StackedBookshelfRequest{
			UserId: userID,
			BookId: bookID,
		}
		bookshelfOutput, err := h.bookClient.StackedBookshelf(ectx, bookshelfInput)
		if err != nil {
			return err
		}
		bs = entity.NewBookshelf(bookshelfOutput.Bookshelf)
		return nil
	})

	if err := eg.Wait(); err != nil {
		util.ErrorHandling(ctx, err)
		return
	}

	res := h.getBookshelfResponse(bs, b)
	ctx.JSON(http.StatusOK, res)
}

// Want - 欲しい本の登録
func (h *bookshelfHandler) Want(ctx *gin.Context) {
	userID := ctx.Param("userID")
	bookID, err := strconv.ParseInt(ctx.Param("bookID"), 10, 64)
	if err != nil {
		util.ErrorHandling(ctx, entity.ErrBadRequest.New(err))
		return
	}

	c := util.SetMetadata(ctx)
	eg, ectx := errgroup.WithContext(c)

	var b *entity.Book
	eg.Go(func() error {
		bookInput := &book.GetBookRequest{
			BookId: bookID,
		}
		bookOutput, err := h.bookClient.GetBook(ectx, bookInput)
		if err != nil {
			return err
		}
		b = entity.NewBook(bookOutput.Book)
		return nil
	})

	var bs *entity.Bookshelf
	eg.Go(func() error {
		bookshelfInput := &book.WantBookshelfRequest{
			UserId: userID,
			BookId: bookID,
		}
		bookshelfOutput, err := h.bookClient.WantBookshelf(ectx, bookshelfInput)
		if err != nil {
			return err
		}
		bs = entity.NewBookshelf(bookshelfOutput.Bookshelf)
		return nil
	})

	if err := eg.Wait(); err != nil {
		util.ErrorHandling(ctx, err)
		return
	}

	res := h.getBookshelfResponse(bs, b)
	ctx.JSON(http.StatusOK, res)
}

// Release - 手放したい本の登録
func (h *bookshelfHandler) Release(ctx *gin.Context) {
	userID := ctx.Param("userID")
	bookID, err := strconv.ParseInt(ctx.Param("bookID"), 10, 64)
	if err != nil {
		util.ErrorHandling(ctx, entity.ErrBadRequest.New(err))
		return
	}

	c := util.SetMetadata(ctx)
	eg, ectx := errgroup.WithContext(c)

	var b *entity.Book
	eg.Go(func() error {
		bookInput := &book.GetBookRequest{
			BookId: bookID,
		}
		bookOutput, err := h.bookClient.GetBook(ectx, bookInput)
		if err != nil {
			return err
		}
		b = entity.NewBook(bookOutput.Book)
		return nil
	})

	var bs *entity.Bookshelf
	eg.Go(func() error {
		bookshelfInput := &book.ReleaseBookshelfRequest{
			UserId: userID,
			BookId: bookID,
		}
		bookshelfOutput, err := h.bookClient.ReleaseBookshelf(ectx, bookshelfInput)
		if err != nil {
			return err
		}
		bs = entity.NewBookshelf(bookshelfOutput.Bookshelf)
		return nil
	})

	if err := eg.Wait(); err != nil {
		util.ErrorHandling(ctx, err)
		return
	}

	res := h.getBookshelfResponse(bs, b)
	ctx.JSON(http.StatusOK, res)
}

// Delete - 本棚から書籍の削除
func (h *bookshelfHandler) Delete(ctx *gin.Context) {
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

	c := util.SetMetadata(ctx)
	_, err = h.bookClient.DeleteBookshelf(c, in)
	if err != nil {
		util.ErrorHandling(ctx, err)
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}

func (h *bookshelfHandler) getBookshelfResponse(bs *entity.Bookshelf, b *entity.Book) *response.BookshelfResponse {
	bookshelf := &response.BookshelfBookshelf{
		ID:         bs.Id,
		Status:     bs.Status().Name(),
		ReadOn:     bs.ReadOn,
		Impression: "",
		CreatedAt:  bs.CreatedAt,
		UpdatedAt:  bs.UpdatedAt,
	}

	return &response.BookshelfResponse{
		ID:           b.Id,
		Title:        b.Title,
		TitleKana:    b.TitleKana,
		Description:  b.Description,
		Isbn:         b.Isbn,
		Publisher:    b.Publisher,
		PublishedOn:  b.PublishedOn,
		ThumbnailURL: b.ThumbnailUrl,
		RakutenURL:   b.RakutenUrl,
		Size:         b.RakutenSize,
		Author:       strings.Join(b.AuthorNames(), "/"),
		AuthorKana:   strings.Join(b.AuthorNameKanas(), "/"),
		CreatedAt:    b.CreatedAt,
		UpdatedAt:    b.UpdatedAt,
		Bookshelf:    bookshelf,
	}
}

func (h *bookshelfHandler) getBookshelfListResponse(
	bss entity.Bookshelves, bm map[int64]*entity.Book, limit, offset, total int64,
) *response.BookshelfListResponse {
	books := make([]*response.BookshelfListBook, 0, len(bss))
	for _, bs := range bss {
		b, ok := bm[bs.BookId]
		if !ok {
			continue
		}

		bookshelf := &response.BookshelfListBookshelf{
			ID:        bs.Id,
			Status:    bs.Status().Name(),
			ReadOn:    bs.ReadOn,
			CreatedAt: bs.CreatedAt,
			UpdatedAt: bs.UpdatedAt,
		}

		book := &response.BookshelfListBook{
			ID:           b.Id,
			Title:        b.Title,
			TitleKana:    b.TitleKana,
			Description:  b.Description,
			Isbn:         b.Isbn,
			Publisher:    b.Publisher,
			PublishedOn:  b.PublishedOn,
			ThumbnailURL: b.ThumbnailUrl,
			RakutenURL:   b.RakutenUrl,
			Size:         b.RakutenSize,
			Author:       strings.Join(b.AuthorNames(), "/"),
			AuthorKana:   strings.Join(b.AuthorNameKanas(), "/"),
			CreatedAt:    b.CreatedAt,
			UpdatedAt:    b.UpdatedAt,
			Bookshelf:    bookshelf,
		}

		books = append(books, book)
	}

	return &response.BookshelfListResponse{
		Books:  books,
		Limit:  limit,
		Offset: offset,
		Total:  total,
	}
}
