package validation

import (
	"strings"
	"testing"

	"github.com/calmato/gran-book/api/server/user/internal/application/input"
	"github.com/golang/mock/gomock"
)

func TestAdminRequestValidation_CreateAdmin(t *testing.T) {
	testCases := map[string]struct {
		Input    *input.CreateAdmin
		Expected bool
	}{
		"ok": {
			Input: &input.CreateAdmin{
				Username:             "test-user",
				Email:                "test-user@calmato.com",
				Password:             "12345678",
				PasswordConfirmation: "12345678",
				Role:                 1,
				LastName:             "テスト",
				FirstName:            "ユーザ",
				LastNameKana:         "てすと",
				FirstNameKana:        "ゆーざ",
			},
			Expected: true,
		},
		"ng_username_required": {
			Input: &input.CreateAdmin{
				Username:             "",
				Email:                "test-user@calmato.com",
				Password:             "12345678",
				PasswordConfirmation: "12345678",
				Role:                 1,
				LastName:             "テスト",
				FirstName:            "ユーザ",
				LastNameKana:         "てすと",
				FirstNameKana:        "ゆーざ",
			},
			Expected: false,
		},
		"ng_username_max": {
			Input: &input.CreateAdmin{
				Username:             strings.Repeat("x", 33),
				Email:                "test-user@calmato.com",
				Password:             "12345678",
				PasswordConfirmation: "12345678",
				Role:                 1,
				LastName:             "テスト",
				FirstName:            "ユーザ",
				LastNameKana:         "てすと",
				FirstNameKana:        "ゆーざ",
			},
			Expected: false,
		},
		"ng_email_required": {
			Input: &input.CreateAdmin{
				Username:             "test-user",
				Email:                "",
				Password:             "12345678",
				PasswordConfirmation: "12345678",
				Role:                 1,
				LastName:             "テスト",
				FirstName:            "ユーザ",
				LastNameKana:         "てすと",
				FirstNameKana:        "ゆーざ",
			},
			Expected: false,
		},
		"ng_email_max": {
			Input: &input.CreateAdmin{
				Username:             "test-user",
				Email:                strings.Repeat("x", 256) + "@calmato.com",
				Password:             "12345678",
				PasswordConfirmation: "12345678",
				Role:                 1,
				LastName:             "テスト",
				FirstName:            "ユーザ",
				LastNameKana:         "てすと",
				FirstNameKana:        "ゆーざ",
			},
			Expected: false,
		},
		"ng_email_format": {
			Input: &input.CreateAdmin{
				Username:             "test-user",
				Email:                "test-user",
				Password:             "12345678",
				PasswordConfirmation: "12345678",
				Role:                 1,
				LastName:             "テスト",
				FirstName:            "ユーザ",
				LastNameKana:         "てすと",
				FirstNameKana:        "ゆーざ",
			},
			Expected: false,
		},
		"ng_password_required": {
			Input: &input.CreateAdmin{
				Username:             "test-user",
				Email:                "test-user@calmato.com",
				Password:             "",
				PasswordConfirmation: "12345678",
				Role:                 1,
				LastName:             "テスト",
				FirstName:            "ユーザ",
				LastNameKana:         "てすと",
				FirstNameKana:        "ゆーざ",
			},
			Expected: false,
		},
		"ng_password_min": {
			Input: &input.CreateAdmin{
				Username:             "test-user",
				Email:                "test-user@calmato.com",
				Password:             strings.Repeat("x", 5),
				PasswordConfirmation: strings.Repeat("x", 5),
				Role:                 1,
				LastName:             "テスト",
				FirstName:            "ユーザ",
				LastNameKana:         "てすと",
				FirstNameKana:        "ゆーざ",
			},
			Expected: false,
		},
		"ng_password_max": {
			Input: &input.CreateAdmin{
				Username:             "test-user",
				Email:                "test-user@calmato.com",
				Password:             strings.Repeat("x", 33),
				PasswordConfirmation: strings.Repeat("x", 33),
				Role:                 1,
				LastName:             "テスト",
				FirstName:            "ユーザ",
				LastNameKana:         "てすと",
				FirstNameKana:        "ゆーざ",
			},
			Expected: false,
		},
		"ng_password_format": {
			Input: &input.CreateAdmin{
				Username:             "test-user",
				Email:                "test-user@calmato.com",
				Password:             "１２３４５６７８",
				PasswordConfirmation: "１２３４５６７８",
				Role:                 1,
				LastName:             "テスト",
				FirstName:            "ユーザ",
				LastNameKana:         "てすと",
				FirstNameKana:        "ゆーざ",
			},
			Expected: false,
		},
		"ng_passwordConfirmation_required": {
			Input: &input.CreateAdmin{
				Username:             "test-user",
				Email:                "test-user@calmato.com",
				Password:             "12345678",
				PasswordConfirmation: "",
				Role:                 1,
				LastName:             "テスト",
				FirstName:            "ユーザ",
				LastNameKana:         "てすと",
				FirstNameKana:        "ゆーざ",
			},
			Expected: false,
		},
		"ng_passwordConfirmation_equal": {
			Input: &input.CreateAdmin{
				Username:             "test-user",
				Email:                "test-user@calmato.com",
				Password:             "12345678",
				PasswordConfirmation: "1234567",
				Role:                 1,
				LastName:             "テスト",
				FirstName:            "ユーザ",
				LastNameKana:         "てすと",
				FirstNameKana:        "ゆーざ",
			},
			Expected: false,
		},
		"ng_role_greater_than": {
			Input: &input.CreateAdmin{
				Username:             "test-user",
				Email:                "test-user@calmato.com",
				Password:             "12345678",
				PasswordConfirmation: "12345678",
				Role:                 0,
				LastName:             "テスト",
				FirstName:            "ユーザ",
				LastNameKana:         "てすと",
				FirstNameKana:        "ゆーざ",
			},
			Expected: false,
		},
		"ng_role_less_than": {
			Input: &input.CreateAdmin{
				Username:             "test-user",
				Email:                "test-user@calmato.com",
				Password:             "12345678",
				PasswordConfirmation: "12345678",
				Role:                 4,
				LastName:             "テスト",
				FirstName:            "ユーザ",
				LastNameKana:         "てすと",
				FirstNameKana:        "ゆーざ",
			},
			Expected: false,
		},
		"ng_lastName_required": {
			Input: &input.CreateAdmin{
				Username:             "test-user",
				Email:                "test-user@calmato.com",
				Password:             "12345678",
				PasswordConfirmation: "12345678",
				Role:                 1,
				LastName:             "",
				FirstName:            "ユーザ",
				LastNameKana:         "てすと",
				FirstNameKana:        "ゆーざ",
			},
			Expected: false,
		},
		"ng_lastName_max": {
			Input: &input.CreateAdmin{
				Username:             "test-user",
				Email:                "test-user@calmato.com",
				Password:             "12345678",
				PasswordConfirmation: "12345678",
				Role:                 1,
				LastName:             strings.Repeat("あ", 17),
				FirstName:            "ユーザ",
				LastNameKana:         "てすと",
				FirstNameKana:        "ゆーざ",
			},
			Expected: false,
		},
		"ng_firstName_required": {
			Input: &input.CreateAdmin{
				Username:             "test-user",
				Email:                "test-user@calmato.com",
				Password:             "12345678",
				PasswordConfirmation: "12345678",
				Role:                 1,
				LastName:             "テスト",
				FirstName:            "",
				LastNameKana:         "てすと",
				FirstNameKana:        "ゆーざ",
			},
			Expected: false,
		},
		"ng_firstName_max": {
			Input: &input.CreateAdmin{
				Username:             "test-user",
				Email:                "test-user@calmato.com",
				Password:             "12345678",
				PasswordConfirmation: "12345678",
				Role:                 1,
				LastName:             "テスト",
				FirstName:            strings.Repeat("あ", 17),
				LastNameKana:         "てすと",
				FirstNameKana:        "ゆーざ",
			},
			Expected: false,
		},
		"ng_lastNameKana_required": {
			Input: &input.CreateAdmin{
				Username:             "test-user",
				Email:                "test-user@calmato.com",
				Password:             "12345678",
				PasswordConfirmation: "12345678",
				Role:                 1,
				LastName:             "テスト",
				FirstName:            "ユーザ",
				LastNameKana:         "",
				FirstNameKana:        "ゆーざ",
			},
			Expected: false,
		},
		"ng_lastNameKana_max": {
			Input: &input.CreateAdmin{
				Username:             "test-user",
				Email:                "test-user@calmato.com",
				Password:             "12345678",
				PasswordConfirmation: "12345678",
				Role:                 1,
				LastName:             "テスト",
				FirstName:            "ユーザ",
				LastNameKana:         strings.Repeat("あ", 33),
				FirstNameKana:        "ゆーざ",
			},
			Expected: false,
		},
		"ng_lastNameKane_format": {
			Input: &input.CreateAdmin{
				Username:             "test-user",
				Email:                "test-user@calmato.com",
				Password:             "12345678",
				PasswordConfirmation: "12345678",
				Role:                 1,
				LastName:             "テスト",
				FirstName:            "ユーザ",
				LastNameKana:         "テスト",
				FirstNameKana:        "ゆーざ",
			},
			Expected: false,
		},
		"ng_firstNameKana_required": {
			Input: &input.CreateAdmin{
				Username:             "test-user",
				Email:                "test-user@calmato.com",
				Password:             "12345678",
				PasswordConfirmation: "12345678",
				Role:                 1,
				LastName:             "テスト",
				FirstName:            "ユーザ",
				LastNameKana:         "てすと",
				FirstNameKana:        "",
			},
			Expected: false,
		},
		"ng_firstNameKana_max": {
			Input: &input.CreateAdmin{
				Username:             "test-user",
				Email:                "test-user@calmato.com",
				Password:             "12345678",
				PasswordConfirmation: "12345678",
				Role:                 1,
				LastName:             "テスト",
				FirstName:            "ユーザ",
				LastNameKana:         "てすと",
				FirstNameKana:        strings.Repeat("あ", 33),
			},
			Expected: false,
		},
		"ng_firstNameKana_format": {
			Input: &input.CreateAdmin{
				Username:             "test-user",
				Email:                "test-user@calmato.com",
				Password:             "12345678",
				PasswordConfirmation: "12345678",
				Role:                 1,
				LastName:             "テスト",
				FirstName:            "ユーザ",
				LastNameKana:         "てすと",
				FirstNameKana:        "ユーザ",
			},
			Expected: false,
		},
	}

	for result, tc := range testCases {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		t.Run(result, func(t *testing.T) {
			target := NewAdminRequestValidation()

			got := target.CreateAdmin(tc.Input)
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

func TestAdminRequestValidation_UpdateAdminRole(t *testing.T) {
	testCases := map[string]struct {
		Input    *input.UpdateAdminRole
		Expected bool
	}{
		"ok": {
			Input: &input.UpdateAdminRole{
				Role: 1,
			},
			Expected: true,
		},
		"ng_role_greater_than": {
			Input: &input.UpdateAdminRole{
				Role: 0,
			},
			Expected: false,
		},
		"ng_role_less_than": {
			Input: &input.UpdateAdminRole{
				Role: 4,
			},
			Expected: false,
		},
	}

	for result, tc := range testCases {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		t.Run(result, func(t *testing.T) {
			target := NewAdminRequestValidation()

			got := target.UpdateAdminRole(tc.Input)
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

func TestAdminRequestValidation_UpdateAdminPassword(t *testing.T) {
	testCases := map[string]struct {
		Input    *input.UpdateAdminPassword
		Expected bool
	}{
		"ok": {
			Input: &input.UpdateAdminPassword{
				Password:             "12345678",
				PasswordConfirmation: "12345678",
			},
			Expected: true,
		},
		"ng_password_required": {
			Input: &input.UpdateAdminPassword{
				Password:             "",
				PasswordConfirmation: "12345678",
			},
			Expected: false,
		},
		"ng_password_min": {
			Input: &input.UpdateAdminPassword{
				Password:             strings.Repeat("x", 5),
				PasswordConfirmation: strings.Repeat("x", 5),
			},
			Expected: false,
		},
		"ng_password_max": {
			Input: &input.UpdateAdminPassword{
				Password:             strings.Repeat("x", 33),
				PasswordConfirmation: strings.Repeat("x", 33),
			},
			Expected: false,
		},
		"ng_password_format": {
			Input: &input.UpdateAdminPassword{
				Password:             "１２３４５６７８",
				PasswordConfirmation: "１２３４５６７８",
			},
			Expected: false,
		},
		"ng_passwordConfirmation_required": {
			Input: &input.UpdateAdminPassword{
				Password:             "12345678",
				PasswordConfirmation: "",
			},
			Expected: false,
		},
		"ng_passwordConfirmation_equal": {
			Input: &input.UpdateAdminPassword{
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
			target := NewAdminRequestValidation()

			got := target.UpdateAdminPassword(tc.Input)
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

func TestAdminRequestValidation_UpdateAdminProfile(t *testing.T) {
	testCases := map[string]struct {
		Input    *input.UpdateAdminProfile
		Expected bool
	}{
		"ok": {
			Input: &input.UpdateAdminProfile{
				Username:      "test-user",
				Email:         "test-user@calmato.com",
				LastName:      "テスト",
				FirstName:     "ユーザ",
				LastNameKana:  "てすと",
				FirstNameKana: "ゆーざ",
			},
			Expected: true,
		},
		"ng_username_required": {
			Input: &input.UpdateAdminProfile{
				Username:      "",
				Email:         "test-user@calmato.com",
				LastName:      "テスト",
				FirstName:     "ユーザ",
				LastNameKana:  "てすと",
				FirstNameKana: "ゆーざ",
			},
			Expected: false,
		},
		"ng_username_max": {
			Input: &input.UpdateAdminProfile{
				Username:      strings.Repeat("x", 33),
				Email:         "test-user@calmato.com",
				LastName:      "テスト",
				FirstName:     "ユーザ",
				LastNameKana:  "てすと",
				FirstNameKana: "ゆーざ",
			},
			Expected: false,
		},
		"ng_email_required": {
			Input: &input.UpdateAdminProfile{
				Username:      "test-user",
				Email:         "",
				LastName:      "テスト",
				FirstName:     "ユーザ",
				LastNameKana:  "てすと",
				FirstNameKana: "ゆーざ",
			},
			Expected: false,
		},
		"ng_email_max": {
			Input: &input.UpdateAdminProfile{
				Username:      "test-user",
				Email:         strings.Repeat("x", 256) + "@calmato.com",
				LastName:      "テスト",
				FirstName:     "ユーザ",
				LastNameKana:  "てすと",
				FirstNameKana: "ゆーざ",
			},
			Expected: false,
		},
		"ng_email_format": {
			Input: &input.UpdateAdminProfile{
				Username:      "test-user",
				Email:         "test-user",
				LastName:      "テスト",
				FirstName:     "ユーザ",
				LastNameKana:  "てすと",
				FirstNameKana: "ゆーざ",
			},
			Expected: false,
		},
		"ng_lastName_required": {
			Input: &input.UpdateAdminProfile{
				Username:      "test-user",
				Email:         "test-user@calmato.com",
				LastName:      "",
				FirstName:     "ユーザ",
				LastNameKana:  "てすと",
				FirstNameKana: "ゆーざ",
			},
			Expected: false,
		},
		"ng_lastName_max": {
			Input: &input.UpdateAdminProfile{
				Username:      "test-user",
				Email:         "test-user@calmato.com",
				LastName:      strings.Repeat("あ", 17),
				FirstName:     "ユーザ",
				LastNameKana:  "てすと",
				FirstNameKana: "ゆーざ",
			},
			Expected: false,
		},
		"ng_firstName_required": {
			Input: &input.UpdateAdminProfile{
				Username:      "test-user",
				Email:         "test-user@calmato.com",
				LastName:      "テスト",
				FirstName:     "",
				LastNameKana:  "てすと",
				FirstNameKana: "ゆーざ",
			},
			Expected: false,
		},
		"ng_firstName_max": {
			Input: &input.UpdateAdminProfile{
				Username:      "test-user",
				Email:         "test-user@calmato.com",
				LastName:      "テスト",
				FirstName:     strings.Repeat("あ", 17),
				LastNameKana:  "てすと",
				FirstNameKana: "ゆーざ",
			},
			Expected: false,
		},
		"ng_lastNameKana_required": {
			Input: &input.UpdateAdminProfile{
				Username:      "test-user",
				Email:         "test-user@calmato.com",
				LastName:      "テスト",
				FirstName:     "ユーザ",
				LastNameKana:  "",
				FirstNameKana: "ゆーざ",
			},
			Expected: false,
		},
		"ng_lastNameKana_max": {
			Input: &input.UpdateAdminProfile{
				Username:      "test-user",
				Email:         "test-user@calmato.com",
				LastName:      "テスト",
				FirstName:     "ユーザ",
				LastNameKana:  strings.Repeat("あ", 33),
				FirstNameKana: "ゆーざ",
			},
			Expected: false,
		},
		"ng_lastNameKane_format": {
			Input: &input.UpdateAdminProfile{
				Username:      "test-user",
				Email:         "test-user@calmato.com",
				LastName:      "テスト",
				FirstName:     "ユーザ",
				LastNameKana:  "テスト",
				FirstNameKana: "ゆーざ",
			},
			Expected: false,
		},
		"ng_firstNameKana_required": {
			Input: &input.UpdateAdminProfile{
				Username:      "test-user",
				Email:         "test-user@calmato.com",
				LastName:      "テスト",
				FirstName:     "ユーザ",
				LastNameKana:  "てすと",
				FirstNameKana: "",
			},
			Expected: false,
		},
		"ng_firstNameKana_max": {
			Input: &input.UpdateAdminProfile{
				Username:      "test-user",
				Email:         "test-user@calmato.com",
				LastName:      "テスト",
				FirstName:     "ユーザ",
				LastNameKana:  "てすと",
				FirstNameKana: strings.Repeat("あ", 33),
			},
			Expected: false,
		},
		"ng_firstNameKana_format": {
			Input: &input.UpdateAdminProfile{
				Username:      "test-user",
				Email:         "test-user@calmato.com",
				LastName:      "テスト",
				FirstName:     "ユーザ",
				LastNameKana:  "てすと",
				FirstNameKana: "ユーザ",
			},
			Expected: false,
		},
	}

	for result, tc := range testCases {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		t.Run(result, func(t *testing.T) {
			target := NewAdminRequestValidation()

			got := target.UpdateAdminProfile(tc.Input)
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
