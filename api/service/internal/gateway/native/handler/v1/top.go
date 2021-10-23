package v1

import (
	"net/http"

	"github.com/calmato/gran-book/api/service/internal/gateway/entity"
	response "github.com/calmato/gran-book/api/service/internal/gateway/native/response/v1"
	"github.com/calmato/gran-book/api/service/internal/gateway/util"
	"github.com/calmato/gran-book/api/service/pkg/datetime"
	"github.com/calmato/gran-book/api/service/proto/book"
	"github.com/calmato/gran-book/api/service/proto/user"
	"github.com/gin-gonic/gin"
)

// getTopUser - ユーザーのトップページ表示用の情報取得
func (h *apiV1Handler) getTopUser(ctx *gin.Context) {
	c := util.SetMetadata(ctx)

	authOutput, err := h.Auth.GetAuth(c, &user.Empty{})
	if err != nil {
		util.ErrorHandling(ctx, err)
		return
	}

	u := entity.NewAuth(authOutput.Auth)

	now := h.now()
	since := datetime.BeginningOfMonth(now).AddDate(-1, 1, 0) // 11ヶ月前の初日
	until := datetime.EndOfMonth(now)                         // 今月末

	resultsInput := &book.ListUserMonthlyResultRequest{
		UserId:    u.Id,
		SinceDate: datetime.FormatDate(since),
		UntilDate: datetime.FormatDate(until),
	}
	resultsOutput, err := h.Book.ListUserMonthlyResult(c, resultsInput)
	if err != nil {
		util.ErrorHandling(ctx, err)
		return
	}

	rs := entity.NewMonthlyResults(resultsOutput.MonthlyResults)
	res := response.NewUserTopResponse(rs.Map(), now)
	ctx.JSON(http.StatusOK, res)
}
