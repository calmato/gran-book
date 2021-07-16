package validation

import (
	"strings"
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
				Title:       strings.Repeat("x", 65),
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
				Description: strings.Repeat("x", 2001),
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
				Importance:  strings.Repeat("x", 65),
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

func TestNotificationRequestValidation_UpdateNotification(t *testing.T) {
	testCases := map[string]struct {
		Input    *input.UpdateNotification
		Expected bool
	}{
		"ok": {
			Input: &input.UpdateNotification{
				Title:       "Gran Book",
				Description: "Gran Bookをリリースされるにあたり、友達紹介キャンペーンを開催します。",
				Importance:  "重要",
			},
			Expected: true,
		},
		"ng_title_required": {
			Input: &input.UpdateNotification{
				Title:       "",
				Description: "Gran Bookをリリースされるにあたり、友達紹介キャンペーンを開催します。",
				Importance:  "重要",
			},
			Expected: false,
		},
		"ng_title_max": {
			Input: &input.UpdateNotification{
				Title:       strings.Repeat("x", 65),
				Description: "Gran Bookをリリースされるにあたり、友達紹介キャンペーンを開催します。",
				Importance:  "重要",
			},
			Expected: false,
		},
		"ng_description_required": {
			Input: &input.UpdateNotification{
				Title:       "Gran Book",
				Description: "",
				Importance:  "重要",
			},
			Expected: false,
		},
		"ng_description_max": {
			Input: &input.UpdateNotification{
				Title:       "Gran Book",
				Description: strings.Repeat("x", 2001),
				Importance:  "重要",
			},
			Expected: false,
		},
		"ng_importance_required": {
			Input: &input.UpdateNotification{
				Title:       "Gran Book",
				Description: "Gran Bookをリリースされるにあたり、友達紹介キャンペーンを開催します。",
				Importance:  "",
			},
			Expected: false,
		},
		"ng_importance_max": {
			Input: &input.UpdateNotification{
				Title:       "Gran Book",
				Description: "",
				Importance:  strings.Repeat("x", 65),
			},
			Expected: false,
		},
	}

	for name, tc := range testCases {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		t.Run(name, func(t *testing.T) {
			target := NewNotificationRequestValidation()

			got := target.UpdateNotification(tc.Input)
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

func TestNotificationRequestValidation_ShowNotification(t *testing.T) {
	testCases := map[string]struct {
		Input    *input.ShowNotification
		Expected bool
	}{
		"ok": {
			Input: &input.ShowNotification{
				ID: 123456789,
			},
			Expected: true,
		},
	}

	for name, tc := range testCases {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		t.Run(name, func(t *testing.T) {
			target := NewNotificationRequestValidation()

			got := target.ShowNotification(tc.Input)
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

func TestNotificationRequestValidation_DeleteNotification(t *testing.T) {
	testCases := map[string]struct {
		Input    *input.DeleteNotification
		Expected bool
	}{
		"ok": {
			Input: &input.DeleteNotification{
				ID: 123456789,
			},
			Expected: true,
		},
	}

	for name, tc := range testCases {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		t.Run(name, func(t *testing.T) {
			target := NewNotificationRequestValidation()

			got := target.DeleteNotification(tc.Input)
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
