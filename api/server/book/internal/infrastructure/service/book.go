package service

import (
	"context"
	"time"

	"github.com/calmato/gran-book/api/server/book/internal/domain"
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

func (s *bookService) List(ctx context.Context, q *domain.ListQuery) ([]*book.Book, error) {
	return s.bookRepository.List(ctx, q)
}

func (s *bookService) ListBookshelf(ctx context.Context, q *domain.ListQuery) ([]*book.Bookshelf, error) {
	return s.bookRepository.ListBookshelf(ctx, q)
}

func (s *bookService) ListReview(ctx context.Context, q *domain.ListQuery) ([]*book.Review, error) {
	return s.bookRepository.ListReview(ctx, q)
}

func (s *bookService) Show(ctx context.Context, bookID int) (*book.Book, error) {
	return s.bookRepository.Show(ctx, bookID)
}

func (s *bookService) ListCount(ctx context.Context, q *domain.ListQuery) (int, error) {
	return s.bookRepository.ListCount(ctx, q)
}

func (s *bookService) ListBookshelfCount(ctx context.Context, q *domain.ListQuery) (int, error) {
	return s.bookRepository.ListBookshelfCount(ctx, q)
}

func (s *bookService) ListReviewCount(ctx context.Context, q *domain.ListQuery) (int, error) {
	return s.bookRepository.ListReviewCount(ctx, q)
}

func (s *bookService) ShowByIsbn(ctx context.Context, isbn string) (*book.Book, error) {
	return s.bookRepository.ShowByIsbn(ctx, isbn)
}

func (s *bookService) ShowBookshelfByUserIDAndBookID(
	ctx context.Context, userID string, bookID int,
) (*book.Bookshelf, error) {
	return s.bookRepository.ShowBookshelfByUserIDAndBookID(ctx, userID, bookID)
}

func (s *bookService) ShowReview(ctx context.Context, reviewID int) (*book.Review, error) {
	return s.bookRepository.ShowReview(ctx, reviewID)
}

func (s *bookService) ShowReviewByUserIDAndBookID(
	ctx context.Context, userID string, bookID int,
) (*book.Review, error) {
	return s.bookRepository.ShowReviewByUserIDAndBookID(ctx, userID, bookID)
}

func (s *bookService) Create(ctx context.Context, b *book.Book) error {
	err := s.associate(ctx, b)
	if err != nil {
		return err
	}

	current := time.Now().Local()

	b.CreatedAt = current
	b.UpdatedAt = current

	err = s.associate(ctx, b)
	if err != nil {
		return err
	}

	return s.bookRepository.Create(ctx, b)
}

func (s *bookService) CreateBookshelf(ctx context.Context, b *book.Bookshelf) error {
	current := time.Now().Local()

	b.CreatedAt = current
	b.UpdatedAt = current

	return s.bookRepository.CreateBookshelf(ctx, b)
}

func (s *bookService) Update(ctx context.Context, b *book.Book) error {
	err := s.associate(ctx, b)
	if err != nil {
		return err
	}

	current := time.Now().Local()

	b.UpdatedAt = current

	return s.bookRepository.Update(ctx, b)
}

func (s *bookService) UpdateBookshelf(ctx context.Context, b *book.Bookshelf) error {
	current := time.Now().Local()

	b.UpdatedAt = current

	return s.bookRepository.UpdateBookshelf(ctx, b)
}

func (s *bookService) MultipleCreate(ctx context.Context, bs []*book.Book) error {
	current := time.Now().Local()

	for _, b := range bs {
		err := s.associate(ctx, b)
		if err != nil {
			return err
		}

		b.CreatedAt = current
		b.UpdatedAt = current
	}

	return s.bookRepository.MultipleCreate(ctx, bs)
}

func (s *bookService) MultipleUpdate(ctx context.Context, bs []*book.Book) error {
	current := time.Now().Local()

	for _, b := range bs {
		err := s.associate(ctx, b)
		if err != nil {
			return err
		}

		b.UpdatedAt = current
	}

	return s.bookRepository.MultipleUpdate(ctx, bs)
}

func (s *bookService) Delete(ctx context.Context, bookID int) error {
	return s.bookRepository.Delete(ctx, bookID)
}

func (s *bookService) DeleteBookshelf(ctx context.Context, bookshelfID int) error {
	return s.bookRepository.DeleteBookshelf(ctx, bookshelfID)
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

func (s *bookService) ValidationReview(ctx context.Context, rv *book.Review) error {
	return s.bookDomainValidation.Review(ctx, rv)
}

func (s *bookService) associate(ctx context.Context, b *book.Book) error {
	current := time.Now().Local()

	for _, a := range b.Authors {
		a.CreatedAt = current
		a.UpdatedAt = current

		err := s.bookRepository.ShowOrCreateAuthor(ctx, a)
		if err != nil {
			return err
		}
	}

	return nil
}
