package repository

import (
	"context"
	"testing"
	"time"

	"github.com/calmato/gran-book/api/service/internal/book/domain/book"
	"github.com/calmato/gran-book/api/service/pkg/database"
	"github.com/calmato/gran-book/api/service/pkg/datetime"
	"github.com/calmato/gran-book/api/service/pkg/test"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBookRepository_List(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mocks, err := test.NewDBMock(ctrl)
	require.NoError(t, err)

	err = mocks.DeleteAll()
	require.NoError(t, err)

	books := make([]*book.Book, 3)
	books[0] = testBook(1, "1234567890123")
	books[1] = testBook(2, "2345678901234")
	books[2] = testBook(3, "3456789012345")

	err = mocks.BookDB.DB.Table(bookTable).Create(&books).Error
	require.NoError(t, err)

	type args struct {
		query *database.ListQuery
	}
	type want struct {
		books book.Books
		isErr bool
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "success",
			args: args{
				query: &database.ListQuery{
					Limit:      20,
					Offset:     0,
					Conditions: []*database.ConditionQuery{},
				},
			},
			want: want{
				books: books,
				isErr: false,
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			target := NewBookRepository(mocks.BookDB, test.Now)

			bs, err := target.List(ctx, tt.args.query)
			assert.Equal(t, tt.want.isErr, err != nil, err)
			for i := range tt.want.books {
				assert.Contains(t, bs, tt.want.books[i])
			}
		})
	}
}

func TestBookRepository_ListBookshelf(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mocks, err := test.NewDBMock(ctrl)
	require.NoError(t, err)

	err = mocks.DeleteAll()
	require.NoError(t, err)

	books := make([]*book.Book, 3)
	books[0] = testBook(1, "1234567890123")
	books[1] = testBook(2, "2345678901234")
	books[2] = testBook(3, "3456789012345")
	err = mocks.BookDB.DB.Table(bookTable).Create(&books).Error
	require.NoError(t, err)

	userID := "12345678-1234-1234-1234-123456789012"
	bookshelves := make([]*book.Bookshelf, 3)
	bookshelves[0] = testBookshelf(1, books[0].ID, userID)
	bookshelves[0].Book = books[0]
	bookshelves[1] = testBookshelf(2, books[1].ID, userID)
	bookshelves[1].Book = books[1]
	bookshelves[2] = testBookshelf(3, books[2].ID, userID)
	bookshelves[2].Book = books[2]

	err = mocks.BookDB.DB.Table(bookshelfTable).Create(&bookshelves).Error
	require.NoError(t, err)

	type args struct {
		query *database.ListQuery
	}
	type want struct {
		bookshelves book.Bookshelves
		isErr       bool
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "success",
			args: args{
				query: &database.ListQuery{
					Limit:      20,
					Offset:     0,
					Conditions: []*database.ConditionQuery{},
				},
			},
			want: want{
				bookshelves: bookshelves,
				isErr:       false,
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			target := NewBookRepository(mocks.BookDB, test.Now)

			bs, err := target.ListBookshelf(ctx, tt.args.query)
			assert.Equal(t, tt.want.isErr, err != nil, err)
			for i := range tt.want.bookshelves {
				assert.Contains(t, bs, tt.want.bookshelves[i])
			}
		})
	}
}

func TestBookRepository_ListReview(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mocks, err := test.NewDBMock(ctrl)
	require.NoError(t, err)

	err = mocks.DeleteAll()
	require.NoError(t, err)

	books := make([]*book.Book, 3)
	books[0] = testBook(1, "1234567890123")
	books[1] = testBook(2, "2345678901234")
	books[2] = testBook(3, "3456789012345")
	err = mocks.BookDB.DB.Table(bookTable).Create(&books).Error
	require.NoError(t, err)

	userID := "12345678-1234-1234-1234-123456789012"
	reviews := make([]*book.Review, 3)
	reviews[0] = testReview(1, books[0].ID, userID)
	reviews[1] = testReview(2, books[1].ID, userID)
	reviews[2] = testReview(3, books[2].ID, userID)

	err = mocks.BookDB.DB.Table(reviewTable).Create(&reviews).Error
	require.NoError(t, err)

	type args struct {
		query *database.ListQuery
	}
	type want struct {
		reviews book.Reviews
		isErr   bool
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "success",
			args: args{
				query: &database.ListQuery{
					Limit:      20,
					Offset:     0,
					Conditions: []*database.ConditionQuery{},
				},
			},
			want: want{
				reviews: reviews,
				isErr:   false,
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			target := NewBookRepository(mocks.BookDB, test.Now)

			rs, err := target.ListReview(ctx, tt.args.query)
			assert.Equal(t, tt.want.isErr, err != nil, err)
			for i := range tt.want.reviews {
				assert.Contains(t, rs, tt.want.reviews[i])
			}
		})
	}
}

func TestBookRepository_Count(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mocks, err := test.NewDBMock(ctrl)
	require.NoError(t, err)

	err = mocks.DeleteAll()
	require.NoError(t, err)

	books := make([]*book.Book, 3)
	books[0] = testBook(1, "1234567890123")
	books[1] = testBook(2, "2345678901234")
	books[2] = testBook(3, "3456789012345")

	err = mocks.BookDB.DB.Table(bookTable).Create(&books).Error
	require.NoError(t, err)

	type args struct {
		query *database.ListQuery
	}
	type want struct {
		count int
		isErr bool
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "success",
			args: args{
				query: &database.ListQuery{
					Limit:      20,
					Offset:     0,
					Conditions: []*database.ConditionQuery{},
				},
			},
			want: want{
				count: 3,
				isErr: false,
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			target := NewBookRepository(mocks.BookDB, test.Now)

			count, err := target.Count(ctx, tt.args.query)
			assert.Equal(t, tt.want.isErr, err != nil, err)
			assert.Equal(t, tt.want.count, count)
		})
	}
}

func TestBookRepository_CountBookshelf(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mocks, err := test.NewDBMock(ctrl)
	require.NoError(t, err)

	err = mocks.DeleteAll()
	require.NoError(t, err)

	books := make([]*book.Book, 3)
	books[0] = testBook(1, "1234567890123")
	books[1] = testBook(2, "2345678901234")
	books[2] = testBook(3, "3456789012345")
	err = mocks.BookDB.DB.Table(bookTable).Create(&books).Error
	require.NoError(t, err)

	userID := "12345678-1234-1234-1234-123456789012"
	bookshelves := make([]*book.Bookshelf, 3)
	bookshelves[0] = testBookshelf(1, books[0].ID, userID)
	bookshelves[0].Book = books[0]
	bookshelves[1] = testBookshelf(2, books[1].ID, userID)
	bookshelves[1].Book = books[1]
	bookshelves[2] = testBookshelf(3, books[2].ID, userID)
	bookshelves[2].Book = books[2]

	err = mocks.BookDB.DB.Table(bookshelfTable).Create(&bookshelves).Error
	require.NoError(t, err)

	type args struct {
		query *database.ListQuery
	}
	type want struct {
		count int
		isErr bool
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "success",
			args: args{
				query: &database.ListQuery{
					Limit:      20,
					Offset:     0,
					Conditions: []*database.ConditionQuery{},
				},
			},
			want: want{
				count: 3,
				isErr: false,
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			target := NewBookRepository(mocks.BookDB, test.Now)

			count, err := target.CountBookshelf(ctx, tt.args.query)
			assert.Equal(t, tt.want.isErr, err != nil, err)
			assert.Equal(t, tt.want.count, count)
		})
	}
}

func TestBookRepository_CountReview(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mocks, err := test.NewDBMock(ctrl)
	require.NoError(t, err)

	err = mocks.DeleteAll()
	require.NoError(t, err)

	books := make([]*book.Book, 3)
	books[0] = testBook(1, "1234567890123")
	books[1] = testBook(2, "2345678901234")
	books[2] = testBook(3, "3456789012345")
	err = mocks.BookDB.DB.Table(bookTable).Create(&books).Error
	require.NoError(t, err)

	userID := "12345678-1234-1234-1234-123456789012"
	reviews := make([]*book.Review, 3)
	reviews[0] = testReview(1, books[0].ID, userID)
	reviews[1] = testReview(2, books[1].ID, userID)
	reviews[2] = testReview(3, books[2].ID, userID)

	err = mocks.BookDB.DB.Table(reviewTable).Create(&reviews).Error
	require.NoError(t, err)

	type args struct {
		query *database.ListQuery
	}
	type want struct {
		count int
		isErr bool
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "success",
			args: args{
				query: &database.ListQuery{
					Limit:      20,
					Offset:     0,
					Conditions: []*database.ConditionQuery{},
				},
			},
			want: want{
				count: 3,
				isErr: false,
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			target := NewBookRepository(mocks.BookDB, test.Now)

			count, err := target.CountReview(ctx, tt.args.query)
			assert.Equal(t, tt.want.isErr, err != nil, err)
			assert.Equal(t, tt.want.count, count)
		})
	}
}

func TestBookRepository_MultiGet(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mocks, err := test.NewDBMock(ctrl)
	require.NoError(t, err)

	err = mocks.DeleteAll()
	require.NoError(t, err)

	books := make([]*book.Book, 3)
	books[0] = testBook(1, "1234567890123")
	books[1] = testBook(2, "2345678901234")
	books[2] = testBook(3, "3456789012345")

	err = mocks.BookDB.DB.Table(bookTable).Create(&books).Error
	require.NoError(t, err)

	type args struct {
		bookIDs []int
	}
	type want struct {
		books book.Books
		isErr bool
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "success",
			args: args{
				bookIDs: []int{1, 2, 3},
			},
			want: want{
				books: books,
				isErr: false,
			},
		},
		{
			name: "success: not found",
			args: args{
				bookIDs: []int{},
			},
			want: want{
				books: []*book.Book{},
				isErr: false,
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			target := NewBookRepository(mocks.BookDB, test.Now)

			bs, err := target.MultiGet(ctx, tt.args.bookIDs)
			assert.Equal(t, tt.want.isErr, err != nil, err)
			for i := range tt.want.books {
				assert.Contains(t, bs, tt.want.books[i])
			}
		})
	}
}

func TestBookRepository_Get(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mocks, err := test.NewDBMock(ctrl)
	require.NoError(t, err)

	err = mocks.DeleteAll()
	require.NoError(t, err)

	expectBook := testBook(1, "1234567890123")
	err = mocks.BookDB.DB.Table(bookTable).Create(expectBook).Error
	require.NoError(t, err)

	type args struct {
		bookID int
	}
	type want struct {
		book  *book.Book
		isErr bool
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "success",
			args: args{
				bookID: 1,
			},
			want: want{
				book:  expectBook,
				isErr: false,
			},
		},
		{
			name: "failed: not found",
			args: args{
				bookID: 2,
			},
			want: want{
				book:  &book.Book{},
				isErr: true,
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			target := NewBookRepository(mocks.BookDB, test.Now)

			b, err := target.Get(ctx, tt.args.bookID)
			assert.Equal(t, tt.want.isErr, err != nil, err)
			assert.Equal(t, tt.want.book, b)
		})
	}
}

func TestBookRepository_GetByIsbn(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mocks, err := test.NewDBMock(ctrl)
	require.NoError(t, err)

	err = mocks.DeleteAll()
	require.NoError(t, err)

	expectBook := testBook(1, "1234567890123")
	err = mocks.BookDB.DB.Table(bookTable).Create(expectBook).Error
	require.NoError(t, err)

	type args struct {
		isbn string
	}
	type want struct {
		book  *book.Book
		isErr bool
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "success",
			args: args{
				isbn: "1234567890123",
			},
			want: want{
				book:  expectBook,
				isErr: false,
			},
		},
		{
			name: "failed: not found",
			args: args{
				isbn: "",
			},
			want: want{
				book:  &book.Book{},
				isErr: true,
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			target := NewBookRepository(mocks.BookDB, test.Now)

			b, err := target.GetByIsbn(ctx, tt.args.isbn)
			assert.Equal(t, tt.want.isErr, err != nil, err)
			assert.Equal(t, tt.want.book, b)
		})
	}
}

func TestBookRepository_GetBookIDByIsbn(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mocks, err := test.NewDBMock(ctrl)
	require.NoError(t, err)

	err = mocks.DeleteAll()
	require.NoError(t, err)

	expectBook := testBook(1, "1234567890123")
	err = mocks.BookDB.DB.Table(bookTable).Create(expectBook).Error
	require.NoError(t, err)

	type args struct {
		isbn string
	}
	type want struct {
		bookID int
		isErr  bool
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "success",
			args: args{
				isbn: "1234567890123",
			},
			want: want{
				bookID: 1,
				isErr:  false,
			},
		},
		{
			name: "failed: not found",
			args: args{
				isbn: "",
			},
			want: want{
				bookID: 0,
				isErr:  true,
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			target := NewBookRepository(mocks.BookDB, test.Now)

			bookID, err := target.GetBookIDByIsbn(ctx, tt.args.isbn)
			assert.Equal(t, tt.want.isErr, err != nil, err)
			assert.Equal(t, tt.want.bookID, bookID)
		})
	}
}

func TestBookRepository_GetBookshelfByUserIDAndBookID(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mocks, err := test.NewDBMock(ctrl)
	require.NoError(t, err)

	err = mocks.DeleteAll()
	require.NoError(t, err)

	books := make([]*book.Book, 1)
	books[0] = testBook(1, "1234567890123")
	err = mocks.BookDB.DB.Table(bookTable).Create(&books).Error
	require.NoError(t, err)

	userID := "12345678-1234-1234-1234-123456789012"
	expectBookshelf := testBookshelf(1, books[0].ID, userID)
	expectBookshelf.Book = books[0]

	err = mocks.BookDB.DB.Table(bookshelfTable).Create(expectBookshelf).Error
	require.NoError(t, err)

	type args struct {
		userID string
		bookID int
	}
	type want struct {
		bookshelf *book.Bookshelf
		isErr     bool
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "success",
			args: args{
				userID: userID,
				bookID: 1,
			},
			want: want{
				bookshelf: expectBookshelf,
				isErr:     false,
			},
		},
		{
			name: "failed: not found",
			args: args{
				userID: "",
				bookID: 0,
			},
			want: want{
				bookshelf: &book.Bookshelf{},
				isErr:     true,
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			target := NewBookRepository(mocks.BookDB, test.Now)

			b, err := target.GetBookshelfByUserIDAndBookID(ctx, tt.args.userID, tt.args.bookID)
			assert.Equal(t, tt.want.isErr, err != nil, err)
			assert.Equal(t, tt.want.bookshelf, b)
		})
	}
}

func TestBookRepository_GetBookshelfIDByUserIDAndBookID(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mocks, err := test.NewDBMock(ctrl)
	require.NoError(t, err)

	err = mocks.DeleteAll()
	require.NoError(t, err)

	books := make([]*book.Book, 1)
	books[0] = testBook(1, "1234567890123")
	err = mocks.BookDB.DB.Table(bookTable).Create(&books).Error
	require.NoError(t, err)

	userID := "12345678-1234-1234-1234-123456789012"
	expectBookshelf := testBookshelf(1, books[0].ID, userID)
	expectBookshelf.Book = books[0]

	err = mocks.BookDB.DB.Table(bookshelfTable).Create(expectBookshelf).Error
	require.NoError(t, err)

	type args struct {
		userID string
		bookID int
	}
	type want struct {
		bookshelfID int
		isErr       bool
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "success",
			args: args{
				userID: userID,
				bookID: 1,
			},
			want: want{
				bookshelfID: 1,
				isErr:       false,
			},
		},
		{
			name: "failed: not found",
			args: args{
				userID: "",
				bookID: 0,
			},
			want: want{
				bookshelfID: 0,
				isErr:       true,
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			target := NewBookRepository(mocks.BookDB, test.Now)

			bookshelfID, err := target.GetBookshelfIDByUserIDAndBookID(ctx, tt.args.userID, tt.args.bookID)
			assert.Equal(t, tt.want.isErr, err != nil, err)
			assert.Equal(t, tt.want.bookshelfID, bookshelfID)
		})
	}
}

func TestBookRepository_GetReview(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mocks, err := test.NewDBMock(ctrl)
	require.NoError(t, err)

	err = mocks.DeleteAll()
	require.NoError(t, err)

	books := make([]*book.Book, 1)
	books[0] = testBook(1, "1234567890123")
	err = mocks.BookDB.DB.Table(bookTable).Create(&books).Error
	require.NoError(t, err)

	userID := "12345678-1234-1234-1234-123456789012"
	expectReview := testReview(1, books[0].ID, userID)

	err = mocks.BookDB.DB.Table(reviewTable).Create(expectReview).Error
	require.NoError(t, err)

	type args struct {
		reviewID int
	}
	type want struct {
		review *book.Review
		isErr  bool
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "success",
			args: args{
				reviewID: 1,
			},
			want: want{
				review: expectReview,
				isErr:  false,
			},
		},
		{
			name: "failed: not found",
			args: args{
				reviewID: 0,
			},
			want: want{
				review: &book.Review{},
				isErr:  true,
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			target := NewBookRepository(mocks.BookDB, test.Now)

			r, err := target.GetReview(ctx, tt.args.reviewID)
			assert.Equal(t, tt.want.isErr, err != nil, err)
			assert.Equal(t, tt.want.review, r)
		})
	}
}

func TestBookRepository_GetReviewByUserIDAndBookID(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mocks, err := test.NewDBMock(ctrl)
	require.NoError(t, err)

	err = mocks.DeleteAll()
	require.NoError(t, err)

	books := make([]*book.Book, 1)
	books[0] = testBook(1, "1234567890123")
	err = mocks.BookDB.DB.Table(bookTable).Create(&books).Error
	require.NoError(t, err)

	userID := "12345678-1234-1234-1234-123456789012"
	expectReview := testReview(1, books[0].ID, userID)

	err = mocks.BookDB.DB.Table(reviewTable).Create(expectReview).Error
	require.NoError(t, err)

	type args struct {
		userID string
		bookID int
	}
	type want struct {
		review *book.Review
		isErr  bool
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "success",
			args: args{
				userID: userID,
				bookID: 1,
			},
			want: want{
				review: expectReview,
				isErr:  false,
			},
		},
		{
			name: "failed: not found",
			args: args{
				userID: "",
				bookID: 0,
			},
			want: want{
				review: &book.Review{},
				isErr:  true,
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			target := NewBookRepository(mocks.BookDB, test.Now)

			r, err := target.GetReviewByUserIDAndBookID(ctx, tt.args.userID, tt.args.bookID)
			assert.Equal(t, tt.want.isErr, err != nil, err)
			assert.Equal(t, tt.want.review, r)
		})
	}
}

func TestBookRepository_GetReviewIDByUserIDAndBookID(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mocks, err := test.NewDBMock(ctrl)
	require.NoError(t, err)

	err = mocks.DeleteAll()
	require.NoError(t, err)

	books := make([]*book.Book, 1)
	books[0] = testBook(1, "1234567890123")
	err = mocks.BookDB.DB.Table(bookTable).Create(&books).Error
	require.NoError(t, err)

	userID := "12345678-1234-1234-1234-123456789012"
	expectReview := testReview(1, books[0].ID, userID)

	err = mocks.BookDB.DB.Table(reviewTable).Create(expectReview).Error
	require.NoError(t, err)

	type args struct {
		userID string
		bookID int
	}
	type want struct {
		reviewID int
		isErr    bool
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "success",
			args: args{
				userID: userID,
				bookID: 1,
			},
			want: want{
				reviewID: 1,
				isErr:    false,
			},
		},
		{
			name: "failed: not found",
			args: args{
				userID: "",
				bookID: 0,
			},
			want: want{
				reviewID: 0,
				isErr:    true,
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			target := NewBookRepository(mocks.BookDB, test.Now)

			reviewID, err := target.GetReviewIDByUserIDAndBookID(ctx, tt.args.userID, tt.args.bookID)
			assert.Equal(t, tt.want.isErr, err != nil, err)
			assert.Equal(t, tt.want.reviewID, reviewID)
		})
	}
}

func TestBookRepository_GetAuthorByName(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mocks, err := test.NewDBMock(ctrl)
	require.NoError(t, err)

	err = mocks.DeleteAll()
	require.NoError(t, err)

	expectAuthor := testAuthor(1, "有沢 ゆう希", "アリサワ ユウキ")
	err = mocks.BookDB.DB.Table(authorTable).Create(expectAuthor).Error
	require.NoError(t, err)

	type args struct {
		name string
	}
	type want struct {
		author *book.Author
		isErr  bool
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "success",
			args: args{
				name: "有沢 ゆう希",
			},
			want: want{
				author: expectAuthor,
				isErr:  false,
			},
		},
		{
			name: "failed: not found",
			args: args{
				name: "",
			},
			want: want{
				author: &book.Author{},
				isErr:  true,
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			target := NewBookRepository(mocks.BookDB, test.Now)

			a, err := target.GetAuthorByName(ctx, tt.args.name)
			assert.Equal(t, tt.want.isErr, err != nil, err)
			assert.Equal(t, tt.want.author, a)
		})
	}
}

func TestBookRepository_GetAuthorIDByName(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mocks, err := test.NewDBMock(ctrl)
	require.NoError(t, err)

	err = mocks.DeleteAll()
	require.NoError(t, err)

	expectAuthor := testAuthor(1, "有沢 ゆう希", "アリサワ ユウキ")
	err = mocks.BookDB.DB.Table(authorTable).Create(expectAuthor).Error
	require.NoError(t, err)

	type args struct {
		name string
	}
	type want struct {
		authorID int
		isErr    bool
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "success",
			args: args{
				name: "有沢 ゆう希",
			},
			want: want{
				authorID: 1,
				isErr:    false,
			},
		},
		{
			name: "failed: not found",
			args: args{
				name: "",
			},
			want: want{
				authorID: 0,
				isErr:    true,
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			target := NewBookRepository(mocks.BookDB, test.Now)

			authorID, err := target.GetAuthorIDByName(ctx, tt.args.name)
			assert.Equal(t, tt.want.isErr, err != nil, err)
			assert.Equal(t, tt.want.authorID, authorID)
		})
	}
}

func TestBookRepository_Create(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mocks, err := test.NewDBMock(ctrl)
	require.NoError(t, err)

	err = mocks.DeleteAll()
	require.NoError(t, err)

	type args struct {
		book *book.Book
	}
	type want struct {
		isErr bool
	}
	tests := []struct {
		name  string
		setup func(t *testing.T, mocks *test.DBMocks)
		args  args
		want  want
	}{
		{
			name:  "success",
			setup: func(t *testing.T, mocks *test.DBMocks) {},
			args: args{
				book: &book.Book{
					Title:          "小説　ちはやふる　上の句",
					TitleKana:      "ショウセツ チハヤフルカミノク",
					Description:    "綾瀬千早は高校入学と同時に、競技かるた部を作ろうと奔走する。幼馴染の太一と仲間を集め、夏の全国大会に出場するためだ。強くなって、新と再会したい。幼い頃かるたを取り合った、新に寄せる千早の秘めた想いに気づきながらも、太一は千早を守り立てる。それぞれの青春を懸けた、一途な情熱の物語が幕開ける。",
					Isbn:           "1234567890123",
					Publisher:      "講談社",
					PublishedOn:    "2018-01-16",
					ThumbnailURL:   "https://thumbnail.image.rakuten.co.jp/@0_mall/book/cabinet/8426/9784062938426.jpg?_ex=120x120",
					RakutenURL:     "https://books.rakuten.co.jp/rb/15271426/",
					RakutenSize:    "コミック",
					RakutenGenreID: "001004008001/001004008003/001019001",
				},
			},
			want: want{
				isErr: false,
			},
		},
		{
			name: "success with association",
			setup: func(t *testing.T, mocks *test.DBMocks) {
				authors := make(book.Authors, 2)
				authors[0] = testAuthor(1, "有沢 ゆう希", "アリサワ ユウキ")
				authors[1] = testAuthor(2, "末次 由紀", "スエツグ ユキ")
				err := mocks.BookDB.DB.Table(authorTable).Create(&authors).Error
				assert.NoError(t, err)
			},
			args: args{
				book: &book.Book{
					Title:          "小説　ちはやふる　上の句",
					TitleKana:      "ショウセツ チハヤフルカミノク",
					Description:    "綾瀬千早は高校入学と同時に、競技かるた部を作ろうと奔走する。幼馴染の太一と仲間を集め、夏の全国大会に出場するためだ。強くなって、新と再会したい。幼い頃かるたを取り合った、新に寄せる千早の秘めた想いに気づきながらも、太一は千早を守り立てる。それぞれの青春を懸けた、一途な情熱の物語が幕開ける。",
					Isbn:           "1234567890123",
					Publisher:      "講談社",
					PublishedOn:    "2018-01-16",
					ThumbnailURL:   "https://thumbnail.image.rakuten.co.jp/@0_mall/book/cabinet/8426/9784062938426.jpg?_ex=120x120",
					RakutenURL:     "https://books.rakuten.co.jp/rb/15271426/",
					RakutenSize:    "コミック",
					RakutenGenreID: "001004008001/001004008003/001019001",
					Authors: []*book.Author{
						{
							Name:     "テスト 著者",
							NameKana: "てすと ちょしゃ",
						},
						{
							Name:     "有沢 ゆう希",
							NameKana: "アリサワ ユウキ",
						},
					},
				},
			},
			want: want{
				isErr: false,
			},
		},
		{
			name:  "failed: internal error",
			setup: func(t *testing.T, mocks *test.DBMocks) {},
			args: args{
				book: &book.Book{},
			},
			want: want{
				isErr: true,
			},
		},
		{
			name:  "failed: internal error with association",
			setup: func(t *testing.T, mocks *test.DBMocks) {},
			args: args{
				book: &book.Book{
					Title:          "小説　ちはやふる　上の句",
					TitleKana:      "ショウセツ チハヤフルカミノク",
					Description:    "綾瀬千早は高校入学と同時に、競技かるた部を作ろうと奔走する。幼馴染の太一と仲間を集め、夏の全国大会に出場するためだ。強くなって、新と再会したい。幼い頃かるたを取り合った、新に寄せる千早の秘めた想いに気づきながらも、太一は千早を守り立てる。それぞれの青春を懸けた、一途な情熱の物語が幕開ける。",
					Isbn:           "1234567890123",
					Publisher:      "講談社",
					PublishedOn:    "2018-01-16",
					ThumbnailURL:   "https://thumbnail.image.rakuten.co.jp/@0_mall/book/cabinet/8426/9784062938426.jpg?_ex=120x120",
					RakutenURL:     "https://books.rakuten.co.jp/rb/15271426/",
					RakutenSize:    "コミック",
					RakutenGenreID: "001004008001/001004008003/001019001",
					Authors: []*book.Author{
						{
							Name:     "",
							NameKana: "",
						},
					},
				},
			},
			want: want{
				isErr: true,
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			err := mocks.Delete(mocks.BookDB, bookTable, authorTable, authorBookTable)
			require.NoError(t, err)
			tt.setup(t, mocks)

			target := NewBookRepository(mocks.BookDB, test.Now)
			err = target.Create(ctx, tt.args.book)
			assert.Equal(t, tt.want.isErr, err != nil, err)
		})
	}
}

func TestBookRepository_MultipleCreate(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mocks, err := test.NewDBMock(ctrl)
	require.NoError(t, err)

	err = mocks.DeleteAll()
	require.NoError(t, err)

	type args struct {
		books book.Books
	}
	type want struct {
		isErr bool
	}
	tests := []struct {
		name  string
		setup func(t *testing.T, mocks *test.DBMocks)
		args  args
		want  want
	}{
		{
			name:  "success",
			setup: func(t *testing.T, mocks *test.DBMocks) {},
			args: args{
				books: []*book.Book{
					{
						Title:          "小説　ちはやふる　上の句",
						TitleKana:      "ショウセツ チハヤフルカミノク",
						Description:    "綾瀬千早は高校入学と同時に、競技かるた部を作ろうと奔走する。幼馴染の太一と仲間を集め、夏の全国大会に出場するためだ。強くなって、新と再会したい。幼い頃かるたを取り合った、新に寄せる千早の秘めた想いに気づきながらも、太一は千早を守り立てる。それぞれの青春を懸けた、一途な情熱の物語が幕開ける。",
						Isbn:           "1234567890123",
						Publisher:      "講談社",
						PublishedOn:    "2018-01-16",
						ThumbnailURL:   "https://thumbnail.image.rakuten.co.jp/@0_mall/book/cabinet/8426/9784062938426.jpg?_ex=120x120",
						RakutenURL:     "https://books.rakuten.co.jp/rb/15271426/",
						RakutenSize:    "コミック",
						RakutenGenreID: "001004008001/001004008003/001019001",
					},
					{
						Title:          "小説　ちはやふる　上の句",
						TitleKana:      "ショウセツ チハヤフルカミノク",
						Description:    "綾瀬千早は高校入学と同時に、競技かるた部を作ろうと奔走する。幼馴染の太一と仲間を集め、夏の全国大会に出場するためだ。強くなって、新と再会したい。幼い頃かるたを取り合った、新に寄せる千早の秘めた想いに気づきながらも、太一は千早を守り立てる。それぞれの青春を懸けた、一途な情熱の物語が幕開ける。",
						Isbn:           "2345678901234",
						Publisher:      "講談社",
						PublishedOn:    "2018-01-16",
						ThumbnailURL:   "https://thumbnail.image.rakuten.co.jp/@0_mall/book/cabinet/8426/9784062938426.jpg?_ex=120x120",
						RakutenURL:     "https://books.rakuten.co.jp/rb/15271426/",
						RakutenSize:    "コミック",
						RakutenGenreID: "001004008001/001004008003/001019001",
					},
				},
			},
			want: want{
				isErr: false,
			},
		},
		{
			name:  "failed: internal error",
			setup: func(t *testing.T, mocks *test.DBMocks) {},
			args: args{
				books: []*book.Book{{}},
			},
			want: want{
				isErr: true,
			},
		},
		{
			name:  "failed: length 0",
			setup: func(t *testing.T, mocks *test.DBMocks) {},
			args: args{
				books: []*book.Book{},
			},
			want: want{
				isErr: true,
			},
		},
		{
			name:  "failed: internal error with association",
			setup: func(t *testing.T, mocks *test.DBMocks) {},
			args: args{
				books: []*book.Book{
					{
						Title:          "小説　ちはやふる　上の句",
						TitleKana:      "ショウセツ チハヤフルカミノク",
						Description:    "綾瀬千早は高校入学と同時に、競技かるた部を作ろうと奔走する。幼馴染の太一と仲間を集め、夏の全国大会に出場するためだ。強くなって、新と再会したい。幼い頃かるたを取り合った、新に寄せる千早の秘めた想いに気づきながらも、太一は千早を守り立てる。それぞれの青春を懸けた、一途な情熱の物語が幕開ける。",
						Isbn:           "1234567890123",
						Publisher:      "講談社",
						PublishedOn:    "2018-01-16",
						ThumbnailURL:   "https://thumbnail.image.rakuten.co.jp/@0_mall/book/cabinet/8426/9784062938426.jpg?_ex=120x120",
						RakutenURL:     "https://books.rakuten.co.jp/rb/15271426/",
						RakutenSize:    "コミック",
						RakutenGenreID: "001004008001/001004008003/001019001",
					},
					{
						Title:          "小説　ちはやふる　上の句",
						TitleKana:      "ショウセツ チハヤフルカミノク",
						Description:    "綾瀬千早は高校入学と同時に、競技かるた部を作ろうと奔走する。幼馴染の太一と仲間を集め、夏の全国大会に出場するためだ。強くなって、新と再会したい。幼い頃かるたを取り合った、新に寄せる千早の秘めた想いに気づきながらも、太一は千早を守り立てる。それぞれの青春を懸けた、一途な情熱の物語が幕開ける。",
						Isbn:           "2345678901234",
						Publisher:      "講談社",
						PublishedOn:    "2018-01-16",
						ThumbnailURL:   "https://thumbnail.image.rakuten.co.jp/@0_mall/book/cabinet/8426/9784062938426.jpg?_ex=120x120",
						RakutenURL:     "https://books.rakuten.co.jp/rb/15271426/",
						RakutenSize:    "コミック",
						RakutenGenreID: "001004008001/001004008003/001019001",
						Authors: []*book.Author{
							{
								Name:     "",
								NameKana: "",
							},
						},
					},
				},
			},
			want: want{
				isErr: true,
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			err := mocks.Delete(mocks.BookDB, bookTable, authorTable, authorBookTable)
			require.NoError(t, err)
			tt.setup(t, mocks)

			target := NewBookRepository(mocks.BookDB, test.Now)
			err = target.MultipleCreate(ctx, tt.args.books)
			assert.Equal(t, tt.want.isErr, err != nil, err)
		})
	}
}

func TestBookRepository_CreateBookshelf(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mocks, err := test.NewDBMock(ctrl)
	require.NoError(t, err)

	err = mocks.DeleteAll()
	require.NoError(t, err)

	userID := "12345678-1234-1234-1234-123456789012"

	type args struct {
		bookshelf *book.Bookshelf
	}
	type want struct {
		isErr bool
	}
	tests := []struct {
		name  string
		setup func(t *testing.T, mocks *test.DBMocks)
		args  args
		want  want
	}{
		{
			name: "success",
			setup: func(t *testing.T, mocks *test.DBMocks) {
				b := testBook(1, "1234567890123")
				err = mocks.BookDB.DB.Table(bookTable).Create(b).Error
				require.NoError(t, err)
			},
			args: args{
				bookshelf: &book.Bookshelf{
					BookID: 1,
					UserID: userID,
					Status: book.ReadStatus,
				},
			},
			want: want{
				isErr: false,
			},
		},
		{
			name: "success to recreate",
			setup: func(t *testing.T, mocks *test.DBMocks) {
				b := testBook(1, "1234567890123")
				err = mocks.BookDB.DB.Table(bookTable).Create(b).Error
				require.NoError(t, err)
				rv := testReview(1, b.ID, userID)
				err := mocks.BookDB.DB.Table(reviewTable).Create(&rv).Error
				assert.NoError(t, err)
			},
			args: args{
				bookshelf: &book.Bookshelf{
					BookID: 1,
					UserID: userID,
					Status: book.ReadStatus,
				},
			},
			want: want{
				isErr: false,
			},
		},
		{
			name: "success with read on",
			setup: func(t *testing.T, mocks *test.DBMocks) {
				b := testBook(1, "1234567890123")
				err = mocks.BookDB.DB.Table(bookTable).Create(b).Error
				require.NoError(t, err)
			},
			args: args{
				bookshelf: &book.Bookshelf{
					BookID: 1,
					UserID: userID,
					Status: book.ReadStatus,
					ReadOn: test.DateMock,
				},
			},
			want: want{
				isErr: false,
			},
		},
		{
			name: "success with association",
			setup: func(t *testing.T, mocks *test.DBMocks) {
				b := testBook(1, "1234567890123")
				err = mocks.BookDB.DB.Table(bookTable).Create(b).Error
				require.NoError(t, err)
			},
			args: args{
				bookshelf: &book.Bookshelf{
					BookID: 1,
					UserID: userID,
					Status: book.ReadStatus,
					Review: &book.Review{
						BookID:     1,
						Impression: "テスト感想です。",
					},
				},
			},
			want: want{
				isErr: false,
			},
		},
		{
			name: "success with association and read on",
			setup: func(t *testing.T, mocks *test.DBMocks) {
				b := testBook(1, "1234567890123")
				err = mocks.BookDB.DB.Table(bookTable).Create(b).Error
				require.NoError(t, err)
			},
			args: args{
				bookshelf: &book.Bookshelf{
					BookID: 1,
					UserID: userID,
					Status: book.ReadStatus,
					ReadOn: test.DateMock,
					Review: &book.Review{
						BookID:     1,
						UserID:     userID,
						Impression: "テスト感想です。",
					},
				},
			},
			want: want{
				isErr: false,
			},
		},
		{
			name: "success to recreate with association",
			setup: func(t *testing.T, mocks *test.DBMocks) {
				b := testBook(1, "1234567890123")
				err = mocks.BookDB.DB.Table(bookTable).Create(b).Error
				require.NoError(t, err)
				rv := testReview(1, b.ID, userID)
				err := mocks.BookDB.DB.Table(reviewTable).Create(&rv).Error
				assert.NoError(t, err)
			},
			args: args{
				bookshelf: &book.Bookshelf{
					BookID: 1,
					UserID: userID,
					Status: book.ReadStatus,
					Review: &book.Review{
						BookID:     1,
						Impression: "テスト感想です。",
					},
				},
			},
			want: want{
				isErr: false,
			},
		},
		{
			name: "success to stacked status with association",
			setup: func(t *testing.T, mocks *test.DBMocks) {
				b := testBook(1, "1234567890123")
				err = mocks.BookDB.DB.Table(bookTable).Create(b).Error
				require.NoError(t, err)
				rv := testReview(1, b.ID, userID)
				err := mocks.BookDB.DB.Table(reviewTable).Create(&rv).Error
				assert.NoError(t, err)
			},
			args: args{
				bookshelf: &book.Bookshelf{
					BookID: 1,
					UserID: userID,
					Status: book.StackedStatus,
					Review: &book.Review{
						BookID:     1,
						Impression: "テスト感想です。",
					},
				},
			},
			want: want{
				isErr: false,
			},
		},
		{
			name:  "failed: internal error",
			setup: func(t *testing.T, mocks *test.DBMocks) {},
			args: args{
				bookshelf: &book.Bookshelf{},
			},
			want: want{
				isErr: true,
			},
		},
		{
			name:  "failed: internal error with association",
			setup: func(t *testing.T, mocks *test.DBMocks) {},
			args: args{
				bookshelf: &book.Bookshelf{
					BookID: 1,
					UserID: userID,
					Status: book.ReadStatus,
					ReadOn: test.DateMock,
					Review: &book.Review{
						BookID: 0,
						UserID: userID,
					},
				},
			},
			want: want{
				isErr: true,
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			err := mocks.Delete(mocks.BookDB, bookshelfTable, reviewTable, bookTable)
			require.NoError(t, err)
			tt.setup(t, mocks)

			target := NewBookRepository(mocks.BookDB, test.Now)
			err = target.CreateBookshelf(ctx, tt.args.bookshelf)
			assert.Equal(t, tt.want.isErr, err != nil, err)
		})
	}
}

func TestBookRepository_Update(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mocks, err := test.NewDBMock(ctrl)
	require.NoError(t, err)

	err = mocks.DeleteAll()
	require.NoError(t, err)

	type args struct {
		book *book.Book
	}
	type want struct {
		isErr bool
	}
	tests := []struct {
		name  string
		setup func(t *testing.T, mocks *test.DBMocks)
		args  args
		want  want
	}{
		{
			name: "success",
			setup: func(t *testing.T, mocks *test.DBMocks) {
				b := testBook(1, "1234567890123")
				err := mocks.BookDB.DB.Table(bookTable).Create(b).Error
				require.NoError(t, err)
			},
			args: args{
				book: &book.Book{
					ID:             1,
					Title:          "小説　ちはやふる　上の句",
					TitleKana:      "ショウセツ チハヤフルカミノク",
					Description:    "綾瀬千早は高校入学と同時に、競技かるた部を作ろうと奔走する。幼馴染の太一と仲間を集め、夏の全国大会に出場するためだ。強くなって、新と再会したい。幼い頃かるたを取り合った、新に寄せる千早の秘めた想いに気づきながらも、太一は千早を守り立てる。それぞれの青春を懸けた、一途な情熱の物語が幕開ける。",
					Isbn:           "1234567890123",
					Publisher:      "講談社",
					PublishedOn:    "2018-01-16",
					ThumbnailURL:   "https://thumbnail.image.rakuten.co.jp/@0_mall/book/cabinet/8426/9784062938426.jpg?_ex=120x120",
					RakutenURL:     "https://books.rakuten.co.jp/rb/15271426/",
					RakutenSize:    "コミック",
					RakutenGenreID: "001004008001/001004008003/001019001",
				},
			},
			want: want{
				isErr: false,
			},
		},
		{
			name: "success with association",
			setup: func(t *testing.T, mocks *test.DBMocks) {
				b := testBook(1, "1234567890123")
				err := mocks.BookDB.DB.Table(bookTable).Create(&b).Error
				require.NoError(t, err)
				as := make([]*book.Author, 2)
				as[0] = testAuthor(1, "有沢 ゆう希", "アリサワ ユウキ")
				as[1] = testAuthor(2, "末次 由紀", "スエツグ ユキ")
				err = mocks.BookDB.DB.Table(authorTable).Create(&as).Error
				require.NoError(t, err)
				bas := make([]*book.BookAuthor, 2)
				bas[0] = &book.BookAuthor{BookID: 1, AuthorID: 1, CreatedAt: test.TimeMock, UpdatedAt: test.TimeMock}
				bas[1] = &book.BookAuthor{BookID: 1, AuthorID: 2, CreatedAt: test.TimeMock, UpdatedAt: test.TimeMock}
				err = mocks.BookDB.DB.Table(authorBookTable).Create(&bas).Error
				require.NoError(t, err)
			},
			args: args{
				book: &book.Book{
					ID:             1,
					Title:          "小説　ちはやふる　上の句",
					TitleKana:      "ショウセツ チハヤフルカミノク",
					Description:    "綾瀬千早は高校入学と同時に、競技かるた部を作ろうと奔走する。幼馴染の太一と仲間を集め、夏の全国大会に出場するためだ。強くなって、新と再会したい。幼い頃かるたを取り合った、新に寄せる千早の秘めた想いに気づきながらも、太一は千早を守り立てる。それぞれの青春を懸けた、一途な情熱の物語が幕開ける。",
					Isbn:           "1234567890123",
					Publisher:      "講談社",
					PublishedOn:    "2018-01-16",
					ThumbnailURL:   "https://thumbnail.image.rakuten.co.jp/@0_mall/book/cabinet/8426/9784062938426.jpg?_ex=120x120",
					RakutenURL:     "https://books.rakuten.co.jp/rb/15271426/",
					RakutenSize:    "コミック",
					RakutenGenreID: "001004008001/001004008003/001019001",
					Authors: []*book.Author{
						{
							Name:     "有沢 ゆう希",
							NameKana: "アリサワ ユウキ",
						},
						{
							Name:     "テスト 著者",
							NameKana: "てすと ちょしゃ",
						},
					},
				},
			},
			want: want{
				isErr: false,
			},
		},
		{
			name:  "failed: internal error",
			setup: func(t *testing.T, mocks *test.DBMocks) {},
			args: args{
				book: &book.Book{},
			},
			want: want{
				isErr: true,
			},
		},
		{
			name: "failed with association",
			setup: func(t *testing.T, mocks *test.DBMocks) {
				b := testBook(1, "1234567890123")
				err := mocks.BookDB.DB.Table(bookTable).Create(b).Error
				require.NoError(t, err)
			},
			args: args{
				book: &book.Book{
					ID:             1,
					Title:          "小説　ちはやふる　上の句",
					TitleKana:      "ショウセツ チハヤフルカミノク",
					Description:    "綾瀬千早は高校入学と同時に、競技かるた部を作ろうと奔走する。幼馴染の太一と仲間を集め、夏の全国大会に出場するためだ。強くなって、新と再会したい。幼い頃かるたを取り合った、新に寄せる千早の秘めた想いに気づきながらも、太一は千早を守り立てる。それぞれの青春を懸けた、一途な情熱の物語が幕開ける。",
					Isbn:           "1234567890123",
					Publisher:      "講談社",
					PublishedOn:    "2018-01-16",
					ThumbnailURL:   "https://thumbnail.image.rakuten.co.jp/@0_mall/book/cabinet/8426/9784062938426.jpg?_ex=120x120",
					RakutenURL:     "https://books.rakuten.co.jp/rb/15271426/",
					RakutenSize:    "コミック",
					RakutenGenreID: "001004008001/001004008003/001019001",
					Authors: []*book.Author{
						{
							Name:     "",
							NameKana: "",
						},
					},
				},
			},
			want: want{
				isErr: true,
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			err := mocks.Delete(mocks.BookDB, authorBookTable, bookTable, authorTable)
			require.NoError(t, err)
			tt.setup(t, mocks)

			target := NewBookRepository(mocks.BookDB, test.Now)
			err = target.Update(ctx, tt.args.book)
			assert.Equal(t, tt.want.isErr, err != nil, err)
		})
	}
}

func TestBookRepository_MultipleUpdate(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mocks, err := test.NewDBMock(ctrl)
	require.NoError(t, err)

	err = mocks.DeleteAll()
	require.NoError(t, err)

	type args struct {
		books book.Books
	}
	type want struct {
		isErr bool
	}
	tests := []struct {
		name  string
		setup func(t *testing.T, mocks *test.DBMocks)
		args  args
		want  want
	}{
		{
			name: "success",
			setup: func(t *testing.T, mocks *test.DBMocks) {
				books := make([]*book.Book, 2)
				books[0] = testBook(1, "1234567890123")
				books[1] = testBook(2, "2345678901234")
				err := mocks.BookDB.DB.Table(bookTable).Create(&books).Error
				require.NoError(t, err)
			},
			args: args{
				books: []*book.Book{
					{
						ID:             1,
						Title:          "小説　ちはやふる　上の句",
						TitleKana:      "ショウセツ チハヤフルカミノク",
						Description:    "綾瀬千早は高校入学と同時に、競技かるた部を作ろうと奔走する。幼馴染の太一と仲間を集め、夏の全国大会に出場するためだ。強くなって、新と再会したい。幼い頃かるたを取り合った、新に寄せる千早の秘めた想いに気づきながらも、太一は千早を守り立てる。それぞれの青春を懸けた、一途な情熱の物語が幕開ける。",
						Isbn:           "1234567890123",
						Publisher:      "講談社",
						PublishedOn:    "2018-01-16",
						ThumbnailURL:   "https://thumbnail.image.rakuten.co.jp/@0_mall/book/cabinet/8426/9784062938426.jpg?_ex=120x120",
						RakutenURL:     "https://books.rakuten.co.jp/rb/15271426/",
						RakutenSize:    "コミック",
						RakutenGenreID: "001004008001/001004008003/001019001",
					},
					{
						ID:             2,
						Title:          "小説　ちはやふる　上の句",
						TitleKana:      "ショウセツ チハヤフルカミノク",
						Description:    "綾瀬千早は高校入学と同時に、競技かるた部を作ろうと奔走する。幼馴染の太一と仲間を集め、夏の全国大会に出場するためだ。強くなって、新と再会したい。幼い頃かるたを取り合った、新に寄せる千早の秘めた想いに気づきながらも、太一は千早を守り立てる。それぞれの青春を懸けた、一途な情熱の物語が幕開ける。",
						Isbn:           "2345678901234",
						Publisher:      "講談社",
						PublishedOn:    "2018-01-16",
						ThumbnailURL:   "https://thumbnail.image.rakuten.co.jp/@0_mall/book/cabinet/8426/9784062938426.jpg?_ex=120x120",
						RakutenURL:     "https://books.rakuten.co.jp/rb/15271426/",
						RakutenSize:    "コミック",
						RakutenGenreID: "001004008001/001004008003/001019001",
					},
				},
			},
			want: want{
				isErr: false,
			},
		},
		{
			name:  "failed: internal error",
			setup: func(t *testing.T, mocks *test.DBMocks) {},
			args: args{
				books: []*book.Book{{}},
			},
			want: want{
				isErr: true,
			},
		},
		{
			name: "failed: length 0",
			setup: func(t *testing.T, mocks *test.DBMocks) {
				books := make([]*book.Book, 2)
				books[0] = testBook(1, "1234567890123")
				books[1] = testBook(2, "2345678901234")
				err := mocks.BookDB.DB.Table(bookTable).Create(&books).Error
				require.NoError(t, err)
			},
			args: args{
				books: []*book.Book{},
			},
			want: want{
				isErr: true,
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			err := mocks.Delete(mocks.BookDB, authorTable, authorBookTable, bookTable)
			require.NoError(t, err)
			tt.setup(t, mocks)

			target := NewBookRepository(mocks.BookDB, test.Now)
			err = target.MultipleUpdate(ctx, tt.args.books)
			assert.Equal(t, tt.want.isErr, err != nil, err)
		})
	}
}

func TestBookRepository_UpdateBookshelf(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mocks, err := test.NewDBMock(ctrl)
	require.NoError(t, err)

	err = mocks.DeleteAll()
	require.NoError(t, err)

	userID := "12345678-1234-1234-1234-123456789012"

	type args struct {
		bookshelf *book.Bookshelf
	}
	type want struct {
		isErr bool
	}
	tests := []struct {
		name  string
		setup func(t *testing.T, mocks *test.DBMocks)
		args  args
		want  want
	}{
		{
			name: "success without read on",
			setup: func(t *testing.T, mocks *test.DBMocks) {
				b := testBook(1, "1234567890123")
				err := mocks.BookDB.DB.Table(bookTable).Create(b).Error
				require.NoError(t, err)

				bookshelf := testBookshelf(1, b.ID, userID)
				err = mocks.BookDB.DB.Table(bookshelfTable).Create(bookshelf).Error
				require.NoError(t, err)
			},
			args: args{
				bookshelf: &book.Bookshelf{
					ID:     1,
					BookID: 1,
					UserID: userID,
					Status: book.ReadStatus,
				},
			},
			want: want{
				isErr: false,
			},
		},
		{
			name:  "success with read on",
			setup: func(t *testing.T, mocks *test.DBMocks) {},
			args: args{
				bookshelf: &book.Bookshelf{
					ID:     1,
					BookID: 1,
					UserID: userID,
					ReadOn: test.DateMock,
					Status: book.ReadStatus,
				},
			},
			want: want{
				isErr: false,
			},
		},
		{
			name:  "success with read on and association",
			setup: func(t *testing.T, mocks *test.DBMocks) {},
			args: args{
				bookshelf: &book.Bookshelf{
					ID:     1,
					BookID: 1,
					UserID: userID,
					ReadOn: test.DateMock,
					Status: book.ReadStatus,
					Review: &book.Review{
						BookID:     1,
						UserID:     userID,
						Impression: "テスト感想です。",
					},
				},
			},
			want: want{
				isErr: false,
			},
		},
		{
			name: "success to recreate with read on",
			setup: func(t *testing.T, mocks *test.DBMocks) {
				rv := testReview(1, 1, userID)
				err := mocks.BookDB.DB.Table(reviewTable).Create(&rv).Error
				require.NoError(t, err)
			},
			args: args{
				bookshelf: &book.Bookshelf{
					ID:     1,
					BookID: 1,
					UserID: userID,
					ReadOn: test.DateMock,
					Status: book.ReadStatus,
				},
			},
			want: want{
				isErr: false,
			},
		},
		{
			name: "success to recreate with read on and association",
			setup: func(t *testing.T, mocks *test.DBMocks) {
				rv := testReview(1, 1, userID)
				err := mocks.BookDB.DB.Table(reviewTable).Create(&rv).Error
				require.NoError(t, err)
			},
			args: args{
				bookshelf: &book.Bookshelf{
					ID:     1,
					BookID: 1,
					UserID: userID,
					ReadOn: test.DateMock,
					Status: book.ReadStatus,
					Review: &book.Review{
						BookID:     1,
						UserID:     userID,
						Impression: "テスト感想です。",
					},
				},
			},
			want: want{
				isErr: false,
			},
		},
		{
			name:  "failed: internal error",
			setup: func(t *testing.T, mocks *test.DBMocks) {},
			args: args{
				bookshelf: &book.Bookshelf{},
			},
			want: want{
				isErr: true,
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			err := mocks.Delete(mocks.BookDB, bookshelfTable, reviewTable)
			require.NoError(t, err)
			tt.setup(t, mocks)

			target := NewBookRepository(mocks.BookDB, test.Now)
			err = target.UpdateBookshelf(ctx, tt.args.bookshelf)
			assert.Equal(t, tt.want.isErr, err != nil, err)
		})
	}
}

func TestBookRepository_Delete(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mocks, err := test.NewDBMock(ctrl)
	require.NoError(t, err)

	err = mocks.DeleteAll()
	require.NoError(t, err)

	type args struct {
		bookID int
	}
	type want struct {
		isErr bool
	}
	tests := []struct {
		name  string
		setup func(t *testing.T, mocks *test.DBMocks)
		args  args
		want  want
	}{
		{
			name: "success",
			setup: func(t *testing.T, mocks *test.DBMocks) {
				b := testBook(1, "1234567890123")
				err = mocks.BookDB.DB.Table(bookTable).Create(b).Error
				require.NoError(t, err)
			},
			args: args{
				bookID: 1,
			},
			want: want{
				isErr: false,
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			err := mocks.Delete(mocks.BookDB, bookTable)
			require.NoError(t, err)
			tt.setup(t, mocks)

			target := NewBookRepository(mocks.BookDB, test.Now)
			err = target.Delete(ctx, tt.args.bookID)
			assert.Equal(t, tt.want.isErr, err != nil, err)

			got, err := target.Get(ctx, tt.args.bookID)
			require.Error(t, err)
			assert.Equal(t, &book.Book{}, got)
		})
	}
}

func TestBookRepository_DeleteBookshelf(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mocks, err := test.NewDBMock(ctrl)
	require.NoError(t, err)

	err = mocks.DeleteAll()
	require.NoError(t, err)

	userID := "12345678-1234-1234-1234-123456789012"

	type args struct {
		bookshelfID int
	}
	type want struct {
		isErr bool
	}
	tests := []struct {
		name  string
		setup func(t *testing.T, mocks *test.DBMocks)
		args  args
		want  want
	}{
		{
			name: "success",
			setup: func(t *testing.T, mocks *test.DBMocks) {
				b := testBook(1, "1234567890123")
				err := mocks.BookDB.DB.Table(bookTable).Create(b).Error
				require.NoError(t, err)
				bookshelf := testBookshelf(1, b.ID, userID)
				err = mocks.BookDB.DB.Table(bookshelfTable).Create(bookshelf).Error
				require.NoError(t, err)
			},
			args: args{
				bookshelfID: 1,
			},
			want: want{
				isErr: false,
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			err := mocks.Delete(mocks.BookDB, bookshelfTable)
			require.NoError(t, err)
			tt.setup(t, mocks)

			target := NewBookRepository(mocks.BookDB, test.Now)
			err = target.DeleteBookshelf(ctx, tt.args.bookshelfID)
			assert.Equal(t, tt.want.isErr, err != nil, err)
		})
	}
}

func TestBookRepository_AggregateReadTotal(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mocks, err := test.NewDBMock(ctrl)
	require.NoError(t, err)

	err = mocks.DeleteAll()
	require.NoError(t, err)

	books := make(book.Books, 6)
	books[0] = testBook(1, "1234567890123")
	books[1] = testBook(2, "2345678901234")
	books[2] = testBook(3, "3456789012345")
	books[3] = testBook(4, "4567890123456")
	books[4] = testBook(5, "5678901234567")
	books[5] = testBook(6, "6789012345678")
	err = mocks.BookDB.DB.Table(bookTable).Create(&books).Error
	require.NoError(t, err)

	userID := "12345678-1234-1234-1234-123456789012"
	bookshelves := make(book.Bookshelves, 6)
	bookshelves[0] = testBookshelfWithReadOn(1, books[0].ID, userID, book.ReadStatus, "2021-08-02")
	bookshelves[1] = testBookshelfWithReadOn(2, books[1].ID, userID, book.ReadingStatus, "2021-08-02")
	bookshelves[2] = testBookshelfWithReadOn(3, books[2].ID, userID, book.ReadStatus, "2021-08-03")
	bookshelves[3] = testBookshelfWithReadOn(4, books[3].ID, userID, book.ReadStatus, "2021-09-15")
	bookshelves[4] = testBookshelfWithReadOn(5, books[4].ID, userID, book.ReadStatus, "2021-09-15")
	bookshelves[5] = testBookshelfWithReadOn(6, books[5].ID, userID, book.ReadStatus, "2020-08-15")
	err = mocks.BookDB.DB.Table(bookshelfTable).Create(&bookshelves).Error
	require.NoError(t, err)

	sinceDate, _ := datetime.ParseDate("2020-08-01")
	untilDate, _ := datetime.ParseDate("2021-09-01")

	type args struct {
		userID string
		since  time.Time
		until  time.Time
	}
	type want struct {
		results book.MonthlyResults
		isErr   bool
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "success",
			args: args{
				userID: userID,
				since:  datetime.BeginningOfMonth(sinceDate),
				until:  datetime.EndOfMonth(untilDate),
			},
			want: want{
				results: book.MonthlyResults{
					{
						Year:      2021,
						Month:     9,
						ReadTotal: 2,
					},
					{
						Year:      2021,
						Month:     8,
						ReadTotal: 2,
					},
					{
						Year:      2020,
						Month:     8,
						ReadTotal: 1,
					},
				},
				isErr: false,
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			target := NewBookRepository(mocks.BookDB, test.Now)

			results, err := target.AggregateReadTotal(ctx, tt.args.userID, tt.args.since, tt.args.until)
			assert.Equal(t, tt.want.isErr, err != nil, err)
			assert.Equal(t, tt.want.results, results)
		})
	}
}

func testBook(id int, isbn string) *book.Book {
	return &book.Book{
		ID:             id,
		Title:          "小説　ちはやふる　上の句",
		TitleKana:      "ショウセツ チハヤフルカミノク",
		Description:    "綾瀬千早は高校入学と同時に、競技かるた部を作ろうと奔走する。幼馴染の太一と仲間を集め、夏の全国大会に出場するためだ。強くなって、新と再会したい。幼い頃かるたを取り合った、新に寄せる千早の秘めた想いに気づきながらも、太一は千早を守り立てる。それぞれの青春を懸けた、一途な情熱の物語が幕開ける。",
		Isbn:           isbn,
		Publisher:      "講談社",
		PublishedOn:    "2018-01-16",
		ThumbnailURL:   "https://thumbnail.image.rakuten.co.jp/@0_mall/book/cabinet/8426/9784062938426.jpg?_ex=120x120",
		RakutenURL:     "https://books.rakuten.co.jp/rb/15271426/",
		RakutenSize:    "コミック",
		RakutenGenreID: "001004008001/001004008003/001019001",
		CreatedAt:      test.TimeMock,
		UpdatedAt:      test.TimeMock,
		Authors:        []*book.Author{},
	}
}

func testBookshelf(id int, bookID int, userID string) *book.Bookshelf {
	return &book.Bookshelf{
		ID:        id,
		BookID:    bookID,
		UserID:    userID,
		ReviewID:  0,
		Status:    book.ReadingStatus,
		ReadOn:    test.DateMock,
		CreatedAt: test.TimeMock,
		UpdatedAt: test.TimeMock,
	}
}

func testBookshelfWithReadOn(id int, bookID int, userID string, status int, readOn string) *book.Bookshelf {
	b := testBookshelf(id, bookID, userID)
	b.Status = status
	b.ReadOn, _ = datetime.ParseDate(readOn)

	return b
}

func testReview(id int, bookID int, userID string) *book.Review {
	return &book.Review{
		ID:         id,
		BookID:     bookID,
		UserID:     userID,
		Score:      3,
		Impression: "テストレビューです",
		CreatedAt:  test.TimeMock,
		UpdatedAt:  test.TimeMock,
	}
}

func testAuthor(id int, name string, nameKana string) *book.Author {
	return &book.Author{
		ID:        id,
		Name:      name,
		NameKana:  nameKana,
		CreatedAt: test.TimeMock,
		UpdatedAt: test.TimeMock,
	}
}
