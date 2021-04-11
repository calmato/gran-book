package application

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/calmato/gran-book/api/server/user/internal/application/input"
	"github.com/calmato/gran-book/api/server/user/internal/domain/user"
	mock_validation "github.com/calmato/gran-book/api/server/user/mock/application/validation"
	mock_user "github.com/calmato/gran-book/api/server/user/mock/domain/user"
	"github.com/golang/mock/gomock"
)

func TestAuthApplication_Authentication(t *testing.T) {
	current := time.Now()

	testCases := map[string]struct {
		Expected struct {
			User  *user.User
			Error error
		}
	}{
		"ok": {
			Expected: struct {
				User  *user.User
				Error error
			}{
				User: &user.User{
					ID:               "00000000-0000-0000-0000-000000000000",
					Username:         "test-user",
					Gender:           0,
					Email:            "test-user@calmato.com",
					PhoneNumber:      "000-0000-0000",
					Role:             0,
					ThumbnailURL:     "",
					SelfIntroduction: "",
					LastName:         "テスト",
					FirstName:        "ユーザ",
					LastNameKana:     "てすと",
					FirstNameKana:    "ゆーざ",
					PostalCode:       "000-0000",
					Prefecture:       "東京都",
					City:             "小金井市",
					AddressLine1:     "貫井北町4-1-1",
					AddressLine2:     "",
					InstanceID:       "",
					Activated:        true,
					CreatedAt:        current,
					UpdatedAt:        current,
				},
				Error: nil,
			},
		},
	}

	for result, tc := range testCases {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		arvm := mock_validation.NewMockAuthRequestValidation(ctrl)

		usm := mock_user.NewMockService(ctrl)
		usm.EXPECT().Authentication(ctx).Return(tc.Expected.User.ID, tc.Expected.Error)
		usm.EXPECT().Show(ctx, tc.Expected.User.ID).Return(tc.Expected.User, tc.Expected.Error)

		t.Run(result, func(t *testing.T) {
			target := NewAuthApplication(arvm, usm)

			got, err := target.Authentication(ctx)
			if !reflect.DeepEqual(err, tc.Expected.Error) {
				t.Fatalf("want %#v, but %#v", tc.Expected.Error, err)
				return
			}

			if !reflect.DeepEqual(got, tc.Expected.User) {
				t.Fatalf("want %#v, but %#v", tc.Expected.User, got)
				return
			}
		})
	}
}

func TestAuthApplication_Create(t *testing.T) {
	testCases := map[string]struct {
		Input    *input.CreateAuth
		Expected struct {
			User  *user.User
			Error error
		}
	}{
		"ok": {
			Input: &input.CreateAuth{
				Username:             "test-user",
				Email:                "test-user@calmato.com",
				Password:             "12345678",
				PasswordConfirmation: "12345678",
			},
			Expected: struct {
				User  *user.User
				Error error
			}{
				User: &user.User{
					Username:  "test-user",
					Email:     "test-user@calmato.com",
					Password:  "12345678",
					Gender:    0,
					Role:      0,
					Activated: true,
				},
				Error: nil,
			},
		},
	}

	for result, tc := range testCases {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		arvm := mock_validation.NewMockAuthRequestValidation(ctrl)
		arvm.EXPECT().CreateAuth(tc.Input).Return(nil)

		usm := mock_user.NewMockService(ctrl)
		usm.EXPECT().Validation(ctx, gomock.Any()).Return(tc.Expected.Error)
		usm.EXPECT().Create(ctx, gomock.Any()).Return(tc.Expected.Error)

		t.Run(result, func(t *testing.T) {
			target := NewAuthApplication(arvm, usm)

			got, err := target.Create(ctx, tc.Input)
			if !reflect.DeepEqual(err, tc.Expected.Error) {
				t.Fatalf("want %#v, but %#v", tc.Expected.Error, err)
				return
			}

			if !reflect.DeepEqual(got, tc.Expected.User) {
				t.Fatalf("want %#v, but %#v", tc.Expected.User, got)
				return
			}
		})
	}
}

func TestAuthApplication_UpdateEmail(t *testing.T) {
	testCases := map[string]struct {
		Input    *input.UpdateAuthEmail
		User     *user.User
		Expected error
	}{
		"ok": {
			Input: &input.UpdateAuthEmail{
				Email: "test-user@calmato.com",
			},
			User:     &user.User{},
			Expected: nil,
		},
	}

	for result, tc := range testCases {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		arvm := mock_validation.NewMockAuthRequestValidation(ctrl)
		arvm.EXPECT().UpdateAuthEmail(tc.Input).Return(nil)

		usm := mock_user.NewMockService(ctrl)
		usm.EXPECT().Validation(ctx, tc.User).Return(tc.Expected)
		usm.EXPECT().Update(ctx, tc.User).Return(tc.Expected)

		t.Run(result, func(t *testing.T) {
			target := NewAuthApplication(arvm, usm)

			got := target.UpdateEmail(ctx, tc.Input, tc.User)
			if !reflect.DeepEqual(got, tc.Expected) {
				t.Fatalf("want %#v, but %#v", tc.Expected, got)
				return
			}
		})
	}
}

func TestAuthApplication_UpdatePassword(t *testing.T) {
	testCases := map[string]struct {
		Input    *input.UpdateAuthPassword
		User     *user.User
		Expected error
	}{
		"ok": {
			Input: &input.UpdateAuthPassword{
				Password:             "12345678",
				PasswordConfirmation: "12345678",
			},
			User: &user.User{
				ID: "00000000-0000-0000-0000-000000000000",
			},
			Expected: nil,
		},
	}

	for result, tc := range testCases {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		arvm := mock_validation.NewMockAuthRequestValidation(ctrl)
		arvm.EXPECT().UpdateAuthPassword(tc.Input).Return(nil)

		usm := mock_user.NewMockService(ctrl)
		usm.EXPECT().UpdatePassword(ctx, tc.User.ID, tc.Input.Password).Return(tc.Expected)

		t.Run(result, func(t *testing.T) {
			target := NewAuthApplication(arvm, usm)

			got := target.UpdatePassword(ctx, tc.Input, tc.User)
			if !reflect.DeepEqual(got, tc.Expected) {
				t.Fatalf("want %#v, but %#v", tc.Expected, got)
				return
			}
		})
	}
}

func TestAuthApplication_UpdateProfile(t *testing.T) {
	testCases := map[string]struct {
		Input    *input.UpdateAuthProfile
		User     *user.User
		Expected error
	}{
		"ok": {
			Input: &input.UpdateAuthProfile{
				Username:         "test-user",
				Gender:           0,
				ThumbnailURL:     "",
				SelfIntroduction: "自己紹介",
			},
			User: &user.User{
				ID: "00000000-0000-0000-0000-000000000000",
			},
			Expected: nil,
		},
	}

	for result, tc := range testCases {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		arvm := mock_validation.NewMockAuthRequestValidation(ctrl)
		arvm.EXPECT().UpdateAuthProfile(tc.Input).Return(nil)

		usm := mock_user.NewMockService(ctrl)
		usm.EXPECT().Validation(ctx, gomock.Any()).Return(tc.Expected)
		usm.EXPECT().Update(ctx, tc.User).Return(tc.Expected)

		t.Run(result, func(t *testing.T) {
			target := NewAuthApplication(arvm, usm)

			got := target.UpdateProfile(ctx, tc.Input, tc.User)
			if !reflect.DeepEqual(got, tc.Expected) {
				t.Fatalf("want %#v, but %#v", tc.Expected, got)
				return
			}
		})
	}
}

func TestAuthApplication_UpdateAddress(t *testing.T) {
	testCases := map[string]struct {
		Input    *input.UpdateAuthAddress
		User     *user.User
		Expected error
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
			User: &user.User{
				ID: "00000000-0000-0000-0000-000000000000",
			},
			Expected: nil,
		},
	}

	for result, tc := range testCases {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		arvm := mock_validation.NewMockAuthRequestValidation(ctrl)
		arvm.EXPECT().UpdateAuthAddress(tc.Input).Return(nil)

		usm := mock_user.NewMockService(ctrl)
		usm.EXPECT().Validation(ctx, gomock.Any()).Return(tc.Expected)
		usm.EXPECT().Update(ctx, tc.User).Return(tc.Expected)

		t.Run(result, func(t *testing.T) {
			target := NewAuthApplication(arvm, usm)

			got := target.UpdateAddress(ctx, tc.Input, tc.User)
			if !reflect.DeepEqual(got, tc.Expected) {
				t.Fatalf("want %#v, but %#v", tc.Expected, got)
				return
			}
		})
	}
}

func TestAuthApplication_UploadThumbnail(t *testing.T) {
	testCases := map[string]struct {
		Input    *input.UploadAuthThumbnail
		User     *user.User
		Expected struct {
			ThumbnailURL string
			Error        error
		}
	}{
		"ok": {
			Input: &input.UploadAuthThumbnail{
				Thumbnail: []byte("あいうえお"),
			},
			User: &user.User{
				ID: "00000000-0000-0000-0000-000000000000",
			},
			Expected: struct {
				ThumbnailURL string
				Error        error
			}{
				ThumbnailURL: "https://google.co.jp",
				Error:        nil,
			},
		},
	}

	for result, tc := range testCases {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		arvm := mock_validation.NewMockAuthRequestValidation(ctrl)
		arvm.EXPECT().UploadAuthThumbnail(tc.Input).Return(nil)

		usm := mock_user.NewMockService(ctrl)
		usm.EXPECT().
			UploadThumbnail(ctx, tc.User.ID, tc.Input.Thumbnail).
			Return(tc.Expected.ThumbnailURL, tc.Expected.Error)

		t.Run(result, func(t *testing.T) {
			target := NewAuthApplication(arvm, usm)

			got, err := target.UploadThumbnail(ctx, tc.Input, tc.User)
			if !reflect.DeepEqual(err, tc.Expected.Error) {
				t.Fatalf("want %#v, but %#v", tc.Expected.Error, err)
				return
			}

			if !reflect.DeepEqual(got, tc.Expected.ThumbnailURL) {
				t.Fatalf("want %#v, but %#v", tc.Expected.ThumbnailURL, got)
				return
			}
		})
	}
}
