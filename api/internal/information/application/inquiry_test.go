package application

import (
	"context"
	"testing"

	"github.com/calmato/gran-book/api/internal/information/domain/inquiry"
	"github.com/calmato/gran-book/api/pkg/test"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestInquiryApplication_Create(t *testing.T) {
	t.Parallel()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	inquiry1 := testInquiry(1)

	tests := []struct {
		name      string
		setup     func(ctx context.Context, t *testing.T, m *test.Mocks)
		inquiry   *inquiry.Inquiry
		expectErr bool
	}{
		{
			name: "success",
			setup: func(ctx context.Context, t *testing.T, m *test.Mocks) {
				m.InquiryDomainValidation.EXPECT().Inquiry(ctx, inquiry1).Return(nil)
				m.InquiryRepository.EXPECT().Create(ctx, inquiry1).Return(nil)
			},
			inquiry:   inquiry1,
			expectErr: false,
		},
		{
			name: "failed to validation error",
			setup: func(ctx context.Context, t *testing.T, m *test.Mocks) {
				m.InquiryDomainValidation.EXPECT().Inquiry(ctx, inquiry1).Return(test.ErrMock)
			},
			inquiry:   inquiry1,
			expectErr: true,
		},
		{
			name: "failed to database error",
			setup: func(ctx context.Context, t *testing.T, m *test.Mocks) {
				m.InquiryDomainValidation.EXPECT().Inquiry(ctx, inquiry1).Return(nil)
				m.InquiryRepository.EXPECT().Create(ctx, inquiry1).Return(test.ErrMock)
			},
			inquiry:   inquiry1,
			expectErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			mocks := test.NewMocks(ctrl)
			tt.setup(ctx, t, mocks)
			target := NewInquiryApplication(mocks.InquiryDomainValidation, mocks.InquiryRepository)

			err := target.Create(ctx, tt.inquiry)
			assert.Equal(t, tt.expectErr, err != nil, err)
		})
	}
}

func testInquiry(id int64) *inquiry.Inquiry {
	return &inquiry.Inquiry{
		ID:          id,
		SenderID:    "00000000-0000-0000-0000-000000000000",
		AdminID:     "11111111-1111-1111-1111-111111111111",
		Subject:     "お問い合わせタイトル",
		Description: "お問い合わせ詳細",
		Email:       "test@calmato.jp",
		IsReplied:   true,
		CreatedAt:   test.Now(),
		UpdatedAt:   test.Now(),
	}
}
