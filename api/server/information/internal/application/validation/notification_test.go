package validation

import (
	"testing"

	"github.com/calmato/gran-book/api/server/information/internal/application/input"
	"github.com/golang/mock/gomock"
)

func TestNotificationRequestValidation_CreateNotification(t *testing.T) {
	testCases := map[string]struct {
		Input    *input.CreateNotification
		Expected bool
	}{
		"ok": {
			Input: &input.CreateNotification{
				Title:       "Gran Book",
				Description: "Gran Bookをリリースされるにあたり、友達紹介キャンペーンを開催します。",
				Importance:  "重要",
			},
			Expected: true,
		},
		"ng_title_required": {
			Input: &input.CreateNotification{
				Title:       "",
				Description: "Gran Bookをリリースされるにあたり、友達紹介キャンペーンを開催します。",
				Importance:  "重要",
			},
			Expected: false,
		},
		"ng_title_max": {
			Input: &input.CreateNotification{
				Title:       "境界値test(65)境界値test(65)境界値test(65)境界値test(65)境界値test(65)境界値test(65",
				Description: "Gran Bookをリリースされるにあたり、友達紹介キャンペーンを開催します。",
				Importance:  "重要",
			},
			Expected: false,
		},
		"ng_description_required": {
			Input: &input.CreateNotification{
				Title:       "Gran Book",
				Description: "",
				Importance:  "重要",
			},
			Expected: false,
		},
		"ng_description_max": {
			Input: &input.CreateNotification{
				Title:       "Gran Book",
				Description: "境界値test(2001)境界値test(2001)境界値test(2001)境界値test(2001)境界値test(2001)境界値test(2001)境界値test(2001)境界値test(2001)境界値test(2001)境界値test(2001)境界値test(2001)境界値test(2001)境界値test(2001)境界値test(2001)境界値test(2001)境界値test(2001)境界値test(2001)境界値test(2001)境界値test(2001)境界値test(2001)境界値test(2001)境界値test(2001)境界値test(2001)境界値test(2001)境界値test(2001)境界値test(2001)境界値test(2001)境界値test(2001)境界値test(2001)境界値test(2001)境界値test(2001)境界値test(2001)境界値test(2001)境界値test(2001)境界値test(2001)境界値test(2001)境界値test(2001)境界値test(2001)境界値test(2001)境界値test(2001)境界値test(2001)境界値test(2001)境界値test(2001)境界値test(2001)境界値test(2001)境界値test(2001)境界値test(2001)境界値test(2001)境界値test(2001)境界値test(2001)境界値test(2001)境界値test(2001)境界値test(2001)境界値test(2001)境界値test(2001)境界値test(2001)境界値test(2001)境界値test(2001)境界値test(2001)境界値test(2001)境界値test(2001)境界値test(2001)境界値test(2001)境界値test(2001)境界値test(2001)境界値test(2001)境界値test(2001)境界値test(2001)境界値test(2001)境界値test(2001)境界値test(2001)境界値test(2001)境界値test(2001)境界値test(2001)境界値test(2001)境界値test(2001)境界値test(2001)境界値test(2001)境界値test(2001)境界値test(2001)境界値test(2001)境界値test(2001)境界値test(2001)境界値test(2001)境界値test(2001)境界値test(2001)境界値test(2001)境界値test(2001)境界値test(2001)境界値test(2001)境界値test(2001)境界値test(2001)境界値test(2001)境界値test(2001)境界値test(2001)境界値test(2001)境界値test(2001)境界値test(2001)境界値test(2001)境界値test(2001)境界値test(2001)境界値test(2001)境界値test(2001)境界値test(2001)境界値test(2001)境界値test(2001)境界値test(2001)境界値test(2001)境界値test(2001)境界値test(2001)境界値test(2001)境界値test(2001)境界値test(2001)境界値test(2001)境界値test(2001)境界値test(2001)境界値test(2001)境界値test(2001)境界値test(2001)境界値test(2001)境界値test(2001)境界値test(2001)境界値test(2001)境界値test(2001)境界値test(2001)境界値test(2001)境界値test(2001)境界値test(2001)境界値test(2001)境界値test(2001)境界値test(2001)境界値test(2001)境界値test(2001)境界値test(2001)境界値test(2001)境界値test(2001)境界値test(2001)境界値test(2001)境界値test(2001)境界値test(2001)境界値test(2001)境界値test(2001)境界値test(2001)境界値test(2001)境界値test(2001)境界値test(2001)境界値test(2001)境界値test(2001)境界値test(2001)境界値test(2001)境界値test(2001)境界値test(2001)境界値test(2001)境界値test(2001",
				Importance:  "重要",
			},
			Expected: false,
		},
		"ng_importance_required": {
			Input: &input.CreateNotification{
				Title:       "Gran Book",
				Description: "Gran Bookをリリースされるにあたり、友達紹介キャンペーンを開催します。",
				Importance:  "",
			},
			Expected: false,
		},
		"ng_importance_max": {
			Input: &input.CreateNotification{
				Title:       "Gran Book",
				Description: "",
				Importance:  "境界値test(65)境界値test(65)境界値test(65)境界値test(65)境界値test(65)境界値test(65",
			},
			Expected: false,
		},
	}

	for name, tc := range testCases {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		t.Run(name, func(t *testing.T) {
			target := NewNotificationRequestValidation()

			got := target.CreateNotification(tc.Input)
			if tc.Expected {
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
