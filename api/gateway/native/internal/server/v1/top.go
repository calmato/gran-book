package v1

import (
	"net/http"
	"time"

	response "github.com/calmato/gran-book/api/gateway/native/internal/response/v1"
	"github.com/gin-gonic/gin"
)

type TopHandler interface {
	UserTop(ctx *gin.Context)
}

type topHandler struct{}

func NewTopHandler() TopHandler {
	return &topHandler{}
}

// UserTop - ユーザーのトップページ表示用の情報取得
func (h *topHandler) UserTop(ctx *gin.Context) {
	// Mock
	res := h.getUserTopResponse()
	ctx.JSON(http.StatusOK, res)
}

func (h *topHandler) getUserTopResponse() *response.UserTopResponse {
	now := time.Now().Local()

	results := []*response.UserTopResponse_MonthlyResult{
		{
			Year:      int32(now.Year()),
			Month:     int32(now.Month()),
			ReadTotal: 10,
		},
		{
			Year:      int32(now.AddDate(0, -1, 0).Year()),
			Month:     int32(now.AddDate(0, -1, 0).Month()),
			ReadTotal: 2,
		},
		{
			Year:      int32(now.AddDate(0, -2, 0).Year()),
			Month:     int32(now.AddDate(0, -2, 0).Month()),
			ReadTotal: 5,
		},
	}

	return &response.UserTopResponse{
		MonthlyResults: results,
	}
}
