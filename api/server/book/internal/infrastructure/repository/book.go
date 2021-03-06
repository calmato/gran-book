package repository

import (
	"context"

	"github.com/calmato/gran-book/api/server/book/internal/domain"
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

func (r *bookRepository) List(ctx context.Context, q *domain.ListQuery) ([]*book.Book, error) {
	bs := []*book.Book{}

	sql := r.client.db.Preload("Authors")
	db := r.client.getListQuery(sql, q)

	err := db.Find(&bs).Error
	if err != nil {
		return nil, exception.ErrorInDatastore.New(err)
	}

	return bs, nil
}

func (r *bookRepository) ListBookshelf(ctx context.Context, q *domain.ListQuery) ([]*book.Bookshelf, error) {
	bss := []*book.Bookshelf{}

	sql := r.client.db.Preload("Book").Preload("Book.Authors")
	db := r.client.getListQuery(sql, q)

	err := db.Find(&bss).Error
	if err != nil {
		return nil, exception.ErrorInDatastore.New(err)
	}

	return bss, nil
}

func (r *bookRepository) ListReview(ctx context.Context, q *domain.ListQuery) ([]*book.Review, error) {
	rvs := []*book.Review{}

	sql := r.client.db
	db := r.client.getListQuery(sql, q)

	err := db.Find(&rvs).Error
	if err != nil {
		return nil, exception.ErrorInDatastore.New(err)
	}

	return rvs, nil
}

func (r *bookRepository) ListCount(ctx context.Context, q *domain.ListQuery) (int, error) {
	sql := r.client.db.Table("books")

	total, err := r.client.getListCount(sql, q)
	if err != nil {
		return 0, exception.ErrorInDatastore.New(err)
	}

	return total, nil
}

func (r *bookRepository) ListBookshelfCount(ctx context.Context, q *domain.ListQuery) (int, error) {
	sql := r.client.db.Table("bookshelves")

	total, err := r.client.getListCount(sql, q)
	if err != nil {
		return 0, exception.ErrorInDatastore.New(err)
	}

	return total, nil
}

func (r *bookRepository) ListReviewCount(ctx context.Context, q *domain.ListQuery) (int, error) {
	sql := r.client.db.Table("reviews")

	total, err := r.client.getListCount(sql, q)
	if err != nil {
		return 0, exception.ErrorInDatastore.New(err)
	}

	return total, nil
}

func (r *bookRepository) Show(ctx context.Context, bookID int) (*book.Book, error) {
	b := &book.Book{}

	err := r.client.db.Preload("Authors").First(b, "id = ?", bookID).Error
	if err != nil {
		return nil, exception.NotFound.New(err)
	}

	return b, nil
}

func (r *bookRepository) ShowByIsbn(ctx context.Context, isbn string) (*book.Book, error) {
	b := &book.Book{}

	err := r.client.db.Preload("Authors").First(b, "isbn = ?", isbn).Error
	if err != nil {
		return nil, exception.NotFound.New(err)
	}

	return b, nil
}

func (r *bookRepository) ShowBookshelfByUserIDAndBookID(
	ctx context.Context, userID string, bookID int,
) (*book.Bookshelf, error) {
	b := &book.Bookshelf{}

	err := r.client.db.
		Preload("Book").Preload("Book.Authors").
		First(b, "user_id = ? AND book_id = ?", userID, bookID).Error
	if err != nil {
		return nil, exception.NotFound.New(err)
	}

	return b, nil
}

func (r *bookRepository) ShowReview(ctx context.Context, reviewID int) (*book.Review, error) {
	rv := &book.Review{}

	err := r.client.db.First(rv, "id = ?", reviewID).Error
	if err != nil {
		return nil, exception.NotFound.New(err)
	}

	return rv, nil
}

func (r *bookRepository) ShowReviewByUserIDAndBookID(
	ctx context.Context, userID string, bookID int,
) (*book.Review, error) {
	rv := &book.Review{}

	err := r.client.db.First(rv, "user_id = ? AND book_id = ?", userID, bookID).Error
	if err != nil {
		return nil, exception.NotFound.New(err)
	}

	return rv, nil
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

	err := tx.Omit(clause.Associations).Create(&b).Error
	if err != nil {
		tx.Rollback()
		return exception.ErrorInDatastore.New(err)
	}

	err = associateBook(tx, b)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (r *bookRepository) CreateBookshelf(ctx context.Context, bs *book.Bookshelf) error {
	tx := r.client.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	if bs.ReadOn.IsZero() {
		err := tx.Omit(clause.Associations, "read_on").Create(&bs).Error
		if err != nil {
			tx.Rollback()
			return exception.ErrorInDatastore.New(err)
		}
	} else {
		err := tx.Omit(clause.Associations).Create(&bs).Error
		if err != nil {
			tx.Rollback()
			return exception.ErrorInDatastore.New(err)
		}
	}

	err := associateBookshelf(tx, bs)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (r *bookRepository) CreateReview(ctx context.Context, rv *book.Review) error {
	err := r.client.db.Create(&rv).Error
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

	if err := tx.Error; err != nil {
		return err
	}

	err := tx.Omit(clause.Associations).Save(&b).Error
	if err != nil {
		tx.Rollback()
		return exception.ErrorInDatastore.New(err)
	}

	err = associateBook(tx, b)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (r *bookRepository) UpdateBookshelf(ctx context.Context, bs *book.Bookshelf) error {
	tx := r.client.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	if bs.ReadOn.IsZero() {
		err := r.client.db.Omit(clause.Associations, "read_on").Save(&bs).Error
		if err != nil {
			return exception.ErrorInDatastore.New(err)
		}
	} else {
		err := r.client.db.Omit(clause.Associations).Save(&bs).Error
		if err != nil {
			return exception.ErrorInDatastore.New(err)
		}
	}

	err := associateBookshelf(tx, bs)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (r *bookRepository) UpdateReview(ctx context.Context, rv *book.Review) error {
	err := r.client.db.Save(&rv).Error
	if err != nil {
		return exception.ErrorInDatastore.New(err)
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
		err := tx.Omit(clause.Associations).Create(b).Error
		if err != nil {
			tx.Rollback()
			return exception.ErrorInDatastore.New(err)
		}

		err = associateBook(tx, b)
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
		err := tx.Omit(clause.Associations).Save(b).Error
		if err != nil {
			tx.Rollback()
			return exception.ErrorInDatastore.New(err)
		}

		err = associateBook(tx, b)
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

func (r *bookRepository) DeleteBookshelf(ctx context.Context, bookshelfID int) error {
	err := r.client.db.Where("id = ?", bookshelfID).Delete(&book.Bookshelf{}).Error
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

func (r *bookRepository) GetReviewIDByUserIDAndBookID(
	ctx context.Context, userID string, bookID int,
) (int, error) {
	rv := &book.Review{}

	err := r.client.db.Select("id").First(rv, "user_id = ? AND book_id = ?", userID, bookID).Error
	if err != nil {
		return 0, exception.NotFound.New(err)
	}

	return rv.ID, nil
}

func associateBook(tx *gorm.DB, b *book.Book) error {
	err := associateAuthor(tx, b)
	if err != nil {
		return err
	}

	return nil
}

func associateBookshelf(tx *gorm.DB, bs *book.Bookshelf) error {
	// 現状の実装内容として、読んだ本のステータスの時のみレビューをすることになってるため
	if bs.Status != book.ReadStatus {
		return nil
	}

	err := associateReview(tx, bs)
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

func associateReview(tx *gorm.DB, bs *book.Bookshelf) error {
	if bs.Review == nil {
		return nil
	}

	if bs.Review.ID == 0 {
		bs.Review.CreatedAt = bs.UpdatedAt
		bs.Review.UpdatedAt = bs.UpdatedAt

		err := tx.Create(&bs.Review).Error
		if err != nil {
			return exception.ErrorInDatastore.New(err)
		}
	} else {
		bs.Review.UpdatedAt = bs.UpdatedAt

		err := tx.Save(&bs.Review).Error
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
