package application

import (
	"context"
	"fmt"
	"testing"

	"github.com/calmato/gran-book/api/server/book/internal/domain/book"
	"github.com/calmato/gran-book/api/server/book/internal/domain/exception"
	"github.com/calmato/gran-book/api/server/book/pkg/database"
	"github.com/calmato/gran-book/api/server/book/pkg/datetime"
	"github.com/calmato/gran-book/api/server/book/pkg/test"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBookApplication_List(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	book1 := testBook(1)
	book2 := testBook(2)

	type args struct {
		query *database.ListQuery
	}
	type want struct {
		books book.Books
		total int
		err   error
	}
	testCases := []struct {
		name  string
		setup func(context.Context, *testing.T, *test.Mocks)
		args  args
		want  want
	}{
		{
			name: "success",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.BookRepository.EXPECT().
					List(ctx, gomock.Any()).
					Return([]*book.Book{book1, book2}, nil)
				mocks.BookRepository.EXPECT().
					Count(ctx, gomock.Any()).
					Return(2, nil)
			},
			args: args{
				query: &database.ListQuery{},
			},
			want: want{
				books: []*book.Book{book1, book2},
				total: 2,
				err:   nil,
			},
		},
		{
			name: "failed: internal error in list",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.BookRepository.EXPECT().
					List(ctx, gomock.Any()).
					Return(nil, exception.ErrorInDatastore.New(test.ErrMock))
			},
			args: args{
				query: &database.ListQuery{},
			},
			want: want{
				books: nil,
				total: 0,
				err:   exception.ErrorInDatastore.New(test.ErrMock),
			},
		},
		{
			name: "failed: internal error in count",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.BookRepository.EXPECT().
					List(ctx, gomock.Any()).
					Return([]*book.Book{book1, book2}, nil)
				mocks.BookRepository.EXPECT().
					Count(ctx, gomock.Any()).
					Return(0, exception.ErrorInDatastore.New(test.ErrMock))
			},
			args: args{
				query: &database.ListQuery{},
			},
			want: want{
				books: nil,
				total: 0,
				err:   exception.ErrorInDatastore.New(test.ErrMock),
			},
		},
	}

	for _, tt := range testCases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			mocks := test.NewMocks(ctrl)
			tt.setup(ctx, t, mocks)
			target := NewBookApplication(
				mocks.BookDomainValidation,
				mocks.BookRepository,
			)

			bs, total, err := target.List(ctx, tt.args.query)
			if tt.want.err != nil {
				require.Equal(t, tt.want.err.Error(), err.Error())
				require.Equal(t, tt.want.books, bs)
				require.Equal(t, tt.want.total, total)
				return
			}
			require.NoError(t, err)
			require.Equal(t, tt.want.books, bs)
			require.Equal(t, tt.want.total, total)
		})
	}
}

func TestBookApplication_ListBookshelf(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	book1 := testBook(1)
	bookshelf1 := testBookshelf(1, book1.ID, "user01")
	bookshelf2 := testBookshelf(2, book1.ID, "user01")

	type args struct {
		query *database.ListQuery
	}
	type want struct {
		bookshelves book.Bookshelves
		total       int
		err         error
	}
	testCases := []struct {
		name  string
		setup func(context.Context, *testing.T, *test.Mocks)
		args  args
		want  want
	}{
		{
			name: "success",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.BookRepository.EXPECT().
					ListBookshelf(ctx, gomock.Any()).
					Return([]*book.Bookshelf{bookshelf1, bookshelf2}, nil)
				mocks.BookRepository.EXPECT().
					CountBookshelf(ctx, gomock.Any()).
					Return(2, nil)
			},
			args: args{
				query: &database.ListQuery{},
			},
			want: want{
				bookshelves: []*book.Bookshelf{bookshelf1, bookshelf2},
				total:       2,
				err:         nil,
			},
		},
		{
			name: "failed: internal error in list",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.BookRepository.EXPECT().
					ListBookshelf(ctx, gomock.Any()).
					Return(nil, exception.ErrorInDatastore.New(test.ErrMock))
			},
			args: args{
				query: &database.ListQuery{},
			},
			want: want{
				bookshelves: nil,
				total:       0,
				err:         exception.ErrorInDatastore.New(test.ErrMock),
			},
		},
		{
			name: "failed: internal error in count",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.BookRepository.EXPECT().
					ListBookshelf(ctx, gomock.Any()).
					Return([]*book.Bookshelf{bookshelf1, bookshelf2}, nil)
				mocks.BookRepository.EXPECT().
					CountBookshelf(ctx, gomock.Any()).
					Return(0, exception.ErrorInDatastore.New(test.ErrMock))
			},
			args: args{
				query: &database.ListQuery{},
			},
			want: want{
				bookshelves: nil,
				total:       0,
				err:         exception.ErrorInDatastore.New(test.ErrMock),
			},
		},
	}

	for _, tt := range testCases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			mocks := test.NewMocks(ctrl)
			tt.setup(ctx, t, mocks)
			target := NewBookApplication(
				mocks.BookDomainValidation,
				mocks.BookRepository,
			)

			bss, total, err := target.ListBookshelf(ctx, tt.args.query)
			if tt.want.err != nil {
				require.Equal(t, tt.want.err.Error(), err.Error())
				require.Equal(t, tt.want.bookshelves, bss)
				require.Equal(t, tt.want.total, total)
				return
			}
			require.NoError(t, err)
			require.Equal(t, tt.want.bookshelves, bss)
			require.Equal(t, tt.want.total, total)
		})
	}
}

func TestBookApplication_ListBookReview(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	book1 := testBook(1)
	review1 := testReview(1, book1.ID, "user01")
	review2 := testReview(2, book1.ID, "user02")

	type args struct {
		bookID int
		limit  int
		offset int
	}
	type want struct {
		reviews book.Reviews
		total   int
		err     error
	}
	testCases := []struct {
		name  string
		setup func(context.Context, *testing.T, *test.Mocks)
		args  args
		want  want
	}{
		{
			name: "success",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.BookRepository.EXPECT().
					ListReview(ctx, gomock.Any()).
					Return([]*book.Review{review1, review2}, nil)
				mocks.BookRepository.EXPECT().
					CountReview(ctx, gomock.Any()).
					Return(2, nil)
			},
			args: args{
				bookID: book1.ID,
				limit:  20,
				offset: 0,
			},
			want: want{
				reviews: []*book.Review{review1, review2},
				total:   2,
				err:     nil,
			},
		},
		{
			name: "failed: internal error in list",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.BookRepository.EXPECT().
					ListReview(ctx, gomock.Any()).
					Return(nil, exception.ErrorInDatastore.New(test.ErrMock))
			},
			args: args{
				bookID: book1.ID,
				limit:  20,
				offset: 0,
			},
			want: want{
				reviews: nil,
				total:   0,
				err:     exception.ErrorInDatastore.New(test.ErrMock),
			},
		},
		{
			name: "failed: internal error in count",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.BookRepository.EXPECT().
					ListReview(ctx, gomock.Any()).
					Return([]*book.Review{review1, review2}, nil)
				mocks.BookRepository.EXPECT().
					CountReview(ctx, gomock.Any()).
					Return(0, exception.ErrorInDatastore.New(test.ErrMock))
			},
			args: args{
				bookID: book1.ID,
				limit:  20,
				offset: 0,
			},
			want: want{
				reviews: nil,
				total:   0,
				err:     exception.ErrorInDatastore.New(test.ErrMock),
			},
		},
	}

	for _, tt := range testCases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			mocks := test.NewMocks(ctrl)
			tt.setup(ctx, t, mocks)
			target := NewBookApplication(
				mocks.BookDomainValidation,
				mocks.BookRepository,
			)

			rs, total, err := target.ListBookReview(ctx, tt.args.bookID, tt.args.limit, tt.args.offset)
			if tt.want.err != nil {
				require.Equal(t, tt.want.err.Error(), err.Error())
				require.Equal(t, tt.want.reviews, rs)
				require.Equal(t, tt.want.total, total)
				return
			}
			require.NoError(t, err)
			require.Equal(t, tt.want.reviews, rs)
			require.Equal(t, tt.want.total, total)
		})
	}
}

func TestBookApplication_ListUserReview(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	book1 := testBook(1)
	book2 := testBook(2)
	review1 := testReview(1, book1.ID, "user01")
	review2 := testReview(2, book2.ID, "user01")

	type args struct {
		userID string
		limit  int
		offset int
	}
	type want struct {
		reviews book.Reviews
		total   int
		err     error
	}
	testCases := []struct {
		name  string
		setup func(context.Context, *testing.T, *test.Mocks)
		args  args
		want  want
	}{
		{
			name: "success",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.BookRepository.EXPECT().
					ListReview(ctx, gomock.Any()).
					Return([]*book.Review{review1, review2}, nil)
				mocks.BookRepository.EXPECT().
					CountReview(ctx, gomock.Any()).
					Return(2, nil)
			},
			args: args{
				userID: "user01",
				limit:  20,
				offset: 0,
			},
			want: want{
				reviews: []*book.Review{review1, review2},
				total:   2,
				err:     nil,
			},
		},
		{
			name: "failed: internal error in list",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.BookRepository.EXPECT().
					ListReview(ctx, gomock.Any()).
					Return(nil, exception.ErrorInDatastore.New(test.ErrMock))
			},
			args: args{
				userID: "user01",
				limit:  20,
				offset: 0,
			},
			want: want{
				reviews: nil,
				total:   0,
				err:     exception.ErrorInDatastore.New(test.ErrMock),
			},
		},
		{
			name: "failed: internal error in count",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.BookRepository.EXPECT().
					ListReview(ctx, gomock.Any()).
					Return([]*book.Review{review1, review2}, nil)
				mocks.BookRepository.EXPECT().
					CountReview(ctx, gomock.Any()).
					Return(0, exception.ErrorInDatastore.New(test.ErrMock))
			},
			args: args{
				userID: "user01",
				limit:  20,
				offset: 0,
			},
			want: want{
				reviews: nil,
				total:   0,
				err:     exception.ErrorInDatastore.New(test.ErrMock),
			},
		},
	}

	for _, tt := range testCases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			mocks := test.NewMocks(ctrl)
			tt.setup(ctx, t, mocks)
			target := NewBookApplication(
				mocks.BookDomainValidation,
				mocks.BookRepository,
			)

			rs, total, err := target.ListUserReview(ctx, tt.args.userID, tt.args.limit, tt.args.offset)
			if tt.want.err != nil {
				require.Equal(t, tt.want.err.Error(), err.Error())
				require.Equal(t, tt.want.reviews, rs)
				require.Equal(t, tt.want.total, total)
				return
			}
			require.NoError(t, err)
			require.Equal(t, tt.want.reviews, rs)
			require.Equal(t, tt.want.total, total)
		})
	}
}

func TestBookApplication_ListUserMonthlyResult(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	results := make(book.MonthlyResults, 2)
	results[0] = &book.MonthlyResult{Year: 2021, Month: 8, ReadTotal: 3}
	results[1] = &book.MonthlyResult{Year: 2021, Month: 9, ReadTotal: 8}

	type args struct {
		userID string
		since  string
		until  string
	}
	type want struct {
		results book.MonthlyResults
		err     error
	}
	tests := []struct {
		name  string
		setup func(context.Context, *testing.T, *test.Mocks)
		args  args
		want  want
	}{
		{
			name: "success",
			setup: func(c context.Context, t *testing.T, m *test.Mocks) {
				sinceDate, _ := datetime.ParseDate("2021-08-01")
				untilDate, _ := datetime.ParseDate("2021-09-01")
				since := datetime.BeginningOfMonth(sinceDate)
				until := datetime.EndOfMonth(untilDate)
				m.BookRepository.EXPECT().
					AggregateReadTotal(ctx, "00000000-0000-0000-0000-000000000000", since, until).
					Return(results, nil)
			},
			args: args{
				userID: "00000000-0000-0000-0000-000000000000",
				since:  "2021-08-01",
				until:  "2021-09-01",
			},
			want: want{
				results: results,
				err:     nil,
			},
		},
		{
			name:  "failed: invalid date format",
			setup: func(c context.Context, t *testing.T, m *test.Mocks) {},
			args: args{
				userID: "00000000-0000-0000-0000-000000000000",
				since:  "2021-08-00",
				until:  "2021-09-01",
			},
			want: want{
				results: nil,
				err:     exception.InvalidRequestValidation.New(errInvalidDateFormat),
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			mocks := test.NewMocks(ctrl)
			tt.setup(ctx, t, mocks)
			target := NewBookApplication(
				mocks.BookDomainValidation,
				mocks.BookRepository,
			)

			rs, err := target.ListUserMonthlyResult(ctx, tt.args.userID, tt.args.since, tt.args.until)
			if tt.want.err != nil {
				fmt.Printf("debug: %+v", tt.want.err)
				require.Equal(t, tt.want.err.Error(), err.Error())
				assert.Equal(t, tt.want.results, rs)
				return
			}
			require.NoError(t, err)
			assert.Equal(t, tt.want.results, rs)
		})
	}
}

func TestBookApplication_MultiGet(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	book1 := testBook(1)
	book2 := testBook(2)

	type args struct {
		bookIDs []int
	}
	type want struct {
		books book.Books
		err   error
	}
	testCases := []struct {
		name  string
		setup func(context.Context, *testing.T, *test.Mocks)
		args  args
		want  want
	}{
		{
			name: "success",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.BookRepository.EXPECT().
					MultiGet(ctx, []int{1, 2}).
					Return([]*book.Book{book1, book2}, nil)
			},
			args: args{
				bookIDs: []int{1, 2},
			},
			want: want{
				books: []*book.Book{book1, book2},
				err:   nil,
			},
		},
	}

	for _, tt := range testCases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			mocks := test.NewMocks(ctrl)
			tt.setup(ctx, t, mocks)
			target := NewBookApplication(
				mocks.BookDomainValidation,
				mocks.BookRepository,
			)

			bs, err := target.MultiGet(ctx, tt.args.bookIDs)
			if tt.want.err != nil {
				require.Equal(t, tt.want.err.Error(), err.Error())
				require.Equal(t, tt.want.books, bs)
				return
			}
			require.NoError(t, err)
			require.Equal(t, tt.want.books, bs)
		})
	}
}

func TestBookApplication_Get(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	book1 := testBook(1)

	type args struct {
		bookID int
	}
	type want struct {
		book *book.Book
		err  error
	}
	testCases := []struct {
		name  string
		setup func(context.Context, *testing.T, *test.Mocks)
		args  args
		want  want
	}{
		{
			name: "success",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.BookRepository.EXPECT().
					Get(ctx, 1).
					Return(book1, nil)
			},
			args: args{
				bookID: 1,
			},
			want: want{
				book: book1,
				err:  nil,
			},
		},
	}

	for _, tt := range testCases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			mocks := test.NewMocks(ctrl)
			tt.setup(ctx, t, mocks)
			target := NewBookApplication(
				mocks.BookDomainValidation,
				mocks.BookRepository,
			)

			b, err := target.Get(ctx, tt.args.bookID)
			if tt.want.err != nil {
				require.Equal(t, tt.want.err.Error(), err.Error())
				require.Equal(t, tt.want.book, b)
				return
			}
			require.NoError(t, err)
			require.Equal(t, tt.want.book, b)
		})
	}
}

func TestBookApplication_GetByIsbn(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	book1 := testBook(1)

	type args struct {
		isbn string
	}
	type want struct {
		book *book.Book
		err  error
	}
	testCases := []struct {
		name  string
		setup func(context.Context, *testing.T, *test.Mocks)
		args  args
		want  want
	}{
		{
			name: "success",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.BookRepository.EXPECT().
					GetByIsbn(ctx, "9784062938426").
					Return(book1, nil)
			},
			args: args{
				isbn: "9784062938426",
			},
			want: want{
				book: book1,
				err:  nil,
			},
		},
	}

	for _, tt := range testCases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			mocks := test.NewMocks(ctrl)
			tt.setup(ctx, t, mocks)
			target := NewBookApplication(
				mocks.BookDomainValidation,
				mocks.BookRepository,
			)

			b, err := target.GetByIsbn(ctx, tt.args.isbn)
			if tt.want.err != nil {
				require.Equal(t, tt.want.err.Error(), err.Error())
				require.Equal(t, tt.want.book, b)
				return
			}
			require.NoError(t, err)
			require.Equal(t, tt.want.book, b)
		})
	}
}

func TestBookApplication_GetBookshelfByUserIDAndBookID(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	book1 := testBook(1)
	bookshelf1 := testBookshelf(1, book1.ID, "user01")

	type args struct {
		userID string
		bookID int
	}
	type want struct {
		bookshelf *book.Bookshelf
		err       error
	}
	testCases := []struct {
		name  string
		setup func(context.Context, *testing.T, *test.Mocks)
		args  args
		want  want
	}{
		{
			name: "success",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.BookRepository.EXPECT().
					GetBookshelfByUserIDAndBookID(ctx, "user01", book1.ID).
					Return(bookshelf1, nil)
			},
			args: args{
				userID: "user01",
				bookID: 1,
			},
			want: want{
				bookshelf: bookshelf1,
				err:       nil,
			},
		},
	}

	for _, tt := range testCases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			mocks := test.NewMocks(ctrl)
			tt.setup(ctx, t, mocks)
			target := NewBookApplication(
				mocks.BookDomainValidation,
				mocks.BookRepository,
			)

			bs, err := target.GetBookshelfByUserIDAndBookID(ctx, tt.args.userID, tt.args.bookID)
			if tt.want.err != nil {
				require.Equal(t, tt.want.err.Error(), err.Error())
				require.Equal(t, tt.want.bookshelf, bs)
				return
			}
			require.NoError(t, err)
			require.Equal(t, tt.want.bookshelf, bs)
		})
	}
}

func TestBookApplication_GetBookshelfByUserIDAndBookIDWithRelated(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	book1 := testBook(1)
	review1 := testReview(1, book1.ID, "user01")
	bookshelf1 := testBookshelf(1, book1.ID, "user01")
	bookshelf1.ReviewID = review1.ID

	type args struct {
		userID string
		bookID int
	}
	type want struct {
		bookshelf *book.Bookshelf
		err       error
	}
	testCases := []struct {
		name  string
		setup func(context.Context, *testing.T, *test.Mocks)
		args  args
		want  want
	}{
		{
			name: "success",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.BookRepository.EXPECT().
					Get(ctx, bookshelf1.BookID).
					Return(book1, nil)
				mocks.BookRepository.EXPECT().
					GetBookshelfByUserIDAndBookID(ctx, "user01", bookshelf1.BookID).
					Return(bookshelf1, nil)
				mocks.BookRepository.EXPECT().
					GetReviewByUserIDAndBookID(ctx, "user01", bookshelf1.ReviewID).
					Return(review1, nil)
			},
			args: args{
				userID: "user01",
				bookID: 1,
			},
			want: want{
				bookshelf: bookshelf1,
				err:       nil,
			},
		},
		{
			name: "success: bookshelf is not found",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.BookRepository.EXPECT().
					Get(ctx, bookshelf1.BookID).
					Return(book1, nil)
				mocks.BookRepository.EXPECT().
					GetBookshelfByUserIDAndBookID(ctx, "user01", bookshelf1.BookID).
					Return(nil, exception.NotFound.New(test.ErrMock))
				mocks.BookRepository.EXPECT().
					GetReviewByUserIDAndBookID(ctx, "user01", bookshelf1.ReviewID).
					Return(review1, nil)
			},
			args: args{
				userID: "user01",
				bookID: 1,
			},
			want: want{
				bookshelf: &book.Bookshelf{Book: book1, Review: review1},
				err:       nil,
			},
		},
		{
			name: "success: review is not found",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.BookRepository.EXPECT().
					Get(ctx, bookshelf1.BookID).
					Return(book1, nil)
				mocks.BookRepository.EXPECT().
					GetBookshelfByUserIDAndBookID(ctx, "user01", bookshelf1.BookID).
					Return(bookshelf1, nil)
				mocks.BookRepository.EXPECT().
					GetReviewByUserIDAndBookID(ctx, "user01", bookshelf1.ReviewID).
					Return(nil, exception.NotFound.New(test.ErrMock))
			},
			args: args{
				userID: "user01",
				bookID: 1,
			},
			want: want{
				bookshelf: bookshelf1,
				err:       nil,
			},
		},
		{
			name: "failed: not found in get book",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.BookRepository.EXPECT().
					Get(ctx, bookshelf1.BookID).
					Return(nil, exception.NotFound.New(test.ErrMock))
			},
			args: args{
				userID: "user01",
				bookID: 1,
			},
			want: want{
				bookshelf: nil,
				err:       exception.NotFound.New(test.ErrMock),
			},
		},
		{
			name: "failed: internal error in get bookshelf",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.BookRepository.EXPECT().
					Get(ctx, bookshelf1.BookID).
					Return(book1, nil)
				mocks.BookRepository.EXPECT().
					GetBookshelfByUserIDAndBookID(ctx, "user01", bookshelf1.BookID).
					Return(nil, exception.ErrorInDatastore.New(test.ErrMock))
			},
			args: args{
				userID: "user01",
				bookID: 1,
			},
			want: want{
				bookshelf: nil,
				err:       exception.ErrorInDatastore.New(test.ErrMock),
			},
		},
		{
			name: "failed: internal error in get review",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.BookRepository.EXPECT().
					Get(ctx, bookshelf1.BookID).
					Return(book1, nil)
				mocks.BookRepository.EXPECT().
					GetBookshelfByUserIDAndBookID(ctx, "user01", bookshelf1.BookID).
					Return(bookshelf1, nil)
				mocks.BookRepository.EXPECT().
					GetReviewByUserIDAndBookID(ctx, "user01", bookshelf1.ReviewID).
					Return(nil, exception.ErrorInDatastore.New(test.ErrMock))
			},
			args: args{
				userID: "user01",
				bookID: 1,
			},
			want: want{
				bookshelf: nil,
				err:       exception.ErrorInDatastore.New(test.ErrMock),
			},
		},
	}

	for _, tt := range testCases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			mocks := test.NewMocks(ctrl)
			tt.setup(ctx, t, mocks)
			target := NewBookApplication(
				mocks.BookDomainValidation,
				mocks.BookRepository,
			)

			bs, err := target.GetBookshelfByUserIDAndBookIDWithRelated(ctx, tt.args.userID, tt.args.bookID)
			if tt.want.err != nil {
				require.Equal(t, tt.want.err.Error(), err.Error())
				require.Equal(t, tt.want.bookshelf, bs)
				return
			}
			require.NoError(t, err)
			require.Equal(t, tt.want.bookshelf, bs)
			require.Equal(t, tt.want.bookshelf.Book, bs.Book)
			require.Equal(t, tt.want.bookshelf.Review, bs.Review)
		})
	}
}

func TestBookApplication_GetReview(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	book1 := testBook(1)
	review1 := testReview(1, book1.ID, "user01")

	type args struct {
		reviewID int
	}
	type want struct {
		review *book.Review
		err    error
	}
	testCases := []struct {
		name  string
		setup func(context.Context, *testing.T, *test.Mocks)
		args  args
		want  want
	}{
		{
			name: "success",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.BookRepository.EXPECT().
					GetReview(ctx, review1.ID).
					Return(review1, nil)
			},
			args: args{
				reviewID: 1,
			},
			want: want{
				review: review1,
				err:    nil,
			},
		},
	}

	for _, tt := range testCases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			mocks := test.NewMocks(ctrl)
			tt.setup(ctx, t, mocks)
			target := NewBookApplication(
				mocks.BookDomainValidation,
				mocks.BookRepository,
			)

			bs, err := target.GetReview(ctx, tt.args.reviewID)
			if tt.want.err != nil {
				require.Equal(t, tt.want.err.Error(), err.Error())
				require.Equal(t, tt.want.review, bs)
				return
			}
			require.NoError(t, err)
			require.Equal(t, tt.want.review, bs)
		})
	}
}

func TestBookApplication_GetReviewByUserIDAndBookID(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	book1 := testBook(1)
	review1 := testReview(1, book1.ID, "user01")

	type args struct {
		userID string
		bookID int
	}
	type want struct {
		review *book.Review
		err    error
	}
	testCases := []struct {
		name  string
		setup func(context.Context, *testing.T, *test.Mocks)
		args  args
		want  want
	}{
		{
			name: "success",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.BookRepository.EXPECT().
					GetReviewByUserIDAndBookID(ctx, "user01", book1.ID).
					Return(review1, nil)
			},
			args: args{
				userID: "user01",
				bookID: 1,
			},
			want: want{
				review: review1,
				err:    nil,
			},
		},
	}

	for _, tt := range testCases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			mocks := test.NewMocks(ctrl)
			tt.setup(ctx, t, mocks)
			target := NewBookApplication(
				mocks.BookDomainValidation,
				mocks.BookRepository,
			)

			r, err := target.GetReviewByUserIDAndBookID(ctx, tt.args.userID, tt.args.bookID)
			if tt.want.err != nil {
				require.Equal(t, tt.want.err.Error(), err.Error())
				require.Equal(t, tt.want.review, r)
				return
			}
			require.NoError(t, err)
			require.Equal(t, tt.want.review, r)
		})
	}
}

func TestBookApplication_Create(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type args struct {
		book *book.Book
	}
	type want struct {
		err error
	}
	testCases := []struct {
		name  string
		setup func(context.Context, *testing.T, *test.Mocks)
		args  args
		want  want
	}{
		{
			name: "success",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.BookDomainValidation.EXPECT().
					Author(ctx, gomock.Any()).
					AnyTimes().Return(nil)
				mocks.BookDomainValidation.EXPECT().
					Book(ctx, gomock.Any()).
					Return(nil)
				mocks.BookRepository.EXPECT().
					Create(ctx, gomock.Any()).
					Return(nil)
			},
			args: args{
				book: &book.Book{
					Title: "小説　ちはやふる　上の句",
					Isbn:  "9784062938426",
					Authors: []*book.Author{
						{
							Name: "有沢 ゆう希",
						},
					},
				},
			},
			want: want{
				err: nil,
			},
		},
		{
			name: "success: authors length is 0",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.BookDomainValidation.EXPECT().
					Book(ctx, gomock.Any()).
					Return(nil)
				mocks.BookRepository.EXPECT().
					Create(ctx, gomock.Any()).
					Return(nil)
			},
			args: args{
				book: &book.Book{
					Title: "小説　ちはやふる　上の句",
					Isbn:  "9784062938426",
				},
			},
			want: want{
				err: nil,
			},
		},
		{
			name: "failed: invalid validation in authors",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.BookDomainValidation.EXPECT().
					Author(ctx, gomock.Any()).
					AnyTimes().Return(exception.InvalidDomainValidation.New(test.ErrMock))
			},
			args: args{
				book: &book.Book{
					Title: "小説　ちはやふる　上の句",
					Isbn:  "9784062938426",
					Authors: []*book.Author{
						{
							Name: "有沢 ゆう希",
						},
					},
				},
			},
			want: want{
				err: exception.InvalidDomainValidation.New(test.ErrMock),
			},
		},
		{
			name: "failed: invalid validation in book",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.BookDomainValidation.EXPECT().
					Author(ctx, gomock.Any()).
					AnyTimes().Return(nil)
				mocks.BookDomainValidation.EXPECT().
					Book(ctx, gomock.Any()).
					Return(exception.InvalidDomainValidation.New(test.ErrMock))
			},
			args: args{
				book: &book.Book{
					Title: "小説　ちはやふる　上の句",
					Isbn:  "9784062938426",
					Authors: []*book.Author{
						{
							Name: "有沢 ゆう希",
						},
					},
				},
			},
			want: want{
				err: exception.InvalidDomainValidation.New(test.ErrMock),
			},
		},
		{
			name: "failed: internal error in create",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.BookDomainValidation.EXPECT().
					Author(ctx, gomock.Any()).
					AnyTimes().Return(nil)
				mocks.BookDomainValidation.EXPECT().
					Book(ctx, gomock.Any()).
					Return(nil)
				mocks.BookRepository.EXPECT().
					Create(ctx, gomock.Any()).
					Return(exception.ErrorInDatastore.New(test.ErrMock))
			},
			args: args{
				book: &book.Book{
					Title: "小説　ちはやふる　上の句",
					Isbn:  "9784062938426",
					Authors: []*book.Author{
						{
							Name: "有沢 ゆう希",
						},
					},
				},
			},
			want: want{
				err: exception.ErrorInDatastore.New(test.ErrMock),
			},
		},
	}

	for _, tt := range testCases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			mocks := test.NewMocks(ctrl)
			tt.setup(ctx, t, mocks)
			target := NewBookApplication(
				mocks.BookDomainValidation,
				mocks.BookRepository,
			)

			err := target.Create(ctx, tt.args.book)
			if tt.want.err != nil {
				require.Equal(t, tt.want.err.Error(), err.Error())
				return
			}
			require.NoError(t, err)
			require.NotZero(t, tt.args.book.CreatedAt)
			require.NotZero(t, tt.args.book.UpdatedAt)
			for _, a := range tt.args.book.Authors {
				require.NotZero(t, a.CreatedAt)
				require.NotZero(t, a.UpdatedAt)
			}
		})
	}
}

func TestBookApplication_CreateBookshelf(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type args struct {
		bookshelf *book.Bookshelf
	}
	type want struct {
		err error
	}
	testCases := []struct {
		name  string
		setup func(context.Context, *testing.T, *test.Mocks)
		args  args
		want  want
	}{
		{
			name: "success",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.BookDomainValidation.EXPECT().
					Review(ctx, gomock.Any()).
					AnyTimes().Return(nil)
				mocks.BookDomainValidation.EXPECT().
					Bookshelf(ctx, gomock.Any()).
					Return(nil)
				mocks.BookRepository.EXPECT().
					CreateBookshelf(ctx, gomock.Any()).
					Return(nil)
			},
			args: args{
				bookshelf: &book.Bookshelf{
					Status: book.ReadingStatus,
					Review: &book.Review{
						Impression: "テストです。",
					},
				},
			},
			want: want{
				err: nil,
			},
		},
		{
			name: "failed: invalid validation in review",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.BookDomainValidation.EXPECT().
					Review(ctx, gomock.Any()).
					AnyTimes().Return(exception.InvalidDomainValidation.New(test.ErrMock))
			},
			args: args{
				bookshelf: &book.Bookshelf{
					Status: book.ReadingStatus,
					Review: &book.Review{
						Impression: "テストです。",
					},
				},
			},
			want: want{
				err: exception.InvalidDomainValidation.New(test.ErrMock),
			},
		},
		{
			name: "failed: invalid validation in bookshelf",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.BookDomainValidation.EXPECT().
					Review(ctx, gomock.Any()).
					AnyTimes().Return(nil)
				mocks.BookDomainValidation.EXPECT().
					Bookshelf(ctx, gomock.Any()).
					Return(exception.InvalidDomainValidation.New(test.ErrMock))
			},
			args: args{
				bookshelf: &book.Bookshelf{
					Status: book.ReadingStatus,
					Review: &book.Review{
						Impression: "テストです。",
					},
				},
			},
			want: want{
				err: exception.InvalidDomainValidation.New(test.ErrMock),
			},
		},
		{
			name: "failed: internal error in create",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.BookDomainValidation.EXPECT().
					Review(ctx, gomock.Any()).
					AnyTimes().Return(nil)
				mocks.BookDomainValidation.EXPECT().
					Bookshelf(ctx, gomock.Any()).
					Return(nil)
				mocks.BookRepository.EXPECT().
					CreateBookshelf(ctx, gomock.Any()).
					Return(exception.ErrorInDatastore.New(test.ErrMock))
			},
			args: args{
				bookshelf: &book.Bookshelf{
					Status: book.ReadingStatus,
					Review: &book.Review{
						Impression: "テストです。",
					},
				},
			},
			want: want{
				err: exception.ErrorInDatastore.New(test.ErrMock),
			},
		},
	}

	for _, tt := range testCases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			mocks := test.NewMocks(ctrl)
			tt.setup(ctx, t, mocks)
			target := NewBookApplication(
				mocks.BookDomainValidation,
				mocks.BookRepository,
			)

			err := target.CreateBookshelf(ctx, tt.args.bookshelf)
			if tt.want.err != nil {
				require.Equal(t, tt.want.err.Error(), err.Error())
				return
			}
			require.NoError(t, err)
			require.NotZero(t, tt.args.bookshelf.CreatedAt)
			require.NotZero(t, tt.args.bookshelf.UpdatedAt)
			if tt.args.bookshelf.Review != nil {
				require.NotZero(t, tt.args.bookshelf.Review.CreatedAt)
				require.NotZero(t, tt.args.bookshelf.Review.UpdatedAt)
			}
		})
	}
}

func TestBookApplication_Update(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type args struct {
		book *book.Book
	}
	type want struct {
		err error
	}
	testCases := []struct {
		name  string
		setup func(context.Context, *testing.T, *test.Mocks)
		args  args
		want  want
	}{
		{
			name: "success",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.BookDomainValidation.EXPECT().
					Author(ctx, gomock.Any()).
					AnyTimes().Return(nil)
				mocks.BookDomainValidation.EXPECT().
					Book(ctx, gomock.Any()).
					Return(nil)
				mocks.BookRepository.EXPECT().
					Update(ctx, gomock.Any()).
					Return(nil)
			},
			args: args{
				book: &book.Book{
					ID:    1,
					Title: "小説　ちはやふる　上の句",
					Isbn:  "9784062938426",
					Authors: []*book.Author{
						{
							Name: "有沢 ゆう希",
						},
					},
				},
			},
			want: want{
				err: nil,
			},
		},
		{
			name: "success: authors length is 0",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.BookDomainValidation.EXPECT().
					Book(ctx, gomock.Any()).
					Return(nil)
				mocks.BookRepository.EXPECT().
					Update(ctx, gomock.Any()).
					Return(nil)
			},
			args: args{
				book: &book.Book{
					ID:    1,
					Title: "小説　ちはやふる　上の句",
					Isbn:  "9784062938426",
				},
			},
			want: want{
				err: nil,
			},
		},
		{
			name: "failed: invalid validation in authors",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.BookDomainValidation.EXPECT().
					Author(ctx, gomock.Any()).
					AnyTimes().Return(exception.InvalidDomainValidation.New(test.ErrMock))
			},
			args: args{
				book: &book.Book{
					ID:    1,
					Title: "小説　ちはやふる　上の句",
					Isbn:  "9784062938426",
					Authors: []*book.Author{
						{
							Name: "有沢 ゆう希",
						},
					},
				},
			},
			want: want{
				err: exception.InvalidDomainValidation.New(test.ErrMock),
			},
		},
		{
			name: "failed: invalid validation in book",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.BookDomainValidation.EXPECT().
					Author(ctx, gomock.Any()).
					AnyTimes().Return(nil)
				mocks.BookDomainValidation.EXPECT().
					Book(ctx, gomock.Any()).
					Return(exception.InvalidDomainValidation.New(test.ErrMock))
			},
			args: args{
				book: &book.Book{
					ID:    1,
					Title: "小説　ちはやふる　上の句",
					Isbn:  "9784062938426",
					Authors: []*book.Author{
						{
							Name: "有沢 ゆう希",
						},
					},
				},
			},
			want: want{
				err: exception.InvalidDomainValidation.New(test.ErrMock),
			},
		},
		{
			name: "failed: internal error in create",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.BookDomainValidation.EXPECT().
					Author(ctx, gomock.Any()).
					AnyTimes().Return(nil)
				mocks.BookDomainValidation.EXPECT().
					Book(ctx, gomock.Any()).
					Return(nil)
				mocks.BookRepository.EXPECT().
					Update(ctx, gomock.Any()).
					Return(exception.ErrorInDatastore.New(test.ErrMock))
			},
			args: args{
				book: &book.Book{
					ID:    1,
					Title: "小説　ちはやふる　上の句",
					Isbn:  "9784062938426",
					Authors: []*book.Author{
						{
							Name: "有沢 ゆう希",
						},
					},
				},
			},
			want: want{
				err: exception.ErrorInDatastore.New(test.ErrMock),
			},
		},
	}

	for _, tt := range testCases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			mocks := test.NewMocks(ctrl)
			tt.setup(ctx, t, mocks)
			target := NewBookApplication(
				mocks.BookDomainValidation,
				mocks.BookRepository,
			)

			err := target.Update(ctx, tt.args.book)
			if tt.want.err != nil {
				require.Equal(t, tt.want.err.Error(), err.Error())
				return
			}
			require.NoError(t, err)
			require.NotZero(t, tt.args.book.UpdatedAt)
			for _, a := range tt.args.book.Authors {
				require.NotZero(t, a.UpdatedAt)
			}
		})
	}
}

func TestBookApplication_UpdateBookshelf(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type args struct {
		bookshelf *book.Bookshelf
	}
	type want struct {
		err error
	}
	testCases := []struct {
		name  string
		setup func(context.Context, *testing.T, *test.Mocks)
		args  args
		want  want
	}{
		{
			name: "success",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.BookDomainValidation.EXPECT().
					Review(ctx, gomock.Any()).
					AnyTimes().Return(nil)
				mocks.BookDomainValidation.EXPECT().
					Bookshelf(ctx, gomock.Any()).
					Return(nil)
				mocks.BookRepository.EXPECT().
					UpdateBookshelf(ctx, gomock.Any()).
					Return(nil)
			},
			args: args{
				bookshelf: &book.Bookshelf{
					Status: book.ReadingStatus,
					Review: &book.Review{
						Impression: "テストです。",
					},
				},
			},
			want: want{
				err: nil,
			},
		},
		{
			name: "failed: invalid validation in review",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.BookDomainValidation.EXPECT().
					Review(ctx, gomock.Any()).
					AnyTimes().Return(exception.InvalidDomainValidation.New(test.ErrMock))
			},
			args: args{
				bookshelf: &book.Bookshelf{
					Status: book.ReadingStatus,
					Review: &book.Review{
						Impression: "テストです。",
					},
				},
			},
			want: want{
				err: exception.InvalidDomainValidation.New(test.ErrMock),
			},
		},
		{
			name: "failed: invalid validation in bookshelf",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.BookDomainValidation.EXPECT().
					Review(ctx, gomock.Any()).
					AnyTimes().Return(nil)
				mocks.BookDomainValidation.EXPECT().
					Bookshelf(ctx, gomock.Any()).
					Return(exception.InvalidDomainValidation.New(test.ErrMock))
			},
			args: args{
				bookshelf: &book.Bookshelf{
					Status: book.ReadingStatus,
					Review: &book.Review{
						Impression: "テストです。",
					},
				},
			},
			want: want{
				err: exception.InvalidDomainValidation.New(test.ErrMock),
			},
		},
		{
			name: "failed: internal error in create",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.BookDomainValidation.EXPECT().
					Review(ctx, gomock.Any()).
					AnyTimes().Return(nil)
				mocks.BookDomainValidation.EXPECT().
					Bookshelf(ctx, gomock.Any()).
					Return(nil)
				mocks.BookRepository.EXPECT().
					UpdateBookshelf(ctx, gomock.Any()).
					Return(exception.ErrorInDatastore.New(test.ErrMock))
			},
			args: args{
				bookshelf: &book.Bookshelf{
					Status: book.ReadingStatus,
					Review: &book.Review{
						Impression: "テストです。",
					},
				},
			},
			want: want{
				err: exception.ErrorInDatastore.New(test.ErrMock),
			},
		},
	}

	for _, tt := range testCases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			mocks := test.NewMocks(ctrl)
			tt.setup(ctx, t, mocks)
			target := NewBookApplication(
				mocks.BookDomainValidation,
				mocks.BookRepository,
			)

			err := target.UpdateBookshelf(ctx, tt.args.bookshelf)
			if tt.want.err != nil {
				require.Equal(t, tt.want.err.Error(), err.Error())
				return
			}
			require.NoError(t, err)
			require.NotZero(t, tt.args.bookshelf.UpdatedAt)
			if tt.args.bookshelf.Review != nil {
				require.NotZero(t, tt.args.bookshelf.Review.UpdatedAt)
			}
		})
	}
}

func TestBookApplication_CreateOrUpdateBookshelf(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type args struct {
		bookshelf *book.Bookshelf
	}
	type want struct {
		err error
	}
	testCases := []struct {
		name  string
		setup func(context.Context, *testing.T, *test.Mocks)
		args  args
		want  want
	}{
		{
			name: "success: create",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.BookDomainValidation.EXPECT().
					Review(ctx, gomock.Any()).
					AnyTimes().Return(nil)
				mocks.BookDomainValidation.EXPECT().
					Bookshelf(ctx, gomock.Any()).
					Return(nil)
				mocks.BookRepository.EXPECT().
					CreateBookshelf(ctx, gomock.Any()).
					Return(nil)
			},
			args: args{
				bookshelf: &book.Bookshelf{
					Status: book.ReadingStatus,
					Review: &book.Review{
						Impression: "テストです。",
					},
				},
			},
			want: want{
				err: nil,
			},
		},
		{
			name: "success: update",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.BookDomainValidation.EXPECT().
					Review(ctx, gomock.Any()).
					AnyTimes().Return(nil)
				mocks.BookDomainValidation.EXPECT().
					Bookshelf(ctx, gomock.Any()).
					Return(nil)
				mocks.BookRepository.EXPECT().
					UpdateBookshelf(ctx, gomock.Any()).
					Return(nil)
			},
			args: args{
				bookshelf: &book.Bookshelf{
					ID:     1,
					Status: book.ReadingStatus,
					Review: &book.Review{
						Impression: "テストです。",
					},
				},
			},
			want: want{
				err: nil,
			},
		},
	}

	for _, tt := range testCases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			mocks := test.NewMocks(ctrl)
			tt.setup(ctx, t, mocks)
			target := NewBookApplication(
				mocks.BookDomainValidation,
				mocks.BookRepository,
			)

			err := target.CreateOrUpdateBookshelf(ctx, tt.args.bookshelf)
			if tt.want.err != nil {
				require.Equal(t, tt.want.err.Error(), err.Error())
				return
			}
			require.NoError(t, err)
			require.NotZero(t, tt.args.bookshelf.UpdatedAt)
		})
	}
}

func TestBookApplication_MultipleCreate(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type args struct {
		books book.Books
	}
	type want struct {
		err error
	}
	testCases := []struct {
		name  string
		setup func(context.Context, *testing.T, *test.Mocks)
		args  args
		want  want
	}{
		{
			name: "success",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.BookDomainValidation.EXPECT().
					Author(ctx, gomock.Any()).
					AnyTimes().Return(nil)
				mocks.BookDomainValidation.EXPECT().
					Book(ctx, gomock.Any()).
					AnyTimes().Return(nil)
				mocks.BookRepository.EXPECT().
					MultipleCreate(ctx, gomock.Any()).
					Return(nil)
			},
			args: args{
				books: []*book.Book{
					{
						Title: "小説　ちはやふる　上の句",
						Isbn:  "9784062938426",
						Authors: []*book.Author{
							{
								Name: "有沢 ゆう希",
							},
						},
					},
					{
						Title: "小説　ちはやふる　下の句",
						Isbn:  "9784062938471",
						Authors: []*book.Author{
							{
								Name: "有沢 ゆう希",
							},
						},
					},
				},
			},
			want: want{
				err: nil,
			},
		},
		{
			name: "success: authors length is 0",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.BookDomainValidation.EXPECT().
					Book(ctx, gomock.Any()).
					AnyTimes().Return(nil)
				mocks.BookRepository.EXPECT().
					MultipleCreate(ctx, gomock.Any()).
					Return(nil)
			},
			args: args{
				books: []*book.Book{
					{
						Title: "小説　ちはやふる　上の句",
						Isbn:  "9784062938426",
					},
					{
						Title: "小説　ちはやふる　下の句",
						Isbn:  "9784062938471",
					},
				},
			},
			want: want{
				err: nil,
			},
		},
		{
			name: "failed: invalid validation in authors",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.BookDomainValidation.EXPECT().
					Author(ctx, gomock.Any()).
					AnyTimes().Return(exception.InvalidDomainValidation.New(test.ErrMock))
			},
			args: args{
				books: []*book.Book{
					{
						Title: "小説　ちはやふる　上の句",
						Isbn:  "9784062938426",
						Authors: []*book.Author{
							{
								Name: "有沢 ゆう希",
							},
						},
					},
					{
						Title: "小説　ちはやふる　下の句",
						Isbn:  "9784062938471",
						Authors: []*book.Author{
							{
								Name: "有沢 ゆう希",
							},
						},
					},
				},
			},
			want: want{
				err: exception.InvalidDomainValidation.New(test.ErrMock),
			},
		},
		{
			name: "failed: invalid validation in book",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.BookDomainValidation.EXPECT().
					Author(ctx, gomock.Any()).
					AnyTimes().Return(nil)
				mocks.BookDomainValidation.EXPECT().
					Book(ctx, gomock.Any()).
					AnyTimes().Return(exception.InvalidDomainValidation.New(test.ErrMock))
			},
			args: args{
				books: []*book.Book{
					{
						Title: "小説　ちはやふる　上の句",
						Isbn:  "9784062938426",
						Authors: []*book.Author{
							{
								Name: "有沢 ゆう希",
							},
						},
					},
					{
						Title: "小説　ちはやふる　下の句",
						Isbn:  "9784062938471",
						Authors: []*book.Author{
							{
								Name: "有沢 ゆう希",
							},
						},
					},
				},
			},
			want: want{
				err: exception.InvalidDomainValidation.New(test.ErrMock),
			},
		},
		{
			name: "failed: internal error in multiple create",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.BookDomainValidation.EXPECT().
					Author(ctx, gomock.Any()).
					AnyTimes().Return(nil)
				mocks.BookDomainValidation.EXPECT().
					Book(ctx, gomock.Any()).
					AnyTimes().Return(nil)
				mocks.BookRepository.EXPECT().
					MultipleCreate(ctx, gomock.Any()).
					Return(exception.ErrorInDatastore.New(test.ErrMock))
			},
			args: args{
				books: []*book.Book{
					{
						Title: "小説　ちはやふる　上の句",
						Isbn:  "9784062938426",
						Authors: []*book.Author{
							{
								Name: "有沢 ゆう希",
							},
						},
					},
					{
						Title: "小説　ちはやふる　下の句",
						Isbn:  "9784062938471",
						Authors: []*book.Author{
							{
								Name: "有沢 ゆう希",
							},
						},
					},
				},
			},
			want: want{
				err: exception.ErrorInDatastore.New(test.ErrMock),
			},
		},
	}

	for _, tt := range testCases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			mocks := test.NewMocks(ctrl)
			tt.setup(ctx, t, mocks)
			target := NewBookApplication(
				mocks.BookDomainValidation,
				mocks.BookRepository,
			)

			err := target.MultipleCreate(ctx, tt.args.books)
			if tt.want.err != nil {
				require.Equal(t, tt.want.err.Error(), err.Error())
				return
			}
			require.NoError(t, err)
			for _, b := range tt.args.books {
				require.NotZero(t, b.CreatedAt)
				require.NotZero(t, b.UpdatedAt)
				for _, a := range b.Authors {
					require.NotZero(t, a.CreatedAt)
					require.NotZero(t, a.UpdatedAt)
				}
			}
		})
	}
}

func TestBookApplication_MultipleUpdate(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type args struct {
		books book.Books
	}
	type want struct {
		err error
	}
	testCases := []struct {
		name  string
		setup func(context.Context, *testing.T, *test.Mocks)
		args  args
		want  want
	}{
		{
			name: "success",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.BookDomainValidation.EXPECT().
					Author(ctx, gomock.Any()).
					AnyTimes().Return(nil)
				mocks.BookDomainValidation.EXPECT().
					Book(ctx, gomock.Any()).
					AnyTimes().Return(nil)
				mocks.BookRepository.EXPECT().
					MultipleUpdate(ctx, gomock.Any()).
					Return(nil)
			},
			args: args{
				books: []*book.Book{
					{
						ID:    1,
						Title: "小説　ちはやふる　上の句",
						Isbn:  "9784062938426",
						Authors: []*book.Author{
							{
								Name: "有沢 ゆう希",
							},
						},
					},
					{
						ID:    2,
						Title: "小説　ちはやふる　下の句",
						Isbn:  "9784062938471",
						Authors: []*book.Author{
							{
								Name: "有沢 ゆう希",
							},
						},
					},
				},
			},
			want: want{
				err: nil,
			},
		},
		{
			name: "success: authors length is 0",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.BookDomainValidation.EXPECT().
					Book(ctx, gomock.Any()).
					AnyTimes().Return(nil)
				mocks.BookRepository.EXPECT().
					MultipleUpdate(ctx, gomock.Any()).
					Return(nil)
			},
			args: args{
				books: []*book.Book{
					{
						ID:    1,
						Title: "小説　ちはやふる　上の句",
						Isbn:  "9784062938426",
					},
					{
						ID:    2,
						Title: "小説　ちはやふる　下の句",
						Isbn:  "9784062938471",
					},
				},
			},
			want: want{
				err: nil,
			},
		},
		{
			name: "failed: invalid validation in authors",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.BookDomainValidation.EXPECT().
					Author(ctx, gomock.Any()).
					AnyTimes().Return(exception.InvalidDomainValidation.New(test.ErrMock))
			},
			args: args{
				books: []*book.Book{
					{
						ID:    1,
						Title: "小説　ちはやふる　上の句",
						Isbn:  "9784062938426",
						Authors: []*book.Author{
							{
								Name: "有沢 ゆう希",
							},
						},
					},
					{
						ID:    2,
						Title: "小説　ちはやふる　下の句",
						Isbn:  "9784062938471",
						Authors: []*book.Author{
							{
								Name: "有沢 ゆう希",
							},
						},
					},
				},
			},
			want: want{
				err: exception.InvalidDomainValidation.New(test.ErrMock),
			},
		},
		{
			name: "failed: invalid validation in book",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.BookDomainValidation.EXPECT().
					Author(ctx, gomock.Any()).
					AnyTimes().Return(nil)
				mocks.BookDomainValidation.EXPECT().
					Book(ctx, gomock.Any()).
					AnyTimes().Return(exception.InvalidDomainValidation.New(test.ErrMock))
			},
			args: args{
				books: []*book.Book{
					{
						ID:    1,
						Title: "小説　ちはやふる　上の句",
						Isbn:  "9784062938426",
						Authors: []*book.Author{
							{
								Name: "有沢 ゆう希",
							},
						},
					},
					{
						ID:    2,
						Title: "小説　ちはやふる　下の句",
						Isbn:  "9784062938471",
						Authors: []*book.Author{
							{
								Name: "有沢 ゆう希",
							},
						},
					},
				},
			},
			want: want{
				err: exception.InvalidDomainValidation.New(test.ErrMock),
			},
		},
		{
			name: "failed: internal error in multiple update",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.BookDomainValidation.EXPECT().
					Author(ctx, gomock.Any()).
					AnyTimes().Return(nil)
				mocks.BookDomainValidation.EXPECT().
					Book(ctx, gomock.Any()).
					AnyTimes().Return(nil)
				mocks.BookRepository.EXPECT().
					MultipleUpdate(ctx, gomock.Any()).
					Return(exception.ErrorInDatastore.New(test.ErrMock))
			},
			args: args{
				books: []*book.Book{
					{
						ID:    1,
						Title: "小説　ちはやふる　上の句",
						Isbn:  "9784062938426",
						Authors: []*book.Author{
							{
								Name: "有沢 ゆう希",
							},
						},
					},
					{
						ID:    2,
						Title: "小説　ちはやふる　下の句",
						Isbn:  "9784062938471",
						Authors: []*book.Author{
							{
								Name: "有沢 ゆう希",
							},
						},
					},
				},
			},
			want: want{
				err: exception.ErrorInDatastore.New(test.ErrMock),
			},
		},
	}

	for _, tt := range testCases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			mocks := test.NewMocks(ctrl)
			tt.setup(ctx, t, mocks)
			target := NewBookApplication(
				mocks.BookDomainValidation,
				mocks.BookRepository,
			)

			err := target.MultipleUpdate(ctx, tt.args.books)
			if tt.want.err != nil {
				require.Equal(t, tt.want.err.Error(), err.Error())
				return
			}
			require.NoError(t, err)
			for _, b := range tt.args.books {
				require.NotZero(t, b.UpdatedAt)
				for _, a := range b.Authors {
					require.NotZero(t, a.UpdatedAt)
				}
			}
		})
	}
}
func TestBookApplication_Delete(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type args struct {
		book *book.Book
	}
	type want struct {
		err error
	}
	testCases := []struct {
		name  string
		setup func(context.Context, *testing.T, *test.Mocks)
		args  args
		want  want
	}{
		{
			name: "success",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.BookRepository.EXPECT().
					Delete(ctx, 1).
					Return(nil)
			},
			args: args{
				book: &book.Book{
					ID: 1,
				},
			},
			want: want{
				err: nil,
			},
		},
	}

	for _, tt := range testCases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			mocks := test.NewMocks(ctrl)
			tt.setup(ctx, t, mocks)
			target := NewBookApplication(
				mocks.BookDomainValidation,
				mocks.BookRepository,
			)

			err := target.Delete(ctx, tt.args.book)
			if tt.want.err != nil {
				require.Equal(t, tt.want.err.Error(), err.Error())
				return
			}
			require.NoError(t, err)
		})
	}
}

func TestBookApplication_DeleteBookshelf(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type args struct {
		bookshelf *book.Bookshelf
	}
	type want struct {
		err error
	}
	testCases := []struct {
		name  string
		setup func(context.Context, *testing.T, *test.Mocks)
		args  args
		want  want
	}{
		{
			name: "success",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.BookRepository.EXPECT().
					DeleteBookshelf(ctx, 1).
					Return(nil)
			},
			args: args{
				bookshelf: &book.Bookshelf{
					ID: 1,
				},
			},
			want: want{
				err: nil,
			},
		},
	}

	for _, tt := range testCases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			mocks := test.NewMocks(ctrl)
			tt.setup(ctx, t, mocks)
			target := NewBookApplication(
				mocks.BookDomainValidation,
				mocks.BookRepository,
			)

			err := target.DeleteBookshelf(ctx, tt.args.bookshelf)
			if tt.want.err != nil {
				require.Equal(t, tt.want.err.Error(), err.Error())
				return
			}
			require.NoError(t, err)
		})
	}
}

func testBook(id int) *book.Book {
	return &book.Book{
		ID:             id,
		Title:          "小説　ちはやふる　上の句",
		TitleKana:      "ショウセツ チハヤフルカミノク",
		Description:    "綾瀬千早は高校入学と同時に、競技かるた部を作ろうと奔走する。幼馴染の太一と仲間を集め、夏の全国大会に出場するためだ。強くなって、新と再会したい。幼い頃かるたを取り合った、新に寄せる千早の秘めた想いに気づきながらも、太一は千早を守り立てる。それぞれの青春を懸けた、一途な情熱の物語が幕開ける。",
		Isbn:           "9784062938426",
		Publisher:      "講談社",
		PublishedOn:    "2018-01-16",
		ThumbnailURL:   "https://thumbnail.image.rakuten.co.jp/@0_mall/book/cabinet/8426/9784062938426.jpg?_ex=120x120",
		RakutenURL:     "https://books.rakuten.co.jp/rb/15271426/",
		RakutenSize:    "コミック",
		RakutenGenreID: "001004008001/001004008003/001019001",
		CreatedAt:      test.TimeMock,
		UpdatedAt:      test.TimeMock,
		Authors: []*book.Author{
			{
				ID:        1,
				Name:      "有沢 ゆう希",
				NameKana:  "アリサワ ユウキ",
				CreatedAt: test.TimeMock,
				UpdatedAt: test.TimeMock,
			},
			{
				ID:        2,
				Name:      "末次 由紀",
				NameKana:  "スエツグ ユキ",
				CreatedAt: test.TimeMock,
				UpdatedAt: test.TimeMock,
			},
		},
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
