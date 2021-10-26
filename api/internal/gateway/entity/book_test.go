package entity

import (
	"testing"

	"github.com/calmato/gran-book/api/pkg/datetime"
	"github.com/calmato/gran-book/api/pkg/test"
	"github.com/calmato/gran-book/api/proto/book"
	"github.com/stretchr/testify/assert"
)

func TestBook(t *testing.T) {
	t.Parallel()
	now := datetime.FormatTime(test.Now())
	tests := []struct {
		name   string
		book   *book.Book
		expect *Book
	}{
		{
			name: "success",
			book: &book.Book{
				Id:             1,
				Title:          "小説　ちはやふる　上の句",
				TitleKana:      "ショウセツ チハヤフルカミノク",
				Description:    "綾瀬千早は高校入学と同時に、競技かるた部を作ろうと奔走する。幼馴染の太一と仲間を集め、夏の全国大会に出場するためだ。強くなって、新と再会したい。幼い頃かるたを取り合った、新に寄せる千早の秘めた想いに気づきながらも、太一は千早を守り立てる。それぞれの青春を懸けた、一途な情熱の物語が幕開ける。",
				Isbn:           "9784062938426",
				Publisher:      "講談社",
				PublishedOn:    "2018-01-16",
				ThumbnailUrl:   "https://thumbnail.image.rakuten.co.jp/@0_mall/book/cabinet/8426/9784062938426.jpg?_ex=120x120",
				RakutenUrl:     "https://books.rakuten.co.jp/rb/15271426/",
				RakutenSize:    "コミック",
				RakutenGenreId: "001004008001/001004008003/001019001",
				CreatedAt:      now,
				UpdatedAt:      now,
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
			expect: &Book{
				Book: &book.Book{
					Id:             1,
					Title:          "小説　ちはやふる　上の句",
					TitleKana:      "ショウセツ チハヤフルカミノク",
					Description:    "綾瀬千早は高校入学と同時に、競技かるた部を作ろうと奔走する。幼馴染の太一と仲間を集め、夏の全国大会に出場するためだ。強くなって、新と再会したい。幼い頃かるたを取り合った、新に寄せる千早の秘めた想いに気づきながらも、太一は千早を守り立てる。それぞれの青春を懸けた、一途な情熱の物語が幕開ける。",
					Isbn:           "9784062938426",
					Publisher:      "講談社",
					PublishedOn:    "2018-01-16",
					ThumbnailUrl:   "https://thumbnail.image.rakuten.co.jp/@0_mall/book/cabinet/8426/9784062938426.jpg?_ex=120x120",
					RakutenUrl:     "https://books.rakuten.co.jp/rb/15271426/",
					RakutenSize:    "コミック",
					RakutenGenreId: "001004008001/001004008003/001019001",
					CreatedAt:      now,
					UpdatedAt:      now,
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
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			actual := NewBook(tt.book)
			assert.Equal(t, tt.expect, actual)
		})
	}
}

func TestBook_AuthorNames(t *testing.T) {
	t.Parallel()
	now := datetime.FormatTime(test.Now())
	tests := []struct {
		name   string
		book   *book.Book
		expect []string
	}{
		{
			name: "success",
			book: &book.Book{
				Id:             1,
				Title:          "小説　ちはやふる　上の句",
				TitleKana:      "ショウセツ チハヤフルカミノク",
				Description:    "綾瀬千早は高校入学と同時に、競技かるた部を作ろうと奔走する。幼馴染の太一と仲間を集め、夏の全国大会に出場するためだ。強くなって、新と再会したい。幼い頃かるたを取り合った、新に寄せる千早の秘めた想いに気づきながらも、太一は千早を守り立てる。それぞれの青春を懸けた、一途な情熱の物語が幕開ける。",
				Isbn:           "9784062938426",
				Publisher:      "講談社",
				PublishedOn:    "2018-01-16",
				ThumbnailUrl:   "https://thumbnail.image.rakuten.co.jp/@0_mall/book/cabinet/8426/9784062938426.jpg?_ex=120x120",
				RakutenUrl:     "https://books.rakuten.co.jp/rb/15271426/",
				RakutenSize:    "コミック",
				RakutenGenreId: "001004008001/001004008003/001019001",
				CreatedAt:      now,
				UpdatedAt:      now,
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
			expect: []string{"有沢 ゆう希", "末次 由紀"},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			actual := NewBook(tt.book)
			assert.Equal(t, tt.expect, actual.AuthorNames())
		})
	}
}

func TestBook_AuthorNameKanas(t *testing.T) {
	t.Parallel()
	now := datetime.FormatTime(test.Now())
	tests := []struct {
		name   string
		book   *book.Book
		expect []string
	}{
		{
			name: "success",
			book: &book.Book{
				Id:             1,
				Title:          "小説　ちはやふる　上の句",
				TitleKana:      "ショウセツ チハヤフルカミノク",
				Description:    "綾瀬千早は高校入学と同時に、競技かるた部を作ろうと奔走する。幼馴染の太一と仲間を集め、夏の全国大会に出場するためだ。強くなって、新と再会したい。幼い頃かるたを取り合った、新に寄せる千早の秘めた想いに気づきながらも、太一は千早を守り立てる。それぞれの青春を懸けた、一途な情熱の物語が幕開ける。",
				Isbn:           "9784062938426",
				Publisher:      "講談社",
				PublishedOn:    "2018-01-16",
				ThumbnailUrl:   "https://thumbnail.image.rakuten.co.jp/@0_mall/book/cabinet/8426/9784062938426.jpg?_ex=120x120",
				RakutenUrl:     "https://books.rakuten.co.jp/rb/15271426/",
				RakutenSize:    "コミック",
				RakutenGenreId: "001004008001/001004008003/001019001",
				CreatedAt:      now,
				UpdatedAt:      now,
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
			expect: []string{"アリサワ ユウキ", "スエツグ ユキ"},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			actual := NewBook(tt.book)
			assert.Equal(t, tt.expect, actual.AuthorNameKanas())
		})
	}
}

func TestBooks(t *testing.T) {
	t.Parallel()
	now := datetime.FormatTime(test.Now())
	tests := []struct {
		name   string
		books  []*book.Book
		expect Books
	}{
		{
			name: "success",
			books: []*book.Book{
				{
					Id:             1,
					Title:          "小説　ちはやふる　上の句",
					TitleKana:      "ショウセツ チハヤフルカミノク",
					Description:    "綾瀬千早は高校入学と同時に、競技かるた部を作ろうと奔走する。幼馴染の太一と仲間を集め、夏の全国大会に出場するためだ。強くなって、新と再会したい。幼い頃かるたを取り合った、新に寄せる千早の秘めた想いに気づきながらも、太一は千早を守り立てる。それぞれの青春を懸けた、一途な情熱の物語が幕開ける。",
					Isbn:           "9784062938426",
					Publisher:      "講談社",
					PublishedOn:    "2018-01-16",
					ThumbnailUrl:   "https://thumbnail.image.rakuten.co.jp/@0_mall/book/cabinet/8426/9784062938426.jpg?_ex=120x120",
					RakutenUrl:     "https://books.rakuten.co.jp/rb/15271426/",
					RakutenSize:    "コミック",
					RakutenGenreId: "001004008001/001004008003/001019001",
					CreatedAt:      now,
					UpdatedAt:      now,
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
			},
			expect: Books{
				{
					Book: &book.Book{
						Id:             1,
						Title:          "小説　ちはやふる　上の句",
						TitleKana:      "ショウセツ チハヤフルカミノク",
						Description:    "綾瀬千早は高校入学と同時に、競技かるた部を作ろうと奔走する。幼馴染の太一と仲間を集め、夏の全国大会に出場するためだ。強くなって、新と再会したい。幼い頃かるたを取り合った、新に寄せる千早の秘めた想いに気づきながらも、太一は千早を守り立てる。それぞれの青春を懸けた、一途な情熱の物語が幕開ける。",
						Isbn:           "9784062938426",
						Publisher:      "講談社",
						PublishedOn:    "2018-01-16",
						ThumbnailUrl:   "https://thumbnail.image.rakuten.co.jp/@0_mall/book/cabinet/8426/9784062938426.jpg?_ex=120x120",
						RakutenUrl:     "https://books.rakuten.co.jp/rb/15271426/",
						RakutenSize:    "コミック",
						RakutenGenreId: "001004008001/001004008003/001019001",
						CreatedAt:      now,
						UpdatedAt:      now,
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
				},
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			actual := NewBooks(tt.books)
			assert.Equal(t, tt.expect, actual)
		})
	}
}

func TestBooks_Map(t *testing.T) {
	t.Parallel()
	now := datetime.FormatTime(test.Now())
	tests := []struct {
		name   string
		books  []*book.Book
		expect map[int64]*Book
	}{
		{
			name: "success",
			books: []*book.Book{
				{
					Id:             1,
					Title:          "小説　ちはやふる　上の句",
					TitleKana:      "ショウセツ チハヤフルカミノク",
					Description:    "綾瀬千早は高校入学と同時に、競技かるた部を作ろうと奔走する。幼馴染の太一と仲間を集め、夏の全国大会に出場するためだ。強くなって、新と再会したい。幼い頃かるたを取り合った、新に寄せる千早の秘めた想いに気づきながらも、太一は千早を守り立てる。それぞれの青春を懸けた、一途な情熱の物語が幕開ける。",
					Isbn:           "9784062938426",
					Publisher:      "講談社",
					PublishedOn:    "2018-01-16",
					ThumbnailUrl:   "https://thumbnail.image.rakuten.co.jp/@0_mall/book/cabinet/8426/9784062938426.jpg?_ex=120x120",
					RakutenUrl:     "https://books.rakuten.co.jp/rb/15271426/",
					RakutenSize:    "コミック",
					RakutenGenreId: "001004008001/001004008003/001019001",
					CreatedAt:      now,
					UpdatedAt:      now,
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
			},
			expect: map[int64]*Book{
				1: {
					Book: &book.Book{
						Id:             1,
						Title:          "小説　ちはやふる　上の句",
						TitleKana:      "ショウセツ チハヤフルカミノク",
						Description:    "綾瀬千早は高校入学と同時に、競技かるた部を作ろうと奔走する。幼馴染の太一と仲間を集め、夏の全国大会に出場するためだ。強くなって、新と再会したい。幼い頃かるたを取り合った、新に寄せる千早の秘めた想いに気づきながらも、太一は千早を守り立てる。それぞれの青春を懸けた、一途な情熱の物語が幕開ける。",
						Isbn:           "9784062938426",
						Publisher:      "講談社",
						PublishedOn:    "2018-01-16",
						ThumbnailUrl:   "https://thumbnail.image.rakuten.co.jp/@0_mall/book/cabinet/8426/9784062938426.jpg?_ex=120x120",
						RakutenUrl:     "https://books.rakuten.co.jp/rb/15271426/",
						RakutenSize:    "コミック",
						RakutenGenreId: "001004008001/001004008003/001019001",
						CreatedAt:      now,
						UpdatedAt:      now,
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
				},
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			actual := NewBooks(tt.books)
			assert.Equal(t, tt.expect, actual.Map())
		})
	}
}
