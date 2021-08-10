package application

import (
	"context"
	"time"

	"github.com/calmato/gran-book/api/server/book/internal/domain/book"
	"github.com/calmato/gran-book/api/server/book/pkg/database"
)

// BookApplication - Bookアプリケーションのインターフェース
type BookApplication interface {
	List(ctx context.Context, q *database.ListQuery) ([]*book.Book, int, error)
	ListBookshelf(ctx context.Context, q *database.ListQuery) ([]*book.Bookshelf, int, error)
	ListBookReview(ctx context.Context, bookID, limit, offset int) ([]*book.Review, int, error)
	ListUserReview(ctx context.Context, userID string, limit, offset int) ([]*book.Review, int, error)
	MultiGet(ctx context.Context, bookIDs []int) ([]*book.Book, error)
	Get(ctx context.Context, bookID int) (*book.Book, error)
	GetByIsbn(ctx context.Context, isbn string) (*book.Book, error)
	GetBookshelfByUserIDAndBookID(ctx context.Context, userID string, bookID int) (*book.Bookshelf, error)
	GetReview(ctx context.Context, reviewID int) (*book.Review, error)
	GetReviewByUserIDAndBookID(ctx context.Context, userID string, bookID int) (*book.Review, error)
	Create(ctx context.Context, b *book.Book) error
	CreateBookshelf(ctx context.Context, bs *book.Bookshelf) error
	Update(ctx context.Context, b *book.Book) error
	UpdateBookshelf(ctx context.Context, bs *book.Bookshelf) error
	MultipleCreate(ctx context.Context, bs []*book.Book) error
	MultipleUpdate(ctx context.Context, bs []*book.Book) error
	Delete(ctx context.Context, b *book.Book) error
	DeleteBookshelf(ctx context.Context, bs *book.Bookshelf) error
}

type bookApplication struct {
	bookDomainValidation book.Validation
	bookRepository       book.Repository
}

// NewBookApplication - BookApplicationの生成
func NewBookApplication(bdv book.Validation, br book.Repository) BookApplication {
	return &bookApplication{
		bookDomainValidation: bdv,
		bookRepository:       br,
	}
}

func (a *bookApplication) List(ctx context.Context, q *database.ListQuery) ([]*book.Book, int, error) {
	bs, err := a.bookRepository.List(ctx, q)
	if err != nil {
		return nil, 0, err
	}

	total, err := a.bookRepository.Count(ctx, q)
	if err != nil {
		return nil, 0, err
	}

	return bs, total, nil
}

func (a *bookApplication) ListBookshelf(ctx context.Context, q *database.ListQuery) ([]*book.Bookshelf, int, error) {
	bss, err := a.bookRepository.ListBookshelf(ctx, q)
	if err != nil {
		return nil, 0, err
	}

	total, err := a.bookRepository.CountBookshelf(ctx, q)
	if err != nil {
		return nil, 0, err
	}

	return bss, total, nil
}

func (a *bookApplication) ListBookReview(
	ctx context.Context, bookID int, limit int, offset int,
) ([]*book.Review, int, error) {
	q := &database.ListQuery{
		Limit:  limit,
		Offset: offset,
		Conditions: []*database.ConditionQuery{
			{
				Field:    "book_id",
				Operator: "==",
				Value:    bookID,
			},
		},
	}

	rs, err := a.bookRepository.ListReview(ctx, q)
	if err != nil {
		return nil, 0, err
	}

	total, err := a.bookRepository.CountReview(ctx, q)
	if err != nil {
		return nil, 0, err
	}

	return rs, total, nil
}

func (a *bookApplication) ListUserReview(
	ctx context.Context, userID string, limit int, offset int,
) ([]*book.Review, int, error) {
	q := &database.ListQuery{
		Limit:  limit,
		Offset: offset,
		Conditions: []*database.ConditionQuery{
			{
				Field:    "user_id",
				Operator: "==",
				Value:    userID,
			},
		},
	}

	rs, err := a.bookRepository.ListReview(ctx, q)
	if err != nil {
		return nil, 0, err
	}

	total, err := a.bookRepository.CountReview(ctx, q)
	if err != nil {
		return nil, 0, err
	}

	return rs, total, nil
}

func (a *bookApplication) MultiGet(ctx context.Context, bookIDs []int) ([]*book.Book, error) {
	return a.bookRepository.MultiGet(ctx, bookIDs)
}

func (a *bookApplication) Get(ctx context.Context, bookID int) (*book.Book, error) {
	return a.bookRepository.Get(ctx, bookID)
}

func (a *bookApplication) GetByIsbn(ctx context.Context, isbn string) (*book.Book, error) {
	return a.bookRepository.GetByIsbn(ctx, isbn)
}

func (a *bookApplication) GetBookshelfByUserIDAndBookID(
	ctx context.Context, userID string, bookID int,
) (*book.Bookshelf, error) {
	return a.bookRepository.GetBookshelfByUserIDAndBookID(ctx, userID, bookID)
}

func (a *bookApplication) GetReview(ctx context.Context, reviewID int) (*book.Review, error) {
	return a.bookRepository.GetReview(ctx, reviewID)
}

func (a *bookApplication) GetReviewByUserIDAndBookID(
	ctx context.Context, userID string, bookID int,
) (*book.Review, error) {
	return a.bookRepository.GetReviewByUserIDAndBookID(ctx, userID, bookID)
}

func (a *bookApplication) Create(ctx context.Context, b *book.Book) error {
	current := time.Now().Local()

	for _, author := range b.Authors {
		err := a.bookDomainValidation.Author(ctx, author)
		if err != nil {
			return err
		}

		if author.ID == 0 {
			author.CreatedAt = current
		}

		author.UpdatedAt = current
	}

	err := a.bookDomainValidation.Book(ctx, b)
	if err != nil {
		return err
	}

	b.CreatedAt = current
	b.UpdatedAt = current

	return a.bookRepository.Create(ctx, b)
}

func (a *bookApplication) CreateBookshelf(ctx context.Context, bs *book.Bookshelf) error {
	current := time.Now().Local()

	if bs.Review != nil {
		err := a.bookDomainValidation.Review(ctx, bs.Review)
		if err != nil {
			return err
		}

		if bs.Review.ID == 0 {
			bs.Review.CreatedAt = current
		}

		bs.Review.UpdatedAt = current
	}

	err := a.bookDomainValidation.Bookshelf(ctx, bs)
	if err != nil {
		return err
	}

	bs.CreatedAt = current
	bs.UpdatedAt = current

	return a.bookRepository.CreateBookshelf(ctx, bs)
}

func (a *bookApplication) Update(ctx context.Context, b *book.Book) error {
	current := time.Now().Local()

	for _, author := range b.Authors {
		err := a.bookDomainValidation.Author(ctx, author)
		if err != nil {
			return err
		}

		if author.ID == 0 {
			author.CreatedAt = current
		}

		author.UpdatedAt = current
	}

	err := a.bookDomainValidation.Book(ctx, b)
	if err != nil {
		return err
	}

	b.UpdatedAt = current

	return a.bookRepository.Update(ctx, b)
}

func (a *bookApplication) UpdateBookshelf(ctx context.Context, bs *book.Bookshelf) error {
	current := time.Now().Local()

	if bs.Review != nil {
		err := a.bookDomainValidation.Review(ctx, bs.Review)
		if err != nil {
			return err
		}

		if bs.Review.ID == 0 {
			bs.Review.CreatedAt = current
		}

		bs.Review.UpdatedAt = current
	}

	err := a.bookDomainValidation.Bookshelf(ctx, bs)
	if err != nil {
		return err
	}

	bs.UpdatedAt = current

	return a.bookRepository.CreateBookshelf(ctx, bs)
}

func (a *bookApplication) MultipleCreate(ctx context.Context, bs []*book.Book) error {
	current := time.Now().Local()

	for _, b := range bs {
		for _, author := range b.Authors {
			err := a.bookDomainValidation.Author(ctx, author)
			if err != nil {
				return err
			}

			if author.ID == 0 {
				author.CreatedAt = current
			}

			author.UpdatedAt = current
		}

		err := a.bookDomainValidation.Book(ctx, b)
		if err != nil {
			return err
		}

		b.CreatedAt = current
		b.UpdatedAt = current
	}

	return a.bookRepository.MultipleCreate(ctx, bs)
}

func (a *bookApplication) MultipleUpdate(ctx context.Context, bs []*book.Book) error {
	current := time.Now().Local()

	for _, b := range bs {
		for _, author := range b.Authors {
			err := a.bookDomainValidation.Author(ctx, author)
			if err != nil {
				return err
			}

			if author.ID == 0 {
				author.CreatedAt = current
			}

			author.UpdatedAt = current
		}

		err := a.bookDomainValidation.Book(ctx, b)
		if err != nil {
			return err
		}

		b.UpdatedAt = current
	}

	return a.bookRepository.MultipleUpdate(ctx, bs)
}

func (a *bookApplication) Delete(ctx context.Context, b *book.Book) error {
	return a.bookRepository.Delete(ctx, b.ID)
}

func (a *bookApplication) DeleteBookshelf(ctx context.Context, bs *book.Bookshelf) error {
	return a.bookRepository.DeleteBookshelf(ctx, bs.ID)
}
