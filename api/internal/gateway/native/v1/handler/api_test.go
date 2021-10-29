package handler

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/calmato/gran-book/api/internal/gateway/entity"
	"github.com/calmato/gran-book/api/pkg/test"
	"github.com/calmato/gran-book/api/proto/user"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestAPIV1Handler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mocks := test.NewMocks(ctrl)
	params := newTestParams(mocks)
	handler := newTestHandler(t, mocks, test.Now)

	actual := NewAPIV1Handler(params, test.Now)
	assert.IsType(t, handler, actual)
}

func TestCurrentUser(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mocks := test.NewMocks(ctrl)
	auth := testAuth("00000000-0000-0000-0000-000000000000")

	type want struct {
		auth  *entity.Auth
		isErr bool
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
				auth:  entity.NewAuth(auth),
				isErr: false,
			},
		},
		{
			name: "failed to get auth",
			setup: func(ctx context.Context, mocks *test.Mocks) {
				mocks.AuthService.EXPECT().GetAuth(ctx, &user.Empty{}).Return(nil, test.ErrMock)
			},
			userID: "00000000-0000-0000-0000-000000000000",
			want: want{
				auth:  nil,
				isErr: true,
			},
		},
		{
			name: "failed to match user id",
			setup: func(ctx context.Context, mocks *test.Mocks) {
				mocks.AuthService.EXPECT().GetAuth(ctx, &user.Empty{}).Return(&user.AuthResponse{Auth: auth}, nil)
			},
			userID: "11111111-1111-1111-1111-111111111111",
			want: want{
				auth:  nil,
				isErr: true,
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			tt.setup(ctx, mocks)
			handler := newTestHandler(t, mocks, test.Now)

			actual, err := handler.currentUser(ctx, tt.userID)
			assert.Equal(t, tt.want.isErr, err != nil)
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
	expect *test.HTTPResponse,
	req *http.Request,
) {
	gin.SetMode("test")
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mocks := test.NewMocks(ctrl)
	w := httptest.NewRecorder()
	ctx, r := gin.CreateTestContext(w)
	setup(ctx, t, mocks, ctrl)

	h := newTestHandler(t, mocks, test.Now)
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
