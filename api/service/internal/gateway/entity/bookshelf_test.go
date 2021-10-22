package entity

import (
	"testing"

	"github.com/calmato/gran-book/api/service/pkg/datetime"
	"github.com/calmato/gran-book/api/service/pkg/test"
	"github.com/calmato/gran-book/api/service/proto/book"
	"github.com/stretchr/testify/assert"
)

func TestBookshelf(t *testing.T) {
	t.Parallel()
	now := datetime.FormatTime(test.Now())
	date := datetime.FormatDate(test.DateMock)
	tests := []struct {
		name      string
		bookshelf *book.Bookshelf
		expect    *Bookshelf
	}{
		{
			name: "success",
			bookshelf: &book.Bookshelf{
				Id:        1,
				BookId:    1,
				UserId:    "00000000-0000-0000-0000-000000000000",
				ReviewId:  0,
				Status:    book.BookshelfStatus_BOOKSHELF_STATUS_READING,
				ReadOn:    date,
				CreatedAt: now,
				UpdatedAt: now,
			},
			expect: &Bookshelf{
				Bookshelf: &book.Bookshelf{
					Id:        1,
					BookId:    1,
					UserId:    "00000000-0000-0000-0000-000000000000",
					ReviewId:  0,
					Status:    book.BookshelfStatus_BOOKSHELF_STATUS_READING,
					ReadOn:    date,
					CreatedAt: now,
					UpdatedAt: now,
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
	now := datetime.FormatTime(test.Now())
	date := datetime.FormatDate(test.DateMock)
	tests := []struct {
		name      string
		bookshelf *book.Bookshelf
		expect    BookshelfStatus
	}{
		{
			name: "success",
			bookshelf: &book.Bookshelf{
				Id:        1,
				BookId:    1,
				UserId:    "00000000-0000-0000-0000-000000000000",
				ReviewId:  0,
				Status:    book.BookshelfStatus_BOOKSHELF_STATUS_READING,
				ReadOn:    date,
				CreatedAt: now,
				UpdatedAt: now,
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
	now := datetime.FormatTime(test.Now())
	date := datetime.FormatDate(test.DateMock)
	tests := []struct {
		name        string
		bookshelves []*book.Bookshelf
		expect      Bookshelves
	}{
		{
			name: "success",
			bookshelves: []*book.Bookshelf{
				{
					Id:        1,
					BookId:    1,
					UserId:    "00000000-0000-0000-0000-000000000000",
					ReviewId:  0,
					Status:    book.BookshelfStatus_BOOKSHELF_STATUS_READING,
					ReadOn:    date,
					CreatedAt: now,
					UpdatedAt: now,
				},
				{
					Id:        2,
					BookId:    2,
					UserId:    "00000000-0000-0000-0000-000000000000",
					ReviewId:  0,
					Status:    book.BookshelfStatus_BOOKSHELF_STATUS_STACKED,
					ReadOn:    date,
					CreatedAt: now,
					UpdatedAt: now,
				},
			},
			expect: Bookshelves{
				{
					Bookshelf: &book.Bookshelf{
						Id:        1,
						BookId:    1,
						UserId:    "00000000-0000-0000-0000-000000000000",
						ReviewId:  0,
						Status:    book.BookshelfStatus_BOOKSHELF_STATUS_READING,
						ReadOn:    date,
						CreatedAt: now,
						UpdatedAt: now,
					},
				},
				{
					Bookshelf: &book.Bookshelf{
						Id:        2,
						BookId:    2,
						UserId:    "00000000-0000-0000-0000-000000000000",
						ReviewId:  0,
						Status:    book.BookshelfStatus_BOOKSHELF_STATUS_STACKED,
						ReadOn:    date,
						CreatedAt: now,
						UpdatedAt: now,
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
	now := datetime.FormatTime(test.Now())
	date := datetime.FormatDate(test.DateMock)
	tests := []struct {
		name        string
		bookshelves []*book.Bookshelf
		expect      []int64
	}{
		{
			name: "success",
			bookshelves: []*book.Bookshelf{
				{
					Id:        1,
					BookId:    1,
					UserId:    "00000000-0000-0000-0000-000000000000",
					ReviewId:  0,
					Status:    book.BookshelfStatus_BOOKSHELF_STATUS_READING,
					ReadOn:    date,
					CreatedAt: now,
					UpdatedAt: now,
				},
				{
					Id:        2,
					BookId:    2,
					UserId:    "00000000-0000-0000-0000-000000000000",
					ReviewId:  0,
					Status:    book.BookshelfStatus_BOOKSHELF_STATUS_STACKED,
					ReadOn:    date,
					CreatedAt: now,
					UpdatedAt: now,
				},
			},
			expect: []int64{1, 2},
		},
		{
			name: "success book id is deprecated",
			bookshelves: []*book.Bookshelf{
				{
					Id:        1,
					BookId:    1,
					UserId:    "00000000-0000-0000-0000-000000000000",
					ReviewId:  0,
					Status:    book.BookshelfStatus_BOOKSHELF_STATUS_READING,
					ReadOn:    date,
					CreatedAt: now,
					UpdatedAt: now,
				},
				{
					Id:        2,
					BookId:    1,
					UserId:    "00000000-0000-0000-0000-000000000000",
					ReviewId:  0,
					Status:    book.BookshelfStatus_BOOKSHELF_STATUS_STACKED,
					ReadOn:    date,
					CreatedAt: now,
					UpdatedAt: now,
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
