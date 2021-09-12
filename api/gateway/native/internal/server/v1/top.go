package v1

import (
	"net/http"
	"time"

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

	res := h.getUserTopResponse(rs.Map(), t)
	ctx.JSON(http.StatusOK, res)
}

func (h *topHandler) getUserTopResponse(
	rsMap map[string]*entity.MonthlyResult, now time.Time,
) *response.UserTopResponse {
	results := make([]*response.UserTopMonthlyResult, 12) // 12ヶ月分
	for i := 0; i < 12; i++ {
		date := now.AddDate(0, -i, 0)
		year := int32(date.Year())
		month := int32(date.Month())

		var total int64
		key := entity.MonthlyResultKey(year, month)
		if r, ok := rsMap[key]; ok {
			total = r.ReadTotal
		}

		result := &response.UserTopMonthlyResult{
			Year:      year,
			Month:     month,
			ReadTotal: total,
		}

		results[i] = result
	}

	return &response.UserTopResponse{
		MonthlyResults: results,
	}
}
