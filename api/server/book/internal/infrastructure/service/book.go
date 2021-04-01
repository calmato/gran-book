package service

import (
	"context"
	"time"

	"github.com/calmato/gran-book/api/server/book/internal/domain/book"
	"github.com/calmato/gran-book/api/server/book/internal/domain/exception"
	"golang.org/x/xerrors"
)

type bookService struct {
	bookDomainValidation book.Validation
	bookRepository       book.Repository
}

// NewBookService - BookServiceの生成
func NewBookService(bdv book.Validation, br book.Repository) book.Service {
	return &bookService{
		bookDomainValidation: bdv,
		bookRepository:       br,
	}
}

func (s *bookService) Show(ctx context.Context, bookID int) (*book.Book, error) {
	return s.bookRepository.Show(ctx, bookID) // TODO: いったんカテゴリとかは取得しない
}

func (s *bookService) ShowByIsbn(ctx context.Context, isbn string) (*book.Book, error) {
	b, err := s.bookRepository.ShowByIsbn(ctx, isbn)
	if err != nil {
		return nil, err
	}

	if b == nil || b.ID == 0 {
		err := xerrors.New("Book is nil.")
		return nil, exception.NotFound.New(err)
	}

	as, err := s.bookRepository.ShowAuthorsByBookID(ctx, b.ID)
	if err != nil {
		return nil, err
	}

	cs, err := s.bookRepository.ShowCategoriesByBookID(ctx, b.ID)
	if err != nil {
		return nil, err
	}

	b.Authors = as
	b.Categories = cs

	return b, nil
}

func (s *bookService) ShowBookshelfByUserIDAndBookID(
	ctx context.Context, userID string, bookID int,
) (*book.Bookshelf, error) {
	return s.ShowBookshelfByUserIDAndBookID(ctx, userID, bookID)
}

func (s *bookService) Create(ctx context.Context, b *book.Book) error {
	current := time.Now()

	b.CreatedAt = current
	b.UpdatedAt = current

	s.associate(ctx, b)
	return s.bookRepository.Create(ctx, b)
}

func (s *bookService) CreateAuthor(ctx context.Context, a *book.Author) error {
	current := time.Now()

	a.CreatedAt = current
	a.UpdatedAt = current

	return s.bookRepository.CreateAuthor(ctx, a)
}

func (s *bookService) CreateBookshelf(ctx context.Context, b *book.Bookshelf) error {
	current := time.Now()

	b.CreatedAt = current
	b.UpdatedAt = current

	return s.bookRepository.CreateBookshelf(ctx, b)
}

func (s *bookService) CreateCategory(ctx context.Context, c *book.Category) error {
	current := time.Now()

	c.CreatedAt = current
	c.UpdatedAt = current

	return s.bookRepository.CreateCategory(ctx, c)
}

func (s *bookService) Update(ctx context.Context, b *book.Book) error {
	current := time.Now()

	b.UpdatedAt = current

	s.associate(ctx, b)
	return s.bookRepository.Update(ctx, b)
}

func (s *bookService) UpdateBookshelf(ctx context.Context, b *book.Bookshelf) error {
	current := time.Now()

	b.UpdatedAt = current

	return s.bookRepository.UpdateBookshelf(ctx, b)
}

func (s *bookService) MultipleCreate(ctx context.Context, bs []*book.Book) error {
	current := time.Now()

	for _, b := range bs {
		b.CreatedAt = current
		b.UpdatedAt = current

		s.associate(ctx, b)
	}

	return s.bookRepository.MultipleCreate(ctx, bs)
}

func (s *bookService) MultipleUpdate(ctx context.Context, bs []*book.Book) error {
	current := time.Now()

	for _, b := range bs {
		b.UpdatedAt = current

		s.associate(ctx, b)
	}

	return s.bookRepository.MultipleUpdate(ctx, bs)
}

func (s *bookService) Validation(ctx context.Context, b *book.Book) error {
	return s.bookDomainValidation.Book(ctx, b)
}

func (s *bookService) ValidationAuthor(ctx context.Context, a *book.Author) error {
	return s.bookDomainValidation.Author(ctx, a)
}

func (s *bookService) ValidationBookshelf(ctx context.Context, b *book.Bookshelf) error {
	return s.bookDomainValidation.Bookshelf(ctx, b)
}

func (s *bookService) ValidationCategory(ctx context.Context, c *book.Category) error {
	return s.bookDomainValidation.Category(ctx, c)
}

func (s *bookService) associate(ctx context.Context, b *book.Book) {
	for _, a := range b.Authors {
		_ = s.CreateAuthor(ctx, a)
	}

	for _, c := range b.Categories {
		_ = s.CreateCategory(ctx, c)
	}
}
