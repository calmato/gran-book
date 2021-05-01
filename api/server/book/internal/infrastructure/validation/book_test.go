package validation

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/calmato/gran-book/api/server/book/internal/domain/book"
	"github.com/calmato/gran-book/api/server/book/lib/datetime"
	mock_book "github.com/calmato/gran-book/api/server/book/mock/domain/book"
	"github.com/golang/mock/gomock"
)

func TestBookService_Book(t *testing.T) {
	testCases := map[string]struct {
		Book     *book.Book
		Expected error
	}{
		"ok": {
			Book: &book.Book{
				ID:           0,
				Title:        "テスト書籍",
				TitleKana:    "てすとしょせき",
				Description:  "本の説明です",
				Isbn:         "1234567890123",
				Publisher:    "テスト著者",
				PublishedOn:  "2021年12月24日",
				ThumbnailURL: "",
				CreatedAt:    time.Time{},
				UpdatedAt:    time.Time{},
				Authors:      []*book.Author{},
				Reviews:      []*book.Review{},
			},
			Expected: nil,
		},
	}

	for result, tc := range testCases {
		t.Run(result, func(t *testing.T) {
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			bvm := mock_book.NewMockRepository(ctrl)
			bvm.EXPECT().GetIDByIsbn(ctx, tc.Book.Isbn).Return(tc.Book.ID, nil)

			t.Run(result, func(t *testing.T) {
				target := NewBookDomainValidation(bvm)

				err := target.Book(ctx, tc.Book)
				if !reflect.DeepEqual(err, tc.Expected) {
					t.Fatalf("want %#v, but %#v", tc.Expected, err)
					return
				}
			})
		})
	}
}

func TestBookService_Author(t *testing.T) {
	testCases := map[string]struct {
		Author   *book.Author
		Expected error
	}{
		"ok": {
			Author: &book.Author{
				ID:        0,
				Name:      "テスト著者",
				NameKana:  "てすとちょしゃ",
				CreatedAt: time.Time{},
				UpdatedAt: time.Time{},
			},
			Expected: nil,
		},
	}

	for result, tc := range testCases {
		t.Run(result, func(t *testing.T) {
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			bvm := mock_book.NewMockRepository(ctrl)

			t.Run(result, func(t *testing.T) {
				target := NewBookDomainValidation(bvm)

				err := target.Author(ctx, tc.Author)
				if !reflect.DeepEqual(err, tc.Expected) {
					t.Fatalf("want %#v, but %#v", tc.Expected, err)
					return
				}
			})
		})
	}
}

func TestBookService_Bookshelf(t *testing.T) {
	testCases := map[string]struct {
		Bookshelf *book.Bookshelf
		Expected  error
	}{
		"ok": {
			Bookshelf: &book.Bookshelf{
				ID:        0,
				BookID:    1,
				UserID:    "00000000-0000-0000-0000-000000000000",
				Status:    5,
				ReadOn:    datetime.StringToDate("2020-01-01"),
				CreatedAt: time.Time{},
				UpdatedAt: time.Time{},
			},
			Expected: nil,
		},
	}

	for result, tc := range testCases {
		t.Run(result, func(t *testing.T) {
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			bvm := mock_book.NewMockRepository(ctrl)
			bvm.EXPECT().
				GetBookshelfIDByUserIDAndBookID(ctx, tc.Bookshelf.UserID, tc.Bookshelf.BookID).
				Return(tc.Bookshelf.ID, nil)

			t.Run(result, func(t *testing.T) {
				target := NewBookDomainValidation(bvm)

				err := target.Bookshelf(ctx, tc.Bookshelf)
				if !reflect.DeepEqual(err, tc.Expected) {
					t.Fatalf("want %#v, but %#v", tc.Expected, err)
					return
				}
			})
		})
	}
}
