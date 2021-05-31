package validation

import (
	"strings"
	"testing"

	"github.com/calmato/gran-book/api/server/user/internal/application/input"
	"github.com/golang/mock/gomock"
)

func TestAdminRequestValidation_ListAdmin(t *testing.T) {
	testCases := map[string]struct {
		Input    *input.ListAdmin
		Expected bool
	}{
		"ok": {
			Input: &input.ListAdmin{
				Limit:     100,
				Offset:    0,
				By:        "email",
				Direction: "asc",
			},
			Expected: true,
		},
		"ng_limit_greater_than": {
			Input: &input.ListAdmin{
				Limit:     -1,
				Offset:    0,
				By:        "",
				Direction: "",
			},
			Expected: false,
		},
		"ng_limit_less_than": {
			Input: &input.ListAdmin{
				Limit:     1001,
				Offset:    0,
				By:        "",
				Direction: "",
			},
			Expected: false,
		},
		"ng_offset_greater_than": {
			Input: &input.ListAdmin{
				Limit:     100,
				Offset:    -1,
				By:        "",
				Direction: "",
			},
			Expected: false,
		},
		"ng_by_other_word": {
			Input: &input.ListAdmin{
				Limit:     100,
				Offset:    0,
				By:        "test",
				Direction: "",
			},
			Expected: false,
		},
		"ng_direction_other_word": {
			Input: &input.ListAdmin{
				Limit:     100,
				Offset:    0,
				By:        "test",
				Direction: "",
			},
			Expected: false,
		},
	}

	for result, tc := range testCases {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		t.Run(result, func(t *testing.T) {
			target := NewAdminRequestValidation()

			got := target.ListAdmin(tc.Input)
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

func TestAdminRequestValidation_SearchAdmin(t *testing.T) {
	testCases := map[string]struct {
		Input    *input.SearchAdmin
		Expected bool
	}{
		"ok": {
			Input: &input.SearchAdmin{
				Limit:     100,
				Offset:    0,
				By:        "email",
				Direction: "asc",
				Field:     "email",
				Value:     "test",
			},
			Expected: true,
		},
		"ng_limit_greater_than": {
			Input: &input.SearchAdmin{
				Limit:     -1,
				Offset:    0,
				By:        "",
				Direction: "",
				Field:     "email",
				Value:     "test",
			},
			Expected: false,
		},
		"ng_limit_less_than": {
			Input: &input.SearchAdmin{
				Limit:     1001,
				Offset:    0,
				By:        "",
				Direction: "",
				Field:     "email",
				Value:     "test",
			},
			Expected: false,
		},
		"ng_offset_greater_than": {
			Input: &input.SearchAdmin{
				Limit:     100,
				Offset:    -1,
				By:        "",
				Direction: "",
				Field:     "email",
				Value:     "test",
			},
			Expected: false,
		},
		"ng_by_other_word": {
			Input: &input.SearchAdmin{
				Limit:     100,
				Offset:    0,
				By:        "test",
				Direction: "",
				Field:     "email",
				Value:     "test",
			},
			Expected: false,
		},
		"ng_direction_other_word": {
			Input: &input.SearchAdmin{
				Limit:     100,
				Offset:    0,
				By:        "test",
				Direction: "",
				Field:     "email",
				Value:     "test",
			},
			Expected: false,
		},
		"ng_field_required": {
			Input: &input.SearchAdmin{
				Limit:     100,
				Offset:    0,
				By:        "test",
				Direction: "",
				Field:     "",
				Value:     "test",
			},
			Expected: false,
		},
		"ng_value_required": {
			Input: &input.SearchAdmin{
				Limit:     100,
				Offset:    0,
				By:        "test",
				Direction: "",
				Field:     "email",
				Value:     "",
			},
			Expected: false,
		},
	}

	for result, tc := range testCases {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		t.Run(result, func(t *testing.T) {
			target := NewAdminRequestValidation()

			got := target.SearchAdmin(tc.Input)
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

func TestAdminRequestValidation_UpdateAdminContact(t *testing.T) {
	testCases := map[string]struct {
		Input    *input.UpdateAdminContact
		Expected bool
	}{
		"ok": {
			Input: &input.UpdateAdminContact{
				Email:       "test-user@calmato.com",
				PhoneNumber: "000-0000-0000",
			},
			Expected: true,
		},
		"ng_email_required": {
			Input: &input.UpdateAdminContact{
				Email:       strings.Repeat("x", 256) + "@calmato.com",
				PhoneNumber: "000-0000-0000",
			},
			Expected: false,
		},
		"ng_email_format": {
			Input: &input.UpdateAdminContact{
				Email:       "test-user",
				PhoneNumber: "000-0000-0000",
			},
			Expected: false,
		},
		"ng_phoneNumber_required": {
			Input: &input.UpdateAdminContact{
				Email:       "test-user@calmato.com",
				PhoneNumber: "",
			},
			Expected: false,
		},
		"ng_phoneNumber_max": {
			Input: &input.UpdateAdminContact{
				Email:       "test-user@calmato.com",
				PhoneNumber: strings.Repeat("x", 17),
			},
			Expected: false,
		},
	}

	for result, tc := range testCases {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		t.Run(result, func(t *testing.T) {
			target := NewAdminRequestValidation()

			got := target.UpdateAdminContact(tc.Input)
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
				LastName:      "テスト",
				FirstName:     "ユーザ",
				LastNameKana:  "てすと",
				FirstNameKana: "ゆーざ",
				Role:          1,
				ThumbnailURL:  "https://www.google.co.jp",
			},
			Expected: true,
		},
		"ng_username_required": {
			Input: &input.UpdateAdminProfile{
				Username:      "",
				LastName:      "テスト",
				FirstName:     "ユーザ",
				LastNameKana:  "てすと",
				FirstNameKana: "ゆーざ",
				Role:          1,
				ThumbnailURL:  "https://www.google.co.jp",
			},
			Expected: false,
		},
		"ng_username_max": {
			Input: &input.UpdateAdminProfile{
				Username:      strings.Repeat("x", 33),
				LastName:      "テスト",
				FirstName:     "ユーザ",
				LastNameKana:  "てすと",
				FirstNameKana: "ゆーざ",
				Role:          1,
				ThumbnailURL:  "https://www.google.co.jp",
			},
			Expected: false,
		},
		"ng_lastName_required": {
			Input: &input.UpdateAdminProfile{
				Username:      "test-user",
				LastName:      "",
				FirstName:     "ユーザ",
				LastNameKana:  "てすと",
				FirstNameKana: "ゆーざ",
				Role:          1,
				ThumbnailURL:  "https://www.google.co.jp",
			},
			Expected: false,
		},
		"ng_lastName_max": {
			Input: &input.UpdateAdminProfile{
				Username:      "test-user",
				LastName:      strings.Repeat("あ", 17),
				FirstName:     "ユーザ",
				LastNameKana:  "てすと",
				FirstNameKana: "ゆーざ",
				Role:          1,
				ThumbnailURL:  "https://www.google.co.jp",
			},
			Expected: false,
		},
		"ng_firstName_required": {
			Input: &input.UpdateAdminProfile{
				Username:      "test-user",
				LastName:      "テスト",
				FirstName:     "",
				LastNameKana:  "てすと",
				FirstNameKana: "ゆーざ",
				Role:          1,
				ThumbnailURL:  "https://www.google.co.jp",
			},
			Expected: false,
		},
		"ng_firstName_max": {
			Input: &input.UpdateAdminProfile{
				Username:      "test-user",
				LastName:      "テスト",
				FirstName:     strings.Repeat("あ", 17),
				LastNameKana:  "てすと",
				FirstNameKana: "ゆーざ",
				Role:          1,
				ThumbnailURL:  "https://www.google.co.jp",
			},
			Expected: false,
		},
		"ng_lastNameKana_required": {
			Input: &input.UpdateAdminProfile{
				Username:      "test-user",
				LastName:      "テスト",
				FirstName:     "ユーザ",
				LastNameKana:  "",
				FirstNameKana: "ゆーざ",
				Role:          1,
				ThumbnailURL:  "https://www.google.co.jp",
			},
			Expected: false,
		},
		"ng_lastNameKana_max": {
			Input: &input.UpdateAdminProfile{
				Username:      "test-user",
				LastName:      "テスト",
				FirstName:     "ユーザ",
				LastNameKana:  strings.Repeat("あ", 33),
				FirstNameKana: "ゆーざ",
				Role:          1,
				ThumbnailURL:  "https://www.google.co.jp",
			},
			Expected: false,
		},
		"ng_lastNameKane_format": {
			Input: &input.UpdateAdminProfile{
				Username:      "test-user",
				LastName:      "テスト",
				FirstName:     "ユーザ",
				LastNameKana:  "テスト",
				FirstNameKana: "ゆーざ",
				Role:          1,
				ThumbnailURL:  "https://www.google.co.jp",
			},
			Expected: false,
		},
		"ng_firstNameKana_required": {
			Input: &input.UpdateAdminProfile{
				Username:      "test-user",
				LastName:      "テスト",
				FirstName:     "ユーザ",
				LastNameKana:  "てすと",
				FirstNameKana: "",
				Role:          1,
				ThumbnailURL:  "https://www.google.co.jp",
			},
			Expected: false,
		},
		"ng_firstNameKana_max": {
			Input: &input.UpdateAdminProfile{
				Username:      "test-user",
				LastName:      "テスト",
				FirstName:     "ユーザ",
				LastNameKana:  "てすと",
				FirstNameKana: strings.Repeat("あ", 33),
				Role:          1,
				ThumbnailURL:  "https://www.google.co.jp",
			},
			Expected: false,
		},
		"ng_firstNameKana_format": {
			Input: &input.UpdateAdminProfile{
				Username:      "test-user",
				LastName:      "テスト",
				FirstName:     "ユーザ",
				LastNameKana:  "てすと",
				FirstNameKana: "ユーザ",
				Role:          1,
				ThumbnailURL:  "https://www.google.co.jp",
			},
			Expected: false,
		},
		"ng_role_greater_than_equal": {
			Input: &input.UpdateAdminProfile{
				Username:      "test-user",
				LastName:      "テスト",
				FirstName:     "ユーザ",
				LastNameKana:  "てすと",
				FirstNameKana: "ゆーざ",
				Role:          0,
				ThumbnailURL:  "https://www.google.co.jp",
			},
			Expected: false,
		},
		"ng_role_less_than_equal": {
			Input: &input.UpdateAdminProfile{
				Username:      "test-user",
				LastName:      "テスト",
				FirstName:     "ユーザ",
				LastNameKana:  "てすと",
				FirstNameKana: "ゆーざ",
				Role:          4,
				ThumbnailURL:  "https://www.google.co.jp",
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

func TestAdminRequestValidation_UploadAdminThumbnail(t *testing.T) {
	testCases := map[string]struct {
		Input    *input.UploadAdminThumbnail
		Expected bool
	}{
		"ok": {
			Input: &input.UploadAdminThumbnail{
				Thumbnail: []byte("あいうえお"),
			},
			Expected: true,
		},
		"ng_thumbnail_required": {
			Input: &input.UploadAdminThumbnail{
				Thumbnail: nil,
			},
			Expected: false,
		},
	}

	for result, tc := range testCases {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		t.Run(result, func(t *testing.T) {
			target := NewAdminRequestValidation()

			got := target.UploadAdminThumbnail(tc.Input)
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
