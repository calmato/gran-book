package v1

import (
	"context"
	"net/http"
	"testing"

	"github.com/calmato/gran-book/api/gateway/native/internal/entity"
	response "github.com/calmato/gran-book/api/gateway/native/internal/response/v1"
	"github.com/calmato/gran-book/api/gateway/native/pkg/datetime"
	"github.com/calmato/gran-book/api/gateway/native/pkg/test"
	"github.com/calmato/gran-book/api/gateway/native/proto/service/book"
	"github.com/calmato/gran-book/api/gateway/native/proto/service/user"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestTop_GetTopUser(t *testing.T) {
	t.Parallel()

	now, err := datetime.ParseTime(test.TimeMock)
	assert.NoError(t, err)

	auth := testAuth("00000000-0000-0000-0000-000000000000")
	results := make([]*book.MonthlyResult, 2)
	results[0] = testMonthlyResult(2021, 8)
	results[1] = testMonthlyResult(2021, 7)

	tests := []struct {
		name   string
		setup  func(ctx context.Context, t *testing.T, mocks *test.Mocks, ctrl *gomock.Controller)
		expect *test.TestResponse
	}{
		{
			name: "success",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks, ctrl *gomock.Controller) {
				mocks.AuthService.EXPECT().
					GetAuth(gomock.Any(), &user.Empty{}).
					Return(&user.AuthResponse{Auth: auth}, nil)
				mocks.BookService.EXPECT().
					ListUserMonthlyResult(gomock.Any(), &book.ListUserMonthlyResultRequest{
						UserId:    "00000000-0000-0000-0000-000000000000",
						SinceDate: "2020-08-01",
						UntilDate: "2021-07-31",
					}).
					Return(&book.UserMonthlyResultListResponse{MonthlyResults: results}, nil)
			},
			expect: &test.TestResponse{
				Code: http.StatusOK,
				Body: response.NewUserTopResponse(
					entity.NewMonthlyResults(results).Map(),
					now,
				),
			},
		},
		{
			name: "failed to get auth",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks, ctrl *gomock.Controller) {
				mocks.AuthService.EXPECT().
					GetAuth(gomock.Any(), &user.Empty{}).
					Return(nil, test.ErrMock)
			},
			expect: &test.TestResponse{
				Code: http.StatusInternalServerError,
			},
		},
		{
			name: "failed to list user monthly result",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks, ctrl *gomock.Controller) {
				mocks.AuthService.EXPECT().
					GetAuth(gomock.Any(), &user.Empty{}).
					Return(&user.AuthResponse{Auth: auth}, nil)
				mocks.BookService.EXPECT().
					ListUserMonthlyResult(gomock.Any(), &book.ListUserMonthlyResultRequest{
						UserId:    "00000000-0000-0000-0000-000000000000",
						SinceDate: "2020-08-01",
						UntilDate: "2021-07-31",
					}).
					Return(nil, test.ErrMock)
			},
			expect: &test.TestResponse{
				Code: http.StatusInternalServerError,
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			path := "/v1/top/user"
			req := test.NewHTTPRequest(t, http.MethodGet, path, nil)
			testHTTP(t, tt.setup, tt.expect, req)
		})
	}
}

func testMonthlyResult(year int32, month int32) *book.MonthlyResult {
	return &book.MonthlyResult{
		Year:      year,
		Month:     month,
		ReadTotal: 8,
	}
}
