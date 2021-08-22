package test

import (
	"testing"

	mock_application "github.com/calmato/gran-book/api/server/book/mock/application"
	mock_book "github.com/calmato/gran-book/api/server/book/mock/domain/book"
	mock_validation "github.com/calmato/gran-book/api/server/book/mock/interface/validation"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

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
