package validation

import (
	"fmt"
	"strings"
	"testing"

	"github.com/calmato/gran-book/api/proto/information"
	"github.com/stretchr/testify/assert"
)

func TestInquiryRequestValdiation_CreateInquiry(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name      string
		req       *information.CreateInquiryRequest
		expectErr bool
	}{
		{
			name: "success",
			req: &information.CreateInquiryRequest{
				UserId:      "00000000-0000-0000-0000-000000000000",
				Subject:     "タイトル",
				Description: "詳細",
				Email:       "test@calmato.jp",
			},
			expectErr: false,
		},
		{
			name: "validation error: UserId.min",
			req: &information.CreateInquiryRequest{
				UserId:      "",
				Subject:     "タイトル",
				Description: "詳細",
				Email:       "test@calmato.jp",
			},
			expectErr: true,
		},
		{
			name: "validation error: Subject.min",
			req: &information.CreateInquiryRequest{
				UserId:      "00000000-0000-0000-0000-000000000000",
				Subject:     "",
				Description: "詳細",
				Email:       "test@calmato.jp",
			},
			expectErr: true,
		},
		{
			name: "validation error: Subject.max",
			req: &information.CreateInquiryRequest{
				UserId:      "00000000-0000-0000-0000-000000000000",
				Subject:     strings.Repeat("x", 65),
				Description: "詳細",
				Email:       "test@calmato.jp",
			},
			expectErr: true,
		},
		{
			name: "validation error: Description.min",
			req: &information.CreateInquiryRequest{
				UserId:      "00000000-0000-0000-0000-000000000000",
				Subject:     "タイトル",
				Description: "",
				Email:       "test@calmato.jp",
			},
			expectErr: true,
		},
		{
			name: "validation error: Description.max",
			req: &information.CreateInquiryRequest{
				UserId:      "00000000-0000-0000-0000-000000000000",
				Subject:     "タイトル",
				Description: strings.Repeat("x", 1001),
				Email:       "test@calmato.jp",
			},
			expectErr: true,
		},
		{
			name: "validation error: email.min",
			req: &information.CreateInquiryRequest{
				UserId:      "00000000-0000-0000-0000-000000000000",
				Subject:     "タイトル",
				Description: "詳細",
				Email:       "",
			},
			expectErr: true,
		},
		{
			name: "validation error: email.max",
			req: &information.CreateInquiryRequest{
				UserId:      "00000000-0000-0000-0000-000000000000",
				Subject:     "タイトル",
				Description: "詳細",
				Email:       fmt.Sprintf("%s@calmato.jp", strings.Repeat("x", 246)),
			},
			expectErr: true,
		},
		{
			name: "validation error: email.format",
			req: &information.CreateInquiryRequest{
				UserId:      "00000000-0000-0000-0000-000000000000",
				Subject:     "タイトル",
				Description: "詳細",
				Email:       "calmato.jp",
			},
			expectErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			target := NewInquiryRequestValidation()
			err := target.CreateInquiry(tt.req)
			assert.Equal(t, tt.expectErr, err != nil, err)
		})
	}
}
