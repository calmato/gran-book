package repository

import (
	"context"

	"github.com/calmato/gran-book/api/server/book/internal/domain/book"
	"github.com/calmato/gran-book/api/server/book/internal/domain/exception"
	"github.com/calmato/gran-book/api/server/book/pkg/array"
	"github.com/calmato/gran-book/api/server/book/pkg/database"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type bookRepository struct {
	client *database.Client
}

// NewBookRepository - BookRepositoryの生成
func NewBookRepository(c *database.Client) book.Repository {
	return &bookRepository{
		client: c,
	}
}

func (r *bookRepository) List(ctx context.Context, q *database.ListQuery) ([]*book.Book, error) {
	bs := []*book.Book{}

	sql := r.client.DB.Preload("Authors")
	err := r.client.GetListQuery("books", sql, q).Find(&bs).Error
	if err != nil {
		return nil, exception.ErrorInDatastore.New(err)
	}

	return bs, nil
}

func (r *bookRepository) ListBookshelf(ctx context.Context, q *database.ListQuery) ([]*book.Bookshelf, error) {
	bss := []*book.Bookshelf{}

	sql := r.client.DB.Preload("Book").Preload("Book.Authors")
	err := r.client.GetListQuery("bookshelves", sql, q).Find(&bss).Error
	if err != nil {
		return nil, exception.ErrorInDatastore.New(err)
	}

	return bss, nil
}

func (r *bookRepository) ListReview(ctx context.Context, q *database.ListQuery) ([]*book.Review, error) {
	rvs := []*book.Review{}

	err := r.client.GetListQuery("reviews", r.client.DB, q).Find(&rvs).Error
	if err != nil {
		return nil, exception.ErrorInDatastore.New(err)
	}

	return rvs, nil
}

func (r *bookRepository) Count(ctx context.Context, q *database.ListQuery) (int, error) {
	total, err := r.client.GetListCount("books", r.client.DB, q)
	if err != nil {
		return 0, exception.ErrorInDatastore.New(err)
	}

	return total, nil
}

func (r *bookRepository) CountBookshelf(ctx context.Context, q *database.ListQuery) (int, error) {
	total, err := r.client.GetListCount("bookshelves", r.client.DB, q)
	if err != nil {
		return 0, exception.ErrorInDatastore.New(err)
	}

	return total, nil
}

func (r *bookRepository) CountReview(ctx context.Context, q *database.ListQuery) (int, error) {
	total, err := r.client.GetListCount("reviews", r.client.DB, q)
	if err != nil {
		return 0, exception.ErrorInDatastore.New(err)
	}

	return total, nil
}

func (r *bookRepository) MultiGet(ctx context.Context, bookIDs []int) ([]*book.Book, error) {
	bs := []*book.Book{}

	err := r.client.DB.Table("books").Where("id IN (?)", bookIDs).Find(&bs).Error
	if err != nil {
		return nil, exception.NotFound.New(err)
	}

	return bs, nil
}

func (r *bookRepository) Get(ctx context.Context, bookID int) (*book.Book, error) {
	b := &book.Book{}

	err := r.client.DB.Table("books").Preload("Authors").First(b, "id = ?", bookID).Error
	if err != nil {
		return nil, exception.NotFound.New(err)
	}

	return b, nil
}

func (r *bookRepository) GetByIsbn(ctx context.Context, isbn string) (*book.Book, error) {
	b := &book.Book{}

	err := r.client.DB.Table("books").Preload("Authors").First(b, "isbn = ?", isbn).Error
	if err != nil {
		return nil, exception.NotFound.New(err)
	}

	return b, nil
}

func (r *bookRepository) GetBookIDByIsbn(ctx context.Context, isbn string) (int, error) {
	b := &book.Book{}

	err := r.client.DB.Table("books").Select("id").First(b, "isbn = ?", isbn).Error
	if err != nil {
		return 0, exception.NotFound.New(err)
	}

	return b.ID, nil
}

func (r *bookRepository) GetBookshelfByUserIDAndBookID(
	ctx context.Context, userID string, bookID int,
) (*book.Bookshelf, error) {
	b := &book.Bookshelf{}

	err := r.client.DB.
		Table("bookshelves").
		Preload("Book").
		Preload("Book.Authors").
		First(b, "user_id = ? AND book_id = ?", userID, bookID).Error
	if err != nil {
		return nil, exception.NotFound.New(err)
	}

	return b, nil
}

func (r *bookRepository) GetBookshelfIDByUserIDAndBookID(
	ctx context.Context, userID string, bookID int,
) (int, error) {
	b := &book.Bookshelf{}

	err := r.client.DB.
		Table("bookshelves").
		Select("id").
		First(b, "user_id = ? AND book_id = ?", userID, bookID).Error
	if err != nil {
		return 0, exception.NotFound.New(err)
	}

	return b.ID, nil
}

func (r *bookRepository) GetReview(ctx context.Context, reviewID int) (*book.Review, error) {
	rv := &book.Review{}

	err := r.client.DB.Table("reviews").First(rv, "id = ?", reviewID).Error
	if err != nil {
		return nil, exception.NotFound.New(err)
	}

	return rv, nil
}

func (r *bookRepository) GetReviewByUserIDAndBookID(
	ctx context.Context, userID string, bookID int,
) (*book.Review, error) {
	rv := &book.Review{}

	err := r.client.DB.
		Table("reviews").
		First(rv, "user_id = ? AND book_id = ?", userID, bookID).Error
	if err != nil {
		return nil, exception.NotFound.New(err)
	}

	return rv, nil
}

func (r *bookRepository) GetReviewIDByUserIDAndBookID(
	ctx context.Context, userID string, bookID int,
) (int, error) {
	rv := &book.Review{}

	err := r.client.DB.
		Table("reviews").
		Select("id").
		First(rv, "user_id = ? AND book_id = ?", userID, bookID).Error
	if err != nil {
		return 0, exception.NotFound.New(err)
	}

	return rv.ID, nil
}

func (r *bookRepository) GetAuthorByName(ctx context.Context, name string) (*book.Author, error) {
	a := &book.Author{}

	err := r.client.DB.Table("authors").Where("name = ?", name).First(a).Error
	if err != nil {
		return nil, exception.ErrorInDatastore.New(err)
	}

	return a, nil
}

func (r *bookRepository) GetAuthorIDByName(ctx context.Context, name string) (int, error) {
	a := &book.Author{}

	err := r.client.DB.Table("authors").Select("id").First(a, "name = ?", name).Error
	if err != nil {
		return 0, exception.NotFound.New(err)
	}

	return a.ID, nil
}

func (r *bookRepository) Create(ctx context.Context, b *book.Book) error {
	tx := r.client.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	err := tx.Table("books").Omit(clause.Associations).Create(&b).Error
	if err != nil {
		tx.Rollback()
		return exception.ErrorInDatastore.New(err)
	}

	err = r.associateBook(tx, b)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (r *bookRepository) CreateBookshelf(ctx context.Context, bs *book.Bookshelf) error {
	tx := r.client.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	err := r.associateBookshelf(tx, bs)
	if err != nil {
		tx.Rollback()
		return err
	}

	if bs.ReadOn.IsZero() {
		err = tx.Table("bookshelves").Omit(clause.Associations, "read_on").Create(&bs).Error
		if err != nil {
			tx.Rollback()
			return exception.ErrorInDatastore.New(err)
		}
	} else {
		err = tx.Table("bookshelves").Omit(clause.Associations).Create(&bs).Error
		if err != nil {
			tx.Rollback()
			return exception.ErrorInDatastore.New(err)
		}
	}

	return tx.Commit().Error
}

func (r *bookRepository) CreateReview(ctx context.Context, rv *book.Review) error {
	err := r.client.DB.Table("reviews").Create(&rv).Error
	if err != nil {
		return exception.ErrorInDatastore.New(err)
	}

	return nil
}

func (r *bookRepository) CreateAuthor(ctx context.Context, a *book.Author) error {
	err := r.client.DB.Table("authors").Create(a).Error
	if err != nil {
		return exception.ErrorInDatastore.New(err)
	}

	return nil
}

func (r *bookRepository) Update(ctx context.Context, b *book.Book) error {
	tx := r.client.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	err := tx.Table("books").Omit(clause.Associations).Save(&b).Error
	if err != nil {
		tx.Rollback()
		return exception.ErrorInDatastore.New(err)
	}

	err = r.associateBook(tx, b)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (r *bookRepository) UpdateBookshelf(ctx context.Context, bs *book.Bookshelf) error {
	tx := r.client.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	err := r.associateBookshelf(tx, bs)
	if err != nil {
		tx.Rollback()
		return err
	}

	if bs.ReadOn.IsZero() {
		err = r.client.DB.Table("bookshelves").Omit(clause.Associations, "read_on").Save(&bs).Error
		if err != nil {
			return exception.ErrorInDatastore.New(err)
		}
	} else {
		err = r.client.DB.Table("bookshelves").Omit(clause.Associations).Save(&bs).Error
		if err != nil {
			return exception.ErrorInDatastore.New(err)
		}
	}

	return tx.Commit().Error
}

func (r *bookRepository) UpdateReview(ctx context.Context, rv *book.Review) error {
	err := r.client.DB.Save(&rv).Error
	if err != nil {
		return exception.ErrorInDatastore.New(err)
	}

	return nil
}

func (r *bookRepository) MultipleCreate(ctx context.Context, bs []*book.Book) error {
	tx := r.client.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	for _, b := range bs {
		err := tx.Table("books").Omit(clause.Associations).Create(b).Error
		if err != nil {
			tx.Rollback()
			return exception.ErrorInDatastore.New(err)
		}

		err = r.associateBook(tx, b)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit().Error
}

func (r *bookRepository) MultipleUpdate(ctx context.Context, bs []*book.Book) error {
	tx := r.client.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	for _, b := range bs {
		err := tx.Table("books").Omit(clause.Associations).Save(b).Error
		if err != nil {
			tx.Rollback()
			return exception.ErrorInDatastore.New(err)
		}

		err = r.associateBook(tx, b)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit().Error
}

func (r *bookRepository) Delete(ctx context.Context, bookID int) error {
	err := r.client.DB.Table("books").Where("id = ?", bookID).Delete(&book.Book{}).Error
	if err != nil {
		return exception.ErrorInDatastore.New(err)
	}

	return nil
}

func (r *bookRepository) DeleteBookshelf(ctx context.Context, bookshelfID int) error {
	err := r.client.DB.Table("bookshelves").Where("id = ?", bookshelfID).Delete(&book.Bookshelf{}).Error
	if err != nil {
		return exception.ErrorInDatastore.New(err)
	}

	return nil
}

func (r *bookRepository) associateBook(tx *gorm.DB, b *book.Book) error {
	err := r.associateAuthor(tx, b)
	if err != nil {
		return err
	}

	return nil
}

func (r *bookRepository) associateBookshelf(tx *gorm.DB, bs *book.Bookshelf) error {
	// 現状の実装内容として、読んだ本のステータスの時のみレビューをすることになってるため
	if bs.Status != book.ReadStatus && bs.ReviewID == 0 {
		return nil
	}

	err := r.associateReview(tx, bs)
	if err != nil {
		return err
	}

	bs.ReviewID = bs.Review.ID

	return nil
}

func (r *bookRepository) associateAuthor(tx *gorm.DB, b *book.Book) error {
	beforeAuthorIDs := []int{}

	// 既存の関連レコード取得
	err := tx.
		Table("authors_books").
		Select("author_id").
		Where("book_id = ?", b.ID).Scan(&beforeAuthorIDs).Error
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
		if isExists, _ := array.Contains(authorID, currentAuthorIDs); !isExists {
			err := tx.
				Table("authors_books").
				Where("book_id = ? AND author_id = ?", b.ID, authorID).
				Delete(&book.BookAuthor{}).Error
			if err != nil {
				return exception.ErrorInDatastore.New(err)
			}
		}
	}

	// 既存レコードとしてない場合、新たに関連レコードの作成
	for _, a := range b.Authors {
		if isExists, _ := array.Contains(beforeAuthorIDs, a.ID); isExists {
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

func (r *bookRepository) associateReview(tx *gorm.DB, bs *book.Bookshelf) error {
	if bs.Review == nil {
		return nil
	}

	if bs.Review.ID == 0 {
		bs.Review.CreatedAt = bs.UpdatedAt
		bs.Review.UpdatedAt = bs.UpdatedAt

		err := tx.Table("reviews").Create(&bs.Review).Error
		if err != nil {
			return exception.ErrorInDatastore.New(err)
		}
	} else {
		bs.Review.UpdatedAt = bs.UpdatedAt

		err := tx.Table("reviews").Save(&bs.Review).Error
		if err != nil {
			return exception.ErrorInDatastore.New(err)
		}
	}

	return nil
}
