package validation

import (
	"strings"
	"testing"

	pb "github.com/calmato/gran-book/api/server/book/proto"
	"github.com/stretchr/testify/assert"
)

func TestBookRequestValidation_ListBookshelf(t *testing.T) {
	t.Parallel()
	type args struct {
		req *pb.ListBookshelfRequest
	}

	testCases := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "success",
			args: args{
				req: &pb.ListBookshelfRequest{
					UserId: "12345678-1234-1234-1234-123456789012",
					Limit:  200,
					Offset: 100,
				},
			},
			want: true,
		},
		{
			name: "validation error: UserId.lte",
			args: args{
				req: &pb.ListBookshelfRequest{
					UserId: "",
					Order: &pb.Order{
						Field:   "created_at",
						OrderBy: pb.OrderBy_ORDER_BY_ASC,
					},
					Limit:  200,
					Offset: 100,
				},
			},
			want: false,
		},
		{
			name: "validation error: Order.OrderBy.in",
			args: args{
				req: &pb.ListBookshelfRequest{
					UserId: "12345678-1234-1234-1234-123456789012",
					Order: &pb.Order{
						Field:   "created_at",
						OrderBy: 2,
					},
					Limit:  200,
					Offset: 100,
				},
			},
			want: false,
		},
		{
			name: "validation error: Limit.lte",
			args: args{
				req: &pb.ListBookshelfRequest{
					UserId: "12345678-1234-1234-1234-123456789012",
					Order: &pb.Order{
						Field:   "created_at",
						OrderBy: pb.OrderBy_ORDER_BY_ASC,
					},
					Limit:  201,
					Offset: 100,
				},
			},
			want: false,
		},
		{
			name: "validation error: Offset.gte",
			args: args{
				req: &pb.ListBookshelfRequest{
					UserId: "12345678-1234-1234-1234-123456789012",
					Order: &pb.Order{
						Field:   "created_at",
						OrderBy: pb.OrderBy_ORDER_BY_ASC,
					},
					Limit:  200,
					Offset: -1,
				},
			},
			want: false,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			target := NewBookRequestValidation()

			got := target.ListBookshelf(tt.args.req)
			switch tt.want {
			case true:
				assert.NoError(t, got)
			case false:
				assert.Error(t, got)
			}
		})
	}
}

func TestBookRequestValidation_ListBookReview(t *testing.T) {
	t.Parallel()
	type args struct {
		req *pb.ListBookReviewRequest
	}

	testCases := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "success",
			args: args{
				req: &pb.ListBookReviewRequest{
					BookId: 1,
					Limit:  200,
					Offset: 100,
				},
			},
			want: true,
		},
		{
			name: "validation error: BookId.gt",
			args: args{
				req: &pb.ListBookReviewRequest{
					BookId: 0,
					Order: &pb.Order{
						Field:   "created_at",
						OrderBy: pb.OrderBy_ORDER_BY_ASC,
					},
					Limit:  200,
					Offset: 100,
				},
			},
			want: false,
		},
		{
			name: "validation error: Order.OrderBy.in",
			args: args{
				req: &pb.ListBookReviewRequest{
					BookId: 1,
					Order: &pb.Order{
						Field:   "created_at",
						OrderBy: 2,
					},
					Limit:  200,
					Offset: 100,
				},
			},
			want: false,
		},
		{
			name: "validation error: Limit.lte",
			args: args{
				req: &pb.ListBookReviewRequest{
					BookId: 1,
					Order: &pb.Order{
						Field:   "created_at",
						OrderBy: pb.OrderBy_ORDER_BY_ASC,
					},
					Limit:  201,
					Offset: 100,
				},
			},
			want: false,
		},
		{
			name: "validation error: Offset.gte",
			args: args{
				req: &pb.ListBookReviewRequest{
					BookId: 1,
					Order: &pb.Order{
						Field:   "created_at",
						OrderBy: pb.OrderBy_ORDER_BY_ASC,
					},
					Limit:  200,
					Offset: -1,
				},
			},
			want: false,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			target := NewBookRequestValidation()

			got := target.ListBookReview(tt.args.req)
			switch tt.want {
			case true:
				assert.NoError(t, got)
			case false:
				assert.Error(t, got)
			}
		})
	}
}

func TestBookRequestValidation_ListUserReview(t *testing.T) {
	t.Parallel()
	type args struct {
		req *pb.ListUserReviewRequest
	}

	testCases := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "success",
			args: args{
				req: &pb.ListUserReviewRequest{
					UserId: "12345678-1234-1234-1234-123456789012",
					Limit:  200,
					Offset: 100,
				},
			},
			want: true,
		},
		{
			name: "validation error: UserId.min_len",
			args: args{
				req: &pb.ListUserReviewRequest{
					UserId: "",
					Order: &pb.Order{
						Field:   "created_at",
						OrderBy: pb.OrderBy_ORDER_BY_ASC,
					},
					Limit:  200,
					Offset: 100,
				},
			},
			want: false,
		},
		{
			name: "validation error: Order.OrderBy.in",
			args: args{
				req: &pb.ListUserReviewRequest{
					UserId: "12345678-1234-1234-1234-123456789012",
					Order: &pb.Order{
						Field:   "created_at",
						OrderBy: 2,
					},
					Limit:  200,
					Offset: 100,
				},
			},
			want: false,
		},
		{
			name: "validation error: Limit.lte",
			args: args{
				req: &pb.ListUserReviewRequest{
					UserId: "12345678-1234-1234-1234-123456789012",
					Order: &pb.Order{
						Field:   "created_at",
						OrderBy: pb.OrderBy_ORDER_BY_ASC,
					},
					Limit:  201,
					Offset: 100,
				},
			},
			want: false,
		},
		{
			name: "validation error: Offset.gte",
			args: args{
				req: &pb.ListUserReviewRequest{
					UserId: "12345678-1234-1234-1234-123456789012",
					Order: &pb.Order{
						Field:   "created_at",
						OrderBy: pb.OrderBy_ORDER_BY_ASC,
					},
					Limit:  200,
					Offset: -1,
				},
			},
			want: false,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			target := NewBookRequestValidation()

			got := target.ListUserReview(tt.args.req)
			switch tt.want {
			case true:
				assert.NoError(t, got)
			case false:
				assert.Error(t, got)
			}
		})
	}
}

func TestBookRequestValidation_MultiGetBooks(t *testing.T) {
	t.Parallel()
	type args struct {
		req *pb.MultiGetBooksRequest
	}

	testCases := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "success",
			args: args{
				req: &pb.MultiGetBooksRequest{
					BookIds: []int64{1, 2},
				},
			},
			want: true,
		},
		{
			name: "validation error: BookIds.min_items",
			args: args{
				req: &pb.MultiGetBooksRequest{
					BookIds: []int64{},
				},
			},
			want: false,
		},
		{
			name: "validation error: BookIds.max_items",
			args: args{
				req: &pb.MultiGetBooksRequest{
					BookIds: func() []int64 {
						bookIDs := make([]int64, 201)
						for i := 0; i < len(bookIDs); i++ {
							bookIDs[i] = int64(i)
						}
						return bookIDs
					}(),
				},
			},
			want: false,
		},
		{
			name: "validation error: BookIds.unique",
			args: args{
				req: &pb.MultiGetBooksRequest{
					BookIds: []int64{1, 1},
				},
			},
			want: false,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			target := NewBookRequestValidation()

			got := target.MultiGetBooks(tt.args.req)
			switch tt.want {
			case true:
				assert.NoError(t, got)
			case false:
				assert.Error(t, got)
			}
		})
	}
}

func TestBookRequestValidation_GetBook(t *testing.T) {
	t.Parallel()
	type args struct {
		req *pb.GetBookRequest
	}

	testCases := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "success",
			args: args{
				req: &pb.GetBookRequest{
					BookId: 1,
				},
			},
			want: true,
		},
		{
			name: "validation error: BookIds.gt",
			args: args{
				req: &pb.GetBookRequest{
					BookId: 0,
				},
			},
			want: false,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			target := NewBookRequestValidation()

			got := target.GetBook(tt.args.req)
			switch tt.want {
			case true:
				assert.NoError(t, got)
			case false:
				assert.Error(t, got)
			}
		})
	}
}

func TestBookRequestValidation_GetBookByIsbn(t *testing.T) {
	t.Parallel()
	type args struct {
		req *pb.GetBookByIsbnRequest
	}

	testCases := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "success",
			args: args{
				req: &pb.GetBookByIsbnRequest{
					Isbn: "9784062938426",
				},
			},
			want: true,
		},
		{
			name: "validation error: Isbn.min_len",
			args: args{
				req: &pb.GetBookByIsbnRequest{
					Isbn: "123456789",
				},
			},
			want: false,
		},
		{
			name: "validation error: Isbn.max_len",
			args: args{
				req: &pb.GetBookByIsbnRequest{
					Isbn: "12345678901234567",
				},
			},
			want: false,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			target := NewBookRequestValidation()

			got := target.GetBookByIsbn(tt.args.req)
			switch tt.want {
			case true:
				assert.NoError(t, got)
			case false:
				assert.Error(t, got)
			}
		})
	}
}

func TestBookRequestValidation_GetBookshelf(t *testing.T) {
	t.Parallel()
	type args struct {
		req *pb.GetBookshelfRequest
	}

	testCases := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "success",
			args: args{
				req: &pb.GetBookshelfRequest{
					UserId: "12345678-1234-1234-1234-123456789012",
					BookId: 1,
				},
			},
			want: true,
		},
		{
			name: "validation error: UserId.min_len",
			args: args{
				req: &pb.GetBookshelfRequest{
					UserId: "",
					BookId: 1,
				},
			},
			want: false,
		},
		{
			name: "validation error: BookId.gt",
			args: args{
				req: &pb.GetBookshelfRequest{
					UserId: "12345678-1234-1234-1234-123456789012",
					BookId: 0,
				},
			},
			want: false,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			target := NewBookRequestValidation()

			got := target.GetBookshelf(tt.args.req)
			switch tt.want {
			case true:
				assert.NoError(t, got)
			case false:
				assert.Error(t, got)
			}
		})
	}
}

func TestBookRequestValidation_GetReview(t *testing.T) {
	t.Parallel()
	type args struct {
		req *pb.GetReviewRequest
	}

	testCases := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "success",
			args: args{
				req: &pb.GetReviewRequest{
					ReviewId: 1,
				},
			},
			want: true,
		},
		{
			name: "validation error: ReviewId.gt",
			args: args{
				req: &pb.GetReviewRequest{
					ReviewId: 0,
				},
			},
			want: false,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			target := NewBookRequestValidation()

			got := target.GetReview(tt.args.req)
			switch tt.want {
			case true:
				assert.NoError(t, got)
			case false:
				assert.Error(t, got)
			}
		})
	}
}

func TestBookRequestValidation_CreateBook(t *testing.T) {
	t.Parallel()
	type args struct {
		req *pb.CreateBookRequest
	}

	testCases := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "success",
			args: args{
				req: &pb.CreateBookRequest{
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
					Authors: []*pb.CreateBookRequest_Author{
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
			want: true,
		},
		{
			name: "validation error: Title.min_len",
			args: args{
				req: &pb.CreateBookRequest{
					Title:          "",
					TitleKana:      "ショウセツ チハヤフルカミノク",
					Description:    "綾瀬千早は高校入学と同時に、競技かるた部を作ろうと奔走する。幼馴染の太一と仲間を集め、夏の全国大会に出場するためだ。強くなって、新と再会したい。幼い頃かるたを取り合った、新に寄せる千早の秘めた想いに気づきながらも、太一は千早を守り立てる。それぞれの青春を懸けた、一途な情熱の物語が幕開ける。",
					Isbn:           "9784062938426",
					Publisher:      "講談社",
					PublishedOn:    "2018-01-16",
					ThumbnailUrl:   "https://thumbnail.image.rakuten.co.jp/@0_mall/book/cabinet/8426/9784062938426.jpg?_ex=120x120",
					RakutenUrl:     "https://books.rakuten.co.jp/rb/15271426/",
					RakutenSize:    "コミック",
					RakutenGenreId: "001004008001/001004008003/001019001",
					Authors: []*pb.CreateBookRequest_Author{
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
			want: false,
		},
		{
			name: "validation error: Title.max_len",
			args: args{
				req: &pb.CreateBookRequest{
					Title:          strings.Repeat("x", 65),
					TitleKana:      "ショウセツ チハヤフルカミノク",
					Description:    "綾瀬千早は高校入学と同時に、競技かるた部を作ろうと奔走する。幼馴染の太一と仲間を集め、夏の全国大会に出場するためだ。強くなって、新と再会したい。幼い頃かるたを取り合った、新に寄せる千早の秘めた想いに気づきながらも、太一は千早を守り立てる。それぞれの青春を懸けた、一途な情熱の物語が幕開ける。",
					Isbn:           "9784062938426",
					Publisher:      "講談社",
					PublishedOn:    "2018-01-16",
					ThumbnailUrl:   "https://thumbnail.image.rakuten.co.jp/@0_mall/book/cabinet/8426/9784062938426.jpg?_ex=120x120",
					RakutenUrl:     "https://books.rakuten.co.jp/rb/15271426/",
					RakutenSize:    "コミック",
					RakutenGenreId: "001004008001/001004008003/001019001",
					Authors: []*pb.CreateBookRequest_Author{
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
			want: false,
		},
		{
			name: "validation error: TitleKana.min_len",
			args: args{
				req: &pb.CreateBookRequest{
					Title:          "小説　ちはやふる　上の句",
					TitleKana:      "",
					Description:    "綾瀬千早は高校入学と同時に、競技かるた部を作ろうと奔走する。幼馴染の太一と仲間を集め、夏の全国大会に出場するためだ。強くなって、新と再会したい。幼い頃かるたを取り合った、新に寄せる千早の秘めた想いに気づきながらも、太一は千早を守り立てる。それぞれの青春を懸けた、一途な情熱の物語が幕開ける。",
					Isbn:           "9784062938426",
					Publisher:      "講談社",
					PublishedOn:    "2018-01-16",
					ThumbnailUrl:   "https://thumbnail.image.rakuten.co.jp/@0_mall/book/cabinet/8426/9784062938426.jpg?_ex=120x120",
					RakutenUrl:     "https://books.rakuten.co.jp/rb/15271426/",
					RakutenSize:    "コミック",
					RakutenGenreId: "001004008001/001004008003/001019001",
					Authors: []*pb.CreateBookRequest_Author{
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
			want: false,
		},
		{
			name: "validation error: TitleKana.max_len",
			args: args{
				req: &pb.CreateBookRequest{
					Title:          "小説　ちはやふる　上の句",
					TitleKana:      strings.Repeat("x", 129),
					Description:    "綾瀬千早は高校入学と同時に、競技かるた部を作ろうと奔走する。幼馴染の太一と仲間を集め、夏の全国大会に出場するためだ。強くなって、新と再会したい。幼い頃かるたを取り合った、新に寄せる千早の秘めた想いに気づきながらも、太一は千早を守り立てる。それぞれの青春を懸けた、一途な情熱の物語が幕開ける。",
					Isbn:           "9784062938426",
					Publisher:      "講談社",
					PublishedOn:    "2018-01-16",
					ThumbnailUrl:   "https://thumbnail.image.rakuten.co.jp/@0_mall/book/cabinet/8426/9784062938426.jpg?_ex=120x120",
					RakutenUrl:     "https://books.rakuten.co.jp/rb/15271426/",
					RakutenSize:    "コミック",
					RakutenGenreId: "001004008001/001004008003/001019001",
					Authors: []*pb.CreateBookRequest_Author{
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
			want: false,
		},
		{
			name: "validation error: Description.max_len",
			args: args{
				req: &pb.CreateBookRequest{
					Title:          "小説　ちはやふる　上の句",
					TitleKana:      "ショウセツ チハヤフルカミノク",
					Description:    strings.Repeat("x", 2001),
					Isbn:           "9784062938426",
					Publisher:      "講談社",
					PublishedOn:    "2018-01-16",
					ThumbnailUrl:   "https://thumbnail.image.rakuten.co.jp/@0_mall/book/cabinet/8426/9784062938426.jpg?_ex=120x120",
					RakutenUrl:     "https://books.rakuten.co.jp/rb/15271426/",
					RakutenSize:    "コミック",
					RakutenGenreId: "001004008001/001004008003/001019001",
					Authors: []*pb.CreateBookRequest_Author{
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
			want: false,
		},
		{
			name: "validation error: Isbn.min_len",
			args: args{
				req: &pb.CreateBookRequest{
					Title:          "小説　ちはやふる　上の句",
					TitleKana:      "ショウセツ チハヤフルカミノク",
					Description:    "綾瀬千早は高校入学と同時に、競技かるた部を作ろうと奔走する。幼馴染の太一と仲間を集め、夏の全国大会に出場するためだ。強くなって、新と再会したい。幼い頃かるたを取り合った、新に寄せる千早の秘めた想いに気づきながらも、太一は千早を守り立てる。それぞれの青春を懸けた、一途な情熱の物語が幕開ける。",
					Isbn:           strings.Repeat("x", 9),
					Publisher:      "講談社",
					PublishedOn:    "2018-01-16",
					ThumbnailUrl:   "https://thumbnail.image.rakuten.co.jp/@0_mall/book/cabinet/8426/9784062938426.jpg?_ex=120x120",
					RakutenUrl:     "https://books.rakuten.co.jp/rb/15271426/",
					RakutenSize:    "コミック",
					RakutenGenreId: "001004008001/001004008003/001019001",
					Authors: []*pb.CreateBookRequest_Author{
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
			want: false,
		},
		{
			name: "validation error: Isbn.max_len",
			args: args{
				req: &pb.CreateBookRequest{
					Title:          "小説　ちはやふる　上の句",
					TitleKana:      "ショウセツ チハヤフルカミノク",
					Description:    "綾瀬千早は高校入学と同時に、競技かるた部を作ろうと奔走する。幼馴染の太一と仲間を集め、夏の全国大会に出場するためだ。強くなって、新と再会したい。幼い頃かるたを取り合った、新に寄せる千早の秘めた想いに気づきながらも、太一は千早を守り立てる。それぞれの青春を懸けた、一途な情熱の物語が幕開ける。",
					Isbn:           strings.Repeat("x", 18),
					Publisher:      "講談社",
					PublishedOn:    "2018-01-16",
					ThumbnailUrl:   "https://thumbnail.image.rakuten.co.jp/@0_mall/book/cabinet/8426/9784062938426.jpg?_ex=120x120",
					RakutenUrl:     "https://books.rakuten.co.jp/rb/15271426/",
					RakutenSize:    "コミック",
					RakutenGenreId: "001004008001/001004008003/001019001",
					Authors: []*pb.CreateBookRequest_Author{
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
			want: false,
		},
		{
			name: "validation error: Publisher.min_len",
			args: args{
				req: &pb.CreateBookRequest{
					Title:          "小説　ちはやふる　上の句",
					TitleKana:      "ショウセツ チハヤフルカミノク",
					Description:    "綾瀬千早は高校入学と同時に、競技かるた部を作ろうと奔走する。幼馴染の太一と仲間を集め、夏の全国大会に出場するためだ。強くなって、新と再会したい。幼い頃かるたを取り合った、新に寄せる千早の秘めた想いに気づきながらも、太一は千早を守り立てる。それぞれの青春を懸けた、一途な情熱の物語が幕開ける。",
					Isbn:           "9784062938426",
					Publisher:      "",
					PublishedOn:    "2018-01-16",
					ThumbnailUrl:   "https://thumbnail.image.rakuten.co.jp/@0_mall/book/cabinet/8426/9784062938426.jpg?_ex=120x120",
					RakutenUrl:     "https://books.rakuten.co.jp/rb/15271426/",
					RakutenSize:    "コミック",
					RakutenGenreId: "001004008001/001004008003/001019001",
					Authors: []*pb.CreateBookRequest_Author{
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
			want: false,
		},
		{
			name: "validation error: Publisher.max_len",
			args: args{
				req: &pb.CreateBookRequest{
					Title:          "小説　ちはやふる　上の句",
					TitleKana:      "ショウセツ チハヤフルカミノク",
					Description:    "綾瀬千早は高校入学と同時に、競技かるた部を作ろうと奔走する。幼馴染の太一と仲間を集め、夏の全国大会に出場するためだ。強くなって、新と再会したい。幼い頃かるたを取り合った、新に寄せる千早の秘めた想いに気づきながらも、太一は千早を守り立てる。それぞれの青春を懸けた、一途な情熱の物語が幕開ける。",
					Isbn:           "9784062938426",
					Publisher:      strings.Repeat("x", 33),
					PublishedOn:    "2018-01-16",
					ThumbnailUrl:   "https://thumbnail.image.rakuten.co.jp/@0_mall/book/cabinet/8426/9784062938426.jpg?_ex=120x120",
					RakutenUrl:     "https://books.rakuten.co.jp/rb/15271426/",
					RakutenSize:    "コミック",
					RakutenGenreId: "001004008001/001004008003/001019001",
					Authors: []*pb.CreateBookRequest_Author{
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
			want: false,
		},
		{
			name: "validation error: PublishedOn.min_len",
			args: args{
				req: &pb.CreateBookRequest{
					Title:          "小説　ちはやふる　上の句",
					TitleKana:      "ショウセツ チハヤフルカミノク",
					Description:    "綾瀬千早は高校入学と同時に、競技かるた部を作ろうと奔走する。幼馴染の太一と仲間を集め、夏の全国大会に出場するためだ。強くなって、新と再会したい。幼い頃かるたを取り合った、新に寄せる千早の秘めた想いに気づきながらも、太一は千早を守り立てる。それぞれの青春を懸けた、一途な情熱の物語が幕開ける。",
					Isbn:           "9784062938426",
					Publisher:      "講談社",
					PublishedOn:    "",
					ThumbnailUrl:   "https://thumbnail.image.rakuten.co.jp/@0_mall/book/cabinet/8426/9784062938426.jpg?_ex=120x120",
					RakutenUrl:     "https://books.rakuten.co.jp/rb/15271426/",
					RakutenSize:    "コミック",
					RakutenGenreId: "001004008001/001004008003/001019001",
					Authors: []*pb.CreateBookRequest_Author{
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
			want: false,
		},
		{
			name: "validation error: PublishedOn.max_len",
			args: args{
				req: &pb.CreateBookRequest{
					Title:          "小説　ちはやふる　上の句",
					TitleKana:      "ショウセツ チハヤフルカミノク",
					Description:    "綾瀬千早は高校入学と同時に、競技かるた部を作ろうと奔走する。幼馴染の太一と仲間を集め、夏の全国大会に出場するためだ。強くなって、新と再会したい。幼い頃かるたを取り合った、新に寄せる千早の秘めた想いに気づきながらも、太一は千早を守り立てる。それぞれの青春を懸けた、一途な情熱の物語が幕開ける。",
					Isbn:           "9784062938426",
					Publisher:      "講談社",
					PublishedOn:    strings.Repeat("x", 17),
					ThumbnailUrl:   "https://thumbnail.image.rakuten.co.jp/@0_mall/book/cabinet/8426/9784062938426.jpg?_ex=120x120",
					RakutenUrl:     "https://books.rakuten.co.jp/rb/15271426/",
					RakutenSize:    "コミック",
					RakutenGenreId: "001004008001/001004008003/001019001",
					Authors: []*pb.CreateBookRequest_Author{
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
			want: false,
		},
		{
			name: "validation error: Authors.Name.min_len",
			args: args{
				req: &pb.CreateBookRequest{
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
					Authors: []*pb.CreateBookRequest_Author{
						{
							Name:     "",
							NameKana: "アリサワ ユウキ",
						},
						{
							Name:     "末次 由紀",
							NameKana: "スエツグ ユキ",
						},
					},
				},
			},
			want: false,
		},
		{
			name: "validation error: Authors.Name.max_len",
			args: args{
				req: &pb.CreateBookRequest{
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
					Authors: []*pb.CreateBookRequest_Author{
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
			},
			want: false,
		},
		{
			name: "validation error: Authors.NameKana.min_len",
			args: args{
				req: &pb.CreateBookRequest{
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
					Authors: []*pb.CreateBookRequest_Author{
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
			},
			want: false,
		},
		{
			name: "validation error: Authors.NameKana.max_len",
			args: args{
				req: &pb.CreateBookRequest{
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
					Authors: []*pb.CreateBookRequest_Author{
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
			},
			want: false,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			target := NewBookRequestValidation()

			got := target.CreateBook(tt.args.req)
			switch tt.want {
			case true:
				assert.NoError(t, got)
			case false:
				assert.Error(t, got)
			}
		})
	}
}

func TestBookRequestValidation_UpdateBook(t *testing.T) {
	t.Parallel()
	type args struct {
		req *pb.UpdateBookRequest
	}

	testCases := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "success",
			args: args{
				req: &pb.UpdateBookRequest{
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
					Authors: []*pb.UpdateBookRequest_Author{
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
			want: true,
		},
		{
			name: "validation error: Title.min_len",
			args: args{
				req: &pb.UpdateBookRequest{
					Title:          "",
					TitleKana:      "ショウセツ チハヤフルカミノク",
					Description:    "綾瀬千早は高校入学と同時に、競技かるた部を作ろうと奔走する。幼馴染の太一と仲間を集め、夏の全国大会に出場するためだ。強くなって、新と再会したい。幼い頃かるたを取り合った、新に寄せる千早の秘めた想いに気づきながらも、太一は千早を守り立てる。それぞれの青春を懸けた、一途な情熱の物語が幕開ける。",
					Isbn:           "9784062938426",
					Publisher:      "講談社",
					PublishedOn:    "2018-01-16",
					ThumbnailUrl:   "https://thumbnail.image.rakuten.co.jp/@0_mall/book/cabinet/8426/9784062938426.jpg?_ex=120x120",
					RakutenUrl:     "https://books.rakuten.co.jp/rb/15271426/",
					RakutenSize:    "コミック",
					RakutenGenreId: "001004008001/001004008003/001019001",
					Authors: []*pb.UpdateBookRequest_Author{
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
			want: false,
		},
		{
			name: "validation error: Title.max_len",
			args: args{
				req: &pb.UpdateBookRequest{
					Title:          strings.Repeat("x", 65),
					TitleKana:      "ショウセツ チハヤフルカミノク",
					Description:    "綾瀬千早は高校入学と同時に、競技かるた部を作ろうと奔走する。幼馴染の太一と仲間を集め、夏の全国大会に出場するためだ。強くなって、新と再会したい。幼い頃かるたを取り合った、新に寄せる千早の秘めた想いに気づきながらも、太一は千早を守り立てる。それぞれの青春を懸けた、一途な情熱の物語が幕開ける。",
					Isbn:           "9784062938426",
					Publisher:      "講談社",
					PublishedOn:    "2018-01-16",
					ThumbnailUrl:   "https://thumbnail.image.rakuten.co.jp/@0_mall/book/cabinet/8426/9784062938426.jpg?_ex=120x120",
					RakutenUrl:     "https://books.rakuten.co.jp/rb/15271426/",
					RakutenSize:    "コミック",
					RakutenGenreId: "001004008001/001004008003/001019001",
					Authors: []*pb.UpdateBookRequest_Author{
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
			want: false,
		},
		{
			name: "validation error: TitleKana.min_len",
			args: args{
				req: &pb.UpdateBookRequest{
					Title:          "小説　ちはやふる　上の句",
					TitleKana:      "",
					Description:    "綾瀬千早は高校入学と同時に、競技かるた部を作ろうと奔走する。幼馴染の太一と仲間を集め、夏の全国大会に出場するためだ。強くなって、新と再会したい。幼い頃かるたを取り合った、新に寄せる千早の秘めた想いに気づきながらも、太一は千早を守り立てる。それぞれの青春を懸けた、一途な情熱の物語が幕開ける。",
					Isbn:           "9784062938426",
					Publisher:      "講談社",
					PublishedOn:    "2018-01-16",
					ThumbnailUrl:   "https://thumbnail.image.rakuten.co.jp/@0_mall/book/cabinet/8426/9784062938426.jpg?_ex=120x120",
					RakutenUrl:     "https://books.rakuten.co.jp/rb/15271426/",
					RakutenSize:    "コミック",
					RakutenGenreId: "001004008001/001004008003/001019001",
					Authors: []*pb.UpdateBookRequest_Author{
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
			want: false,
		},
		{
			name: "validation error: TitleKana.max_len",
			args: args{
				req: &pb.UpdateBookRequest{
					Title:          "小説　ちはやふる　上の句",
					TitleKana:      strings.Repeat("x", 129),
					Description:    "綾瀬千早は高校入学と同時に、競技かるた部を作ろうと奔走する。幼馴染の太一と仲間を集め、夏の全国大会に出場するためだ。強くなって、新と再会したい。幼い頃かるたを取り合った、新に寄せる千早の秘めた想いに気づきながらも、太一は千早を守り立てる。それぞれの青春を懸けた、一途な情熱の物語が幕開ける。",
					Isbn:           "9784062938426",
					Publisher:      "講談社",
					PublishedOn:    "2018-01-16",
					ThumbnailUrl:   "https://thumbnail.image.rakuten.co.jp/@0_mall/book/cabinet/8426/9784062938426.jpg?_ex=120x120",
					RakutenUrl:     "https://books.rakuten.co.jp/rb/15271426/",
					RakutenSize:    "コミック",
					RakutenGenreId: "001004008001/001004008003/001019001",
					Authors: []*pb.UpdateBookRequest_Author{
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
			want: false,
		},
		{
			name: "validation error: Description.max_len",
			args: args{
				req: &pb.UpdateBookRequest{
					Title:          "小説　ちはやふる　上の句",
					TitleKana:      "ショウセツ チハヤフルカミノク",
					Description:    strings.Repeat("x", 2001),
					Isbn:           "9784062938426",
					Publisher:      "講談社",
					PublishedOn:    "2018-01-16",
					ThumbnailUrl:   "https://thumbnail.image.rakuten.co.jp/@0_mall/book/cabinet/8426/9784062938426.jpg?_ex=120x120",
					RakutenUrl:     "https://books.rakuten.co.jp/rb/15271426/",
					RakutenSize:    "コミック",
					RakutenGenreId: "001004008001/001004008003/001019001",
					Authors: []*pb.UpdateBookRequest_Author{
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
			want: false,
		},
		{
			name: "validation error: Isbn.min_len",
			args: args{
				req: &pb.UpdateBookRequest{
					Title:          "小説　ちはやふる　上の句",
					TitleKana:      "ショウセツ チハヤフルカミノク",
					Description:    "綾瀬千早は高校入学と同時に、競技かるた部を作ろうと奔走する。幼馴染の太一と仲間を集め、夏の全国大会に出場するためだ。強くなって、新と再会したい。幼い頃かるたを取り合った、新に寄せる千早の秘めた想いに気づきながらも、太一は千早を守り立てる。それぞれの青春を懸けた、一途な情熱の物語が幕開ける。",
					Isbn:           strings.Repeat("x", 9),
					Publisher:      "講談社",
					PublishedOn:    "2018-01-16",
					ThumbnailUrl:   "https://thumbnail.image.rakuten.co.jp/@0_mall/book/cabinet/8426/9784062938426.jpg?_ex=120x120",
					RakutenUrl:     "https://books.rakuten.co.jp/rb/15271426/",
					RakutenSize:    "コミック",
					RakutenGenreId: "001004008001/001004008003/001019001",
					Authors: []*pb.UpdateBookRequest_Author{
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
			want: false,
		},
		{
			name: "validation error: Isbn.max_len",
			args: args{
				req: &pb.UpdateBookRequest{
					Title:          "小説　ちはやふる　上の句",
					TitleKana:      "ショウセツ チハヤフルカミノク",
					Description:    "綾瀬千早は高校入学と同時に、競技かるた部を作ろうと奔走する。幼馴染の太一と仲間を集め、夏の全国大会に出場するためだ。強くなって、新と再会したい。幼い頃かるたを取り合った、新に寄せる千早の秘めた想いに気づきながらも、太一は千早を守り立てる。それぞれの青春を懸けた、一途な情熱の物語が幕開ける。",
					Isbn:           strings.Repeat("x", 18),
					Publisher:      "講談社",
					PublishedOn:    "2018-01-16",
					ThumbnailUrl:   "https://thumbnail.image.rakuten.co.jp/@0_mall/book/cabinet/8426/9784062938426.jpg?_ex=120x120",
					RakutenUrl:     "https://books.rakuten.co.jp/rb/15271426/",
					RakutenSize:    "コミック",
					RakutenGenreId: "001004008001/001004008003/001019001",
					Authors: []*pb.UpdateBookRequest_Author{
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
			want: false,
		},
		{
			name: "validation error: Publisher.min_len",
			args: args{
				req: &pb.UpdateBookRequest{
					Title:          "小説　ちはやふる　上の句",
					TitleKana:      "ショウセツ チハヤフルカミノク",
					Description:    "綾瀬千早は高校入学と同時に、競技かるた部を作ろうと奔走する。幼馴染の太一と仲間を集め、夏の全国大会に出場するためだ。強くなって、新と再会したい。幼い頃かるたを取り合った、新に寄せる千早の秘めた想いに気づきながらも、太一は千早を守り立てる。それぞれの青春を懸けた、一途な情熱の物語が幕開ける。",
					Isbn:           "9784062938426",
					Publisher:      "",
					PublishedOn:    "2018-01-16",
					ThumbnailUrl:   "https://thumbnail.image.rakuten.co.jp/@0_mall/book/cabinet/8426/9784062938426.jpg?_ex=120x120",
					RakutenUrl:     "https://books.rakuten.co.jp/rb/15271426/",
					RakutenSize:    "コミック",
					RakutenGenreId: "001004008001/001004008003/001019001",
					Authors: []*pb.UpdateBookRequest_Author{
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
			want: false,
		},
		{
			name: "validation error: Publisher.max_len",
			args: args{
				req: &pb.UpdateBookRequest{
					Title:          "小説　ちはやふる　上の句",
					TitleKana:      "ショウセツ チハヤフルカミノク",
					Description:    "綾瀬千早は高校入学と同時に、競技かるた部を作ろうと奔走する。幼馴染の太一と仲間を集め、夏の全国大会に出場するためだ。強くなって、新と再会したい。幼い頃かるたを取り合った、新に寄せる千早の秘めた想いに気づきながらも、太一は千早を守り立てる。それぞれの青春を懸けた、一途な情熱の物語が幕開ける。",
					Isbn:           "9784062938426",
					Publisher:      strings.Repeat("x", 33),
					PublishedOn:    "2018-01-16",
					ThumbnailUrl:   "https://thumbnail.image.rakuten.co.jp/@0_mall/book/cabinet/8426/9784062938426.jpg?_ex=120x120",
					RakutenUrl:     "https://books.rakuten.co.jp/rb/15271426/",
					RakutenSize:    "コミック",
					RakutenGenreId: "001004008001/001004008003/001019001",
					Authors: []*pb.UpdateBookRequest_Author{
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
			want: false,
		},
		{
			name: "validation error: PublishedOn.min_len",
			args: args{
				req: &pb.UpdateBookRequest{
					Title:          "小説　ちはやふる　上の句",
					TitleKana:      "ショウセツ チハヤフルカミノク",
					Description:    "綾瀬千早は高校入学と同時に、競技かるた部を作ろうと奔走する。幼馴染の太一と仲間を集め、夏の全国大会に出場するためだ。強くなって、新と再会したい。幼い頃かるたを取り合った、新に寄せる千早の秘めた想いに気づきながらも、太一は千早を守り立てる。それぞれの青春を懸けた、一途な情熱の物語が幕開ける。",
					Isbn:           "9784062938426",
					Publisher:      "講談社",
					PublishedOn:    "",
					ThumbnailUrl:   "https://thumbnail.image.rakuten.co.jp/@0_mall/book/cabinet/8426/9784062938426.jpg?_ex=120x120",
					RakutenUrl:     "https://books.rakuten.co.jp/rb/15271426/",
					RakutenSize:    "コミック",
					RakutenGenreId: "001004008001/001004008003/001019001",
					Authors: []*pb.UpdateBookRequest_Author{
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
			want: false,
		},
		{
			name: "validation error: PublishedOn.max_len",
			args: args{
				req: &pb.UpdateBookRequest{
					Title:          "小説　ちはやふる　上の句",
					TitleKana:      "ショウセツ チハヤフルカミノク",
					Description:    "綾瀬千早は高校入学と同時に、競技かるた部を作ろうと奔走する。幼馴染の太一と仲間を集め、夏の全国大会に出場するためだ。強くなって、新と再会したい。幼い頃かるたを取り合った、新に寄せる千早の秘めた想いに気づきながらも、太一は千早を守り立てる。それぞれの青春を懸けた、一途な情熱の物語が幕開ける。",
					Isbn:           "9784062938426",
					Publisher:      "講談社",
					PublishedOn:    strings.Repeat("x", 17),
					ThumbnailUrl:   "https://thumbnail.image.rakuten.co.jp/@0_mall/book/cabinet/8426/9784062938426.jpg?_ex=120x120",
					RakutenUrl:     "https://books.rakuten.co.jp/rb/15271426/",
					RakutenSize:    "コミック",
					RakutenGenreId: "001004008001/001004008003/001019001",
					Authors: []*pb.UpdateBookRequest_Author{
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
			want: false,
		},
		{
			name: "validation error: Authors.Name.min_len",
			args: args{
				req: &pb.UpdateBookRequest{
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
					Authors: []*pb.UpdateBookRequest_Author{
						{
							Name:     "",
							NameKana: "アリサワ ユウキ",
						},
						{
							Name:     "末次 由紀",
							NameKana: "スエツグ ユキ",
						},
					},
				},
			},
			want: false,
		},
		{
			name: "validation error: Authors.Name.max_len",
			args: args{
				req: &pb.UpdateBookRequest{
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
					Authors: []*pb.UpdateBookRequest_Author{
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
			},
			want: false,
		},
		{
			name: "validation error: Authors.NameKana.min_len",
			args: args{
				req: &pb.UpdateBookRequest{
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
					Authors: []*pb.UpdateBookRequest_Author{
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
			},
			want: false,
		},
		{
			name: "validation error: Authors.NameKana.max_len",
			args: args{
				req: &pb.UpdateBookRequest{
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
					Authors: []*pb.UpdateBookRequest_Author{
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
			},
			want: false,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			target := NewBookRequestValidation()

			got := target.UpdateBook(tt.args.req)
			switch tt.want {
			case true:
				assert.NoError(t, got)
			case false:
				assert.Error(t, got)
			}
		})
	}
}

func TestBookRequestValidation_DeleteBook(t *testing.T) {
	t.Parallel()
	type args struct {
		req *pb.DeleteBookRequest
	}

	testCases := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "success",
			args: args{
				req: &pb.DeleteBookRequest{
					BookId: 1,
				},
			},
			want: true,
		},
		{
			name: "validation error: BookIds.gt",
			args: args{
				req: &pb.DeleteBookRequest{
					BookId: 0,
				},
			},
			want: false,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			target := NewBookRequestValidation()

			got := target.DeleteBook(tt.args.req)
			switch tt.want {
			case true:
				assert.NoError(t, got)
			case false:
				assert.Error(t, got)
			}
		})
	}
}

func TestBookRequestValidation_DeleteBookshelf(t *testing.T) {
	t.Parallel()
	type args struct {
		req *pb.DeleteBookshelfRequest
	}

	testCases := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "success",
			args: args{
				req: &pb.DeleteBookshelfRequest{
					UserId: "12345678-1234-1234-1234-123456789012",
					BookId: 1,
				},
			},
			want: true,
		},
		{
			name: "validation error: UserId.min_len",
			args: args{
				req: &pb.DeleteBookshelfRequest{
					UserId: "",
					BookId: 1,
				},
			},
			want: false,
		},
		{
			name: "validation error: BookId.gt",
			args: args{
				req: &pb.DeleteBookshelfRequest{
					UserId: "12345678-1234-1234-1234-123456789012",
					BookId: 0,
				},
			},
			want: false,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			target := NewBookRequestValidation()

			got := target.DeleteBookshelf(tt.args.req)
			switch tt.want {
			case true:
				assert.NoError(t, got)
			case false:
				assert.Error(t, got)
			}
		})
	}
}

func TestBookRequestValidation_ReadBookshelf(t *testing.T) {
	t.Parallel()
	type args struct {
		req *pb.ReadBookshelfRequest
	}

	testCases := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "success",
			args: args{
				req: &pb.ReadBookshelfRequest{
					UserId:     "12345678-1234-1234-1234-123456789012",
					BookId:     1,
					ReadOn:     "2021-08-01",
					Impression: "感想です",
				},
			},
			want: true,
		},
		{
			name: "validation error: UserId.min_len",
			args: args{
				req: &pb.ReadBookshelfRequest{
					UserId:     "",
					BookId:     1,
					ReadOn:     "2021-08-01",
					Impression: "感想です",
				},
			},
			want: false,
		},
		{
			name: "validation error: BookId.gt",
			args: args{
				req: &pb.ReadBookshelfRequest{
					UserId:     "12345678-1234-1234-1234-123456789012",
					BookId:     0,
					ReadOn:     "2021-08-01",
					Impression: "感想です",
				},
			},
			want: false,
		},
		{
			name: "validation error: Impression.max_len",
			args: args{
				req: &pb.ReadBookshelfRequest{
					UserId:     "12345678-1234-1234-1234-123456789012",
					BookId:     1,
					ReadOn:     "2021-08-01",
					Impression: strings.Repeat("x", 1001),
				},
			},
			want: false,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			target := NewBookRequestValidation()

			got := target.ReadBookshelf(tt.args.req)
			switch tt.want {
			case true:
				assert.NoError(t, got)
			case false:
				assert.Error(t, got)
			}
		})
	}
}

func TestBookRequestValidation_ReadingBookshelf(t *testing.T) {
	t.Parallel()
	type args struct {
		req *pb.ReadingBookshelfRequest
	}

	testCases := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "success",
			args: args{
				req: &pb.ReadingBookshelfRequest{
					UserId: "12345678-1234-1234-1234-123456789012",
					BookId: 1,
				},
			},
			want: true,
		},
		{
			name: "validation error: UserId.min_len",
			args: args{
				req: &pb.ReadingBookshelfRequest{
					UserId: "",
					BookId: 1,
				},
			},
			want: false,
		},
		{
			name: "validation error: BookId.gt",
			args: args{
				req: &pb.ReadingBookshelfRequest{
					UserId: "12345678-1234-1234-1234-123456789012",
					BookId: 0,
				},
			},
			want: false,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			target := NewBookRequestValidation()

			got := target.ReadingBookshelf(tt.args.req)
			switch tt.want {
			case true:
				assert.NoError(t, got)
			case false:
				assert.Error(t, got)
			}
		})
	}
}

func TestBookRequestValidation_StackedBookshelf(t *testing.T) {
	t.Parallel()
	type args struct {
		req *pb.StackedBookshelfRequest
	}

	testCases := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "success",
			args: args{
				req: &pb.StackedBookshelfRequest{
					UserId: "12345678-1234-1234-1234-123456789012",
					BookId: 1,
				},
			},
			want: true,
		},
		{
			name: "validation error: UserId.min_len",
			args: args{
				req: &pb.StackedBookshelfRequest{
					UserId: "",
					BookId: 1,
				},
			},
			want: false,
		},
		{
			name: "validation error: BookId.gt",
			args: args{
				req: &pb.StackedBookshelfRequest{
					UserId: "12345678-1234-1234-1234-123456789012",
					BookId: 0,
				},
			},
			want: false,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			target := NewBookRequestValidation()

			got := target.StackedBookshelf(tt.args.req)
			switch tt.want {
			case true:
				assert.NoError(t, got)
			case false:
				assert.Error(t, got)
			}
		})
	}
}

func TestBookRequestValidation_WantBookshelf(t *testing.T) {
	t.Parallel()
	type args struct {
		req *pb.WantBookshelfRequest
	}

	testCases := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "success",
			args: args{
				req: &pb.WantBookshelfRequest{
					UserId: "12345678-1234-1234-1234-123456789012",
					BookId: 1,
				},
			},
			want: true,
		},
		{
			name: "validation error: UserId.min_len",
			args: args{
				req: &pb.WantBookshelfRequest{
					UserId: "",
					BookId: 1,
				},
			},
			want: false,
		},
		{
			name: "validation error: BookId.gt",
			args: args{
				req: &pb.WantBookshelfRequest{
					UserId: "12345678-1234-1234-1234-123456789012",
					BookId: 0,
				},
			},
			want: false,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			target := NewBookRequestValidation()

			got := target.WantBookshelf(tt.args.req)
			switch tt.want {
			case true:
				assert.NoError(t, got)
			case false:
				assert.Error(t, got)
			}
		})
	}
}

func TestBookRequestValidation_ReleaseBookshelf(t *testing.T) {
	t.Parallel()
	type args struct {
		req *pb.ReleaseBookshelfRequest
	}

	testCases := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "success",
			args: args{
				req: &pb.ReleaseBookshelfRequest{
					UserId: "12345678-1234-1234-1234-123456789012",
					BookId: 1,
				},
			},
			want: true,
		},
		{
			name: "validation error: UserId.min_len",
			args: args{
				req: &pb.ReleaseBookshelfRequest{
					UserId: "",
					BookId: 1,
				},
			},
			want: false,
		},
		{
			name: "validation error: BookId.gt",
			args: args{
				req: &pb.ReleaseBookshelfRequest{
					UserId: "12345678-1234-1234-1234-123456789012",
					BookId: 0,
				},
			},
			want: false,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			target := NewBookRequestValidation()

			got := target.ReleaseBookshelf(tt.args.req)
			switch tt.want {
			case true:
				assert.NoError(t, got)
			case false:
				assert.Error(t, got)
			}
		})
	}
}
