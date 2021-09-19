package v2

import (
	"context"
	"net/http"
	"testing"

	"github.com/calmato/gran-book/api/gateway/native/internal/entity"
	response "github.com/calmato/gran-book/api/gateway/native/internal/response/v2"
	"github.com/calmato/gran-book/api/gateway/native/pkg/test"
	"github.com/calmato/gran-book/api/gateway/native/proto/service/book"
	"github.com/calmato/gran-book/api/gateway/native/proto/service/user"
	"github.com/golang/mock/gomock"
)

func TestBook_GetBook(t *testing.T) {
	t.Parallel()

	book1 := testBook(1)
	users := make([]*user.User, 2)
	users[0] = testUser("00000000-0000-0000-0000-000000000000")
	users[1] = testUser("11111111-1111-1111-1111-111111111111")
	reviews := make([]*book.Review, 2)
	reviews[0] = testReview(1, book1.Id, users[0].Id)
	reviews[1] = testReview(2, book1.Id, users[1].Id)

	tests := []struct {
		name   string
		setup  func(ctx context.Context, t *testing.T, mocks *test.Mocks, ctrl *gomock.Controller)
		bookID string
		query  string
		expect *test.TestResponse
	}{
		{
			name: "id:success",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks, ctrl *gomock.Controller) {
				mocks.BookService.EXPECT().
					GetBook(gomock.Any(), &book.GetBookRequest{BookId: 1}).
					Return(&book.BookResponse{Book: book1}, nil)
				mocks.BookService.EXPECT().
					ListBookReview(gomock.Any(), &book.ListBookReviewRequest{BookId: 1, Limit: 20, Offset: 0}).
					Return(&book.ReviewListResponse{Reviews: reviews, Total: 2, Limit: 20, Offset: 0}, nil)
				mocks.UserService.EXPECT().
					MultiGetUser(gomock.Any(), &user.MultiGetUserRequest{UserIds: []string{
						"00000000-0000-0000-0000-000000000000",
						"11111111-1111-1111-1111-111111111111",
					}}).
					Return(&user.UserListResponse{Users: users, Total: 2, Limit: 2, Offset: 0}, nil)
			},
			bookID: "1",
			query:  "?key=id",
			expect: &test.TestResponse{
				Code: http.StatusOK,
				Body: response.NewBookResponse(
					entity.NewBook(book1),
					entity.NewReviews(reviews),
					entity.NewUsers(users).Map(),
					20,
					0,
					2,
				),
			},
		},
		{
			name:   "id:failed to invalid book id",
			setup:  func(ctx context.Context, t *testing.T, mocks *test.Mocks, ctrl *gomock.Controller) {},
			bookID: "1.1",
			query:  "?key=id",
			expect: &test.TestResponse{
				Code: http.StatusBadRequest,
			},
		},
		{
			name: "id:failed to get book",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks, ctrl *gomock.Controller) {
				mocks.BookService.EXPECT().
					GetBook(gomock.Any(), &book.GetBookRequest{BookId: 1}).
					Return(nil, test.ErrMock)
			},
			bookID: "1",
			query:  "?key=id",
			expect: &test.TestResponse{
				Code: http.StatusInternalServerError,
			},
		},
		{
			name: "id:failed to list book review",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks, ctrl *gomock.Controller) {
				mocks.BookService.EXPECT().
					GetBook(gomock.Any(), &book.GetBookRequest{BookId: 1}).
					Return(&book.BookResponse{Book: book1}, nil)
				mocks.BookService.EXPECT().
					ListBookReview(gomock.Any(), &book.ListBookReviewRequest{BookId: 1, Limit: 20, Offset: 0}).
					Return(nil, test.ErrMock)
			},
			bookID: "1",
			query:  "?key=id",
			expect: &test.TestResponse{
				Code: http.StatusInternalServerError,
			},
		},
		{
			name: "id:failed to multi get user",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks, ctrl *gomock.Controller) {
				mocks.BookService.EXPECT().
					GetBook(gomock.Any(), &book.GetBookRequest{BookId: 1}).
					Return(&book.BookResponse{Book: book1}, nil)
				mocks.BookService.EXPECT().
					ListBookReview(gomock.Any(), &book.ListBookReviewRequest{BookId: 1, Limit: 20, Offset: 0}).
					Return(&book.ReviewListResponse{Reviews: reviews, Total: 2, Limit: 20, Offset: 0}, nil)
				mocks.UserService.EXPECT().
					MultiGetUser(gomock.Any(), &user.MultiGetUserRequest{UserIds: []string{
						"00000000-0000-0000-0000-000000000000",
						"11111111-1111-1111-1111-111111111111",
					}}).
					Return(nil, test.ErrMock)
			},
			bookID: "1",
			query:  "?key=id",
			expect: &test.TestResponse{
				Code: http.StatusInternalServerError,
			},
		},
		{
			name: "isbn:success",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks, ctrl *gomock.Controller) {
				mocks.BookService.EXPECT().
					GetBookByIsbn(gomock.Any(), &book.GetBookByIsbnRequest{Isbn: "9784062938426"}).
					Return(&book.BookResponse{Book: book1}, nil)
				mocks.BookService.EXPECT().
					ListBookReview(gomock.Any(), &book.ListBookReviewRequest{BookId: 1, Limit: 20, Offset: 0}).
					Return(&book.ReviewListResponse{Reviews: reviews, Total: 2, Limit: 20, Offset: 0}, nil)
				mocks.UserService.EXPECT().
					MultiGetUser(gomock.Any(), &user.MultiGetUserRequest{UserIds: []string{
						"00000000-0000-0000-0000-000000000000",
						"11111111-1111-1111-1111-111111111111",
					}}).
					Return(&user.UserListResponse{Users: users, Total: 2, Limit: 2, Offset: 0}, nil)
			},
			bookID: "9784062938426",
			query:  "?key=isbn",
			expect: &test.TestResponse{
				Code: http.StatusOK,
				Body: response.NewBookResponse(
					entity.NewBook(book1),
					entity.NewReviews(reviews),
					entity.NewUsers(users).Map(),
					20,
					0,
					2,
				),
			},
		},
		{
			name: "isbn:failed to get book by isbn",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks, ctrl *gomock.Controller) {
				mocks.BookService.EXPECT().
					GetBookByIsbn(gomock.Any(), &book.GetBookByIsbnRequest{Isbn: "9784062938426"}).
					Return(nil, test.ErrMock)
			},
			bookID: "9784062938426",
			query:  "?key=isbn",
			expect: &test.TestResponse{
				Code: http.StatusInternalServerError,
			},
		},
		{
			name: "isbn:failed to list book review",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks, ctrl *gomock.Controller) {
				mocks.BookService.EXPECT().
					GetBookByIsbn(gomock.Any(), &book.GetBookByIsbnRequest{Isbn: "9784062938426"}).
					Return(&book.BookResponse{Book: book1}, nil)
				mocks.BookService.EXPECT().
					ListBookReview(gomock.Any(), &book.ListBookReviewRequest{BookId: 1, Limit: 20, Offset: 0}).
					Return(nil, test.ErrMock)
			},
			bookID: "9784062938426",
			query:  "?key=isbn",
			expect: &test.TestResponse{
				Code: http.StatusInternalServerError,
			},
		},
		{
			name: "isbn:failed to multi get user",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks, ctrl *gomock.Controller) {
				mocks.BookService.EXPECT().
					GetBookByIsbn(gomock.Any(), &book.GetBookByIsbnRequest{Isbn: "9784062938426"}).
					Return(&book.BookResponse{Book: book1}, nil)
				mocks.BookService.EXPECT().
					ListBookReview(gomock.Any(), &book.ListBookReviewRequest{BookId: 1, Limit: 20, Offset: 0}).
					Return(&book.ReviewListResponse{Reviews: reviews, Total: 2, Limit: 20, Offset: 0}, nil)
				mocks.UserService.EXPECT().
					MultiGetUser(gomock.Any(), &user.MultiGetUserRequest{UserIds: []string{
						"00000000-0000-0000-0000-000000000000",
						"11111111-1111-1111-1111-111111111111",
					}}).
					Return(nil, test.ErrMock)
			},
			bookID: "9784062938426",
			query:  "?key=isbn",
			expect: &test.TestResponse{
				Code: http.StatusInternalServerError,
			},
		},
		{
			name:   "other:failed to invalid key",
			setup:  func(ctx context.Context, t *testing.T, mocks *test.Mocks, ctrl *gomock.Controller) {},
			bookID: "1",
			query:  "?key=bookid",
			expect: &test.TestResponse{
				Code: http.StatusBadRequest,
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			path := "/v2/books/" + tt.bookID + tt.query
			req := test.NewHTTPRequest(t, http.MethodGet, path, nil)
			testHTTP(t, tt.setup, tt.expect, req)
		})
	}
}

func testBook(id int64) *book.Book {
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
	}
}

func testReview(id int64, bookID int64, userID string) *book.Review {
	return &book.Review{
		Id:         id,
		BookId:     bookID,
		UserId:     userID,
		Score:      3,
		Impression: "テストレビューです",
		CreatedAt:  test.TimeMock,
		UpdatedAt:  test.TimeMock,
	}
}

func testUser(id string) *user.User {
	return &user.User{
		Id:               id,
		Username:         "テストユーザー",
		Gender:           user.Gender_GENDER_MAN,
		Email:            "test-user@calmato.jp",
		PhoneNumber:      "000-0000-0000",
		ThumbnailUrl:     "https://go.dev/images/gophers/ladder.svg",
		SelfIntroduction: "テストコードです。",
		LastName:         "テスト",
		FirstName:        "ユーザー",
		LastNameKana:     "てすと",
		FirstNameKana:    "ゆーざー",
		CreatedAt:        test.TimeMock,
		UpdatedAt:        test.TimeMock,
	}
}
