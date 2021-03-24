package validation

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/calmato/gran-book/api/server/book/internal/domain/book"
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
				Publisher:    "テスト出版社",
				Title:        "テスト",
				Description:  "",
				Isbn:         "",
				ThumbnailURL: "",
				Version:      "0.0.1",
				PublishedOn:  time.Date(2020, 1, 1, 0, 0, 0, 0, time.Local),
				CreatedAt:    time.Time{},
				UpdatedAt:    time.Time{},
				Bookshelves:  []*book.Bookshelf{},
				Authors: []*book.Author{{
					Name: "テスト著者",
				}},
				Categories: []*book.Category{{
					Name: "コミック",
				}},
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

			t.Run(result, func(t *testing.T) {
				target := NewBookDomainValidation()

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

			t.Run(result, func(t *testing.T) {
				target := NewBookDomainValidation()

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
				ID:         0,
				BookID:     1,
				UserID:     "00000000-0000-0000-0000-000000000000",
				Status:     5,
				Impression: "感想です",
				CreatedAt:  time.Time{},
				UpdatedAt:  time.Time{},
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

			t.Run(result, func(t *testing.T) {
				target := NewBookDomainValidation()

				err := target.Bookshelf(ctx, tc.Bookshelf)
				if !reflect.DeepEqual(err, tc.Expected) {
					t.Fatalf("want %#v, but %#v", tc.Expected, err)
					return
				}
			})
		})
	}
}

func TestBookService_Category(t *testing.T) {
	testCases := map[string]struct {
		Category *book.Category
		Expected error
	}{
		"ok": {
			Category: &book.Category{
				ID:        0,
				Name:      "テストカテゴリ",
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

			t.Run(result, func(t *testing.T) {
				target := NewBookDomainValidation()

				err := target.Category(ctx, tc.Category)
				if !reflect.DeepEqual(err, tc.Expected) {
					t.Fatalf("want %#v, but %#v", tc.Expected, err)
					return
				}
			})
		})
	}
}
