package v1

import (
	"testing"
	"time"

	"github.com/calmato/gran-book/api/gateway/native/internal/entity"
	"github.com/calmato/gran-book/api/gateway/native/pkg/datetime"
	"github.com/calmato/gran-book/api/gateway/native/pkg/test"
	"github.com/calmato/gran-book/api/gateway/native/proto/service/book"
	"github.com/stretchr/testify/assert"
)

func TestUserTopResponse(t *testing.T) {
	t.Parallel()
	now, _ := datetime.ParseTime(test.TimeMock)
	type args struct {
		results map[string]*entity.MonthlyResult
		now     time.Time
	}
	tests := []struct {
		name   string
		args   args
		expect *UserTopResponse
	}{
		{
			name: "success",
			args: args{
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
				now: now,
			},
			expect: &UserTopResponse{
				MonthlyResults: []*userTopMonthlyResult{
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
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			actual := NewUserTopResponse(tt.args.results, tt.args.now)
			assert.Equal(t, tt.expect, actual)
		})
	}
}