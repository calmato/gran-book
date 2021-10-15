package test

import (
	"errors"

	mock_user_application "github.com/calmato/gran-book/api/service/mock/user/application"
	mock_user "github.com/calmato/gran-book/api/service/mock/user/domain/user"
	mock_user_validation "github.com/calmato/gran-book/api/service/mock/user/interface/validation"
	"github.com/calmato/gran-book/api/service/pkg/database"
	"github.com/calmato/gran-book/api/service/pkg/datetime"
	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/proto"
)

var (
	ErrMock     = errors.New("some error")
	TimeMock, _ = datetime.ParseTime("2021-07-24 20:00:00")
	DateMock, _ = datetime.ParseDate("2021-07-24")
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
