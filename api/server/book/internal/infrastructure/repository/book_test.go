package repository

import (
	"context"
	"testing"

	"github.com/calmato/gran-book/api/server/book/internal/domain/book"
	"github.com/calmato/gran-book/api/server/book/pkg/database"
	"github.com/calmato/gran-book/api/server/book/pkg/test"
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

	err = mocks.DeleteAll(mocks.BookDB, "books")
	require.NoError(t, err)

	books := make([]*book.Book, 3)
	books[0] = testBook(1, "1234567890123")
	books[1] = testBook(2, "2345678901234")
	books[2] = testBook(3, "3456789012345")

	err = mocks.BookDB.DB.Table("books").Create(&books).Error
	require.NoError(t, err)

	type args struct {
		query *database.ListQuery
	}
	type want struct {
		books []*book.Book
		isErr bool
	}
	testCases := []struct {
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

	for _, tt := range testCases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			target := NewBookRepository(mocks.BookDB)

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

	err = mocks.DeleteAll(mocks.UserDB, "users")
	require.NoError(t, err)
	err = mocks.DeleteAll(mocks.BookDB, "bookshelves", "books")
	require.NoError(t, err)

	userID, err := mocks.CreateUser()
	require.NoError(t, err)
	books := make([]*book.Book, 3)
	books[0] = testBook(1, "1234567890123")
	books[1] = testBook(2, "2345678901234")
	books[2] = testBook(3, "3456789012345")
	err = mocks.BookDB.DB.Table("books").Create(&books).Error
	require.NoError(t, err)

	bookshelves := make([]*book.Bookshelf, 3)
	bookshelves[0] = testBookshelf(1, books[0].ID, userID)
	bookshelves[0].Book = books[0]
	bookshelves[1] = testBookshelf(2, books[1].ID, userID)
	bookshelves[1].Book = books[1]
	bookshelves[2] = testBookshelf(3, books[2].ID, userID)
	bookshelves[2].Book = books[2]

	err = mocks.BookDB.DB.Table("bookshelves").Create(&bookshelves).Error
	require.NoError(t, err)

	type args struct {
		query *database.ListQuery
	}
	type want struct {
		bookshelves []*book.Bookshelf
		isErr       bool
	}
	testCases := []struct {
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

	for _, tt := range testCases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			target := NewBookRepository(mocks.BookDB)

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

	err = mocks.DeleteAll(mocks.UserDB, "users")
	require.NoError(t, err)
	err = mocks.DeleteAll(mocks.BookDB, "reviews", "books")
	require.NoError(t, err)

	userID, err := mocks.CreateUser()
	require.NoError(t, err)
	books := make([]*book.Book, 3)
	books[0] = testBook(1, "1234567890123")
	books[1] = testBook(2, "2345678901234")
	books[2] = testBook(3, "3456789012345")
	err = mocks.BookDB.DB.Table("books").Create(&books).Error
	require.NoError(t, err)

	reviews := make([]*book.Review, 3)
	reviews[0] = testReview(1, books[0].ID, userID)
	reviews[1] = testReview(2, books[1].ID, userID)
	reviews[2] = testReview(3, books[2].ID, userID)

	err = mocks.BookDB.DB.Table("reviews").Create(&reviews).Error
	require.NoError(t, err)

	type args struct {
		query *database.ListQuery
	}
	type want struct {
		reviews []*book.Review
		isErr   bool
	}
	testCases := []struct {
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

	for _, tt := range testCases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			target := NewBookRepository(mocks.BookDB)

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

	err = mocks.DeleteAll(mocks.BookDB, "books")
	require.NoError(t, err)

	books := make([]*book.Book, 3)
	books[0] = testBook(1, "1234567890123")
	books[1] = testBook(2, "2345678901234")
	books[2] = testBook(3, "3456789012345")

	err = mocks.BookDB.DB.Table("books").Create(&books).Error
	require.NoError(t, err)

	type args struct {
		query *database.ListQuery
	}
	type want struct {
		count int
		isErr bool
	}
	testCases := []struct {
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

	for _, tt := range testCases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			target := NewBookRepository(mocks.BookDB)

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

	err = mocks.DeleteAll(mocks.UserDB, "users")
	require.NoError(t, err)
	err = mocks.DeleteAll(mocks.BookDB, "bookshelves", "books")
	require.NoError(t, err)

	userID, err := mocks.CreateUser()
	require.NoError(t, err)
	books := make([]*book.Book, 3)
	books[0] = testBook(1, "1234567890123")
	books[1] = testBook(2, "2345678901234")
	books[2] = testBook(3, "3456789012345")
	err = mocks.BookDB.DB.Table("books").Create(&books).Error
	require.NoError(t, err)

	bookshelves := make([]*book.Bookshelf, 3)
	bookshelves[0] = testBookshelf(1, books[0].ID, userID)
	bookshelves[0].Book = books[0]
	bookshelves[1] = testBookshelf(2, books[1].ID, userID)
	bookshelves[1].Book = books[1]
	bookshelves[2] = testBookshelf(3, books[2].ID, userID)
	bookshelves[2].Book = books[2]

	err = mocks.BookDB.DB.Table("bookshelves").Create(&bookshelves).Error
	require.NoError(t, err)

	type args struct {
		query *database.ListQuery
	}
	type want struct {
		count int
		isErr bool
	}
	testCases := []struct {
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

	for _, tt := range testCases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			target := NewBookRepository(mocks.BookDB)

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

	err = mocks.DeleteAll(mocks.UserDB, "users")
	require.NoError(t, err)
	err = mocks.DeleteAll(mocks.BookDB, "reviews", "books")
	require.NoError(t, err)

	userID, err := mocks.CreateUser()
	require.NoError(t, err)
	books := make([]*book.Book, 3)
	books[0] = testBook(1, "1234567890123")
	books[1] = testBook(2, "2345678901234")
	books[2] = testBook(3, "3456789012345")
	err = mocks.BookDB.DB.Table("books").Create(&books).Error
	require.NoError(t, err)

	reviews := make([]*book.Review, 3)
	reviews[0] = testReview(1, books[0].ID, userID)
	reviews[1] = testReview(2, books[1].ID, userID)
	reviews[2] = testReview(3, books[2].ID, userID)

	err = mocks.BookDB.DB.Table("reviews").Create(&reviews).Error
	require.NoError(t, err)

	type args struct {
		query *database.ListQuery
	}
	type want struct {
		count int
		isErr bool
	}
	testCases := []struct {
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

	for _, tt := range testCases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			target := NewBookRepository(mocks.BookDB)

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

	err = mocks.DeleteAll(mocks.BookDB, "books")
	require.NoError(t, err)

	books := make([]*book.Book, 3)
	books[0] = testBook(1, "1234567890123")
	books[1] = testBook(2, "2345678901234")
	books[2] = testBook(3, "3456789012345")

	err = mocks.BookDB.DB.Table("books").Create(&books).Error
	require.NoError(t, err)

	type args struct {
		bookIDs []int
	}
	type want struct {
		books []*book.Book
		isErr bool
	}
	testCases := []struct {
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

	for _, tt := range testCases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			target := NewBookRepository(mocks.BookDB)

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

	err = mocks.DeleteAll(mocks.BookDB, "books")
	require.NoError(t, err)

	expectBook := testBook(1, "1234567890123")
	err = mocks.BookDB.DB.Table("books").Create(expectBook).Error
	require.NoError(t, err)

	type args struct {
		bookID int
	}
	type want struct {
		book  *book.Book
		isErr bool
	}
	testCases := []struct {
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
				book:  nil,
				isErr: true,
			},
		},
	}

	for _, tt := range testCases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			target := NewBookRepository(mocks.BookDB)

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

	err = mocks.DeleteAll(mocks.BookDB, "books")
	require.NoError(t, err)

	expectBook := testBook(1, "1234567890123")
	err = mocks.BookDB.DB.Table("books").Create(expectBook).Error
	require.NoError(t, err)

	type args struct {
		isbn string
	}
	type want struct {
		book  *book.Book
		isErr bool
	}
	testCases := []struct {
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
				book:  nil,
				isErr: true,
			},
		},
	}

	for _, tt := range testCases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			target := NewBookRepository(mocks.BookDB)

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

	err = mocks.DeleteAll(mocks.BookDB, "books")
	require.NoError(t, err)

	expectBook := testBook(1, "1234567890123")
	err = mocks.BookDB.DB.Table("books").Create(expectBook).Error
	require.NoError(t, err)

	type args struct {
		isbn string
	}
	type want struct {
		bookID int
		isErr  bool
	}
	testCases := []struct {
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

	for _, tt := range testCases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			target := NewBookRepository(mocks.BookDB)

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

	err = mocks.DeleteAll(mocks.UserDB, "users")
	require.NoError(t, err)
	err = mocks.DeleteAll(mocks.BookDB, "bookshelves", "books")
	require.NoError(t, err)

	userID, err := mocks.CreateUser()
	require.NoError(t, err)
	books := make([]*book.Book, 1)
	books[0] = testBook(1, "1234567890123")
	err = mocks.BookDB.DB.Table("books").Create(&books).Error
	require.NoError(t, err)

	expectBookshelf := testBookshelf(1, books[0].ID, userID)
	expectBookshelf.Book = books[0]

	err = mocks.BookDB.DB.Table("bookshelves").Create(expectBookshelf).Error
	require.NoError(t, err)

	type args struct {
		userID string
		bookID int
	}
	type want struct {
		bookshelf *book.Bookshelf
		isErr     bool
	}
	testCases := []struct {
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
				bookshelf: nil,
				isErr:     true,
			},
		},
	}

	for _, tt := range testCases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			target := NewBookRepository(mocks.BookDB)

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

	err = mocks.DeleteAll(mocks.UserDB, "users")
	require.NoError(t, err)
	err = mocks.DeleteAll(mocks.BookDB, "bookshelves", "books")
	require.NoError(t, err)

	userID, err := mocks.CreateUser()
	require.NoError(t, err)
	books := make([]*book.Book, 1)
	books[0] = testBook(1, "1234567890123")
	err = mocks.BookDB.DB.Table("books").Create(&books).Error
	require.NoError(t, err)

	expectBookshelf := testBookshelf(1, books[0].ID, userID)
	expectBookshelf.Book = books[0]

	err = mocks.BookDB.DB.Table("bookshelves").Create(expectBookshelf).Error
	require.NoError(t, err)

	type args struct {
		userID string
		bookID int
	}
	type want struct {
		bookshelfID int
		isErr       bool
	}
	testCases := []struct {
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

	for _, tt := range testCases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			target := NewBookRepository(mocks.BookDB)

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

	err = mocks.DeleteAll(mocks.UserDB, "users")
	require.NoError(t, err)
	err = mocks.DeleteAll(mocks.BookDB, "reviews", "books")
	require.NoError(t, err)

	userID, err := mocks.CreateUser()
	require.NoError(t, err)
	books := make([]*book.Book, 1)
	books[0] = testBook(1, "1234567890123")
	err = mocks.BookDB.DB.Table("books").Create(&books).Error
	require.NoError(t, err)

	expectReview := testReview(1, books[0].ID, userID)

	err = mocks.BookDB.DB.Table("reviews").Create(expectReview).Error
	require.NoError(t, err)

	type args struct {
		reviewID int
	}
	type want struct {
		review *book.Review
		isErr  bool
	}
	testCases := []struct {
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
				review: nil,
				isErr:  true,
			},
		},
	}

	for _, tt := range testCases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			target := NewBookRepository(mocks.BookDB)

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

	err = mocks.DeleteAll(mocks.UserDB, "users")
	require.NoError(t, err)
	err = mocks.DeleteAll(mocks.BookDB, "reviews", "books")
	require.NoError(t, err)

	userID, err := mocks.CreateUser()
	require.NoError(t, err)
	books := make([]*book.Book, 1)
	books[0] = testBook(1, "1234567890123")
	err = mocks.BookDB.DB.Table("books").Create(&books).Error
	require.NoError(t, err)

	expectReview := testReview(1, books[0].ID, userID)

	err = mocks.BookDB.DB.Table("reviews").Create(expectReview).Error
	require.NoError(t, err)

	type args struct {
		userID string
		bookID int
	}
	type want struct {
		review *book.Review
		isErr  bool
	}
	testCases := []struct {
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
				review: nil,
				isErr:  true,
			},
		},
	}

	for _, tt := range testCases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			target := NewBookRepository(mocks.BookDB)

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

	err = mocks.DeleteAll(mocks.UserDB, "users")
	require.NoError(t, err)
	err = mocks.DeleteAll(mocks.BookDB, "reviews", "books")
	require.NoError(t, err)

	userID, err := mocks.CreateUser()
	require.NoError(t, err)
	books := make([]*book.Book, 1)
	books[0] = testBook(1, "1234567890123")
	err = mocks.BookDB.DB.Table("books").Create(&books).Error
	require.NoError(t, err)

	expectReview := testReview(1, books[0].ID, userID)

	err = mocks.BookDB.DB.Table("reviews").Create(expectReview).Error
	require.NoError(t, err)

	type args struct {
		userID string
		bookID int
	}
	type want struct {
		reviewID int
		isErr    bool
	}
	testCases := []struct {
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

	for _, tt := range testCases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			target := NewBookRepository(mocks.BookDB)

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

	err = mocks.DeleteAll(mocks.BookDB, "authors")
	require.NoError(t, err)

	expectAuthor := testAuthor(1, "有沢 ゆう希")
	err = mocks.BookDB.DB.Table("authors").Create(expectAuthor).Error
	require.NoError(t, err)

	type args struct {
		name string
	}
	type want struct {
		author *book.Author
		isErr  bool
	}
	testCases := []struct {
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
				author: nil,
				isErr:  true,
			},
		},
	}

	for _, tt := range testCases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			target := NewBookRepository(mocks.BookDB)

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

	err = mocks.DeleteAll(mocks.BookDB, "authors")
	require.NoError(t, err)

	expectAuthor := testAuthor(1, "有沢 ゆう希")
	err = mocks.BookDB.DB.Table("authors").Create(expectAuthor).Error
	require.NoError(t, err)

	type args struct {
		name string
	}
	type want struct {
		authorID int
		isErr    bool
	}
	testCases := []struct {
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

	for _, tt := range testCases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			target := NewBookRepository(mocks.BookDB)

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

	type args struct {
		book *book.Book
	}
	type want struct {
		isErr bool
	}
	testCases := []struct {
		name string
		args args
		want want
	}{
		{
			name: "success",
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
					CreatedAt:      test.TimeMock,
					UpdatedAt:      test.TimeMock,
				},
			},
			want: want{
				isErr: false,
			},
		},
		{
			name: "failed: internal error",
			args: args{
				book: &book.Book{},
			},
			want: want{
				isErr: true,
			},
		},
	}

	for _, tt := range testCases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			target := NewBookRepository(mocks.BookDB)

			err = mocks.DeleteAll(mocks.BookDB, "books")
			require.NoError(t, err)

			err := target.Create(ctx, tt.args.book)
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

	err = mocks.DeleteAll(mocks.BookDB, "books")
	require.NoError(t, err)

	userID, err := mocks.CreateUser()
	require.NoError(t, err)

	b := testBook(1, "1234567890123")
	err = mocks.BookDB.DB.Table("books").Create(b).Error
	require.NoError(t, err)

	type args struct {
		bookshelf *book.Bookshelf
	}
	type want struct {
		isErr bool
	}
	testCases := []struct {
		name string
		args args
		want want
	}{
		{
			name: "success",
			args: args{
				bookshelf: &book.Bookshelf{
					BookID:    b.ID,
					UserID:    userID,
					Status:    book.ReadStatus,
					CreatedAt: test.TimeMock,
					UpdatedAt: test.TimeMock,
				},
			},
			want: want{
				isErr: false,
			},
		},
		{
			name: "failed: internal error",
			args: args{
				bookshelf: &book.Bookshelf{},
			},
			want: want{
				isErr: true,
			},
		},
	}

	for _, tt := range testCases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			target := NewBookRepository(mocks.BookDB)

			err := mocks.DeleteAll(mocks.BookDB, "bookshelves")
			require.NoError(t, err)

			err = target.CreateBookshelf(ctx, tt.args.bookshelf)
			assert.Equal(t, tt.want.isErr, err != nil, err)
		})
	}
}

func TestBookRepository_CreateReview(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mocks, err := test.NewDBMock(ctrl)
	require.NoError(t, err)

	err = mocks.DeleteAll(mocks.BookDB, "books")
	require.NoError(t, err)

	userID, err := mocks.CreateUser()
	require.NoError(t, err)

	b := testBook(1, "1234567890123")
	err = mocks.BookDB.DB.Table("books").Create(b).Error
	require.NoError(t, err)

	type args struct {
		review *book.Review
	}
	type want struct {
		isErr bool
	}
	testCases := []struct {
		name string
		args args
		want want
	}{
		{
			name: "success",
			args: args{
				review: &book.Review{
					BookID:     b.ID,
					UserID:     userID,
					Score:      3,
					Impression: "テストレビューです",
					CreatedAt:  test.TimeMock,
					UpdatedAt:  test.TimeMock,
				},
			},
			want: want{
				isErr: false,
			},
		},
		{
			name: "failed: internal error",
			args: args{
				review: &book.Review{},
			},
			want: want{
				isErr: true,
			},
		},
	}

	for _, tt := range testCases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			target := NewBookRepository(mocks.BookDB)

			err := mocks.DeleteAll(mocks.BookDB, "reviews")
			require.NoError(t, err)

			err = target.CreateReview(ctx, tt.args.review)
			assert.Equal(t, tt.want.isErr, err != nil, err)
		})
	}
}

func TestBookRepository_CreateAuthor(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mocks, err := test.NewDBMock(ctrl)
	require.NoError(t, err)

	type args struct {
		author *book.Author
	}
	type want struct {
		isErr bool
	}
	testCases := []struct {
		name string
		args args
		want want
	}{
		{
			name: "success",
			args: args{
				author: &book.Author{
					Name:      "有沢 ゆう希",
					NameKana:  "アリサワ ユウキ",
					CreatedAt: test.TimeMock,
					UpdatedAt: test.TimeMock,
				},
			},
			want: want{
				isErr: false,
			},
		},
		{
			name: "failed: internal error",
			args: args{
				author: &book.Author{},
			},
			want: want{
				isErr: true,
			},
		},
	}

	for _, tt := range testCases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			target := NewBookRepository(mocks.BookDB)

			err := mocks.DeleteAll(mocks.BookDB, "authors")
			require.NoError(t, err)

			err = target.CreateAuthor(ctx, tt.args.author)
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

	b := testBook(1, "1234567890123")

	type args struct {
		book *book.Book
	}
	type want struct {
		isErr bool
	}
	testCases := []struct {
		name string
		args args
		want want
	}{
		{
			name: "success",
			args: args{
				book: &book.Book{
					ID:             b.ID,
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
					CreatedAt:      test.TimeMock,
					UpdatedAt:      test.TimeMock,
				},
			},
			want: want{
				isErr: false,
			},
		},
		{
			name: "failed: internal error",
			args: args{
				book: &book.Book{},
			},
			want: want{
				isErr: true,
			},
		},
	}

	for _, tt := range testCases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			target := NewBookRepository(mocks.BookDB)

			err = mocks.DeleteAll(mocks.BookDB, "books")
			require.NoError(t, err)

			err = mocks.BookDB.DB.Table("books").Create(b).Error
			require.NoError(t, err)

			err := target.Update(ctx, tt.args.book)
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

	err = mocks.DeleteAll(mocks.BookDB, "books")
	require.NoError(t, err)

	userID, err := mocks.CreateUser()
	require.NoError(t, err)

	b := testBook(1, "1234567890123")
	err = mocks.BookDB.DB.Table("books").Create(b).Error
	require.NoError(t, err)

	bookshelf := testBookshelf(1, b.ID, userID)

	type args struct {
		bookshelf *book.Bookshelf
	}
	type want struct {
		isErr bool
	}
	testCases := []struct {
		name string
		args args
		want want
	}{
		{
			name: "success",
			args: args{
				bookshelf: &book.Bookshelf{
					ID:        bookshelf.ID,
					BookID:    b.ID,
					UserID:    userID,
					Status:    book.ReadStatus,
					CreatedAt: test.TimeMock,
					UpdatedAt: test.TimeMock,
				},
			},
			want: want{
				isErr: false,
			},
		},
		{
			name: "failed: internal error",
			args: args{
				bookshelf: &book.Bookshelf{},
			},
			want: want{
				isErr: true,
			},
		},
	}

	for _, tt := range testCases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			target := NewBookRepository(mocks.BookDB)

			err := mocks.DeleteAll(mocks.BookDB, "bookshelves")
			require.NoError(t, err)

			err = mocks.BookDB.DB.Table("bookshelves").Create(bookshelf).Error
			require.NoError(t, err)

			err = target.UpdateBookshelf(ctx, tt.args.bookshelf)
			assert.Equal(t, tt.want.isErr, err != nil, err)
		})
	}
}

func TestBookRepository_UpdateReview(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mocks, err := test.NewDBMock(ctrl)
	require.NoError(t, err)

	err = mocks.DeleteAll(mocks.BookDB, "books")
	require.NoError(t, err)

	userID, err := mocks.CreateUser()
	require.NoError(t, err)

	b := testBook(1, "1234567890123")
	err = mocks.BookDB.DB.Table("books").Create(b).Error
	require.NoError(t, err)

	review := testReview(1, b.ID, userID)

	type args struct {
		review *book.Review
	}
	type want struct {
		isErr bool
	}
	testCases := []struct {
		name string
		args args
		want want
	}{
		{
			name: "success",
			args: args{
				review: &book.Review{
					ID:         review.ID,
					BookID:     b.ID,
					UserID:     userID,
					Score:      3,
					Impression: "テストレビューです",
					CreatedAt:  test.TimeMock,
					UpdatedAt:  test.TimeMock,
				},
			},
			want: want{
				isErr: false,
			},
		},
		{
			name: "failed: internal error",
			args: args{
				review: &book.Review{},
			},
			want: want{
				isErr: true,
			},
		},
	}

	for _, tt := range testCases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			target := NewBookRepository(mocks.BookDB)

			err := mocks.DeleteAll(mocks.BookDB, "reviews")
			require.NoError(t, err)

			err = mocks.BookDB.DB.Table("reviews").Create(review).Error
			require.NoError(t, err)

			err = target.UpdateReview(ctx, tt.args.review)
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

	type args struct {
		books []*book.Book
	}
	type want struct {
		isErr bool
	}
	testCases := []struct {
		name string
		args args
		want want
	}{
		{
			name: "success",
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
						CreatedAt:      test.TimeMock,
						UpdatedAt:      test.TimeMock,
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
						CreatedAt:      test.TimeMock,
						UpdatedAt:      test.TimeMock,
					},
				},
			},
			want: want{
				isErr: false,
			},
		},
		{
			name: "success: not found",
			args: args{
				books: []*book.Book{},
			},
			want: want{
				isErr: false,
			},
		},
		{
			name: "failed: internal error",
			args: args{
				books: []*book.Book{{}},
			},
			want: want{
				isErr: true,
			},
		},
	}

	for _, tt := range testCases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			target := NewBookRepository(mocks.BookDB)

			err = mocks.DeleteAll(mocks.BookDB, "books")
			require.NoError(t, err)

			err := target.MultipleCreate(ctx, tt.args.books)
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

	books := make([]*book.Book, 2)
	books[0] = testBook(1, "1234567890123")
	books[1] = testBook(2, "2345678901234")

	type args struct {
		books []*book.Book
	}
	type want struct {
		isErr bool
	}
	testCases := []struct {
		name string
		args args
		want want
	}{
		{
			name: "success",
			args: args{
				books: []*book.Book{
					{
						ID:             books[0].ID,
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
						CreatedAt:      test.TimeMock,
						UpdatedAt:      test.TimeMock,
					},
					{
						ID:             books[1].ID,
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
						CreatedAt:      test.TimeMock,
						UpdatedAt:      test.TimeMock,
					},
				},
			},
			want: want{
				isErr: false,
			},
		},
		{
			name: "success: not found",
			args: args{
				books: []*book.Book{},
			},
			want: want{
				isErr: false,
			},
		},
		{
			name: "failed: internal error",
			args: args{
				books: []*book.Book{{}},
			},
			want: want{
				isErr: true,
			},
		},
	}

	for _, tt := range testCases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			target := NewBookRepository(mocks.BookDB)

			err = mocks.DeleteAll(mocks.BookDB, "books")
			require.NoError(t, err)

			err = mocks.BookDB.DB.Table("books").Create(&books).Error
			require.NoError(t, err)

			err := target.MultipleUpdate(ctx, tt.args.books)
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

	b := testBook(1, "1234567890123")

	type args struct {
		bookID int
	}
	type want struct {
		isErr bool
	}
	testCases := []struct {
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
				isErr: false,
			},
		},
	}

	for _, tt := range testCases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			target := NewBookRepository(mocks.BookDB)

			err = mocks.DeleteAll(mocks.BookDB, "books")
			require.NoError(t, err)

			err = mocks.BookDB.DB.Table("books").Create(b).Error
			require.NoError(t, err)

			err := target.Delete(ctx, tt.args.bookID)
			assert.Equal(t, tt.want.isErr, err != nil, err)

			got, err := target.Get(ctx, tt.args.bookID)
			require.Error(t, err)
			assert.Nil(t, got)
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

	err = mocks.DeleteAll(mocks.BookDB, "books")
	require.NoError(t, err)

	userID, err := mocks.CreateUser()
	require.NoError(t, err)

	b := testBook(1, "1234567890123")
	err = mocks.BookDB.DB.Table("books").Create(b).Error
	require.NoError(t, err)

	bookshelf := testBookshelf(1, b.ID, userID)

	type args struct {
		bookshelfID int
	}
	type want struct {
		isErr bool
	}
	testCases := []struct {
		name string
		args args
		want want
	}{
		{
			name: "success",
			args: args{
				bookshelfID: 1,
			},
			want: want{
				isErr: false,
			},
		},
	}

	for _, tt := range testCases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			target := NewBookRepository(mocks.BookDB)

			err := mocks.DeleteAll(mocks.BookDB, "bookshelves")
			require.NoError(t, err)

			err = mocks.BookDB.DB.Table("bookshelves").Create(bookshelf).Error
			require.NoError(t, err)

			err = target.DeleteBookshelf(ctx, tt.args.bookshelfID)
			assert.Equal(t, tt.want.isErr, err != nil, err)
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

func testAuthor(id int, name string) *book.Author {
	return &book.Author{
		ID:        id,
		Name:      name,
		NameKana:  name,
		CreatedAt: test.TimeMock,
		UpdatedAt: test.TimeMock,
	}
}
