package inquiry

import (
	"testing"

	"github.com/calmato/gran-book/api/pkg/datetime"
	"github.com/calmato/gran-book/api/pkg/test"
	"github.com/calmato/gran-book/api/proto/information"
	"github.com/stretchr/testify/assert"
)

func TestInquiries_Proto(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name      string
		inquiries Inquiries
		expect    []*information.Inquiry
	}{
		{
			name: "success",
			inquiries: Inquiries{
				{
					ID:          1,
					SenderID:    "00000000-0000-0000-0000-000000000000",
					AdminID:     "11111111-1111-1111-1111-111111111111",
					Subject:     "お問い合わせタイトル",
					Description: "お問い合わせ詳細",
					Email:       "test@calmato.jp",
					IsReplied:   true,
					CreatedAt:   test.Now(),
					UpdatedAt:   test.Now(),
				},
				{
					ID:          2,
					SenderID:    "00000000-0000-0000-0000-000000000000",
					AdminID:     "",
					Subject:     "お問い合わせタイトル",
					Description: "お問い合わせ詳細",
					Email:       "test@calmato.jp",
					IsReplied:   false,
					CreatedAt:   test.Now(),
					UpdatedAt:   test.Now(),
				},
			},
			expect: []*information.Inquiry{
				{
					Id:          1,
					SenderId:    "00000000-0000-0000-0000-000000000000",
					AdminId:     "11111111-1111-1111-1111-111111111111",
					Subject:     "お問い合わせタイトル",
					Description: "お問い合わせ詳細",
					Email:       "test@calmato.jp",
					IsReplied:   true,
					CreatedAt:   datetime.FormatTime(test.Now()),
					UpdatedAt:   datetime.FormatTime(test.Now()),
				},
				{
					Id:          2,
					SenderId:    "00000000-0000-0000-0000-000000000000",
					AdminId:     "",
					Subject:     "お問い合わせタイトル",
					Description: "お問い合わせ詳細",
					Email:       "test@calmato.jp",
					IsReplied:   false,
					CreatedAt:   datetime.FormatTime(test.Now()),
					UpdatedAt:   datetime.FormatTime(test.Now()),
				},
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.expect, tt.inquiries.Proto())
		})
	}
}
