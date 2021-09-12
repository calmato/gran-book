package v1

import (
	"net/http"
	"strings"

	"github.com/calmato/gran-book/api/gateway/native/internal/entity"
	request "github.com/calmato/gran-book/api/gateway/native/internal/request/v1"
	response "github.com/calmato/gran-book/api/gateway/native/internal/response/v1"
	"github.com/calmato/gran-book/api/gateway/native/internal/server/util"
	"github.com/calmato/gran-book/api/gateway/native/proto/service/book"
	"github.com/calmato/gran-book/api/gateway/native/proto/service/user"
	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
)

type BookHandler interface {
	Get(ctx *gin.Context)
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
}

type bookHandler struct {
	authClient user.AuthServiceClient
	bookClient book.BookServiceClient
}

func NewBookHandler(ac user.AuthServiceClient, bc book.BookServiceClient) BookHandler {
	return &bookHandler{
		authClient: ac,
		bookClient: bc,
	}
}

// Get - 書籍情報取得 (ISBN指定) ※廃止予定
func (h *bookHandler) Get(ctx *gin.Context) {
	c := util.SetMetadata(ctx)
	isbn := ctx.Param("bookID")

	eg, ectx := errgroup.WithContext(c)

	var a *entity.Auth
	eg.Go(func() error {
		out, err := h.authClient.GetAuth(ectx, &user.Empty{})
		if err != nil {
			return err
		}
		a = entity.NewAuth(out.Auth)
		return nil
	})

	var b *entity.Book
	eg.Go(func() error {
		in := &book.GetBookByIsbnRequest{
			Isbn: isbn,
		}
		out, err := h.bookClient.GetBookByIsbn(ectx, in)
		if err != nil {
			return err
		}
		b = entity.NewBook(out.Book)
		return nil
	})

	if err := eg.Wait(); err != nil {
		util.ErrorHandling(ctx, err)
		return
	}

	in := &book.GetBookshelfRequest{
		BookId: b.Id,
		UserId: a.Id,
	}
	out, err := h.bookClient.GetBookshelf(c, in)
	if err != nil && !util.IsNotFound(err) {
		util.ErrorHandling(ctx, err)
		return
	}

	bs := entity.NewBookshelf(out.Bookshelf)
	res := response.NewBookResponse(b, bs)
	ctx.JSON(http.StatusOK, res)
}

// Create - 書籍登録
func (h *bookHandler) Create(ctx *gin.Context) {
	c := util.SetMetadata(ctx)

	req := &request.CreateBookRequest{}
	if err := ctx.BindJSON(req); err != nil {
		util.ErrorHandling(ctx, entity.ErrBadRequest.New(err))
		return
	}

	authorNames := strings.Split(req.Author, "/")
	authorNameKanas := strings.Split(req.AuthorKana, "/")
	authors := make([]*book.CreateBookRequest_Author, len(authorNames))
	for i := range authorNames {
		author := &book.CreateBookRequest_Author{
			Name:     authorNames[i],
			NameKana: authorNameKanas[i],
		}

		authors[i] = author
	}

	in := &book.CreateBookRequest{
		Title:          req.Title,
		TitleKana:      req.TitleKana,
		Description:    req.ItemCaption,
		Isbn:           req.Isbn,
		Publisher:      req.PublisherName,
		PublishedOn:    req.SalesDate,
		ThumbnailUrl:   h.getThumbnailURLByRequest(req.SmallImageURL, req.MediumImageURL, req.LargeImageURL),
		RakutenUrl:     req.ItemURL,
		RakutenSize:    req.Size,
		RakutenGenreId: req.BooksGenreID,
		Authors:        authors,
	}
	out, err := h.bookClient.CreateBook(c, in)
	if err != nil {
		util.ErrorHandling(ctx, err)
		return
	}

	b := entity.NewBook(out.Book)
	res := response.NewBookResponse(b, nil)
	ctx.JSON(http.StatusOK, res)
}

// Update - 書籍更新
func (h *bookHandler) Update(ctx *gin.Context) {
	c := util.SetMetadata(ctx)

	req := &request.UpdateBookRequest{}
	if err := ctx.BindJSON(req); err != nil {
		util.ErrorHandling(ctx, entity.ErrBadRequest.New(err))
		return
	}

	authorNames := strings.Split(req.Author, "/")
	authorNameKanas := strings.Split(req.AuthorKana, "/")
	authors := make([]*book.UpdateBookRequest_Author, len(authorNames))
	for i := range authorNames {
		author := &book.UpdateBookRequest_Author{
			Name:     authorNames[i],
			NameKana: authorNameKanas[i],
		}

		authors[i] = author
	}

	in := &book.UpdateBookRequest{
		Title:          req.Title,
		TitleKana:      req.TitleKana,
		Description:    req.ItemCaption,
		Isbn:           req.Isbn,
		Publisher:      req.PublisherName,
		PublishedOn:    req.SalesDate,
		ThumbnailUrl:   h.getThumbnailURLByRequest(req.SmallImageURL, req.MediumImageURL, req.LargeImageURL),
		RakutenUrl:     req.ItemURL,
		RakutenSize:    req.Size,
		RakutenGenreId: req.BooksGenreID,
		Authors:        authors,
	}
	out, err := h.bookClient.UpdateBook(c, in)
	if err != nil {
		util.ErrorHandling(ctx, err)
		return
	}

	b := entity.NewBook(out.Book)
	res := response.NewBookResponse(b, nil)
	ctx.JSON(http.StatusOK, res)
}

func (h *bookHandler) getThumbnailURLByRequest(smallImageURL, mediumImageURL, largeImageURL string) string {
	if largeImageURL != "" {
		return largeImageURL
	}

	if mediumImageURL != "" {
		return mediumImageURL
	}

	if smallImageURL != "" {
		return smallImageURL
	}

	return ""
}
