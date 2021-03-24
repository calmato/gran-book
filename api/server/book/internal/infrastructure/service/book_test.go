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

func TestBookService_ShowByIsbn(t *testing.T) {
	current := time.Now()

	testCases := map[string]struct {
		Isbn     string
		Expected struct {
			Book  *book.Book
			Error error
		}
	}{
		"ok": {
			Isbn: "978-1-234-56789-7",
			Expected: struct {
				Book  *book.Book
				Error error
			}{
				Book: &book.Book{
					ID:           1,
					Title:        "テスト書籍",
					Description:  "本の説明",
					Isbn:         "978-1-234-56789-7",
					ThumbnailURL: "",
					Version:      "0.0.1",
					Publisher:    "テスト出版社",
					PublishedOn:  time.Date(2020, 1, 1, 0, 0, 0, 0, time.Local),
					CreatedAt:    current,
					UpdatedAt:    current,
					Bookshelves:  []*book.Bookshelf{},
					Authors: []*book.Author{{
						ID:        1,
						Name:      "テスト著者",
						CreatedAt: current,
						UpdatedAt: current,
					}},
					Categories: []*book.Category{{
						ID:        1,
						Name:      "コミック",
						CreatedAt: current,
						UpdatedAt: current,
					}},
				},
				Error: nil,
			},
		},
	}

	for result, tc := range testCases {
		t.Run(result, func(t *testing.T) {
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			bvm := mock_book.NewMockValidation(ctrl)

			brm := mock_book.NewMockRepository(ctrl)
			brm.EXPECT().ShowByIsbn(ctx, tc.Isbn).Return(tc.Expected.Book, tc.Expected.Error)
			brm.EXPECT().ShowAuthorsByBookID(ctx, tc.Expected.Book.ID).Return(tc.Expected.Book.Authors, nil)
			brm.EXPECT().ShowCategoriesByBookID(ctx, tc.Expected.Book.ID).Return(tc.Expected.Book.Categories, nil)

			t.Run(result, func(t *testing.T) {
				target := NewBookService(bvm, brm)

				got, err := target.ShowByIsbn(ctx, tc.Isbn)
				if !reflect.DeepEqual(err, tc.Expected.Error) {
					t.Fatalf("want %#v, but %#v", tc.Expected.Error, err)
					return
				}

				if !reflect.DeepEqual(got, tc.Expected.Book) {
					t.Fatalf("want %#v, but %#v", tc.Expected.Book, got)
					return
				}
			})
		})
	}
}

func TestBookService_Create(t *testing.T) {
	testCases := map[string]struct {
		Book     *book.Book
		Expected error
	}{
		"ok": {
			Book: &book.Book{
				ID:           0,
				Title:        "テスト",
				Description:  "",
				Isbn:         "",
				ThumbnailURL: "",
				Version:      "0.0.1",
				Publisher:    "テスト出版社",
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

			bvm := mock_book.NewMockValidation(ctrl)

			brm := mock_book.NewMockRepository(ctrl)
			brm.EXPECT().Create(ctx, tc.Book).Return(tc.Expected)
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

func TestBookService_Update(t *testing.T) {
	current := time.Now()

	testCases := map[string]struct {
		Book     *book.Book
		Expected error
	}{
		"ok": {
			Book: &book.Book{
				ID:           1,
				Title:        "テスト書籍",
				Description:  "本の説明",
				Isbn:         "978-1-234-56789-7",
				ThumbnailURL: "",
				Version:      "0.0.1",
				Publisher:    "テスト出版社",
				PublishedOn:  time.Date(2020, 1, 1, 0, 0, 0, 0, time.Local),
				CreatedAt:    current,
				UpdatedAt:    current,
				Bookshelves:  []*book.Bookshelf{},
				Authors: []*book.Author{{
					ID:        1,
					Name:      "テスト著者",
					CreatedAt: current,
					UpdatedAt: current,
				}},
				Categories: []*book.Category{{
					ID:        1,
					Name:      "コミック",
					CreatedAt: current,
					UpdatedAt: current,
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

			brm := mock_book.NewMockRepository(ctrl)
			brm.EXPECT().Update(ctx, tc.Book).Return(tc.Expected)
			for _, a := range tc.Book.Authors {
				brm.EXPECT().CreateAuthor(ctx, a).Return(tc.Expected)
			}
			for _, c := range tc.Book.Categories {
				brm.EXPECT().CreateCategory(ctx, c).Return(tc.Expected)
			}

			t.Run(result, func(t *testing.T) {
				target := NewBookService(bvm, brm)

				err := target.Update(ctx, tc.Book)
				if !reflect.DeepEqual(err, tc.Expected) {
					t.Fatalf("want %#v, but %#v", tc.Expected, err)
					return
				}
			})
		})
	}
}

func TestBookService_MultipleCreate(t *testing.T) {
	testCases := map[string]struct {
		Books    []*book.Book
		Expected error
	}{
		"ok": {
			Books: []*book.Book{
				{
					ID:           0,
					Title:        "テスト",
					Description:  "",
					Isbn:         "",
					ThumbnailURL: "",
					Version:      "0.0.1",
					Publisher:    "テスト出版社",
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

			brm := mock_book.NewMockRepository(ctrl)
			brm.EXPECT().MultipleCreate(ctx, tc.Books).Return(tc.Expected)
			for _, b := range tc.Books {
				for _, a := range b.Authors {
					brm.EXPECT().CreateAuthor(ctx, a).Return(tc.Expected)
				}
				for _, c := range b.Categories {
					brm.EXPECT().CreateCategory(ctx, c).Return(tc.Expected)
				}
			}

			t.Run(result, func(t *testing.T) {
				target := NewBookService(bvm, brm)

				err := target.MultipleCreate(ctx, tc.Books)
				if !reflect.DeepEqual(err, tc.Expected) {
					t.Fatalf("want %#v, but %#v", tc.Expected, err)
					return
				}
			})
		})
	}
}

func TestBookService_MultipleUpdate(t *testing.T) {
	current := time.Now()

	testCases := map[string]struct {
		Books    []*book.Book
		Expected error
	}{
		"ok": {
			Books: []*book.Book{
				{
					ID:           1,
					Title:        "テスト書籍",
					Description:  "本の説明",
					Isbn:         "978-1-234-56789-7",
					ThumbnailURL: "",
					Version:      "0.0.1",
					Publisher:    "テスト出版社",
					PublishedOn:  time.Date(2020, 1, 1, 0, 0, 0, 0, time.Local),
					CreatedAt:    current,
					UpdatedAt:    current,
					Bookshelves:  []*book.Bookshelf{},
					Authors: []*book.Author{{
						ID:        1,
						Name:      "テスト著者",
						CreatedAt: current,
						UpdatedAt: current,
					}},
					Categories: []*book.Category{{
						ID:        1,
						Name:      "コミック",
						CreatedAt: current,
						UpdatedAt: current,
					}},
				},
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

			brm := mock_book.NewMockRepository(ctrl)
			brm.EXPECT().MultipleUpdate(ctx, tc.Books).Return(tc.Expected)
			for _, b := range tc.Books {
				for _, a := range b.Authors {
					brm.EXPECT().CreateAuthor(ctx, a).Return(tc.Expected)
				}
				for _, c := range b.Categories {
					brm.EXPECT().CreateCategory(ctx, c).Return(tc.Expected)
				}
			}

			t.Run(result, func(t *testing.T) {
				target := NewBookService(bvm, brm)

				err := target.MultipleUpdate(ctx, tc.Books)
				if !reflect.DeepEqual(err, tc.Expected) {
					t.Fatalf("want %#v, but %#v", tc.Expected, err)
					return
				}
			})
		})
	}
}

func TestBookService_Validation(t *testing.T) {
	testCases := map[string]struct {
		Book     *book.Book
		Expected error
	}{
		"ok": {
			Book: &book.Book{
				ID:           0,
				Title:        "テスト",
				Description:  "",
				Isbn:         "",
				ThumbnailURL: "",
				Version:      "0.0.1",
				Publisher:    "テスト出版社",
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

			bvm := mock_book.NewMockValidation(ctrl)
			bvm.EXPECT().Book(ctx, tc.Book).Return(tc.Expected)

			brm := mock_book.NewMockRepository(ctrl)

			t.Run(result, func(t *testing.T) {
				target := NewBookService(bvm, brm)

				err := target.Validation(ctx, tc.Book)
				if !reflect.DeepEqual(err, tc.Expected) {
					t.Fatalf("want %#v, but %#v", tc.Expected, err)
					return
				}
			})
		})
	}
}

func TestBookService_ValidationAuthor(t *testing.T) {
	testCases := map[string]struct {
		Author   *book.Author
		Expected error
	}{
		"ok": {
			Author: &book.Author{
				Name: "テスト著者",
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
			bvm.EXPECT().Author(ctx, tc.Author).Return(tc.Expected)

			brm := mock_book.NewMockRepository(ctrl)

			t.Run(result, func(t *testing.T) {
				target := NewBookService(bvm, brm)

				err := target.ValidationAuthor(ctx, tc.Author)
				if !reflect.DeepEqual(err, tc.Expected) {
					t.Fatalf("want %#v, but %#v", tc.Expected, err)
					return
				}
			})
		})
	}
}

func TestBookService_ValidationCategory(t *testing.T) {
	testCases := map[string]struct {
		Category *book.Category
		Expected error
	}{
		"ok": {
			Category: &book.Category{
				Name: "コミック",
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
			bvm.EXPECT().Category(ctx, tc.Category).Return(tc.Expected)

			brm := mock_book.NewMockRepository(ctrl)

			t.Run(result, func(t *testing.T) {
				target := NewBookService(bvm, brm)

				err := target.ValidationCategory(ctx, tc.Category)
				if !reflect.DeepEqual(err, tc.Expected) {
					t.Fatalf("want %#v, but %#v", tc.Expected, err)
					return
				}
			})
		})
	}
}
