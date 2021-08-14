package test

import (
	"fmt"
	"testing"
	"time"

	mock_application "github.com/calmato/gran-book/api/server/book/mock/application"
	mock_book "github.com/calmato/gran-book/api/server/book/mock/domain/book"
	mock_validation "github.com/calmato/gran-book/api/server/book/mock/interface/validation"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

var (
	ErrMock = fmt.Errorf("some error")

	jst, _   = time.LoadLocation("Asia/Tokyo")
	TimeMock = time.Date(2021, time.Month(7), 24, 20, 0, 0, 0, jst)
	DateMock = time.Date(2021, time.Month(7), 24, 0, 0, 0, 0, jst)
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

func NewMocks(ctrl *gomock.Controller) *Mocks {
	return &Mocks{
		BookApplication:       mock_application.NewMockBookApplication(ctrl),
		BookDomainValidation:  mock_book.NewMockValidation(ctrl),
		BookRepository:        mock_book.NewMockRepository(ctrl),
		BookRequestValidation: mock_validation.NewMockBookRequestValidation(ctrl),
	}
}

func TestGRPC(t *testing.T, expect *TestResponse, res interface{}, err error) {
	if expect.Code != codes.OK {
		require.Error(t, err)

		status, ok := status.FromError(err)
		require.True(t, ok)
		require.Equal(t, expect.Code.String(), status.Code().String())
		return
	}

	require.NoError(t, err)
	require.Equal(t, expect.Message, res)
}
