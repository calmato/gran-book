package application

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/calmato/gran-book/api/server/user/internal/application/input"
	"github.com/calmato/gran-book/api/server/user/internal/application/output"
	"github.com/calmato/gran-book/api/server/user/internal/domain/user"
	mock_validation "github.com/calmato/gran-book/api/server/user/mock/application/validation"
	mock_user "github.com/calmato/gran-book/api/server/user/mock/domain/user"
	"github.com/golang/mock/gomock"
)

func TestAdminApplication_List(t *testing.T) {
	type args struct {
		input *input.ListAdmin
	}
	type want struct {
		users  []*user.User
		output *output.ListQuery
		err    error
	}

	current := time.Now().Local()
	testCases := map[string]struct {
		args args
		want want
	}{
		"ok": {
			args: args{
				input: &input.ListAdmin{
					Limit:  100,
					Offset: 0,
				},
			},
			want: want{
				users: []*user.User{
					{
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
				},
				output: &output.ListQuery{
					Limit:  100,
					Offset: 0,
					Total:  1,
					Order:  nil,
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

		arvm := mock_validation.NewMockAdminRequestValidation(ctrl)
		arvm.EXPECT().ListAdmin(tc.args.input).Return(nil)

		usm := mock_user.NewMockService(ctrl)
		usm.EXPECT().List(ctx, gomock.Any()).Return(tc.want.users, tc.want.err)
		usm.EXPECT().ListCount(ctx, gomock.Any()).Return(tc.want.output.Total, tc.want.err)

		t.Run(name, func(t *testing.T) {
			target := NewAdminApplication(arvm, usm)

			users, _, err := target.List(ctx, tc.args.input)
			if !reflect.DeepEqual(err, tc.want.err) {
				t.Fatalf("want %#v, but %#v", tc.want.err, err)
				return
			}

			if !reflect.DeepEqual(users, tc.want.users) {
				t.Fatalf("want %#v, but %#v", tc.want.users, users)
				return
			}
		})
	}
}

func TestAdminApplication_Search(t *testing.T) {
	type args struct {
		input *input.SearchAdmin
	}
	type want struct {
		users  []*user.User
		output *output.ListQuery
		err    error
	}

	current := time.Now().Local()

	testCases := map[string]struct {
		args args
		want want
	}{
		"ok": {
			args: args{
				input: &input.SearchAdmin{
					Limit:  100,
					Offset: 0,
					Field:  "email",
					Value:  "test-user@calmato.com",
				},
			},
			want: want{
				users: []*user.User{
					{
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
				},
				output: &output.ListQuery{
					Limit:  100,
					Offset: 0,
					Total:  1,
					Order:  nil,
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

		arvm := mock_validation.NewMockAdminRequestValidation(ctrl)
		arvm.EXPECT().SearchAdmin(tc.args.input).Return(nil)

		usm := mock_user.NewMockService(ctrl)
		usm.EXPECT().List(ctx, gomock.Any()).Return(tc.want.users, tc.want.err)
		usm.EXPECT().ListCount(ctx, gomock.Any()).Return(tc.want.output.Total, tc.want.err)

		t.Run(name, func(t *testing.T) {
			target := NewAdminApplication(arvm, usm)

			users, _, err := target.Search(ctx, tc.args.input)
			if !reflect.DeepEqual(err, tc.want.err) {
				t.Fatalf("want %#v, but %#v", tc.want.err, err)
				return
			}

			if !reflect.DeepEqual(users, tc.want.users) {
				t.Fatalf("want %#v, but %#v", tc.want.users, users)
				return
			}
		})
	}
}

func TestAdminApplication_Show(t *testing.T) {
	type args struct {
		uid string
	}
	type want struct {
		user *user.User
		err  error
	}

	current := time.Now()
	testCases := map[string]struct {
		args args
		want want
	}{
		"ok": {
			args: args{
				uid: "00000000-0000-0000-0000-000000000000",
			},
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

		arvm := mock_validation.NewMockAdminRequestValidation(ctrl)

		usm := mock_user.NewMockService(ctrl)
		usm.EXPECT().Show(ctx, tc.args.uid).Return(tc.want.user, tc.want.err)

		t.Run(name, func(t *testing.T) {
			target := NewAdminApplication(arvm, usm)

			user, err := target.Show(ctx, tc.args.uid)
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

func TestAdminApplication_Create(t *testing.T) {
	type args struct {
		input *input.CreateAdmin
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
				input: &input.CreateAdmin{
					Username:      "test-user",
					Email:         "test@calmato.com",
					Password:      "12345678",
					Role:          1,
					LastName:      "テスト",
					FirstName:     "ユーザ",
					LastNameKana:  "てすと",
					FirstNameKana: "ゆーざ",
				},
			},
			want: want{
				user: &user.User{
					Username:      "test-user",
					Email:         "test@calmato.com",
					Password:      "12345678",
					Gender:        0,
					Role:          1,
					LastName:      "テスト",
					FirstName:     "ユーザ",
					LastNameKana:  "てすと",
					FirstNameKana: "ゆーざ",
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

		arvm := mock_validation.NewMockAdminRequestValidation(ctrl)
		arvm.EXPECT().CreateAdmin(tc.args.input).Return(nil)

		usm := mock_user.NewMockService(ctrl)
		usm.EXPECT().Validation(ctx, gomock.Any()).Return(nil)
		usm.EXPECT().Create(ctx, gomock.Any()).Return(tc.want.err)

		t.Run(name, func(t *testing.T) {
			target := NewAdminApplication(arvm, usm)

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

func TestAdminApplication_UpdateContact(t *testing.T) {
	type args struct {
		uid   string
		input *input.UpdateAdminContact
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
				uid: "12345678-1234-1234-1234-12345678901234",
				input: &input.UpdateAdminContact{
					Email:       "test-user@calmato.jp",
					PhoneNumber: "000-0000-0000",
				},
			},
			want: want{
				user: &user.User{
					ID:          "12345678-1234-1234-1234-12345678901234",
					Email:       "test-user@calmato.jp",
					PhoneNumber: "000-0000-0000",
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

		arvm := mock_validation.NewMockAdminRequestValidation(ctrl)
		arvm.EXPECT().UpdateAdminContact(tc.args.input).Return(nil)

		usm := mock_user.NewMockService(ctrl)
		usm.EXPECT().Validation(ctx, gomock.Any()).Return(nil)
		usm.EXPECT().Show(ctx, tc.args.uid).Return(tc.want.user, nil)
		usm.EXPECT().Update(ctx, tc.want.user).Return(tc.want.err)

		t.Run(name, func(t *testing.T) {
			target := NewAdminApplication(arvm, usm)

			user, err := target.UpdateContact(ctx, tc.args.input, tc.args.uid)
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

func TestAdminApplication_UpdatePassword(t *testing.T) {
	type args struct {
		uid   string
		input *input.UpdateAdminPassword
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
				uid: "12345678-1234-1234-1234-12345678901234",
				input: &input.UpdateAdminPassword{
					Password:             "12345678",
					PasswordConfirmation: "12345678",
				},
			},
			want: want{
				user: &user.User{
					ID:       "12345678-1234-1234-1234-12345678901234",
					Password: "12345678",
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

		arvm := mock_validation.NewMockAdminRequestValidation(ctrl)
		arvm.EXPECT().UpdateAdminPassword(tc.args.input).Return(nil)

		usm := mock_user.NewMockService(ctrl)
		usm.EXPECT().Show(ctx, tc.args.uid).Return(tc.want.user, nil)
		usm.EXPECT().UpdatePassword(ctx, tc.want.user.ID, tc.want.user.Password).Return(tc.want.err)

		t.Run(name, func(t *testing.T) {
			target := NewAdminApplication(arvm, usm)

			user, err := target.UpdatePassword(ctx, tc.args.input, tc.args.uid)
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

func TestAdminApplication_UpdateProfile(t *testing.T) {
	type args struct {
		uid   string
		input *input.UpdateAdminProfile
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
				uid: "12345678-1234-1234-1234-12345678901234",
				input: &input.UpdateAdminProfile{
					Username:      "test-user",
					LastName:      "テスト",
					FirstName:     "ユーザ",
					LastNameKana:  "てすと",
					FirstNameKana: "ゆーざ",
					Role:          1,
					ThumbnailURL:  "https://www.google.co.jp",
				},
			},
			want: want{
				user: &user.User{
					ID:            "12345678-1234-1234-1234-12345678901234",
					Username:      "test-user",
					Email:         "test@calmato.com",
					LastName:      "テスト",
					FirstName:     "ユーザ",
					LastNameKana:  "てすと",
					FirstNameKana: "ゆーざ",
					Role:          1,
					ThumbnailURL:  "https://www.google.co.jp",
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

		arvm := mock_validation.NewMockAdminRequestValidation(ctrl)
		arvm.EXPECT().UpdateAdminProfile(tc.args.input).Return(nil)

		usm := mock_user.NewMockService(ctrl)
		usm.EXPECT().Validation(ctx, gomock.Any()).Return(nil)
		usm.EXPECT().Show(ctx, tc.args.uid).Return(tc.want.user, nil)
		usm.EXPECT().Update(ctx, tc.want.user).Return(tc.want.err)

		t.Run(name, func(t *testing.T) {
			target := NewAdminApplication(arvm, usm)

			user, err := target.UpdateProfile(ctx, tc.args.input, tc.args.uid)
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

func TestAdminApplication_Delete(t *testing.T) {
	type args struct {
		uid string
	}

	testCases := map[string]struct {
		args args
		want error
	}{
		"ok": {
			args: args{
				uid: "00000000-0000-0000-0000-000000000000",
			},
			want: nil,
		},
	}

	for name, tc := range testCases {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		u := &user.User{
			ID: tc.args.uid,
		}

		arvm := mock_validation.NewMockAdminRequestValidation(ctrl)

		usm := mock_user.NewMockService(ctrl)
		usm.EXPECT().Show(ctx, tc.args.uid).Return(u, nil)
		usm.EXPECT().Update(ctx, u).Return(tc.want)

		t.Run(name, func(t *testing.T) {
			target := NewAdminApplication(arvm, usm)

			err := target.Delete(ctx, tc.args.uid)
			if !reflect.DeepEqual(err, tc.want) {
				t.Fatalf("want %#v, but %#v", tc.want, err)
				return
			}
		})
	}
}
