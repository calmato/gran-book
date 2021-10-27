package entity

import (
	"testing"
	"time"

	"github.com/calmato/gran-book/api/internal/gateway/entity"
	"github.com/calmato/gran-book/api/pkg/test"
	"github.com/calmato/gran-book/api/proto/book"
	"github.com/stretchr/testify/assert"
)

func TestMonthlyResults(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		results map[string]*entity.MonthlyResult
		now     time.Time
		expect  MonthlyResults
	}{
		{
			name: "success",
			results: map[string]*entity.MonthlyResult{
				"2021-07": {
					MonthlyResult: &book.MonthlyResult{
						Year:      2021,
						Month:     7,
						ReadTotal: 3,
					},
				},
				"2021-05": {
					MonthlyResult: &book.MonthlyResult{
						Year:      2021,
						Month:     5,
						ReadTotal: 8,
					},
				},
			},
			now: test.Now(),
			expect: MonthlyResults{
				{Year: 2021, Month: 7, ReadTotal: 3},
				{Year: 2021, Month: 6, ReadTotal: 0},
				{Year: 2021, Month: 5, ReadTotal: 8},
				{Year: 2021, Month: 4, ReadTotal: 0},
				{Year: 2021, Month: 3, ReadTotal: 0},
				{Year: 2021, Month: 2, ReadTotal: 0},
				{Year: 2021, Month: 1, ReadTotal: 0},
				{Year: 2020, Month: 12, ReadTotal: 0},
				{Year: 2020, Month: 11, ReadTotal: 0},
				{Year: 2020, Month: 10, ReadTotal: 0},
				{Year: 2020, Month: 9, ReadTotal: 0},
				{Year: 2020, Month: 8, ReadTotal: 0},
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			actual := NewMonthlyResults(tt.results, tt.now)
			assert.Equal(t, tt.expect, actual)
		})
	}
}
