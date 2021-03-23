package application

import (
	"context"
	"reflect"
	"testing"

	"github.com/calmato/gran-book/api/server/book/internal/application/input"
	"github.com/calmato/gran-book/api/server/book/internal/domain/book"
	"github.com/calmato/gran-book/api/server/book/lib/datetime"
	mock_validation "github.com/calmato/gran-book/api/server/book/mock/application/validation"
	mock_book "github.com/calmato/gran-book/api/server/book/mock/domain/book"
	"github.com/golang/mock/gomock"
)

func TestBookApplication_Create(t *testing.T) {
	testCases := map[string]struct {
		Input    *input.BookItem
		Expected struct {
			Book  *book.Book
			Error error
		}
	}{
		"ok": {
			Input: &input.BookItem{
				Title:        "テスト書籍",
				Description:  "書籍の説明",
				Isbn:         "08881516881516315501",
				ThumbnailURL: "",
				Version:      "1.5.4.0.preview.3",
				Publisher:    "テスト出版社",
				PublishedOn:  "2021-01-01",
				Authors: []*input.BookAuthor{{
					Name: "テスト著者",
				}},
				Categories: []*input.BookCategory{{
					Name: "Comics & Graphic Novels",
				}},
			},
			Expected: struct {
				Book  *book.Book
				Error error
			}{
				Book: &book.Book{
					Title:        "テスト書籍",
					Description:  "書籍の説明",
					Isbn:         "08881516881516315501",
					ThumbnailURL: "",
					Version:      "1.5.4.0.preview.3",
					PublishedOn:  datetime.StringToDate("2021-01-01"),
					Publisher: &book.Publisher{
						Name: "テスト出版社",
					},
					Authors: []*book.Author{{
						Name: "テスト著者",
					}},
					Categories: []*book.Category{{
						Name: "Comics & Graphic Novels",
					}},
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
		brv.EXPECT().BookItem(tc.Input).Return(nil)

		bsm := mock_book.NewMockService(ctrl)
		bsm.EXPECT().Create(ctx, gomock.Any()).Return(tc.Expected.Error)
		bsm.EXPECT().Validation(ctx, gomock.Any()).Return(tc.Expected.Error)
		bsm.EXPECT().ValidationAuthor(ctx, gomock.Any()).Return(tc.Expected.Error)
		bsm.EXPECT().ValidationCategory(ctx, gomock.Any()).Return(tc.Expected.Error)
		bsm.EXPECT().ValidationPublisher(ctx, gomock.Any()).Return(tc.Expected.Error)

		t.Run(result, func(t *testing.T) {
			target := NewBookApplication(brv, bsm)

			got, err := target.Create(ctx, tc.Input)
			if !reflect.DeepEqual(err, tc.Expected.Error) {
				t.Fatalf("want %#v, but %#v", tc.Expected.Error, err)
				return
			}

			if !reflect.DeepEqual(got, tc.Expected.Book) {
				t.Fatalf("want %#v, but %#v", tc.Expected.Book, got)
				return
			}
		})
	}
}

func TestBookApplication_Update(t *testing.T) {
	testCases := map[string]struct {
		Input    *input.BookItem
		Expected struct {
			Book  *book.Book
			Error error
		}
	}{
		"ok": {
			Input: &input.BookItem{
				Title:        "テスト書籍",
				Description:  "書籍の説明",
				Isbn:         "08881516881516315501",
				ThumbnailURL: "",
				Version:      "1.5.4.0.preview.3",
				Publisher:    "テスト出版社",
				PublishedOn:  "2021-01-01",
				Authors: []*input.BookAuthor{{
					Name: "テスト著者",
				}},
				Categories: []*input.BookCategory{{
					Name: "Comics & Graphic Novels",
				}},
			},
			Expected: struct {
				Book  *book.Book
				Error error
			}{
				Book: &book.Book{
					Title:        "テスト書籍",
					Description:  "書籍の説明",
					Isbn:         "08881516881516315501",
					ThumbnailURL: "",
					Version:      "1.5.4.0.preview.3",
					PublishedOn:  datetime.StringToDate("2021-01-01"),
					Publisher: &book.Publisher{
						Name: "テスト出版社",
					},
					Authors: []*book.Author{{
						Name: "テスト著者",
					}},
					Categories: []*book.Category{{
						Name: "Comics & Graphic Novels",
					}},
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
		brv.EXPECT().BookItem(tc.Input).Return(nil)

		bsm := mock_book.NewMockService(ctrl)
		bsm.EXPECT().ShowByIsbn(ctx, tc.Input.Isbn).Return(tc.Expected.Book, tc.Expected.Error)
		bsm.EXPECT().Update(ctx, tc.Expected.Book).Return(tc.Expected.Error)
		bsm.EXPECT().Validation(ctx, gomock.Any()).Return(tc.Expected.Error)
		bsm.EXPECT().ValidationAuthor(ctx, gomock.Any()).Return(tc.Expected.Error)
		bsm.EXPECT().ValidationCategory(ctx, gomock.Any()).Return(tc.Expected.Error)
		bsm.EXPECT().ValidationPublisher(ctx, gomock.Any()).Return(tc.Expected.Error)

		t.Run(result, func(t *testing.T) {
			target := NewBookApplication(brv, bsm)

			got, err := target.Update(ctx, tc.Input)
			if !reflect.DeepEqual(err, tc.Expected.Error) {
				t.Fatalf("want %#v, but %#v", tc.Expected.Error, err)
				return
			}

			if !reflect.DeepEqual(got, tc.Expected.Book) {
				t.Fatalf("want %#v, but %#v", tc.Expected.Book, got)
				return
			}
		})
	}
}

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
				Books: []*input.BookItem{
					{
						Title:        "テスト書籍",
						Description:  "書籍の説明",
						Isbn:         "08881516881516315501",
						ThumbnailURL: "",
						Version:      "1.5.5.0.preview.3",
						Publisher:    "テスト出版社",
						PublishedOn:  "2021-01-01",
						Authors: []*input.BookAuthor{{
							Name: "テスト著者",
						}},
						Categories: []*input.BookCategory{{
							Name: "Comics & Graphic Novels",
						}},
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
						PublishedOn:  datetime.StringToDate("2021-01-01"),
						Publisher: &book.Publisher{
							Name: "テスト出版社",
						},
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
		bsm.EXPECT().ValidationPublisher(ctx, gomock.Any()).Return(tc.Expected.Error)

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
