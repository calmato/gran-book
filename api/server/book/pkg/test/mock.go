package test

import (
	"fmt"
	"time"

	mock_application "github.com/calmato/gran-book/api/server/book/mock/application"
	mock_book "github.com/calmato/gran-book/api/server/book/mock/domain/book"
	mock_validation "github.com/calmato/gran-book/api/server/book/mock/interface/validation"
	"github.com/calmato/gran-book/api/server/book/pkg/database"
	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/proto"
)

var (
	ErrMock = fmt.Errorf("some error")

	jst, _   = time.LoadLocation("Asia/Tokyo")
	TimeMock = time.Date(2021, time.Month(7), 24, 20, 0, 0, 0, jst).Local()
	DateMock = time.Date(2021, time.Month(7), 24, 0, 0, 0, 0, jst).Local()
)

type TestResponse struct {
	Code    codes.Code
	Message proto.Message
}

type Mocks struct {
	BookApplication       *mock_application.MockBookApplication
	BookDomainValidation  *mock_book.MockValidation
	BookRepository        *mock_book.MockRepository
	BookRequestValidation *mock_validation.MockBookRequestValidation
}

type DBMocks struct {
	UserDB *database.Client
	BookDB *database.Client
}
