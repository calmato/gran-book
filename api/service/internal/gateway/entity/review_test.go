package entity

import (
	"testing"

	"github.com/calmato/gran-book/api/service/pkg/datetime"
	"github.com/calmato/gran-book/api/service/pkg/test"
	"github.com/calmato/gran-book/api/service/proto/book"
	"github.com/stretchr/testify/assert"
)

func TestReview(t *testing.T) {
	t.Parallel()
	now := datetime.FormatTime(test.Now())
	tests := []struct {
		name   string
		review *book.Review
		expect *Review
	}{
		{
			name: "success",
			review: &book.Review{
				Id:         1,
				BookId:     1,
				UserId:     "00000000-0000-0000-0000-000000000000",
				Score:      3,
				Impression: "テストレビューです",
				CreatedAt:  now,
				UpdatedAt:  now,
			},
			expect: &Review{
				Review: &book.Review{
					Id:         1,
					BookId:     1,
					UserId:     "00000000-0000-0000-0000-000000000000",
					Score:      3,
					Impression: "テストレビューです",
					CreatedAt:  now,
					UpdatedAt:  now,
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
	now := datetime.FormatTime(test.Now())
	tests := []struct {
		name    string
		reviews []*book.Review
		expect  Reviews
	}{
		{
			name: "success",
			reviews: []*book.Review{
				{
					Id:         1,
					BookId:     1,
					UserId:     "00000000-0000-0000-0000-000000000000",
					Score:      3,
					Impression: "テストレビューです",
					CreatedAt:  now,
					UpdatedAt:  now,
				},
				{
					Id:         2,
					BookId:     2,
					UserId:     "00000000-0000-0000-0000-000000000000",
					Score:      3,
					Impression: "テストレビューです",
					CreatedAt:  now,
					UpdatedAt:  now,
				},
			},
			expect: Reviews{
				{
					Review: &book.Review{
						Id:         1,
						BookId:     1,
						UserId:     "00000000-0000-0000-0000-000000000000",
						Score:      3,
						Impression: "テストレビューです",
						CreatedAt:  now,
						UpdatedAt:  now,
					},
				},
				{
					Review: &book.Review{
						Id:         2,
						BookId:     2,
						UserId:     "00000000-0000-0000-0000-000000000000",
						Score:      3,
						Impression: "テストレビューです",
						CreatedAt:  now,
						UpdatedAt:  now,
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
	now := datetime.FormatTime(test.Now())
	tests := []struct {
		name    string
		reviews []*book.Review
		expect  []string
	}{
		{
			name: "success",
			reviews: []*book.Review{
				{
					Id:         1,
					BookId:     1,
					UserId:     "00000000-0000-0000-0000-000000000000",
					Score:      3,
					Impression: "テストレビューです",
					CreatedAt:  now,
					UpdatedAt:  now,
				},
				{
					Id:         2,
					BookId:     2,
					UserId:     "00000000-0000-0000-0000-000000000000",
					Score:      3,
					Impression: "テストレビューです",
					CreatedAt:  now,
					UpdatedAt:  now,
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
	now := datetime.FormatTime(test.Now())
	tests := []struct {
		name    string
		reviews []*book.Review
		expect  []int64
	}{
		{
			name: "success",
			reviews: []*book.Review{
				{
					Id:         1,
					BookId:     1,
					UserId:     "00000000-0000-0000-0000-000000000000",
					Score:      3,
					Impression: "テストレビューです",
					CreatedAt:  now,
					UpdatedAt:  now,
				},
				{
					Id:         2,
					BookId:     2,
					UserId:     "00000000-0000-0000-0000-000000000000",
					Score:      3,
					Impression: "テストレビューです",
					CreatedAt:  now,
					UpdatedAt:  now,
				},
			},
			expect: []int64{1, 2},
		},
		{
			name: "success book id is deprecated",
			reviews: []*book.Review{
				{
					Id:         1,
					BookId:     1,
					UserId:     "00000000-0000-0000-0000-000000000000",
					Score:      3,
					Impression: "テストレビューです",
					CreatedAt:  now,
					UpdatedAt:  now,
				},
				{
					Id:         2,
					BookId:     1,
					UserId:     "00000000-0000-0000-0000-000000000000",
					Score:      3,
					Impression: "テストレビューです",
					CreatedAt:  now,
					UpdatedAt:  now,
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
