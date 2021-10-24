package v1

import (
	"context"
	"net/http"
	"strconv"
	"testing"

	gentity "github.com/calmato/gran-book/api/service/internal/gateway/entity"
	"github.com/calmato/gran-book/api/service/internal/gateway/native/entity"
	request "github.com/calmato/gran-book/api/service/internal/gateway/native/request/v1"
	response "github.com/calmato/gran-book/api/service/internal/gateway/native/response/v1"
	"github.com/calmato/gran-book/api/service/pkg/datetime"
	"github.com/calmato/gran-book/api/service/pkg/test"
	"github.com/calmato/gran-book/api/service/proto/book"
	"github.com/golang/mock/gomock"
)

func TestBookshelf_ListBookshelf(t *testing.T) {
	t.Parallel()

	user1 := testAuth("00000000-0000-0000-0000-000000000000")
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
		expect *test.HTTPResponse
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
			expect: &test.HTTPResponse{
				Code: http.StatusOK,
				Body: response.NewBookshelfListResponse(
					gentity.NewBookshelves(bookshelves),
					gentity.NewBooks(books).Map(),
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
			expect: &test.HTTPResponse{
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
			expect: &test.HTTPResponse{
				Code: http.StatusInternalServerError,
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			path := "/v1/users/00000000-0000-0000-0000-000000000000/books" + tt.query
			req := test.NewHTTPRequest(t, http.MethodGet, path, nil)
			testHTTP(t, tt.setup, tt.expect, req)
		})
	}
}

func TestBookshelf_GetBookshelf(t *testing.T) {
	t.Parallel()

	user1 := testAuth("00000000-0000-0000-0000-000000000000")
	book1 := testBook(1)
	bookshelf1 := testBookshelf(1, book1.Id, user1.Id)

	tests := []struct {
		name   string
		setup  func(ctx context.Context, t *testing.T, mocks *test.Mocks, ctrl *gomock.Controller)
		bookID string
		expect *test.HTTPResponse
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
			},
			bookID: "1",
			expect: &test.HTTPResponse{
				Code: http.StatusOK,
				Body: response.NewBookshelfResponse(
					gentity.NewBookshelf(bookshelf1),
					gentity.NewBook(book1),
				),
			},
		},
		{
			name:   "failed to invalid book param",
			setup:  func(ctx context.Context, t *testing.T, mocks *test.Mocks, ctrl *gomock.Controller) {},
			bookID: "1.1",
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
					GetBookshelf(gomock.Any(), &book.GetBookshelfRequest{
						UserId: "00000000-0000-0000-0000-000000000000",
						BookId: 1,
					}).
					Return(&book.BookshelfResponse{Bookshelf: bookshelf1}, nil)
			},
			bookID: "1",
			expect: &test.HTTPResponse{
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
			},
			bookID: "1",
			expect: &test.HTTPResponse{
				Code: http.StatusInternalServerError,
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			path := "/v1/users/00000000-0000-0000-0000-000000000000/books/" + tt.bookID
			req := test.NewHTTPRequest(t, http.MethodGet, path, nil)
			testHTTP(t, tt.setup, tt.expect, req)
		})
	}
}

func TestBookshelf_ReadBookshelf(t *testing.T) {
	t.Parallel()

	user1 := testAuth("00000000-0000-0000-0000-000000000000")
	book1 := testBook(1)
	bookshelf1 := testBookshelf(1, book1.Id, user1.Id)

	tests := []struct {
		name   string
		setup  func(ctx context.Context, t *testing.T, mocks *test.Mocks, ctrl *gomock.Controller)
		bookID string
		req    *request.ReadBookshelfRequest
		expect *test.HTTPResponse
	}{
		{
			name: "success",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks, ctrl *gomock.Controller) {
				mocks.BookService.EXPECT().
					GetBook(gomock.Any(), &book.GetBookRequest{BookId: 1}).
					Return(&book.BookResponse{Book: book1}, nil)
				mocks.BookService.EXPECT().
					ReadBookshelf(gomock.Any(), &book.ReadBookshelfRequest{
						UserId:     "00000000-0000-0000-0000-000000000000",
						BookId:     1,
						Impression: "テストの感想です",
						ReadOn:     "2021-08-02",
					}).
					Return(&book.BookshelfResponse{Bookshelf: bookshelf1}, nil)
			},
			bookID: "1",
			req: &request.ReadBookshelfRequest{
				Impression: "テストの感想です",
				ReadOn:     "2021-08-02",
			},
			expect: &test.HTTPResponse{
				Code: http.StatusOK,
				Body: response.NewBookshelfResponse(
					gentity.NewBookshelf(bookshelf1),
					gentity.NewBook(book1),
				),
			},
		},
		{
			name:   "failed to invalid book param",
			setup:  func(ctx context.Context, t *testing.T, mocks *test.Mocks, ctrl *gomock.Controller) {},
			bookID: "1.1",
			req:    &request.ReadBookshelfRequest{},
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
					ReadBookshelf(gomock.Any(), &book.ReadBookshelfRequest{
						UserId:     "00000000-0000-0000-0000-000000000000",
						BookId:     1,
						Impression: "テストの感想です",
						ReadOn:     "2021-08-02",
					}).
					Return(&book.BookshelfResponse{Bookshelf: bookshelf1}, nil)
			},
			bookID: "1",
			req: &request.ReadBookshelfRequest{
				Impression: "テストの感想です",
				ReadOn:     "2021-08-02",
			},
			expect: &test.HTTPResponse{
				Code: http.StatusInternalServerError,
			},
		},
		{
			name: "failed to read bookshelf",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks, ctrl *gomock.Controller) {
				mocks.BookService.EXPECT().
					GetBook(gomock.Any(), &book.GetBookRequest{BookId: 1}).
					Return(&book.BookResponse{Book: book1}, nil)
				mocks.BookService.EXPECT().
					ReadBookshelf(gomock.Any(), &book.ReadBookshelfRequest{
						UserId:     "00000000-0000-0000-0000-000000000000",
						BookId:     1,
						Impression: "テストの感想です",
						ReadOn:     "2021-08-02",
					}).
					Return(nil, test.ErrMock)
			},
			bookID: "1",
			req: &request.ReadBookshelfRequest{
				Impression: "テストの感想です",
				ReadOn:     "2021-08-02",
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
			path := "/v1/users/00000000-0000-0000-0000-000000000000/books/" + tt.bookID + "/read"
			req := test.NewHTTPRequest(t, http.MethodPost, path, tt.req)
			testHTTP(t, tt.setup, tt.expect, req)
		})
	}
}

func TestBookshelf_ReadingBookshelf(t *testing.T) {
	t.Parallel()

	user1 := testAuth("00000000-0000-0000-0000-000000000000")
	book1 := testBook(1)
	bookshelf1 := testBookshelf(1, book1.Id, user1.Id)

	tests := []struct {
		name   string
		setup  func(ctx context.Context, t *testing.T, mocks *test.Mocks, ctrl *gomock.Controller)
		bookID string
		expect *test.HTTPResponse
	}{
		{
			name: "success",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks, ctrl *gomock.Controller) {
				mocks.BookService.EXPECT().
					GetBook(gomock.Any(), &book.GetBookRequest{BookId: 1}).
					Return(&book.BookResponse{Book: book1}, nil)
				mocks.BookService.EXPECT().
					ReadingBookshelf(gomock.Any(), &book.ReadingBookshelfRequest{
						UserId: "00000000-0000-0000-0000-000000000000",
						BookId: 1,
					}).
					Return(&book.BookshelfResponse{Bookshelf: bookshelf1}, nil)
			},
			bookID: "1",
			expect: &test.HTTPResponse{
				Code: http.StatusOK,
				Body: response.NewBookshelfResponse(
					gentity.NewBookshelf(bookshelf1),
					gentity.NewBook(book1),
				),
			},
		},
		{
			name:   "failed to invalid book param",
			setup:  func(ctx context.Context, t *testing.T, mocks *test.Mocks, ctrl *gomock.Controller) {},
			bookID: "1.1",
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
					ReadingBookshelf(gomock.Any(), &book.ReadingBookshelfRequest{
						UserId: "00000000-0000-0000-0000-000000000000",
						BookId: 1,
					}).
					Return(&book.BookshelfResponse{Bookshelf: bookshelf1}, nil)
			},
			bookID: "1",
			expect: &test.HTTPResponse{
				Code: http.StatusInternalServerError,
			},
		},
		{
			name: "failed to reading bookshelf",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks, ctrl *gomock.Controller) {
				mocks.BookService.EXPECT().
					GetBook(gomock.Any(), &book.GetBookRequest{BookId: 1}).
					Return(&book.BookResponse{Book: book1}, nil)
				mocks.BookService.EXPECT().
					ReadingBookshelf(gomock.Any(), &book.ReadingBookshelfRequest{
						UserId: "00000000-0000-0000-0000-000000000000",
						BookId: 1,
					}).
					Return(nil, test.ErrMock)
			},
			bookID: "1",
			expect: &test.HTTPResponse{
				Code: http.StatusInternalServerError,
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			path := "/v1/users/00000000-0000-0000-0000-000000000000/books/" + tt.bookID + "/reading"
			req := test.NewHTTPRequest(t, http.MethodPost, path, nil)
			testHTTP(t, tt.setup, tt.expect, req)
		})
	}
}

func TestBookshelf_StackedBookshelf(t *testing.T) {
	t.Parallel()

	user1 := testAuth("00000000-0000-0000-0000-000000000000")
	book1 := testBook(1)
	bookshelf1 := testBookshelf(1, book1.Id, user1.Id)

	tests := []struct {
		name   string
		setup  func(ctx context.Context, t *testing.T, mocks *test.Mocks, ctrl *gomock.Controller)
		bookID string
		expect *test.HTTPResponse
	}{
		{
			name: "success",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks, ctrl *gomock.Controller) {
				mocks.BookService.EXPECT().
					GetBook(gomock.Any(), &book.GetBookRequest{BookId: 1}).
					Return(&book.BookResponse{Book: book1}, nil)
				mocks.BookService.EXPECT().
					StackedBookshelf(gomock.Any(), &book.StackedBookshelfRequest{
						UserId: "00000000-0000-0000-0000-000000000000",
						BookId: 1,
					}).
					Return(&book.BookshelfResponse{Bookshelf: bookshelf1}, nil)
			},
			bookID: "1",
			expect: &test.HTTPResponse{
				Code: http.StatusOK,
				Body: response.NewBookshelfResponse(
					gentity.NewBookshelf(bookshelf1),
					gentity.NewBook(book1),
				),
			},
		},
		{
			name:   "failed to invalid book param",
			setup:  func(ctx context.Context, t *testing.T, mocks *test.Mocks, ctrl *gomock.Controller) {},
			bookID: "1.1",
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
					StackedBookshelf(gomock.Any(), &book.StackedBookshelfRequest{
						UserId: "00000000-0000-0000-0000-000000000000",
						BookId: 1,
					}).
					Return(&book.BookshelfResponse{Bookshelf: bookshelf1}, nil)
			},
			bookID: "1",
			expect: &test.HTTPResponse{
				Code: http.StatusInternalServerError,
			},
		},
		{
			name: "failed to stacked bookshelf",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks, ctrl *gomock.Controller) {
				mocks.BookService.EXPECT().
					GetBook(gomock.Any(), &book.GetBookRequest{BookId: 1}).
					Return(&book.BookResponse{Book: book1}, nil)
				mocks.BookService.EXPECT().
					StackedBookshelf(gomock.Any(), &book.StackedBookshelfRequest{
						UserId: "00000000-0000-0000-0000-000000000000",
						BookId: 1,
					}).
					Return(nil, test.ErrMock)
			},
			bookID: "1",
			expect: &test.HTTPResponse{
				Code: http.StatusInternalServerError,
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			path := "/v1/users/00000000-0000-0000-0000-000000000000/books/" + tt.bookID + "/stack"
			req := test.NewHTTPRequest(t, http.MethodPost, path, nil)
			testHTTP(t, tt.setup, tt.expect, req)
		})
	}
}

func TestBookshelf_WantBookshelf(t *testing.T) {
	t.Parallel()

	user1 := testAuth("00000000-0000-0000-0000-000000000000")
	book1 := testBook(1)
	bookshelf1 := testBookshelf(1, book1.Id, user1.Id)

	tests := []struct {
		name   string
		setup  func(ctx context.Context, t *testing.T, mocks *test.Mocks, ctrl *gomock.Controller)
		bookID string
		expect *test.HTTPResponse
	}{
		{
			name: "success",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks, ctrl *gomock.Controller) {
				mocks.BookService.EXPECT().
					GetBook(gomock.Any(), &book.GetBookRequest{BookId: 1}).
					Return(&book.BookResponse{Book: book1}, nil)
				mocks.BookService.EXPECT().
					WantBookshelf(gomock.Any(), &book.WantBookshelfRequest{
						UserId: "00000000-0000-0000-0000-000000000000",
						BookId: 1,
					}).
					Return(&book.BookshelfResponse{Bookshelf: bookshelf1}, nil)
			},
			bookID: "1",
			expect: &test.HTTPResponse{
				Code: http.StatusOK,
				Body: response.NewBookshelfResponse(
					gentity.NewBookshelf(bookshelf1),
					gentity.NewBook(book1),
				),
			},
		},
		{
			name:   "failed to invalid book param",
			setup:  func(ctx context.Context, t *testing.T, mocks *test.Mocks, ctrl *gomock.Controller) {},
			bookID: "1.1",
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
					WantBookshelf(gomock.Any(), &book.WantBookshelfRequest{
						UserId: "00000000-0000-0000-0000-000000000000",
						BookId: 1,
					}).
					Return(&book.BookshelfResponse{Bookshelf: bookshelf1}, nil)
			},
			bookID: "1",
			expect: &test.HTTPResponse{
				Code: http.StatusInternalServerError,
			},
		},
		{
			name: "failed to want bookshelf",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks, ctrl *gomock.Controller) {
				mocks.BookService.EXPECT().
					GetBook(gomock.Any(), &book.GetBookRequest{BookId: 1}).
					Return(&book.BookResponse{Book: book1}, nil)
				mocks.BookService.EXPECT().
					WantBookshelf(gomock.Any(), &book.WantBookshelfRequest{
						UserId: "00000000-0000-0000-0000-000000000000",
						BookId: 1,
					}).
					Return(nil, test.ErrMock)
			},
			bookID: "1",
			expect: &test.HTTPResponse{
				Code: http.StatusInternalServerError,
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			path := "/v1/users/00000000-0000-0000-0000-000000000000/books/" + tt.bookID + "/want"
			req := test.NewHTTPRequest(t, http.MethodPost, path, nil)
			testHTTP(t, tt.setup, tt.expect, req)
		})
	}
}

func TestBookshelf_ReleaseBookshelf(t *testing.T) {
	t.Parallel()

	user1 := testAuth("00000000-0000-0000-0000-000000000000")
	book1 := testBook(1)
	bookshelf1 := testBookshelf(1, book1.Id, user1.Id)

	tests := []struct {
		name   string
		setup  func(ctx context.Context, t *testing.T, mocks *test.Mocks, ctrl *gomock.Controller)
		bookID string
		expect *test.HTTPResponse
	}{
		{
			name: "success",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks, ctrl *gomock.Controller) {
				mocks.BookService.EXPECT().
					GetBook(gomock.Any(), &book.GetBookRequest{BookId: 1}).
					Return(&book.BookResponse{Book: book1}, nil)
				mocks.BookService.EXPECT().
					ReleaseBookshelf(gomock.Any(), &book.ReleaseBookshelfRequest{
						UserId: "00000000-0000-0000-0000-000000000000",
						BookId: 1,
					}).
					Return(&book.BookshelfResponse{Bookshelf: bookshelf1}, nil)
			},
			bookID: "1",
			expect: &test.HTTPResponse{
				Code: http.StatusOK,
				Body: response.NewBookshelfResponse(
					gentity.NewBookshelf(bookshelf1),
					gentity.NewBook(book1),
				),
			},
		},
		{
			name:   "failed to invalid book param",
			setup:  func(ctx context.Context, t *testing.T, mocks *test.Mocks, ctrl *gomock.Controller) {},
			bookID: "1.1",
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
					ReleaseBookshelf(gomock.Any(), &book.ReleaseBookshelfRequest{
						UserId: "00000000-0000-0000-0000-000000000000",
						BookId: 1,
					}).
					Return(&book.BookshelfResponse{Bookshelf: bookshelf1}, nil)
			},
			bookID: "1",
			expect: &test.HTTPResponse{
				Code: http.StatusInternalServerError,
			},
		},
		{
			name: "failed to release bookshelf",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks, ctrl *gomock.Controller) {
				mocks.BookService.EXPECT().
					GetBook(gomock.Any(), &book.GetBookRequest{BookId: 1}).
					Return(&book.BookResponse{Book: book1}, nil)
				mocks.BookService.EXPECT().
					ReleaseBookshelf(gomock.Any(), &book.ReleaseBookshelfRequest{
						UserId: "00000000-0000-0000-0000-000000000000",
						BookId: 1,
					}).
					Return(nil, test.ErrMock)
			},
			bookID: "1",
			expect: &test.HTTPResponse{
				Code: http.StatusInternalServerError,
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			path := "/v1/users/00000000-0000-0000-0000-000000000000/books/" + tt.bookID + "/release"
			req := test.NewHTTPRequest(t, http.MethodPost, path, nil)
			testHTTP(t, tt.setup, tt.expect, req)
		})
	}
}

func TestBookshelf_DeleteBookshelf(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		setup  func(ctx context.Context, t *testing.T, mocks *test.Mocks, ctrl *gomock.Controller)
		bookID string
		expect *test.HTTPResponse
	}{
		{
			name: "success",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks, ctrl *gomock.Controller) {
				mocks.BookService.EXPECT().
					DeleteBookshelf(gomock.Any(), &book.DeleteBookshelfRequest{
						UserId: "00000000-0000-0000-0000-000000000000",
						BookId: 1,
					}).
					Return(&book.Empty{}, nil)
			},
			bookID: "1",
			expect: &test.HTTPResponse{
				Code: http.StatusNoContent,
			},
		},
		{
			name:   "failed to invalid book param",
			setup:  func(ctx context.Context, t *testing.T, mocks *test.Mocks, ctrl *gomock.Controller) {},
			bookID: "1.1",
			expect: &test.HTTPResponse{
				Code: http.StatusBadRequest,
			},
		},
		{
			name: "failed to delete bookshelf",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks, ctrl *gomock.Controller) {
				mocks.BookService.EXPECT().
					DeleteBookshelf(gomock.Any(), &book.DeleteBookshelfRequest{
						UserId: "00000000-0000-0000-0000-000000000000",
						BookId: 1,
					}).
					Return(nil, test.ErrMock)
			},
			bookID: "1",
			expect: &test.HTTPResponse{
				Code: http.StatusInternalServerError,
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			path := "/v1/users/00000000-0000-0000-0000-000000000000/books/" + tt.bookID
			req := test.NewHTTPRequest(t, http.MethodDelete, path, nil)
			testHTTP(t, tt.setup, tt.expect, req)
		})
	}
}

func testBookshelf(id int64, bookID int64, userID string) *book.Bookshelf {
	now := datetime.FormatTime(test.TimeMock)
	return &book.Bookshelf{
		Id:        id,
		BookId:    bookID,
		UserId:    userID,
		ReviewId:  0,
		Status:    book.BookshelfStatus_BOOKSHELF_STATUS_READING,
		CreatedAt: now,
		UpdatedAt: now,
	}
}
