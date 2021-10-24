package entity

import (
	"time"

	"github.com/calmato/gran-book/api/service/internal/gateway/entity"
)

type MonthlyResult struct {
	Year      int32 `json:"year"`      // 年
	Month     int32 `json:"month"`     // 月
	ReadTotal int64 `json:"readTotal"` // 読んだ本の合計
}

type MonthlyResults []*MonthlyResult

func NewMonthlyResult(r *entity.MonthlyResult, date time.Time) *MonthlyResult {
	var total int64
	if r != nil {
		total = r.ReadTotal
	}

	return &MonthlyResult{
		Year:      int32(date.Year()),
		Month:     int32(date.Month()),
		ReadTotal: total,
	}
}

func NewMonthlyResults(rm map[string]*entity.MonthlyResult, now time.Time) MonthlyResults {
	const monthPerYear = 12 // 12ヶ月分
	res := make(MonthlyResults, monthPerYear)
	for i := 0; i < 12; i++ {
		date := now.AddDate(0, -i, 0)
		key := entity.MonthlyResultKey(int32(date.Year()), int32(date.Month()))
		res[i] = NewMonthlyResult(rm[key], date)
	}
	return res
}
