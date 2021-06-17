package validation

import (
	"strings"
	"testing"

	"github.com/calmato/gran-book/api/server/inquiry/internal/application/input"
	"github.com/golang/mock/gomock"
)

func TestInquiryRequeestValidation_CreateInquiry(t *testing.T) {
	testCases := map[string]struct {
		Input    *input.Inquiry
		Expected bool
	}{
		"ok": {
			Input: &input.Inquiry{
				SenderId:    "00000000-0000-0000-0000-000000000000",
				Subject:     "本が検索しても見つからない",
				Description: "読みたい本を検索しても見つかりません。どうしたら追加できますか？",
				Email:       "calmato.dev@gmail.com",
			},
			Expected: true,
		},
		"ng_userId_required": {
			Input: &input.Inquiry{
				SenderId:    "",
				Subject:     "本が検索しても見つからない",
				Description: "読みたい本を検索しても見つかりません。どうしたら追加できますか？",
				Email:       "calmato.dev@gmail.com",
			},
			Expected: false,
		},
		"ng_subject_required": {
			Input: &input.Inquiry{
				SenderId:    "00000000-0000-0000-0000-000000000000",
				Subject:     "",
				Description: "読みたい本を検索しても見つかりません。どうしたら追加できますか？",
				Email:       "calmato.dev@gmail.com",
			},
			Expected: false,
		},
		"ng_subject_max": {
			Input: &input.Inquiry{
				SenderId:    "00000000-0000-0000-0000-000000000000",
				Subject:     strings.Repeat("x", 65),
				Description: "読みたい本を検索しても見つかりません。どうしたら追加できますか？",
				Email:       "calmato.dev@gmail.com",
			},
			Expected: false,
		},
		"ng_description_required": {
			Input: &input.Inquiry{
				SenderId:    "00000000-0000-0000-0000-000000000000",
				Subject:     "本が検索しても見つからない",
				Description: "",
				Email:       "calmato.dev@gmail.com",
			},
			Expected: false,
		},
		"ng_description_max": {
			Input: &input.Inquiry{
				SenderId:    "00000000-0000-0000-0000-000000000000",
				Subject:     "本が検索しても見つからない",
				Description: strings.Repeat("x", 1001),
				Email:       "calmato.dev@gmail.com",
			},
			Expected: false,
		},
		"ng_email_required": {
			Input: &input.Inquiry{
				SenderId:    "00000000-0000-0000-0000-000000000000",
				Subject:     "本が検索しても見つからない",
				Description: "読みたい本を検索しても見つかりません。どうしたら追加できますか？",
				Email:       "",
			},
			Expected: false,
		},
		"ng_email_max": {
			Input: &input.Inquiry{
				SenderId:    "00000000-0000-0000-0000-000000000000",
				Subject:     "本が検索しても見つからない",
				Description: "読みたい本を検索しても見つかりません。どうしたら追加できますか？",
				Email:       strings.Repeat("x", 257),
			},
			Expected: false,
		},
	}

	for name, tc := range testCases {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		t.Run(name, func(t *testing.T) {
			target := NewInquiryRequestValidation()

			got := target.Inquiry(tc.Input)
			if !tc.Expected {
				if got != nil {
					t.Fatalf("Incorrect result: %#v", got)
				}
			} else {
				if got == nil {
					t.Fatalf("Incorrect result: result is nil")
				}
			}
		})
	}
}
