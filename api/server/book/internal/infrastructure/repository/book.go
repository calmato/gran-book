package repository

import (
	"context"

	"github.com/calmato/gran-book/api/server/book/internal/domain/book"
	"github.com/calmato/gran-book/api/server/book/internal/domain/exception"
)

type bookRepository struct {
	client *Client
}

// NewBookRepository - BookRepositoryの生成
func NewBookRepository(c *Client) book.Repository {
	return &bookRepository{
		client: c,
	}
}

func (r *bookRepository) Create(ctx context.Context, b *book.Book) error {
	err := r.client.db.Create(&b).Error
	if err != nil {
		return exception.ErrorInDatastore.New(err)
	}

	if len(b.Authors) > 0 {
		r.client.db.Model(&b).Association("authors_books").Append(&b.Authors)
	}

	if len(b.Categories) > 0 {
		r.client.db.Model(&b).Association("books_categories").Append(&b.Categories)
	}

	return nil
}

func (r *bookRepository) CreateAuthor(ctx context.Context, a *book.Author) error {
	err := r.client.db.Table("authors").Where("name = ?", a.Name).FirstOrCreate(&a).Error
	if err != nil {
		return exception.ErrorInDatastore.New(err)
	}

	return nil
}

func (r *bookRepository) CreateBookshelf(ctx context.Context, b *book.Bookshelf) error {
	err := r.client.db.Create(&b).Error
	if err != nil {
		return exception.ErrorInDatastore.New(err)
	}

	return nil
}

func (r *bookRepository) CreateCategory(ctx context.Context, c *book.Category) error {
	err := r.client.db.Table("categories").Where("name = ?", c.Name).FirstOrCreate(&c).Error
	if err != nil {
		return exception.ErrorInDatastore.New(err)
	}

	return nil
}

func (r *bookRepository) CreatePublisher(ctx context.Context, p *book.Publisher) error {
	err := r.client.db.Table("publishers").Where("name = ?", p.Name).FirstOrCreate(&p).Error
	if err != nil {
		return exception.ErrorInDatastore.New(err)
	}

	return nil
}
