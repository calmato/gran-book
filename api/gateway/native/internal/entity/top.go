package entity

import pb "github.com/calmato/gran-book/api/gateway/native/proto"

type MonthlyResult struct {
	*pb.MonthlyResult
}

type MonthlyResults []*MonthlyResult

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
