package service

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/calmato/gran-book/api/server/book/internal/domain/book"
	mock_book "github.com/calmato/gran-book/api/server/book/mock/domain/book"
	"github.com/golang/mock/gomock"
)

func TestBookService_Create(t *testing.T) {
	testCases := map[string]struct {
		Book     *book.Book
		Expected error
	}{
		"ok": {
			Book: &book.Book{
				ID:           0,
				PublisherID:  0,
				Title:        "テスト",
				Description:  "",
				Isbn:         "",
				ThumbnailURL: "",
				Version:      "0.0.1",
				PublishedOn:  time.Date(2020, 1, 1, 0, 0, 0, 0, time.Local),
				CreatedAt:    time.Time{},
				UpdatedAt:    time.Time{},
				Publisher: &book.Publisher{
					Name: "テスト出版社",
				},
				Bookshelfs: []*book.Bookshelf{},
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

			bvm := mock_book.NewMockValidation(ctrl)
			bvm.EXPECT().Book(ctx, tc.Book).Return(nil)
			bvm.EXPECT().Publisher(ctx, tc.Book.Publisher).Return(nil)
			for _, a := range tc.Book.Authors {
				bvm.EXPECT().Author(ctx, a).Return(nil)
			}
			for _, c := range tc.Book.Categories {
				bvm.EXPECT().Category(ctx, c).Return(nil)
			}

			brm := mock_book.NewMockRepository(ctrl)
			brm.EXPECT().Create(ctx, tc.Book).Return(tc.Expected)
			brm.EXPECT().CreatePublisher(ctx, tc.Book.Publisher).Return(tc.Expected)
			for _, a := range tc.Book.Authors {
				brm.EXPECT().CreateAuthor(ctx, a).Return(tc.Expected)
			}
			for _, c := range tc.Book.Categories {
				brm.EXPECT().CreateCategory(ctx, c).Return(tc.Expected)
			}

			t.Run(result, func(t *testing.T) {
				target := NewBookService(bvm, brm)

				err := target.Create(ctx, tc.Book)
				if !reflect.DeepEqual(err, tc.Expected) {
					t.Fatalf("want %#v, but %#v", tc.Expected, err)
					return
				}
			})
		})
	}
}

func TestBookService_CreateAuthor(t *testing.T) {
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

			bvm := mock_book.NewMockValidation(ctrl)
			bvm.EXPECT().Author(ctx, tc.Author).Return(nil)

			brm := mock_book.NewMockRepository(ctrl)
			brm.EXPECT().CreateAuthor(ctx, tc.Author).Return(tc.Expected)

			t.Run(result, func(t *testing.T) {
				target := NewBookService(bvm, brm)

				err := target.CreateAuthor(ctx, tc.Author)
				if !reflect.DeepEqual(err, tc.Expected) {
					t.Fatalf("want %#v, but %#v", tc.Expected, err)
					return
				}
			})
		})
	}
}

func TestBookService_CreateBookshelf(t *testing.T) {
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

			bvm := mock_book.NewMockValidation(ctrl)
			bvm.EXPECT().Bookshelf(ctx, tc.Bookshelf).Return(nil)

			brm := mock_book.NewMockRepository(ctrl)
			brm.EXPECT().CreateBookshelf(ctx, tc.Bookshelf).Return(tc.Expected)

			t.Run(result, func(t *testing.T) {
				target := NewBookService(bvm, brm)

				err := target.CreateBookshelf(ctx, tc.Bookshelf)
				if !reflect.DeepEqual(err, tc.Expected) {
					t.Fatalf("want %#v, but %#v", tc.Expected, err)
					return
				}
			})
		})
	}
}

func TestBookService_CreateCategory(t *testing.T) {
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

			bvm := mock_book.NewMockValidation(ctrl)
			bvm.EXPECT().Category(ctx, tc.Category).Return(nil)

			brm := mock_book.NewMockRepository(ctrl)
			brm.EXPECT().CreateCategory(ctx, tc.Category).Return(tc.Expected)

			t.Run(result, func(t *testing.T) {
				target := NewBookService(bvm, brm)

				err := target.CreateCategory(ctx, tc.Category)
				if !reflect.DeepEqual(err, tc.Expected) {
					t.Fatalf("want %#v, but %#v", tc.Expected, err)
					return
				}
			})
		})
	}
}

func TestBookService_CreatePublisher(t *testing.T) {
	testCases := map[string]struct {
		Publisher *book.Publisher
		Expected  error
	}{
		"ok": {
			Publisher: &book.Publisher{
				ID:        0,
				Name:      "テスト出版社",
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

			bvm := mock_book.NewMockValidation(ctrl)
			bvm.EXPECT().Publisher(ctx, tc.Publisher).Return(nil)

			brm := mock_book.NewMockRepository(ctrl)
			brm.EXPECT().CreatePublisher(ctx, tc.Publisher).Return(tc.Expected)

			t.Run(result, func(t *testing.T) {
				target := NewBookService(bvm, brm)

				err := target.CreatePublisher(ctx, tc.Publisher)
				if !reflect.DeepEqual(err, tc.Expected) {
					t.Fatalf("want %#v, but %#v", tc.Expected, err)
					return
				}
			})
		})
	}
}
