package service

import (
	"context"
	"time"

	"github.com/calmato/gran-book/api/server/book/internal/domain/book"
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

func (s *bookService) ShowByIsbn(ctx context.Context, isbn string) (*book.Book, error) {
	return s.bookRepository.ShowByIsbn(ctx, isbn)
}

func (s *bookService) ShowByTitleAndPublisher(
	ctx context.Context, title string, publisher string,
) (*book.Book, error) {
	return s.bookRepository.ShowByTitleAndPublisher(ctx, title, publisher)
}

func (s *bookService) Create(ctx context.Context, b *book.Book) error {
	err := s.bookDomainValidation.Book(ctx, b)
	if err != nil {
		return err
	}

	current := time.Now()

	b.CreatedAt = current
	b.UpdatedAt = current

	if b.Publisher != nil {
		_ = s.CreatePublisher(ctx, b.Publisher)
		b.PublisherID = b.Publisher.ID
	}

	for _, a := range b.Authors {
		_ = s.CreateAuthor(ctx, a)
	}

	for _, c := range b.Categories {
		_ = s.CreateCategory(ctx, c)
	}

	return s.bookRepository.Create(ctx, b)
}

func (s *bookService) CreateAuthor(ctx context.Context, a *book.Author) error {
	err := s.bookDomainValidation.Author(ctx, a)
	if err != nil {
		return err
	}

	current := time.Now()

	a.CreatedAt = current
	a.UpdatedAt = current

	return s.bookRepository.CreateAuthor(ctx, a)
}

func (s *bookService) CreateBookshelf(ctx context.Context, b *book.Bookshelf) error {
	err := s.bookDomainValidation.Bookshelf(ctx, b)
	if err != nil {
		return err
	}

	current := time.Now()

	b.CreatedAt = current
	b.UpdatedAt = current

	return s.bookRepository.CreateBookshelf(ctx, b)
}

func (s *bookService) CreateCategory(ctx context.Context, c *book.Category) error {
	err := s.bookDomainValidation.Category(ctx, c)
	if err != nil {
		return err
	}

	current := time.Now()

	c.CreatedAt = current
	c.UpdatedAt = current

	return s.bookRepository.CreateCategory(ctx, c)
}

func (s *bookService) CreatePublisher(ctx context.Context, p *book.Publisher) error {
	err := s.bookDomainValidation.Publisher(ctx, p)
	if err != nil {
		return err
	}

	current := time.Now()

	p.CreatedAt = current
	p.UpdatedAt = current

	return s.bookRepository.CreatePublisher(ctx, p)
}

func (s *bookService) Update(ctx context.Context, b *book.Book) error {
	err := s.bookDomainValidation.Book(ctx, b)
	if err != nil {
		return err
	}

	current := time.Now()

	b.UpdatedAt = current

	if b.Publisher != nil {
		_ = s.CreatePublisher(ctx, b.Publisher)
		b.PublisherID = b.Publisher.ID
	}

	for _, a := range b.Authors {
		_ = s.CreateAuthor(ctx, a)
	}

	for _, c := range b.Categories {
		_ = s.CreateCategory(ctx, c)
	}

	return s.bookRepository.Update(ctx, b)
}
