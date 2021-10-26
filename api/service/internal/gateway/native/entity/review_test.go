package entity

import (
	"testing"

	"github.com/calmato/gran-book/api/service/internal/gateway/entity"
	"github.com/calmato/gran-book/api/service/pkg/datetime"
	"github.com/calmato/gran-book/api/service/pkg/test"
	"github.com/calmato/gran-book/api/service/proto/book"
	"github.com/calmato/gran-book/api/service/proto/user"
	"github.com/stretchr/testify/assert"
)

func TestBookReviews(t *testing.T) {
	t.Parallel()
	now := datetime.FormatTime(test.TimeMock)
	tests := []struct {
		name    string
		reviews entity.Reviews
		users   map[string]*entity.User
		expect  BookReviews
	}{
		{
			name: "success",
			reviews: entity.Reviews{
				{
					Review: &book.Review{
						Id:         1,
						BookId:     1,
						UserId:     "00000000-0000-0000-0000-000000000000",
						Score:      3,
						Impression: "テストレビューです",
						CreatedAt:  now,
						UpdatedAt:  now,
					},
				},
				{
					Review: &book.Review{
						Id:         2,
						BookId:     2,
						UserId:     "00000000-0000-0000-0000-000000000000",
						Score:      3,
						Impression: "テストレビューです",
						CreatedAt:  now,
						UpdatedAt:  now,
					},
				},
			},
			users: map[string]*entity.User{
				"00000000-0000-0000-0000-000000000000": {
					User: &user.User{
						Id:               "00000000-0000-0000-0000-000000000000",
						Username:         "テストユーザー",
						Gender:           user.Gender_GENDER_MAN,
						Email:            "test-user01@calmato.jp",
						PhoneNumber:      "000-0000-0000",
						ThumbnailUrl:     "https://go.dev/images/gophers/ladder.svg",
						SelfIntroduction: "テストコードです。",
						LastName:         "テスト",
						FirstName:        "ユーザー",
						LastNameKana:     "てすと",
						FirstNameKana:    "ゆーざー",
						CreatedAt:        now,
						UpdatedAt:        now,
					},
				},
			},
			expect: BookReviews{
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
		{
			name: "success users is length 0",
			reviews: entity.Reviews{
				{
					Review: &book.Review{
						Id:         1,
						BookId:     1,
						UserId:     "00000000-0000-0000-0000-000000000000",
						Score:      3,
						Impression: "テストレビューです",
						CreatedAt:  now,
						UpdatedAt:  now,
					},
				},
				{
					Review: &book.Review{
						Id:         2,
						BookId:     2,
						UserId:     "00000000-0000-0000-0000-000000000000",
						Score:      3,
						Impression: "テストレビューです",
						CreatedAt:  now,
						UpdatedAt:  now,
					},
				},
			},
			users: map[string]*entity.User{},
			expect: BookReviews{
				{
					ID:         1,
					Impression: "テストレビューです",
					User: &BookReviewUser{
						Username: "unknown",
					},
					CreatedAt: now,
					UpdatedAt: now,
				},
				{
					ID:         2,
					Impression: "テストレビューです",
					User: &BookReviewUser{
						Username: "unknown",
					},
					CreatedAt: now,
					UpdatedAt: now,
				},
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			actual := NewBookReviews(tt.reviews, tt.users)
			assert.Equal(t, tt.expect, actual)
		})
	}
}

func TestUserReviews(t *testing.T) {
	t.Parallel()
	now := datetime.FormatTime(test.TimeMock)
	tests := []struct {
		name    string
		reviews entity.Reviews
		books   map[int64]*entity.Book
		expect  UserReviews
	}{
		{
			name: "success",
			reviews: entity.Reviews{
				{
					Review: &book.Review{
						Id:         1,
						BookId:     1,
						UserId:     "00000000-0000-0000-0000-000000000000",
						Score:      3,
						Impression: "テストレビューです",
						CreatedAt:  now,
						UpdatedAt:  now,
					},
				},
				{
					Review: &book.Review{
						Id:         2,
						BookId:     2,
						UserId:     "00000000-0000-0000-0000-000000000000",
						Score:      3,
						Impression: "テストレビューです",
						CreatedAt:  now,
						UpdatedAt:  now,
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
			expect: UserReviews{
				{
					ID:         1,
					Impression: "テストレビューです",
					Book: &UserReviewBook{
						ID:           1,
						Title:        "小説　ちはやふる　上の句",
						ThumbnailURL: "https://thumbnail.image.rakuten.co.jp/@0_mall/book/cabinet/8426/9784062938426.jpg?_ex=120x120",
					},
					CreatedAt: now,
					UpdatedAt: now,
				},
				{
					ID:         2,
					Impression: "テストレビューです",
					Book: &UserReviewBook{
						ID:           2,
						Title:        "小説　ちはやふる　上の句",
						ThumbnailURL: "https://thumbnail.image.rakuten.co.jp/@0_mall/book/cabinet/8426/9784062938426.jpg?_ex=120x120",
					},
					CreatedAt: now,
					UpdatedAt: now,
				},
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			actual := NewUserReviews(tt.reviews, tt.books)
			assert.Equal(t, tt.expect, actual)
		})
	}
}
