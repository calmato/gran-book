package test

import (
	"errors"
	"time"

	mock_book "github.com/calmato/gran-book/api/gateway/native/mock/book"
	mock_chat "github.com/calmato/gran-book/api/gateway/native/mock/chat"
	mock_user "github.com/calmato/gran-book/api/gateway/native/mock/user"
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
	AdminService *mock_user.MockAdminServiceClient
	AuthService  *mock_user.MockAuthServiceClient
	BookService  *mock_book.MockBookServiceClient
	ChatService  *mock_chat.MockChatServiceClient
	UserService  *mock_user.MockUserServiceClient
}
