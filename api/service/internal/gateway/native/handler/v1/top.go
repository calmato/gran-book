package v1

import (
	"net/http"

	gentity "github.com/calmato/gran-book/api/service/internal/gateway/entity"
	"github.com/calmato/gran-book/api/service/internal/gateway/native/entity"
	response "github.com/calmato/gran-book/api/service/internal/gateway/native/response/v1"
	"github.com/calmato/gran-book/api/service/internal/gateway/util"
	"github.com/calmato/gran-book/api/service/pkg/datetime"
	"github.com/calmato/gran-book/api/service/proto/book"
	"github.com/gin-gonic/gin"
)

// getTopUser - ユーザーのトップページ表示用の情報取得
func (h *apiV1Handler) getTopUser(ctx *gin.Context) {
	c := util.SetMetadata(ctx)

	a, err := h.authGetAuth(c)
	if err != nil {
		util.ErrorHandling(ctx, err)
		return
	}

	now := h.now()
	since := datetime.BeginningOfMonth(now).AddDate(-1, 1, 0) // 11ヶ月前の初日
	until := datetime.EndOfMonth(now)                         // 今月末

	in := &book.ListUserMonthlyResultRequest{
		UserId:    a.Id,
		SinceDate: datetime.FormatDate(since),
		UntilDate: datetime.FormatDate(until),
	}
	out, err := h.Book.ListUserMonthlyResult(c, in)
	if err != nil {
		util.ErrorHandling(ctx, err)
		return
	}
	rs := gentity.NewMonthlyResults(out.MonthlyResults)

	res := &response.UserTopResponse{
		MonthlyResults: entity.NewMonthlyResults(rs.Map(), now),
	}
	ctx.JSON(http.StatusOK, res)
}
