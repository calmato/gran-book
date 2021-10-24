package v1

import (
	"net/http"
	"strings"

	gentity "github.com/calmato/gran-book/api/service/internal/gateway/entity"
	request "github.com/calmato/gran-book/api/service/internal/gateway/native/request/v1"
	response "github.com/calmato/gran-book/api/service/internal/gateway/native/response/v1"
	"github.com/calmato/gran-book/api/service/internal/gateway/util"
	"github.com/calmato/gran-book/api/service/pkg/exception"
	"github.com/calmato/gran-book/api/service/proto/book"
	"github.com/calmato/gran-book/api/service/proto/user"
	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
)

// getBook - 書籍情報取得 (ISBN指定) ※廃止予定
func (h *apiV1Handler) getBook(ctx *gin.Context) {
	c := util.SetMetadata(ctx)
	isbn := ctx.Param("bookID")

	eg, ectx := errgroup.WithContext(c)

	var a *gentity.Auth
	eg.Go(func() error {
		out, err := h.Auth.GetAuth(ectx, &user.Empty{})
		if err != nil {
			return err
		}
		a = gentity.NewAuth(out.Auth)
		return nil
	})

	var b *gentity.Book
	eg.Go(func() error {
		in := &book.GetBookByIsbnRequest{
			Isbn: isbn,
		}
		out, err := h.Book.GetBookByIsbn(ectx, in)
		if err != nil {
			return err
		}
		b = gentity.NewBook(out.Book)
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
	out, err := h.Book.GetBookshelf(c, in)
	if err != nil && !util.IsNotFound(err) {
		util.ErrorHandling(ctx, err)
		return
	}

	bs := gentity.NewBookshelf(out.Bookshelf)
	res := response.NewBookResponse(b, bs)
	ctx.JSON(http.StatusOK, res)
}

// createBook - 書籍登録
func (h *apiV1Handler) createBook(ctx *gin.Context) {
	c := util.SetMetadata(ctx)

	req := &request.CreateBookRequest{}
	if err := ctx.BindJSON(req); err != nil {
		util.ErrorHandling(ctx, exception.ErrInvalidArgument.New(err))
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
	out, err := h.Book.CreateBook(c, in)
	if err != nil {
		util.ErrorHandling(ctx, err)
		return
	}

	b := gentity.NewBook(out.Book)
	res := response.NewBookResponse(b, nil)
	ctx.JSON(http.StatusOK, res)
}

// updateBook - 書籍更新
func (h *apiV1Handler) updateBook(ctx *gin.Context) {
	c := util.SetMetadata(ctx)

	req := &request.UpdateBookRequest{}
	if err := ctx.BindJSON(req); err != nil {
		util.ErrorHandling(ctx, exception.ErrInvalidArgument.New(err))
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
	out, err := h.Book.UpdateBook(c, in)
	if err != nil {
		util.ErrorHandling(ctx, err)
		return
	}

	b := gentity.NewBook(out.Book)
	res := response.NewBookResponse(b, nil)
	ctx.JSON(http.StatusOK, res)
}

func (h *apiV1Handler) getThumbnailURLByRequest(smallImageURL, mediumImageURL, largeImageURL string) string {
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
