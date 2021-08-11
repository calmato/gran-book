package application

import (
	"context"
	"testing"
	"time"

	"github.com/calmato/gran-book/api/server/book/internal/domain/book"
	"github.com/calmato/gran-book/api/server/book/internal/domain/exception"
	"github.com/calmato/gran-book/api/server/book/pkg/database"
	"github.com/calmato/gran-book/api/server/book/pkg/test"
	"github.com/golang/mock/gomock"
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
		books []*book.Book
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
		bookshelves []*book.Bookshelf
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
					Count(ctx, gomock.Any()).
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
					Count(ctx, gomock.Any()).
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
		reviews []*book.Review
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
					Count(ctx, gomock.Any()).
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
					Count(ctx, gomock.Any()).
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
		reviews []*book.Review
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
					Count(ctx, gomock.Any()).
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
					Count(ctx, gomock.Any()).
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
		books []*book.Book
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

func testBook(id int) *book.Book {
	current := time.Now().Local()

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
		CreatedAt:      current,
		UpdatedAt:      current,
		Authors: []*book.Author{
			{
				ID:        1,
				Name:      "有沢 ゆう希",
				NameKana:  "アリサワ ユウキ",
				CreatedAt: current,
				UpdatedAt: current,
			},
			{
				ID:        2,
				Name:      "末次 由紀",
				NameKana:  "スエツグ ユキ",
				CreatedAt: current,
				UpdatedAt: current,
			},
		},
	}
}

func testBookshelf(id int, bookID int, userID string) *book.Bookshelf {
	current := time.Now().Local()

	return &book.Bookshelf{
		ID:        id,
		BookID:    bookID,
		UserID:    userID,
		ReviewID:  0,
		Status:    book.ReadingStatus,
		CreatedAt: current,
		UpdatedAt: current,
	}
}

func testReview(id int, bookID int, userID string) *book.Review {
	current := time.Now().Local()

	return &book.Review{
		ID:         id,
		BookID:     bookID,
		UserID:     userID,
		Score:      3,
		Impression: "テストレビューです",
		CreatedAt:  current,
		UpdatedAt:  current,
	}
}
