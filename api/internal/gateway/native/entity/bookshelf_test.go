package entity

import (
	"testing"

	"github.com/calmato/gran-book/api/internal/gateway/entity"
	"github.com/calmato/gran-book/api/pkg/datetime"
	"github.com/calmato/gran-book/api/pkg/test"
	"github.com/calmato/gran-book/api/proto/book"
	"github.com/stretchr/testify/assert"
)

func TestBookOnBookshelf_Fill(t *testing.T) {
	t.Parallel()
	now := datetime.FormatTime(test.TimeMock)
	tests := []struct {
		name    string
		book    *BookOnBookshelf
		reviews BookReviews
		limit   int64
		offset  int64
		total   int64
		expect  *BookOnBookshelf
	}{
		{
			name: "success",
			book: &BookOnBookshelf{},
			reviews: BookReviews{
				{
					ID:         1,
					Impression: "テストレビューです",
					User: &BookReviewUser{
						ID:           "00000000-0000-0000-0000-000000000000",
						Username:     "テストユーザー",
						ThumbnailURL: "https://go.dev/images/gophers/ladder.svg",
					},
					CreatedAt: now,
					UpdatedAt: now,
				},
			},
			limit:  100,
			offset: 20,
			total:  50,
			expect: &BookOnBookshelf{
				Reviews: BookReviews{
					{
						ID:         1,
						Impression: "テストレビューです",
						User: &BookReviewUser{
							ID:           "00000000-0000-0000-0000-000000000000",
							Username:     "テストユーザー",
							ThumbnailURL: "https://go.dev/images/gophers/ladder.svg",
						},
						CreatedAt: now,
						UpdatedAt: now,
					},
				},
				ReviewLimit:  100,
				ReviewOffset: 20,
				ReviewTotal:  50,
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			tt.book.Fill(tt.reviews, tt.limit, tt.offset, tt.total)
			assert.Equal(t, tt.expect, tt.book)
		})
	}
}

func TestBooksOnBookshelf(t *testing.T) {
	t.Parallel()
	now := datetime.FormatTime(test.TimeMock)
	tests := []struct {
		name        string
		bookshelves entity.Bookshelves
		books       map[int64]*entity.Book
		expect      BooksOnBookshelf
	}{
		{
			name: "success",
			bookshelves: entity.Bookshelves{
				{
					Bookshelf: &book.Bookshelf{
						Id:        1,
						BookId:    1,
						UserId:    "00000000-0000-0000-0000-000000000000",
						ReviewId:  0,
						Status:    book.BookshelfStatus_BOOKSHELF_STATUS_READING,
						ReadOn:    datetime.FormatDate(test.DateMock),
						CreatedAt: now,
						UpdatedAt: now,
					},
				},
				{
					Bookshelf: &book.Bookshelf{
						Id:        2,
						BookId:    2,
						UserId:    "00000000-0000-0000-0000-000000000000",
						ReviewId:  0,
						Status:    book.BookshelfStatus_BOOKSHELF_STATUS_STACKED,
						ReadOn:    datetime.FormatDate(test.DateMock),
						CreatedAt: now,
						UpdatedAt: now,
					},
				},
			},
			books: map[int64]*entity.Book{
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
				2: {
					Book: &book.Book{
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
			expect: BooksOnBookshelf{
				{
					ID:           1,
					Title:        "小説　ちはやふる　上の句",
					TitleKana:    "ショウセツ チハヤフルカミノク",
					Description:  "綾瀬千早は高校入学と同時に、競技かるた部を作ろうと奔走する。幼馴染の太一と仲間を集め、夏の全国大会に出場するためだ。強くなって、新と再会したい。幼い頃かるたを取り合った、新に寄せる千早の秘めた想いに気づきながらも、太一は千早を守り立てる。それぞれの青春を懸けた、一途な情熱の物語が幕開ける。",
					Isbn:         "9784062938426",
					Publisher:    "講談社",
					PublishedOn:  "2018-01-16",
					ThumbnailURL: "https://thumbnail.image.rakuten.co.jp/@0_mall/book/cabinet/8426/9784062938426.jpg?_ex=120x120",
					RakutenURL:   "https://books.rakuten.co.jp/rb/15271426/",
					Size:         "コミック",
					Author:       "有沢 ゆう希/末次 由紀",
					AuthorKana:   "アリサワ ユウキ/スエツグ ユキ",
					Bookshelf: &Bookshelf{
						ReviewID:  0,
						Status:    entity.BookshelfStatusReading.Name(),
						ReadOn:    datetime.FormatDate(test.DateMock),
						CreatedAt: now,
						UpdatedAt: now,
					},
					CreatedAt: now,
					UpdatedAt: now,
				},
				{
					ID:           2,
					Title:        "小説　ちはやふる　上の句",
					TitleKana:    "ショウセツ チハヤフルカミノク",
					Description:  "綾瀬千早は高校入学と同時に、競技かるた部を作ろうと奔走する。幼馴染の太一と仲間を集め、夏の全国大会に出場するためだ。強くなって、新と再会したい。幼い頃かるたを取り合った、新に寄せる千早の秘めた想いに気づきながらも、太一は千早を守り立てる。それぞれの青春を懸けた、一途な情熱の物語が幕開ける。",
					Isbn:         "9784062938426",
					Publisher:    "講談社",
					PublishedOn:  "2018-01-16",
					ThumbnailURL: "https://thumbnail.image.rakuten.co.jp/@0_mall/book/cabinet/8426/9784062938426.jpg?_ex=120x120",
					RakutenURL:   "https://books.rakuten.co.jp/rb/15271426/",
					Size:         "コミック",
					Author:       "有沢 ゆう希/末次 由紀",
					AuthorKana:   "アリサワ ユウキ/スエツグ ユキ",
					Bookshelf: &Bookshelf{
						ReviewID:  0,
						Status:    entity.BookshelfStatusStacked.Name(),
						ReadOn:    datetime.FormatDate(test.DateMock),
						CreatedAt: now,
						UpdatedAt: now,
					},
					CreatedAt: now,
					UpdatedAt: now,
				},
			},
		},
		{
			name: "success books length is 0",
			bookshelves: entity.Bookshelves{
				{
					Bookshelf: &book.Bookshelf{
						Id:        1,
						BookId:    1,
						UserId:    "00000000-0000-0000-0000-000000000000",
						ReviewId:  0,
						Status:    book.BookshelfStatus_BOOKSHELF_STATUS_READING,
						ReadOn:    datetime.FormatDate(test.DateMock),
						CreatedAt: now,
						UpdatedAt: now,
					},
				},
				{
					Bookshelf: &book.Bookshelf{
						Id:        2,
						BookId:    2,
						UserId:    "00000000-0000-0000-0000-000000000000",
						ReviewId:  0,
						Status:    book.BookshelfStatus_BOOKSHELF_STATUS_STACKED,
						ReadOn:    datetime.FormatDate(test.DateMock),
						CreatedAt: now,
						UpdatedAt: now,
					},
				},
			},
			books:  map[int64]*entity.Book{},
			expect: BooksOnBookshelf{},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			actual := NewBooksOnBookshelf(tt.books, tt.bookshelves)
			assert.Equal(t, tt.expect, actual)
		})
	}
}
