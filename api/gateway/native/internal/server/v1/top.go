package v1

import (
	"net/http"

	"github.com/calmato/gran-book/api/gateway/native/internal/entity"
	response "github.com/calmato/gran-book/api/gateway/native/internal/response/v1"
	"github.com/calmato/gran-book/api/gateway/native/internal/server/util"
	"github.com/calmato/gran-book/api/gateway/native/pkg/datetime"
	pb "github.com/calmato/gran-book/api/gateway/native/proto"
	"github.com/gin-gonic/gin"
)

type TopHandler interface {
	UserTop(ctx *gin.Context)
}

type topHandler struct {
	authClient pb.AuthServiceClient
	bookClient pb.BookServiceClient
}

func NewTopHandler(ac pb.AuthServiceClient, bc pb.BookServiceClient) TopHandler {
	return &topHandler{
		authClient: ac,
		bookClient: bc,
	}
}

// UserTop - ユーザーのトップページ表示用の情報取得
func (h *topHandler) UserTop(ctx *gin.Context) {
	now := datetime.Now()

	authOutput, err := h.authClient.GetAuth(ctx, &pb.Empty{})
	if err != nil {
		util.ErrorHandling(ctx, err)
		return
	}

	u := entity.NewAuth(authOutput.Auth)

	t := datetime.New(now)
	since := t.BeginningOfMonth().AddDate(-1, 1, 0) // 11ヶ月前の初日
	until := t.EndOfMonth()                         // 今月末

	resultsInput := &pb.ListUserMonthlyResultRequest{
		UserId:    u.Id,
		SinceDate: since.String(),
		UntilDate: until.String(),
	}

	resultsOutput, err := h.bookClient.ListUserMonthlyResult(ctx, resultsInput)
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
