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

func TestBookApplication_Show(t *testing.T) {
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
					Bookshelf:    &book.Bookshelf{},
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

			b, err := target.Show(ctx, tc.Isbn)
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
			Book      *book.Book
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
				Book      *book.Book
				Bookshelf *book.Bookshelf
				Error     error
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
				Bookshelf: &book.Bookshelf{
					UserID: "00000000-0000-0000-0000-000000000000",
					BookID: 1,
					Status: 1,
					ReadOn: datetime.StringToDate("2020-01-01"),
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
				Book      *book.Book
				Bookshelf *book.Bookshelf
				Error     error
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
