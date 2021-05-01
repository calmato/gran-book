package application

import (
	"context"

	"github.com/calmato/gran-book/api/server/book/internal/application/input"
	"github.com/calmato/gran-book/api/server/book/internal/application/output"
	"github.com/calmato/gran-book/api/server/book/internal/application/validation"
	"github.com/calmato/gran-book/api/server/book/internal/domain"
	"github.com/calmato/gran-book/api/server/book/internal/domain/book"
	"github.com/calmato/gran-book/api/server/book/lib/datetime"
)

// BookApplication - Bookアプリケーションのインターフェース
type BookApplication interface {
	ListBookshelf(ctx context.Context, in *input.ListBookshelf) ([]*book.Bookshelf, *output.ListQuery, error)
	Show(ctx context.Context, isbn string) (*book.Book, error)
	ShowBookshelf(ctx context.Context, userID string, bookID int) (*book.Bookshelf, error)
	Create(ctx context.Context, in *input.Book) (*book.Book, error)
	Update(ctx context.Context, in *input.Book) (*book.Book, error)
	CreateOrUpdateBookshelf(ctx context.Context, in *input.Bookshelf) (*book.Bookshelf, error)
	Delete(ctx context.Context, bookID int) error
	DeleteBookshelf(ctx context.Context, bookID int, uid string) error
}

type bookApplication struct {
	bookRequestValidation validation.BookRequestValidation
	bookService           book.Service
}

// NewBookApplication - BookApplicationの生成
func NewBookApplication(brv validation.BookRequestValidation, bs book.Service) BookApplication {
	return &bookApplication{
		bookRequestValidation: brv,
		bookService:           bs,
	}
}

func (a *bookApplication) ListBookshelf(
	ctx context.Context, in *input.ListBookshelf,
) ([]*book.Bookshelf, *output.ListQuery, error) {
	err := a.bookRequestValidation.ListBookshelf(in)
	if err != nil {
		return nil, nil, err
	}

	query := &domain.ListQuery{
		Limit:  in.Limit,
		Offset: in.Offset,
		Conditions: []*domain.QueryCondition{
			{
				Field:    "user_id",
				Operator: "==",
				Value:    in.UserID,
			},
		},
	}

	bss, err := a.bookService.ListBookshelf(ctx, query)
	if err != nil {
		return nil, nil, err
	}

	total, err := a.bookService.ListBookshelfCount(ctx, query)
	if err != nil {
		return nil, nil, err
	}

	out := &output.ListQuery{
		Limit:  in.Limit,
		Offset: in.Offset,
		Total:  total,
	}

	return bss, out, nil
}

func (a *bookApplication) Show(ctx context.Context, isbn string) (*book.Book, error) {
	return a.bookService.ShowByIsbn(ctx, isbn)
}

func (a *bookApplication) ShowBookshelf(ctx context.Context, userID string, bookID int) (*book.Bookshelf, error) {
	bs, err := a.bookService.ShowBookshelfByUserIDAndBookID(ctx, userID, bookID)
	if err != nil {
		return nil, err
	}

	rv, _ := a.bookService.ShowReviewByUserIDAndBookID(ctx, userID, bookID)

	bs.Review = rv

	return bs, nil
}

func (a *bookApplication) Create(ctx context.Context, in *input.Book) (*book.Book, error) {
	err := a.bookRequestValidation.Book(in)
	if err != nil {
		return nil, err
	}

	as := make([]*book.Author, len(in.Authors))
	for i, v := range in.Authors {
		author := &book.Author{
			Name:     v.Name,
			NameKana: v.NameKana,
		}

		err = a.bookService.ValidationAuthor(ctx, author)
		if err != nil {
			return nil, err
		}

		as[i] = author
	}

	b := &book.Book{
		Title:          in.Title,
		TitleKana:      in.TitleKana,
		Description:    in.Description,
		Isbn:           in.Isbn,
		Publisher:      in.Publisher,
		PublishedOn:    in.PublishedOn,
		ThumbnailURL:   in.ThumbnailURL,
		RakutenURL:     in.RakutenURL,
		RakutenGenreID: in.RakutenGenreID,
		Authors:        as,
	}

	err = a.bookService.Validation(ctx, b)
	if err != nil {
		return nil, err
	}

	err = a.bookService.Create(ctx, b)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func (a *bookApplication) Update(ctx context.Context, in *input.Book) (*book.Book, error) {
	err := a.bookRequestValidation.Book(in)
	if err != nil {
		return nil, err
	}

	b, err := a.bookService.ShowByIsbn(ctx, in.Isbn)
	if err != nil {
		return nil, err
	}

	as := make([]*book.Author, len(in.Authors))
	for i, v := range in.Authors {
		author := &book.Author{
			Name:     v.Name,
			NameKana: v.NameKana,
		}

		err = a.bookService.ValidationAuthor(ctx, author)
		if err != nil {
			return nil, err
		}

		as[i] = author
	}

	b.Title = in.Title
	b.TitleKana = in.TitleKana
	b.Description = in.Description
	b.Isbn = in.Isbn
	b.Publisher = in.Publisher
	b.PublishedOn = in.PublishedOn
	b.ThumbnailURL = in.ThumbnailURL
	b.RakutenURL = in.RakutenURL
	b.RakutenGenreID = in.RakutenGenreID
	b.Authors = as

	err = a.bookService.Validation(ctx, b)
	if err != nil {
		return nil, err
	}

	err = a.bookService.Update(ctx, b)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func (a *bookApplication) CreateOrUpdateBookshelf(
	ctx context.Context, in *input.Bookshelf,
) (*book.Bookshelf, error) {
	err := a.bookRequestValidation.Bookshelf(in)
	if err != nil {
		return nil, err
	}

	b, err := a.bookService.Show(ctx, in.BookID)
	if err != nil {
		return nil, err
	}

	rv, _ := a.bookService.ShowReviewByUserIDAndBookID(ctx, in.UserID, b.ID)
	bs, _ := a.bookService.ShowBookshelfByUserIDAndBookID(ctx, in.UserID, b.ID)
	if bs == nil {
		bs = &book.Bookshelf{}
	}

	bs.BookID = b.ID
	bs.UserID = in.UserID
	bs.Status = in.Status
	bs.ReadOn = datetime.StringToDate(in.ReadOn)
	bs.Book = b
	bs.Review = rv

	err = a.bookService.ValidationBookshelf(ctx, bs)
	if err != nil {
		return nil, err
	}

	if bs.ID == 0 {
		err = a.bookService.CreateBookshelf(ctx, bs)
		if err != nil {
			return nil, err
		}
	} else {
		err = a.bookService.UpdateBookshelf(ctx, bs)
		if err != nil {
			return nil, err
		}
	}

	return bs, nil
}

func (a *bookApplication) Delete(ctx context.Context, bookID int) error {
	b, err := a.bookService.Show(ctx, bookID)
	if err != nil {
		return err
	}

	return a.bookService.Delete(ctx, b.ID)
}

func (a *bookApplication) DeleteBookshelf(ctx context.Context, bookID int, uid string) error {
	b, err := a.bookService.ShowBookshelfByUserIDAndBookID(ctx, uid, bookID)
	if err != nil {
		return err
	}

	return a.bookService.DeleteBookshelf(ctx, b.ID)
}
