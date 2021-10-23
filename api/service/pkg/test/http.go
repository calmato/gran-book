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

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

/**
 * NewHTTPRequest - HTTP Request(application/json)を生成
 */
func NewHTTPRequest(t *testing.T, method, path string, body interface{}) *http.Request {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var buf []byte
	if body != nil {
		var err error
		buf, err = json.Marshal(body)
		require.NoError(t, err, err)
	}

	req, err := http.NewRequest(method, path, bytes.NewReader(buf))
	require.NoError(t, err, err)

	req.Header.Add("Content-Type", "application/json")
	return req
}

/**
 * NewMultipartRequset - HTTP Request(multipart/form-data)を生成
 */
func NewMultipartRequest(t *testing.T, method, path, field string) *http.Request {
	buf := &bytes.Buffer{}
	writer := multipart.NewWriter(buf)
	defer writer.Close()

	filepath := getFilepath(t)
	file, err := os.Open(filepath)
	require.NoError(t, err, err)
	defer file.Close()

	header := textproto.MIMEHeader{}
	header.Set("Content-Disposition", fmt.Sprintf(`form-data; name="%s"; filename="%s"`, field, filename))
	header.Set("Content-Type", "multipart/form-data")
	part, err := writer.CreatePart(header)
	require.NoError(t, err, err)

	_, err = io.Copy(part, file)
	require.NoError(t, err, err)

	req, err := http.NewRequest(method, path, buf)
	require.NoError(t, err, err)

	req.Header.Add("Content-Type", writer.FormDataContentType())
	return req
}

/**
 * TestHTTP - HTTP Responseの検証
 */
func TestHTTP(t *testing.T, expect *HTTPResponse, res *httptest.ResponseRecorder) {
	require.Equal(t, expect.Code, res.Code)

	if isError(res) || expect.Body == nil {
		return
	}

	body, err := json.Marshal(expect.Body)
	require.NoError(t, err, err)
	require.Equal(t, string(body), res.Body.String())
}

/**
 * Private Methods
 */
func getFilepath(t *testing.T) string {
	dir, err := os.Getwd()
	assert.NoError(t, err)

	strs := strings.Split(dir, "api/service/internal/gateway")
	if len(strs) == 0 {
		t.Fatal("test: invalid file path")
	}

	return filepath.Join(strs[0], "/api/service/pkg/test", filename)
}

func isError(res *httptest.ResponseRecorder) bool {
	return res.Code < 200 || 300 <= res.Code
}
