package entity

import (
	"testing"

	"github.com/calmato/gran-book/api/gateway/native/pkg/test"
	pb "github.com/calmato/gran-book/api/gateway/native/proto"
	"github.com/stretchr/testify/assert"
)

func TestReviews(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name          string
		reviews       Reviews
		expectUserIDs []string
		expectBookIDs []int64
	}{
		{
			name: "success",
			reviews: Reviews{
				{
					Review: &pb.Review{
						Id:         1,
						BookId:     1,
						UserId:     "00000000-0000-0000-0000-000000000000",
						Score:      3,
						Impression: "テストレビューです",
						CreatedAt:  test.TimeMock,
						UpdatedAt:  test.TimeMock,
					},
				},
				{
					Review: &pb.Review{
						Id:         2,
						BookId:     2,
						UserId:     "00000000-0000-0000-0000-000000000000",
						Score:      3,
						Impression: "テストレビューです",
						CreatedAt:  test.TimeMock,
						UpdatedAt:  test.TimeMock,
					},
				},
			},
			expectUserIDs: []string{"00000000-0000-0000-0000-000000000000"},
			expectBookIDs: []int64{1, 2},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.expectUserIDs, tt.reviews.UserIDs())
			assert.Equal(t, tt.expectBookIDs, tt.reviews.BookIDs())
		})
	}
}
