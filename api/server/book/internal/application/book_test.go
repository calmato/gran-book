package application

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/calmato/gran-book/api/server/book/internal/application/input"
	"github.com/calmato/gran-book/api/server/book/internal/application/output"
	"github.com/calmato/gran-book/api/server/book/internal/domain"
	"github.com/calmato/gran-book/api/server/book/internal/domain/book"
	"github.com/calmato/gran-book/api/server/book/lib/datetime"
	mock_validation "github.com/calmato/gran-book/api/server/book/mock/application/validation"
	mock_book "github.com/calmato/gran-book/api/server/book/mock/domain/book"
	"github.com/golang/mock/gomock"
)

func TestBookApplication_ListByBookIDs(t *testing.T) {
	type args struct {
		input *input.ListBookByBookIDs
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
				input: &input.ListBookByBookIDs{
					BookIDs: []int{1},
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
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		brv := mock_validation.NewMockBookRequestValidation(ctrl)
		brv.EXPECT().ListBookByBookIDs(tc.args.input).Return(nil)

		bsm := mock_book.NewMockService(ctrl)
		bsm.EXPECT().List(ctx, gomock.Any()).Return(tc.want.books, tc.want.err)

		t.Run(name, func(t *testing.T) {
			target := NewBookApplication(brv, bsm)

			bs, err := target.ListByBookIDs(ctx, tc.args.input)
			if !reflect.DeepEqual(err, tc.want.err) {
				t.Fatalf("want %#v, but %#v", tc.want.err, err)
				return
			}

			if !reflect.DeepEqual(bs, tc.want.books) {
				t.Fatalf("want %#v, but %#v", tc.want.books, bs)
				return
			}
		})
	}
}

func TestBookApplication_ListBookshelf(t *testing.T) {
	type args struct {
		input *input.ListBookshelf
	}
	type want struct {
		bookshelves []*book.Bookshelf
		output      *output.ListQuery
		err         error
	}

	current := time.Now().Local()
	testCases := map[string]struct {
		args args
		want want
	}{
		"ok": {
			args: args{
				input: &input.ListBookshelf{
					UserID: "00000000-0000-0000-0000-000000000000",
					Limit:  100,
					Offset: 0,
				},
			},
			want: want{
				bookshelves: []*book.Bookshelf{
					{
						ID:        1,
						UserID:    "00000000-0000-0000-0000-000000000000",
						BookID:    1,
						Status:    1,
						ReadOn:    datetime.StringToDate("2020-01-01"),
						CreatedAt: current,
						UpdatedAt: current,
					},
				},
				output: &output.ListQuery{
					Limit:  100,
					Offset: 0,
					Total:  1,
				},
				err: nil,
			},
		},
	}

	for name, tc := range testCases {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		cs := []*domain.QueryCondition{
			{
				Field:    "user_id",
				Operator: "==",
				Value:    tc.args.input.UserID,
			},
		}

		q := &domain.ListQuery{
			Limit:      tc.args.input.Limit,
			Offset:     tc.args.input.Offset,
			Conditions: cs,
		}

		brv := mock_validation.NewMockBookRequestValidation(ctrl)
		brv.EXPECT().ListBookshelf(tc.args.input).Return(nil)

		bsm := mock_book.NewMockService(ctrl)
		bsm.EXPECT().ListBookshelf(ctx, q).Return(tc.want.bookshelves, tc.want.err)
		bsm.EXPECT().ListBookshelfCount(ctx, q).Return(tc.want.output.Total, nil)

		t.Run(name, func(t *testing.T) {
			target := NewBookApplication(brv, bsm)

			bss, _, err := target.ListBookshelf(ctx, tc.args.input)
			if !reflect.DeepEqual(err, tc.want.err) {
				t.Fatalf("want %#v, but %#v", tc.want.err, err)
				return
			}

			if !reflect.DeepEqual(bss, tc.want.bookshelves) {
				t.Fatalf("want %#v, but %#v", tc.want.bookshelves, bss)
				return
			}
		})
	}
}

func TestBookApplication_ListBookReview(t *testing.T) {
	type args struct {
		input *input.ListBookReview
	}
	type want struct {
		reviews []*book.Review
		output  *output.ListQuery
		err     error
	}

	testCases := map[string]struct {
		args args
		want want
	}{
		"ok": {
			args: args{
				input: &input.ListBookReview{
					BookID:    1,
					Limit:     0,
					Offset:    0,
					By:        "",
					Direction: "",
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
				output: &output.ListQuery{
					Limit:  100,
					Offset: 0,
					Total:  1,
				},
				err: nil,
			},
		},
	}

	for name, tc := range testCases {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		q := &domain.ListQuery{
			Limit:  tc.args.input.Limit,
			Offset: tc.args.input.Offset,
			Conditions: []*domain.QueryCondition{
				{
					Field:    "book_id",
					Operator: "==",
					Value:    tc.args.input.BookID,
				},
			},
		}

		if tc.args.input.By != "" {
			o := &domain.QueryOrder{
				By:        tc.args.input.By,
				Direction: tc.args.input.Direction,
			}

			q.Order = o
		}

		brv := mock_validation.NewMockBookRequestValidation(ctrl)
		brv.EXPECT().ListBookReview(tc.args.input).Return(nil)

		bsm := mock_book.NewMockService(ctrl)
		bsm.EXPECT().ListReview(ctx, q).Return(tc.want.reviews, tc.want.err)
		bsm.EXPECT().ListReviewCount(ctx, q).Return(tc.want.output.Total, tc.want.err)

		t.Run(name, func(t *testing.T) {
			target := NewBookApplication(brv, bsm)

			rvs, _, err := target.ListBookReview(ctx, tc.args.input)
			if !reflect.DeepEqual(err, tc.want.err) {
				t.Fatalf("want %#v, but %#v", tc.want.err, err)
				return
			}

			if !reflect.DeepEqual(rvs, tc.want.reviews) {
				t.Fatalf("want %#v, but %#v", tc.want.reviews, rvs)
				return
			}
		})
	}
}

func TestBookApplication_ListUserReview(t *testing.T) {
	type args struct {
		input *input.ListUserReview
	}
	type want struct {
		reviews []*book.Review
		output  *output.ListQuery
		err     error
	}

	testCases := map[string]struct {
		args args
		want want
	}{
		"ok": {
			args: args{
				input: &input.ListUserReview{
					UserID:    "00000000-0000-0000-0000-000000000000",
					Limit:     0,
					Offset:    0,
					By:        "",
					Direction: "",
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
				output: &output.ListQuery{
					Limit:  100,
					Offset: 0,
					Total:  1,
				},
				err: nil,
			},
		},
	}

	for name, tc := range testCases {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		q := &domain.ListQuery{
			Limit:  tc.args.input.Limit,
			Offset: tc.args.input.Offset,
			Conditions: []*domain.QueryCondition{
				{
					Field:    "user_id",
					Operator: "==",
					Value:    tc.args.input.UserID,
				},
			},
		}

		if tc.args.input.By != "" {
			o := &domain.QueryOrder{
				By:        tc.args.input.By,
				Direction: tc.args.input.Direction,
			}

			q.Order = o
		}

		brv := mock_validation.NewMockBookRequestValidation(ctrl)
		brv.EXPECT().ListUserReview(tc.args.input).Return(nil)

		bsm := mock_book.NewMockService(ctrl)
		bsm.EXPECT().ListReview(ctx, q).Return(tc.want.reviews, tc.want.err)
		bsm.EXPECT().ListReviewCount(ctx, q).Return(tc.want.output.Total, tc.want.err)

		t.Run(name, func(t *testing.T) {
			target := NewBookApplication(brv, bsm)

			rvs, _, err := target.ListUserReview(ctx, tc.args.input)
			if !reflect.DeepEqual(err, tc.want.err) {
				t.Fatalf("want %#v, but %#v", tc.want.err, err)
				return
			}

			if !reflect.DeepEqual(rvs, tc.want.reviews) {
				t.Fatalf("want %#v, but %#v", tc.want.reviews, rvs)
				return
			}
		})
	}
}

func TestBookApplication_Show(t *testing.T) {
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
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		brv := mock_validation.NewMockBookRequestValidation(ctrl)

		bsm := mock_book.NewMockService(ctrl)
		bsm.EXPECT().Show(ctx, tc.want.book.ID).Return(tc.want.book, tc.want.err)

		t.Run(name, func(t *testing.T) {
			target := NewBookApplication(brv, bsm)

			b, err := target.Show(ctx, tc.args.bookID)
			if !reflect.DeepEqual(err, tc.want.err) {
				t.Fatalf("want %#v, but %#v", tc.want.err, err)
				return
			}

			if !reflect.DeepEqual(b, tc.want.book) {
				t.Fatalf("want %#v, but %#v", tc.want.book, b)
				return
			}
		})
	}
}

func TestBookApplication_ShowByIsbn(t *testing.T) {
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
				isbn: "1234567890123",
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
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		brv := mock_validation.NewMockBookRequestValidation(ctrl)

		bsm := mock_book.NewMockService(ctrl)
		bsm.EXPECT().ShowByIsbn(ctx, tc.want.book.Isbn).Return(tc.want.book, tc.want.err)

		t.Run(name, func(t *testing.T) {
			target := NewBookApplication(brv, bsm)

			b, err := target.ShowByIsbn(ctx, tc.args.isbn)
			if !reflect.DeepEqual(err, tc.want.err) {
				t.Fatalf("want %#v, but %#v", tc.want.err, err)
				return
			}

			if !reflect.DeepEqual(b, tc.want.book) {
				t.Fatalf("want %#v, but %#v", tc.want.book, b)
				return
			}
		})
	}
}

func TestBookApplication_ShowBookshelf(t *testing.T) {
	type args struct {
		userID string
		bookID int
	}
	type want struct {
		bookshelf *book.Bookshelf
		err       error
	}

	current := time.Now().Local()
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
					ID:        1,
					UserID:    "00000000-0000-0000-0000-000000000000",
					BookID:    1,
					Status:    1,
					ReadOn:    datetime.StringToDate("2020-01-01"),
					CreatedAt: current,
					UpdatedAt: current,
				},
				err: nil,
			},
		},
	}

	for name, tc := range testCases {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		brv := mock_validation.NewMockBookRequestValidation(ctrl)

		bsm := mock_book.NewMockService(ctrl)
		bsm.EXPECT().
			ShowBookshelfByUserIDAndBookID(ctx, tc.want.bookshelf.UserID, tc.want.bookshelf.BookID).
			Return(tc.want.bookshelf, tc.want.err)
		bsm.EXPECT().
			ShowReviewByUserIDAndBookID(ctx, tc.want.bookshelf.UserID, tc.want.bookshelf.BookID).
			Return(tc.want.bookshelf.Review, tc.want.err)

		t.Run(name, func(t *testing.T) {
			target := NewBookApplication(brv, bsm)

			bs, err := target.ShowBookshelf(ctx, tc.args.userID, tc.args.bookID)
			if !reflect.DeepEqual(err, tc.want.err) {
				t.Fatalf("want %#v, but %#v", tc.want.err, err)
				return
			}

			if !reflect.DeepEqual(bs, tc.want.bookshelf) {
				t.Fatalf("want %#v, but %#v", tc.want.bookshelf, bs)
				return
			}
		})
	}
}

func TestBookApplication_ShowReview(t *testing.T) {
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
					Score:      5,
					Impression: "書籍の感想です。",
				},
				err: nil,
			},
		},
	}

	for name, tc := range testCases {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		brv := mock_validation.NewMockBookRequestValidation(ctrl)

		bsm := mock_book.NewMockService(ctrl)
		bsm.EXPECT().ShowReview(ctx, tc.want.review.ID).Return(tc.want.review, tc.want.err)

		t.Run(name, func(t *testing.T) {
			target := NewBookApplication(brv, bsm)

			rv, err := target.ShowReview(ctx, tc.args.reviewID)
			if !reflect.DeepEqual(err, tc.want.err) {
				t.Fatalf("want %#v, but %#v", tc.want.err, err)
				return
			}

			if !reflect.DeepEqual(rv, tc.want.review) {
				t.Fatalf("want %#v, but %#v", tc.want.review, rv)
				return
			}
		})
	}
}

func TestBookApplication_Create(t *testing.T) {
	type args struct {
		input *input.Book
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
				input: &input.Book{
					Title:          "小説　ちはやふる　上の句",
					TitleKana:      "ショウセツ チハヤフルカミノク",
					Description:    "綾瀬千早は高校入学と同時に、競技かるた部を作ろうと奔走する。幼馴染の太一と仲間を集め、夏の全国大会に出場するためだ。強くなって、新と再会したい。幼い頃かるたを取り合った、新に寄せる千早の秘めた想いに気づきながらも、太一は千早を守り立てる。それぞれの青春を懸けた、一途な情熱の物語が幕開ける。",
					Isbn:           "9784062938426",
					Publisher:      "講談社",
					PublishedOn:    "2021年12月24日",
					ThumbnailURL:   "https://thumbnail.image.rakuten.co.jp/@0_mall/book/cabinet/8426/9784062938426.jpg?_ex=120x120",
					RakutenURL:     "https://books.rakuten.co.jp/rb/15271426/",
					RakutenSize:    "コミック",
					RakutenGenreID: "001004008001/001004008003/001019001",
					Authors: []*input.BookAuthor{
						{
							Name:     "有沢 ゆう希",
							NameKana: "アリサワ ユウキ",
						},
						{
							Name:     "末次 由紀",
							NameKana: "スエツグ ユキ",
						},
					},
				},
			},
			want: want{
				book: &book.Book{
					Title:          "小説　ちはやふる　上の句",
					TitleKana:      "ショウセツ チハヤフルカミノク",
					Description:    "綾瀬千早は高校入学と同時に、競技かるた部を作ろうと奔走する。幼馴染の太一と仲間を集め、夏の全国大会に出場するためだ。強くなって、新と再会したい。幼い頃かるたを取り合った、新に寄せる千早の秘めた想いに気づきながらも、太一は千早を守り立てる。それぞれの青春を懸けた、一途な情熱の物語が幕開ける。",
					Isbn:           "9784062938426",
					Publisher:      "講談社",
					PublishedOn:    "2021年12月24日",
					ThumbnailURL:   "https://thumbnail.image.rakuten.co.jp/@0_mall/book/cabinet/8426/9784062938426.jpg?_ex=120x120",
					RakutenURL:     "https://books.rakuten.co.jp/rb/15271426/",
					RakutenSize:    "コミック",
					RakutenGenreID: "001004008001/001004008003/001019001",
					Authors: []*book.Author{
						{
							Name:     "有沢 ゆう希",
							NameKana: "アリサワ ユウキ",
						},
						{
							Name:     "末次 由紀",
							NameKana: "スエツグ ユキ",
						},
					},
				},
				err: nil,
			},
		},
	}

	for name, tc := range testCases {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		brv := mock_validation.NewMockBookRequestValidation(ctrl)
		brv.EXPECT().Book(tc.args.input)

		bsm := mock_book.NewMockService(ctrl)
		bsm.EXPECT().Validation(ctx, gomock.Any()).Return(nil)
		bsm.EXPECT().ValidationAuthor(ctx, gomock.Any()).Return(nil).AnyTimes()
		bsm.EXPECT().Create(ctx, gomock.Any()).Return(tc.want.err)

		t.Run(name, func(t *testing.T) {
			target := NewBookApplication(brv, bsm)

			b, err := target.Create(ctx, tc.args.input)
			if !reflect.DeepEqual(err, tc.want.err) {
				t.Fatalf("want %#v, but %#v", tc.want.err, err)
				return
			}

			if !reflect.DeepEqual(b, tc.want.book) {
				t.Fatalf("want %#v, but %#v", tc.want.book, b)
				return
			}
		})
	}
}

func TestBookApplication_Update(t *testing.T) {
	type args struct {
		input *input.Book
	}
	type want struct {
		book *book.Book
		err  error
	}

	current := time.Now().Local()
	testCases := map[string]struct {
		args args
		want want
	}{
		"ok": {
			args: args{
				input: &input.Book{
					Title:          "小説　ちはやふる　上の句",
					TitleKana:      "ショウセツ チハヤフルカミノク",
					Description:    "綾瀬千早は高校入学と同時に、競技かるた部を作ろうと奔走する。幼馴染の太一と仲間を集め、夏の全国大会に出場するためだ。強くなって、新と再会したい。幼い頃かるたを取り合った、新に寄せる千早の秘めた想いに気づきながらも、太一は千早を守り立てる。それぞれの青春を懸けた、一途な情熱の物語が幕開ける。",
					Isbn:           "9784062938426",
					Publisher:      "講談社",
					PublishedOn:    "2018-01-16",
					ThumbnailURL:   "https://thumbnail.image.rakuten.co.jp/@0_mall/book/cabinet/8426/9784062938426.jpg?_ex=120x120",
					RakutenURL:     "https://books.rakuten.co.jp/rb/15271426/",
					RakutenSize:    "コミック",
					RakutenGenreID: "001004008001/001004008003/001019001",
					Authors: []*input.BookAuthor{
						{
							Name:     "有沢 ゆう希",
							NameKana: "アリサワ ユウキ",
						},
						{
							Name:     "末次 由紀",
							NameKana: "スエツグ ユキ",
						},
					},
				},
			},
			want: want{
				book: &book.Book{
					ID:             1,
					Title:          "小説　ちはやふる　上の句",
					TitleKana:      "ショウセツ チハヤフルカミノク",
					Description:    "綾瀬千早は高校入学と同時に、競技かるた部を作ろうと奔走する。幼馴染の太一と仲間を集め、夏の全国大会に出場するためだ。強くなって、新と再会したい。幼い頃かるたを取り合った、新に寄せる千早の秘めた想いに気づきながらも、太一は千早を守り立てる。それぞれの青春を懸けた、一途な情熱の物語が幕開ける。",
					Isbn:           "9784062938426",
					Publisher:      "講談社",
					PublishedOn:    "2021年12月24日",
					ThumbnailURL:   "https://thumbnail.image.rakuten.co.jp/@0_mall/book/cabinet/8426/9784062938426.jpg?_ex=120x120",
					RakutenURL:     "https://books.rakuten.co.jp/rb/15271426/",
					RakutenSize:    "コミック",
					RakutenGenreID: "001004008001/001004008003/001019001",
					CreatedAt:      current,
					UpdatedAt:      current,
					Authors: []*book.Author{
						{
							ID:        1,
							Name:      "有沢 ゆう希",
							NameKana:  "アリサワ ユウキ",
							CreatedAt: current,
							UpdatedAt: current,
						},
						{
							ID:        2,
							Name:      "末次 由紀",
							NameKana:  "スエツグ ユキ",
							CreatedAt: current,
							UpdatedAt: current,
						},
					},
				},
				err: nil,
			},
		},
	}

	for name, tc := range testCases {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		brv := mock_validation.NewMockBookRequestValidation(ctrl)
		brv.EXPECT().Book(tc.args.input)

		bsm := mock_book.NewMockService(ctrl)
		bsm.EXPECT().Validation(ctx, gomock.Any()).Return(nil)
		bsm.EXPECT().ValidationAuthor(ctx, gomock.Any()).Return(nil).AnyTimes()
		bsm.EXPECT().ShowByIsbn(ctx, tc.want.book.Isbn).Return(tc.want.book, nil)
		bsm.EXPECT().Update(ctx, tc.want.book).Return(tc.want.err)

		t.Run(name, func(t *testing.T) {
			target := NewBookApplication(brv, bsm)

			b, err := target.Update(ctx, tc.args.input)
			if !reflect.DeepEqual(err, tc.want.err) {
				t.Fatalf("want %#v, but %#v", tc.want.err, err)
				return
			}

			if !reflect.DeepEqual(b, tc.want.book) {
				t.Fatalf("want %#v, but %#v", tc.want.book, b)
				return
			}
		})
	}
}

func TestBookApplication_CreateOrUpdateBookshelf(t *testing.T) {
	type args struct {
		input *input.Bookshelf
	}
	type want struct {
		bookshelf *book.Bookshelf
		err       error
	}

	current := time.Now().Local()
	testCases := map[string]struct {
		action string
		args   args
		want   want
	}{
		"ok_create": {
			action: "create",
			args: args{
				input: &input.Bookshelf{
					UserID: "00000000-0000-0000-0000-000000000000",
					BookID: 1,
					Status: 1,
					ReadOn: "2020-01-01",
				},
			},
			want: want{
				bookshelf: &book.Bookshelf{
					UserID: "00000000-0000-0000-0000-000000000000",
					BookID: 1,
					Status: 1,
					ReadOn: datetime.StringToDate("2020-01-01"),
					Book: &book.Book{
						ID:           1,
						Title:        "テスト書籍",
						TitleKana:    "てすとしょせき",
						Description:  "本の説明です",
						Isbn:         "1234567890123",
						Publisher:    "テスト著者",
						PublishedOn:  "2021年12月24日",
						ThumbnailURL: "",
					},
				},
				err: nil,
			},
		},
		"ok_update": {
			action: "update",
			args: args{
				input: &input.Bookshelf{
					UserID: "00000000-0000-0000-0000-000000000000",
					BookID: 1,
					Status: 1,
					ReadOn: "2020-01-01",
				},
			},
			want: want{
				bookshelf: &book.Bookshelf{
					ID:        1,
					UserID:    "00000000-0000-0000-0000-000000000000",
					BookID:    1,
					Status:    1,
					ReadOn:    datetime.StringToDate("2020-01-01"),
					CreatedAt: current,
					UpdatedAt: current,
					Book: &book.Book{
						ID:           1,
						Title:        "テスト書籍",
						TitleKana:    "てすとしょせき",
						Description:  "本の説明です",
						Isbn:         "1234567890123",
						Publisher:    "テスト著者",
						PublishedOn:  "2021年12月24日",
						ThumbnailURL: "",
					},
				},
				err: nil,
			},
		},
	}

	for name, tc := range testCases {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		brv := mock_validation.NewMockBookRequestValidation(ctrl)
		brv.EXPECT().Bookshelf(tc.args.input).Return(nil)

		bsm := mock_book.NewMockService(ctrl)
		bsm.EXPECT().Show(ctx, tc.want.bookshelf.BookID).Return(tc.want.bookshelf.Book, nil)
		bsm.EXPECT().ValidationBookshelf(ctx, gomock.Any()).Return(nil)
		bsm.EXPECT().
			ShowReviewByUserIDAndBookID(ctx, tc.want.bookshelf.UserID, tc.want.bookshelf.BookID).
			Return(tc.want.bookshelf.Review, nil)

		switch tc.action {
		case "create":
			bsm.EXPECT().
				ShowBookshelfByUserIDAndBookID(ctx, tc.want.bookshelf.UserID, tc.want.bookshelf.BookID).
				Return(nil, nil)
			bsm.EXPECT().CreateBookshelf(ctx, gomock.Any()).Return(tc.want.err)
		case "update":
			bsm.EXPECT().
				ShowBookshelfByUserIDAndBookID(ctx, tc.want.bookshelf.UserID, tc.want.bookshelf.BookID).
				Return(tc.want.bookshelf, nil)
			bsm.EXPECT().UpdateBookshelf(ctx, tc.want.bookshelf).Return(tc.want.err)
		}

		t.Run(name, func(t *testing.T) {
			target := NewBookApplication(brv, bsm)

			bs, err := target.CreateOrUpdateBookshelf(ctx, tc.args.input)
			if !reflect.DeepEqual(err, tc.want.err) {
				t.Fatalf("want %#v, but %#v", tc.want.err, err)
				return
			}

			if !reflect.DeepEqual(bs, tc.want.bookshelf) {
				t.Fatalf("want %#v, but %#v", tc.want.bookshelf, bs)
				return
			}
		})
	}
}

func TestBookApplication_Delete(t *testing.T) {
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
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		b := &book.Book{
			ID: tc.args.bookID,
		}

		brv := mock_validation.NewMockBookRequestValidation(ctrl)

		bsm := mock_book.NewMockService(ctrl)
		bsm.EXPECT().Show(ctx, tc.args.bookID).Return(b, nil)
		bsm.EXPECT().Delete(ctx, tc.args.bookID).Return(tc.want)

		t.Run(name, func(t *testing.T) {
			target := NewBookApplication(brv, bsm)

			got := target.Delete(ctx, tc.args.bookID)
			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("want %#v, but %#v", tc.want, got)
				return
			}
		})
	}
}

func TestBookApplication_DeleteBookshelf(t *testing.T) {
	type args struct {
		bookID int
		uid    string
	}

	testCases := map[string]struct {
		args args
		want error
	}{
		"ok": {
			args: args{
				bookID: 1,
				uid:    "00000000-0000-0000-0000-000000000000",
			},
			want: nil,
		},
	}

	for name, tc := range testCases {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		b := &book.Bookshelf{
			ID: 1,
		}

		brv := mock_validation.NewMockBookRequestValidation(ctrl)

		bsm := mock_book.NewMockService(ctrl)
		bsm.EXPECT().ShowBookshelfByUserIDAndBookID(ctx, tc.args.uid, tc.args.bookID).Return(b, nil)
		bsm.EXPECT().DeleteBookshelf(ctx, b.ID).Return(tc.want)

		t.Run(name, func(t *testing.T) {
			target := NewBookApplication(brv, bsm)

			got := target.DeleteBookshelf(ctx, tc.args.bookID, tc.args.uid)
			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("want %#v, but %#v", tc.want, got)
				return
			}
		})
	}
}
