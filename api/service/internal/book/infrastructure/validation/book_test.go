package validation

import (
	"context"
	"testing"

	"github.com/calmato/gran-book/api/service/internal/book/domain/book"
	"github.com/calmato/gran-book/api/service/pkg/test"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestBookDomainValidation_Book(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type args struct {
		book *book.Book
	}
	tests := []struct {
		name   string
		setup  func(ctx context.Context, t *testing.T, mocks *test.Mocks)
		args   args
		expect bool
	}{
		{
			name: "success if id does not exists when create",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.BookRepository.EXPECT().GetBookIDByIsbn(ctx, "1234567890123").Return(0, test.ErrMock)
			},
			args: args{
				book: &book.Book{
					ID:   0,
					Isbn: "1234567890123",
				},
			},
			expect: true,
		},
		{
			name: "success if id does not exists when update",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.BookRepository.EXPECT().GetBookIDByIsbn(ctx, "1234567890123").Return(1, nil)
			},
			args: args{
				book: &book.Book{
					ID:   1,
					Isbn: "1234567890123",
				},
			},
			expect: true,
		},
		{
			name: "failed if id exists when create",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.BookRepository.EXPECT().GetBookIDByIsbn(ctx, "1234567890123").Return(1, nil)
			},
			args: args{
				book: &book.Book{
					ID:   0,
					Isbn: "1234567890123",
				},
			},
			expect: false,
		},
		{
			name: "failed if id is not match when update",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.BookRepository.EXPECT().GetBookIDByIsbn(ctx, "1234567890123").Return(2, nil)
			},
			args: args{
				book: &book.Book{
					ID:   0,
					Isbn: "1234567890123",
				},
			},
			expect: false,
		},
		{
			name:  "success if isbn is zero value",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {},
			args: args{
				book: &book.Book{Isbn: ""},
			},
			expect: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			mocks := test.NewMocks(ctrl)
			tt.setup(ctx, t, mocks)
			target := NewBookDomainValidation(mocks.BookRepository)

			err := target.Book(ctx, tt.args.book)
			assert.Equal(t, tt.expect, err == nil)
		})
	}
}

func TestBookDomainValidation_Author(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type args struct {
		author *book.Author
	}
	tests := []struct {
		name   string
		setup  func(ctx context.Context, t *testing.T, mocks *test.Mocks)
		args   args
		expect bool
	}{
		{
			name:  "success",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {},
			args: args{
				author: &book.Author{
					Name:     "テスト著者",
					NameKana: "てすとちょしゃ",
				},
			},
			expect: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			mocks := test.NewMocks(ctrl)
			tt.setup(ctx, t, mocks)
			target := NewBookDomainValidation(mocks.BookRepository)

			err := target.Author(ctx, tt.args.author)
			assert.Equal(t, tt.expect, err == nil)
		})
	}
}

func TestBookDomainValidation_Bookshelf(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type args struct {
		bookshelf *book.Bookshelf
	}
	tests := []struct {
		name   string
		setup  func(ctx context.Context, t *testing.T, mocks *test.Mocks)
		args   args
		expect bool
	}{
		{
			name: "success if id does not exists when create",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.BookRepository.EXPECT().
					GetBookshelfIDByUserIDAndBookID(ctx, "00000000-0000-0000-0000-000000000000", 1).
					Return(0, test.ErrMock)
			},
			args: args{
				bookshelf: &book.Bookshelf{
					ID:     0,
					BookID: 1,
					UserID: "00000000-0000-0000-0000-000000000000",
					Status: book.ReadingStatus,
					ReadOn: test.DateMock,
				},
			},
			expect: true,
		},
		{
			name: "success if id does not exists when update",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.BookRepository.EXPECT().
					GetBookshelfIDByUserIDAndBookID(ctx, "00000000-0000-0000-0000-000000000000", 1).
					Return(1, nil)
			},
			args: args{
				bookshelf: &book.Bookshelf{
					ID:     1,
					BookID: 1,
					UserID: "00000000-0000-0000-0000-000000000000",
					Status: book.ReadingStatus,
					ReadOn: test.DateMock,
				},
			},
			expect: true,
		},
		{
			name: "failed if id exists when create",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.BookRepository.EXPECT().
					GetBookshelfIDByUserIDAndBookID(ctx, "00000000-0000-0000-0000-000000000000", 1).
					Return(1, nil)
			},
			args: args{
				bookshelf: &book.Bookshelf{
					ID:     0,
					BookID: 1,
					UserID: "00000000-0000-0000-0000-000000000000",
					Status: book.ReadingStatus,
					ReadOn: test.DateMock,
				},
			},
			expect: false,
		},
		{
			name: "failed if id is not match when update",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.BookRepository.EXPECT().
					GetBookshelfIDByUserIDAndBookID(ctx, "00000000-0000-0000-0000-000000000000", 1).
					Return(2, nil)
			},
			args: args{
				bookshelf: &book.Bookshelf{
					ID:     1,
					BookID: 1,
					UserID: "00000000-0000-0000-0000-000000000000",
					Status: book.ReadingStatus,
					ReadOn: test.DateMock,
				},
			},
			expect: false,
		},
		{
			name:  "success if user id is zero value",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {},
			args: args{
				bookshelf: &book.Bookshelf{
					BookID: 1,
					UserID: "",
					Status: book.ReadingStatus,
					ReadOn: test.DateMock,
				},
			},
			expect: true,
		},
		{
			name:  "success if book id is zero value",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {},
			args: args{
				bookshelf: &book.Bookshelf{
					BookID: 0,
					UserID: "00000000-0000-0000-0000-000000000000",
					Status: book.ReadingStatus,
					ReadOn: test.DateMock,
				},
			},
			expect: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			mocks := test.NewMocks(ctrl)
			tt.setup(ctx, t, mocks)
			target := NewBookDomainValidation(mocks.BookRepository)

			err := target.Bookshelf(ctx, tt.args.bookshelf)
			assert.Equal(t, tt.expect, err == nil)
		})
	}
}

func TestBookDomainValidation_Review(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type args struct {
		review *book.Review
	}
	tests := []struct {
		name   string
		setup  func(ctx context.Context, t *testing.T, mocks *test.Mocks)
		args   args
		expect bool
	}{
		{
			name: "success if id does not exists when create",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.BookRepository.EXPECT().
					GetReviewIDByUserIDAndBookID(ctx, "00000000-0000-0000-0000-000000000000", 1).
					Return(0, test.ErrMock)
			},
			args: args{
				review: &book.Review{
					ID:         0,
					BookID:     1,
					UserID:     "00000000-0000-0000-0000-000000000000",
					Score:      5,
					Impression: "書籍の感想です。",
				},
			},
			expect: true,
		},
		{
			name: "success if id does not exists when update",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.BookRepository.EXPECT().
					GetReviewIDByUserIDAndBookID(ctx, "00000000-0000-0000-0000-000000000000", 1).
					Return(1, nil)
			},
			args: args{
				review: &book.Review{
					ID:         1,
					BookID:     1,
					UserID:     "00000000-0000-0000-0000-000000000000",
					Score:      5,
					Impression: "書籍の感想です。",
				},
			},
			expect: true,
		},
		{
			name: "failed if id exists when create",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.BookRepository.EXPECT().
					GetReviewIDByUserIDAndBookID(ctx, "00000000-0000-0000-0000-000000000000", 1).
					Return(1, nil)
			},
			args: args{
				review: &book.Review{
					ID:         0,
					BookID:     1,
					UserID:     "00000000-0000-0000-0000-000000000000",
					Score:      5,
					Impression: "書籍の感想です。",
				},
			},
			expect: false,
		},
		{
			name: "failed if id is not match when update",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.BookRepository.EXPECT().
					GetReviewIDByUserIDAndBookID(ctx, "00000000-0000-0000-0000-000000000000", 1).
					Return(2, nil)
			},
			args: args{
				review: &book.Review{
					ID:         1,
					BookID:     1,
					UserID:     "00000000-0000-0000-0000-000000000000",
					Score:      5,
					Impression: "書籍の感想です。",
				},
			},
			expect: false,
		},
		{
			name:  "success if user id is zero value",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {},
			args: args{
				review: &book.Review{
					ID:         0,
					BookID:     1,
					UserID:     "",
					Score:      5,
					Impression: "書籍の感想です。",
				},
			},
			expect: true,
		},
		{
			name:  "success if book id is zero value",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {},
			args: args{
				review: &book.Review{
					ID:         0,
					BookID:     0,
					UserID:     "00000000-0000-0000-0000-000000000000",
					Score:      5,
					Impression: "書籍の感想です。",
				},
			},
			expect: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			mocks := test.NewMocks(ctrl)
			tt.setup(ctx, t, mocks)
			target := NewBookDomainValidation(mocks.BookRepository)

			err := target.Review(ctx, tt.args.review)
			assert.Equal(t, tt.expect, err == nil)
		})
	}
}
