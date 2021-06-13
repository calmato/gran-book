package validation

import (
	"strings"
	"testing"

	"github.com/calmato/gran-book/api/server/user/internal/application/input"
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

	for name, tc := range testCases {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		t.Run(name, func(t *testing.T) {
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

	for name, tc := range testCases {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		t.Run(name, func(t *testing.T) {
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

	for name, tc := range testCases {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		t.Run(name, func(t *testing.T) {
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
				ThumbnailURL:     "",
				SelfIntroduction: "自己紹介",
			},
			Expected: true,
		},
		"ng_username_required": {
			Input: &input.UpdateAuthProfile{
				Username:         "",
				Gender:           0,
				ThumbnailURL:     "",
				SelfIntroduction: "自己紹介",
			},
			Expected: false,
		},
		"ng_username_max": {
			Input: &input.UpdateAuthProfile{
				Username:         strings.Repeat("x", 33),
				Gender:           0,
				ThumbnailURL:     "",
				SelfIntroduction: "自己紹介",
			},
			Expected: false,
		},
		"ng_gender_greater_than": {
			Input: &input.UpdateAuthProfile{
				Username:         "test-user",
				Gender:           -1,
				ThumbnailURL:     "",
				SelfIntroduction: "自己紹介",
			},
			Expected: false,
		},
		"ng_gender_less_than": {
			Input: &input.UpdateAuthProfile{
				Username:         "test-user",
				Gender:           3,
				ThumbnailURL:     "",
				SelfIntroduction: "自己紹介",
			},
			Expected: false,
		},
		"ng_selfintroduction_max": {
			Input: &input.UpdateAuthProfile{
				Username:         "test-user",
				Gender:           0,
				ThumbnailURL:     "",
				SelfIntroduction: strings.Repeat("x", 257),
			},
			Expected: false,
		},
	}

	for name, tc := range testCases {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		t.Run(name, func(t *testing.T) {
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

func TestAuthRequestValidation_UpdateAuthAddress(t *testing.T) {
	testCases := map[string]struct {
		Input    *input.UpdateAuthAddress
		Expected bool
	}{
		"ok": {
			Input: &input.UpdateAuthAddress{
				LastName:      "テスト",
				FirstName:     "ユーザ",
				LastNameKana:  "てすと",
				FirstNameKana: "ゆーざ",
				PhoneNumber:   "000-0000-0000",
				PostalCode:    "000-0000",
				Prefecture:    "東京都",
				City:          "小金井市",
				AddressLine1:  "貫井北町4-1-1",
				AddressLine2:  "",
			},
			Expected: true,
		},
		"ng_lastName_required": {
			Input: &input.UpdateAuthAddress{
				LastName:      "",
				FirstName:     "ユーザ",
				LastNameKana:  "てすと",
				FirstNameKana: "ゆーざ",
				PhoneNumber:   "000-0000-0000",
				PostalCode:    "000-0000",
				Prefecture:    "東京都",
				City:          "小金井市",
			},
			Expected: false,
		},
		"ng_lastName_max": {
			Input: &input.UpdateAuthAddress{
				LastName:      strings.Repeat("x", 17),
				FirstName:     "ユーザ",
				LastNameKana:  "てすと",
				FirstNameKana: "ゆーざ",
				PhoneNumber:   "000-0000-0000",
				PostalCode:    "000-0000",
				Prefecture:    "東京都",
				City:          "小金井市",
			},
			Expected: false,
		},
		"ng_firstName_required": {
			Input: &input.UpdateAuthAddress{
				LastName:      "テスト",
				FirstName:     "",
				LastNameKana:  "てすと",
				FirstNameKana: "ゆーざ",
				PhoneNumber:   "000-0000-0000",
				PostalCode:    "000-0000",
				Prefecture:    "東京都",
				City:          "小金井市",
				AddressLine1:  "貫井北町4-1-1",
				AddressLine2:  "",
			},
			Expected: false,
		},
		"ng_firstName_max": {
			Input: &input.UpdateAuthAddress{
				LastName:      "テスト",
				FirstName:     strings.Repeat("x", 17),
				LastNameKana:  "てすと",
				FirstNameKana: "ゆーざ",
				PhoneNumber:   "000-0000-0000",
				PostalCode:    "000-0000",
				Prefecture:    "東京都",
				City:          "小金井市",
				AddressLine1:  "貫井北町4-1-1",
				AddressLine2:  "",
			},
			Expected: false,
		},
		"ng_lastNameKana_required": {
			Input: &input.UpdateAuthAddress{
				LastName:      "テスト",
				FirstName:     "ユーザ",
				LastNameKana:  "",
				FirstNameKana: "ゆーざ",
				PhoneNumber:   "000-0000-0000",
				PostalCode:    "000-0000",
				Prefecture:    "東京都",
				City:          "小金井市",
				AddressLine1:  "貫井北町4-1-1",
				AddressLine2:  "",
			},
			Expected: false,
		},
		"ng_lastNameKana_max": {
			Input: &input.UpdateAuthAddress{
				LastName:      "テスト",
				FirstName:     "ユーザ",
				LastNameKana:  strings.Repeat("あ", 33),
				FirstNameKana: "ゆーざ",
				PhoneNumber:   "000-0000-0000",
				PostalCode:    "000-0000",
				Prefecture:    "東京都",
				City:          "小金井市",
				AddressLine1:  "貫井北町4-1-1",
				AddressLine2:  "",
			},
			Expected: false,
		},
		"ng_lastNameKane_format": {
			Input: &input.UpdateAuthAddress{
				LastName:      "テスト",
				FirstName:     "ユーザ",
				LastNameKana:  "Test",
				FirstNameKana: "ゆーざ",
				PhoneNumber:   "000-0000-0000",
				PostalCode:    "000-0000",
				Prefecture:    "東京都",
				City:          "小金井市",
				AddressLine1:  "貫井北町4-1-1",
				AddressLine2:  "",
			},
			Expected: false,
		},
		"ng_firstNameKana_required": {
			Input: &input.UpdateAuthAddress{
				LastName:      "テスト",
				FirstName:     "ユーザ",
				LastNameKana:  "てすと",
				FirstNameKana: "",
				PhoneNumber:   "000-0000-0000",
				PostalCode:    "000-0000",
				Prefecture:    "東京都",
				City:          "小金井市",
				AddressLine1:  "貫井北町4-1-1",
				AddressLine2:  "",
			},
			Expected: false,
		},
		"ng_firstNameKana_max": {
			Input: &input.UpdateAuthAddress{
				LastName:      "テスト",
				FirstName:     "ユーザ",
				LastNameKana:  "てすと",
				FirstNameKana: strings.Repeat("あ", 33),
				PhoneNumber:   "000-0000-0000",
				PostalCode:    "000-0000",
				Prefecture:    "東京都",
				City:          "小金井市",
				AddressLine1:  "貫井北町4-1-1",
				AddressLine2:  "",
			},
			Expected: false,
		},
		"ng_firstNameKana_format": {
			Input: &input.UpdateAuthAddress{
				LastName:      "テスト",
				FirstName:     "ユーザ",
				LastNameKana:  "てすと",
				FirstNameKana: "User",
				PhoneNumber:   "000-0000-0000",
				PostalCode:    "000-0000",
				Prefecture:    "東京都",
				City:          "小金井市",
				AddressLine1:  "貫井北町4-1-1",
				AddressLine2:  "",
			},
			Expected: false,
		},
		"ng_phoneNumber_required": {
			Input: &input.UpdateAuthAddress{
				LastName:      "テスト",
				FirstName:     "ユーザ",
				LastNameKana:  "てすと",
				FirstNameKana: "ゆーざ",
				PhoneNumber:   "",
				PostalCode:    "000-0000",
				Prefecture:    "東京都",
				City:          "小金井市",
				AddressLine1:  "貫井北町4-1-1",
				AddressLine2:  "",
			},
			Expected: false,
		},
		"ng_phoneNumber_max": {
			Input: &input.UpdateAuthAddress{
				LastName:      "テスト",
				FirstName:     "ユーザ",
				LastNameKana:  "てすと",
				FirstNameKana: "ゆーざ",
				PhoneNumber:   strings.Repeat("0", 17),
				PostalCode:    "000-0000",
				Prefecture:    "東京都",
				City:          "小金井市",
				AddressLine1:  "貫井北町4-1-1",
				AddressLine2:  "",
			},
			Expected: false,
		},
		"ng_postalCode_required": {
			Input: &input.UpdateAuthAddress{
				LastName:      "テスト",
				FirstName:     "ユーザ",
				LastNameKana:  "てすと",
				FirstNameKana: "ゆーざ",
				PhoneNumber:   "000-0000-0000",
				PostalCode:    "",
				Prefecture:    "東京都",
				City:          "小金井市",
				AddressLine1:  "貫井北町4-1-1",
				AddressLine2:  "",
			},
			Expected: false,
		},
		"ng_postalCode_max": {
			Input: &input.UpdateAuthAddress{
				LastName:      "テスト",
				FirstName:     "ユーザ",
				LastNameKana:  "てすと",
				FirstNameKana: "ゆーざ",
				PhoneNumber:   "000-0000-0000",
				PostalCode:    strings.Repeat("0", 17),
				Prefecture:    "東京都",
				City:          "小金井市",
				AddressLine1:  "貫井北町4-1-1",
				AddressLine2:  "",
			},
			Expected: false,
		},
		"ng_prefecture_required": {
			Input: &input.UpdateAuthAddress{
				LastName:      "テスト",
				FirstName:     "ユーザ",
				LastNameKana:  "てすと",
				FirstNameKana: "ゆーざ",
				PhoneNumber:   "000-0000-0000",
				PostalCode:    "000-0000",
				Prefecture:    "",
				City:          "小金井市",
				AddressLine1:  "貫井北町4-1-1",
				AddressLine2:  "",
			},
			Expected: false,
		},
		"ng_prefecture_max": {
			Input: &input.UpdateAuthAddress{
				LastName:      "テスト",
				FirstName:     "ユーザ",
				LastNameKana:  "てすと",
				FirstNameKana: "ゆーざ",
				PhoneNumber:   "000-0000-0000",
				PostalCode:    "000-0000",
				Prefecture:    strings.Repeat("x", 33),
				City:          "小金井市",
				AddressLine1:  "貫井北町4-1-1",
				AddressLine2:  "",
			},
			Expected: false,
		},
		"ng_city_required": {
			Input: &input.UpdateAuthAddress{
				LastName:      "テスト",
				FirstName:     "ユーザ",
				LastNameKana:  "てすと",
				FirstNameKana: "ゆーざ",
				PhoneNumber:   "000-0000-0000",
				PostalCode:    "000-0000",
				Prefecture:    "東京都",
				City:          "",
				AddressLine1:  "貫井北町4-1-1",
				AddressLine2:  "",
			},
			Expected: false,
		},
		"ng_city_max": {
			Input: &input.UpdateAuthAddress{
				LastName:      "テスト",
				FirstName:     "ユーザ",
				LastNameKana:  "てすと",
				FirstNameKana: "ゆーざ",
				PhoneNumber:   "000-0000-0000",
				PostalCode:    "000-0000",
				Prefecture:    "東京都",
				City:          strings.Repeat("x", 33),
				AddressLine1:  "貫井北町4-1-1",
				AddressLine2:  "",
			},
			Expected: false,
		},
		"ng_addressLine1_required": {
			Input: &input.UpdateAuthAddress{
				LastName:      "テスト",
				FirstName:     "ユーザ",
				LastNameKana:  "てすと",
				FirstNameKana: "ゆーざ",
				PhoneNumber:   "000-0000-0000",
				PostalCode:    "000-0000",
				Prefecture:    "東京都",
				City:          "小金井市",
				AddressLine1:  "",
				AddressLine2:  "",
			},
			Expected: false,
		},
		"ng_addressLine1_max": {
			Input: &input.UpdateAuthAddress{
				LastName:      "テスト",
				FirstName:     "ユーザ",
				LastNameKana:  "てすと",
				FirstNameKana: "ゆーざ",
				PhoneNumber:   "000-0000-0000",
				PostalCode:    "000-0000",
				Prefecture:    "東京都",
				City:          "小金井市",
				AddressLine1:  strings.Repeat("x", 65),
				AddressLine2:  "",
			},
			Expected: false,
		},
		"ng_addressLine2_max": {
			Input: &input.UpdateAuthAddress{
				LastName:      "テスト",
				FirstName:     "ユーザ",
				LastNameKana:  "てすと",
				FirstNameKana: "ゆーざ",
				PhoneNumber:   "000-0000-0000",
				PostalCode:    "000-0000",
				Prefecture:    "東京都",
				City:          "小金井市",
				AddressLine1:  "貫井北町4-1-1",
				AddressLine2:  strings.Repeat("x", 65),
			},
			Expected: false,
		},
	}

	for name, tc := range testCases {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		t.Run(name, func(t *testing.T) {
			target := NewAuthRequestValidation()

			got := target.UpdateAuthAddress(tc.Input)
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

func TestAuthRequestValidation_UploadAuthThumbnail(t *testing.T) {
	testCases := map[string]struct {
		Input    *input.UploadAuthThumbnail
		Expected bool
	}{
		"ok": {
			Input: &input.UploadAuthThumbnail{
				Thumbnail: []byte("あいうえお"),
			},
			Expected: true,
		},
		"ng_thumbnail_required": {
			Input: &input.UploadAuthThumbnail{
				Thumbnail: nil,
			},
			Expected: false,
		},
	}

	for name, tc := range testCases {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		t.Run(name, func(t *testing.T) {
			target := NewAuthRequestValidation()

			got := target.UploadAuthThumbnail(tc.Input)
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

func TestAuthRequestValidation_RegisterDevice(t *testing.T) {
	testCases := map[string]struct {
		Input    *input.RegisterAuthDevice
		Expected bool
	}{
		"ok": {
			Input: &input.RegisterAuthDevice{
				InstanceID: "cTP0f6Y_Q26VG9TbTjReZz:APA91bG6Ns9A5DsXaMcImyyNImS4VD",
			},
			Expected: true,
		},
		"ng_instanceId_required": {
			Input: &input.RegisterAuthDevice{
				InstanceID: "",
			},
			Expected: false,
		},
	}

	for name, tc := range testCases {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		t.Run(name, func(t *testing.T) {
			target := NewAuthRequestValidation()

			got := target.RegisterAuthDevice(tc.Input)
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
