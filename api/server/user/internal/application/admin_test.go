package application

import (
	"context"
	"reflect"
	"testing"

	"github.com/calmato/gran-book/api/server/user/internal/application/input"
	"github.com/calmato/gran-book/api/server/user/internal/domain/user"
	mock_validation "github.com/calmato/gran-book/api/server/user/mock/application/validation"
	mock_user "github.com/calmato/gran-book/api/server/user/mock/domain/user"
	"github.com/golang/mock/gomock"
)

func TestAdminApplication_Create(t *testing.T) {
	testCases := map[string]struct {
		Input    *input.CreateAdmin
		Expected struct {
			User  *user.User
			Error error
		}
	}{
		"ok": {
			Input: &input.CreateAdmin{
				Username:      "test-user",
				Email:         "test@calmato.com",
				Password:      "12345678",
				Role:          1,
				LastName:      "テスト",
				FirstName:     "ユーザ",
				LastNameKana:  "てすと",
				FirstNameKana: "ゆーざ",
			},
			Expected: struct {
				User  *user.User
				Error error
			}{
				User: &user.User{
					Username:      "test-user",
					Email:         "test@calmato.com",
					Password:      "12345678",
					Gender:        0,
					Role:          1,
					LastName:      "テスト",
					FirstName:     "ユーザ",
					LastNameKana:  "てすと",
					FirstNameKana: "ゆーざ",
					Activated:     true,
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

		arvm := mock_validation.NewMockAdminRequestValidation(ctrl)
		arvm.EXPECT().CreateAdmin(tc.Input).Return(nil)

		usm := mock_user.NewMockService(ctrl)
		usm.EXPECT().Create(ctx, gomock.Any()).Return(tc.Expected.Error)

		t.Run(result, func(t *testing.T) {
			target := NewAdminApplication(arvm, usm)

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

func TestAdminApplication_UpdateRole(t *testing.T) {
	testCases := map[string]struct {
		Input    *input.UpdateAdminRole
		ID       string
		Expected struct {
			User  *user.User
			Error error
		}
	}{
		"ok": {
			Input: &input.UpdateAdminRole{
				Role: 2,
			},
			ID: "12345678-1234-1234-1234-12345678901234",
			Expected: struct {
				User  *user.User
				Error error
			}{
				User: &user.User{
					ID:   "12345678-1234-1234-1234-12345678901234",
					Role: 2,
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

		arvm := mock_validation.NewMockAdminRequestValidation(ctrl)
		arvm.EXPECT().UpdateAdminRole(tc.Input).Return(nil)

		usm := mock_user.NewMockService(ctrl)
		usm.EXPECT().Show(ctx, tc.ID).Return(tc.Expected.User, tc.Expected.Error)
		usm.EXPECT().Update(ctx, tc.Expected.User).Return(tc.Expected.Error)

		t.Run(result, func(t *testing.T) {
			target := NewAdminApplication(arvm, usm)

			got, err := target.UpdateRole(ctx, tc.Input, tc.ID)
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

func TestAdminApplication_UpdatePassword(t *testing.T) {
	testCases := map[string]struct {
		Input    *input.UpdateAdminPassword
		ID       string
		Expected struct {
			User  *user.User
			Error error
		}
	}{
		"ok": {
			Input: &input.UpdateAdminPassword{
				Password:             "12345678",
				PasswordConfirmation: "12345678",
			},
			ID: "12345678-1234-1234-1234-12345678901234",
			Expected: struct {
				User  *user.User
				Error error
			}{
				User: &user.User{
					ID:       "12345678-1234-1234-1234-12345678901234",
					Password: "12345678",
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

		arvm := mock_validation.NewMockAdminRequestValidation(ctrl)
		arvm.EXPECT().UpdateAdminPassword(tc.Input).Return(nil)

		usm := mock_user.NewMockService(ctrl)
		usm.EXPECT().Show(ctx, tc.ID).Return(tc.Expected.User, tc.Expected.Error)
		usm.EXPECT().UpdatePassword(ctx, tc.Expected.User.ID, tc.Expected.User.Password).Return(tc.Expected.Error)

		t.Run(result, func(t *testing.T) {
			target := NewAdminApplication(arvm, usm)

			got, err := target.UpdatePassword(ctx, tc.Input, tc.ID)
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

func TestAdminApplication_UpdateProfile(t *testing.T) {
	testCases := map[string]struct {
		Input    *input.UpdateAdminProfile
		ID       string
		Expected struct {
			User  *user.User
			Error error
		}
	}{
		"ok": {
			Input: &input.UpdateAdminProfile{
				Username:      "test-user",
				Email:         "test@calmato.com",
				LastName:      "テスト",
				FirstName:     "ユーザ",
				LastNameKana:  "てすと",
				FirstNameKana: "ゆーざ",
			},
			ID: "12345678-1234-1234-1234-12345678901234",
			Expected: struct {
				User  *user.User
				Error error
			}{
				User: &user.User{
					ID:            "12345678-1234-1234-1234-12345678901234",
					Username:      "test-user",
					Email:         "test@calmato.com",
					LastName:      "テスト",
					FirstName:     "ユーザ",
					LastNameKana:  "てすと",
					FirstNameKana: "ゆーざ",
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

		arvm := mock_validation.NewMockAdminRequestValidation(ctrl)
		arvm.EXPECT().UpdateAdminProfile(tc.Input).Return(nil)

		usm := mock_user.NewMockService(ctrl)
		usm.EXPECT().Show(ctx, tc.ID).Return(tc.Expected.User, tc.Expected.Error)
		usm.EXPECT().Update(ctx, tc.Expected.User).Return(tc.Expected.Error)

		t.Run(result, func(t *testing.T) {
			target := NewAdminApplication(arvm, usm)

			got, err := target.UpdateProfile(ctx, tc.Input, tc.ID)
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
