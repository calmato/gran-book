package v2

import (
	"context"
	"net/http"
	"strconv"
	"testing"

	"github.com/calmato/gran-book/api/gateway/native/internal/entity"
	response "github.com/calmato/gran-book/api/gateway/native/internal/response/v2"
	"github.com/calmato/gran-book/api/gateway/native/pkg/test"
	"github.com/calmato/gran-book/api/gateway/native/proto/service/book"
	"github.com/calmato/gran-book/api/gateway/native/proto/service/user"
	"github.com/golang/mock/gomock"
)

func TestBookshelf_ListBookshelf(t *testing.T) {
	t.Parallel()

	user1 := testUser("00000000-0000-0000-0000-000000000000")
	books := make([]*book.Book, 2)
	books[0] = testBook(1)
	books[1] = testBook(2)
	bookshelves := make([]*book.Bookshelf, 2)
	bookshelves[0] = testBookshelf(1, books[0].Id, user1.Id)
	bookshelves[1] = testBookshelf(2, books[1].Id, user1.Id)

	tests := []struct {
		name   string
		setup  func(ctx context.Context, t *testing.T, mocks *test.Mocks, ctrl *gomock.Controller)
		query  string
		expect *test.TestResponse
	}{
		{
			name: "success",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks, ctrl *gomock.Controller) {
				limit, _ := strconv.ParseInt(entity.ListLimitDefault, 10, 64)
				offset, _ := strconv.ParseInt(entity.ListOffsetDefault, 10, 64)
				mocks.BookService.EXPECT().
					ListBookshelf(gomock.Any(), &book.ListBookshelfRequest{
						UserId: "00000000-0000-0000-0000-000000000000",
						Limit:  limit,
						Offset: offset,
					}).
					Return(&book.BookshelfListResponse{
						Bookshelves: bookshelves,
						Limit:       limit,
						Offset:      offset,
						Total:       2,
					}, nil)
				mocks.BookService.EXPECT().
					MultiGetBooks(gomock.Any(), &book.MultiGetBooksRequest{BookIds: []int64{1, 2}}).
					Return(&book.BookListResponse{Books: books}, nil)
			},
			query: "",
			expect: &test.TestResponse{
				Code: http.StatusOK,
				Body: response.NewBookshelfListResponse(
					entity.NewBookshelves(bookshelves),
					entity.NewBooks(books).Map(),
					100,
					0,
					2,
				),
			},
		},
		{
			name:  "failed to invalid limit query",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks, ctrl *gomock.Controller) {},
			query: "?limit=1.1",
			expect: &test.TestResponse{
				Code: http.StatusBadRequest,
			},
		},
		{
			name:  "failed to invalid offset query",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks, ctrl *gomock.Controller) {},
			query: "?offset=0.1",
			expect: &test.TestResponse{
				Code: http.StatusBadRequest,
			},
		},
		{
			name: "failed to list bookshelf",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks, ctrl *gomock.Controller) {
				limit, _ := strconv.ParseInt(entity.ListLimitDefault, 10, 64)
				offset, _ := strconv.ParseInt(entity.ListOffsetDefault, 10, 64)
				mocks.BookService.EXPECT().
					ListBookshelf(gomock.Any(), &book.ListBookshelfRequest{
						UserId: "00000000-0000-0000-0000-000000000000",
						Limit:  limit,
						Offset: offset,
					}).
					Return(nil, test.ErrMock)
			},
			query: "",
			expect: &test.TestResponse{
				Code: http.StatusInternalServerError,
			},
		},
		{
			name: "failed to multi get books",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks, ctrl *gomock.Controller) {
				limit, _ := strconv.ParseInt(entity.ListLimitDefault, 10, 64)
				offset, _ := strconv.ParseInt(entity.ListOffsetDefault, 10, 64)
				mocks.BookService.EXPECT().
					ListBookshelf(gomock.Any(), &book.ListBookshelfRequest{
						UserId: "00000000-0000-0000-0000-000000000000",
						Limit:  limit,
						Offset: offset,
					}).
					Return(&book.BookshelfListResponse{
						Bookshelves: bookshelves,
						Limit:       limit,
						Offset:      offset,
						Total:       2,
					}, nil)
				mocks.BookService.EXPECT().
					MultiGetBooks(gomock.Any(), &book.MultiGetBooksRequest{BookIds: []int64{1, 2}}).
					Return(nil, test.ErrMock)
			},
			query: "",
			expect: &test.TestResponse{
				Code: http.StatusInternalServerError,
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			path := "/v2/users/00000000-0000-0000-0000-000000000000/books" + tt.query
			req := test.NewHTTPRequest(t, http.MethodGet, path, nil)
			testHTTP(t, tt.setup, tt.expect, req)
		})
	}
}

func TestBookshelf_GetBookshelf(t *testing.T) {
	t.Parallel()

	users := make([]*user.User, 2)
	users[0] = testUser("00000000-0000-0000-0000-000000000000")
	users[1] = testUser("11111111-1111-1111-1111-111111111111")
	book1 := testBook(1)
	bookshelf1 := testBookshelf(1, book1.Id, users[0].Id)
	reviews := make([]*book.Review, 2)
	reviews[0] = testReview(1, book1.Id, users[0].Id)
	reviews[1] = testReview(2, book1.Id, users[1].Id)

	tests := []struct {
		name   string
		setup  func(ctx context.Context, t *testing.T, mocks *test.Mocks, ctrl *gomock.Controller)
		bookID string
		expect *test.TestResponse
	}{
		{
			name: "success",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks, ctrl *gomock.Controller) {
				mocks.BookService.EXPECT().
					GetBook(gomock.Any(), &book.GetBookRequest{BookId: 1}).
					Return(&book.BookResponse{Book: book1}, nil)
				mocks.BookService.EXPECT().
					GetBookshelf(gomock.Any(), &book.GetBookshelfRequest{
						UserId: "00000000-0000-0000-0000-000000000000",
						BookId: 1,
					}).
					Return(&book.BookshelfResponse{Bookshelf: bookshelf1}, nil)
				mocks.BookService.EXPECT().
					ListBookReview(gomock.Any(), &book.ListBookReviewRequest{BookId: 1, Limit: 20, Offset: 0}).
					Return(&book.ReviewListResponse{Reviews: reviews, Total: 2, Limit: 20, Offset: 0}, nil)
				mocks.UserService.EXPECT().
					MultiGetUser(gomock.Any(), &user.MultiGetUserRequest{
						UserIds: []string{
							"00000000-0000-0000-0000-000000000000",
							"11111111-1111-1111-1111-111111111111",
						},
					}).
					Return(&user.UserListResponse{Users: users, Total: 2, Limit: 2, Offset: 0}, nil)
			},
			bookID: "1",
			expect: &test.TestResponse{
				Code: http.StatusOK,
				Body: response.NewBookshelfResponse(
					entity.NewBookshelf(bookshelf1),
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
			name:   "failed to invalid book param",
			setup:  func(ctx context.Context, t *testing.T, mocks *test.Mocks, ctrl *gomock.Controller) {},
			bookID: "1.1",
			expect: &test.TestResponse{
				Code: http.StatusBadRequest,
			},
		},
		{
			name: "failed to get book",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks, ctrl *gomock.Controller) {
				mocks.BookService.EXPECT().
					GetBook(gomock.Any(), &book.GetBookRequest{BookId: 1}).
					Return(nil, test.ErrMock)
				mocks.BookService.EXPECT().
					GetBookshelf(gomock.Any(), &book.GetBookshelfRequest{
						UserId: "00000000-0000-0000-0000-000000000000",
						BookId: 1,
					}).
					Return(&book.BookshelfResponse{Bookshelf: bookshelf1}, nil)
				mocks.BookService.EXPECT().
					ListBookReview(gomock.Any(), &book.ListBookReviewRequest{BookId: 1, Limit: 20, Offset: 0}).
					Return(&book.ReviewListResponse{Reviews: reviews, Total: 2, Limit: 20, Offset: 0}, nil)
			},
			bookID: "1",
			expect: &test.TestResponse{
				Code: http.StatusInternalServerError,
			},
		},
		{
			name: "failed to get bookshelf",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks, ctrl *gomock.Controller) {
				mocks.BookService.EXPECT().
					GetBook(gomock.Any(), &book.GetBookRequest{BookId: 1}).
					Return(&book.BookResponse{Book: book1}, nil)
				mocks.BookService.EXPECT().
					GetBookshelf(gomock.Any(), &book.GetBookshelfRequest{
						UserId: "00000000-0000-0000-0000-000000000000",
						BookId: 1,
					}).
					Return(nil, test.ErrMock)
				mocks.BookService.EXPECT().
					ListBookReview(gomock.Any(), &book.ListBookReviewRequest{BookId: 1, Limit: 20, Offset: 0}).
					Return(&book.ReviewListResponse{Reviews: reviews, Total: 2, Limit: 20, Offset: 0}, nil)
			},
			bookID: "1",
			expect: &test.TestResponse{
				Code: http.StatusInternalServerError,
			},
		},
		{
			name: "failed to list book review",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks, ctrl *gomock.Controller) {
				mocks.BookService.EXPECT().
					GetBook(gomock.Any(), &book.GetBookRequest{BookId: 1}).
					Return(&book.BookResponse{Book: book1}, nil)
				mocks.BookService.EXPECT().
					GetBookshelf(gomock.Any(), &book.GetBookshelfRequest{
						UserId: "00000000-0000-0000-0000-000000000000",
						BookId: 1,
					}).
					Return(&book.BookshelfResponse{Bookshelf: bookshelf1}, nil)
				mocks.BookService.EXPECT().
					ListBookReview(gomock.Any(), &book.ListBookReviewRequest{BookId: 1, Limit: 20, Offset: 0}).
					Return(nil, test.ErrMock)
			},
			bookID: "1",
			expect: &test.TestResponse{
				Code: http.StatusInternalServerError,
			},
		},
		{
			name: "failed to multi get user",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks, ctrl *gomock.Controller) {
				mocks.BookService.EXPECT().
					GetBook(gomock.Any(), &book.GetBookRequest{BookId: 1}).
					Return(&book.BookResponse{Book: book1}, nil)
				mocks.BookService.EXPECT().
					GetBookshelf(gomock.Any(), &book.GetBookshelfRequest{
						UserId: "00000000-0000-0000-0000-000000000000",
						BookId: 1,
					}).
					Return(&book.BookshelfResponse{Bookshelf: bookshelf1}, nil)
				mocks.BookService.EXPECT().
					ListBookReview(gomock.Any(), &book.ListBookReviewRequest{BookId: 1, Limit: 20, Offset: 0}).
					Return(&book.ReviewListResponse{Reviews: reviews, Total: 2, Limit: 20, Offset: 0}, nil)
				mocks.UserService.EXPECT().
					MultiGetUser(gomock.Any(), &user.MultiGetUserRequest{
						UserIds: []string{
							"00000000-0000-0000-0000-000000000000",
							"11111111-1111-1111-1111-111111111111",
						},
					}).
					Return(nil, test.ErrMock)
			},
			bookID: "1",
			expect: &test.TestResponse{
				Code: http.StatusInternalServerError,
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			path := "/v2/users/00000000-0000-0000-0000-000000000000/books/" + tt.bookID
			req := test.NewHTTPRequest(t, http.MethodGet, path, nil)
			testHTTP(t, tt.setup, tt.expect, req)
		})
	}
}

func testBookshelf(id int64, bookID int64, userID string) *book.Bookshelf {
	return &book.Bookshelf{
		Id:        id,
		BookId:    bookID,
		UserId:    userID,
		ReviewId:  0,
		Status:    book.BookshelfStatus_BOOKSHELF_STATUS_READING,
		CreatedAt: test.TimeMock,
		UpdatedAt: test.TimeMock,
	}
}
