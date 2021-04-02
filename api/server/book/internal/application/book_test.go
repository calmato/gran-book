package application

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/calmato/gran-book/api/server/book/internal/application/input"
	"github.com/calmato/gran-book/api/server/book/internal/domain/book"
	"github.com/calmato/gran-book/api/server/book/lib/datetime"
	mock_validation "github.com/calmato/gran-book/api/server/book/mock/application/validation"
	mock_book "github.com/calmato/gran-book/api/server/book/mock/domain/book"
	"github.com/golang/mock/gomock"
)

func TestBookApplication_MultipleCreateAndUpdate(t *testing.T) {
	testCases := map[string]struct {
		Input    *input.CreateAndUpdateBooks
		Expected struct {
			Books []*book.Book
			Error error
		}
	}{
		"ok": {
			Input: &input.CreateAndUpdateBooks{
				Books: []*input.Book{
					{
						Title:        "テスト書籍",
						Description:  "書籍の説明",
						Isbn:         "08881516881516315501",
						ThumbnailURL: "",
						Version:      "1.5.5.0.preview.3",
						Publisher:    "テスト出版社",
						PublishedOn:  "2021-01-01",
						Authors:      []string{"テスト著者"},
						Categories:   []string{"Comics & Graphic Novels"},
					},
				},
			},
			Expected: struct {
				Books []*book.Book
				Error error
			}{
				Books: []*book.Book{
					{
						Title:        "テスト書籍",
						Description:  "書籍の説明",
						Isbn:         "08881516881516315501",
						ThumbnailURL: "",
						Version:      "1.5.4.0.preview.3",
						Publisher:    "テスト出版社",
						PublishedOn:  datetime.StringToDate("2021-01-01"),
						Authors: []*book.Author{{
							Name: "テスト著者",
						}},
						Categories: []*book.Category{{
							Name: "Comics & Graphic Novels",
						}},
					},
				},
				Error: nil,
			},
		},
	}

	for result, tc := range testCases {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		brv := mock_validation.NewMockBookRequestValidation(ctrl)
		brv.EXPECT().CreateAndUpdateBooks(tc.Input).Return(nil)

		bsm := mock_book.NewMockService(ctrl)
		for i, b := range tc.Input.Books {
			bsm.EXPECT().ShowByIsbn(ctx, b.Isbn).Return(tc.Expected.Books[i], tc.Expected.Error)
		}
		bsm.EXPECT().MultipleCreate(ctx, gomock.Any()).Return(tc.Expected.Error)
		bsm.EXPECT().MultipleUpdate(ctx, gomock.Any()).Return(tc.Expected.Error)
		bsm.EXPECT().Validation(ctx, gomock.Any()).Return(tc.Expected.Error)
		bsm.EXPECT().ValidationAuthor(ctx, gomock.Any()).Return(tc.Expected.Error)
		bsm.EXPECT().ValidationCategory(ctx, gomock.Any()).Return(tc.Expected.Error)

		t.Run(result, func(t *testing.T) {
			target := NewBookApplication(brv, bsm)

			got, err := target.MultipleCreateAndUpdate(ctx, tc.Input)
			if !reflect.DeepEqual(err, tc.Expected.Error) {
				t.Fatalf("want %#v, but %#v", tc.Expected.Error, err)
				return
			}

			if !reflect.DeepEqual(got, tc.Expected.Books) {
				t.Fatalf("want %#v, but %#v", tc.Expected.Books, got)
				return
			}
		})
	}
}

func TestBookApplication_CreateOrUpdateBookshelf(t *testing.T) {
	current := time.Now()

	testCases := map[string]struct {
		Input    *input.Bookshelf
		Action   string
		Expected struct {
			Book      *book.Book
			Bookshelf *book.Bookshelf
			Error     error
		}
	}{
		"ok_create": {
			Input: &input.Bookshelf{
				UserID:     "00000000-0000-0000-0000-000000000000",
				BookID:     1,
				Status:     1,
				Impression: "感想です",
				ReadOn:     "2020-01-01",
			},
			Action: "create",
			Expected: struct {
				Book      *book.Book
				Bookshelf *book.Bookshelf
				Error     error
			}{
				Book: &book.Book{
					Title:        "テスト書籍",
					Description:  "書籍の説明",
					Isbn:         "08881516881516315501",
					ThumbnailURL: "",
					Version:      "1.5.4.0.preview.3",
					Publisher:    "テスト出版社",
					PublishedOn:  datetime.StringToDate("2021-01-01"),
					Authors: []*book.Author{{
						Name: "テスト著者",
					}},
					Categories: []*book.Category{{
						Name: "Comics & Graphic Novels",
					}},
				},
				Bookshelf: &book.Bookshelf{
					UserID:     "00000000-0000-0000-0000-000000000000",
					BookID:     1,
					Status:     1,
					Impression: "感想です",
					ReadOn:     datetime.StringToDate("2020-01-01"),
				},
				Error: nil,
			},
		},
		"ok_update": {
			Input: &input.Bookshelf{
				UserID:     "00000000-0000-0000-0000-000000000000",
				BookID:     1,
				Status:     1,
				Impression: "感想です",
				ReadOn:     "2020-01-01",
			},
			Action: "update",
			Expected: struct {
				Book      *book.Book
				Bookshelf *book.Bookshelf
				Error     error
			}{
				Book: &book.Book{
					Title:        "テスト書籍",
					Description:  "書籍の説明",
					Isbn:         "08881516881516315501",
					ThumbnailURL: "",
					Version:      "1.5.4.0.preview.3",
					Publisher:    "テスト出版社",
					PublishedOn:  datetime.StringToDate("2021-01-01"),
					Authors: []*book.Author{{
						Name: "テスト著者",
					}},
					Categories: []*book.Category{{
						Name: "Comics & Graphic Novels",
					}},
				},
				Bookshelf: &book.Bookshelf{
					ID:         1,
					UserID:     "00000000-0000-0000-0000-000000000000",
					BookID:     1,
					Status:     1,
					Impression: "感想です",
					ReadOn:     datetime.StringToDate("2020-01-01"),
					CreatedAt:  current,
					UpdatedAt:  current,
				},
				Error: nil,
			},
		},
	}

	for result, tc := range testCases {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		brv := mock_validation.NewMockBookRequestValidation(ctrl)
		brv.EXPECT().Bookshelf(tc.Input).Return(nil)

		bsm := mock_book.NewMockService(ctrl)
		bsm.EXPECT().Show(ctx, tc.Input.BookID).Return(tc.Expected.Book, tc.Expected.Error)
		bsm.EXPECT().ValidationBookshelf(ctx, gomock.Any()).Return(tc.Expected.Error)

		switch tc.Action {
		case "create":
			bsm.EXPECT().
				ShowBookshelfByUserIDAndBookID(ctx, tc.Input.UserID, tc.Input.BookID).
				Return(nil, tc.Expected.Error)
			bsm.EXPECT().CreateBookshelf(ctx, gomock.Any()).Return(tc.Expected.Error)
		case "update":
			bsm.EXPECT().
				ShowBookshelfByUserIDAndBookID(ctx, tc.Input.UserID, tc.Input.BookID).
				Return(tc.Expected.Bookshelf, tc.Expected.Error)
			bsm.EXPECT().UpdateBookshelf(ctx, tc.Expected.Bookshelf).Return(tc.Expected.Error)
		}

		t.Run(result, func(t *testing.T) {
			target := NewBookApplication(brv, bsm)

			b, bs, err := target.CreateOrUpdateBookshelf(ctx, tc.Input)
			if !reflect.DeepEqual(err, tc.Expected.Error) {
				t.Fatalf("want %#v, but %#v", tc.Expected.Error, err)
				return
			}

			if !reflect.DeepEqual(b, tc.Expected.Book) {
				t.Fatalf("want %#v, but %#v", tc.Expected.Book, b)
				return
			}

			if !reflect.DeepEqual(bs, tc.Expected.Bookshelf) {
				t.Fatalf("want %#v, but %#v", tc.Expected.Bookshelf, bs)
				return
			}
		})
	}
}
