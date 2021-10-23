package v1

import (
	"time"

	"github.com/calmato/gran-book/api/service/internal/gateway/entity"
)

// ユーザーのトップページ表示用の情報
type UserTopResponse struct {
	MonthlyResults []*userTopMonthlyResult `json:"monthlyResultsList"` // 月毎の読書実績一覧
}

func NewUserTopResponse(rm map[string]*entity.MonthlyResult, now time.Time) *UserTopResponse {
	return &UserTopResponse{
		MonthlyResults: newUserTopMonthlyResults(rm, now),
	}
}

type userTopMonthlyResult struct {
	Year      int32 `json:"year"`      // 年
	Month     int32 `json:"month"`     // 月
	ReadTotal int64 `json:"readTotal"` // 読んだ本の合計
}

func newUserTopMonthlyResult(r *entity.MonthlyResult, date time.Time) *userTopMonthlyResult {
	var total int64
	if r != nil {
		total = r.ReadTotal
	}

	return &userTopMonthlyResult{
		Year:      int32(date.Year()),
		Month:     int32(date.Month()),
		ReadTotal: total,
	}
}

func newUserTopMonthlyResults(rm map[string]*entity.MonthlyResult, now time.Time) []*userTopMonthlyResult {
	res := make([]*userTopMonthlyResult, 12) // 12ヶ月分
	for i := 0; i < 12; i++ {
		date := now.AddDate(0, -i, 0)
		key := entity.MonthlyResultKey(int32(date.Year()), int32(date.Month()))

		r := rm[key]
		res[i] = newUserTopMonthlyResult(r, date)
	}
	return res
}
