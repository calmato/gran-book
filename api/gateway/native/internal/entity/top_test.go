package entity

import (
	"testing"

	pb "github.com/calmato/gran-book/api/gateway/native/proto/service/book"
	"github.com/stretchr/testify/assert"
)

func TestMonthlyResultKey(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name   string
		year   int32
		month  int32
		expect string
	}{
		{
			name:   "success: 2021-08",
			year:   2021,
			month:  8,
			expect: "2021-08",
		},
		{
			name:   "success: 2021-10",
			year:   2021,
			month:  10,
			expect: "2021-10",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.expect, MonthlyResultKey(tt.year, tt.month))
		})
	}
}

func TestMonthlyResult(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name      string
		result    *pb.MonthlyResult
		expect    *MonthlyResult
		expectKey string
	}{
		{
			name: "success",
			result: &pb.MonthlyResult{
				Year:      2021,
				Month:     8,
				ReadTotal: 8,
			},
			expect: &MonthlyResult{
				MonthlyResult: &pb.MonthlyResult{
					Year:      2021,
					Month:     8,
					ReadTotal: 8,
				},
			},
			expectKey: "2021-08",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			actual := NewMonthlyResult(tt.result)
			assert.Equal(t, tt.expect, actual)
		})
	}
}

func TestMonthlyResults(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name      string
		results   []*pb.MonthlyResult
		expect    MonthlyResults
		expectMap map[string]*MonthlyResult
	}{
		{
			name: "success",
			results: []*pb.MonthlyResult{
				{
					Year:      2021,
					Month:     8,
					ReadTotal: 8,
				},
				{
					Year:      2021,
					Month:     7,
					ReadTotal: 3,
				},
			},
			expect: MonthlyResults{
				{
					MonthlyResult: &pb.MonthlyResult{
						Year:      2021,
						Month:     8,
						ReadTotal: 8,
					},
				},
				{
					MonthlyResult: &pb.MonthlyResult{
						Year:      2021,
						Month:     7,
						ReadTotal: 3,
					},
				},
			},
			expectMap: map[string]*MonthlyResult{
				"2021-07": {
					MonthlyResult: &pb.MonthlyResult{
						Year:      2021,
						Month:     7,
						ReadTotal: 3,
					},
				},
				"2021-08": {
					MonthlyResult: &pb.MonthlyResult{
						Year:      2021,
						Month:     8,
						ReadTotal: 8,
					},
				},
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			actual := NewMonthlyResults(tt.results)
			assert.Equal(t, tt.expect, actual)
			assert.Equal(t, tt.expectMap, actual.Map())
		})
	}
}
