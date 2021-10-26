package entity

import (
	"testing"

	"github.com/calmato/gran-book/api/service/internal/gateway/entity"
	"github.com/calmato/gran-book/api/service/pkg/datetime"
	"github.com/calmato/gran-book/api/service/pkg/test"
	"github.com/calmato/gran-book/api/service/proto/book"
	"github.com/stretchr/testify/assert"
)

func TestBook(t *testing.T) {
	t.Parallel()
	now := datetime.FormatTime(test.TimeMock)
	tests := []struct {
		name    string
		book    *entity.Book
		reviews BookReviews
		users   map[string]*entity.User
		expect  *Book
	}{
		{
			name: "success",
			book: &entity.Book{
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
				{
					ID:         2,
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
			expect: &Book{
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
				CreatedAt:    now,
				UpdatedAt:    now,
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
					{
						ID:         2,
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
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			actual := NewBook(tt.book, tt.reviews)
			assert.Equal(t, tt.expect, actual)
		})
	}
}

func TestBook_Fill(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name   string
		book   *Book
		limit  int64
		offset int64
		total  int64
		expect *Book
	}{
		{
			name:   "success",
			book:   &Book{},
			limit:  100,
			offset: 20,
			total:  50,
			expect: &Book{
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
			tt.book.Fill(tt.limit, tt.offset, tt.total)
			assert.Equal(t, tt.expect, tt.book)
		})
	}
}
