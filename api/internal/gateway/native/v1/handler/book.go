package handler

import (
	"context"
	"fmt"
	"net/http"
	"strings"

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

// getBook - 書籍情報取得 (ISBN指定) ※廃止予定
func (h *apiV1Handler) getBook(ctx *gin.Context) {
	c := util.SetMetadata(ctx)

	isbn := ctx.Param("bookID")

	eg, ectx := errgroup.WithContext(c)
	var a *gentity.Auth
	eg.Go(func() (err error) {
		a, err = h.authGetAuth(ectx)
		return
	})
	var b *gentity.Book
	eg.Go(func() (err error) {
		b, err = h.bookGetBookByIsbn(ectx, isbn)
		return
	})
	if err := eg.Wait(); err != nil {
		util.ErrorHandling(ctx, err)
		return
	}

	bs, err := h.bookGetBookshelf(c, b.Id, a.Id)
	if err != nil {
		util.ErrorHandling(ctx, err)
		return
	}
	res := &response.BookResponse{
		Book: entity.NewBook(b, bs),
	}
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

	res := &response.BookResponse{
		Book: entity.NewBook(b, nil),
	}
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

	res := &response.BookResponse{
		Book: entity.NewBook(b, nil),
	}
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

func (h *apiV1Handler) bookMultiGetBooks(ctx context.Context, bookIDs []int64) (gentity.Books, error) {
	in := &book.MultiGetBooksRequest{
		BookIds: bookIDs,
	}
	out, err := h.Book.MultiGetBooks(ctx, in)
	if err != nil {
		return nil, err
	}

	return gentity.NewBooks(out.Books), nil
}

func (h *apiV1Handler) bookGetBook(ctx context.Context, bookID int64) (*gentity.Book, error) {
	in := &book.GetBookRequest{
		BookId: bookID,
	}
	out, err := h.Book.GetBook(ctx, in)
	if err != nil {
		return nil, err
	}
	if out.Book == nil {
		err := fmt.Errorf("book is not found: %d", bookID)
		return nil, exception.ErrNotFound.New(err)
	}

	return gentity.NewBook(out.Book), nil
}

func (h *apiV1Handler) bookGetBookByIsbn(ctx context.Context, isbn string) (*gentity.Book, error) {
	in := &book.GetBookByIsbnRequest{
		Isbn: isbn,
	}
	out, err := h.Book.GetBookByIsbn(ctx, in)
	if err != nil {
		return nil, err
	}
	if out.Book == nil {
		err := fmt.Errorf("book is not found: %d", isbn)
		return nil, exception.ErrNotFound.New(err)
	}

	return gentity.NewBook(out.Book), nil
}

func (h *apiV1Handler) bookListBookReview(
	ctx context.Context, bookID, limit, offset int64,
) (gentity.Reviews, int64, error) {
	in := &book.ListBookReviewRequest{
		BookId: bookID,
		Limit:  limit,
		Offset: offset,
	}
	out, err := h.Book.ListBookReview(ctx, in)
	if err != nil {
		return nil, 0, err
	}

	return gentity.NewReviews(out.Reviews), out.Total, nil
}

func (h *apiV1Handler) bookListUserReview(
	ctx context.Context, userID string, limit, offset int64,
) (gentity.Reviews, int64, error) {
	in := &book.ListUserReviewRequest{
		UserId: userID,
		Limit:  limit,
		Offset: offset,
	}
	out, err := h.Book.ListUserReview(ctx, in)
	if err != nil {
		return nil, 0, err
	}

	return gentity.NewReviews(out.Reviews), out.Total, nil
}

func (h *apiV1Handler) bookGetReview(ctx context.Context, reviewID int64) (*gentity.Review, error) {
	in := &book.GetReviewRequest{
		ReviewId: reviewID,
	}
	out, err := h.Book.GetReview(ctx, in)
	if err != nil {
		return nil, err
	}
	if out.Review == nil {
		err := fmt.Errorf("review is not found: %d", reviewID)
		return nil, exception.ErrNotFound.New(err)
	}

	return gentity.NewReview(out.Review), nil
}
