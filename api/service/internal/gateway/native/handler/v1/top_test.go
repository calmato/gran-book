package v1

import (
	"context"
	"net/http"
	"testing"

	"github.com/calmato/gran-book/api/service/internal/gateway/entity"
	response "github.com/calmato/gran-book/api/service/internal/gateway/native/response/v1"
	"github.com/calmato/gran-book/api/service/pkg/test"
	"github.com/calmato/gran-book/api/service/proto/book"
	"github.com/calmato/gran-book/api/service/proto/user"
	"github.com/golang/mock/gomock"
)

func TestTop_GetTopUser(t *testing.T) {
	t.Parallel()

	auth := testAuth("00000000-0000-0000-0000-000000000000")
	results := make([]*book.MonthlyResult, 2)
	results[0] = testMonthlyResult(2021, 8)
	results[1] = testMonthlyResult(2021, 7)

	tests := []struct {
		name   string
		setup  func(ctx context.Context, t *testing.T, mocks *test.Mocks, ctrl *gomock.Controller)
		expect *test.HTTPResponse
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
			expect: &test.HTTPResponse{
				Code: http.StatusOK,
				Body: response.NewUserTopResponse(
					entity.NewMonthlyResults(results).Map(),
					test.Now(),
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
			expect: &test.HTTPResponse{
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
			expect: &test.HTTPResponse{
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
