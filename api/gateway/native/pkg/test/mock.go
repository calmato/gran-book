package test

import (
	"errors"
	"time"

	mock_book "github.com/calmato/gran-book/api/gateway/native/mock/book"
	mock_chat "github.com/calmato/gran-book/api/gateway/native/mock/chat"
	mock_user "github.com/calmato/gran-book/api/gateway/native/mock/user"
	"github.com/calmato/gran-book/api/gateway/native/pkg/datetime"
)

const filename = "calmato.png"

var (
	ErrMock = errors.New("some error")

	jst, _   = time.LoadLocation("Asia/Tokyo")
	TimeMock = datetime.FormatTime(time.Date(2021, time.Month(7), 24, 20, 0, 0, 0, jst))
	DateMock = datetime.FormatDate(time.Date(2021, time.Month(7), 24, 0, 0, 0, 0, jst))
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
