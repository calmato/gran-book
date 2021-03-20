package repository

import (
	"context"

	"github.com/calmato/gran-book/api/server/book/internal/domain"
	"github.com/calmato/gran-book/api/server/book/internal/domain/book"
	"github.com/calmato/gran-book/api/server/book/internal/domain/exception"
	"gorm.io/gorm/clause"
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

func (r *bookRepository) ShowByIsbn(ctx context.Context, isbn string) (*book.Book, error) {
	b := &book.Book{}

	err := r.client.db.First(b, "isbn = ?", isbn).Error
	if err != nil {
		return nil, exception.NotFound.New(err)
	}

	return b, nil
}

func (r *bookRepository) ShowByTitleAndPublisher(
	ctx context.Context, title string, publisher string,
) (*book.Book, error) {
	b := &book.Book{}

	sql := r.client.db.Table("books").Joins("LEFT JOIN publishers ON publishers.id = books.publisher_id")

	err := sql.First(b, "books.title = ? AND publishers.name = ?", title, publisher).Error
	if err != nil {
		return nil, exception.NotFound.New(err)
	}

	return b, nil
}

func (r *bookRepository) Create(ctx context.Context, b *book.Book) error {
	err := r.client.db.Omit(clause.Associations).Create(&b).Error
	if err != nil {
		return exception.ErrorInDatastore.New(err)
	}

	_ = r.AssociateAuthor(ctx, b)
	_ = r.AssociateCategory(ctx, b)

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

func (r *bookRepository) Update(ctx context.Context, b *book.Book) error {
	err := r.client.db.Omit(clause.Associations).Save(&b).Error
	if err != nil {
		return exception.ErrorInDatastore.New(err)
	}

	_ = r.AssociateAuthor(ctx, b)
	_ = r.AssociateCategory(ctx, b)

	return nil
}

func (r *bookRepository) AssociateAuthor(ctx context.Context, b *book.Book) error {
	beforeAuthorIDs := []int{}
	q := &domain.ListQuery{
		Conditions: []*domain.QueryCondition{{
			Field:    "book_id",
			Operator: "==",
			Value:    b.ID,
		}},
	}

	// 既存の関連レコード取得
	sql := r.client.db.Table("authors_books").Select("author_id")
	db := r.client.getListQuery(sql, q)

	err := db.Scan(&beforeAuthorIDs).Error
	if err != nil {
		return exception.ErrorInDatastore.New(err)
	}

	// 現在のAuthorID一覧の作成
	currentAuthorIDs := make([]int, len(b.Authors))
	for i, a := range b.Authors {
		currentAuthorIDs[i] = a.ID
	}

	// 不要なもの削除
	for _, authorID := range beforeAuthorIDs {
		if !isContain(authorID, currentAuthorIDs) {
			stmt := "DELETE FROM authors_books WHERE book_id = ? AND author_id = ?"
			err := r.client.db.Exec(stmt, b.ID, authorID).Error
			if err != nil {
				return exception.ErrorInDatastore.New(err)
			}
		}
	}

	// 既存レコードとしてない場合、新たに関連レコードの作成
	for _, a := range b.Authors {
		if isContain(a.ID, beforeAuthorIDs) {
			continue
		}

		ba := &book.BookAuthor{
			BookID:    b.ID,
			AuthorID:  a.ID,
			CreatedAt: b.CreatedAt,
			UpdatedAt: b.UpdatedAt,
		}

		err := r.client.db.Table("authors_books").Create(&ba).Error
		if err != nil {
			return exception.ErrorInDatastore.New(err)
		}
	}

	return nil
}

func (r *bookRepository) AssociateCategory(ctx context.Context, b *book.Book) error {
	beforeCategoryIDs := []int{}
	q := &domain.ListQuery{
		Conditions: []*domain.QueryCondition{{
			Field:    "book_id",
			Operator: "==",
			Value:    b.ID,
		}},
	}

	// 既存の関連レコード取得
	sql := r.client.db.Table("books_categories").Select("category_id")
	db := r.client.getListQuery(sql, q)

	err := db.Scan(&beforeCategoryIDs).Error
	if err != nil {
		return exception.ErrorInDatastore.New(err)
	}

	// 現在のCategoryID一覧の作成
	currentCategoryIDs := make([]int, len(b.Categories))
	for i, a := range b.Categories {
		currentCategoryIDs[i] = a.ID
	}

	// 不要なもの削除
	for _, categoryID := range beforeCategoryIDs {
		if !isContain(categoryID, currentCategoryIDs) {
			stmt := "DELETE FROM books_categories WHERE book_id = ? AND category_id = ?"
			err := r.client.db.Exec(stmt, b.ID, categoryID).Error
			if err != nil {
				return exception.ErrorInDatastore.New(err)
			}
		}
	}

	// 既存レコードとしてない場合、新たに関連レコードの作成
	for _, c := range b.Categories {
		if isContain(c.ID, beforeCategoryIDs) {
			continue
		}

		bc := &book.BookCategory{
			BookID:     b.ID,
			CategoryID: c.ID,
			CreatedAt:  b.CreatedAt,
			UpdatedAt:  b.UpdatedAt,
		}

		err := r.client.db.Table("books_categories").Create(&bc).Error
		if err != nil {
			return exception.ErrorInDatastore.New(err)
		}
	}

	return nil
}

func isContain(target int, ids []int) bool {
	for _, id := range ids {
		if target == id {
			return true
		}
	}

	return false
}
