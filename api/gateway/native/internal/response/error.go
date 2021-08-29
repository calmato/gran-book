package response

import (
	"github.com/calmato/gran-book/api/gateway/native/internal/entity"
)

// ErrorResponse - エラーのレスポンス
type ErrorResponse struct {
	StatusCode int              `json:"status"`  // ステータスコード
	ErrorCode  entity.ErrorCode `json:"code"`    // エラーコード
	Message    string           `json:"message"` // エラー概要
	Detail     string           `json:"detail"`  // エラー詳細
}
