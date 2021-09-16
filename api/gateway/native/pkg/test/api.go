package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"net/textproto"
	"os"
	"path/filepath"
	"strings"
	"testing"

	mock_book "github.com/calmato/gran-book/api/gateway/native/mock/book"
	mock_chat "github.com/calmato/gran-book/api/gateway/native/mock/chat"
	mock_user "github.com/calmato/gran-book/api/gateway/native/mock/user"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func NewHTTPMock(
	t *testing.T,
	method, path string,
	body interface{},
) (*gin.Context, *httptest.ResponseRecorder, *Mocks) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var buf []byte
	if body != nil {
		var err error
		buf, err = json.Marshal(body)
		require.NoError(t, err)
	}

	ctx, res := newTestHTTP(t, method, path, bytes.NewReader(buf))
	ctx.Request.Header.Add("Content-Type", "application/json")
	mocks := NewMocks(ctrl)
	return ctx, res, mocks
}

func NewMultipartMock(
	t *testing.T,
	ctrl *gomock.Controller,
	method, path, field string,
) (*gin.Context, *httptest.ResponseRecorder, *Mocks) {
	buf := &bytes.Buffer{}
	writer := multipart.NewWriter(buf)
	defer writer.Close()

	filepath := getFilepath(t)
	file, err := os.Open(filepath)
	require.NoError(t, err)
	defer file.Close()

	header := textproto.MIMEHeader{}
	header.Set("Content-Disposition", fmt.Sprintf(`form-data; name="%s"; filename="%s"`, field, filename))
	header.Set("Content-Type", "multipart/form-data")
	part, err := writer.CreatePart(header)
	require.NoError(t, err)

	_, err = io.Copy(part, file)
	require.NoError(t, err)

	ctx, res := newTestHTTP(t, method, path, buf)
	ctx.Request.Header.Add("Content-Type", writer.FormDataContentType())
	mocks := NewMocks(ctrl)
	fmt.Printf("api=%+v\n", ctx.Request.Header)
	return ctx, res, mocks
}

func newTestHTTP(
	t *testing.T,
	method, path string,
	body io.Reader,
) (*gin.Context, *httptest.ResponseRecorder) {
	req, err := http.NewRequest(method, path, body)
	require.NoError(t, err)

	res := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(res)
	ctx.Request = req

	return ctx, res
}

func getFilepath(t *testing.T) string {
	dir, err := os.Getwd()
	assert.NoError(t, err)

	strs := strings.Split(dir, "api/gateway/native")
	if len(strs) == 0 {
		t.Fatal("test: invalid file path")
	}

	return filepath.Join(strs[0], "/api/gateway/native/pkg/test", filename)
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
