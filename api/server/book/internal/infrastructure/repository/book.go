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

func (r *bookRepository) ListAuthorByBookID(ctx context.Context, bookID int) ([]*book.Author, error) {
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

func (r *bookRepository) Show(ctx context.Context, bookID int) (*book.Book, error) {
	b := &book.Book{}

	err := r.client.db.First(b, "id = ?", bookID).Error
	if err != nil {
		return nil, exception.NotFound.New(err)
	}

	return b, nil
}

func (r *bookRepository) ShowByIsbn(ctx context.Context, isbn string) (*book.Book, error) {
	b := &book.Book{}

	err := r.client.db.First(b, "isbn = ?", isbn).Error
	if err != nil {
		return nil, exception.NotFound.New(err)
	}

	return b, nil
}

func (r *bookRepository) ShowBookshelfByUserIDAndBookID(
	ctx context.Context, userID string, bookID int,
) (*book.Bookshelf, error) {
	b := &book.Bookshelf{}

	err := r.client.db.First(b, "user_id = ? AND book_id = ?", userID, bookID).Error
	if err != nil {
		return nil, exception.NotFound.New(err)
	}

	return b, nil
}

func (r *bookRepository) ShowOrCreateAuthor(ctx context.Context, a *book.Author) error {
	err := r.client.db.Table("authors").Where("name = ?", a.Name).FirstOrCreate(&a).Error
	if err != nil {
		return exception.ErrorInDatastore.New(err)
	}

	return nil
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

func (r *bookRepository) CreateBookshelf(ctx context.Context, b *book.Bookshelf) error {
	if b.ReadOn.IsZero() {
		err := r.client.db.Omit("read_on").Create(&b).Error
		if err != nil {
			return exception.ErrorInDatastore.New(err)
		}
	} else {
		err := r.client.db.Create(&b).Error
		if err != nil {
			return exception.ErrorInDatastore.New(err)
		}
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

	if err := tx.Error; err != nil {
		return err
	}

	if b.PublishedOn.IsZero() {
		err := tx.Omit(clause.Associations, "published_on").Save(&b).Error
		if err != nil {
			tx.Rollback()
			return exception.ErrorInDatastore.New(err)
		}
	} else {
		err := tx.Omit(clause.Associations).Save(&b).Error
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

func (r *bookRepository) UpdateBookshelf(ctx context.Context, b *book.Bookshelf) error {
	if b.ReadOn.IsZero() {
		err := r.client.db.Omit("read_on").Save(&b).Error
		if err != nil {
			return exception.ErrorInDatastore.New(err)
		}
	} else {
		err := r.client.db.Save(&b).Error
		if err != nil {
			return exception.ErrorInDatastore.New(err)
		}
	}

	return nil
}

func (r *bookRepository) MultipleCreate(ctx context.Context, bs []*book.Book) error {
	tx := r.client.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	for _, b := range bs {
		if b.PublishedOn.IsZero() {
			err := tx.Omit(clause.Associations, "published_on").Create(b).Error
			if err != nil {
				tx.Rollback()
				return exception.ErrorInDatastore.New(err)
			}
		} else {
			err := tx.Omit(clause.Associations).Create(b).Error
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

	if err := tx.Error; err != nil {
		return err
	}

	for _, b := range bs {
		if b.PublishedOn.IsZero() {
			err := tx.Omit(clause.Associations, "published_on").Save(b).Error
			if err != nil {
				tx.Rollback()
				return exception.ErrorInDatastore.New(err)
			}
		} else {
			err := tx.Omit(clause.Associations).Save(b).Error
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
	}

	return tx.Commit().Error
}

func (r *bookRepository) Delete(ctx context.Context, bookID int) error {
	err := r.client.db.Where("id = ?", bookID).Delete(&book.Book{}).Error
	if err != nil {
		return exception.ErrorInDatastore.New(err)
	}

	return nil
}

func (r *bookRepository) GetIDByIsbn(ctx context.Context, isbn string) (int, error) {
	b := &book.Book{}

	err := r.client.db.Select("id").First(b, "isbn = ?", isbn).Error
	if err != nil {
		return 0, exception.NotFound.New(err)
	}

	return b.ID, nil
}

func (r *bookRepository) GetAuthorIDByName(ctx context.Context, name string) (int, error) {
	a := &book.Author{}

	err := r.client.db.Select("id").First(a, "name = ?", name).Error
	if err != nil {
		return 0, exception.NotFound.New(err)
	}

	return a.ID, nil
}

func (r *bookRepository) GetBookshelfIDByUserIDAndBookID(
	ctx context.Context, userID string, bookID int,
) (int, error) {
	b := &book.Bookshelf{}

	err := r.client.db.Select("id").First(b, "user_id = ? AND book_id = ?", userID, bookID).Error
	if err != nil {
		return 0, exception.NotFound.New(err)
	}

	return b.ID, nil
}

func associate(tx *gorm.DB, b *book.Book) error {
	err := associateAuthor(tx, b)
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
			CreatedAt: b.UpdatedAt,
			UpdatedAt: b.UpdatedAt,
		}

		err := tx.Table("authors_books").Create(&ba).Error
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
