package v1

import "github.com/calmato/gran-book/api/service/internal/gateway/native/entity"

// ユーザーのトップページ表示用の情報
type UserTopResponse struct {
	MonthlyResults entity.MonthlyResults `json:"monthlyResultsList"` // 月毎の読書実績一覧
}
