package server

import (
	"context"
	"testing"

	"github.com/calmato/gran-book/api/pkg/exception"
	"github.com/calmato/gran-book/api/pkg/test"
	pb "github.com/calmato/gran-book/api/proto/information"
	"github.com/golang/mock/gomock"
	"google.golang.org/grpc/codes"
)

func TestInquiryServer_CreateInquiry(t *testing.T) {
	t.Parallel()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tests := []struct {
		name   string
		setup  func(ctx context.Context, t *testing.T, m *test.Mocks)
		req    *pb.CreateInquiryRequest
		expect *test.Response
	}{
		{
			name: "success",
			setup: func(ctx context.Context, t *testing.T, m *test.Mocks) {
				m.InquiryRequestValidation.EXPECT().CreateInquiry(gomock.Any()).Return(nil)
				m.InquiryApplication.EXPECT().Create(ctx, gomock.Any()).Return(nil)
			},
			req: &pb.CreateInquiryRequest{
				UserId:      "00000000-0000-0000-0000-000000000000",
				Subject:     "タイトル",
				Description: "詳細",
				Email:       "test@calmato.jp",
			},
			expect: &test.Response{
				Code: codes.OK,
				Message: &pb.InquiryResponse{
					Inquiry: &pb.Inquiry{
						Id:          0,
						SenderId:    "00000000-0000-0000-0000-000000000000",
						AdminId:     "",
						Subject:     "タイトル",
						Description: "詳細",
						Email:       "test@calmato.jp",
						IsReplied:   false,
						CreatedAt:   "",
						UpdatedAt:   "",
					},
				},
			},
		},
		{
			name: "failed: invalid argument",
			setup: func(ctx context.Context, t *testing.T, m *test.Mocks) {
				err := exception.ErrInvalidRequestValidation.New(test.ErrMock)
				m.InquiryRequestValidation.EXPECT().CreateInquiry(gomock.Any()).Return(err)
			},
			req: &pb.CreateInquiryRequest{},
			expect: &test.Response{
				Code: codes.InvalidArgument,
			},
		},
		{
			name: "failed: internal error",
			setup: func(ctx context.Context, t *testing.T, m *test.Mocks) {
				err := exception.ErrInDatastore.New(test.ErrMock)
				m.InquiryRequestValidation.EXPECT().CreateInquiry(gomock.Any()).Return(nil)
				m.InquiryApplication.EXPECT().Create(ctx, gomock.Any()).Return(err)
			},
			req: &pb.CreateInquiryRequest{
				UserId:      "00000000-0000-0000-0000-000000000000",
				Subject:     "タイトル",
				Description: "詳細",
				Email:       "test@calmato.jp",
			},
			expect: &test.Response{
				Code: codes.Internal,
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			mocks := test.NewMocks(ctrl)
			tt.setup(ctx, t, mocks)

			target := NewInquiryServer(mocks.InquiryRequestValidation, mocks.InquiryApplication)
			res, err := target.CreateInquiry(ctx, tt.req)
			test.GRPC(t, tt.expect, res, err)
		})
	}
}
