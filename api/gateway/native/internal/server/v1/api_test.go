package v1

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/calmato/gran-book/api/gateway/native/internal/entity"
	"github.com/calmato/gran-book/api/gateway/native/pkg/datetime"
	"github.com/calmato/gran-book/api/gateway/native/pkg/test"
	"github.com/calmato/gran-book/api/gateway/native/proto/service/user"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAPIV1Handler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	current, err := datetime.ParseTime(test.TimeMock)
	require.NoError(t, err)
	now := func() time.Time {
		return current
	}

	mocks := test.NewMocks(ctrl)
	params := newTestParams(mocks)
	handler := newTestHandler(t, mocks, now)

	actual := NewAPIV1Handler(params, now)
	assert.IsType(t, handler, actual)
}

func TestCurrentUser(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	current, err := datetime.ParseTime(test.TimeMock)
	require.NoError(t, err)
	now := func() time.Time {
		return current
	}

	mocks := test.NewMocks(ctrl)
	auth := testAuth("00000000-0000-0000-0000-000000000000")

	type want struct {
		auth *entity.Auth
		err  error
	}
	tests := []struct {
		name   string
		setup  func(ctx context.Context, mocks *test.Mocks)
		userID string
		want   want
	}{
		{
			name: "success",
			setup: func(ctx context.Context, mocks *test.Mocks) {
				mocks.AuthService.EXPECT().GetAuth(ctx, &user.Empty{}).Return(&user.AuthResponse{Auth: auth}, nil)
			},
			userID: "00000000-0000-0000-0000-000000000000",
			want: want{
				auth: entity.NewAuth(auth),
				err:  nil,
			},
		},
		{
			name: "failed to get auth",
			setup: func(ctx context.Context, mocks *test.Mocks) {
				mocks.AuthService.EXPECT().GetAuth(ctx, &user.Empty{}).Return(nil, test.ErrMock)
			},
			userID: "00000000-0000-0000-0000-000000000000",
			want: want{
				auth: nil,
				err:  test.ErrMock,
			},
		},
		{
			name: "failed to match user id",
			setup: func(ctx context.Context, mocks *test.Mocks) {
				mocks.AuthService.EXPECT().GetAuth(ctx, &user.Empty{}).Return(&user.AuthResponse{Auth: auth}, nil)
			},
			userID: "11111111-1111-1111-1111-111111111111",
			want: want{
				auth: nil,
				err:  entity.ErrForbidden.New(errUnmatchUserID),
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			tt.setup(ctx, mocks)
			handler := newTestHandler(t, mocks, now)

			actual, err := handler.currentUser(ctx, tt.userID)
			assert.ErrorIs(t, tt.want.err, err)
			assert.Equal(t, tt.want.auth, actual)
		})
	}
}

/**
 * Test Methods
 */
func testHTTP(
	t *testing.T,
	setup func(ctx context.Context, t *testing.T, mocks *test.Mocks, ctrl *gomock.Controller),
	expect *test.TestResponse,
	req *http.Request,
) {
	gin.SetMode("test")
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	current, err := datetime.ParseTime(test.TimeMock)
	require.NoError(t, err)
	now := func() time.Time {
		return current
	}

	mocks := test.NewMocks(ctrl)
	w := httptest.NewRecorder()
	ctx, r := gin.CreateTestContext(w)
	setup(ctx, t, mocks, ctrl)

	h := newTestHandler(t, mocks, now)
	newTestRoutes(h, r)

	r.ServeHTTP(w, req)
	test.TestHTTP(t, expect, w)
}

func newTestHandler(t *testing.T, mocks *test.Mocks, now func() time.Time) *apiV1Handler {
	params := newTestParams(mocks)
	return &apiV1Handler{*params, now}
}

func newTestRoutes(h *apiV1Handler, r *gin.Engine) {
	h.Routes(r.Group(""))
	h.NonAuthRoutes(r.Group(""))
}

func newTestParams(mocks *test.Mocks) *Params {
	return &Params{
		Auth: mocks.AuthService,
		Book: mocks.BookService,
		Chat: mocks.ChatService,
		User: mocks.UserService,
	}
}
