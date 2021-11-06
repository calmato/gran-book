package handler

import (
	"context"
	"net/http"
	"testing"

	gentity "github.com/calmato/gran-book/api/internal/gateway/entity"
	"github.com/calmato/gran-book/api/internal/gateway/native/v1/entity"
	response "github.com/calmato/gran-book/api/internal/gateway/native/v1/response"
	"github.com/calmato/gran-book/api/pkg/datetime"
	"github.com/calmato/gran-book/api/pkg/test"
	"github.com/calmato/gran-book/api/proto/book"
	"github.com/calmato/gran-book/api/proto/user"
	"github.com/golang/mock/gomock"
)

func TestReview_ListReviewByBook(t *testing.T) {
	t.Parallel()

	users := make([]*user.User, 2)
	users[0] = testUser("00000000-0000-0000-0000-000000000000")
	users[1] = testUser("11111111-1111-1111-1111-111111111111")
	book1 := testBook(1)
	reviews := make([]*book.Review, 2)
	reviews[0] = testReview(1, book1.Id, users[0].Id)
	reviews[1] = testReview(2, book1.Id, users[1].Id)

	tests := []struct {
		name   string
		setup  func(ctx context.Context, t *testing.T, mocks *test.Mocks, ctrl *gomock.Controller)
		bookID string
		query  string
		expect *test.HTTPResponse
	}{
		{
			name: "success",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks, ctrl *gomock.Controller) {
				mocks.BookService.EXPECT().
					ListBookReview(gomock.Any(), &book.ListBookReviewRequest{
						BookId: 1,
						Limit:  100,
						Offset: 0,
					}).
					Return(&book.ReviewListResponse{
						Reviews: reviews,
						Limit:   100,
						Offset:  0,
						Total:   2,
					}, nil)
				mocks.UserService.EXPECT().
					MultiGetUser(gomock.Any(), &user.MultiGetUserRequest{
						UserIds: []string{
							"00000000-0000-0000-0000-000000000000",
							"11111111-1111-1111-1111-111111111111",
						},
					}).
					Return(&user.UserListResponse{
						Users:  users,
						Limit:  2,
						Offset: 0,
						Total:  2,
					}, nil)
			},
			bookID: "1",
			query:  "",
			expect: &test.HTTPResponse{
				Code: http.StatusOK,
				Body: &response.BookReviewListResponse{
					Reviews: entity.NewBookReviews(gentity.NewReviews(reviews), gentity.NewUsers(users).Map()),
					Limit:   100,
					Offset:  0,
					Total:   2,
				},
			},
		},
		{
			name:   "failed to invalid limit query",
			setup:  func(ctx context.Context, t *testing.T, mocks *test.Mocks, ctrl *gomock.Controller) {},
			bookID: "1",
			query:  "?limit=1.1",
			expect: &test.HTTPResponse{
				Code: http.StatusBadRequest,
			},
		},
		{
			name:   "failed to invalid offset query",
			setup:  func(ctx context.Context, t *testing.T, mocks *test.Mocks, ctrl *gomock.Controller) {},
			bookID: "1",
			query:  "?offset=0.1",
			expect: &test.HTTPResponse{
				Code: http.StatusBadRequest,
			},
		},
		{
			name:   "failed to invalid book id param",
			setup:  func(ctx context.Context, t *testing.T, mocks *test.Mocks, ctrl *gomock.Controller) {},
			bookID: "1.1",
			query:  "",
			expect: &test.HTTPResponse{
				Code: http.StatusBadRequest,
			},
		},
		{
			name: "failed to list book review",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks, ctrl *gomock.Controller) {
				mocks.BookService.EXPECT().
					ListBookReview(gomock.Any(), &book.ListBookReviewRequest{
						BookId: 1,
						Limit:  100,
						Offset: 0,
					}).
					Return(nil, test.ErrMock)
			},
			bookID: "1",
			query:  "",
			expect: &test.HTTPResponse{
				Code: http.StatusInternalServerError,
			},
		},
		{
			name: "failed to multi get user",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks, ctrl *gomock.Controller) {
				mocks.BookService.EXPECT().
					ListBookReview(gomock.Any(), &book.ListBookReviewRequest{
						BookId: 1,
						Limit:  100,
						Offset: 0,
					}).
					Return(&book.ReviewListResponse{
						Reviews: reviews,
						Limit:   100,
						Offset:  0,
						Total:   2,
					}, nil)
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
			query:  "",
			expect: &test.HTTPResponse{
				Code: http.StatusInternalServerError,
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			path := "/v1/books/" + tt.bookID + "/reviews" + tt.query
			req := test.NewHTTPRequest(t, http.MethodGet, path, nil)
			testHTTP(t, tt.setup, tt.expect, req)
		})
	}
}

func TestReview_ListReviewByUser(t *testing.T) {
	t.Parallel()

	user1 := testUser("00000000-0000-0000-0000-000000000000")
	books := make([]*book.Book, 2)
	books[0] = testBook(1)
	books[1] = testBook(2)
	reviews := make([]*book.Review, 2)
	reviews[0] = testReview(1, books[0].Id, user1.Id)
	reviews[1] = testReview(2, books[1].Id, user1.Id)

	tests := []struct {
		name   string
		setup  func(ctx context.Context, t *testing.T, mocks *test.Mocks, ctrl *gomock.Controller)
		query  string
		expect *test.HTTPResponse
	}{
		{
			name: "success",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks, ctrl *gomock.Controller) {
				mocks.BookService.EXPECT().
					ListUserReview(gomock.Any(), &book.ListUserReviewRequest{
						UserId: "00000000-0000-0000-0000-000000000000",
						Limit:  100,
						Offset: 0,
					}).
					Return(&book.ReviewListResponse{
						Reviews: reviews,
						Limit:   100,
						Offset:  0,
						Total:   2,
					}, nil)
				mocks.BookService.EXPECT().
					MultiGetBooks(gomock.Any(), &book.MultiGetBooksRequest{BookIds: []int64{1, 2}}).
					Return(&book.BookListResponse{Books: books}, nil)
			},
			query: "",
			expect: &test.HTTPResponse{
				Code: http.StatusOK,
				Body: &response.UserReviewListResponse{
					Reviews: entity.NewUserReviews(gentity.NewReviews(reviews), gentity.NewBooks(books).Map()),
					Limit:   100,
					Offset:  0,
					Total:   2,
				},
			},
		},
		{
			name:  "failed to invalid limit query",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks, ctrl *gomock.Controller) {},
			query: "?limit=1.1",
			expect: &test.HTTPResponse{
				Code: http.StatusBadRequest,
			},
		},
		{
			name:  "failed to invalid offset query",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks, ctrl *gomock.Controller) {},
			query: "?offset=0.1",
			expect: &test.HTTPResponse{
				Code: http.StatusBadRequest,
			},
		},
		{
			name: "failed to list user review",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks, ctrl *gomock.Controller) {
				mocks.BookService.EXPECT().
					ListUserReview(gomock.Any(), &book.ListUserReviewRequest{
						UserId: "00000000-0000-0000-0000-000000000000",
						Limit:  100,
						Offset: 0,
					}).
					Return(nil, test.ErrMock)
			},
			query: "",
			expect: &test.HTTPResponse{
				Code: http.StatusInternalServerError,
			},
		},
		{
			name: "failed to multi get books",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks, ctrl *gomock.Controller) {
				mocks.BookService.EXPECT().
					ListUserReview(gomock.Any(), &book.ListUserReviewRequest{
						UserId: "00000000-0000-0000-0000-000000000000",
						Limit:  100,
						Offset: 0,
					}).
					Return(&book.ReviewListResponse{
						Reviews: reviews,
						Limit:   100,
						Offset:  0,
						Total:   2,
					}, nil)
				mocks.BookService.EXPECT().
					MultiGetBooks(gomock.Any(), &book.MultiGetBooksRequest{BookIds: []int64{1, 2}}).
					Return(nil, test.ErrMock)
			},
			query: "",
			expect: &test.HTTPResponse{
				Code: http.StatusInternalServerError,
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			path := "/v1/users/00000000-0000-0000-0000-000000000000/reviews" + tt.query
			req := test.NewHTTPRequest(t, http.MethodGet, path, nil)
			testHTTP(t, tt.setup, tt.expect, req)
		})
	}
}

func TestReview_getBookReview(t *testing.T) {
	t.Parallel()

	user1 := testUser("00000000-0000-0000-0000-000000000000")
	book1 := testBook(1)
	review1 := testReview(1, book1.Id, user1.Id)

	tests := []struct {
		name     string
		setup    func(ctx context.Context, t *testing.T, mocks *test.Mocks, ctrl *gomock.Controller)
		bookID   string
		reviewID string
		expect   *test.HTTPResponse
	}{
		{
			name: "success",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks, ctrl *gomock.Controller) {
				mocks.BookService.EXPECT().
					GetBook(gomock.Any(), &book.GetBookRequest{BookId: 1}).
					Return(&book.BookResponse{Book: book1}, nil)
				mocks.BookService.EXPECT().
					GetReview(gomock.Any(), &book.GetReviewRequest{ReviewId: 1}).
					Return(&book.ReviewResponse{Review: review1}, nil)
				mocks.UserService.EXPECT().
					GetUser(gomock.Any(), &user.GetUserRequest{UserId: "00000000-0000-0000-0000-000000000000"}).
					Return(&user.UserResponse{User: user1}, nil)
			},
			bookID:   "1",
			reviewID: "1",
			expect: &test.HTTPResponse{
				Code: http.StatusOK,
				Body: &response.BookReviewResponse{
					BookReview: entity.NewBookReview(gentity.NewReview(review1), gentity.NewUser(user1)),
				},
			},
		},
		{
			name:     "failed to invalid book id param",
			setup:    func(ctx context.Context, t *testing.T, mocks *test.Mocks, ctrl *gomock.Controller) {},
			bookID:   "1.1",
			reviewID: "1",
			expect: &test.HTTPResponse{
				Code: http.StatusBadRequest,
			},
		},
		{
			name:     "failed to invalid review id param",
			setup:    func(ctx context.Context, t *testing.T, mocks *test.Mocks, ctrl *gomock.Controller) {},
			bookID:   "1",
			reviewID: "1.1",
			expect: &test.HTTPResponse{
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
					GetReview(gomock.Any(), &book.GetReviewRequest{ReviewId: 1}).
					Return(&book.ReviewResponse{Review: review1}, nil)
			},
			bookID:   "1",
			reviewID: "1",
			expect: &test.HTTPResponse{
				Code: http.StatusInternalServerError,
			},
		},
		{
			name: "failed to get review",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks, ctrl *gomock.Controller) {
				mocks.BookService.EXPECT().
					GetBook(gomock.Any(), &book.GetBookRequest{BookId: 1}).
					Return(&book.BookResponse{Book: book1}, nil)
				mocks.BookService.EXPECT().
					GetReview(gomock.Any(), &book.GetReviewRequest{ReviewId: 1}).
					Return(nil, test.ErrMock)
			},
			bookID:   "1",
			reviewID: "1",
			expect: &test.HTTPResponse{
				Code: http.StatusInternalServerError,
			},
		},
		{
			name: "failed to get user",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks, ctrl *gomock.Controller) {
				mocks.BookService.EXPECT().
					GetBook(gomock.Any(), &book.GetBookRequest{BookId: 1}).
					Return(&book.BookResponse{Book: book1}, nil)
				mocks.BookService.EXPECT().
					GetReview(gomock.Any(), &book.GetReviewRequest{ReviewId: 1}).
					Return(&book.ReviewResponse{Review: review1}, nil)
				mocks.UserService.EXPECT().
					GetUser(gomock.Any(), &user.GetUserRequest{UserId: "00000000-0000-0000-0000-000000000000"}).
					Return(nil, test.ErrMock)
			},
			bookID:   "1",
			reviewID: "1",
			expect: &test.HTTPResponse{
				Code: http.StatusInternalServerError,
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			path := "/v1/books/" + tt.bookID + "/reviews/" + tt.reviewID
			req := test.NewHTTPRequest(t, http.MethodGet, path, nil)
			testHTTP(t, tt.setup, tt.expect, req)
		})
	}
}

func TestReview_getUserReview(t *testing.T) {
	t.Parallel()

	user1 := testUser("00000000-0000-0000-0000-000000000000")
	user2 := testUser("11111111-1111-1111-1111-111111111111")
	book1 := testBook(1)
	review1 := testReview(1, book1.Id, user1.Id)
	review2 := testReview(2, book1.Id, user2.Id)

	tests := []struct {
		name     string
		setup    func(ctx context.Context, t *testing.T, mocks *test.Mocks, ctrl *gomock.Controller)
		reviewID string
		expect   *test.HTTPResponse
	}{
		{
			name: "success",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks, ctrl *gomock.Controller) {
				mocks.BookService.EXPECT().
					GetReview(gomock.Any(), &book.GetReviewRequest{ReviewId: 1}).
					Return(&book.ReviewResponse{Review: review1}, nil)
				mocks.BookService.EXPECT().
					GetBook(gomock.Any(), &book.GetBookRequest{BookId: 1}).
					Return(&book.BookResponse{Book: book1}, nil)
			},
			reviewID: "1",
			expect: &test.HTTPResponse{
				Code: http.StatusOK,
				Body: &response.UserReviewResponse{
					UserReview: entity.NewUserReview(gentity.NewReview(review1), gentity.NewBook(book1)),
				},
			},
		},
		{
			name:     "failed to invalid review id param",
			setup:    func(ctx context.Context, t *testing.T, mocks *test.Mocks, ctrl *gomock.Controller) {},
			reviewID: "1.1",
			expect: &test.HTTPResponse{
				Code: http.StatusBadRequest,
			},
		},
		{
			name: "failed to get review",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks, ctrl *gomock.Controller) {
				mocks.BookService.EXPECT().
					GetReview(gomock.Any(), &book.GetReviewRequest{ReviewId: 1}).
					Return(nil, test.ErrMock)
			},
			reviewID: "1",
			expect: &test.HTTPResponse{
				Code: http.StatusInternalServerError,
			},
		},
		{
			name: "failed to invalid user id in review",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks, ctrl *gomock.Controller) {
				mocks.BookService.EXPECT().
					GetReview(gomock.Any(), &book.GetReviewRequest{ReviewId: 1}).
					Return(&book.ReviewResponse{Review: review2}, nil)
			},
			reviewID: "1",
			expect: &test.HTTPResponse{
				Code: http.StatusBadRequest,
			},
		},
		{
			name: "failed to get book",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks, ctrl *gomock.Controller) {
				mocks.BookService.EXPECT().
					GetReview(gomock.Any(), &book.GetReviewRequest{ReviewId: 1}).
					Return(&book.ReviewResponse{Review: review1}, nil)
				mocks.BookService.EXPECT().
					GetBook(gomock.Any(), &book.GetBookRequest{BookId: 1}).
					Return(nil, test.ErrMock)
			},
			reviewID: "1",
			expect: &test.HTTPResponse{
				Code: http.StatusInternalServerError,
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			path := "/v1/users/00000000-0000-0000-0000-000000000000/reviews/" + tt.reviewID
			req := test.NewHTTPRequest(t, http.MethodGet, path, nil)
			testHTTP(t, tt.setup, tt.expect, req)
		})
	}
}

func testReview(id int64, bookID int64, userID string) *book.Review {
	now := datetime.FormatTime(test.TimeMock)
	return &book.Review{
		Id:         id,
		BookId:     bookID,
		UserId:     userID,
		Score:      3,
		Impression: "テストレビューです",
		CreatedAt:  now,
		UpdatedAt:  now,
	}
}
