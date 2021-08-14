package validation

import (
	"context"
	"reflect"
	"testing"

	"github.com/calmato/gran-book/api/server/book/internal/domain/book"
	"github.com/calmato/gran-book/api/server/book/lib/datetime"
	mock_book "github.com/calmato/gran-book/api/server/book/mock/domain/book"
	"github.com/golang/mock/gomock"
)

func TestBookService_Book(t *testing.T) {
	type args struct {
		book *book.Book
	}

	testCases := map[string]struct {
		args args
		want error
	}{
		"ok": {
			args: args{
				book: &book.Book{
					Title:        "テスト書籍",
					TitleKana:    "てすとしょせき",
					Description:  "本の説明です",
					Isbn:         "1234567890123",
					Publisher:    "テスト著者",
					PublishedOn:  "2021年12月24日",
					ThumbnailURL: "",
				},
			},
			want: nil,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			bvm := mock_book.NewMockRepository(ctrl)
			bvm.EXPECT().GetBookIDByIsbn(ctx, tc.args.book.Isbn).Return(tc.args.book.ID, nil)

			t.Run(name, func(t *testing.T) {
				target := NewBookDomainValidation(bvm)

				err := target.Book(ctx, tc.args.book)
				if !reflect.DeepEqual(err, tc.want) {
					t.Fatalf("want %#v, but %#v", tc.want, err)
					return
				}
			})
		})
	}
}

func TestBookService_Author(t *testing.T) {
	type args struct {
		author *book.Author
	}

	testCases := map[string]struct {
		args args
		want error
	}{
		"ok": {
			args: args{
				author: &book.Author{
					Name:     "テスト著者",
					NameKana: "てすとちょしゃ",
				},
			},
			want: nil,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			bvm := mock_book.NewMockRepository(ctrl)

			t.Run(name, func(t *testing.T) {
				target := NewBookDomainValidation(bvm)

				err := target.Author(ctx, tc.args.author)
				if !reflect.DeepEqual(err, tc.want) {
					t.Fatalf("want %#v, but %#v", tc.want, err)
					return
				}
			})
		})
	}
}

func TestBookService_Bookshelf(t *testing.T) {
	type args struct {
		bookshelf *book.Bookshelf
	}

	testCases := map[string]struct {
		args args
		want error
	}{
		"ok": {
			args: args{
				bookshelf: &book.Bookshelf{
					BookID: 1,
					UserID: "00000000-0000-0000-0000-000000000000",
					Status: 5,
					ReadOn: datetime.StringToDate("2020-01-01"),
				},
			},
			want: nil,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			bvm := mock_book.NewMockRepository(ctrl)
			bvm.EXPECT().
				GetBookshelfIDByUserIDAndBookID(ctx, tc.args.bookshelf.UserID, tc.args.bookshelf.BookID).
				Return(tc.args.bookshelf.ID, nil)

			t.Run(name, func(t *testing.T) {
				target := NewBookDomainValidation(bvm)

				err := target.Bookshelf(ctx, tc.args.bookshelf)
				if !reflect.DeepEqual(err, tc.want) {
					t.Fatalf("want %#v, but %#v", tc.want, err)
					return
				}
			})
		})
	}
}

func TestBookService_Review(t *testing.T) {
	type args struct {
		review *book.Review
	}

	testCases := map[string]struct {
		args args
		want error
	}{
		"ok": {
			args: args{
				review: &book.Review{
					BookID:     1,
					UserID:     "00000000-0000-0000-0000-000000000000",
					Score:      5,
					Impression: "書籍の感想です。",
				},
			},
			want: nil,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			bvm := mock_book.NewMockRepository(ctrl)
			bvm.EXPECT().
				GetReviewIDByUserIDAndBookID(ctx, tc.args.review.UserID, tc.args.review.BookID).
				Return(tc.args.review.ID, nil)

			t.Run(name, func(t *testing.T) {
				target := NewBookDomainValidation(bvm)

				err := target.Review(ctx, tc.args.review)
				if !reflect.DeepEqual(err, tc.want) {
					t.Fatalf("want %#v, but %#v", tc.want, err)
					return
				}
			})
		})
	}
}
