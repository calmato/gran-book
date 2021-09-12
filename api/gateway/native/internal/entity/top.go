package entity

import (
	"fmt"

	pb "github.com/calmato/gran-book/api/gateway/native/proto/service/book"
)

type MonthlyResult struct {
	*pb.MonthlyResult
}

type MonthlyResults []*MonthlyResult

func MonthlyResultKey(year int32, month int32) string {
	return fmt.Sprintf("%04d-%02d", year, month)
}

func NewMonthlyResult(r *pb.MonthlyResult) *MonthlyResult {
	return &MonthlyResult{r}
}

func NewMonthlyResults(rs []*pb.MonthlyResult) MonthlyResults {
	res := make(MonthlyResults, len(rs))
	for i := range rs {
		res[i] = NewMonthlyResult(rs[i])
	}
	return res
}

func (rs MonthlyResults) Map() map[string]*MonthlyResult {
	res := make(map[string]*MonthlyResult, len(rs))
	for _, r := range rs {
		key := MonthlyResultKey(r.Year, r.Month)
		res[key] = r
	}
	return res
}
