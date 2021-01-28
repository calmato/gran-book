package service

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/calmato/gran-book/api/user/internal/domain/exception"
	"github.com/calmato/gran-book/api/user/internal/domain/user"
	mock_user "github.com/calmato/gran-book/api/user/mock/domain/user"
	"github.com/golang/mock/gomock"
)

func TestUserService_Authentication(t *testing.T) {
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
					CreatedAt:        current,
					UpdatedAt:        current,
				},
				Error: nil,
			},
		},
		"ng_unauthorized": {
			Expected: struct {
				User  *user.User
				Error error
			}{
				User:  nil,
				Error: exception.Unauthorized.New(nil),
			},
		},
	}

	for result, tc := range testCases {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		uvm := mock_user.NewMockValidation(ctrl)

		urm := mock_user.NewMockRepository(ctrl)
		urm.EXPECT().Authentication(ctx).Return(tc.Expected.User, tc.Expected.Error)

		uum := mock_user.NewMockUploader(ctrl)

		t.Run(result, func(t *testing.T) {
			target := NewUserService(uvm, urm, uum)

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

func TestUserService_Create(t *testing.T) {
	testCases := map[string]struct {
		User     *user.User
		Expected error
	}{
		"ok": {
			User: &user.User{
				ID:               "",
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
				CreatedAt:        time.Time{},
				UpdatedAt:        time.Time{},
			},
			Expected: nil,
		},
	}

	for result, tc := range testCases {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		uvm := mock_user.NewMockValidation(ctrl)
		uvm.EXPECT().User(ctx, tc.User).Return(nil)

		urm := mock_user.NewMockRepository(ctrl)
		urm.EXPECT().Create(ctx, tc.User).Return(tc.Expected)

		uum := mock_user.NewMockUploader(ctrl)

		t.Run(result, func(t *testing.T) {
			target := NewUserService(uvm, urm, uum)

			got := target.Create(ctx, tc.User)
			if !reflect.DeepEqual(got, tc.Expected) {
				t.Fatalf("want %#v, but %#v", tc.Expected, got)
				return
			}

			if tc.User.ID == "" {
				t.Fatal("User.ID must be not null")
				return
			}

			if tc.User.CreatedAt.IsZero() {
				t.Fatal("User.CreatedAt must be not null")
				return
			}

			if tc.User.UpdatedAt.IsZero() {
				t.Fatal("User.UpdatedAt must be not null")
				return
			}
		})
	}
}

func TestUserService_Update(t *testing.T) {
	current := time.Now()

	testCases := map[string]struct {
		User     *user.User
		Expected error
	}{
		"ok": {
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
				CreatedAt:        current,
				UpdatedAt:        current,
			},
			Expected: nil,
		},
	}

	for result, tc := range testCases {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		uvm := mock_user.NewMockValidation(ctrl)
		uvm.EXPECT().User(ctx, tc.User).Return(nil)

		urm := mock_user.NewMockRepository(ctrl)
		urm.EXPECT().Update(ctx, tc.User).Return(tc.Expected)

		uum := mock_user.NewMockUploader(ctrl)

		t.Run(result, func(t *testing.T) {
			target := NewUserService(uvm, urm, uum)

			got := target.Update(ctx, tc.User)
			if !reflect.DeepEqual(got, tc.Expected) {
				t.Fatalf("want %#v, but %#v", tc.Expected, got)
				return
			}

			if tc.User.UpdatedAt == current {
				t.Fatal("User.UpdatedAt must be changed")
				return
			}
		})
	}
}

func TestUserService_UpdatePassword(t *testing.T) {
	testCases := map[string]struct {
		UID      string
		Password string
		Expected error
	}{
		"ok": {
			UID:      "00000000-0000-0000-0000-000000000000",
			Password: "12345678",
			Expected: nil,
		},
	}

	for result, tc := range testCases {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		uvm := mock_user.NewMockValidation(ctrl)

		urm := mock_user.NewMockRepository(ctrl)
		urm.EXPECT().UpdatePassword(ctx, tc.UID, tc.Password).Return(tc.Expected)

		uum := mock_user.NewMockUploader(ctrl)

		t.Run(result, func(t *testing.T) {
			target := NewUserService(uvm, urm, uum)

			got := target.UpdatePassword(ctx, tc.UID, tc.Password)
			if !reflect.DeepEqual(got, tc.Expected) {
				t.Fatalf("want %#v, but %#v", tc.Expected, got)
				return
			}
		})
	}
}

func TestUserService_UpdateThumbnail(t *testing.T) {
	testCases := map[string]struct {
		UID       string
		Thumbnail []byte
		Expected  struct {
			ThumbailURL string
			Error       error
		}
	}{
		"ok": {
			UID:       "00000000-0000-0000-0000-000000000000",
			Thumbnail: []byte{},
			Expected: struct {
				ThumbailURL string
				Error       error
			}{
				ThumbailURL: "https://calmato.com/user_thumbnails",
				Error:       nil,
			},
		},
	}

	for result, tc := range testCases {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		uvm := mock_user.NewMockValidation(ctrl)

		urm := mock_user.NewMockRepository(ctrl)

		uum := mock_user.NewMockUploader(ctrl)
		uum.EXPECT().Thumbnail(ctx, tc.UID, tc.Thumbnail).Return(tc.Expected.ThumbailURL, tc.Expected.Error)

		t.Run(result, func(t *testing.T) {
			target := NewUserService(uvm, urm, uum)

			got, err := target.UploadThumbnail(ctx, tc.UID, tc.Thumbnail)
			if !reflect.DeepEqual(err, tc.Expected.Error) {
				t.Fatalf("want %#v, but %#v", tc.Expected.Error, err)
				return
			}

			if !reflect.DeepEqual(got, tc.Expected.ThumbailURL) {
				t.Fatalf("want %#v, but %#v", tc.Expected.ThumbailURL, got)
				return
			}
		})
	}
}
