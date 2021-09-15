package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	mock_book "github.com/calmato/gran-book/api/gateway/native/mock/book"
	mock_chat "github.com/calmato/gran-book/api/gateway/native/mock/chat"
	mock_user "github.com/calmato/gran-book/api/gateway/native/mock/user"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func NewHTTPClient(t *testing.T, body interface{}) (*gin.Context, *httptest.ResponseRecorder, *Mocks) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var jb []byte
	if body != nil {
		var err error
		jb, err = json.Marshal(body)
		require.NoError(t, err)
	}

	req, err := http.NewRequest("", "", bytes.NewReader(jb))
	require.NoError(t, err)

	res := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(res)
	ctx.Request = req

	mocks := NewMocks(ctrl)

	return ctx, res, mocks
}

func TestHTTP(t *testing.T, expect *TestResponse, res *httptest.ResponseRecorder) {
	require.Equal(t, expect.Code, res.Code)

	if isError(res) || expect.Body == nil {
		return
	}

	body, err := json.Marshal(expect.Body)
	require.NoError(t, err)
	require.Equal(t, string(body), res.Body.String())
}

func isError(res *httptest.ResponseRecorder) bool {
	return res.Code < 200 || 300 <= res.Code
}

func NewMocks(ctrl *gomock.Controller) *Mocks {
	return &Mocks{
		AdminService: mock_user.NewMockAdminServiceClient(ctrl),
		AuthService:  mock_user.NewMockAuthServiceClient(ctrl),
		BookService:  mock_book.NewMockBookServiceClient(ctrl),
		ChatService:  mock_chat.NewMockChatServiceClient(ctrl),
		UserService:  mock_user.NewMockUserServiceClient(ctrl),
	}
}
