package v2

import (
	"testing"

	"github.com/calmato/gran-book/api/service/internal/gateway/entity"
	"github.com/calmato/gran-book/api/service/pkg/datetime"
	"github.com/calmato/gran-book/api/service/pkg/test"
	"github.com/calmato/gran-book/api/service/proto/book"
	"github.com/calmato/gran-book/api/service/proto/user"
	"github.com/stretchr/testify/assert"
)

func TestBookResponse(t *testing.T) {
	t.Parallel()
	now := datetime.FormatTime(test.TimeMock)
	type args struct {
		book    *entity.Book
		reviews entity.Reviews
		users   map[string]*entity.User
		limit   int64
		offset  int64
		total   int64
	}
	tests := []struct {
		name   string
		args   args
		expect *BookResponse
	}{
		{
			name: "success",
			args: args{
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
							UserId:     "11111111-1111-1111-1111-111111111111",
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
					"11111111-1111-1111-1111-111111111111": {
						User: &user.User{
							Id:               "11111111-1111-1111-1111-111111111111",
							Username:         "テストユーザー",
							Gender:           user.Gender_GENDER_MAN,
							Email:            "test-user02@calmato.jp",
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
				limit:  100,
				offset: 0,
				total:  2,
			},
			expect: &BookResponse{
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
				Reviews: []*bookReview{
					{
						ID:         1,
						Impression: "テストレビューです",
						User: &bookUser{
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
						User: &bookUser{
							ID:           "11111111-1111-1111-1111-111111111111",
							Username:     "テストユーザー",
							ThumbnailURL: "https://go.dev/images/gophers/ladder.svg",
						},
						CreatedAt: now,
						UpdatedAt: now,
					},
				},
				ReviewLimit:  100,
				ReviewOffset: 0,
				ReviewTotal:  2,
			},
		},
		{
			name: "success users is length 0",
			args: args{
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
							UserId:     "11111111-1111-1111-1111-111111111111",
							Score:      3,
							Impression: "テストレビューです",
							CreatedAt:  now,
							UpdatedAt:  now,
						},
					},
				},
				users:  map[string]*entity.User{},
				limit:  100,
				offset: 0,
				total:  2,
			},
			expect: &BookResponse{
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
				Reviews: []*bookReview{
					{
						ID:         1,
						Impression: "テストレビューです",
						User: &bookUser{
							Username: "unknown",
						},
						CreatedAt: now,
						UpdatedAt: now,
					},
					{
						ID:         2,
						Impression: "テストレビューです",
						User: &bookUser{
							Username: "unknown",
						},
						CreatedAt: now,
						UpdatedAt: now,
					},
				},
				ReviewLimit:  100,
				ReviewOffset: 0,
				ReviewTotal:  2,
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			actual := NewBookResponse(
				tt.args.book,
				tt.args.reviews,
				tt.args.users,
				tt.args.limit,
				tt.args.offset,
				tt.args.total,
			)
			assert.Equal(t, tt.expect, actual)
		})
	}
}
