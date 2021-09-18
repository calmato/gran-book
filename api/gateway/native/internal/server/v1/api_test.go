package v1

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/calmato/gran-book/api/gateway/native/pkg/test"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
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

	_ = newTestHandler(r, mocks)
	r.ServeHTTP(w, req)
	test.TestHTTP(t, expect, w)
}

func newTestHandler(r *gin.Engine, mocks *test.Mocks) *apiV1Handler {
	params := newTestParams(mocks)

	h := &apiV1Handler{*params}
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
