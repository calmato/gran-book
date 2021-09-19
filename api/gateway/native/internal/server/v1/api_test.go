package v1

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/calmato/gran-book/api/gateway/native/pkg/datetime"
	"github.com/calmato/gran-book/api/gateway/native/pkg/test"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func testHTTP(
	t *testing.T,
	setup func(ctx context.Context, t *testing.T, mocks *test.Mocks, ctrl *gomock.Controller),
	expect *test.TestResponse,
	req *http.Request,
) {
	gin.SetMode("test")
	ctrl := gomock.NewController(t)
	mocks := test.NewMocks(ctrl)
	w := httptest.NewRecorder()
	ctx, r := gin.CreateTestContext(w)
	setup(ctx, t, mocks, ctrl)

	_ = newTestHandler(t, r, mocks)
	r.ServeHTTP(w, req)
	test.TestHTTP(t, expect, w)
}

func newTestHandler(t *testing.T, r *gin.Engine, mocks *test.Mocks) *apiV1Handler {
	params := newTestParams(mocks)
	current, err := datetime.ParseTime(test.TimeMock)
	now := func() time.Time {
		return current
	}
	require.NoError(t, err)

	h := &apiV1Handler{*params, now}
	h.Routes(r.Group(""))
	h.NonAuthRoutes(r.Group(""))

	return h
}

func newTestParams(mocks *test.Mocks) *Params {
	return &Params{
		Auth: mocks.AuthService,
		Book: mocks.BookService,
		Chat: mocks.ChatService,
		User: mocks.UserService,
	}
}
