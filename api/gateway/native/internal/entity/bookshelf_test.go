package entity

import (
	"testing"

	"github.com/calmato/gran-book/api/gateway/native/pkg/test"
	pb "github.com/calmato/gran-book/api/gateway/native/proto"
	"github.com/stretchr/testify/assert"
)

func TestBookshelves(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name          string
		bookshelves   Bookshelves
		expectBookIDs []int64
	}{
		{
			name: "success",
			bookshelves: Bookshelves{
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
			expectBookIDs: []int64{1, 2},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.expectBookIDs, tt.bookshelves.BookIDs())
		})
	}
}
