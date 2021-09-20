package application

import (
	"context"
	"errors"
	"time"

	"github.com/calmato/gran-book/api/server/book/internal/domain/book"
	"github.com/calmato/gran-book/api/server/book/internal/domain/exception"
	"github.com/calmato/gran-book/api/server/book/pkg/database"
	"github.com/calmato/gran-book/api/server/book/pkg/datetime"
)

// BookApplication - Bookアプリケーションのインターフェース
type BookApplication interface {
	List(ctx context.Context, q *database.ListQuery) (book.Books, int, error)
	ListBookshelf(ctx context.Context, q *database.ListQuery) (book.Bookshelves, int, error)
	ListBookReview(ctx context.Context, bookID, limit, offset int) (book.Reviews, int, error)
	ListUserReview(ctx context.Context, userID string, limit, offset int) (book.Reviews, int, error)
	ListUserMonthlyResult(ctx context.Context, userID, since, until string) (book.MonthlyResults, error)
	MultiGet(ctx context.Context, bookIDs []int) (book.Books, error)
	Get(ctx context.Context, bookID int) (*book.Book, error)
	GetByIsbn(ctx context.Context, isbn string) (*book.Book, error)
	GetBookshelfByUserIDAndBookID(ctx context.Context, userID string, bookID int) (*book.Bookshelf, error)
	GetBookshelfByUserIDAndBookIDWithRelated(ctx context.Context, userID string, bookID int) (*book.Bookshelf, error)
	GetReview(ctx context.Context, reviewID int) (*book.Review, error)
	GetReviewByUserIDAndBookID(ctx context.Context, userID string, bookID int) (*book.Review, error)
	Create(ctx context.Context, b *book.Book) error
	CreateBookshelf(ctx context.Context, bs *book.Bookshelf) error
	Update(ctx context.Context, b *book.Book) error
	UpdateBookshelf(ctx context.Context, bs *book.Bookshelf) error
	CreateOrUpdateBookshelf(ctx context.Context, bs *book.Bookshelf) error
	MultipleCreate(ctx context.Context, bs book.Books) error
	MultipleUpdate(ctx context.Context, bs book.Books) error
	Delete(ctx context.Context, b *book.Book) error
	DeleteBookshelf(ctx context.Context, bs *book.Bookshelf) error
}

type bookApplication struct {
	bookDomainValidation book.Validation
	bookRepository       book.Repository
}

var errInvalidDateFormat = errors.New("application: invalid date format")

// NewBookApplication - BookApplicationの生成
func NewBookApplication(bdv book.Validation, br book.Repository) BookApplication {
	return &bookApplication{
		bookDomainValidation: bdv,
		bookRepository:       br,
	}
}

func (a *bookApplication) List(ctx context.Context, q *database.ListQuery) (book.Books, int, error) {
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

func (a *bookApplication) ListBookshelf(ctx context.Context, q *database.ListQuery) (book.Bookshelves, int, error) {
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
) (book.Reviews, int, error) {
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
) (book.Reviews, int, error) {
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

func (a *bookApplication) ListUserMonthlyResult(
	ctx context.Context, userID, sinceDate, untilDate string,
) (book.MonthlyResults, error) {
	since, _ := datetime.ParseDate(sinceDate)
	until, _ := datetime.ParseDate(untilDate)
	if since.IsZero() || until.IsZero() {
		return nil, exception.InvalidRequestValidation.New(errInvalidDateFormat)
	}

	return a.bookRepository.AggregateReadTotal(ctx, userID, since, until)
}

func (a *bookApplication) MultiGet(ctx context.Context, bookIDs []int) (book.Books, error) {
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

func (a *bookApplication) GetBookshelfByUserIDAndBookIDWithRelated(
	ctx context.Context, userID string, bookID int,
) (*book.Bookshelf, error) {
	b, err := a.bookRepository.Get(ctx, bookID)
	if err != nil {
		return nil, err
	}

	bs, err := a.bookRepository.GetBookshelfByUserIDAndBookID(ctx, userID, bookID)
	if err != nil {
		if err.(exception.CustomError).Code() != exception.NotFound {
			return nil, err
		}

		bs = &book.Bookshelf{}
	}

	r, err := a.bookRepository.GetReviewByUserIDAndBookID(ctx, userID, bookID)
	if err != nil {
		if err.(exception.CustomError).Code() != exception.NotFound {
			return nil, err
		}

		r = &book.Review{}
	}

	bs.Book = b
	bs.Review = r

	return bs, nil
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

	return a.bookRepository.UpdateBookshelf(ctx, bs)
}

func (a *bookApplication) CreateOrUpdateBookshelf(ctx context.Context, bs *book.Bookshelf) error {
	if bs.ID == 0 {
		return a.CreateBookshelf(ctx, bs)
	}

	return a.UpdateBookshelf(ctx, bs)
}

func (a *bookApplication) MultipleCreate(ctx context.Context, bs book.Books) error {
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

func (a *bookApplication) MultipleUpdate(ctx context.Context, bs book.Books) error {
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
