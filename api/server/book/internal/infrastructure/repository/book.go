package repository

import (
	"context"
	"strings"

	"github.com/calmato/gran-book/api/server/book/internal/domain/book"
	"github.com/calmato/gran-book/api/server/book/internal/domain/exception"
	"gorm.io/gorm"
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

func (r *bookRepository) ShowPublisherByBookID(ctx context.Context, bookID int) (*book.Publisher, error) {
	p := &book.Publisher{}

	columns := []string{
		"publishers.id",
		"publishers.name",
		"publishers.created_at",
		"publishers.updated_at",
	}
	sql := r.client.db.
		Table("publishers").
		Select(strings.Join(columns, ", ")).
		Joins("LEFT JOIN books ON publishers.id = books.publisher_id").
		Where("books.id = ?", bookID)

	err := sql.First(&p).Error
	if err != nil {
		return nil, exception.ErrorInDatastore.New(err)
	}

	return p, nil
}

func (r *bookRepository) ShowAuthorsByBookID(ctx context.Context, bookID int) ([]*book.Author, error) {
	as := []*book.Author{}

	columns := []string{
		"authors.id",
		"authors.name",
		"authors.created_at",
		"authors.updated_at",
	}
	sql := r.client.db.
		Table("authors").
		Select(strings.Join(columns, ", ")).
		Joins("LEFT JOIN authors_books ON authors_books.author_id = authors.id").
		Where("authors_books.book_id = ?", bookID)

	err := sql.Scan(&as).Error
	if err != nil {
		return nil, exception.ErrorInDatastore.New(err)
	}

	return as, nil
}

func (r *bookRepository) ShowCategoriesByBookID(ctx context.Context, bookID int) ([]*book.Category, error) {
	cs := []*book.Category{}

	columns := []string{
		"categories.id",
		"categories.name",
		"categories.created_at",
		"categories.updated_at",
	}
	sql := r.client.db.
		Table("categories").
		Select(strings.Join(columns, ", ")).
		Joins("LEFT JOIN books_categories ON books_categories.category_id = categories.id").
		Where("books_categories.book_id = ?", bookID)

	err := sql.Scan(&cs).Error
	if err != nil {
		return nil, exception.ErrorInDatastore.New(err)
	}

	return cs, nil
}

func (r *bookRepository) Create(ctx context.Context, b *book.Book) error {
	tx := r.client.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	if b.PublishedOn.IsZero() {
		err := tx.Omit(clause.Associations, "published_on").Create(&b).Error
		if err != nil {
			tx.Rollback()
			return exception.ErrorInDatastore.New(err)
		}
	} else {
		err := tx.Omit(clause.Associations).Create(&b).Error
		if err != nil {
			tx.Rollback()
			return exception.ErrorInDatastore.New(err)
		}
	}

	err := associate(tx, b)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
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
	tx := r.client.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	err := tx.Error
	if err != nil {
		return err
	}

	err = r.client.db.Omit(clause.Associations).Save(&b).Error
	if err != nil {
		return exception.ErrorInDatastore.New(err)
	}

	err = associate(tx, b)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (r *bookRepository) MultipleCreate(ctx context.Context, bs []*book.Book) error {
	tx := r.client.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	err := tx.Error
	if err != nil {
		return err
	}

	for _, b := range bs {
		if b.PublishedOn.IsZero() {
			err = tx.Omit(clause.Associations, "published_on").Create(b).Error
			if err != nil {
				tx.Rollback()
				return exception.ErrorInDatastore.New(err)
			}
		} else {
			err = tx.Omit(clause.Associations).Create(b).Error
			if err != nil {
				tx.Rollback()
				return exception.ErrorInDatastore.New(err)
			}
		}

		err = associate(tx, b)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit().Error
}

func (r *bookRepository) MultipleUpdate(ctx context.Context, bs []*book.Book) error {
	tx := r.client.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	err := tx.Error
	if err != nil {
		return err
	}

	for _, b := range bs {
		err = r.client.db.Omit(clause.Associations).Save(b).Error
		if err != nil {
			return exception.ErrorInDatastore.New(err)
		}

		err = associate(tx, b)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit().Error
}

func associate(tx *gorm.DB, b *book.Book) error {
	err := associateAuthor(tx, b)
	if err != nil {
		return err
	}

	err = associateCategory(tx, b)
	if err != nil {
		return err
	}

	return nil
}

func associateAuthor(tx *gorm.DB, b *book.Book) error {
	beforeAuthorIDs := []int{}

	// 既存の関連レコード取得
	db := tx.Table("authors_books").Select("author_id").Where("book_id = ?", b.ID)
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
			sql := "DELETE FROM authors_books WHERE book_id = ? AND author_id = ?"
			err := tx.Exec(sql, b.ID, authorID).Error
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

		err := tx.Table("authors_books").Create(&ba).Error
		if err != nil {
			return exception.ErrorInDatastore.New(err)
		}
	}

	return nil
}

func associateCategory(tx *gorm.DB, b *book.Book) error {
	beforeCategoryIDs := []int{}

	// 既存の関連レコード取得
	db := tx.Table("books_categories").Select("category_id").Where("book_id = ?", b.ID)
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
			sql := "DELETE FROM books_categories WHERE book_id = ? AND category_id = ?"
			err := tx.Exec(sql, b.ID, categoryID).Error
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

		err := tx.Table("books_categories").Create(&bc).Error
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
