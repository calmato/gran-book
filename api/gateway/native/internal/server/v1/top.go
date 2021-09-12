package v1

import (
	"net/http"

	"github.com/calmato/gran-book/api/gateway/native/internal/entity"
	response "github.com/calmato/gran-book/api/gateway/native/internal/response/v1"
	"github.com/calmato/gran-book/api/gateway/native/internal/server/util"
	"github.com/calmato/gran-book/api/gateway/native/pkg/datetime"
	"github.com/calmato/gran-book/api/gateway/native/proto/service/book"
	"github.com/calmato/gran-book/api/gateway/native/proto/service/user"
	"github.com/gin-gonic/gin"
)

type TopHandler interface {
	UserTop(ctx *gin.Context)
}

type topHandler struct {
	authClient user.AuthServiceClient
	bookClient book.BookServiceClient
}

func NewTopHandler(ac user.AuthServiceClient, bc book.BookServiceClient) TopHandler {
	return &topHandler{
		authClient: ac,
		bookClient: bc,
	}
}

// UserTop - ユーザーのトップページ表示用の情報取得
func (h *topHandler) UserTop(ctx *gin.Context) {
	c := util.SetMetadata(ctx)
	authOutput, err := h.authClient.GetAuth(c, &user.Empty{})
	if err != nil {
		util.ErrorHandling(ctx, err)
		return
	}

	u := entity.NewAuth(authOutput.Auth)

	t := datetime.Now()
	since := datetime.BeginningOfMonth(t).AddDate(-1, 1, 0) // 11ヶ月前の初日
	until := datetime.EndOfMonth(t)                         // 今月末

	resultsInput := &book.ListUserMonthlyResultRequest{
		UserId:    u.Id,
		SinceDate: datetime.FormatDate(since),
		UntilDate: datetime.FormatDate(until),
	}

	resultsOutput, err := h.bookClient.ListUserMonthlyResult(c, resultsInput)
	if err != nil {
		util.ErrorHandling(ctx, err)
		return
	}

	rs := entity.NewMonthlyResults(resultsOutput.MonthlyResults)

	res := h.getUserTopResponse(rs)
	ctx.JSON(http.StatusOK, res)
}

func (h *topHandler) getUserTopResponse(rs entity.MonthlyResults) *response.UserTopResponse {
	results := make([]*response.UserTopMonthlyResult, len(rs))
	for i, r := range rs {
		result := &response.UserTopMonthlyResult{
			Year:      r.Year,
			Month:     r.Month,
			ReadTotal: r.ReadTotal,
		}

		results[i] = result
	}

	return &response.UserTopResponse{
		MonthlyResults: results,
	}
}
