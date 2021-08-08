package response

import (
	"github.com/calmato/gran-book/api/gateway/native/internal/entity"
)

// ErrorResponse - エラーのレスポンス
type ErrorResponse struct {
	StatusCode int              `json:"status"`
	ErrorCode  entity.ErrorCode `json:"code"`
	Message    string           `json:"message"`
	Detail     string           `json:"detail"`
}
