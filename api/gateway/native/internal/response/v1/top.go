package v1

// ユーザーのトップページ表示用の情報
type UserTopResponse struct {
	MonthlyResults []*UserTopMonthlyResult `json:"monthlyResultsList"` // 月毎の読書実績一覧
}

type UserTopMonthlyResult struct {
	Year      int32 `json:"year"`      // 年
	Month     int32 `json:"month"`     // 月
	ReadTotal int64 `json:"readTotal"` // 読んだ本の合計
}
