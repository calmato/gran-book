package handler

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/calmato/gran-book/api/pkg/test"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestAPIV2Handler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mocks := test.NewMocks(ctrl)
	params := newTestParams(mocks)
	handler := newTestHandler(t, mocks, test.Now)

	actual := NewAPIV2Handler(params, test.Now)
	assert.IsType(t, handler, actual)
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

func newTestHandler(t *testing.T, mocks *test.Mocks, now func() time.Time) *apiV2Handler {
	params := newTestParams(mocks)
	return &apiV2Handler{*params, now}
}

func newTestRoutes(h *apiV2Handler, r *gin.Engine) {
	h.Routes(r.Group(""))
}

func newTestParams(mocks *test.Mocks) *Params {
	return &Params{
		Auth: mocks.AuthService,
		Book: mocks.BookService,
		User: mocks.UserService,
	}
}
