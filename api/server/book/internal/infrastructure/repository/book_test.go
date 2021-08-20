package repository

import (
	"context"
	"fmt"
	"testing"

	"github.com/calmato/gran-book/api/server/book/internal/domain/book"
	"github.com/calmato/gran-book/api/server/book/pkg/database"
	"github.com/calmato/gran-book/api/server/book/pkg/test"
	"github.com/golang/mock/gomock"
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
			require.Equal(t, tt.want.isErr, err != nil)
			for i := range tt.want.books {
				require.Contains(t, bs, tt.want.books[i])
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
			require.Equal(t, tt.want.isErr, err != nil)
			for i := range tt.want.bookshelves {
				fmt.Printf("debug: %v === %v", tt.want.bookshelves[i], bs[i])
				require.Contains(t, bs, tt.want.bookshelves[i])
			}
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
