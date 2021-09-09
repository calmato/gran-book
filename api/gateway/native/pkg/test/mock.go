package test

import (
	"errors"
	"time"

	"github.com/calmato/gran-book/api/gateway/native/mock"
)

var (
	ErrMock = errors.New("some error")

	jst, _   = time.LoadLocation("Asia/Tokyo")
	TimeMock = time.Date(2021, time.Month(7), 24, 20, 0, 0, 0, jst).Local().String()
	DateMock = time.Date(2021, time.Month(7), 24, 0, 0, 0, 0, jst).Local().String()
)

type TestResponse struct {
	Code int
	Body interface{}
}

type Mocks struct {
	AdminService *mock.MockAdminServiceClient
	AuthService  *mock.MockAuthServiceClient
	BookService  *mock.MockBookServiceClient
	ChatService  *mock.MockChatServiceClient
	UserService  *mock.MockUserServiceClient
}
