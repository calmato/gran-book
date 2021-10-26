package v1

import (
	"context"
	"net/http"
	"testing"

	gentity "github.com/calmato/gran-book/api/internal/gateway/entity"
	request "github.com/calmato/gran-book/api/internal/gateway/native/request/v1"
	response "github.com/calmato/gran-book/api/internal/gateway/native/response/v1"
	"github.com/calmato/gran-book/api/pkg/datetime"
	"github.com/calmato/gran-book/api/pkg/test"
	"github.com/calmato/gran-book/api/proto/book"
	"github.com/calmato/gran-book/api/proto/user"
	"github.com/golang/mock/gomock"
)

func TestBook_GetBook(t *testing.T) {
	t.Parallel()

	user1 := testAuth("00000000-0000-0000-0000-000000000000")
	book1 := testBook(1)
	bookshelf1 := testBookshelf(1, book1.Id, user1.Id)

	tests := []struct {
		name   string
		setup  func(ctx context.Context, t *testing.T, mocks *test.Mocks, ctrl *gomock.Controller)
		expect *test.HTTPResponse
	}{
		{
			name: "success",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks, ctrl *gomock.Controller) {
				mocks.AuthService.EXPECT().
					GetAuth(gomock.Any(), &user.Empty{}).
					Return(&user.AuthResponse{Auth: user1}, nil)
				mocks.BookService.EXPECT().
					GetBookByIsbn(gomock.Any(), &book.GetBookByIsbnRequest{Isbn: "9784062938426"}).
					Return(&book.BookResponse{Book: book1}, nil)
				mocks.BookService.EXPECT().
					GetBookshelf(gomock.Any(), &book.GetBookshelfRequest{UserId: user1.Id, BookId: book1.Id}).
					Return(&book.BookshelfResponse{Bookshelf: bookshelf1}, nil)
			},
			expect: &test.HTTPResponse{
				Code: http.StatusOK,
				Body: response.NewBookResponse(
					gentity.NewBook(book1),
					gentity.NewBookshelf(bookshelf1),
				),
			},
		},
		{
			name: "failed to get auth",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks, ctrl *gomock.Controller) {
				mocks.AuthService.EXPECT().
					GetAuth(gomock.Any(), &user.Empty{}).
					Return(nil, test.ErrMock)
				mocks.BookService.EXPECT().
					GetBookByIsbn(gomock.Any(), &book.GetBookByIsbnRequest{Isbn: "9784062938426"}).
					Return(&book.BookResponse{Book: book1}, nil)
			},
			expect: &test.HTTPResponse{
				Code: http.StatusInternalServerError,
			},
		},
		{
			name: "failed to get book by isbn",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks, ctrl *gomock.Controller) {
				mocks.AuthService.EXPECT().
					GetAuth(gomock.Any(), &user.Empty{}).
					Return(&user.AuthResponse{Auth: user1}, nil)
				mocks.BookService.EXPECT().
					GetBookByIsbn(gomock.Any(), &book.GetBookByIsbnRequest{Isbn: "9784062938426"}).
					Return(nil, test.ErrMock)
			},
			expect: &test.HTTPResponse{
				Code: http.StatusInternalServerError,
			},
		},
		{
			name: "failed to get bookshelf",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks, ctrl *gomock.Controller) {
				mocks.AuthService.EXPECT().
					GetAuth(gomock.Any(), &user.Empty{}).
					Return(&user.AuthResponse{Auth: user1}, nil)
				mocks.BookService.EXPECT().
					GetBookByIsbn(gomock.Any(), &book.GetBookByIsbnRequest{Isbn: "9784062938426"}).
					Return(&book.BookResponse{Book: book1}, nil)
				mocks.BookService.EXPECT().
					GetBookshelf(gomock.Any(), &book.GetBookshelfRequest{UserId: user1.Id, BookId: book1.Id}).
					Return(nil, test.ErrMock)
			},
			expect: &test.HTTPResponse{
				Code: http.StatusInternalServerError,
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			path := "/v1/books/9784062938426"
			req := test.NewHTTPRequest(t, http.MethodGet, path, nil)
			testHTTP(t, tt.setup, tt.expect, req)
		})
	}
}

func TestBook_CreateBook(t *testing.T) {
	t.Parallel()

	now := datetime.FormatTime(test.TimeMock)
	book1 := testBook(1)

	tests := []struct {
		name   string
		setup  func(ctx context.Context, t *testing.T, mocks *test.Mocks, ctrl *gomock.Controller)
		req    *request.CreateBookRequest
		expect *test.HTTPResponse
	}{
		{
			name: "success",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks, ctrl *gomock.Controller) {
				mocks.BookService.EXPECT().
					CreateBook(gomock.Any(), &book.CreateBookRequest{
						Title:          "小説　ちはやふる　上の句",
						TitleKana:      "ショウセツ チハヤフルカミノク",
						Description:    "綾瀬千早は高校入学と同時に、競技かるた部を作ろうと奔走する。幼馴染の太一と仲間を集め、夏の全国大会に出場するためだ。強くなって、新と再会したい。幼い頃かるたを取り合った、に寄早の秘めた想いに気づきながらも、太一は千早を守り立てる。それぞれの青春を懸けた、一途な情熱の物語が幕開ける。",
						Isbn:           "9784062938426",
						Publisher:      "講談社",
						PublishedOn:    "2018-01-16",
						ThumbnailUrl:   "https://thumbnail.image.rakuten.co.jp/@0_mall/book/cabinet/8426/9784062938426.jpg?_ex=120x120",
						RakutenUrl:     "https://books.rakuten.co.jp/rb/15271426/",
						RakutenSize:    "コミック",
						RakutenGenreId: "001004008001/001004008003/001019001",
						Authors: []*book.CreateBookRequest_Author{
							{
								Name:     "有沢 ゆう希",
								NameKana: "アリサワ ユウキ",
							},
							{
								Name:     "末次 由紀",
								NameKana: "スエツグ ユキ",
							},
						},
					}).
					Return(&book.BookResponse{Book: book1}, nil)
			},
			req: &request.CreateBookRequest{
				Title:          "小説　ちはやふる　上の句",
				TitleKana:      "ショウセツ チハヤフルカミノク",
				ItemCaption:    "綾瀬千早は高校入学と同時に、競技かるた部を作ろうと奔走する。幼馴染の太一と仲間を集め、夏の全国大会に出場するためだ。強くなって、新と再会したい。幼い頃かるたを取り合った、に寄早の秘めた想いに気づきながらも、太一は千早を守り立てる。それぞれの青春を懸けた、一途な情熱の物語が幕開ける。",
				Isbn:           "9784062938426",
				PublisherName:  "講談社",
				SalesDate:      "2018-01-16",
				SmallImageURL:  "https://thumbnail.image.rakuten.co.jp/@0_mall/book/cabinet/8426/9784062938426.jpg?_ex=120x120",
				MediumImageURL: "https://thumbnail.image.rakuten.co.jp/@0_mall/book/cabinet/8426/9784062938426.jpg?_ex=120x120",
				LargeImageURL:  "https://thumbnail.image.rakuten.co.jp/@0_mall/book/cabinet/8426/9784062938426.jpg?_ex=120x120",
				ItemURL:        "https://books.rakuten.co.jp/rb/15271426/",
				Size:           "コミック",
				BooksGenreID:   "001004008001/001004008003/001019001",
				Author:         "有沢 ゆう希/末次 由紀",
				AuthorKana:     "アリサワ ユウキ/スエツグ ユキ",
			},
			expect: &test.HTTPResponse{
				Code: http.StatusOK,
				Body: &response.BookResponse{
					ID:           1,
					Title:        "小説　ちはやふる　上の句",
					TitleKana:    "ショウセツ チハヤフルカミノク",
					Description:  "綾瀬千早は高校入学と同時に、競技かるた部を作ろうと奔走する。幼馴染の太一と仲間を集め、夏の全国大会に出場するためだ。強くなって、新と再会したい。幼い頃かるたを取り合った、に寄早の秘めた想いに気づきながらも、太一は千早を守り立てる。それぞれの青春を懸けた、一途な情熱の物語が幕開ける。",
					Isbn:         "9784062938426",
					Publisher:    "講談社",
					PublishedOn:  "2018-01-16",
					ThumbnailURL: "https://thumbnail.image.rakuten.co.jp/@0_mall/book/cabinet/8426/9784062938426.jpg?_ex=120x120",
					RakutenURL:   "https://books.rakuten.co.jp/rb/15271426/",
					Size:         "コミック",
					Author:       "有沢 ゆう希/末次 由紀",
					AuthorKana:   "アリサワ ユウキ/スエツグ ユキ",
					Bookshelf:    nil,
					CreatedAt:    now,
					UpdatedAt:    now,
				},
			},
		},
		{
			name:  "failed to invalid argument",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks, ctrl *gomock.Controller) {},
			req:   &request.CreateBookRequest{},
			expect: &test.HTTPResponse{
				Code: http.StatusBadRequest,
			},
		},
		{
			name: "failed to create book",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks, ctrl *gomock.Controller) {
				mocks.BookService.EXPECT().
					CreateBook(gomock.Any(), &book.CreateBookRequest{
						Title:          "小説　ちはやふる　上の句",
						TitleKana:      "ショウセツ チハヤフルカミノク",
						Description:    "綾瀬千早は高校入学と同時に、競技かるた部を作ろうと奔走する。幼馴染の太一と仲間を集め、夏の全国大会に出場するためだ。強くなって、新と再会したい。幼い頃かるたを取り合った、に寄早の秘めた想いに気づきながらも、太一は千早を守り立てる。それぞれの青春を懸けた、一途な情熱の物語が幕開ける。",
						Isbn:           "9784062938426",
						Publisher:      "講談社",
						PublishedOn:    "2018-01-16",
						ThumbnailUrl:   "https://thumbnail.image.rakuten.co.jp/@0_mall/book/cabinet/8426/9784062938426.jpg?_ex=120x120",
						RakutenUrl:     "https://books.rakuten.co.jp/rb/15271426/",
						RakutenSize:    "コミック",
						RakutenGenreId: "001004008001/001004008003/001019001",
						Authors: []*book.CreateBookRequest_Author{
							{
								Name:     "有沢 ゆう希",
								NameKana: "アリサワ ユウキ",
							},
							{
								Name:     "末次 由紀",
								NameKana: "スエツグ ユキ",
							},
						},
					}).
					Return(nil, test.ErrMock)
			},
			req: &request.CreateBookRequest{
				Title:          "小説　ちはやふる　上の句",
				TitleKana:      "ショウセツ チハヤフルカミノク",
				ItemCaption:    "綾瀬千早は高校入学と同時に、競技かるた部を作ろうと奔走する。幼馴染の太一と仲間を集め、夏の全国大会に出場するためだ。強くなって、新と再会したい。幼い頃かるたを取り合った、に寄早の秘めた想いに気づきながらも、太一は千早を守り立てる。それぞれの青春を懸けた、一途な情熱の物語が幕開ける。",
				Isbn:           "9784062938426",
				PublisherName:  "講談社",
				SalesDate:      "2018-01-16",
				SmallImageURL:  "https://thumbnail.image.rakuten.co.jp/@0_mall/book/cabinet/8426/9784062938426.jpg?_ex=120x120",
				MediumImageURL: "",
				LargeImageURL:  "",
				ItemURL:        "https://books.rakuten.co.jp/rb/15271426/",
				Size:           "コミック",
				BooksGenreID:   "001004008001/001004008003/001019001",
				Author:         "有沢 ゆう希/末次 由紀",
				AuthorKana:     "アリサワ ユウキ/スエツグ ユキ",
			},
			expect: &test.HTTPResponse{
				Code: http.StatusInternalServerError,
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			path := "/v1/books"
			req := test.NewHTTPRequest(t, http.MethodPost, path, tt.req)
			testHTTP(t, tt.setup, tt.expect, req)
		})
	}
}

func TestBook_UpdateBook(t *testing.T) {
	t.Parallel()

	now := datetime.FormatTime(test.TimeMock)
	book1 := testBook(1)

	tests := []struct {
		name   string
		setup  func(ctx context.Context, t *testing.T, mocks *test.Mocks, ctrl *gomock.Controller)
		req    *request.UpdateBookRequest
		expect *test.HTTPResponse
	}{
		{
			name: "success",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks, ctrl *gomock.Controller) {
				mocks.BookService.EXPECT().
					UpdateBook(gomock.Any(), &book.UpdateBookRequest{
						Title:          "小説　ちはやふる　上の句",
						TitleKana:      "ショウセツ チハヤフルカミノク",
						Description:    "綾瀬千早は高校入学と同時に、競技かるた部を作ろうと奔走する。幼馴染の太一と仲間を集め、夏の全国大会に出場するためだ。強くなって、新と再会したい。幼い頃かるたを取り合った、に寄早の秘めた想いに気づきながらも、太一は千早を守り立てる。それぞれの青春を懸けた、一途な情熱の物語が幕開ける。",
						Isbn:           "9784062938426",
						Publisher:      "講談社",
						PublishedOn:    "2018-01-16",
						ThumbnailUrl:   "https://thumbnail.image.rakuten.co.jp/@0_mall/book/cabinet/8426/9784062938426.jpg?_ex=120x120",
						RakutenUrl:     "https://books.rakuten.co.jp/rb/15271426/",
						RakutenSize:    "コミック",
						RakutenGenreId: "001004008001/001004008003/001019001",
						Authors: []*book.UpdateBookRequest_Author{
							{
								Name:     "有沢 ゆう希",
								NameKana: "アリサワ ユウキ",
							},
							{
								Name:     "末次 由紀",
								NameKana: "スエツグ ユキ",
							},
						},
					}).
					Return(&book.BookResponse{Book: book1}, nil)
			},
			req: &request.UpdateBookRequest{
				Title:          "小説　ちはやふる　上の句",
				TitleKana:      "ショウセツ チハヤフルカミノク",
				ItemCaption:    "綾瀬千早は高校入学と同時に、競技かるた部を作ろうと奔走する。幼馴染の太一と仲間を集め、夏の全国大会に出場するためだ。強くなって、新と再会したい。幼い頃かるたを取り合った、に寄早の秘めた想いに気づきながらも、太一は千早を守り立てる。それぞれの青春を懸けた、一途な情熱の物語が幕開ける。",
				Isbn:           "9784062938426",
				PublisherName:  "講談社",
				SalesDate:      "2018-01-16",
				SmallImageURL:  "https://thumbnail.image.rakuten.co.jp/@0_mall/book/cabinet/8426/9784062938426.jpg?_ex=120x120",
				MediumImageURL: "https://thumbnail.image.rakuten.co.jp/@0_mall/book/cabinet/8426/9784062938426.jpg?_ex=120x120",
				LargeImageURL:  "",
				ItemURL:        "https://books.rakuten.co.jp/rb/15271426/",
				Size:           "コミック",
				BooksGenreID:   "001004008001/001004008003/001019001",
				Author:         "有沢 ゆう希/末次 由紀",
				AuthorKana:     "アリサワ ユウキ/スエツグ ユキ",
			},
			expect: &test.HTTPResponse{
				Code: http.StatusOK,
				Body: &response.BookResponse{
					ID:           1,
					Title:        "小説　ちはやふる　上の句",
					TitleKana:    "ショウセツ チハヤフルカミノク",
					Description:  "綾瀬千早は高校入学と同時に、競技かるた部を作ろうと奔走する。幼馴染の太一と仲間を集め、夏の全国大会に出場するためだ。強くなって、新と再会したい。幼い頃かるたを取り合った、に寄早の秘めた想いに気づきながらも、太一は千早を守り立てる。それぞれの青春を懸けた、一途な情熱の物語が幕開ける。",
					Isbn:         "9784062938426",
					Publisher:    "講談社",
					PublishedOn:  "2018-01-16",
					ThumbnailURL: "https://thumbnail.image.rakuten.co.jp/@0_mall/book/cabinet/8426/9784062938426.jpg?_ex=120x120",
					RakutenURL:   "https://books.rakuten.co.jp/rb/15271426/",
					Size:         "コミック",
					Author:       "有沢 ゆう希/末次 由紀",
					AuthorKana:   "アリサワ ユウキ/スエツグ ユキ",
					Bookshelf:    nil,
					CreatedAt:    now,
					UpdatedAt:    now,
				},
			},
		},
		{
			name:  "failed to invalid argument",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks, ctrl *gomock.Controller) {},
			req:   &request.UpdateBookRequest{},
			expect: &test.HTTPResponse{
				Code: http.StatusBadRequest,
			},
		},
		{
			name: "failed to create book",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks, ctrl *gomock.Controller) {
				mocks.BookService.EXPECT().
					UpdateBook(gomock.Any(), &book.UpdateBookRequest{
						Title:          "小説　ちはやふる　上の句",
						TitleKana:      "ショウセツ チハヤフルカミノク",
						Description:    "綾瀬千早は高校入学と同時に、競技かるた部を作ろうと奔走する。幼馴染の太一と仲間を集め、夏の全国大会に出場するためだ。強くなって、新と再会したい。幼い頃かるたを取り合った、に寄早の秘めた想いに気づきながらも、太一は千早を守り立てる。それぞれの青春を懸けた、一途な情熱の物語が幕開ける。",
						Isbn:           "9784062938426",
						Publisher:      "講談社",
						PublishedOn:    "2018-01-16",
						ThumbnailUrl:   "",
						RakutenUrl:     "https://books.rakuten.co.jp/rb/15271426/",
						RakutenSize:    "コミック",
						RakutenGenreId: "001004008001/001004008003/001019001",
						Authors: []*book.UpdateBookRequest_Author{
							{
								Name:     "有沢 ゆう希",
								NameKana: "アリサワ ユウキ",
							},
							{
								Name:     "末次 由紀",
								NameKana: "スエツグ ユキ",
							},
						},
					}).
					Return(nil, test.ErrMock)
			},
			req: &request.UpdateBookRequest{
				Title:          "小説　ちはやふる　上の句",
				TitleKana:      "ショウセツ チハヤフルカミノク",
				ItemCaption:    "綾瀬千早は高校入学と同時に、競技かるた部を作ろうと奔走する。幼馴染の太一と仲間を集め、夏の全国大会に出場するためだ。強くなって、新と再会したい。幼い頃かるたを取り合った、に寄早の秘めた想いに気づきながらも、太一は千早を守り立てる。それぞれの青春を懸けた、一途な情熱の物語が幕開ける。",
				Isbn:           "9784062938426",
				PublisherName:  "講談社",
				SalesDate:      "2018-01-16",
				SmallImageURL:  "",
				MediumImageURL: "",
				LargeImageURL:  "",
				ItemURL:        "https://books.rakuten.co.jp/rb/15271426/",
				Size:           "コミック",
				BooksGenreID:   "001004008001/001004008003/001019001",
				Author:         "有沢 ゆう希/末次 由紀",
				AuthorKana:     "アリサワ ユウキ/スエツグ ユキ",
			},
			expect: &test.HTTPResponse{
				Code: http.StatusInternalServerError,
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			path := "/v1/books"
			req := test.NewHTTPRequest(t, http.MethodPatch, path, tt.req)
			testHTTP(t, tt.setup, tt.expect, req)
		})
	}
}

func testBook(id int64) *book.Book {
	now := datetime.FormatTime(test.TimeMock)
	return &book.Book{
		Id:             id,
		Title:          "小説　ちはやふる　上の句",
		TitleKana:      "ショウセツ チハヤフルカミノク",
		Description:    "綾瀬千早は高校入学と同時に、競技かるた部を作ろうと奔走する。幼馴染の太一と仲間を集め、夏の全国大会に出場するためだ。強くなって、新と再会したい。幼い頃かるたを取り合った、に寄早の秘めた想いに気づきながらも、太一は千早を守り立てる。それぞれの青春を懸けた、一途な情熱の物語が幕開ける。",
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
	}
}
