package application

import (
	"context"

	"github.com/calmato/gran-book/api/server/book/internal/application/input"
	"github.com/calmato/gran-book/api/server/book/internal/application/validation"
	"github.com/calmato/gran-book/api/server/book/internal/domain/book"
	"github.com/calmato/gran-book/api/server/book/lib/datetime"
)

// BookApplication - Bookアプリケーションのインターフェース
type BookApplication interface {
	Create(ctx context.Context, in *input.CreateBook) (*book.Book, error)
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
func (a *bookApplication) Create(ctx context.Context, in *input.CreateBook) (*book.Book, error) {
	err := a.bookRequestValidation.CreateBook(in)
	if err != nil {
		return nil, err
	}

	p := &book.Publisher{
		Name: in.Publisher,
	}

	as := make([]*book.Author, len(in.Authors))
	for i, v := range in.Authors {
		author := &book.Author{
			Name: v.Name,
		}

		as[i] = author
	}

	cs := make([]*book.Category, len(in.Categories))
	for i, v := range in.Categories {
		c := &book.Category{
			Name: v.Name,
		}

		cs[i] = c
	}

	b := &book.Book{
		Title:        in.Title,
		Description:  in.Description,
		Isbn:         in.Isbn,
		ThumbnailURL: in.ThumbnailURL,
		Version:      in.Version,
		PublishedOn:  datetime.StringToTime(in.PublishedOn),
		Publisher:    p,
		Authors:      as,
		Categories:   cs,
	}

	err = a.bookService.Create(ctx, b)
	if err != nil {
		return nil, err
	}

	return b, nil
}
