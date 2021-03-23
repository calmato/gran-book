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
	Create(ctx context.Context, in *input.BookItem) (*book.Book, error)
	Update(ctx context.Context, in *input.BookItem) (*book.Book, error)
	MultipleCreateAndUpdate(ctx context.Context, in *input.CreateAndUpdateBooks) ([]*book.Book, error)
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

func (a *bookApplication) Create(ctx context.Context, in *input.BookItem) (*book.Book, error) {
	err := a.bookRequestValidation.BookItem(in)
	if err != nil {
		return nil, err
	}

	b, err := a.initializeBook(ctx, in)
	if err != nil {
		return nil, err
	}

	err = a.bookService.Create(ctx, b)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func (a *bookApplication) Update(ctx context.Context, in *input.BookItem) (*book.Book, error) {
	err := a.bookRequestValidation.BookItem(in)
	if err != nil {
		return nil, err
	}

	b, err := a.bookService.ShowByIsbn(ctx, in.Isbn)
	if err != nil {
		return nil, err
	}

	newBook, err := a.initializeBook(ctx, in)
	if err != nil {
		return nil, err
	}

	if newBook.Publisher != nil {
		b.PublisherID = newBook.Publisher.ID
		b.Publisher = newBook.Publisher
	}

	b.Authors = newBook.Authors
	b.Categories = newBook.Categories

	err = a.bookService.Update(ctx, b)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func (a *bookApplication) MultipleCreateAndUpdate(
	ctx context.Context, in *input.CreateAndUpdateBooks,
) ([]*book.Book, error) {
	err := a.bookRequestValidation.CreateAndUpdateBooks(in)
	if err != nil {
		return nil, err
	}

	bs := []*book.Book{}
	nonCreatedBooks := []*book.Book{}
	updateRequiredBooks := []*book.Book{}

	for _, v := range in.Books {
		b, _ := a.bookService.ShowByIsbn(ctx, v.Isbn)
		if b == nil {
			b, err := a.initializeBook(ctx, v)
			if err != nil {
				return nil, err
			}

			nonCreatedBooks = append(nonCreatedBooks, b)
			continue
		}

		// 既存データとバージョンが異なっている場合のみ更新
		if b.Version != v.Version {
			newBook, err := a.initializeBook(ctx, v)
			if err != nil {
				return nil, err
			}

			if newBook.Publisher != nil {
				b.PublisherID = newBook.Publisher.ID
				b.Publisher = newBook.Publisher
			}

			b.Authors = newBook.Authors
			b.Categories = newBook.Categories

			updateRequiredBooks = append(updateRequiredBooks, b)
			continue
		}

		bs = append(bs, b)
	}

	err = a.bookService.MultipleCreate(ctx, nonCreatedBooks)
	if err != nil {
		return nil, err
	}

	err = a.bookService.MultipleUpdate(ctx, updateRequiredBooks)
	if err != nil {
		return nil, err
	}

	bs = append(bs, nonCreatedBooks...)
	bs = append(bs, updateRequiredBooks...)

	return bs, nil
}

func (a *bookApplication) initializeBook(ctx context.Context, in *input.BookItem) (*book.Book, error) {
	as := make([]*book.Author, len(in.Authors))
	for i, v := range in.Authors {
		author := &book.Author{
			Name: v.Name,
		}

		err := a.bookService.ValidationAuthor(ctx, author)
		if err != nil {
			return nil, err
		}

		as[i] = author
	}

	cs := make([]*book.Category, len(in.Categories))
	for i, v := range in.Categories {
		c := &book.Category{
			Name: v.Name,
		}

		err := a.bookService.ValidationCategory(ctx, c)
		if err != nil {
			return nil, err
		}

		cs[i] = c
	}

	b := &book.Book{
		Title:        in.Title,
		Description:  in.Description,
		Isbn:         in.Isbn,
		ThumbnailURL: in.ThumbnailURL,
		Version:      in.Version,
		PublishedOn:  datetime.StringToDate(in.PublishedOn),
		Authors:      as,
		Categories:   cs,
	}

	if in.Publisher != "" {
		p := &book.Publisher{
			Name: in.Publisher,
		}

		err := a.bookService.ValidationPublisher(ctx, p)
		if err != nil {
			return nil, err
		}

		b.Publisher = p
	}

	err := a.bookService.Validation(ctx, b)
	if err != nil {
		return nil, err
	}

	return b, nil
}
