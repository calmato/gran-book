package entity

import (
	"testing"

	"github.com/calmato/gran-book/api/gateway/native/pkg/test"
	pb "github.com/calmato/gran-book/api/gateway/native/proto/service/book"
	"github.com/stretchr/testify/assert"
)

func TestBookshelf(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name      string
		bookshelf *pb.Bookshelf
		expect    *Bookshelf
	}{
		{
			name: "success",
			bookshelf: &pb.Bookshelf{
				Id:        1,
				BookId:    1,
				UserId:    "00000000-0000-0000-0000-000000000000",
				ReviewId:  0,
				Status:    pb.BookshelfStatus_BOOKSHELF_STATUS_READING,
				ReadOn:    test.DateMock,
				CreatedAt: test.TimeMock,
				UpdatedAt: test.TimeMock,
			},
			expect: &Bookshelf{
				Bookshelf: &pb.Bookshelf{
					Id:        1,
					BookId:    1,
					UserId:    "00000000-0000-0000-0000-000000000000",
					ReviewId:  0,
					Status:    pb.BookshelfStatus_BOOKSHELF_STATUS_READING,
					ReadOn:    test.DateMock,
					CreatedAt: test.TimeMock,
					UpdatedAt: test.TimeMock,
				},
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			actual := NewBookshelf(tt.bookshelf)
			assert.Equal(t, tt.expect, actual)
		})
	}
}

func TestBookshelf_Status(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name      string
		bookshelf *pb.Bookshelf
		expect    BookshelfStatus
	}{
		{
			name: "success",
			bookshelf: &pb.Bookshelf{
				Id:        1,
				BookId:    1,
				UserId:    "00000000-0000-0000-0000-000000000000",
				ReviewId:  0,
				Status:    pb.BookshelfStatus_BOOKSHELF_STATUS_READING,
				ReadOn:    test.DateMock,
				CreatedAt: test.TimeMock,
				UpdatedAt: test.TimeMock,
			},
			expect: BookshelfStatusReading,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			actual := NewBookshelf(tt.bookshelf)
			assert.Equal(t, tt.expect, actual.Status())
		})
	}
}

func TestBookshelves(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name        string
		bookshelves []*pb.Bookshelf
		expect      Bookshelves
	}{
		{
			name: "success",
			bookshelves: []*pb.Bookshelf{
				{
					Id:        1,
					BookId:    1,
					UserId:    "00000000-0000-0000-0000-000000000000",
					ReviewId:  0,
					Status:    pb.BookshelfStatus_BOOKSHELF_STATUS_READING,
					ReadOn:    test.DateMock,
					CreatedAt: test.TimeMock,
					UpdatedAt: test.TimeMock,
				},
				{
					Id:        2,
					BookId:    2,
					UserId:    "00000000-0000-0000-0000-000000000000",
					ReviewId:  0,
					Status:    pb.BookshelfStatus_BOOKSHELF_STATUS_STACKED,
					ReadOn:    test.DateMock,
					CreatedAt: test.TimeMock,
					UpdatedAt: test.TimeMock,
				},
			},
			expect: Bookshelves{
				{
					Bookshelf: &pb.Bookshelf{
						Id:        1,
						BookId:    1,
						UserId:    "00000000-0000-0000-0000-000000000000",
						ReviewId:  0,
						Status:    pb.BookshelfStatus_BOOKSHELF_STATUS_READING,
						ReadOn:    test.DateMock,
						CreatedAt: test.TimeMock,
						UpdatedAt: test.TimeMock,
					},
				},
				{
					Bookshelf: &pb.Bookshelf{
						Id:        2,
						BookId:    2,
						UserId:    "00000000-0000-0000-0000-000000000000",
						ReviewId:  0,
						Status:    pb.BookshelfStatus_BOOKSHELF_STATUS_STACKED,
						ReadOn:    test.DateMock,
						CreatedAt: test.TimeMock,
						UpdatedAt: test.TimeMock,
					},
				},
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			actual := NewBookshelves(tt.bookshelves)
			assert.Equal(t, tt.expect, actual)
		})
	}
}

func TestBookshelves_BookIDs(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name        string
		bookshelves []*pb.Bookshelf
		expect      []int64
	}{
		{
			name: "success",
			bookshelves: []*pb.Bookshelf{
				{
					Id:        1,
					BookId:    1,
					UserId:    "00000000-0000-0000-0000-000000000000",
					ReviewId:  0,
					Status:    pb.BookshelfStatus_BOOKSHELF_STATUS_READING,
					ReadOn:    test.DateMock,
					CreatedAt: test.TimeMock,
					UpdatedAt: test.TimeMock,
				},
				{
					Id:        2,
					BookId:    2,
					UserId:    "00000000-0000-0000-0000-000000000000",
					ReviewId:  0,
					Status:    pb.BookshelfStatus_BOOKSHELF_STATUS_STACKED,
					ReadOn:    test.DateMock,
					CreatedAt: test.TimeMock,
					UpdatedAt: test.TimeMock,
				},
			},
			expect: []int64{1, 2},
		},
		{
			name: "success book id is deprecated",
			bookshelves: []*pb.Bookshelf{
				{
					Id:        1,
					BookId:    1,
					UserId:    "00000000-0000-0000-0000-000000000000",
					ReviewId:  0,
					Status:    pb.BookshelfStatus_BOOKSHELF_STATUS_READING,
					ReadOn:    test.DateMock,
					CreatedAt: test.TimeMock,
					UpdatedAt: test.TimeMock,
				},
				{
					Id:        2,
					BookId:    1,
					UserId:    "00000000-0000-0000-0000-000000000000",
					ReviewId:  0,
					Status:    pb.BookshelfStatus_BOOKSHELF_STATUS_STACKED,
					ReadOn:    test.DateMock,
					CreatedAt: test.TimeMock,
					UpdatedAt: test.TimeMock,
				},
			},
			expect: []int64{1},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			actual := NewBookshelves(tt.bookshelves)
			assert.Equal(t, tt.expect, actual.BookIDs())
		})
	}
}
