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

func (s *bookService) ShowByIsbn(ctx context.Context, isbn string) (*book.Book, error) {
	b, err := s.bookRepository.ShowByIsbn(ctx, isbn)
	if err != nil {
		return nil, err
	}

	if b == nil || b.ID == 0 {
		err := xerrors.New("Book is nil.")
		return nil, exception.NotFound.New(err)
	}

	p, err := s.bookRepository.ShowPublisherByBookID(ctx, b.ID)
	if err != nil {
		return nil, err
	}

	as, err := s.bookRepository.ShowAuthorsByBookID(ctx, b.ID)
	if err != nil {
		return nil, err
	}

	cs, err := s.bookRepository.ShowCategoriesByBookID(ctx, b.ID)
	if err != nil {
		return nil, err
	}

	b.PublisherID = p.ID
	b.Publisher = p
	b.Authors = as
	b.Categories = cs

	return b, nil
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

func (s *bookService) CreatePublisher(ctx context.Context, p *book.Publisher) error {
	current := time.Now()

	p.CreatedAt = current
	p.UpdatedAt = current

	return s.bookRepository.CreatePublisher(ctx, p)
}

func (s *bookService) Update(ctx context.Context, b *book.Book) error {
	current := time.Now()

	b.UpdatedAt = current

	s.associate(ctx, b)
	return s.bookRepository.Update(ctx, b)
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

func (s *bookService) ValidationCategory(ctx context.Context, c *book.Category) error {
	return s.bookDomainValidation.Category(ctx, c)
}

func (s *bookService) ValidationPublisher(ctx context.Context, p *book.Publisher) error {
	return s.bookDomainValidation.Publisher(ctx, p)
}

func (s *bookService) associate(ctx context.Context, b *book.Book) {
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
}
