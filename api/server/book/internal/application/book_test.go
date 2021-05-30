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
	testCases := map[string]struct {
		Input    *input.ListBookByBookIDs
		Expected struct {
			Books []*book.Book
			Error error
		}
	}{
		"ok": {
			Input: &input.ListBookByBookIDs{
				BookIDs: []int{1},
			},
			Expected: struct {
				Books []*book.Book
				Error error
			}{
				Books: []*book.Book{
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
		brv.EXPECT().ListBookByBookIDs(tc.Input).Return(nil)

		bsm := mock_book.NewMockService(ctrl)
		bsm.EXPECT().List(ctx, gomock.Any()).Return(tc.Expected.Books, tc.Expected.Error)

		t.Run(result, func(t *testing.T) {
			target := NewBookApplication(brv, bsm)

			bs, err := target.ListByBookIDs(ctx, tc.Input)
			if !reflect.DeepEqual(err, tc.Expected.Error) {
				t.Fatalf("want %#v, but %#v", tc.Expected.Error, err)
				return
			}

			if !reflect.DeepEqual(bs, tc.Expected.Books) {
				t.Fatalf("want %#v, but %#v", tc.Expected.Books, bs)
				return
			}
		})
	}
}

func TestBookApplication_ListBookshelf(t *testing.T) {
	current := time.Now()

	testCases := map[string]struct {
		Input    *input.ListBookshelf
		Expected struct {
			Bookshelves []*book.Bookshelf
			Output      *output.ListQuery
			Error       error
		}
	}{
		"ok": {
			Input: &input.ListBookshelf{
				UserID: "00000000-0000-0000-0000-000000000000",
				Limit:  100,
				Offset: 0,
			},
			Expected: struct {
				Bookshelves []*book.Bookshelf
				Output      *output.ListQuery
				Error       error
			}{
				Bookshelves: []*book.Bookshelf{
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
				Output: &output.ListQuery{
					Limit:  100,
					Offset: 0,
					Total:  1,
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

		cs := []*domain.QueryCondition{
			{
				Field:    "user_id",
				Operator: "==",
				Value:    tc.Input.UserID,
			},
		}

		q := &domain.ListQuery{
			Limit:      tc.Input.Limit,
			Offset:     tc.Input.Offset,
			Conditions: cs,
		}

		brv := mock_validation.NewMockBookRequestValidation(ctrl)
		brv.EXPECT().ListBookshelf(tc.Input).Return(nil)

		bsm := mock_book.NewMockService(ctrl)
		bsm.EXPECT().ListBookshelf(ctx, q).Return(tc.Expected.Bookshelves, tc.Expected.Error)
		bsm.EXPECT().ListBookshelfCount(ctx, q).Return(tc.Expected.Output.Total, tc.Expected.Error)

		t.Run(result, func(t *testing.T) {
			target := NewBookApplication(brv, bsm)

			bss, _, err := target.ListBookshelf(ctx, tc.Input)
			if !reflect.DeepEqual(err, tc.Expected.Error) {
				t.Fatalf("want %#v, but %#v", tc.Expected.Error, err)
				return
			}

			if !reflect.DeepEqual(bss, tc.Expected.Bookshelves) {
				t.Fatalf("want %#v, but %#v", tc.Expected.Bookshelves, bss)
				return
			}
		})
	}
}

func TestBookApplication_ListBookReview(t *testing.T) {
	testCases := map[string]struct {
		Input    *input.ListBookReview
		Expected struct {
			Reviews []*book.Review
			Output  *output.ListQuery
			Error   error
		}
	}{
		"ok": {
			Input: &input.ListBookReview{
				BookID:    1,
				Limit:     0,
				Offset:    0,
				By:        "",
				Direction: "",
			},
			Expected: struct {
				Reviews []*book.Review
				Output  *output.ListQuery
				Error   error
			}{
				Reviews: []*book.Review{
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
				Output: &output.ListQuery{
					Limit:  100,
					Offset: 0,
					Total:  1,
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

		q := &domain.ListQuery{
			Limit:  tc.Input.Limit,
			Offset: tc.Input.Offset,
			Conditions: []*domain.QueryCondition{
				{
					Field:    "book_id",
					Operator: "==",
					Value:    tc.Input.BookID,
				},
			},
		}

		if tc.Input.By != "" {
			o := &domain.QueryOrder{
				By:        tc.Input.By,
				Direction: tc.Input.Direction,
			}

			q.Order = o
		}

		brv := mock_validation.NewMockBookRequestValidation(ctrl)
		brv.EXPECT().ListBookReview(tc.Input).Return(nil)

		bsm := mock_book.NewMockService(ctrl)
		bsm.EXPECT().ListReview(ctx, q).Return(tc.Expected.Reviews, tc.Expected.Error)
		bsm.EXPECT().ListReviewCount(ctx, q).Return(tc.Expected.Output.Total, tc.Expected.Error)

		t.Run(result, func(t *testing.T) {
			target := NewBookApplication(brv, bsm)

			rvs, _, err := target.ListBookReview(ctx, tc.Input)
			if !reflect.DeepEqual(err, tc.Expected.Error) {
				t.Fatalf("want %#v, but %#v", tc.Expected.Error, err)
				return
			}

			if !reflect.DeepEqual(rvs, tc.Expected.Reviews) {
				t.Fatalf("want %#v, but %#v", tc.Expected.Reviews, rvs)
				return
			}
		})
	}
}

func TestBookApplication_ListUserReview(t *testing.T) {
	testCases := map[string]struct {
		Input    *input.ListUserReview
		Expected struct {
			Reviews []*book.Review
			Output  *output.ListQuery
			Error   error
		}
	}{
		"ok": {
			Input: &input.ListUserReview{
				UserID:    "00000000-0000-0000-0000-000000000000",
				Limit:     0,
				Offset:    0,
				By:        "",
				Direction: "",
			},
			Expected: struct {
				Reviews []*book.Review
				Output  *output.ListQuery
				Error   error
			}{
				Reviews: []*book.Review{
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
				Output: &output.ListQuery{
					Limit:  100,
					Offset: 0,
					Total:  1,
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

		q := &domain.ListQuery{
			Limit:  tc.Input.Limit,
			Offset: tc.Input.Offset,
			Conditions: []*domain.QueryCondition{
				{
					Field:    "user_id",
					Operator: "==",
					Value:    tc.Input.UserID,
				},
			},
		}

		if tc.Input.By != "" {
			o := &domain.QueryOrder{
				By:        tc.Input.By,
				Direction: tc.Input.Direction,
			}

			q.Order = o
		}

		brv := mock_validation.NewMockBookRequestValidation(ctrl)
		brv.EXPECT().ListUserReview(tc.Input).Return(nil)

		bsm := mock_book.NewMockService(ctrl)
		bsm.EXPECT().ListReview(ctx, q).Return(tc.Expected.Reviews, tc.Expected.Error)
		bsm.EXPECT().ListReviewCount(ctx, q).Return(tc.Expected.Output.Total, tc.Expected.Error)

		t.Run(result, func(t *testing.T) {
			target := NewBookApplication(brv, bsm)

			rvs, _, err := target.ListUserReview(ctx, tc.Input)
			if !reflect.DeepEqual(err, tc.Expected.Error) {
				t.Fatalf("want %#v, but %#v", tc.Expected.Error, err)
				return
			}

			if !reflect.DeepEqual(rvs, tc.Expected.Reviews) {
				t.Fatalf("want %#v, but %#v", tc.Expected.Reviews, rvs)
				return
			}
		})
	}
}

func TestBookApplication_Show(t *testing.T) {
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

		bsm := mock_book.NewMockService(ctrl)
		bsm.EXPECT().Show(ctx, tc.BookID).Return(tc.Expected.Book, tc.Expected.Error)

		t.Run(result, func(t *testing.T) {
			target := NewBookApplication(brv, bsm)

			b, err := target.Show(ctx, tc.BookID)
			if !reflect.DeepEqual(err, tc.Expected.Error) {
				t.Fatalf("want %#v, but %#v", tc.Expected.Error, err)
				return
			}

			if !reflect.DeepEqual(b, tc.Expected.Book) {
				t.Fatalf("want %#v, but %#v", tc.Expected.Book, b)
				return
			}
		})
	}
}

func TestBookApplication_ShowByIsbn(t *testing.T) {
	testCases := map[string]struct {
		Isbn     string
		Expected struct {
			Book  *book.Book
			Error error
		}
	}{
		"ok": {
			Isbn: "1234567890123",
			Expected: struct {
				Book  *book.Book
				Error error
			}{
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

		bsm := mock_book.NewMockService(ctrl)
		bsm.EXPECT().ShowByIsbn(ctx, tc.Isbn).Return(tc.Expected.Book, tc.Expected.Error)

		t.Run(result, func(t *testing.T) {
			target := NewBookApplication(brv, bsm)

			b, err := target.ShowByIsbn(ctx, tc.Isbn)
			if !reflect.DeepEqual(err, tc.Expected.Error) {
				t.Fatalf("want %#v, but %#v", tc.Expected.Error, err)
				return
			}

			if !reflect.DeepEqual(b, tc.Expected.Book) {
				t.Fatalf("want %#v, but %#v", tc.Expected.Book, b)
				return
			}
		})
	}
}

func TestBookApplication_ShowBookshelf(t *testing.T) {
	current := time.Now()

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
					ID:        1,
					UserID:    "00000000-0000-0000-0000-000000000000",
					BookID:    1,
					Status:    1,
					ReadOn:    datetime.StringToDate("2020-01-01"),
					CreatedAt: current,
					UpdatedAt: current,
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

		bsm := mock_book.NewMockService(ctrl)
		bsm.EXPECT().
			ShowBookshelfByUserIDAndBookID(ctx, tc.UserID, tc.BookID).
			Return(tc.Expected.Bookshelf, tc.Expected.Error)
		bsm.EXPECT().
			ShowReviewByUserIDAndBookID(ctx, tc.UserID, tc.BookID).
			Return(tc.Expected.Bookshelf.Review, tc.Expected.Error)

		t.Run(result, func(t *testing.T) {
			target := NewBookApplication(brv, bsm)

			bs, err := target.ShowBookshelf(ctx, tc.UserID, tc.BookID)
			if !reflect.DeepEqual(err, tc.Expected.Error) {
				t.Fatalf("want %#v, but %#v", tc.Expected.Error, err)
				return
			}

			if !reflect.DeepEqual(bs, tc.Expected.Bookshelf) {
				t.Fatalf("want %#v, but %#v", tc.Expected.Bookshelf, bs)
				return
			}
		})
	}
}

func TestBookApplication_ShowReview(t *testing.T) {
	testCases := map[string]struct {
		ReviewID int
		Expected struct {
			Review *book.Review
			Error  error
		}
	}{
		"ok": {
			ReviewID: 1,
			Expected: struct {
				Review *book.Review
				Error  error
			}{
				Review: &book.Review{
					ID:         1,
					BookID:     1,
					UserID:     "00000000-0000-0000-0000-000000000000",
					Score:      5,
					Impression: "書籍の感想です。",
					CreatedAt:  time.Time{},
					UpdatedAt:  time.Time{},
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

		bsm := mock_book.NewMockService(ctrl)
		bsm.EXPECT().ShowReview(ctx, tc.ReviewID).Return(tc.Expected.Review, tc.Expected.Error)

		t.Run(result, func(t *testing.T) {
			target := NewBookApplication(brv, bsm)

			rv, err := target.ShowReview(ctx, tc.ReviewID)
			if !reflect.DeepEqual(err, tc.Expected.Error) {
				t.Fatalf("want %#v, but %#v", tc.Expected.Error, err)
				return
			}

			if !reflect.DeepEqual(rv, tc.Expected.Review) {
				t.Fatalf("want %#v, but %#v", tc.Expected.Review, rv)
				return
			}
		})
	}
}

func TestBookApplication_Create(t *testing.T) {
	testCases := map[string]struct {
		Input    *input.Book
		Expected struct {
			Book  *book.Book
			Error error
		}
	}{
		"ok": {
			Input: &input.Book{
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
			Expected: struct {
				Book  *book.Book
				Error error
			}{
				Book: &book.Book{
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
		brv.EXPECT().Book(tc.Input)

		bsm := mock_book.NewMockService(ctrl)
		bsm.EXPECT().Validation(ctx, gomock.Any()).Return(nil)
		bsm.EXPECT().ValidationAuthor(ctx, gomock.Any()).Return(nil).AnyTimes()
		bsm.EXPECT().Create(ctx, gomock.Any()).Return(tc.Expected.Error)

		t.Run(result, func(t *testing.T) {
			target := NewBookApplication(brv, bsm)

			b, err := target.Create(ctx, tc.Input)
			if !reflect.DeepEqual(err, tc.Expected.Error) {
				t.Fatalf("want %#v, but %#v", tc.Expected.Error, err)
				return
			}

			if !reflect.DeepEqual(b, tc.Expected.Book) {
				t.Fatalf("want %#v, but %#v", tc.Expected.Book, b)
				return
			}
		})
	}
}

func TestBookApplication_Update(t *testing.T) {
	current := time.Now()

	testCases := map[string]struct {
		Input    *input.Book
		Expected struct {
			Book  *book.Book
			Error error
		}
	}{
		"ok": {
			Input: &input.Book{
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
			Expected: struct {
				Book  *book.Book
				Error error
			}{
				Book: &book.Book{
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
		brv.EXPECT().Book(tc.Input)

		bsm := mock_book.NewMockService(ctrl)
		bsm.EXPECT().Validation(ctx, gomock.Any()).Return(nil)
		bsm.EXPECT().ValidationAuthor(ctx, gomock.Any()).Return(nil).AnyTimes()
		bsm.EXPECT().ShowByIsbn(ctx, tc.Input.Isbn).Return(tc.Expected.Book, nil)
		bsm.EXPECT().Update(ctx, tc.Expected.Book).Return(tc.Expected.Error)

		t.Run(result, func(t *testing.T) {
			target := NewBookApplication(brv, bsm)

			b, err := target.Update(ctx, tc.Input)
			if !reflect.DeepEqual(err, tc.Expected.Error) {
				t.Fatalf("want %#v, but %#v", tc.Expected.Error, err)
				return
			}

			if !reflect.DeepEqual(b, tc.Expected.Book) {
				t.Fatalf("want %#v, but %#v", tc.Expected.Book, b)
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
			Bookshelf *book.Bookshelf
			Error     error
		}
	}{
		"ok_create": {
			Input: &input.Bookshelf{
				UserID: "00000000-0000-0000-0000-000000000000",
				BookID: 1,
				Status: 1,
				ReadOn: "2020-01-01",
			},
			Action: "create",
			Expected: struct {
				Bookshelf *book.Bookshelf
				Error     error
			}{
				Bookshelf: &book.Bookshelf{
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
						CreatedAt:    time.Time{},
						UpdatedAt:    time.Time{},
						Authors:      []*book.Author{},
						Reviews:      []*book.Review{},
					},
				},
				Error: nil,
			},
		},
		"ok_update": {
			Input: &input.Bookshelf{
				UserID: "00000000-0000-0000-0000-000000000000",
				BookID: 1,
				Status: 1,
				ReadOn: "2020-01-01",
			},
			Action: "update",
			Expected: struct {
				Bookshelf *book.Bookshelf
				Error     error
			}{
				Bookshelf: &book.Bookshelf{
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
						CreatedAt:    time.Time{},
						UpdatedAt:    time.Time{},
						Authors:      []*book.Author{},
						Reviews:      []*book.Review{},
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
		brv.EXPECT().Bookshelf(tc.Input).Return(nil)

		bsm := mock_book.NewMockService(ctrl)
		bsm.EXPECT().Show(ctx, tc.Input.BookID).Return(tc.Expected.Bookshelf.Book, tc.Expected.Error)
		bsm.EXPECT().ValidationBookshelf(ctx, gomock.Any()).Return(tc.Expected.Error)
		bsm.EXPECT().
			ShowReviewByUserIDAndBookID(ctx, tc.Input.UserID, tc.Input.BookID).
			Return(tc.Expected.Bookshelf.Review, tc.Expected.Error)

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

			bs, err := target.CreateOrUpdateBookshelf(ctx, tc.Input)
			if !reflect.DeepEqual(err, tc.Expected.Error) {
				t.Fatalf("want %#v, but %#v", tc.Expected.Error, err)
				return
			}

			if !reflect.DeepEqual(bs, tc.Expected.Bookshelf) {
				t.Fatalf("want %#v, but %#v", tc.Expected.Bookshelf, bs)
				return
			}
		})
	}
}

func TestBookApplication_Delete(t *testing.T) {
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
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		b := &book.Book{
			ID: tc.BookID,
		}

		brv := mock_validation.NewMockBookRequestValidation(ctrl)

		bsm := mock_book.NewMockService(ctrl)
		bsm.EXPECT().Show(ctx, tc.BookID).Return(b, nil)
		bsm.EXPECT().Delete(ctx, tc.BookID).Return(tc.Expected)

		t.Run(result, func(t *testing.T) {
			target := NewBookApplication(brv, bsm)

			got := target.Delete(ctx, tc.BookID)
			if !reflect.DeepEqual(got, tc.Expected) {
				t.Fatalf("want %#v, but %#v", tc.Expected, got)
				return
			}
		})
	}
}

func TestBookApplication_DeleteBookshelf(t *testing.T) {
	testCases := map[string]struct {
		BookID   int
		UID      string
		Expected error
	}{
		"ok": {
			BookID:   1,
			UID:      "00000000-0000-0000-0000-000000000000",
			Expected: nil,
		},
	}

	for result, tc := range testCases {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		b := &book.Bookshelf{
			ID: 1,
		}

		brv := mock_validation.NewMockBookRequestValidation(ctrl)

		bsm := mock_book.NewMockService(ctrl)
		bsm.EXPECT().ShowBookshelfByUserIDAndBookID(ctx, tc.UID, tc.BookID).Return(b, nil)
		bsm.EXPECT().DeleteBookshelf(ctx, b.ID).Return(tc.Expected)

		t.Run(result, func(t *testing.T) {
			target := NewBookApplication(brv, bsm)

			got := target.DeleteBookshelf(ctx, tc.BookID, tc.UID)
			if !reflect.DeepEqual(got, tc.Expected) {
				t.Fatalf("want %#v, but %#v", tc.Expected, got)
				return
			}
		})
	}
}
