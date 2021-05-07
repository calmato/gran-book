package validation

import (
	"strings"
	"testing"

	"github.com/calmato/gran-book/api/server/book/internal/application/input"
	"github.com/golang/mock/gomock"
)

func TestBookRequestValidation_Book(t *testing.T) {
	testCases := map[string]struct {
		Input    *input.Book
		Expected bool
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
			Expected: true,
		},
		"ng_title_required": {
			Input: &input.Book{
				Title:          "",
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
			Expected: false,
		},
		"ng_title_max": {
			Input: &input.Book{
				Title:          strings.Repeat("x", 65),
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
			Expected: false,
		},
		"ng_titleKana_required": {
			Input: &input.Book{
				Title:          "小説　ちはやふる　上の句",
				TitleKana:      "",
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
			Expected: false,
		},
		"ng_titleKana_max": {
			Input: &input.Book{
				Title:          "小説　ちはやふる　上の句",
				TitleKana:      strings.Repeat("x", 129),
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
			Expected: false,
		},
		"ng_description_max": {
			Input: &input.Book{
				Title:          "小説　ちはやふる　上の句",
				TitleKana:      "ショウセツ チハヤフルカミノク",
				Description:    strings.Repeat("x", 2001),
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
			Expected: false,
		},
		"ng_isbn_required": {
			Input: &input.Book{
				Title:          "小説　ちはやふる　上の句",
				TitleKana:      "ショウセツ チハヤフルカミノク",
				Description:    "綾瀬千早は高校入学と同時に、競技かるた部を作ろうと奔走する。幼馴染の太一と仲間を集め、夏の全国大会に出場するためだ。強くなって、新と再会したい。幼い頃かるたを取り合った、新に寄せる千早の秘めた想いに気づきながらも、太一は千早を守り立てる。それぞれの青春を懸けた、一途な情熱の物語が幕開ける。",
				Isbn:           "",
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
			Expected: false,
		},
		"ng_isbn_max": {
			Input: &input.Book{
				Title:          "小説　ちはやふる　上の句",
				TitleKana:      "ショウセツ チハヤフルカミノク",
				Description:    "綾瀬千早は高校入学と同時に、競技かるた部を作ろうと奔走する。幼馴染の太一と仲間を集め、夏の全国大会に出場するためだ。強くなって、新と再会したい。幼い頃かるたを取り合った、新に寄せる千早の秘めた想いに気づきながらも、太一は千早を守り立てる。それぞれの青春を懸けた、一途な情熱の物語が幕開ける。",
				Isbn:           strings.Repeat("x", 17),
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
			Expected: false,
		},
		"ng_publisher_required": {
			Input: &input.Book{
				Title:          "小説　ちはやふる　上の句",
				TitleKana:      "ショウセツ チハヤフルカミノク",
				Description:    "綾瀬千早は高校入学と同時に、競技かるた部を作ろうと奔走する。幼馴染の太一と仲間を集め、夏の全国大会に出場するためだ。強くなって、新と再会したい。幼い頃かるたを取り合った、新に寄せる千早の秘めた想いに気づきながらも、太一は千早を守り立てる。それぞれの青春を懸けた、一途な情熱の物語が幕開ける。",
				Isbn:           "9784062938426",
				Publisher:      "",
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
			Expected: false,
		},
		"ng_publisher_max": {
			Input: &input.Book{
				Title:          "小説　ちはやふる　上の句",
				TitleKana:      "ショウセツ チハヤフルカミノク",
				Description:    "綾瀬千早は高校入学と同時に、競技かるた部を作ろうと奔走する。幼馴染の太一と仲間を集め、夏の全国大会に出場するためだ。強くなって、新と再会したい。幼い頃かるたを取り合った、新に寄せる千早の秘めた想いに気づきながらも、太一は千早を守り立てる。それぞれの青春を懸けた、一途な情熱の物語が幕開ける。",
				Isbn:           "9784062938426",
				Publisher:      strings.Repeat("x", 33),
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
			Expected: false,
		},
		"ng_publishedOn_required": {
			Input: &input.Book{
				Title:          "小説　ちはやふる　上の句",
				TitleKana:      "ショウセツ チハヤフルカミノク",
				Description:    "綾瀬千早は高校入学と同時に、競技かるた部を作ろうと奔走する。幼馴染の太一と仲間を集め、夏の全国大会に出場するためだ。強くなって、新と再会したい。幼い頃かるたを取り合った、新に寄せる千早の秘めた想いに気づきながらも、太一は千早を守り立てる。それぞれの青春を懸けた、一途な情熱の物語が幕開ける。",
				Isbn:           "9784062938426",
				Publisher:      "講談社",
				PublishedOn:    "",
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
			Expected: false,
		},
		"ng_authors_name_required": {
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
						Name:     "",
						NameKana: "スエツグ ユキ",
					},
				},
			},
			Expected: false,
		},
		"ng_authors_name_max": {
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
						Name:     strings.Repeat("x", 33),
						NameKana: "アリサワ ユウキ",
					},
					{
						Name:     "末次 由紀",
						NameKana: "スエツグ ユキ",
					},
				},
			},
			Expected: false,
		},
		"ng_authors_nameKana_required": {
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
						NameKana: "",
					},
				},
			},
			Expected: false,
		},
		"ng_authors_nameKana_max": {
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
						NameKana: strings.Repeat("x", 65),
					},
					{
						Name:     "末次 由紀",
						NameKana: "スエツグ ユキ",
					},
				},
			},
			Expected: false,
		},
	}

	for result, tc := range testCases {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		t.Run(result, func(t *testing.T) {
			target := NewBookRequestValidation()

			got := target.Book(tc.Input)
			if tc.Expected {
				if got != nil {
					t.Fatalf("Incorrect result: %#v", got)
				}
			} else {
				if got == nil {
					t.Fatalf("Incorrect result: result is nil")
				}
			}
		})
	}
}

func TestBookRequestValidation_Bookshelf(t *testing.T) {
	testCases := map[string]struct {
		Input    *input.Bookshelf
		Expected bool
	}{
		"ok": {
			Input: &input.Bookshelf{
				UserID:     "00000000-0000-0000-0000-000000000000",
				BookID:     1,
				Status:     1,
				ReadOn:     "2020-01-01",
				Impression: "本の感想です。",
			},
			Expected: true,
		},
		"ng_userId_required": {
			Input: &input.Bookshelf{
				UserID:     "",
				BookID:     1,
				Status:     1,
				ReadOn:     "2020-01-01",
				Impression: "本の感想です。",
			},
			Expected: false,
		},
		"ng_bookId_required": {
			Input: &input.Bookshelf{
				UserID:     "00000000-0000-0000-0000-000000000000",
				BookID:     0,
				Status:     1,
				ReadOn:     "2020-01-01",
				Impression: "本の感想です。",
			},
			Expected: false,
		},
		"ng_bookId_greater_than_equal": {
			Input: &input.Bookshelf{
				UserID:     "00000000-0000-0000-0000-000000000000",
				BookID:     0,
				Status:     1,
				ReadOn:     "2020-01-01",
				Impression: "本の感想です。",
			},
			Expected: false,
		},
		"ng_status_greater_than_equal": {
			Input: &input.Bookshelf{
				UserID:     "00000000-0000-0000-0000-000000000000",
				BookID:     1,
				Status:     -1,
				ReadOn:     "2020-01-01",
				Impression: "本の感想です。",
			},
			Expected: false,
		},
		"ng_status_less_than_equal": {
			Input: &input.Bookshelf{
				UserID:     "00000000-0000-0000-0000-000000000000",
				BookID:     1,
				Status:     6,
				ReadOn:     "2020-01-01",
				Impression: "本の感想です。",
			},
			Expected: false,
		},
		"ng_impression_max": {
			Input: &input.Bookshelf{
				UserID:     "00000000-0000-0000-0000-000000000000",
				BookID:     1,
				Status:     6,
				ReadOn:     "2020-01-01",
				Impression: strings.Repeat("x", 1001),
			},
			Expected: false,
		},
	}

	for result, tc := range testCases {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		t.Run(result, func(t *testing.T) {
			target := NewBookRequestValidation()

			got := target.Bookshelf(tc.Input)
			if tc.Expected {
				if got != nil {
					t.Fatalf("Incorrect result: %#v", got)
				}
			} else {
				if got == nil {
					t.Fatalf("Incorrect result: result is nil")
				}
			}
		})
	}
}

func TestBookRequestValidation_ListBookByBookIDs(t *testing.T) {
	testCases := map[string]struct {
		Input    *input.ListBookByBookIDs
		Expected bool
	}{
		"ok": {
			Input: &input.ListBookByBookIDs{
				BookIDs: []int{1, 2},
			},
			Expected: true,
		},
		"ng_bookIds_unique": {
			Input: &input.ListBookByBookIDs{
				BookIDs: []int{1, 1},
			},
			Expected: false,
		},
		"ng_bookIds_required": {
			Input: &input.ListBookByBookIDs{
				BookIDs: []int{0, 1},
			},
			Expected: false,
		},
		"ng_bookIds_greater_than_equal": {
			Input: &input.ListBookByBookIDs{
				BookIDs: []int{0},
			},
			Expected: false,
		},
	}

	for result, tc := range testCases {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		t.Run(result, func(t *testing.T) {
			target := NewBookRequestValidation()

			got := target.ListBookByBookIDs(tc.Input)
			if tc.Expected {
				if got != nil {
					t.Fatalf("Incorrect result: %#v", got)
				}
			} else {
				if got == nil {
					t.Fatalf("Incorrect result: result is nil")
				}
			}
		})
	}
}

func TestBookRequestValidation_ListBookshelf(t *testing.T) {
	testCases := map[string]struct {
		Input    *input.ListBookshelf
		Expected bool
	}{
		"ok": {
			Input: &input.ListBookshelf{
				UserID: "00000000-0000-0000-0000-000000000000",
				Limit:  0,
				Offset: 0,
			},
			Expected: true,
		},
		"ng_userId_required": {
			Input: &input.ListBookshelf{
				UserID: "",
				Limit:  0,
				Offset: 0,
			},
			Expected: false,
		},
		"ng_limit_greater_than_equal": {
			Input: &input.ListBookshelf{
				UserID: "00000000-0000-0000-0000-000000000000",
				Limit:  -1,
				Offset: 0,
			},
			Expected: false,
		},
		"ng_limit_less_than_equal": {
			Input: &input.ListBookshelf{
				UserID: "00000000-0000-0000-0000-000000000000",
				Limit:  1001,
				Offset: 0,
			},
			Expected: false,
		},
		"ng_offset_greater_than_equal": {
			Input: &input.ListBookshelf{
				UserID: "00000000-0000-0000-0000-000000000000",
				Limit:  0,
				Offset: -1,
			},
			Expected: false,
		},
	}

	for result, tc := range testCases {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		t.Run(result, func(t *testing.T) {
			target := NewBookRequestValidation()

			got := target.ListBookshelf(tc.Input)
			if tc.Expected {
				if got != nil {
					t.Fatalf("Incorrect result: %#v", got)
				}
			} else {
				if got == nil {
					t.Fatalf("Incorrect result: result is nil")
				}
			}
		})
	}
}

func TestBookRequestValidation_ListBookReview(t *testing.T) {
	testCases := map[string]struct {
		Input    *input.ListBookReview
		Expected bool
	}{
		"ok": {
			Input: &input.ListBookReview{
				BookID:    1,
				Limit:     0,
				Offset:    0,
				By:        "",
				Direction: "",
			},
			Expected: true,
		},
		"ng_bookId_required": {
			Input: &input.ListBookReview{
				BookID:    0,
				Limit:     0,
				Offset:    0,
				By:        "",
				Direction: "",
			},
			Expected: false,
		},
		"ng_bookId_greater_than_equal": {
			Input: &input.ListBookReview{
				BookID:    0,
				Limit:     0,
				Offset:    0,
				By:        "",
				Direction: "",
			},
			Expected: false,
		},
		"ng_limit_greater_than_equal": {
			Input: &input.ListBookReview{
				BookID:    1,
				Limit:     -1,
				Offset:    0,
				By:        "",
				Direction: "",
			},
			Expected: false,
		},
		"ng_limit_less_than_equal": {
			Input: &input.ListBookReview{
				BookID:    1,
				Limit:     1001,
				Offset:    0,
				By:        "",
				Direction: "",
			},
			Expected: false,
		},
		"ng_offset_greater_than_equal": {
			Input: &input.ListBookReview{
				BookID:    1,
				Limit:     0,
				Offset:    -1,
				By:        "",
				Direction: "",
			},
			Expected: false,
		},
		"ng_by_oneof": {
			Input: &input.ListBookReview{
				BookID:    1,
				Limit:     0,
				Offset:    -1,
				By:        "userId",
				Direction: "asc",
			},
			Expected: false,
		},
		"ng_direction_oneof": {
			Input: &input.ListBookReview{
				BookID:    1,
				Limit:     0,
				Offset:    -1,
				By:        "id",
				Direction: "test",
			},
			Expected: false,
		},
	}

	for result, tc := range testCases {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		t.Run(result, func(t *testing.T) {
			target := NewBookRequestValidation()

			got := target.ListBookReview(tc.Input)
			if tc.Expected {
				if got != nil {
					t.Fatalf("Incorrect result: %#v", got)
				}
			} else {
				if got == nil {
					t.Fatalf("Incorrect result: result is nil")
				}
			}
		})
	}
}

func TestBookRequestValidation_ListUserReview(t *testing.T) {
	testCases := map[string]struct {
		Input    *input.ListUserReview
		Expected bool
	}{
		"ok": {
			Input: &input.ListUserReview{
				UserID:    "00000000-0000-0000-0000-000000000000",
				Limit:     0,
				Offset:    0,
				By:        "",
				Direction: "",
			},
			Expected: true,
		},
		"ng_userId_required": {
			Input: &input.ListUserReview{
				UserID:    "",
				Limit:     0,
				Offset:    0,
				By:        "",
				Direction: "",
			},
			Expected: false,
		},
		"ng_limit_greater_than_equal": {
			Input: &input.ListUserReview{
				UserID:    "00000000-0000-0000-0000-000000000000",
				Limit:     -1,
				Offset:    0,
				By:        "",
				Direction: "",
			},
			Expected: false,
		},
		"ng_limit_less_than_equal": {
			Input: &input.ListUserReview{
				UserID:    "00000000-0000-0000-0000-000000000000",
				Limit:     1001,
				Offset:    0,
				By:        "",
				Direction: "",
			},
			Expected: false,
		},
		"ng_offset_greater_than_equal": {
			Input: &input.ListUserReview{
				UserID:    "00000000-0000-0000-0000-000000000000",
				Limit:     0,
				Offset:    -1,
				By:        "",
				Direction: "",
			},
			Expected: false,
		},
		"ng_by_oneof": {
			Input: &input.ListUserReview{
				UserID:    "00000000-0000-0000-0000-000000000000",
				Limit:     0,
				Offset:    -1,
				By:        "userId",
				Direction: "asc",
			},
			Expected: false,
		},
		"ng_direction_oneof": {
			Input: &input.ListUserReview{
				UserID:    "00000000-0000-0000-0000-000000000000",
				Limit:     0,
				Offset:    -1,
				By:        "id",
				Direction: "test",
			},
			Expected: false,
		},
	}

	for result, tc := range testCases {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		t.Run(result, func(t *testing.T) {
			target := NewBookRequestValidation()

			got := target.ListUserReview(tc.Input)
			if tc.Expected {
				if got != nil {
					t.Fatalf("Incorrect result: %#v", got)
				}
			} else {
				if got == nil {
					t.Fatalf("Incorrect result: result is nil")
				}
			}
		})
	}
}
