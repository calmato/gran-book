package server

import (
	"context"
	"testing"

	"github.com/calmato/gran-book/api/service/internal/book/domain/book"
	"github.com/calmato/gran-book/api/service/pkg/exception"
	"github.com/calmato/gran-book/api/service/pkg/test"
	pb "github.com/calmato/gran-book/api/service/proto/book"
	"github.com/golang/mock/gomock"
	"google.golang.org/grpc/codes"
)

func TestBookServer_ListBookshelf(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	book1 := testBook(1)
	bookshelf1 := testBookshelf(book1.ID, 1, "user01")
	bookshelf2 := testBookshelf(book1.ID, 2, "user01")

	type args struct {
		req *pb.ListBookshelfRequest
	}
	testCases := []struct {
		name  string
		setup func(context.Context, *testing.T, *test.Mocks)
		args  args
		want  *test.Response
	}{
		{
			name: "success",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.BookRequestValidation.EXPECT().
					ListBookshelf(gomock.Any()).
					Return(nil)
				mocks.BookApplication.EXPECT().
					ListBookshelf(ctx, gomock.Any()).
					Return([]*book.Bookshelf{bookshelf1, bookshelf2}, 2, nil)
			},
			args: args{
				req: &pb.ListBookshelfRequest{
					UserId: "user01",
					Limit:  100,
					Offset: 0,
				},
			},
			want: &test.Response{
				Code:    codes.OK,
				Message: getBookshelfListResponse([]*book.Bookshelf{bookshelf1, bookshelf2}, 100, 0, 2),
			},
		},
		{
			name: "success with order",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.BookRequestValidation.EXPECT().
					ListBookshelf(gomock.Any()).
					Return(nil)
				mocks.BookApplication.EXPECT().
					ListBookshelf(ctx, gomock.Any()).
					Return([]*book.Bookshelf{bookshelf1, bookshelf2}, 2, nil)
			},
			args: args{
				req: &pb.ListBookshelfRequest{
					UserId: "user01",
					Limit:  100,
					Offset: 0,
					Order: &pb.Order{
						Field:   "created_at",
						OrderBy: pb.OrderBy_ORDER_BY_ASC,
					},
				},
			},
			want: &test.Response{
				Code:    codes.OK,
				Message: getBookshelfListResponse([]*book.Bookshelf{bookshelf1, bookshelf2}, 100, 0, 2),
			},
		},
		{
			name: "failed: invalid argument",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.BookRequestValidation.EXPECT().
					ListBookshelf(gomock.Any()).
					Return(exception.ErrInvalidRequestValidation.New(test.ErrMock))
			},
			args: args{
				req: &pb.ListBookshelfRequest{},
			},
			want: &test.Response{
				Code:    codes.InvalidArgument,
				Message: nil,
			},
		},
		{
			name: "failed: internel error",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.BookRequestValidation.EXPECT().
					ListBookshelf(gomock.Any()).
					Return(nil)
				mocks.BookApplication.EXPECT().
					ListBookshelf(ctx, gomock.Any()).
					Return(nil, 0, exception.ErrInDatastore.New(test.ErrMock))
			},
			args: args{
				req: &pb.ListBookshelfRequest{
					UserId: "user01",
					Limit:  100,
					Offset: 0,
				},
			},
			want: &test.Response{
				Code:    codes.Internal,
				Message: nil,
			},
		},
	}

	for _, tt := range testCases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			mocks := test.NewMocks(ctrl)
			tt.setup(ctx, t, mocks)
			target := NewBookServer(mocks.BookRequestValidation, mocks.BookApplication)

			res, err := target.ListBookshelf(ctx, tt.args.req)
			test.GRPC(t, tt.want, res, err)
		})
	}
}

func TestBookServer_ListBookReview(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	book1 := testBook(1)
	review1 := testReview(book1.ID, 1, "user01")
	review2 := testReview(book1.ID, 2, "user02")

	type args struct {
		req *pb.ListBookReviewRequest
	}
	testCases := []struct {
		name  string
		setup func(context.Context, *testing.T, *test.Mocks)
		args  args
		want  *test.Response
	}{
		{
			name: "success",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.BookRequestValidation.EXPECT().
					ListBookReview(gomock.Any()).
					Return(nil)
				mocks.BookApplication.EXPECT().
					ListBookReview(ctx, 1, 100, 0).
					Return([]*book.Review{review1, review2}, 2, nil)
			},
			args: args{
				req: &pb.ListBookReviewRequest{
					BookId: 1,
					Limit:  100,
					Offset: 0,
				},
			},
			want: &test.Response{
				Code:    codes.OK,
				Message: getReviewListResponse([]*book.Review{review1, review2}, 100, 0, 2),
			},
		},
		{
			name: "failed: invalid argument",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.BookRequestValidation.EXPECT().
					ListBookReview(gomock.Any()).
					Return(exception.ErrInvalidRequestValidation.New(test.ErrMock))
			},
			args: args{
				req: &pb.ListBookReviewRequest{},
			},
			want: &test.Response{
				Code:    codes.InvalidArgument,
				Message: nil,
			},
		},
		{
			name: "failed: internel error",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.BookRequestValidation.EXPECT().
					ListBookReview(gomock.Any()).
					Return(nil)
				mocks.BookApplication.EXPECT().
					ListBookReview(ctx, 1, 100, 0).
					Return(nil, 0, exception.ErrInDatastore.New(test.ErrMock))
			},
			args: args{
				req: &pb.ListBookReviewRequest{
					BookId: 1,
					Limit:  100,
					Offset: 0,
				},
			},
			want: &test.Response{
				Code:    codes.Internal,
				Message: nil,
			},
		},
	}

	for _, tt := range testCases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			mocks := test.NewMocks(ctrl)
			tt.setup(ctx, t, mocks)
			target := NewBookServer(mocks.BookRequestValidation, mocks.BookApplication)

			res, err := target.ListBookReview(ctx, tt.args.req)
			test.GRPC(t, tt.want, res, err)
		})
	}
}

func TestBookServer_ListUserReview(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	book1 := testBook(1)
	book2 := testBook(2)
	review1 := testReview(book1.ID, 1, "user01")
	review2 := testReview(book2.ID, 2, "user01")

	type args struct {
		req *pb.ListUserReviewRequest
	}
	testCases := []struct {
		name  string
		setup func(context.Context, *testing.T, *test.Mocks)
		args  args
		want  *test.Response
	}{
		{
			name: "success",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.BookRequestValidation.EXPECT().
					ListUserReview(gomock.Any()).
					Return(nil)
				mocks.BookApplication.EXPECT().
					ListUserReview(ctx, "user01", 100, 0).
					Return([]*book.Review{review1, review2}, 2, nil)
			},
			args: args{
				req: &pb.ListUserReviewRequest{
					UserId: "user01",
					Limit:  100,
					Offset: 0,
				},
			},
			want: &test.Response{
				Code:    codes.OK,
				Message: getReviewListResponse([]*book.Review{review1, review2}, 100, 0, 2),
			},
		},
		{
			name: "failed: invalid argument",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.BookRequestValidation.EXPECT().
					ListUserReview(gomock.Any()).
					Return(exception.ErrInvalidRequestValidation.New(test.ErrMock))
			},
			args: args{
				req: &pb.ListUserReviewRequest{},
			},
			want: &test.Response{
				Code:    codes.InvalidArgument,
				Message: nil,
			},
		},
		{
			name: "failed: internel error",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.BookRequestValidation.EXPECT().
					ListUserReview(gomock.Any()).
					Return(nil)
				mocks.BookApplication.EXPECT().
					ListUserReview(ctx, "user01", 100, 0).
					Return(nil, 0, exception.ErrInDatastore.New(test.ErrMock))
			},
			args: args{
				req: &pb.ListUserReviewRequest{
					UserId: "user01",
					Limit:  100,
					Offset: 0,
				},
			},
			want: &test.Response{
				Code:    codes.Internal,
				Message: nil,
			},
		},
	}

	for _, tt := range testCases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			mocks := test.NewMocks(ctrl)
			tt.setup(ctx, t, mocks)
			target := NewBookServer(mocks.BookRequestValidation, mocks.BookApplication)

			res, err := target.ListUserReview(ctx, tt.args.req)
			test.GRPC(t, tt.want, res, err)
		})
	}
}
func TestBookServer_ListUserMonthlyResult(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	results := make(book.MonthlyResults, 2)
	results[0] = &book.MonthlyResult{Year: 2021, Month: 9, ReadTotal: 3}
	results[1] = &book.MonthlyResult{Year: 2021, Month: 8, ReadTotal: 8}

	type args struct {
		req *pb.ListUserMonthlyResultRequest
	}
	testCases := []struct {
		name  string
		setup func(context.Context, *testing.T, *test.Mocks)
		args  args
		want  *test.Response
	}{
		{
			name: "success",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.BookRequestValidation.EXPECT().
					ListUserMonthlyResult(gomock.Any()).
					Return(nil)
				mocks.BookApplication.EXPECT().
					ListUserMonthlyResult(ctx, "user01", "2021-08-01", "2021-09-30").
					Return(results, nil)
			},
			args: args{
				req: &pb.ListUserMonthlyResultRequest{
					UserId:    "user01",
					SinceDate: "2021-08-01",
					UntilDate: "2021-09-30",
				},
			},
			want: &test.Response{
				Code:    codes.OK,
				Message: getUserMonthlyResultListResponse(results),
			},
		},
		{
			name: "failed: invalid argument",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.BookRequestValidation.EXPECT().
					ListUserMonthlyResult(gomock.Any()).
					Return(exception.ErrInvalidRequestValidation.New(test.ErrMock))
			},
			args: args{
				req: &pb.ListUserMonthlyResultRequest{},
			},
			want: &test.Response{
				Code:    codes.InvalidArgument,
				Message: nil,
			},
		},
		{
			name: "failed: internel error",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.BookRequestValidation.EXPECT().
					ListUserMonthlyResult(gomock.Any()).
					Return(nil)
				mocks.BookApplication.EXPECT().
					ListUserMonthlyResult(ctx, "user01", "2021-08-01", "2021-09-30").
					Return(nil, exception.ErrInDatastore.New(test.ErrMock))
			},
			args: args{
				req: &pb.ListUserMonthlyResultRequest{
					UserId:    "user01",
					SinceDate: "2021-08-01",
					UntilDate: "2021-09-30",
				},
			},
			want: &test.Response{
				Code:    codes.Internal,
				Message: nil,
			},
		},
	}

	for _, tt := range testCases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			mocks := test.NewMocks(ctrl)
			tt.setup(ctx, t, mocks)
			target := NewBookServer(mocks.BookRequestValidation, mocks.BookApplication)

			res, err := target.ListUserMonthlyResult(ctx, tt.args.req)
			test.GRPC(t, tt.want, res, err)
		})
	}
}

func TestBookServer_MultiGetBooks(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	book1 := testBook(1)
	book2 := testBook(2)

	type args struct {
		req *pb.MultiGetBooksRequest
	}
	testCases := []struct {
		name  string
		setup func(context.Context, *testing.T, *test.Mocks)
		args  args
		want  *test.Response
	}{
		{
			name: "success",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.BookRequestValidation.EXPECT().
					MultiGetBooks(gomock.Any()).
					Return(nil)
				mocks.BookApplication.EXPECT().
					MultiGet(ctx, []int{1, 2}).
					Return([]*book.Book{book1, book2}, nil)
			},
			args: args{
				req: &pb.MultiGetBooksRequest{
					BookIds: []int64{1, 2},
				},
			},
			want: &test.Response{
				Code:    codes.OK,
				Message: getBookListResponse([]*book.Book{book1, book2}),
			},
		},
		{
			name: "failed: invalid argument",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.BookRequestValidation.EXPECT().
					MultiGetBooks(gomock.Any()).
					Return(exception.ErrInvalidRequestValidation.New(test.ErrMock))
			},
			args: args{
				req: &pb.MultiGetBooksRequest{},
			},
			want: &test.Response{
				Code:    codes.InvalidArgument,
				Message: nil,
			},
		},
		{
			name: "failed: internel error",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.BookRequestValidation.EXPECT().
					MultiGetBooks(gomock.Any()).
					Return(nil)
				mocks.BookApplication.EXPECT().
					MultiGet(ctx, []int{1, 2}).
					Return(nil, exception.ErrInDatastore.New(test.ErrMock))
			},
			args: args{
				req: &pb.MultiGetBooksRequest{
					BookIds: []int64{1, 2},
				},
			},
			want: &test.Response{
				Code:    codes.Internal,
				Message: nil,
			},
		},
	}

	for _, tt := range testCases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			mocks := test.NewMocks(ctrl)
			tt.setup(ctx, t, mocks)
			target := NewBookServer(mocks.BookRequestValidation, mocks.BookApplication)

			res, err := target.MultiGetBooks(ctx, tt.args.req)
			test.GRPC(t, tt.want, res, err)
		})
	}
}

func TestBookServer_GetBook(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	book1 := testBook(1)

	type args struct {
		req *pb.GetBookRequest
	}
	testCases := []struct {
		name  string
		setup func(context.Context, *testing.T, *test.Mocks)
		args  args
		want  *test.Response
	}{
		{
			name: "success",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.BookRequestValidation.EXPECT().
					GetBook(gomock.Any()).
					Return(nil)
				mocks.BookApplication.EXPECT().
					Get(ctx, 1).
					Return(book1, nil)
			},
			args: args{
				req: &pb.GetBookRequest{
					BookId: 1,
				},
			},
			want: &test.Response{
				Code:    codes.OK,
				Message: getBookResponse(book1),
			},
		},
		{
			name: "failed: invalid argument",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.BookRequestValidation.EXPECT().
					GetBook(gomock.Any()).
					Return(exception.ErrInvalidRequestValidation.New(test.ErrMock))
			},
			args: args{
				req: &pb.GetBookRequest{},
			},
			want: &test.Response{
				Code:    codes.InvalidArgument,
				Message: nil,
			},
		},
		{
			name: "failed: internel error",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.BookRequestValidation.EXPECT().
					GetBook(gomock.Any()).
					Return(nil)
				mocks.BookApplication.EXPECT().
					Get(ctx, 1).
					Return(nil, exception.ErrInDatastore.New(test.ErrMock))
			},
			args: args{
				req: &pb.GetBookRequest{
					BookId: 1,
				},
			},
			want: &test.Response{
				Code:    codes.Internal,
				Message: nil,
			},
		},
	}

	for _, tt := range testCases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			mocks := test.NewMocks(ctrl)
			tt.setup(ctx, t, mocks)
			target := NewBookServer(mocks.BookRequestValidation, mocks.BookApplication)

			res, err := target.GetBook(ctx, tt.args.req)
			test.GRPC(t, tt.want, res, err)
		})
	}
}

func TestBookServer_GetBookByIsbn(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	book1 := testBook(1)

	type args struct {
		req *pb.GetBookByIsbnRequest
	}
	testCases := []struct {
		name  string
		setup func(context.Context, *testing.T, *test.Mocks)
		args  args
		want  *test.Response
	}{
		{
			name: "success",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.BookRequestValidation.EXPECT().
					GetBookByIsbn(gomock.Any()).
					Return(nil)
				mocks.BookApplication.EXPECT().
					GetByIsbn(ctx, "9784062938426").
					Return(book1, nil)
			},
			args: args{
				req: &pb.GetBookByIsbnRequest{
					Isbn: "9784062938426",
				},
			},
			want: &test.Response{
				Code:    codes.OK,
				Message: getBookResponse(book1),
			},
		},
		{
			name: "failed: invalid argument",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.BookRequestValidation.EXPECT().
					GetBookByIsbn(gomock.Any()).
					Return(exception.ErrInvalidRequestValidation.New(test.ErrMock))
			},
			args: args{
				req: &pb.GetBookByIsbnRequest{},
			},
			want: &test.Response{
				Code:    codes.InvalidArgument,
				Message: nil,
			},
		},
		{
			name: "failed: internel error",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.BookRequestValidation.EXPECT().
					GetBookByIsbn(gomock.Any()).
					Return(nil)
				mocks.BookApplication.EXPECT().
					GetByIsbn(ctx, "9784062938426").
					Return(nil, exception.ErrInDatastore.New(test.ErrMock))
			},
			args: args{
				req: &pb.GetBookByIsbnRequest{
					Isbn: "9784062938426",
				},
			},
			want: &test.Response{
				Code:    codes.Internal,
				Message: nil,
			},
		},
	}

	for _, tt := range testCases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			mocks := test.NewMocks(ctrl)
			tt.setup(ctx, t, mocks)
			target := NewBookServer(mocks.BookRequestValidation, mocks.BookApplication)

			res, err := target.GetBookByIsbn(ctx, tt.args.req)
			test.GRPC(t, tt.want, res, err)
		})
	}
}

func TestBookServer_GetBookshelf(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	book1 := testBook(1)
	bookshelf1 := testBookshelf(book1.ID, 1, "user01")

	type args struct {
		req *pb.GetBookshelfRequest
	}
	testCases := []struct {
		name  string
		setup func(context.Context, *testing.T, *test.Mocks)
		args  args
		want  *test.Response
	}{
		{
			name: "success",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.BookRequestValidation.EXPECT().
					GetBookshelf(gomock.Any()).
					Return(nil)
				mocks.BookApplication.EXPECT().
					GetBookshelfByUserIDAndBookID(ctx, "user01", 1).
					Return(bookshelf1, nil)
			},
			args: args{
				req: &pb.GetBookshelfRequest{
					BookId: 1,
					UserId: "user01",
				},
			},
			want: &test.Response{
				Code:    codes.OK,
				Message: getBookshelfResponse(bookshelf1),
			},
		},
		{
			name: "failed: invalid argument",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.BookRequestValidation.EXPECT().
					GetBookshelf(gomock.Any()).
					Return(exception.ErrInvalidRequestValidation.New(test.ErrMock))
			},
			args: args{
				req: &pb.GetBookshelfRequest{},
			},
			want: &test.Response{
				Code:    codes.InvalidArgument,
				Message: nil,
			},
		},
		{
			name: "failed: internel error",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.BookRequestValidation.EXPECT().
					GetBookshelf(gomock.Any()).
					Return(nil)
				mocks.BookApplication.EXPECT().
					GetBookshelfByUserIDAndBookID(ctx, "user01", 1).
					Return(nil, exception.ErrInDatastore.New(test.ErrMock))
			},
			args: args{
				req: &pb.GetBookshelfRequest{
					UserId: "user01",
					BookId: 1,
				},
			},
			want: &test.Response{
				Code:    codes.Internal,
				Message: nil,
			},
		},
	}

	for _, tt := range testCases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			mocks := test.NewMocks(ctrl)
			tt.setup(ctx, t, mocks)
			target := NewBookServer(mocks.BookRequestValidation, mocks.BookApplication)

			res, err := target.GetBookshelf(ctx, tt.args.req)
			test.GRPC(t, tt.want, res, err)
		})
	}
}

func TestBookServer_GetReview(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	book1 := testBook(1)
	review1 := testReview(book1.ID, 1, "user01")

	type args struct {
		req *pb.GetReviewRequest
	}
	testCases := []struct {
		name  string
		setup func(context.Context, *testing.T, *test.Mocks)
		args  args
		want  *test.Response
	}{
		{
			name: "success",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.BookRequestValidation.EXPECT().
					GetReview(gomock.Any()).
					Return(nil)
				mocks.BookApplication.EXPECT().
					GetReview(ctx, 1).
					Return(review1, nil)
			},
			args: args{
				req: &pb.GetReviewRequest{
					ReviewId: 1,
				},
			},
			want: &test.Response{
				Code:    codes.OK,
				Message: getReviewResponse(review1),
			},
		},
		{
			name: "failed: invalid argument",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.BookRequestValidation.EXPECT().
					GetReview(gomock.Any()).
					Return(exception.ErrInvalidRequestValidation.New(test.ErrMock))
			},
			args: args{
				req: &pb.GetReviewRequest{},
			},
			want: &test.Response{
				Code:    codes.InvalidArgument,
				Message: nil,
			},
		},
		{
			name: "failed: internel error",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.BookRequestValidation.EXPECT().
					GetReview(gomock.Any()).
					Return(nil)
				mocks.BookApplication.EXPECT().
					GetReview(ctx, 1).
					Return(nil, exception.ErrInDatastore.New(test.ErrMock))
			},
			args: args{
				req: &pb.GetReviewRequest{
					ReviewId: 1,
				},
			},
			want: &test.Response{
				Code:    codes.Internal,
				Message: nil,
			},
		},
	}

	for _, tt := range testCases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			mocks := test.NewMocks(ctrl)
			tt.setup(ctx, t, mocks)
			target := NewBookServer(mocks.BookRequestValidation, mocks.BookApplication)

			res, err := target.GetReview(ctx, tt.args.req)
			test.GRPC(t, tt.want, res, err)
		})
	}
}

func TestBookServer_CreateBook(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type args struct {
		req *pb.CreateBookRequest
	}
	testCases := []struct {
		name  string
		setup func(context.Context, *testing.T, *test.Mocks)
		args  args
		want  *test.Response
	}{
		{
			name: "success",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.BookRequestValidation.EXPECT().
					CreateBook(gomock.Any()).
					Return(nil)
				mocks.BookApplication.EXPECT().
					Create(ctx, gomock.Any()).
					Return(nil)
			},
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
			want: &test.Response{
				Code: codes.OK,
				Message: &pb.BookResponse{
					Book: &pb.Book{
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
						CreatedAt:      "",
						UpdatedAt:      "",
						Authors: []*pb.Author{
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
		},
		{
			name: "failed: invalid argument",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.BookRequestValidation.EXPECT().
					CreateBook(gomock.Any()).
					Return(exception.ErrInvalidRequestValidation.New(test.ErrMock))
			},
			args: args{
				req: &pb.CreateBookRequest{},
			},
			want: &test.Response{
				Code:    codes.InvalidArgument,
				Message: nil,
			},
		},
		{
			name: "failed: internel error",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.BookRequestValidation.EXPECT().
					CreateBook(gomock.Any()).
					Return(nil)
				mocks.BookApplication.EXPECT().
					Create(ctx, gomock.Any()).
					Return(exception.ErrInDatastore.New(test.ErrMock))
			},
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
			want: &test.Response{
				Code:    codes.Internal,
				Message: nil,
			},
		},
	}

	for _, tt := range testCases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			mocks := test.NewMocks(ctrl)
			tt.setup(ctx, t, mocks)
			target := NewBookServer(mocks.BookRequestValidation, mocks.BookApplication)

			res, err := target.CreateBook(ctx, tt.args.req)
			test.GRPC(t, tt.want, res, err)
		})
	}
}

func TestBookServer_UpdateBook(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	book1 := testBook(1)

	type args struct {
		req *pb.UpdateBookRequest
	}
	testCases := []struct {
		name  string
		setup func(context.Context, *testing.T, *test.Mocks)
		args  args
		want  *test.Response
	}{
		{
			name: "success",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.BookRequestValidation.EXPECT().
					UpdateBook(gomock.Any()).
					Return(nil)
				mocks.BookApplication.EXPECT().
					GetByIsbn(ctx, "9784062938426").
					Return(book1, nil)
				mocks.BookApplication.EXPECT().
					Update(ctx, book1).
					Return(nil)
			},
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
			want: &test.Response{
				Code:    codes.OK,
				Message: getBookResponse(book1),
			},
		},
		{
			name: "failed: invalid argument",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.BookRequestValidation.EXPECT().
					UpdateBook(gomock.Any()).
					Return(exception.ErrInvalidRequestValidation.New(test.ErrMock))
			},
			args: args{
				req: &pb.UpdateBookRequest{},
			},
			want: &test.Response{
				Code:    codes.InvalidArgument,
				Message: nil,
			},
		},
		{
			name: "failed: not found",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.BookRequestValidation.EXPECT().
					UpdateBook(gomock.Any()).
					Return(nil)
				mocks.BookApplication.EXPECT().
					GetByIsbn(ctx, "9784062938426").
					Return(nil, exception.ErrNotFound.New(test.ErrMock))
			},
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
			want: &test.Response{
				Code:    codes.NotFound,
				Message: nil,
			},
		},
		{
			name: "failed: internel error",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.BookRequestValidation.EXPECT().
					UpdateBook(gomock.Any()).
					Return(nil)
				mocks.BookApplication.EXPECT().
					GetByIsbn(ctx, "9784062938426").
					Return(book1, nil)
				mocks.BookApplication.EXPECT().
					Update(ctx, gomock.Any()).
					Return(exception.ErrInDatastore.New(test.ErrMock))
			},
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
			want: &test.Response{
				Code:    codes.Internal,
				Message: nil,
			},
		},
	}

	for _, tt := range testCases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			mocks := test.NewMocks(ctrl)
			tt.setup(ctx, t, mocks)
			target := NewBookServer(mocks.BookRequestValidation, mocks.BookApplication)

			res, err := target.UpdateBook(ctx, tt.args.req)
			test.GRPC(t, tt.want, res, err)
		})
	}
}

func TestBookServer_ReadBookshelf(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	book1 := testBook(1)

	type args struct {
		req *pb.ReadBookshelfRequest
	}
	testCases := []struct {
		name  string
		setup func(context.Context, *testing.T, *test.Mocks)
		args  args
		want  *test.Response
	}{
		{
			name: "success",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				bs := &book.Bookshelf{
					Book:   book1,
					Review: &book.Review{},
				}
				mocks.BookRequestValidation.EXPECT().
					ReadBookshelf(gomock.Any()).
					Return(nil)
				mocks.BookApplication.EXPECT().
					GetBookshelfByUserIDAndBookIDWithRelated(ctx, "user01", 1).
					Return(bs, nil)
				mocks.BookApplication.EXPECT().
					CreateOrUpdateBookshelf(ctx, bs).
					Return(nil)
			},
			args: args{
				req: &pb.ReadBookshelfRequest{
					UserId:     "user01",
					BookId:     1,
					ReadOn:     "2021-08-01",
					Impression: "テスト感想です。",
				},
			},
			want: &test.Response{
				Code: codes.OK,
				Message: &pb.BookshelfResponse{
					Bookshelf: &pb.Bookshelf{
						BookId:    1,
						UserId:    "user01",
						Status:    pb.BookshelfStatus_BOOKSHELF_STATUS_READ,
						ReadOn:    "2021-08-01",
						CreatedAt: "",
						UpdatedAt: "",
					},
				},
			},
		},
		{
			name: "failed: invalid argument",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.BookRequestValidation.EXPECT().
					ReadBookshelf(gomock.Any()).
					Return(exception.ErrInvalidRequestValidation.New(test.ErrMock))
			},
			args: args{
				req: &pb.ReadBookshelfRequest{},
			},
			want: &test.Response{
				Code:    codes.InvalidArgument,
				Message: nil,
			},
		}, {
			name: "failed: not found",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.BookRequestValidation.EXPECT().
					ReadBookshelf(gomock.Any()).
					Return(nil)
				mocks.BookApplication.EXPECT().
					GetBookshelfByUserIDAndBookIDWithRelated(ctx, "user01", 1).
					Return(nil, exception.ErrNotFound.New(test.ErrMock))
			},
			args: args{
				req: &pb.ReadBookshelfRequest{
					UserId: "user01",
					BookId: 1,
				},
			},
			want: &test.Response{
				Code:    codes.NotFound,
				Message: nil,
			},
		},

		{
			name: "failed: internel error",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				bs := &book.Bookshelf{
					Book:   book1,
					Review: &book.Review{},
				}
				mocks.BookRequestValidation.EXPECT().
					ReadBookshelf(gomock.Any()).
					Return(nil)
				mocks.BookApplication.EXPECT().
					GetBookshelfByUserIDAndBookIDWithRelated(ctx, "user01", 1).
					Return(bs, nil)
				mocks.BookApplication.EXPECT().
					CreateOrUpdateBookshelf(ctx, bs).
					Return(exception.ErrInDatastore.New(test.ErrMock))
			},
			args: args{
				req: &pb.ReadBookshelfRequest{
					UserId:     "user01",
					BookId:     1,
					ReadOn:     "20210801",
					Impression: "テスト感想です。",
				},
			},
			want: &test.Response{
				Code:    codes.Internal,
				Message: nil,
			},
		},
	}

	for _, tt := range testCases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			mocks := test.NewMocks(ctrl)
			tt.setup(ctx, t, mocks)
			target := NewBookServer(mocks.BookRequestValidation, mocks.BookApplication)

			res, err := target.ReadBookshelf(ctx, tt.args.req)
			test.GRPC(t, tt.want, res, err)
		})
	}
}

func TestBookServer_ReadingBookshelf(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	book1 := testBook(1)

	type args struct {
		req *pb.ReadingBookshelfRequest
	}
	testCases := []struct {
		name  string
		setup func(context.Context, *testing.T, *test.Mocks)
		args  args
		want  *test.Response
	}{
		{
			name: "success",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				bs := &book.Bookshelf{
					Book:   book1,
					Review: &book.Review{},
				}
				mocks.BookRequestValidation.EXPECT().
					ReadingBookshelf(gomock.Any()).
					Return(nil)
				mocks.BookApplication.EXPECT().
					GetBookshelfByUserIDAndBookIDWithRelated(ctx, "user01", 1).
					Return(bs, nil)
				mocks.BookApplication.EXPECT().
					CreateOrUpdateBookshelf(ctx, bs).
					Return(nil)
			},
			args: args{
				req: &pb.ReadingBookshelfRequest{
					UserId: "user01",
					BookId: 1,
				},
			},
			want: &test.Response{
				Code: codes.OK,
				Message: &pb.BookshelfResponse{
					Bookshelf: &pb.Bookshelf{
						BookId:    1,
						UserId:    "user01",
						Status:    pb.BookshelfStatus_BOOKSHELF_STATUS_READING,
						ReadOn:    "",
						CreatedAt: "",
						UpdatedAt: "",
					},
				},
			},
		},
		{
			name: "failed: invalid argument",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.BookRequestValidation.EXPECT().
					ReadingBookshelf(gomock.Any()).
					Return(exception.ErrInvalidRequestValidation.New(test.ErrMock))
			},
			args: args{
				req: &pb.ReadingBookshelfRequest{},
			},
			want: &test.Response{
				Code:    codes.InvalidArgument,
				Message: nil,
			},
		}, {
			name: "failed: not found",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.BookRequestValidation.EXPECT().
					ReadingBookshelf(gomock.Any()).
					Return(nil)
				mocks.BookApplication.EXPECT().
					GetBookshelfByUserIDAndBookIDWithRelated(ctx, "user01", 1).
					Return(nil, exception.ErrNotFound.New(test.ErrMock))
			},
			args: args{
				req: &pb.ReadingBookshelfRequest{
					UserId: "user01",
					BookId: 1,
				},
			},
			want: &test.Response{
				Code:    codes.NotFound,
				Message: nil,
			},
		},

		{
			name: "failed: internel error",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				bs := &book.Bookshelf{
					Book:   book1,
					Review: &book.Review{},
				}
				mocks.BookRequestValidation.EXPECT().
					ReadingBookshelf(gomock.Any()).
					Return(nil)
				mocks.BookApplication.EXPECT().
					GetBookshelfByUserIDAndBookIDWithRelated(ctx, "user01", 1).
					Return(bs, nil)
				mocks.BookApplication.EXPECT().
					CreateOrUpdateBookshelf(ctx, bs).
					Return(exception.ErrInDatastore.New(test.ErrMock))
			},
			args: args{
				req: &pb.ReadingBookshelfRequest{
					UserId: "user01",
					BookId: 1,
				},
			},
			want: &test.Response{
				Code:    codes.Internal,
				Message: nil,
			},
		},
	}

	for _, tt := range testCases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			mocks := test.NewMocks(ctrl)
			tt.setup(ctx, t, mocks)
			target := NewBookServer(mocks.BookRequestValidation, mocks.BookApplication)

			res, err := target.ReadingBookshelf(ctx, tt.args.req)
			test.GRPC(t, tt.want, res, err)
		})
	}
}

func TestBookServer_StackedBookshelf(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	book1 := testBook(1)

	type args struct {
		req *pb.StackedBookshelfRequest
	}
	testCases := []struct {
		name  string
		setup func(context.Context, *testing.T, *test.Mocks)
		args  args
		want  *test.Response
	}{
		{
			name: "success",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				bs := &book.Bookshelf{
					Book:   book1,
					Review: &book.Review{},
				}
				mocks.BookRequestValidation.EXPECT().
					StackedBookshelf(gomock.Any()).
					Return(nil)
				mocks.BookApplication.EXPECT().
					GetBookshelfByUserIDAndBookIDWithRelated(ctx, "user01", 1).
					Return(bs, nil)
				mocks.BookApplication.EXPECT().
					CreateOrUpdateBookshelf(ctx, bs).
					Return(nil)
			},
			args: args{
				req: &pb.StackedBookshelfRequest{
					UserId: "user01",
					BookId: 1,
				},
			},
			want: &test.Response{
				Code: codes.OK,
				Message: &pb.BookshelfResponse{
					Bookshelf: &pb.Bookshelf{
						BookId:    1,
						UserId:    "user01",
						Status:    pb.BookshelfStatus_BOOKSHELF_STATUS_STACKED,
						ReadOn:    "",
						CreatedAt: "",
						UpdatedAt: "",
					},
				},
			},
		},
		{
			name: "failed: invalid argument",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.BookRequestValidation.EXPECT().
					StackedBookshelf(gomock.Any()).
					Return(exception.ErrInvalidRequestValidation.New(test.ErrMock))
			},
			args: args{
				req: &pb.StackedBookshelfRequest{},
			},
			want: &test.Response{
				Code:    codes.InvalidArgument,
				Message: nil,
			},
		}, {
			name: "failed: not found",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.BookRequestValidation.EXPECT().
					StackedBookshelf(gomock.Any()).
					Return(nil)
				mocks.BookApplication.EXPECT().
					GetBookshelfByUserIDAndBookIDWithRelated(ctx, "user01", 1).
					Return(nil, exception.ErrNotFound.New(test.ErrMock))
			},
			args: args{
				req: &pb.StackedBookshelfRequest{
					UserId: "user01",
					BookId: 1,
				},
			},
			want: &test.Response{
				Code:    codes.NotFound,
				Message: nil,
			},
		},

		{
			name: "failed: internel error",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				bs := &book.Bookshelf{
					Book:   book1,
					Review: &book.Review{},
				}
				mocks.BookRequestValidation.EXPECT().
					StackedBookshelf(gomock.Any()).
					Return(nil)
				mocks.BookApplication.EXPECT().
					GetBookshelfByUserIDAndBookIDWithRelated(ctx, "user01", 1).
					Return(bs, nil)
				mocks.BookApplication.EXPECT().
					CreateOrUpdateBookshelf(ctx, bs).
					Return(exception.ErrInDatastore.New(test.ErrMock))
			},
			args: args{
				req: &pb.StackedBookshelfRequest{
					UserId: "user01",
					BookId: 1,
				},
			},
			want: &test.Response{
				Code:    codes.Internal,
				Message: nil,
			},
		},
	}

	for _, tt := range testCases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			mocks := test.NewMocks(ctrl)
			tt.setup(ctx, t, mocks)
			target := NewBookServer(mocks.BookRequestValidation, mocks.BookApplication)

			res, err := target.StackedBookshelf(ctx, tt.args.req)
			test.GRPC(t, tt.want, res, err)
		})
	}
}

func TestBookServer_ReleaseBookshelf(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	book1 := testBook(1)

	type args struct {
		req *pb.ReleaseBookshelfRequest
	}
	testCases := []struct {
		name  string
		setup func(context.Context, *testing.T, *test.Mocks)
		args  args
		want  *test.Response
	}{
		{
			name: "success",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				bs := &book.Bookshelf{
					Book:   book1,
					Review: &book.Review{},
				}
				mocks.BookRequestValidation.EXPECT().
					ReleaseBookshelf(gomock.Any()).
					Return(nil)
				mocks.BookApplication.EXPECT().
					GetBookshelfByUserIDAndBookIDWithRelated(ctx, "user01", 1).
					Return(bs, nil)
				mocks.BookApplication.EXPECT().
					CreateOrUpdateBookshelf(ctx, bs).
					Return(nil)
			},
			args: args{
				req: &pb.ReleaseBookshelfRequest{
					UserId: "user01",
					BookId: 1,
				},
			},
			want: &test.Response{
				Code: codes.OK,
				Message: &pb.BookshelfResponse{
					Bookshelf: &pb.Bookshelf{
						BookId:    1,
						UserId:    "user01",
						Status:    pb.BookshelfStatus_BOOKSHELF_STATUS_RELEASE,
						ReadOn:    "",
						CreatedAt: "",
						UpdatedAt: "",
					},
				},
			},
		},
		{
			name: "failed: invalid argument",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.BookRequestValidation.EXPECT().
					ReleaseBookshelf(gomock.Any()).
					Return(exception.ErrInvalidRequestValidation.New(test.ErrMock))
			},
			args: args{
				req: &pb.ReleaseBookshelfRequest{},
			},
			want: &test.Response{
				Code:    codes.InvalidArgument,
				Message: nil,
			},
		}, {
			name: "failed: not found",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.BookRequestValidation.EXPECT().
					ReleaseBookshelf(gomock.Any()).
					Return(nil)
				mocks.BookApplication.EXPECT().
					GetBookshelfByUserIDAndBookIDWithRelated(ctx, "user01", 1).
					Return(nil, exception.ErrNotFound.New(test.ErrMock))
			},
			args: args{
				req: &pb.ReleaseBookshelfRequest{
					UserId: "user01",
					BookId: 1,
				},
			},
			want: &test.Response{
				Code:    codes.NotFound,
				Message: nil,
			},
		},

		{
			name: "failed: internel error",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				bs := &book.Bookshelf{
					Book:   book1,
					Review: &book.Review{},
				}
				mocks.BookRequestValidation.EXPECT().
					ReleaseBookshelf(gomock.Any()).
					Return(nil)
				mocks.BookApplication.EXPECT().
					GetBookshelfByUserIDAndBookIDWithRelated(ctx, "user01", 1).
					Return(bs, nil)
				mocks.BookApplication.EXPECT().
					CreateOrUpdateBookshelf(ctx, bs).
					Return(exception.ErrInDatastore.New(test.ErrMock))
			},
			args: args{
				req: &pb.ReleaseBookshelfRequest{
					UserId: "user01",
					BookId: 1,
				},
			},
			want: &test.Response{
				Code:    codes.Internal,
				Message: nil,
			},
		},
	}

	for _, tt := range testCases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			mocks := test.NewMocks(ctrl)
			tt.setup(ctx, t, mocks)
			target := NewBookServer(mocks.BookRequestValidation, mocks.BookApplication)

			res, err := target.ReleaseBookshelf(ctx, tt.args.req)
			test.GRPC(t, tt.want, res, err)
		})
	}
}

func TestBookServer_WantBookshelf(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	book1 := testBook(1)

	type args struct {
		req *pb.WantBookshelfRequest
	}
	testCases := []struct {
		name  string
		setup func(context.Context, *testing.T, *test.Mocks)
		args  args
		want  *test.Response
	}{
		{
			name: "success",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				bs := &book.Bookshelf{
					Book:   book1,
					Review: &book.Review{},
				}
				mocks.BookRequestValidation.EXPECT().
					WantBookshelf(gomock.Any()).
					Return(nil)
				mocks.BookApplication.EXPECT().
					GetBookshelfByUserIDAndBookIDWithRelated(ctx, "user01", 1).
					Return(bs, nil)
				mocks.BookApplication.EXPECT().
					CreateOrUpdateBookshelf(ctx, bs).
					Return(nil)
			},
			args: args{
				req: &pb.WantBookshelfRequest{
					UserId: "user01",
					BookId: 1,
				},
			},
			want: &test.Response{
				Code: codes.OK,
				Message: &pb.BookshelfResponse{
					Bookshelf: &pb.Bookshelf{
						BookId:    1,
						UserId:    "user01",
						Status:    pb.BookshelfStatus_BOOKSHELF_STATUS_WANT,
						ReadOn:    "",
						CreatedAt: "",
						UpdatedAt: "",
					},
				},
			},
		},
		{
			name: "failed: invalid argument",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.BookRequestValidation.EXPECT().
					WantBookshelf(gomock.Any()).
					Return(exception.ErrInvalidRequestValidation.New(test.ErrMock))
			},
			args: args{
				req: &pb.WantBookshelfRequest{},
			},
			want: &test.Response{
				Code:    codes.InvalidArgument,
				Message: nil,
			},
		}, {
			name: "failed: not found",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.BookRequestValidation.EXPECT().
					WantBookshelf(gomock.Any()).
					Return(nil)
				mocks.BookApplication.EXPECT().
					GetBookshelfByUserIDAndBookIDWithRelated(ctx, "user01", 1).
					Return(nil, exception.ErrNotFound.New(test.ErrMock))
			},
			args: args{
				req: &pb.WantBookshelfRequest{
					UserId: "user01",
					BookId: 1,
				},
			},
			want: &test.Response{
				Code:    codes.NotFound,
				Message: nil,
			},
		},

		{
			name: "failed: internel error",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				bs := &book.Bookshelf{
					Book:   book1,
					Review: &book.Review{},
				}
				mocks.BookRequestValidation.EXPECT().
					WantBookshelf(gomock.Any()).
					Return(nil)
				mocks.BookApplication.EXPECT().
					GetBookshelfByUserIDAndBookIDWithRelated(ctx, "user01", 1).
					Return(bs, nil)
				mocks.BookApplication.EXPECT().
					CreateOrUpdateBookshelf(ctx, bs).
					Return(exception.ErrInDatastore.New(test.ErrMock))
			},
			args: args{
				req: &pb.WantBookshelfRequest{
					UserId: "user01",
					BookId: 1,
				},
			},
			want: &test.Response{
				Code:    codes.Internal,
				Message: nil,
			},
		},
	}

	for _, tt := range testCases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			mocks := test.NewMocks(ctrl)
			tt.setup(ctx, t, mocks)
			target := NewBookServer(mocks.BookRequestValidation, mocks.BookApplication)

			res, err := target.WantBookshelf(ctx, tt.args.req)
			test.GRPC(t, tt.want, res, err)
		})
	}
}

func TestBookServer_DeleteBook(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	book1 := testBook(1)

	type args struct {
		req *pb.DeleteBookRequest
	}
	testCases := []struct {
		name  string
		setup func(context.Context, *testing.T, *test.Mocks)
		args  args
		want  *test.Response
	}{
		{
			name: "success",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.BookRequestValidation.EXPECT().
					DeleteBook(gomock.Any()).
					Return(nil)
				mocks.BookApplication.EXPECT().
					Get(ctx, 1).
					Return(book1, nil)
				mocks.BookApplication.EXPECT().
					Delete(ctx, book1).
					Return(nil)
			},
			args: args{
				req: &pb.DeleteBookRequest{
					BookId: 1,
				},
			},
			want: &test.Response{
				Code:    codes.OK,
				Message: &pb.Empty{},
			},
		},
		{
			name: "failed: invalid argument",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.BookRequestValidation.EXPECT().
					DeleteBook(gomock.Any()).
					Return(exception.ErrInvalidRequestValidation.New(test.ErrMock))
			},
			args: args{
				req: &pb.DeleteBookRequest{},
			},
			want: &test.Response{
				Code:    codes.InvalidArgument,
				Message: nil,
			},
		}, {
			name: "failed: not found",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.BookRequestValidation.EXPECT().
					DeleteBook(gomock.Any()).
					Return(nil)
				mocks.BookApplication.EXPECT().
					Get(ctx, 1).
					Return(nil, exception.ErrNotFound.New(test.ErrMock))
			},
			args: args{
				req: &pb.DeleteBookRequest{
					BookId: 1,
				},
			},
			want: &test.Response{
				Code:    codes.NotFound,
				Message: nil,
			},
		},

		{
			name: "failed: internel error",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.BookRequestValidation.EXPECT().
					DeleteBook(gomock.Any()).
					Return(nil)
				mocks.BookApplication.EXPECT().
					Get(ctx, 1).
					Return(book1, nil)
				mocks.BookApplication.EXPECT().
					Delete(ctx, book1).
					Return(exception.ErrInDatastore.New(test.ErrMock))
			},
			args: args{
				req: &pb.DeleteBookRequest{
					BookId: 1,
				},
			},
			want: &test.Response{
				Code:    codes.Internal,
				Message: nil,
			},
		},
	}

	for _, tt := range testCases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			mocks := test.NewMocks(ctrl)
			tt.setup(ctx, t, mocks)
			target := NewBookServer(mocks.BookRequestValidation, mocks.BookApplication)

			res, err := target.DeleteBook(ctx, tt.args.req)
			test.GRPC(t, tt.want, res, err)
		})
	}
}

func TestBookServer_DeleteBookshelf(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	book1 := testBook(1)
	bookshelf1 := testBookshelf(1, book1.ID, "user01")

	type args struct {
		req *pb.DeleteBookshelfRequest
	}
	testCases := []struct {
		name  string
		setup func(context.Context, *testing.T, *test.Mocks)
		args  args
		want  *test.Response
	}{
		{
			name: "success",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.BookRequestValidation.EXPECT().
					DeleteBookshelf(gomock.Any()).
					Return(nil)
				mocks.BookApplication.EXPECT().
					GetBookshelfByUserIDAndBookID(ctx, "user01", 1).
					Return(bookshelf1, nil)
				mocks.BookApplication.EXPECT().
					DeleteBookshelf(ctx, bookshelf1).
					Return(nil)
			},
			args: args{
				req: &pb.DeleteBookshelfRequest{
					UserId: "user01",
					BookId: 1,
				},
			},
			want: &test.Response{
				Code:    codes.OK,
				Message: &pb.Empty{},
			},
		},
		{
			name: "failed: invalid argument",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.BookRequestValidation.EXPECT().
					DeleteBookshelf(gomock.Any()).
					Return(exception.ErrInvalidRequestValidation.New(test.ErrMock))
			},
			args: args{
				req: &pb.DeleteBookshelfRequest{},
			},
			want: &test.Response{
				Code:    codes.InvalidArgument,
				Message: nil,
			},
		}, {
			name: "failed: not found",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.BookRequestValidation.EXPECT().
					DeleteBookshelf(gomock.Any()).
					Return(nil)
				mocks.BookApplication.EXPECT().
					GetBookshelfByUserIDAndBookID(ctx, "user01", 1).
					Return(nil, exception.ErrNotFound.New(test.ErrMock))
			},
			args: args{
				req: &pb.DeleteBookshelfRequest{
					UserId: "user01",
					BookId: 1,
				},
			},
			want: &test.Response{
				Code:    codes.NotFound,
				Message: nil,
			},
		},

		{
			name: "failed: internel error",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.BookRequestValidation.EXPECT().
					DeleteBookshelf(gomock.Any()).
					Return(nil)
				mocks.BookApplication.EXPECT().
					GetBookshelfByUserIDAndBookID(ctx, "user01", 1).
					Return(bookshelf1, nil)
				mocks.BookApplication.EXPECT().
					DeleteBookshelf(ctx, bookshelf1).
					Return(exception.ErrInDatastore.New(test.ErrMock))
			},
			args: args{
				req: &pb.DeleteBookshelfRequest{
					UserId: "user01",
					BookId: 1,
				},
			},
			want: &test.Response{
				Code:    codes.Internal,
				Message: nil,
			},
		},
	}

	for _, tt := range testCases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			mocks := test.NewMocks(ctrl)
			tt.setup(ctx, t, mocks)
			target := NewBookServer(mocks.BookRequestValidation, mocks.BookApplication)

			res, err := target.DeleteBookshelf(ctx, tt.args.req)
			test.GRPC(t, tt.want, res, err)
		})
	}
}

func testBook(id int) *book.Book {
	return &book.Book{
		ID:             id,
		Title:          "小説　ちはやふる　上の句",
		TitleKana:      "ショウセツ チハヤフルカミノク",
		Description:    "綾瀬千早は高校入学と同時に、競技かるた部を作ろうと奔走する。幼馴染の太一と仲間を集め、夏の全国大会に出場するためだ。強くなって、新と再会したい。幼い頃かるたを取り合った、新に寄せる千早の秘めた想いに気づきながらも、太一は千早を守り立てる。それぞれの青春を懸けた、一途な情熱の物語が幕開ける。",
		Isbn:           "9784062938426",
		Publisher:      "講談社",
		PublishedOn:    "2018-01-16",
		ThumbnailURL:   "https://thumbnail.image.rakuten.co.jp/@0_mall/book/cabinet/8426/9784062938426.jpg?_ex=120x120",
		RakutenURL:     "https://books.rakuten.co.jp/rb/15271426/",
		RakutenSize:    "コミック",
		RakutenGenreID: "001004008001/001004008003/001019001",
		CreatedAt:      test.TimeMock,
		UpdatedAt:      test.TimeMock,
		Authors: []*book.Author{
			{
				ID:        1,
				Name:      "有沢 ゆう希",
				NameKana:  "アリサワ ユウキ",
				CreatedAt: test.TimeMock,
				UpdatedAt: test.TimeMock,
			},
			{
				ID:        2,
				Name:      "末次 由紀",
				NameKana:  "スエツグ ユキ",
				CreatedAt: test.TimeMock,
				UpdatedAt: test.TimeMock,
			},
		},
	}
}

func testBookshelf(id int, bookID int, userID string) *book.Bookshelf {
	return &book.Bookshelf{
		ID:        id,
		BookID:    bookID,
		UserID:    userID,
		ReviewID:  0,
		Status:    book.ReadingStatus,
		CreatedAt: test.TimeMock,
		UpdatedAt: test.TimeMock,
	}
}

func testReview(id int, bookID int, userID string) *book.Review {
	return &book.Review{
		ID:         id,
		BookID:     bookID,
		UserID:     userID,
		Score:      3,
		Impression: "テストレビューです",
		CreatedAt:  test.TimeMock,
		UpdatedAt:  test.TimeMock,
	}
}
