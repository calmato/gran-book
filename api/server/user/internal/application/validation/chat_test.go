package validation

import (
	"strings"
	"testing"

	"github.com/calmato/gran-book/api/server/user/internal/application/input"
	"github.com/golang/mock/gomock"
)

func TestChatRequestValidation_CreateRoom(t *testing.T) {
	testCases := map[string]struct {
		Input    *input.CreateRoom
		Expected bool
	}{
		"ok": {
			Input: &input.CreateRoom{
				UserIDs: []string{
					"09c90850-fa2f-4fae-95e2-922d25f90c2c",
					"78b10377-4e3e-48fd-892c-36bcd7c91f3b",
				},
			},
			Expected: true,
		},
		"ng_userIds_unique": {
			Input: &input.CreateRoom{
				UserIDs: []string{
					"09c90850-fa2f-4fae-95e2-922d25f90c2c",
					"09c90850-fa2f-4fae-95e2-922d25f90c2c",
				},
			},
			Expected: false,
		},
		"ng_userIds_required": {
			Input: &input.CreateRoom{
				UserIDs: []string{
					"09c90850-fa2f-4fae-95e2-922d25f90c2c",
					"",
				},
			},
			Expected: false,
		},
		"ng_userIds_max": {
			Input: &input.CreateRoom{
				UserIDs: []string{
					strings.Repeat("x", 37),
					"78b10377-4e3e-48fd-892c-36bcd7c91f3b",
				},
			},
			Expected: false,
		},
	}

	for name, tc := range testCases {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		t.Run(name, func(t *testing.T) {
			target := NewChatRequestValidation()

			got := target.CreateRoom(tc.Input)
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

func TestChatRequestValidation_CreateTextMessage(t *testing.T) {
	testCases := map[string]struct {
		Input    *input.CreateTextMessage
		Expected bool
	}{
		"ok": {
			Input: &input.CreateTextMessage{
				Text: "テストメッセージ",
			},
			Expected: true,
		},
		"ng_text_required": {
			Input: &input.CreateTextMessage{
				Text: "",
			},
			Expected: false,
		},
		"ng_text_max": {
			Input: &input.CreateTextMessage{
				Text: strings.Repeat("x", 1001),
			},
			Expected: false,
		},
	}

	for name, tc := range testCases {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		t.Run(name, func(t *testing.T) {
			target := NewChatRequestValidation()

			got := target.CreateTextMessage(tc.Input)
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

func TestChatRequestValidation_CreateImageMessage(t *testing.T) {
	testCases := map[string]struct {
		Input    *input.CreateImageMessage
		Expected bool
	}{
		"ok": {
			Input: &input.CreateImageMessage{
				Image: []byte("あいうえお"),
			},
			Expected: true,
		},
		"ng_text_required": {
			Input: &input.CreateImageMessage{
				Image: nil,
			},
			Expected: false,
		},
	}

	for name, tc := range testCases {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		t.Run(name, func(t *testing.T) {
			target := NewChatRequestValidation()

			got := target.CreateImageMessage(tc.Input)
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
