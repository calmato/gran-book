package v1

import (
	"testing"

	"github.com/calmato/gran-book/api/gateway/native/internal/entity"
	"github.com/calmato/gran-book/api/gateway/native/pkg/test"
	"github.com/calmato/gran-book/api/gateway/native/proto/service/book"
	"github.com/calmato/gran-book/api/gateway/native/proto/service/user"
	"github.com/stretchr/testify/assert"
)

func TestBookReviewResponse(t *testing.T) {
	t.Parallel()
	type args struct {
		review *entity.Review
		user   *entity.User
	}
	tests := []struct {
		name   string
		args   args
		expect *BookReviewResponse
	}{
		{
			name: "success",
			args: args{
				review: &entity.Review{
					Review: &book.Review{
						Id:         1,
						BookId:     1,
						UserId:     "00000000-0000-0000-0000-000000000000",
						Score:      3,
						Impression: "テストレビューです",
						CreatedAt:  test.TimeMock,
						UpdatedAt:  test.TimeMock,
					},
				},
				user: &entity.User{
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
						CreatedAt:        test.TimeMock,
						UpdatedAt:        test.TimeMock,
					},
				},
			},
			expect: &BookReviewResponse{
				ID:         1,
				Impression: "テストレビューです",
				User: &bookReviewUser{
					ID:           "00000000-0000-0000-0000-000000000000",
					Username:     "テストユーザー",
					ThumbnailURL: "https://go.dev/images/gophers/ladder.svg",
				},
				CreatedAt: test.TimeMock,
				UpdatedAt: test.TimeMock,
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			actual := NewBookReviewResponse(tt.args.review, tt.args.user)
			assert.Equal(t, tt.expect, actual)
		})
	}
}

func TestBookReviewListResponse(t *testing.T) {
	t.Parallel()
	type args struct {
		reviews entity.Reviews
		users   map[string]*entity.User
		limit   int64
		offset  int64
		total   int64
	}
	tests := []struct {
		name   string
		args   args
		expect *BookReviewListResponse
	}{
		{
			name: "success",
			args: args{
				reviews: entity.Reviews{
					{
						Review: &book.Review{
							Id:         1,
							BookId:     1,
							UserId:     "00000000-0000-0000-0000-000000000000",
							Score:      3,
							Impression: "テストレビューです",
							CreatedAt:  test.TimeMock,
							UpdatedAt:  test.TimeMock,
						},
					},
					{
						Review: &book.Review{
							Id:         2,
							BookId:     2,
							UserId:     "00000000-0000-0000-0000-000000000000",
							Score:      3,
							Impression: "テストレビューです",
							CreatedAt:  test.TimeMock,
							UpdatedAt:  test.TimeMock,
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
							CreatedAt:        test.TimeMock,
							UpdatedAt:        test.TimeMock,
						},
					},
				},
				limit:  100,
				offset: 0,
				total:  2,
			},
			expect: &BookReviewListResponse{
				Reviews: []*bookReviewListReview{
					{
						ID:         1,
						Impression: "テストレビューです",
						User: &bookReviewListUser{
							ID:           "00000000-0000-0000-0000-000000000000",
							Username:     "テストユーザー",
							ThumbnailURL: "https://go.dev/images/gophers/ladder.svg",
						},
						CreatedAt: test.TimeMock,
						UpdatedAt: test.TimeMock,
					},
					{
						ID:         2,
						Impression: "テストレビューです",
						User: &bookReviewListUser{
							ID:           "00000000-0000-0000-0000-000000000000",
							Username:     "テストユーザー",
							ThumbnailURL: "https://go.dev/images/gophers/ladder.svg",
						},
						CreatedAt: test.TimeMock,
						UpdatedAt: test.TimeMock,
					},
				},
				Limit:  100,
				Offset: 0,
				Total:  2,
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			actual := NewBookReviewListResponse(
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

func TestUserReviewResponse(t *testing.T) {
	t.Parallel()
	type args struct {
		review *entity.Review
		book   *entity.Book
	}
	tests := []struct {
		name   string
		args   args
		expect *UserReviewResponse
	}{
		{
			name: "success",
			args: args{
				review: &entity.Review{
					Review: &book.Review{
						Id:         1,
						BookId:     1,
						UserId:     "00000000-0000-0000-0000-000000000000",
						Score:      3,
						Impression: "テストレビューです",
						CreatedAt:  test.TimeMock,
						UpdatedAt:  test.TimeMock,
					},
				},
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
						CreatedAt:      test.TimeMock,
						UpdatedAt:      test.TimeMock,
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
			expect: &UserReviewResponse{
				ID:         1,
				Impression: "テストレビューです",
				Book: &userReviewBook{
					ID:           1,
					Title:        "小説　ちはやふる　上の句",
					ThumbnailURL: "https://thumbnail.image.rakuten.co.jp/@0_mall/book/cabinet/8426/9784062938426.jpg?_ex=120x120",
				},
				CreatedAt: test.TimeMock,
				UpdatedAt: test.TimeMock,
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			actual := NewUserReviewResponse(tt.args.review, tt.args.book)
			assert.Equal(t, tt.expect, actual)
		})
	}
}

func TestUserReviewListResponse(t *testing.T) {
	t.Parallel()
	type args struct {
		reviews entity.Reviews
		books   map[int64]*entity.Book
		limit   int64
		offset  int64
		total   int64
	}
	tests := []struct {
		name   string
		args   args
		expect *UserReviewListResponse
	}{
		{
			name: "success",
			args: args{
				reviews: entity.Reviews{
					{
						Review: &book.Review{
							Id:         1,
							BookId:     1,
							UserId:     "00000000-0000-0000-0000-000000000000",
							Score:      3,
							Impression: "テストレビューです",
							CreatedAt:  test.TimeMock,
							UpdatedAt:  test.TimeMock,
						},
					},
					{
						Review: &book.Review{
							Id:         2,
							BookId:     2,
							UserId:     "00000000-0000-0000-0000-000000000000",
							Score:      3,
							Impression: "テストレビューです",
							CreatedAt:  test.TimeMock,
							UpdatedAt:  test.TimeMock,
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
							CreatedAt:      test.TimeMock,
							UpdatedAt:      test.TimeMock,
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
							CreatedAt:      test.TimeMock,
							UpdatedAt:      test.TimeMock,
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
				limit:  100,
				offset: 0,
				total:  2,
			},
			expect: &UserReviewListResponse{
				Reviews: []*userReviewListReview{
					{
						ID:         1,
						Impression: "テストレビューです",
						Book: &userReviewListBook{
							ID:           1,
							Title:        "小説　ちはやふる　上の句",
							ThumbnailURL: "https://thumbnail.image.rakuten.co.jp/@0_mall/book/cabinet/8426/9784062938426.jpg?_ex=120x120",
						},
						CreatedAt: test.TimeMock,
						UpdatedAt: test.TimeMock,
					},
					{
						ID:         2,
						Impression: "テストレビューです",
						Book: &userReviewListBook{
							ID:           2,
							Title:        "小説　ちはやふる　上の句",
							ThumbnailURL: "https://thumbnail.image.rakuten.co.jp/@0_mall/book/cabinet/8426/9784062938426.jpg?_ex=120x120",
						},
						CreatedAt: test.TimeMock,
						UpdatedAt: test.TimeMock,
					},
				},
				Limit:  100,
				Offset: 0,
				Total:  2,
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			actual := NewUserReviewListResponse(
				tt.args.reviews,
				tt.args.books,
				tt.args.limit,
				tt.args.offset,
				tt.args.total,
			)
			assert.Equal(t, tt.expect, actual)
		})
	}
}
