package book

import (
	"testing"
	"time"

	"github.com/calmato/gran-book/api/pkg/datetime"
	"github.com/calmato/gran-book/api/proto/book"
	"github.com/stretchr/testify/assert"
)

func TestBook(t *testing.T) {
	t.Parallel()
	now := time.Now().Local()
	tests := []struct {
		name        string
		book        *Book
		expectProto *book.Book
	}{
		{
			name: "success",
			book: &Book{
				ID:             1,
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
				CreatedAt:      now,
				UpdatedAt:      now,
				Authors: Authors{
					{
						ID:        1,
						Name:      "有沢 ゆう希",
						NameKana:  "アリサワ ユウキ",
						CreatedAt: now,
						UpdatedAt: now,
					},
					{
						ID:        2,
						Name:      "末次 由紀",
						NameKana:  "スエツグ ユキ",
						CreatedAt: now,
						UpdatedAt: now,
					},
				},
			},
			expectProto: &book.Book{
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
				CreatedAt:      datetime.FormatTime(now),
				UpdatedAt:      datetime.FormatTime(now),
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
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.expectProto, tt.book.Proto())
		})
	}
}

func TestBooks(t *testing.T) {
	t.Parallel()
	now := time.Now().Local()
	tests := []struct {
		name        string
		books       Books
		expectProto []*book.Book
	}{
		{
			name: "success",
			books: Books{
				{
					ID:             1,
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
					CreatedAt:      now,
					UpdatedAt:      now,
					Authors: Authors{
						{
							ID:        1,
							Name:      "有沢 ゆう希",
							NameKana:  "アリサワ ユウキ",
							CreatedAt: now,
							UpdatedAt: now,
						},
						{
							ID:        2,
							Name:      "末次 由紀",
							NameKana:  "スエツグ ユキ",
							CreatedAt: now,
							UpdatedAt: now,
						},
					},
				},
				{
					ID:             2,
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
					CreatedAt:      now,
					UpdatedAt:      now,
					Authors: Authors{
						{
							ID:        1,
							Name:      "有沢 ゆう希",
							NameKana:  "アリサワ ユウキ",
							CreatedAt: now,
							UpdatedAt: now,
						},
						{
							ID:        2,
							Name:      "末次 由紀",
							NameKana:  "スエツグ ユキ",
							CreatedAt: now,
							UpdatedAt: now,
						},
					},
				},
			},
			expectProto: []*book.Book{
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
					CreatedAt:      datetime.FormatTime(now),
					UpdatedAt:      datetime.FormatTime(now),
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
				{
					Id:             2,
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
					CreatedAt:      datetime.FormatTime(now),
					UpdatedAt:      datetime.FormatTime(now),
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
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.expectProto, tt.books.Proto())
		})
	}
}

func TestMonthlyResult(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name        string
		result      *MonthlyResult
		expectProto *book.MonthlyResult
	}{
		{
			name: "success",
			result: &MonthlyResult{
				Year:      2021,
				Month:     8,
				ReadTotal: 10,
			},
			expectProto: &book.MonthlyResult{
				Year:      2021,
				Month:     8,
				ReadTotal: 10,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.expectProto, tt.result.Proto())
		})
	}
}

func TestMonthlyResults(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name        string
		results     MonthlyResults
		expectProto []*book.MonthlyResult
	}{
		{
			name: "success",
			results: MonthlyResults{
				{
					Year:      2021,
					Month:     8,
					ReadTotal: 10,
				},
				{
					Year:      2021,
					Month:     9,
					ReadTotal: 5,
				},
			},
			expectProto: []*book.MonthlyResult{
				{
					Year:      2021,
					Month:     8,
					ReadTotal: 10,
				},
				{
					Year:      2021,
					Month:     9,
					ReadTotal: 5,
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.expectProto, tt.results.Proto())
		})
	}
}

func TestAuthor(t *testing.T) {
	t.Parallel()
	now := time.Now().Local()
	tests := []struct {
		name        string
		author      *Author
		expectProto *book.Author
	}{
		{
			name: "success",
			author: &Author{
				ID:        1,
				Name:      "有沢 ゆう希",
				NameKana:  "アリサワ ユウキ",
				CreatedAt: now,
				UpdatedAt: now,
			},
			expectProto: &book.Author{
				Name:     "有沢 ゆう希",
				NameKana: "アリサワ ユウキ",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.expectProto, tt.author.Proto())
		})
	}
}

func TestAuthors(t *testing.T) {
	t.Parallel()
	now := time.Now().Local()
	tests := []struct {
		name        string
		authors     Authors
		expectProto []*book.Author
	}{
		{
			name: "success",
			authors: Authors{
				{
					ID:        1,
					Name:      "有沢 ゆう希",
					NameKana:  "アリサワ ユウキ",
					CreatedAt: now,
					UpdatedAt: now,
				},
				{
					ID:        2,
					Name:      "末次 由紀",
					NameKana:  "スエツグ ユキ",
					CreatedAt: now,
					UpdatedAt: now,
				},
			},
			expectProto: []*book.Author{
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
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.expectProto, tt.authors.Proto())
		})
	}
}

func TestBookshelf(t *testing.T) {
	t.Parallel()
	now := time.Now().Local()
	tests := []struct {
		name        string
		bookshelf   *Bookshelf
		expectProto *book.Bookshelf
	}{
		{
			name: "success",
			bookshelf: &Bookshelf{
				ID:        1,
				BookID:    1,
				UserID:    "00000000-0000-0000-0000-000000000000",
				ReviewID:  0,
				Status:    ReadingStatus,
				ReadOn:    now,
				CreatedAt: now,
				UpdatedAt: now,
			},
			expectProto: &book.Bookshelf{
				Id:        1,
				BookId:    1,
				UserId:    "00000000-0000-0000-0000-000000000000",
				ReviewId:  0,
				Status:    book.BookshelfStatus_BOOKSHELF_STATUS_READING,
				ReadOn:    datetime.FormatDate(now),
				CreatedAt: datetime.FormatTime(now),
				UpdatedAt: datetime.FormatTime(now),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.expectProto, tt.bookshelf.Proto())
		})
	}
}

func TestBookshelves(t *testing.T) {
	t.Parallel()
	now := time.Now().Local()
	tests := []struct {
		name        string
		bookshelves Bookshelves
		expectProto []*book.Bookshelf
	}{
		{
			name: "success",
			bookshelves: Bookshelves{
				{
					ID:        1,
					BookID:    1,
					UserID:    "00000000-0000-0000-0000-000000000000",
					ReviewID:  0,
					Status:    ReadingStatus,
					ReadOn:    now,
					CreatedAt: now,
					UpdatedAt: now,
				},
				{
					ID:        2,
					BookID:    2,
					UserID:    "00000000-0000-0000-0000-000000000000",
					ReviewID:  0,
					Status:    StackedStatus,
					ReadOn:    now,
					CreatedAt: now,
					UpdatedAt: now,
				},
			},
			expectProto: []*book.Bookshelf{
				{
					Id:        1,
					BookId:    1,
					UserId:    "00000000-0000-0000-0000-000000000000",
					ReviewId:  0,
					Status:    book.BookshelfStatus_BOOKSHELF_STATUS_READING,
					ReadOn:    datetime.FormatDate(now),
					CreatedAt: datetime.FormatTime(now),
					UpdatedAt: datetime.FormatTime(now),
				},
				{
					Id:        2,
					BookId:    2,
					UserId:    "00000000-0000-0000-0000-000000000000",
					ReviewId:  0,
					Status:    book.BookshelfStatus_BOOKSHELF_STATUS_STACKED,
					ReadOn:    datetime.FormatDate(now),
					CreatedAt: datetime.FormatTime(now),
					UpdatedAt: datetime.FormatTime(now),
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.expectProto, tt.bookshelves.Proto())
		})
	}
}

func TestReview(t *testing.T) {
	t.Parallel()
	now := time.Now().Local()
	tests := []struct {
		name        string
		review      *Review
		expectProto *book.Review
	}{
		{
			name: "success",
			review: &Review{
				ID:         1,
				BookID:     1,
				UserID:     "00000000-0000-0000-0000-000000000000",
				Score:      3,
				Impression: "テストレビューです",
				CreatedAt:  now,
				UpdatedAt:  now,
			},
			expectProto: &book.Review{
				Id:         1,
				BookId:     1,
				UserId:     "00000000-0000-0000-0000-000000000000",
				Score:      3,
				Impression: "テストレビューです",
				CreatedAt:  datetime.FormatTime(now),
				UpdatedAt:  datetime.FormatTime(now),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.expectProto, tt.review.Proto())
		})
	}
}

func TestReviews(t *testing.T) {
	t.Parallel()
	now := time.Now().Local()
	tests := []struct {
		name        string
		reviews     Reviews
		expectProto []*book.Review
	}{
		{
			name: "success",
			reviews: Reviews{
				{
					ID:         1,
					BookID:     1,
					UserID:     "00000000-0000-0000-0000-000000000000",
					Score:      3,
					Impression: "テストレビューです",
					CreatedAt:  now,
					UpdatedAt:  now,
				},
				{
					ID:         2,
					BookID:     2,
					UserID:     "00000000-0000-0000-0000-000000000000",
					Score:      3,
					Impression: "テストレビューです",
					CreatedAt:  now,
					UpdatedAt:  now,
				},
			},
			expectProto: []*book.Review{
				{
					Id:         1,
					BookId:     1,
					UserId:     "00000000-0000-0000-0000-000000000000",
					Score:      3,
					Impression: "テストレビューです",
					CreatedAt:  datetime.FormatTime(now),
					UpdatedAt:  datetime.FormatTime(now),
				},
				{
					Id:         2,
					BookId:     2,
					UserId:     "00000000-0000-0000-0000-000000000000",
					Score:      3,
					Impression: "テストレビューです",
					CreatedAt:  datetime.FormatTime(now),
					UpdatedAt:  datetime.FormatTime(now),
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.expectProto, tt.reviews.Proto())
		})
	}
}
