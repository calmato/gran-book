package service

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/calmato/gran-book/api/server/book/internal/domain"
	"github.com/calmato/gran-book/api/server/book/internal/domain/book"
	mock_book "github.com/calmato/gran-book/api/server/book/mock/domain/book"
	"github.com/golang/mock/gomock"
)

func TestBookService_ListBookshelf(t *testing.T) {
	testCases := map[string]struct {
		Query    *domain.ListQuery
		Expected struct {
			Bookshelves []*book.Bookshelf
			Error       error
		}
	}{
		"ok": {
			Query: &domain.ListQuery{
				Limit:      100,
				Offset:     0,
				Order:      nil,
				Conditions: []*domain.QueryCondition{},
			},
			Expected: struct {
				Bookshelves []*book.Bookshelf
				Error       error
			}{
				Bookshelves: []*book.Bookshelf{
					{
						ID:        0,
						BookID:    1,
						UserID:    "00000000-0000-0000-0000-000000000000",
						Status:    5,
						CreatedAt: time.Time{},
						UpdatedAt: time.Time{},
					},
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
			brm.EXPECT().ListBookshelf(ctx, tc.Query).Return(tc.Expected.Bookshelves, tc.Expected.Error)

			t.Run(result, func(t *testing.T) {
				target := NewBookService(bvm, brm)

				got, err := target.ListBookshelf(ctx, tc.Query)
				if !reflect.DeepEqual(err, tc.Expected.Error) {
					t.Fatalf("want %#v, but %#v", tc.Expected.Error, err)
					return
				}

				if !reflect.DeepEqual(got, tc.Expected.Bookshelves) {
					t.Fatalf("want %#v, but %#v", tc.Expected.Bookshelves, got)
					return
				}
			})
		})
	}
}

func TestUserService_ListCount(t *testing.T) {
	testCases := map[string]struct {
		Query    *domain.ListQuery
		Expected struct {
			Count int
			Error error
		}
	}{
		"ok": {
			Query: &domain.ListQuery{
				Limit:      100,
				Offset:     0,
				Order:      nil,
				Conditions: []*domain.QueryCondition{},
			},
			Expected: struct {
				Count int
				Error error
			}{
				Count: 1,
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
			brm.EXPECT().ListBookshelfCount(ctx, tc.Query).Return(tc.Expected.Count, tc.Expected.Error)

			t.Run(result, func(t *testing.T) {
				target := NewBookService(bvm, brm)

				got, err := target.ListBookshelfCount(ctx, tc.Query)
				if !reflect.DeepEqual(err, tc.Expected.Error) {
					t.Fatalf("want %#v, but %#v", tc.Expected.Error, err)
					return
				}

				if !reflect.DeepEqual(got, tc.Expected.Count) {
					t.Fatalf("want %#v, but %#v", tc.Expected.Count, got)
					return
				}
			})
		})
	}
}

func TestBookService_Show(t *testing.T) {
	testCases := map[string]struct {
		BookID   int
		Expected struct {
			Book  *book.Book
			Error error
		}
	}{
		"ok": {
			BookID: 1,
			Expected: struct {
				Book  *book.Book
				Error error
			}{
				Book: &book.Book{
					ID:           1,
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
					Bookshelf:    &book.Bookshelf{},
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
			brm.EXPECT().Show(ctx, tc.BookID).Return(tc.Expected.Book, tc.Expected.Error)
			brm.EXPECT().ListAuthorByBookID(ctx, tc.Expected.Book.ID).Return(tc.Expected.Book.Authors, nil)

			t.Run(result, func(t *testing.T) {
				target := NewBookService(bvm, brm)

				got, err := target.Show(ctx, tc.BookID)
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

func TestBookService_ShowByIsbn(t *testing.T) {
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
					Bookshelf:    &book.Bookshelf{},
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
			brm.EXPECT().ListAuthorByBookID(ctx, tc.Expected.Book.ID).Return(tc.Expected.Book.Authors, nil)

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

func TestBookService_ShowBookshelfByUserIDAndBookID(t *testing.T) {
	testCases := map[string]struct {
		UserID   string
		BookID   int
		Expected struct {
			Bookshelf *book.Bookshelf
			Error     error
		}
	}{
		"ok": {
			UserID: "00000000-0000-0000-0000-000000000000",
			BookID: 1,
			Expected: struct {
				Bookshelf *book.Bookshelf
				Error     error
			}{
				Bookshelf: &book.Bookshelf{
					ID:        0,
					BookID:    1,
					UserID:    "00000000-0000-0000-0000-000000000000",
					Status:    5,
					CreatedAt: time.Time{},
					UpdatedAt: time.Time{},
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
			brm.EXPECT().
				ShowBookshelfByUserIDAndBookID(ctx, tc.UserID, tc.BookID).
				Return(tc.Expected.Bookshelf, tc.Expected.Error)

			t.Run(result, func(t *testing.T) {
				target := NewBookService(bvm, brm)

				got, err := target.ShowBookshelfByUserIDAndBookID(ctx, tc.UserID, tc.BookID)
				if !reflect.DeepEqual(err, tc.Expected.Error) {
					t.Fatalf("want %#v, but %#v", tc.Expected.Error, err)
					return
				}

				if !reflect.DeepEqual(got, tc.Expected.Bookshelf) {
					t.Fatalf("want %#v, but %#v", tc.Expected.Bookshelf, got)
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
				Title:        "テスト書籍",
				TitleKana:    "てすとしょせき",
				Description:  "本の説明です",
				Isbn:         "1234567890123",
				Publisher:    "テスト著者",
				PublishedOn:  "2021年12月24日",
				ThumbnailURL: "",
				Authors:      []*book.Author{},
				Reviews:      []*book.Review{},
				Bookshelf:    &book.Bookshelf{},
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
				brm.EXPECT().ShowOrCreateAuthor(ctx, a).Return(tc.Expected)
			}

			t.Run(result, func(t *testing.T) {
				target := NewBookService(bvm, brm)

				err := target.Create(ctx, tc.Book)
				if !reflect.DeepEqual(err, tc.Expected) {
					t.Fatalf("want %#v, but %#v", tc.Expected, err)
					return
				}

				if tc.Book.CreatedAt.IsZero() {
					t.Fatal("CreatedAt must be not null")
					return
				}

				if tc.Book.UpdatedAt.IsZero() {
					t.Fatal("UpdatedAt must be not null")
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
				BookID: 1,
				UserID: "00000000-0000-0000-0000-000000000000",
				Status: 5,
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

				if tc.Bookshelf.CreatedAt.IsZero() {
					t.Fatal("CreatedAt must be not null")
					return
				}

				if tc.Bookshelf.UpdatedAt.IsZero() {
					t.Fatal("UpdatedAt must be not null")
					return
				}
			})
		})
	}
}

func TestBookService_Update(t *testing.T) {
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
				Publisher:    "テスト出版社",
				PublishedOn:  "2021年12月24日",
				Bookshelf:    &book.Bookshelf{},
				Reviews:      []*book.Review{},
				Authors: []*book.Author{{
					ID:       1,
					Name:     "テスト著者",
					NameKana: "てすとちょしゃ",
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
				brm.EXPECT().ShowOrCreateAuthor(ctx, a).Return(tc.Expected)
			}

			t.Run(result, func(t *testing.T) {
				target := NewBookService(bvm, brm)

				err := target.Update(ctx, tc.Book)
				if !reflect.DeepEqual(err, tc.Expected) {
					t.Fatalf("want %#v, but %#v", tc.Expected, err)
					return
				}

				if tc.Book.UpdatedAt.IsZero() {
					t.Fatal("UpdatedAt must be not null")
					return
				}
			})
		})
	}
}

func TestBookService_UpdateBookshelf(t *testing.T) {
	testCases := map[string]struct {
		Bookshelf *book.Bookshelf
		Expected  error
	}{
		"ok": {
			Bookshelf: &book.Bookshelf{
				ID:     1,
				BookID: 1,
				UserID: "00000000-0000-0000-0000-000000000000",
				Status: 5,
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
			brm.EXPECT().UpdateBookshelf(ctx, tc.Bookshelf).Return(tc.Expected)

			t.Run(result, func(t *testing.T) {
				target := NewBookService(bvm, brm)

				err := target.UpdateBookshelf(ctx, tc.Bookshelf)
				if !reflect.DeepEqual(err, tc.Expected) {
					t.Fatalf("want %#v, but %#v", tc.Expected, err)
					return
				}

				if tc.Bookshelf.UpdatedAt.IsZero() {
					t.Fatal("UpdatedAt must be not null")
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
					Title:        "テスト書籍",
					TitleKana:    "てすとしょせき",
					Description:  "本の説明です",
					Isbn:         "1234567890123",
					Publisher:    "テスト著者",
					PublishedOn:  "2021年12月24日",
					ThumbnailURL: "",
					Authors:      []*book.Author{},
					Reviews:      []*book.Review{},
					Bookshelf:    &book.Bookshelf{},
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
					brm.EXPECT().ShowOrCreateAuthor(ctx, a).Return(tc.Expected)
				}
			}

			t.Run(result, func(t *testing.T) {
				target := NewBookService(bvm, brm)

				err := target.MultipleCreate(ctx, tc.Books)
				if !reflect.DeepEqual(err, tc.Expected) {
					t.Fatalf("want %#v, but %#v", tc.Expected, err)
					return
				}

				for _, v := range tc.Books {
					if v.CreatedAt.IsZero() {
						t.Fatal("CreatedAt must be not null")
						return
					}

					if v.UpdatedAt.IsZero() {
						t.Fatal("UpdatedAt must be not null")
						return
					}
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
					Publisher:    "テスト出版社",
					PublishedOn:  "2021年12月24日",
					CreatedAt:    current,
					UpdatedAt:    current,
					Bookshelf:    &book.Bookshelf{},
					Reviews:      []*book.Review{},
					Authors: []*book.Author{{
						ID:        1,
						Name:      "テスト著者",
						NameKana:  "てすとちょしゃ",
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
					brm.EXPECT().ShowOrCreateAuthor(ctx, a).Return(tc.Expected)
				}
			}

			t.Run(result, func(t *testing.T) {
				target := NewBookService(bvm, brm)

				err := target.MultipleUpdate(ctx, tc.Books)
				if !reflect.DeepEqual(err, tc.Expected) {
					t.Fatalf("want %#v, but %#v", tc.Expected, err)
					return
				}

				for _, v := range tc.Books {
					if v.UpdatedAt.IsZero() {
						t.Fatal("UpdatedAt must be not null")
						return
					}
				}
			})
		})
	}
}

func TestBookService_Delete(t *testing.T) {
	testCases := map[string]struct {
		BookID   int
		Expected error
	}{
		"ok": {
			BookID:   1,
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
			brm.EXPECT().Delete(ctx, tc.BookID).Return(tc.Expected)

			t.Run(result, func(t *testing.T) {
				target := NewBookService(bvm, brm)

				err := target.Delete(ctx, tc.BookID)
				if !reflect.DeepEqual(err, tc.Expected) {
					t.Fatalf("want %#v, but %#v", tc.Expected, err)
					return
				}
			})
		})
	}
}

func TestBookService_DeleteBookshelf(t *testing.T) {
	testCases := map[string]struct {
		BookshelfID int
		Expected    error
	}{
		"ok": {
			BookshelfID: 1,
			Expected:    nil,
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
			brm.EXPECT().DeleteBookshelf(ctx, tc.BookshelfID).Return(tc.Expected)

			t.Run(result, func(t *testing.T) {
				target := NewBookService(bvm, brm)

				err := target.DeleteBookshelf(ctx, tc.BookshelfID)
				if !reflect.DeepEqual(err, tc.Expected) {
					t.Fatalf("want %#v, but %#v", tc.Expected, err)
					return
				}
			})
		})
	}
}

func TestBookService_Validation(t *testing.T) {
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
				Publisher:    "テスト出版社",
				PublishedOn:  "2021年12月24日",
				CreatedAt:    current,
				UpdatedAt:    current,
				Bookshelf:    &book.Bookshelf{},
				Reviews:      []*book.Review{},
				Authors: []*book.Author{{
					ID:        1,
					Name:      "テスト著者",
					NameKana:  "てすとちょしゃ",
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

func TestBookService_ValidationBookshelf(t *testing.T) {
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
			bvm.EXPECT().Bookshelf(ctx, tc.Bookshelf).Return(tc.Expected)

			brm := mock_book.NewMockRepository(ctrl)

			t.Run(result, func(t *testing.T) {
				target := NewBookService(bvm, brm)

				err := target.ValidationBookshelf(ctx, tc.Bookshelf)
				if !reflect.DeepEqual(err, tc.Expected) {
					t.Fatalf("want %#v, but %#v", tc.Expected, err)
					return
				}
			})
		})
	}
}
