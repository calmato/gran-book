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
	type want struct {
		user *user.User
		err  error
	}

	current := time.Now().Local()
	testCases := map[string]struct {
		want want
	}{
		"ok": {
			want: want{
				user: &user.User{
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
					CreatedAt:        current,
					UpdatedAt:        current,
				},
				err: nil,
			},
		},
	}

	for name, tc := range testCases {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		arvm := mock_validation.NewMockAuthRequestValidation(ctrl)

		usm := mock_user.NewMockService(ctrl)
		usm.EXPECT().Authentication(ctx).Return(tc.want.user.ID, nil)
		usm.EXPECT().Show(ctx, tc.want.user.ID).Return(tc.want.user, tc.want.err)

		t.Run(name, func(t *testing.T) {
			target := NewAuthApplication(arvm, usm)

			user, err := target.Authentication(ctx)
			if !reflect.DeepEqual(err, tc.want.err) {
				t.Fatalf("want %#v, but %#v", tc.want.err, err)
				return
			}

			if !reflect.DeepEqual(user, tc.want.user) {
				t.Fatalf("want %#v, but %#v", tc.want.user, user)
				return
			}
		})
	}
}

func TestAuthApplication_Create(t *testing.T) {
	type args struct {
		input *input.CreateAuth
	}
	type want struct {
		user *user.User
		err  error
	}

	testCases := map[string]struct {
		args args
		want want
	}{
		"ok": {
			args: args{
				input: &input.CreateAuth{
					Username:             "test-user",
					Email:                "test-user@calmato.com",
					Password:             "12345678",
					PasswordConfirmation: "12345678",
				},
			},
			want: want{
				user: &user.User{
					Username: "test-user",
					Email:    "test-user@calmato.com",
					Password: "12345678",
					Gender:   0,
					Role:     0,
				},
				err: nil,
			},
		},
	}

	for name, tc := range testCases {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		arvm := mock_validation.NewMockAuthRequestValidation(ctrl)
		arvm.EXPECT().CreateAuth(tc.args.input).Return(nil)

		usm := mock_user.NewMockService(ctrl)
		usm.EXPECT().Validation(ctx, gomock.Any()).Return(nil)
		usm.EXPECT().Create(ctx, gomock.Any()).Return(tc.want.err)

		t.Run(name, func(t *testing.T) {
			target := NewAuthApplication(arvm, usm)

			user, err := target.Create(ctx, tc.args.input)
			if !reflect.DeepEqual(err, tc.want.err) {
				t.Fatalf("want %#v, but %#v", tc.want.err, err)
				return
			}

			if !reflect.DeepEqual(user, tc.want.user) {
				t.Fatalf("want %#v, but %#v", tc.want.user, user)
				return
			}
		})
	}
}

func TestAuthApplication_UpdateEmail(t *testing.T) {
	type args struct {
		input *input.UpdateAuthEmail
		user  *user.User
	}

	testCases := map[string]struct {
		args args
		want error
	}{
		"ok": {
			args: args{
				input: &input.UpdateAuthEmail{
					Email: "test-user@calmato.com",
				},
				user: &user.User{},
			},
			want: nil,
		},
	}

	for name, tc := range testCases {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		arvm := mock_validation.NewMockAuthRequestValidation(ctrl)
		arvm.EXPECT().UpdateAuthEmail(tc.args.input).Return(nil)

		usm := mock_user.NewMockService(ctrl)
		usm.EXPECT().Validation(ctx, tc.args.user).Return(nil)
		usm.EXPECT().Update(ctx, tc.args.user).Return(tc.want)

		t.Run(name, func(t *testing.T) {
			target := NewAuthApplication(arvm, usm)

			got := target.UpdateEmail(ctx, tc.args.input, tc.args.user)
			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("want %#v, but %#v", tc.want, got)
				return
			}

			if tc.args.user.Email != tc.args.input.Email {
				t.Fatalf("want %#v, but %#v", tc.args.input.Email, tc.args.user.Email)
				return
			}
		})
	}
}

func TestAuthApplication_UpdatePassword(t *testing.T) {
	type args struct {
		input *input.UpdateAuthPassword
		user  *user.User
	}

	testCases := map[string]struct {
		args args
		want error
	}{
		"ok": {
			args: args{
				input: &input.UpdateAuthPassword{
					Password:             "12345678",
					PasswordConfirmation: "12345678",
				},
				user: &user.User{
					ID: "00000000-0000-0000-0000-000000000000",
				},
			},
			want: nil,
		},
	}

	for name, tc := range testCases {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		arvm := mock_validation.NewMockAuthRequestValidation(ctrl)
		arvm.EXPECT().UpdateAuthPassword(tc.args.input).Return(nil)

		usm := mock_user.NewMockService(ctrl)
		usm.EXPECT().UpdatePassword(ctx, tc.args.user.ID, tc.args.input.Password).Return(tc.want)

		t.Run(name, func(t *testing.T) {
			target := NewAuthApplication(arvm, usm)

			got := target.UpdatePassword(ctx, tc.args.input, tc.args.user)
			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("want %#v, but %#v", tc.want, got)
				return
			}
		})
	}
}

func TestAuthApplication_UpdateProfile(t *testing.T) {
	type args struct {
		input *input.UpdateAuthProfile
		user  *user.User
	}

	testCases := map[string]struct {
		args args
		want error
	}{
		"ok": {
			args: args{
				input: &input.UpdateAuthProfile{
					Username:         "test-user",
					Gender:           0,
					ThumbnailURL:     "",
					SelfIntroduction: "自己紹介",
				},
				user: &user.User{
					ID: "00000000-0000-0000-0000-000000000000",
				},
			},
			want: nil,
		},
	}

	for name, tc := range testCases {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		arvm := mock_validation.NewMockAuthRequestValidation(ctrl)
		arvm.EXPECT().UpdateAuthProfile(tc.args.input).Return(nil)

		usm := mock_user.NewMockService(ctrl)
		usm.EXPECT().Validation(ctx, gomock.Any()).Return(nil)
		usm.EXPECT().Update(ctx, tc.args.user).Return(tc.want)

		t.Run(name, func(t *testing.T) {
			target := NewAuthApplication(arvm, usm)

			got := target.UpdateProfile(ctx, tc.args.input, tc.args.user)
			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("want %#v, but %#v", tc.want, got)
				return
			}

			if tc.args.user.Username != tc.args.input.Username {
				t.Fatalf("want %#v, but %#v", tc.args.input.Username, tc.args.user.Username)
				return
			}

			if tc.args.user.Gender != tc.args.input.Gender {
				t.Fatalf("want %#v, but %#v", tc.args.input.Gender, tc.args.user.Gender)
				return
			}

			if tc.args.user.ThumbnailURL != tc.args.input.ThumbnailURL {
				t.Fatalf("want %#v, but %#v", tc.args.input.ThumbnailURL, tc.args.user.ThumbnailURL)
				return
			}

			if tc.args.user.SelfIntroduction != tc.args.input.SelfIntroduction {
				t.Fatalf("want %#v, but %#v", tc.args.input.SelfIntroduction, tc.args.user.SelfIntroduction)
				return
			}
		})
	}
}

func TestAuthApplication_UpdateAddress(t *testing.T) {
	type args struct {
		input *input.UpdateAuthAddress
		user  *user.User
	}

	testCases := map[string]struct {
		args args
		want error
	}{
		"ok": {
			args: args{
				input: &input.UpdateAuthAddress{
					LastName:      "テスト",
					FirstName:     "ユーザ",
					LastNameKana:  "てすと",
					FirstNameKana: "ゆーざ",
					PhoneNumber:   "000-0000-0000",
					PostalCode:    "000-0000",
					Prefecture:    "東京都",
					City:          "小金井市",
					AddressLine1:  "貫井北町4-1-1",
					AddressLine2:  "自然科学棟 N110",
				},
				user: &user.User{
					ID: "00000000-0000-0000-0000-000000000000",
				},
			},
			want: nil,
		},
	}

	for name, tc := range testCases {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		arvm := mock_validation.NewMockAuthRequestValidation(ctrl)
		arvm.EXPECT().UpdateAuthAddress(tc.args.input).Return(nil)

		usm := mock_user.NewMockService(ctrl)
		usm.EXPECT().Validation(ctx, gomock.Any()).Return(nil)
		usm.EXPECT().Update(ctx, tc.args.user).Return(tc.want)

		t.Run(name, func(t *testing.T) {
			target := NewAuthApplication(arvm, usm)

			got := target.UpdateAddress(ctx, tc.args.input, tc.args.user)
			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("want %#v, but %#v", tc.want, got)
				return
			}

			if tc.args.user.LastName != tc.args.input.LastName {
				t.Fatalf("want %#v, but %#v", tc.args.input.LastName, tc.args.user.LastName)
				return
			}

			if tc.args.user.FirstName != tc.args.input.FirstName {
				t.Fatalf("want %#v, but %#v", tc.args.input.FirstName, tc.args.user.FirstName)
				return
			}

			if tc.args.user.LastNameKana != tc.args.input.LastNameKana {
				t.Fatalf("want %#v, but %#v", tc.args.input.LastNameKana, tc.args.user.LastNameKana)
				return
			}

			if tc.args.user.FirstNameKana != tc.args.input.FirstNameKana {
				t.Fatalf("want %#v, but %#v", tc.args.input.FirstNameKana, tc.args.user.FirstNameKana)
				return
			}

			if tc.args.user.PhoneNumber != tc.args.input.PhoneNumber {
				t.Fatalf("want %#v, but %#v", tc.args.input.PhoneNumber, tc.args.user.PhoneNumber)
				return
			}

			if tc.args.user.PostalCode != tc.args.input.PostalCode {
				t.Fatalf("want %#v, but %#v", tc.args.input.PostalCode, tc.args.user.PostalCode)
				return
			}

			if tc.args.user.Prefecture != tc.args.input.Prefecture {
				t.Fatalf("want %#v, but %#v", tc.args.input.Prefecture, tc.args.user.Prefecture)
				return
			}

			if tc.args.user.City != tc.args.input.City {
				t.Fatalf("want %#v, but %#v", tc.args.input.City, tc.args.user.City)
				return
			}

			if tc.args.user.AddressLine1 != tc.args.input.AddressLine1 {
				t.Fatalf("want %#v, but %#v", tc.args.input.AddressLine1, tc.args.user.AddressLine1)
				return
			}

			if tc.args.user.AddressLine2 != tc.args.input.AddressLine2 {
				t.Fatalf("want %#v, but %#v", tc.args.input.AddressLine2, tc.args.user.AddressLine2)
				return
			}
		})
	}
}

func TestAuthApplication_UploadThumbnail(t *testing.T) {
	type args struct {
		input *input.UploadAuthThumbnail
		user  *user.User
	}
	type want struct {
		thumbnailURL string
		err          error
	}

	testCases := map[string]struct {
		args args
		want want
	}{
		"ok": {
			args: args{
				input: &input.UploadAuthThumbnail{
					Thumbnail: []byte("あいうえお"),
				},
				user: &user.User{
					ID: "00000000-0000-0000-0000-000000000000",
				},
			},
			want: want{
				thumbnailURL: "https://google.co.jp",
				err:          nil,
			},
		},
	}

	for name, tc := range testCases {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		arvm := mock_validation.NewMockAuthRequestValidation(ctrl)
		arvm.EXPECT().UploadAuthThumbnail(tc.args.input).Return(nil)

		usm := mock_user.NewMockService(ctrl)
		usm.EXPECT().
			UploadThumbnail(ctx, tc.args.user.ID, tc.args.input.Thumbnail).
			Return(tc.want.thumbnailURL, tc.want.err)

		t.Run(name, func(t *testing.T) {
			target := NewAuthApplication(arvm, usm)

			thumbnailURL, err := target.UploadThumbnail(ctx, tc.args.input, tc.args.user)
			if !reflect.DeepEqual(err, tc.want.err) {
				t.Fatalf("want %#v, but %#v", tc.want.err, err)
				return
			}

			if !reflect.DeepEqual(thumbnailURL, tc.want.thumbnailURL) {
				t.Fatalf("want %#v, but %#v", tc.want.thumbnailURL, thumbnailURL)
				return
			}
		})
	}
}

func TestAuthApplication_Delete(t *testing.T) {
	type args struct {
		user *user.User
	}

	testCases := map[string]struct {
		args args
		want error
	}{
		"ok": {
			args: args{
				user: &user.User{
					ID: "00000000-0000-0000-0000-000000000000",
				},
			},
			want: nil,
		},
	}

	for name, tc := range testCases {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		arvm := mock_validation.NewMockAuthRequestValidation(ctrl)

		usm := mock_user.NewMockService(ctrl)
		usm.EXPECT().Delete(ctx, tc.args.user.ID).Return(tc.want)

		t.Run(name, func(t *testing.T) {
			target := NewAuthApplication(arvm, usm)

			err := target.Delete(ctx, tc.args.user)
			if !reflect.DeepEqual(err, tc.want) {
				t.Fatalf("want %#v, but %#v", tc.want, err)
				return
			}
		})
	}
}

func TestAuthApplication_RegisterDevice(t *testing.T) {
	type args struct {
		input *input.RegisterAuthDevice
		user  *user.User
	}

	testCases := map[string]struct {
		args args
		want error
	}{
		"ok": {
			args: args{
				input: &input.RegisterAuthDevice{
					InstanceID: "cTP0f6Y_Q26VG9TbTjReZz:APA91bG6Ns9A5DsXaMcImyyNImS4VD",
				},
				user: &user.User{
					ID: "00000000-0000-0000-0000-000000000000",
				},
			},
			want: nil,
		},
	}

	for name, tc := range testCases {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		arvm := mock_validation.NewMockAuthRequestValidation(ctrl)
		arvm.EXPECT().RegisterAuthDevice(tc.args.input).Return(nil)

		usm := mock_user.NewMockService(ctrl)
		usm.EXPECT().Update(ctx, tc.args.user).Return(tc.want)

		t.Run(name, func(t *testing.T) {
			target := NewAuthApplication(arvm, usm)

			err := target.RegisterDevice(ctx, tc.args.input, tc.args.user)
			if !reflect.DeepEqual(err, tc.want) {
				t.Fatalf("want %#v, but %#v", tc.want, err)
				return
			}

			if tc.args.user.InstanceID != tc.args.input.InstanceID {
				t.Fatalf("want %#v, but %#v", tc.args.input.InstanceID, tc.args.user.InstanceID)
				return
			}
		})
	}
}
