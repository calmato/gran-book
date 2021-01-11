package validation

import (
	"strings"
	"testing"

	"github.com/calmato/gran-book/api/user/internal/application/input"
	"github.com/golang/mock/gomock"
)

func TestAuthRequestValidation_CreateAuth(t *testing.T) {
	testCases := map[string]struct {
		Input    *input.CreateAuth
		Expected bool
	}{
		"ok": {
			Input: &input.CreateAuth{
				Username:             "test-user",
				Email:                "test-user@calmato.com",
				Password:             "12345678",
				PasswordConfirmation: "12345678",
			},
			Expected: true,
		},
		"ng_username_required": {
			Input: &input.CreateAuth{
				Username:             "",
				Email:                "test-user@calmato.com",
				Password:             "12345678",
				PasswordConfirmation: "12345678",
			},
			Expected: false,
		},
		"ng_username_max": {
			Input: &input.CreateAuth{
				Username:             strings.Repeat("x", 33),
				Email:                "test-user@calmato.com",
				Password:             "12345678",
				PasswordConfirmation: "12345678",
			},
			Expected: false,
		},
		"ng_email_required": {
			Input: &input.CreateAuth{
				Username:             "test-user",
				Email:                "",
				Password:             "12345678",
				PasswordConfirmation: "12345678",
			},
			Expected: false,
		},
		"ng_email_max": {
			Input: &input.CreateAuth{
				Username:             "test-user",
				Email:                strings.Repeat("x", 256) + "@calmato.com",
				Password:             "12345678",
				PasswordConfirmation: "12345678",
			},
			Expected: false,
		},
		"ng_email_format": {
			Input: &input.CreateAuth{
				Username:             "test-user",
				Email:                "test-user",
				Password:             "12345678",
				PasswordConfirmation: "12345678",
			},
			Expected: false,
		},
		"ng_password_required": {
			Input: &input.CreateAuth{
				Username:             "test-user",
				Email:                "test-user@calmato.com",
				Password:             "",
				PasswordConfirmation: "12345678",
			},
			Expected: false,
		},
		"ng_password_min": {
			Input: &input.CreateAuth{
				Username:             "test-user",
				Email:                "test-user@calmato.com",
				Password:             strings.Repeat("x", 5),
				PasswordConfirmation: strings.Repeat("x", 5),
			},
			Expected: false,
		},
		"ng_password_max": {
			Input: &input.CreateAuth{
				Username:             "test-user",
				Email:                "test-user@calmato.com",
				Password:             strings.Repeat("x", 33),
				PasswordConfirmation: strings.Repeat("x", 33),
			},
			Expected: false,
		},
		"ng_password_format": {
			Input: &input.CreateAuth{
				Username:             "test-user",
				Email:                "test-user@calmato.com",
				Password:             "１２３４５６７８",
				PasswordConfirmation: "１２３４５６７８",
			},
			Expected: false,
		},
		"ng_passwordConfirmation_required": {
			Input: &input.CreateAuth{
				Username:             "test-user",
				Email:                "test-user@calmato.com",
				Password:             "12345678",
				PasswordConfirmation: "",
			},
			Expected: false,
		},
		"ng_passwordConfirmation_equal": {
			Input: &input.CreateAuth{
				Username:             "test-user",
				Email:                "test-user@calmato.com",
				Password:             "12345678",
				PasswordConfirmation: "1234567",
			},
			Expected: false,
		},
	}

	for result, tc := range testCases {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		t.Run(result, func(t *testing.T) {
			target := NewAuthRequestValidation()

			got := target.CreateAuth(tc.Input)
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

func TestAuthRequestValidation_UpdateAuthEmail(t *testing.T) {
	testCases := map[string]struct {
		Input    *input.UpdateAuthEmail
		Expected bool
	}{
		"ok": {
			Input: &input.UpdateAuthEmail{
				Email: "test-user@calmato.com",
			},
			Expected: true,
		},
		"ng_email_required": {
			Input: &input.UpdateAuthEmail{
				Email: "",
			},
			Expected: false,
		},
		"ng_email_max": {
			Input: &input.UpdateAuthEmail{
				Email: strings.Repeat("x", 256) + "@calmato.com",
			},
			Expected: false,
		},
		"ng_email_format": {
			Input: &input.UpdateAuthEmail{
				Email: "test-user",
			},
			Expected: false,
		},
	}

	for result, tc := range testCases {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		t.Run(result, func(t *testing.T) {
			target := NewAuthRequestValidation()

			got := target.UpdateAuthEmail(tc.Input)
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

func TestAuthRequestValidation_UpdateAuthPassword(t *testing.T) {
	testCases := map[string]struct {
		Input    *input.UpdateAuthPassword
		Expected bool
	}{
		"ok": {
			Input: &input.UpdateAuthPassword{
				Password:             "12345678",
				PasswordConfirmation: "12345678",
			},
			Expected: true,
		},
		"ng_password_required": {
			Input: &input.UpdateAuthPassword{
				Password:             "",
				PasswordConfirmation: "12345678",
			},
			Expected: false,
		},
		"ng_password_min": {
			Input: &input.UpdateAuthPassword{
				Password:             strings.Repeat("x", 5),
				PasswordConfirmation: strings.Repeat("x", 5),
			},
			Expected: false,
		},
		"ng_password_max": {
			Input: &input.UpdateAuthPassword{
				Password:             strings.Repeat("x", 33),
				PasswordConfirmation: strings.Repeat("x", 33),
			},
			Expected: false,
		},
		"ng_password_format": {
			Input: &input.UpdateAuthPassword{
				Password:             "１２３４５６７８",
				PasswordConfirmation: "１２３４５６７８",
			},
			Expected: false,
		},
		"ng_passwordConfirmation_required": {
			Input: &input.UpdateAuthPassword{
				Password:             "12345678",
				PasswordConfirmation: "",
			},
			Expected: false,
		},
		"ng_passwordConfirmation_equal": {
			Input: &input.UpdateAuthPassword{
				Password:             "12345678",
				PasswordConfirmation: "1234567",
			},
			Expected: false,
		},
	}

	for result, tc := range testCases {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		t.Run(result, func(t *testing.T) {
			target := NewAuthRequestValidation()

			got := target.UpdateAuthPassword(tc.Input)
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

func TestAuthRequestValidation_UpdateAuthProfile(t *testing.T) {
	testCases := map[string]struct {
		Input    *input.UpdateAuthProfile
		Expected bool
	}{
		"ok": {
			Input: &input.UpdateAuthProfile{
				Username:         "test-user",
				Gender:           0,
				Thumbnail:        "",
				SelfIntroduction: "自己紹介",
			},
			Expected: true,
		},
		"ng_username_required": {
			Input: &input.UpdateAuthProfile{
				Username:         "",
				Gender:           0,
				Thumbnail:        "",
				SelfIntroduction: "自己紹介",
			},
			Expected: false,
		},
		"ng_username_max": {
			Input: &input.UpdateAuthProfile{
				Username:         strings.Repeat("x", 33),
				Gender:           0,
				Thumbnail:        "",
				SelfIntroduction: "自己紹介",
			},
			Expected: false,
		},
		"ng_gender_greater_than": {
			Input: &input.UpdateAuthProfile{
				Username:         "test-user",
				Gender:           -1,
				Thumbnail:        "",
				SelfIntroduction: "自己紹介",
			},
			Expected: false,
		},
		"ng_gender_less_than": {
			Input: &input.UpdateAuthProfile{
				Username:         "test-user",
				Gender:           3,
				Thumbnail:        "",
				SelfIntroduction: "自己紹介",
			},
			Expected: false,
		},
		"ng_thumbnail_format": {
			Input: &input.UpdateAuthProfile{
				Username:         "test-user",
				Gender:           0,
				Thumbnail:        "invalida-thumbnail",
				SelfIntroduction: "自己紹介",
			},
			Expected: false,
		},
		"ng_selfintroduction_max": {
			Input: &input.UpdateAuthProfile{
				Username:         "test-user",
				Gender:           0,
				Thumbnail:        "",
				SelfIntroduction: strings.Repeat("x", 257),
			},
			Expected: false,
		},
	}

	for result, tc := range testCases {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		t.Run(result, func(t *testing.T) {
			target := NewAuthRequestValidation()

			got := target.UpdateAuthProfile(tc.Input)
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
