package repository

import (
	"context"
	"errors"
	"time"

	"github.com/calmato/gran-book/api/internal/book/domain/book"
	"github.com/calmato/gran-book/api/pkg/array"
	"github.com/calmato/gran-book/api/pkg/database"
	"github.com/calmato/gran-book/api/pkg/exception"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type bookRepository struct {
	client *database.Client
	now    func() time.Time
}

// NewBookRepository - BookRepositoryの生成
func NewBookRepository(c *database.Client, now func() time.Time) book.Repository {
	return &bookRepository{
		client: c,
		now:    now,
	}
}

func (r *bookRepository) List(ctx context.Context, q *database.ListQuery) (book.Books, error) {
	bs := book.Books{}

	sql := r.client.DB.Preload("Authors")
	err := r.client.GetListQuery(bookTable, sql, q).Find(&bs).Error
	return bs, exception.ToDBError(err)
}

func (r *bookRepository) ListBookshelf(ctx context.Context, q *database.ListQuery) (book.Bookshelves, error) {
	bss := book.Bookshelves{}

	sql := r.client.DB.Preload("Book").Preload("Book.Authors")
	err := r.client.GetListQuery(bookshelfTable, sql, q).Find(&bss).Error
	return bss, exception.ToDBError(err)
}

func (r *bookRepository) ListReview(ctx context.Context, q *database.ListQuery) (book.Reviews, error) {
	rvs := book.Reviews{}

	err := r.client.GetListQuery(reviewTable, r.client.DB, q).Find(&rvs).Error
	return rvs, exception.ToDBError(err)
}

func (r *bookRepository) Count(ctx context.Context, q *database.ListQuery) (int, error) {
	total, err := r.client.GetListCount(bookTable, r.client.DB, q)
	return total, exception.ToDBError(err)
}

func (r *bookRepository) CountBookshelf(ctx context.Context, q *database.ListQuery) (int, error) {
	total, err := r.client.GetListCount(bookshelfTable, r.client.DB, q)
	return total, exception.ToDBError(err)
}

func (r *bookRepository) CountReview(ctx context.Context, q *database.ListQuery) (int, error) {
	total, err := r.client.GetListCount(reviewTable, r.client.DB, q)
	return total, exception.ToDBError(err)
}

func (r *bookRepository) MultiGet(ctx context.Context, bookIDs []int) (book.Books, error) {
	bs := book.Books{}

	err := r.client.DB.Table(bookTable).Preload("Authors").Where("id IN (?)", bookIDs).Find(&bs).Error
	return bs, exception.ToDBError(err)
}

func (r *bookRepository) Get(ctx context.Context, bookID int) (*book.Book, error) {
	b := &book.Book{}

	err := r.client.DB.Table(bookTable).Preload("Authors").First(b, "id = ?", bookID).Error

	return b, exception.ToDBError(err)
}

func (r *bookRepository) GetByIsbn(ctx context.Context, isbn string) (*book.Book, error) {
	b := &book.Book{}

	err := r.client.DB.Table(bookTable).Preload("Authors").First(b, "isbn = ?", isbn).Error
	return b, exception.ToDBError(err)
}

func (r *bookRepository) GetBookIDByIsbn(ctx context.Context, isbn string) (int, error) {
	b := &book.Book{}

	err := r.client.DB.Table(bookTable).Select("id").First(b, "isbn = ?", isbn).Error
	return b.ID, exception.ToDBError(err)
}

func (r *bookRepository) GetBookshelfByUserIDAndBookID(
	ctx context.Context, userID string, bookID int,
) (*book.Bookshelf, error) {
	b := &book.Bookshelf{}

	err := r.client.DB.
		Table(bookshelfTable).
		Preload("Book").
		Preload("Book.Authors").
		First(b, "user_id = ? AND book_id = ?", userID, bookID).Error
	return b, exception.ToDBError(err)
}

func (r *bookRepository) GetBookshelfIDByUserIDAndBookID(
	ctx context.Context, userID string, bookID int,
) (int, error) {
	b := &book.Bookshelf{}

	err := r.client.DB.
		Table(bookshelfTable).
		Select("id").
		First(b, "user_id = ? AND book_id = ?", userID, bookID).Error
	return b.ID, exception.ToDBError(err)
}

func (r *bookRepository) GetReview(ctx context.Context, reviewID int) (*book.Review, error) {
	rv := &book.Review{}

	err := r.client.DB.Table(reviewTable).First(rv, "id = ?", reviewID).Error
	return rv, exception.ToDBError(err)
}

func (r *bookRepository) GetReviewByUserIDAndBookID(
	ctx context.Context, userID string, bookID int,
) (*book.Review, error) {
	rv := &book.Review{}

	err := r.client.DB.
		Table(reviewTable).
		First(rv, "user_id = ? AND book_id = ?", userID, bookID).Error
	return rv, exception.ToDBError(err)
}

func (r *bookRepository) GetReviewIDByUserIDAndBookID(
	ctx context.Context, userID string, bookID int,
) (int, error) {
	rv := &book.Review{}

	err := r.client.DB.
		Table(reviewTable).
		Select("id").
		First(rv, "user_id = ? AND book_id = ?", userID, bookID).Error
	return rv.ID, exception.ToDBError(err)
}

func (r *bookRepository) GetAuthorByName(ctx context.Context, name string) (*book.Author, error) {
	a := &book.Author{}

	err := r.client.DB.Table(authorTable).Where("name = ?", name).First(a).Error
	return a, exception.ToDBError(err)
}

func (r *bookRepository) GetAuthorIDByName(ctx context.Context, name string) (int, error) {
	a := &book.Author{}

	err := r.client.DB.Table(authorTable).Select("id").First(a, "name = ?", name).Error
	return a.ID, exception.ToDBError(err)
}

func (r *bookRepository) Create(ctx context.Context, b *book.Book) error {
	tx, err := r.client.Begin()
	if err != nil {
		return exception.ToDBError(err)
	}
	defer r.client.Close(tx)

	err = r.createBook(tx, b)
	if err != nil {
		tx.Rollback()
		return exception.ToDBError(err)
	}

	err = r.createBookAssociation(tx, b)
	if err != nil {
		tx.Rollback()
		return exception.ToDBError(err)
	}

	return exception.ToDBError(tx.Commit().Error)
}

func (r *bookRepository) MultipleCreate(ctx context.Context, bs book.Books) error {
	tx, err := r.client.Begin()
	if err != nil {
		return exception.ToDBError(err)
	}
	defer r.client.Close(tx)

	err = r.createBook(tx, bs...)
	if err != nil {
		tx.Rollback()
		return exception.ToDBError(err)
	}

	err = r.createBookAssociation(tx, bs...)
	if err != nil {
		tx.Rollback()
		return exception.ToDBError(err)
	}

	return exception.ToDBError(tx.Commit().Error)
}

func (r *bookRepository) createBook(tx *gorm.DB, bs ...*book.Book) error {
	now := r.now()
	for _, b := range bs {
		b.CreatedAt = now
		b.UpdatedAt = now
	}

	return tx.Table(bookTable).Omit(clause.Associations).Create(&bs).Error
}

func (r *bookRepository) createBookAssociation(tx *gorm.DB, bs ...*book.Book) error {
	bas := make([]*book.BookAuthor, 0)
	for _, b := range bs {
		// 著者情報の取得 or 新規登録
		err := r.getOrCreateAuthor(tx, b.Authors...)
		if err != nil {
			return err
		}

		// 書籍と著者の関連付けレコードの作成
		for _, a := range b.Authors {
			ba := &book.BookAuthor{
				BookID:   b.ID,
				AuthorID: a.ID,
			}
			bas = append(bas, ba)
		}
	}

	if len(bas) == 0 {
		return nil
	}

	return r.createAuthorBook(tx, bas...)
}

func (r *bookRepository) CreateBookshelf(ctx context.Context, bs *book.Bookshelf) error {
	tx, err := r.client.Begin()
	if err != nil {
		return exception.ToDBError(err)
	}
	defer r.client.Close(tx)

	err = r.createBookshelfAssociation(tx, bs)
	if err != nil {
		tx.Rollback()
		return exception.ToDBError(err)
	}

	err = r.createBookshelf(tx, bs)
	if err != nil {
		tx.Rollback()
		return exception.ToDBError(err)
	}

	return exception.ToDBError(tx.Commit().Error)
}

func (r *bookRepository) createBookshelf(tx *gorm.DB, bs *book.Bookshelf) error {
	now := r.now()
	bs.CreatedAt = now
	bs.UpdatedAt = now

	sql := tx.Table(bookshelfTable)
	if bs.ReadOn.IsZero() {
		sql = sql.Omit(clause.Associations, "read_on")
	} else {
		sql = sql.Omit(clause.Associations)
	}

	return sql.Create(bs).Error
}

func (r *bookRepository) createBookshelfAssociation(tx *gorm.DB, bs *book.Bookshelf) error {
	// 本棚へ再追加時、以前レビューした情報が残っている場合があるため
	rv := &book.Review{}

	err := tx.Table(reviewTable).First(rv, "user_id = ? AND book_id = ?", bs.UserID, bs.BookID).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	// 以前にレビューをしていた場合、関連付けを追加
	if err == nil {
		bs.ReviewID = rv.ID
	}

	// 現状の実装内容として、読んだ本のステータスの時のみレビューをすることになってるため
	if bs.Status != book.ReadStatus {
		if rv.ID != 0 {
			bs.Review = rv
		}

		return nil
	}

	// 読んだ本のステータスのときは、レビューをするためUpsert/Delete処理
	// Review == nil -> Delete, Review != nil -> Upsert
	if bs.Review == nil || bs.Review.Impression == "" {
		if rv.ID == 0 {
			return nil
		}

		return tx.Table(reviewTable).Delete(rv, "id = ?", rv.ID).Error
	}

	rv.BookID = bs.BookID
	rv.UserID = bs.UserID
	rv.Impression = bs.Review.Impression

	err = r.upsertReview(tx, rv)
	if err != nil {
		return err
	}

	bs.Review = rv
	bs.ReviewID = rv.ID
	return nil
}

func (r *bookRepository) getOrCreateAuthor(tx *gorm.DB, as ...*book.Author) error {
	now := r.now()
	for _, a := range as {
		a.CreatedAt = now
		a.UpdatedAt = now

		err := tx.Table(authorTable).Where("name = ? AND name_kana = ?", a.Name, a.NameKana).FirstOrCreate(a).Error
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *bookRepository) createAuthorBook(tx *gorm.DB, bas ...*book.BookAuthor) error {
	now := r.now()
	for _, ba := range bas {
		ba.CreatedAt = now
		ba.UpdatedAt = now
	}

	return tx.Table(authorBookTable).Create(&bas).Error
}

func (r *bookRepository) Update(ctx context.Context, b *book.Book) error {
	tx, err := r.client.Begin()
	if err != nil {
		return exception.ToDBError(err)
	}
	defer r.client.Close(tx)

	err = r.updateBook(tx, b)
	if err != nil {
		tx.Rollback()
		return exception.ToDBError(err)
	}

	err = r.updateBookAssociation(tx, b)
	if err != nil {
		tx.Rollback()
		return exception.ToDBError(err)
	}

	return exception.ToDBError(tx.Commit().Error)
}

func (r *bookRepository) MultipleUpdate(ctx context.Context, bs book.Books) error {
	tx, err := r.client.Begin()
	if err != nil {
		return exception.ToDBError(err)
	}
	defer r.client.Close(tx)

	err = r.updateBook(tx, bs...)
	if err != nil {
		tx.Rollback()
		return exception.ToDBError(err)
	}

	err = r.updateBookAssociation(tx, bs...)
	if err != nil {
		tx.Rollback()
		return exception.ToDBError(err)
	}

	return exception.ToDBError(tx.Commit().Error)
}

func (r *bookRepository) updateBook(tx *gorm.DB, bs ...*book.Book) error {
	now := r.now()
	for _, b := range bs {
		b.CreatedAt = now
		b.UpdatedAt = now
	}

	return tx.Table(bookTable).Omit(clause.Associations).Save(&bs).Error
}

func (r *bookRepository) updateBookAssociation(tx *gorm.DB, bs ...*book.Book) error {
	// 著者情報の取得 or 新規登録, 現在のAuthorID一覧の作成
	currentAuthorIDs := make([]int, 0)
	for _, b := range bs {
		err := r.getOrCreateAuthor(tx, b.Authors...)
		if err != nil {
			return err
		}

		authorIDs := make([]int, len(b.Authors))
		for i := range b.Authors {
			authorIDs[i] = b.Authors[i].ID
		}

		currentAuthorIDs = append(currentAuthorIDs, authorIDs...)
	}

	// 既存の関連レコード取得
	for _, b := range bs {
		beforeAuthorIDs := []int{}
		err := tx.Table(authorBookTable).Select("author_id").Where("book_id = ?", b.ID).Find(&beforeAuthorIDs).Error
		if err != nil {
			return err
		}

		// 関連レコードが既存にない場合、新たに関連レコードを作成
		bas := make([]*book.BookAuthor, 0)
		for _, authorID := range currentAuthorIDs {
			if isExists, _ := array.Contains(beforeAuthorIDs, authorID); isExists {
				continue
			}

			ba := &book.BookAuthor{
				BookID:   b.ID,
				AuthorID: authorID,
			}
			bas = append(bas, ba)
		}

		if len(bas) > 0 {
			err = r.createAuthorBook(tx, bas...)
			if err != nil {
				return err
			}
		}

		// 不要な関連レコードを削除
		for _, authorID := range beforeAuthorIDs {
			if isExists, _ := array.Contains(currentAuthorIDs, authorID); isExists {
				continue
			}

			err := r.deleteAuthorBook(tx, b.ID, authorID)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (r *bookRepository) UpdateBookshelf(ctx context.Context, bs *book.Bookshelf) error {
	tx, err := r.client.Begin()
	if err != nil {
		return exception.ToDBError(err)
	}
	defer r.client.Close(tx)

	err = r.updateBookshelfAssociation(tx, bs)
	if err != nil {
		tx.Rollback()
		return exception.ToDBError(err)
	}

	err = r.updateBookshelf(tx, bs)
	if err != nil {
		tx.Rollback()
		return exception.ToDBError(err)
	}

	return exception.ToDBError(tx.Commit().Error)
}

func (r *bookRepository) updateBookshelf(tx *gorm.DB, bs *book.Bookshelf) error {
	now := r.now()
	bs.CreatedAt = now
	bs.UpdatedAt = now

	sql := tx.Table(bookshelfTable)
	if bs.ReadOn.IsZero() {
		sql = sql.Omit(clause.Associations, "read_on")
	} else {
		sql = sql.Omit(clause.Associations)
	}

	return sql.Save(bs).Error
}

func (r *bookRepository) updateBookshelfAssociation(tx *gorm.DB, bs *book.Bookshelf) error {
	// 更新時、以前レビューした情報が残っている場合があるため
	rv := &book.Review{}

	err := tx.Table(reviewTable).First(rv, "user_id = ? AND book_id = ?", bs.UserID, bs.BookID).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	// 以前にレビューをしていた場合、関連付けを追加
	if err == nil {
		bs.ReviewID = rv.ID
	}

	// 現状の実装内容として、読んだ本のステータスの時のみレビューをすることになってるため
	if bs.Status != book.ReadStatus {
		return nil
	}

	// 読んだ本のステータスのときは、レビューをするためUpsert/Delete処理
	// Review == nil -> Delete, Review != nil -> Upsert
	if bs.Review == nil || bs.Review.Impression == "" {
		if rv.ID == 0 {
			return nil
		}

		return tx.Table(reviewTable).Delete(rv, "id = ?", rv.ID).Error
	}

	rv.BookID = bs.BookID
	rv.UserID = bs.UserID
	rv.Impression = bs.Review.Impression

	err = r.upsertReview(tx, rv)
	if err != nil {
		return err
	}

	bs.Review = rv
	bs.ReviewID = rv.ID
	return nil
}

func (r *bookRepository) upsertReview(tx *gorm.DB, rvs ...*book.Review) error {
	now := r.now()

	for _, rv := range rvs {
		rv.UpdatedAt = now

		var err error
		if rv.ID == 0 {
			rv.CreatedAt = now
			err = tx.Table(reviewTable).Create(rv).Error
		} else {
			err = tx.Table(reviewTable).Save(rv).Error
		}

		if err != nil {
			return err
		}
	}

	return nil
}

func (r *bookRepository) Delete(ctx context.Context, bookID int) error {
	err := r.client.DB.Table(bookTable).Where("id = ?", bookID).Delete(&book.Book{}).Error
	return exception.ToDBError(err)
}

func (r *bookRepository) DeleteBookshelf(ctx context.Context, bookshelfID int) error {
	err := r.client.DB.Table(bookshelfTable).Where("id = ?", bookshelfID).Delete(&book.Bookshelf{}).Error
	return exception.ToDBError(err)
}

func (r *bookRepository) deleteAuthorBook(tx *gorm.DB, bookID int, authorID int) error {
	return tx.Table(authorBookTable).
		Where("book_id = ? AND author_id = ?", bookID, authorID).
		Delete(&book.BookAuthor{}).Error
}

func (r *bookRepository) AggregateReadTotal(
	ctx context.Context, userID string, since, until time.Time,
) (book.MonthlyResults, error) {
	rs := book.MonthlyResults{}

	err := r.client.DB.Table(bookshelfTable).
		Select("DATE_FORMAT(read_on, '%Y') AS year, DATE_FORMAT(read_on, '%m') AS month, COUNT(id) AS read_total").
		Where("user_id = ?", userID).
		Where("status = ?", book.ReadStatus).
		Where("read_on IS NOT NULL").
		Where("read_on >= ?", since).
		Where("read_on <= ?", until).
		Group("year, month").
		Order("year DESC, month DESC").
		Find(&rs).Error
	return rs, exception.ToDBError(err)
}
