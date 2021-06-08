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

func TestBookService_List(t *testing.T) {
	type args struct {
		query *domain.ListQuery
	}
	type want struct {
		books []*book.Book
		err   error
	}

	testCases := map[string]struct {
		args args
		want want
	}{
		"ok": {
			args: args{
				query: &domain.ListQuery{
					Limit:      100,
					Offset:     0,
					Order:      nil,
					Conditions: []*domain.QueryCondition{},
				},
			},
			want: want{
				books: []*book.Book{
					{
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
					},
				},
				err: nil,
			},
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			bvm := mock_book.NewMockValidation(ctrl)

			brm := mock_book.NewMockRepository(ctrl)
			brm.EXPECT().List(ctx, tc.args.query).Return(tc.want.books, tc.want.err)

			t.Run(name, func(t *testing.T) {
				target := NewBookService(bvm, brm)

				books, err := target.List(ctx, tc.args.query)
				if !reflect.DeepEqual(err, tc.want.err) {
					t.Fatalf("want %#v, but %#v", tc.want.err, err)
					return
				}

				if !reflect.DeepEqual(books, tc.want.books) {
					t.Fatalf("want %#v, but %#v", tc.want.books, books)
					return
				}
			})
		})
	}
}

func TestBookService_ListBookshelf(t *testing.T) {
	type args struct {
		query *domain.ListQuery
	}
	type want struct {
		bookshelves []*book.Bookshelf
		err         error
	}

	testCases := map[string]struct {
		args args
		want want
	}{
		"ok": {
			args: args{
				query: &domain.ListQuery{
					Limit:      100,
					Offset:     0,
					Order:      nil,
					Conditions: []*domain.QueryCondition{},
				},
			},
			want: want{
				bookshelves: []*book.Bookshelf{
					{
						ID:        0,
						BookID:    1,
						UserID:    "00000000-0000-0000-0000-000000000000",
						Status:    5,
						CreatedAt: time.Time{},
						UpdatedAt: time.Time{},
					},
				},
				err: nil,
			},
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			bvm := mock_book.NewMockValidation(ctrl)

			brm := mock_book.NewMockRepository(ctrl)
			brm.EXPECT().ListBookshelf(ctx, tc.args.query).Return(tc.want.bookshelves, tc.want.err)

			t.Run(name, func(t *testing.T) {
				target := NewBookService(bvm, brm)

				bookshelves, err := target.ListBookshelf(ctx, tc.args.query)
				if !reflect.DeepEqual(err, tc.want.err) {
					t.Fatalf("want %#v, but %#v", tc.want.err, err)
					return
				}

				if !reflect.DeepEqual(bookshelves, tc.want.bookshelves) {
					t.Fatalf("want %#v, but %#v", tc.want.bookshelves, bookshelves)
					return
				}
			})
		})
	}
}

func TestBookService_ListReview(t *testing.T) {
	type args struct {
		query *domain.ListQuery
	}
	type want struct {
		reviews []*book.Review
		err     error
	}

	testCases := map[string]struct {
		args args
		want want
	}{
		"ok": {
			args: args{
				query: &domain.ListQuery{
					Limit:      100,
					Offset:     0,
					Order:      nil,
					Conditions: []*domain.QueryCondition{},
				},
			},
			want: want{
				reviews: []*book.Review{
					{
						ID:         1,
						BookID:     1,
						UserID:     "00000000-0000-0000-0000-000000000000",
						Score:      5,
						Impression: "書籍の感想です。",
						CreatedAt:  time.Time{},
						UpdatedAt:  time.Time{},
					},
				},
				err: nil,
			},
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			bvm := mock_book.NewMockValidation(ctrl)

			brm := mock_book.NewMockRepository(ctrl)
			brm.EXPECT().ListReview(ctx, tc.args.query).Return(tc.want.reviews, tc.want.err)

			t.Run(name, func(t *testing.T) {
				target := NewBookService(bvm, brm)

				reviews, err := target.ListReview(ctx, tc.args.query)
				if !reflect.DeepEqual(err, tc.want.err) {
					t.Fatalf("want %#v, but %#v", tc.want.err, err)
					return
				}

				if !reflect.DeepEqual(reviews, tc.want.reviews) {
					t.Fatalf("want %#v, but %#v", tc.want.reviews, reviews)
					return
				}
			})
		})
	}
}

func TestUserService_ListCount(t *testing.T) {
	type args struct {
		query *domain.ListQuery
	}
	type want struct {
		count int
		err   error
	}

	testCases := map[string]struct {
		args args
		want want
	}{
		"ok": {
			args: args{
				query: &domain.ListQuery{
					Limit:      100,
					Offset:     0,
					Order:      nil,
					Conditions: []*domain.QueryCondition{},
				},
			},
			want: want{
				count: 1,
				err:   nil,
			},
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			bvm := mock_book.NewMockValidation(ctrl)

			brm := mock_book.NewMockRepository(ctrl)
			brm.EXPECT().ListCount(ctx, tc.args.query).Return(tc.want.count, tc.want.err)

			t.Run(name, func(t *testing.T) {
				target := NewBookService(bvm, brm)

				count, err := target.ListCount(ctx, tc.args.query)
				if !reflect.DeepEqual(err, tc.want.err) {
					t.Fatalf("want %#v, but %#v", tc.want.err, err)
					return
				}

				if !reflect.DeepEqual(count, tc.want.count) {
					t.Fatalf("want %#v, but %#v", tc.want.count, count)
					return
				}
			})
		})
	}
}

func TestUserService_ListBookshelfCount(t *testing.T) {
	type args struct {
		query *domain.ListQuery
	}
	type want struct {
		count int
		err   error
	}

	testCases := map[string]struct {
		args args
		want want
	}{
		"ok": {
			args: args{
				query: &domain.ListQuery{
					Limit:      100,
					Offset:     0,
					Order:      nil,
					Conditions: []*domain.QueryCondition{},
				},
			},
			want: want{
				count: 1,
				err:   nil,
			},
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			bvm := mock_book.NewMockValidation(ctrl)

			brm := mock_book.NewMockRepository(ctrl)
			brm.EXPECT().ListBookshelfCount(ctx, tc.args.query).Return(tc.want.count, tc.want.err)

			t.Run(name, func(t *testing.T) {
				target := NewBookService(bvm, brm)

				count, err := target.ListBookshelfCount(ctx, tc.args.query)
				if !reflect.DeepEqual(err, tc.want.err) {
					t.Fatalf("want %#v, but %#v", tc.want.err, err)
					return
				}

				if !reflect.DeepEqual(count, tc.want.count) {
					t.Fatalf("want %#v, but %#v", tc.want.count, count)
					return
				}
			})
		})
	}
}

func TestUserService_ListReviewCount(t *testing.T) {
	type args struct {
		query *domain.ListQuery
	}
	type want struct {
		count int
		err   error
	}

	testCases := map[string]struct {
		args args
		want want
	}{
		"ok": {
			args: args{
				query: &domain.ListQuery{
					Limit:      100,
					Offset:     0,
					Order:      nil,
					Conditions: []*domain.QueryCondition{},
				},
			},
			want: want{
				count: 1,
				err:   nil,
			},
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			bvm := mock_book.NewMockValidation(ctrl)

			brm := mock_book.NewMockRepository(ctrl)
			brm.EXPECT().ListReviewCount(ctx, tc.args.query).Return(tc.want.count, tc.want.err)

			t.Run(name, func(t *testing.T) {
				target := NewBookService(bvm, brm)

				count, err := target.ListReviewCount(ctx, tc.args.query)
				if !reflect.DeepEqual(err, tc.want.err) {
					t.Fatalf("want %#v, but %#v", tc.want.err, err)
					return
				}

				if !reflect.DeepEqual(count, tc.want.count) {
					t.Fatalf("want %#v, but %#v", tc.want.count, count)
					return
				}
			})
		})
	}
}

func TestBookService_Show(t *testing.T) {
	type args struct {
		bookID int
	}
	type want struct {
		book *book.Book
		err  error
	}

	testCases := map[string]struct {
		args args
		want want
	}{
		"ok": {
			args: args{
				bookID: 1,
			},
			want: want{
				book: &book.Book{
					ID:           1,
					Title:        "テスト書籍",
					TitleKana:    "てすとしょせき",
					Description:  "本の説明です",
					Isbn:         "1234567890123",
					Publisher:    "テスト著者",
					PublishedOn:  "2021年12月24日",
					ThumbnailURL: "",
				},
				err: nil,
			},
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			bvm := mock_book.NewMockValidation(ctrl)

			brm := mock_book.NewMockRepository(ctrl)
			brm.EXPECT().Show(ctx, tc.want.book.ID).Return(tc.want.book, tc.want.err)

			t.Run(name, func(t *testing.T) {
				target := NewBookService(bvm, brm)

				book, err := target.Show(ctx, tc.args.bookID)
				if !reflect.DeepEqual(err, tc.want.err) {
					t.Fatalf("want %#v, but %#v", tc.want.err, err)
					return
				}

				if !reflect.DeepEqual(book, tc.want.book) {
					t.Fatalf("want %#v, but %#v", tc.want.book, book)
					return
				}
			})
		})
	}
}

func TestBookService_ShowByIsbn(t *testing.T) {
	type args struct {
		isbn string
	}
	type want struct {
		book *book.Book
		err  error
	}

	testCases := map[string]struct {
		args args
		want want
	}{
		"ok": {
			args: args{
				isbn: "978-1-234-56789-7",
			},
			want: want{
				book: &book.Book{
					ID:           1,
					Title:        "テスト書籍",
					TitleKana:    "てすとしょせき",
					Description:  "本の説明です",
					Isbn:         "1234567890123",
					Publisher:    "テスト著者",
					PublishedOn:  "2021年12月24日",
					ThumbnailURL: "",
				},
				err: nil,
			},
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			bvm := mock_book.NewMockValidation(ctrl)

			brm := mock_book.NewMockRepository(ctrl)
			brm.EXPECT().ShowByIsbn(ctx, tc.args.isbn).Return(tc.want.book, tc.want.err)

			t.Run(name, func(t *testing.T) {
				target := NewBookService(bvm, brm)

				book, err := target.ShowByIsbn(ctx, tc.args.isbn)
				if !reflect.DeepEqual(err, tc.want.err) {
					t.Fatalf("want %#v, but %#v", tc.want.err, err)
					return
				}

				if !reflect.DeepEqual(book, tc.want.book) {
					t.Fatalf("want %#v, but %#v", tc.want.book, book)
					return
				}
			})
		})
	}
}

func TestBookService_ShowBookshelfByUserIDAndBookID(t *testing.T) {
	type args struct {
		userID string
		bookID int
	}
	type want struct {
		bookshelf *book.Bookshelf
		err       error
	}

	testCases := map[string]struct {
		args args
		want want
	}{
		"ok": {
			args: args{
				userID: "00000000-0000-0000-0000-000000000000",
				bookID: 1,
			},
			want: want{
				bookshelf: &book.Bookshelf{
					ID:     1,
					BookID: 1,
					UserID: "00000000-0000-0000-0000-000000000000",
					Status: 5,
				},
				err: nil,
			},
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			bvm := mock_book.NewMockValidation(ctrl)

			brm := mock_book.NewMockRepository(ctrl)
			brm.EXPECT().
				ShowBookshelfByUserIDAndBookID(ctx, tc.args.userID, tc.args.bookID).
				Return(tc.want.bookshelf, tc.want.err)

			t.Run(name, func(t *testing.T) {
				target := NewBookService(bvm, brm)

				bookshelf, err := target.ShowBookshelfByUserIDAndBookID(ctx, tc.args.userID, tc.args.bookID)
				if !reflect.DeepEqual(err, tc.want.err) {
					t.Fatalf("want %#v, but %#v", tc.want.err, err)
					return
				}

				if !reflect.DeepEqual(bookshelf, tc.want.bookshelf) {
					t.Fatalf("want %#v, but %#v", tc.want.bookshelf, bookshelf)
					return
				}
			})
		})
	}
}

func TestBookService_ShowReview(t *testing.T) {
	type args struct {
		reviewID int
	}
	type want struct {
		review *book.Review
		err    error
	}

	testCases := map[string]struct {
		args args
		want want
	}{
		"ok": {
			args: args{
				reviewID: 1,
			},
			want: want{
				review: &book.Review{
					ID:         1,
					BookID:     1,
					UserID:     "00000000-0000-0000-0000-000000000000",
					Score:      1,
					Impression: "読んだ感想です。",
				},
				err: nil,
			},
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			bvm := mock_book.NewMockValidation(ctrl)

			brm := mock_book.NewMockRepository(ctrl)
			brm.EXPECT().ShowReview(ctx, tc.args.reviewID).Return(tc.want.review, tc.want.err)

			t.Run(name, func(t *testing.T) {
				target := NewBookService(bvm, brm)

				review, err := target.ShowReview(ctx, tc.args.reviewID)
				if !reflect.DeepEqual(err, tc.want.err) {
					t.Fatalf("want %#v, but %#v", tc.want.err, err)
					return
				}

				if !reflect.DeepEqual(review, tc.want.review) {
					t.Fatalf("want %#v, but %#v", tc.want.review, review)
					return
				}
			})
		})
	}
}

func TestBookService_ShowReviewByUserIDAndBookID(t *testing.T) {
	type args struct {
		userID string
		bookID int
	}
	type want struct {
		review *book.Review
		err    error
	}

	testCases := map[string]struct {
		args args
		want want
	}{
		"ok": {
			args: args{
				userID: "00000000-0000-0000-0000-000000000000",
				bookID: 1,
			},
			want: want{
				review: &book.Review{
					ID:         1,
					BookID:     1,
					UserID:     "00000000-0000-0000-0000-000000000000",
					Score:      1,
					Impression: "読んだ感想です。",
				},
				err: nil,
			},
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			bvm := mock_book.NewMockValidation(ctrl)

			brm := mock_book.NewMockRepository(ctrl)
			brm.EXPECT().
				ShowReviewByUserIDAndBookID(ctx, tc.args.userID, tc.args.bookID).
				Return(tc.want.review, tc.want.err)

			t.Run(name, func(t *testing.T) {
				target := NewBookService(bvm, brm)

				review, err := target.ShowReviewByUserIDAndBookID(ctx, tc.args.userID, tc.args.bookID)
				if !reflect.DeepEqual(err, tc.want.err) {
					t.Fatalf("want %#v, but %#v", tc.want.err, err)
					return
				}

				if !reflect.DeepEqual(review, tc.want.review) {
					t.Fatalf("want %#v, but %#v", tc.want.review, review)
					return
				}
			})
		})
	}
}

func TestBookService_Create(t *testing.T) {
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
					Authors:      []*book.Author{},
					Reviews:      []*book.Review{},
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

			bvm := mock_book.NewMockValidation(ctrl)

			brm := mock_book.NewMockRepository(ctrl)
			brm.EXPECT().Create(ctx, tc.args.book).Return(tc.want)
			for _, a := range tc.args.book.Authors {
				brm.EXPECT().ShowOrCreateAuthor(ctx, a).Return(nil)
			}

			t.Run(name, func(t *testing.T) {
				target := NewBookService(bvm, brm)

				err := target.Create(ctx, tc.args.book)
				if !reflect.DeepEqual(err, tc.want) {
					t.Fatalf("want %#v, but %#v", tc.want, err)
					return
				}

				if tc.args.book.CreatedAt.IsZero() {
					t.Fatal("CreatedAt must be not null")
					return
				}

				if tc.args.book.UpdatedAt.IsZero() {
					t.Fatal("UpdatedAt must be not null")
					return
				}
			})
		})
	}
}

func TestBookService_CreateBookshelf(t *testing.T) {
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

			bvm := mock_book.NewMockValidation(ctrl)

			brm := mock_book.NewMockRepository(ctrl)
			brm.EXPECT().CreateBookshelf(ctx, tc.args.bookshelf).Return(tc.want)

			t.Run(name, func(t *testing.T) {
				target := NewBookService(bvm, brm)

				err := target.CreateBookshelf(ctx, tc.args.bookshelf)
				if !reflect.DeepEqual(err, tc.want) {
					t.Fatalf("want %#v, but %#v", tc.want, err)
					return
				}

				if tc.args.bookshelf.CreatedAt.IsZero() {
					t.Fatal("CreatedAt must be not null")
					return
				}

				if tc.args.bookshelf.UpdatedAt.IsZero() {
					t.Fatal("UpdatedAt must be not null")
					return
				}
			})
		})
	}
}

func TestBookService_Update(t *testing.T) {
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
					ID:           1,
					Title:        "テスト書籍",
					Description:  "本の説明",
					Isbn:         "978-1-234-56789-7",
					ThumbnailURL: "",
					Publisher:    "テスト出版社",
					PublishedOn:  "2021年12月24日",
					Authors: []*book.Author{{
						ID:       1,
						Name:     "テスト著者",
						NameKana: "てすとちょしゃ",
					}},
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

			bvm := mock_book.NewMockValidation(ctrl)

			brm := mock_book.NewMockRepository(ctrl)
			brm.EXPECT().Update(ctx, tc.args.book).Return(tc.want)
			for _, a := range tc.args.book.Authors {
				brm.EXPECT().ShowOrCreateAuthor(ctx, a).Return(nil)
			}

			t.Run(name, func(t *testing.T) {
				target := NewBookService(bvm, brm)

				err := target.Update(ctx, tc.args.book)
				if !reflect.DeepEqual(err, tc.want) {
					t.Fatalf("want %#v, but %#v", tc.want, err)
					return
				}

				if tc.args.book.UpdatedAt.IsZero() {
					t.Fatal("UpdatedAt must be not null")
					return
				}
			})
		})
	}
}

func TestBookService_UpdateBookshelf(t *testing.T) {
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
					ID:     1,
					BookID: 1,
					UserID: "00000000-0000-0000-0000-000000000000",
					Status: 5,
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

			bvm := mock_book.NewMockValidation(ctrl)

			brm := mock_book.NewMockRepository(ctrl)
			brm.EXPECT().UpdateBookshelf(ctx, tc.args.bookshelf).Return(tc.want)

			t.Run(name, func(t *testing.T) {
				target := NewBookService(bvm, brm)

				err := target.UpdateBookshelf(ctx, tc.args.bookshelf)
				if !reflect.DeepEqual(err, tc.want) {
					t.Fatalf("want %#v, but %#v", tc.want, err)
					return
				}

				if tc.args.bookshelf.UpdatedAt.IsZero() {
					t.Fatal("UpdatedAt must be not null")
					return
				}
			})
		})
	}
}

func TestBookService_MultipleCreate(t *testing.T) {
	type args struct {
		books []*book.Book
	}

	testCases := map[string]struct {
		args args
		want error
	}{
		"ok": {
			args: args{
				books: []*book.Book{
					{
						Title:        "テスト書籍",
						TitleKana:    "てすとしょせき",
						Description:  "本の説明です",
						Isbn:         "1234567890123",
						Publisher:    "テスト著者",
						PublishedOn:  "2021年12月24日",
						ThumbnailURL: "",
					},
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

			bvm := mock_book.NewMockValidation(ctrl)

			brm := mock_book.NewMockRepository(ctrl)
			brm.EXPECT().MultipleCreate(ctx, tc.args.books).Return(tc.want)
			for _, b := range tc.args.books {
				for _, a := range b.Authors {
					brm.EXPECT().ShowOrCreateAuthor(ctx, a).Return(nil)
				}
			}

			t.Run(name, func(t *testing.T) {
				target := NewBookService(bvm, brm)

				err := target.MultipleCreate(ctx, tc.args.books)
				if !reflect.DeepEqual(err, tc.want) {
					t.Fatalf("want %#v, but %#v", tc.want, err)
					return
				}

				for _, v := range tc.args.books {
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
	type args struct {
		books []*book.Book
	}

	current := time.Now().Local()
	testCases := map[string]struct {
		args args
		want error
	}{
		"ok": {
			args: args{
				books: []*book.Book{
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
						Authors: []*book.Author{{
							ID:        1,
							Name:      "テスト著者",
							NameKana:  "てすとちょしゃ",
							CreatedAt: current,
							UpdatedAt: current,
						}},
					},
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

			bvm := mock_book.NewMockValidation(ctrl)

			brm := mock_book.NewMockRepository(ctrl)
			brm.EXPECT().MultipleUpdate(ctx, tc.args.books).Return(tc.want)
			for _, b := range tc.args.books {
				for _, a := range b.Authors {
					brm.EXPECT().ShowOrCreateAuthor(ctx, a).Return(nil)
				}
			}

			t.Run(name, func(t *testing.T) {
				target := NewBookService(bvm, brm)

				err := target.MultipleUpdate(ctx, tc.args.books)
				if !reflect.DeepEqual(err, tc.want) {
					t.Fatalf("want %#v, but %#v", tc.want, err)
					return
				}

				for _, v := range tc.args.books {
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
	type args struct {
		bookID int
	}

	testCases := map[string]struct {
		args args
		want error
	}{
		"ok": {
			args: args{
				bookID: 1,
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

			bvm := mock_book.NewMockValidation(ctrl)

			brm := mock_book.NewMockRepository(ctrl)
			brm.EXPECT().Delete(ctx, tc.args.bookID).Return(tc.want)

			t.Run(name, func(t *testing.T) {
				target := NewBookService(bvm, brm)

				err := target.Delete(ctx, tc.args.bookID)
				if !reflect.DeepEqual(err, tc.want) {
					t.Fatalf("want %#v, but %#v", tc.want, err)
					return
				}
			})
		})
	}
}

func TestBookService_DeleteBookshelf(t *testing.T) {
	type args struct {
		bookshelfID int
	}

	testCases := map[string]struct {
		args args
		want error
	}{
		"ok": {
			args: args{
				bookshelfID: 1,
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

			bvm := mock_book.NewMockValidation(ctrl)

			brm := mock_book.NewMockRepository(ctrl)
			brm.EXPECT().DeleteBookshelf(ctx, tc.args.bookshelfID).Return(tc.want)

			t.Run(name, func(t *testing.T) {
				target := NewBookService(bvm, brm)

				err := target.DeleteBookshelf(ctx, tc.args.bookshelfID)
				if !reflect.DeepEqual(err, tc.want) {
					t.Fatalf("want %#v, but %#v", tc.want, err)
					return
				}
			})
		})
	}
}

func TestBookService_Validation(t *testing.T) {
	type args struct {
		book *book.Book
	}

	current := time.Now().Local()
	testCases := map[string]struct {
		args args
		want error
	}{
		"ok": {
			args: args{
				book: &book.Book{
					ID:           1,
					Title:        "テスト書籍",
					Description:  "本の説明",
					Isbn:         "978-1-234-56789-7",
					ThumbnailURL: "",
					Publisher:    "テスト出版社",
					PublishedOn:  "2021年12月24日",
					CreatedAt:    current,
					UpdatedAt:    current,
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
			want: nil,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			bvm := mock_book.NewMockValidation(ctrl)
			bvm.EXPECT().Book(ctx, tc.args.book).Return(tc.want)

			brm := mock_book.NewMockRepository(ctrl)

			t.Run(name, func(t *testing.T) {
				target := NewBookService(bvm, brm)

				err := target.Validation(ctx, tc.args.book)
				if !reflect.DeepEqual(err, tc.want) {
					t.Fatalf("want %#v, but %#v", tc.want, err)
					return
				}
			})
		})
	}
}

func TestBookService_ValidationAuthor(t *testing.T) {
	type args struct {
		author *book.Author
	}

	current := time.Now().Local()
	testCases := map[string]struct {
		args args
		want error
	}{
		"ok": {
			args: args{
				author: &book.Author{
					Name:      "テスト著者",
					CreatedAt: current,
					UpdatedAt: current,
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

			bvm := mock_book.NewMockValidation(ctrl)
			bvm.EXPECT().Author(ctx, tc.args.author).Return(tc.want)

			brm := mock_book.NewMockRepository(ctrl)

			t.Run(name, func(t *testing.T) {
				target := NewBookService(bvm, brm)

				err := target.ValidationAuthor(ctx, tc.args.author)
				if !reflect.DeepEqual(err, tc.want) {
					t.Fatalf("want %#v, but %#v", tc.want, err)
					return
				}
			})
		})
	}
}

func TestBookService_ValidationBookshelf(t *testing.T) {
	type args struct {
		bookshelf *book.Bookshelf
	}

	current := time.Now().Local()
	testCases := map[string]struct {
		args args
		want error
	}{
		"ok": {
			args: args{
				bookshelf: &book.Bookshelf{
					ID:        0,
					BookID:    1,
					UserID:    "00000000-0000-0000-0000-000000000000",
					Status:    5,
					CreatedAt: current,
					UpdatedAt: current,
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

			bvm := mock_book.NewMockValidation(ctrl)
			bvm.EXPECT().Bookshelf(ctx, tc.args.bookshelf).Return(tc.want)

			brm := mock_book.NewMockRepository(ctrl)

			t.Run(name, func(t *testing.T) {
				target := NewBookService(bvm, brm)

				err := target.ValidationBookshelf(ctx, tc.args.bookshelf)
				if !reflect.DeepEqual(err, tc.want) {
					t.Fatalf("want %#v, but %#v", tc.want, err)
					return
				}
			})
		})
	}
}

func TestBookService_ValidationReview(t *testing.T) {
	type args struct {
		review *book.Review
	}

	current := time.Now().Local()
	testCases := map[string]struct {
		args args
		want error
	}{
		"ok": {
			args: args{
				review: &book.Review{
					ID:         0,
					BookID:     1,
					UserID:     "00000000-0000-0000-0000-000000000000",
					Score:      5,
					Impression: "書籍の感想です。",
					CreatedAt:  current,
					UpdatedAt:  current,
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

			bvm := mock_book.NewMockValidation(ctrl)
			bvm.EXPECT().Review(ctx, tc.args.review).Return(tc.want)

			brm := mock_book.NewMockRepository(ctrl)

			t.Run(name, func(t *testing.T) {
				target := NewBookService(bvm, brm)

				err := target.ValidationReview(ctx, tc.args.review)
				if !reflect.DeepEqual(err, tc.want) {
					t.Fatalf("want %#v, but %#v", tc.want, err)
					return
				}
			})
		})
	}
}
