package test

import (
	"errors"
	"time"

	mock_user_application "github.com/calmato/gran-book/api/service/mock/user/application"
	mock_user "github.com/calmato/gran-book/api/service/mock/user/domain/user"
	mock_user_validation "github.com/calmato/gran-book/api/service/mock/user/interface/validation"
	"github.com/calmato/gran-book/api/service/pkg/database"
	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/proto"
)

var (
	ErrMock = errors.New("some error")

	jst, _   = time.LoadLocation("Asia/Tokyo")
	TimeMock = time.Date(2021, time.Month(7), 24, 20, 0, 0, 0, jst).Local()
	DateMock = time.Date(2021, time.Month(7), 24, 0, 0, 0, 0, jst).Local()
)

type TestResponse struct {
	Code    codes.Code
	Message proto.Message
}

type Mocks struct {
	AdminRequestValidation *mock_user_validation.MockAdminRequestValidation
	AuthRequestValidation  *mock_user_validation.MockAuthRequestValidation
	UserApplication        *mock_user_application.MockUserApplication
	UserDomainValidation   *mock_user.MockValidation
	UserRepository         *mock_user.MockRepository
	UserRequestValidation  *mock_user_validation.MockUserRequestValidation
	UserUploader           *mock_user.MockUploader
}

type DBMocks struct {
	UserDB *database.Client
	BookDB *database.Client
}
