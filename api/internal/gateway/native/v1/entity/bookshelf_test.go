package entity

import (
	"testing"

	"github.com/calmato/gran-book/api/internal/gateway/entity"
	"github.com/calmato/gran-book/api/pkg/datetime"
	"github.com/calmato/gran-book/api/pkg/test"
	"github.com/calmato/gran-book/api/proto/book"
	"github.com/stretchr/testify/assert"
)

func TestBookshelf(t *testing.T) {
	t.Parallel()
	now := datetime.FormatTime(test.TimeMock)
	tests := []struct {
		name      string
		bookshelf *entity.Bookshelf
		expect    *Bookshelf
	}{
		{
			name: "success",
			bookshelf: &entity.Bookshelf{
				Bookshelf: &book.Bookshelf{
					Id:        1,
					BookId:    1,
					UserId:    "00000000-0000-0000-0000-000000000000",
					ReviewId:  0,
					Status:    book.BookshelfStatus_BOOKSHELF_STATUS_READING,
					ReadOn:    datetime.FormatDate(test.DateMock),
					CreatedAt: now,
					UpdatedAt: now,
				},
			},
			expect: &Bookshelf{
				ID:         1,
				Status:     entity.BookshelfStatusReading.Name(),
				ReadOn:     datetime.FormatDate(test.DateMock),
				Impression: "",
				CreatedAt:  now,
				UpdatedAt:  now,
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
