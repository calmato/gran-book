package test

import (
	"fmt"
	"time"

	"github.com/calmato/gran-book/api/server/information/pkg/database"
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

type Mocks struct{}

type DBMocks struct {
	UserDB        *database.Client
	InformationDB *database.Client
}
