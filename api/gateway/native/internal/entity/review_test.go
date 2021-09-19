package entity

import (
	"testing"

	"github.com/calmato/gran-book/api/gateway/native/pkg/test"
	pb "github.com/calmato/gran-book/api/gateway/native/proto/service/book"
	"github.com/stretchr/testify/assert"
)

func TestReview(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name   string
		review *pb.Review
		expect *Review
	}{
		{
			name: "success",
			review: &pb.Review{
				Id:         1,
				BookId:     1,
				UserId:     "00000000-0000-0000-0000-000000000000",
				Score:      3,
				Impression: "テストレビューです",
				CreatedAt:  test.TimeMock,
				UpdatedAt:  test.TimeMock,
			},
			expect: &Review{
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
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			actual := NewReview(tt.review)
			assert.Equal(t, tt.expect, actual)
		})
	}
}

func TestReviews(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		reviews []*pb.Review
		expect  Reviews
	}{
		{
			name: "success",
			reviews: []*pb.Review{
				{
					Id:         1,
					BookId:     1,
					UserId:     "00000000-0000-0000-0000-000000000000",
					Score:      3,
					Impression: "テストレビューです",
					CreatedAt:  test.TimeMock,
					UpdatedAt:  test.TimeMock,
				},
				{
					Id:         2,
					BookId:     2,
					UserId:     "00000000-0000-0000-0000-000000000000",
					Score:      3,
					Impression: "テストレビューです",
					CreatedAt:  test.TimeMock,
					UpdatedAt:  test.TimeMock,
				},
			},
			expect: Reviews{
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
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			actual := NewReviews(tt.reviews)
			assert.Equal(t, tt.expect, actual)
		})
	}
}

func TestReviews_UserIDs(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		reviews []*pb.Review
		expect  []string
	}{
		{
			name: "success",
			reviews: []*pb.Review{
				{
					Id:         1,
					BookId:     1,
					UserId:     "00000000-0000-0000-0000-000000000000",
					Score:      3,
					Impression: "テストレビューです",
					CreatedAt:  test.TimeMock,
					UpdatedAt:  test.TimeMock,
				},
				{
					Id:         2,
					BookId:     2,
					UserId:     "00000000-0000-0000-0000-000000000000",
					Score:      3,
					Impression: "テストレビューです",
					CreatedAt:  test.TimeMock,
					UpdatedAt:  test.TimeMock,
				},
			},
			expect: []string{"00000000-0000-0000-0000-000000000000"},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			actual := NewReviews(tt.reviews)
			assert.Equal(t, tt.expect, actual.UserIDs())
		})
	}
}

func TestReviews_BookIDs(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		reviews []*pb.Review
		expect  []int64
	}{
		{
			name: "success",
			reviews: []*pb.Review{
				{
					Id:         1,
					BookId:     1,
					UserId:     "00000000-0000-0000-0000-000000000000",
					Score:      3,
					Impression: "テストレビューです",
					CreatedAt:  test.TimeMock,
					UpdatedAt:  test.TimeMock,
				},
				{
					Id:         2,
					BookId:     2,
					UserId:     "00000000-0000-0000-0000-000000000000",
					Score:      3,
					Impression: "テストレビューです",
					CreatedAt:  test.TimeMock,
					UpdatedAt:  test.TimeMock,
				},
			},
			expect: []int64{1, 2},
		},
		{
			name: "success book id is deprecated",
			reviews: []*pb.Review{
				{
					Id:         1,
					BookId:     1,
					UserId:     "00000000-0000-0000-0000-000000000000",
					Score:      3,
					Impression: "テストレビューです",
					CreatedAt:  test.TimeMock,
					UpdatedAt:  test.TimeMock,
				},
				{
					Id:         2,
					BookId:     1,
					UserId:     "00000000-0000-0000-0000-000000000000",
					Score:      3,
					Impression: "テストレビューです",
					CreatedAt:  test.TimeMock,
					UpdatedAt:  test.TimeMock,
				},
			},
			expect: []int64{1},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			actual := NewReviews(tt.reviews)
			assert.Equal(t, tt.expect, actual.BookIDs())
		})
	}
}
