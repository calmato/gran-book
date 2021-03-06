package application

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/calmato/gran-book/api/server/user/internal/application/output"
	"github.com/calmato/gran-book/api/server/user/internal/domain/user"
	mock_validation "github.com/calmato/gran-book/api/server/user/mock/application/validation"
	mock_user "github.com/calmato/gran-book/api/server/user/mock/domain/user"
	"github.com/golang/mock/gomock"
)

func TestUserApplication_GetUserProfile(t *testing.T) {
	current := time.Now()

	testCases := map[string]struct {
		UID      string
		CUID     string
		Expected struct {
			User   *user.User
			Output *output.UserProfile
			Error  error
		}
	}{
		"ok": {
			UID:  "00000000-0000-0000-0000-000000000000",
			CUID: "11111111-1111-1111-1111-111111111111",
			Expected: struct {
				User   *user.User
				Output *output.UserProfile
				Error  error
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
				Output: &output.UserProfile{
					IsFollow:      true,
					IsFollower:    false,
					FollowCount:   1,
					FollowerCount: 3,
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

		urvm := mock_validation.NewMockUserRequestValidation(ctrl)

		usm := mock_user.NewMockService(ctrl)
		usm.EXPECT().Show(ctx, tc.UID).Return(tc.Expected.User, tc.Expected.Error)
		usm.EXPECT().
			ListFriendCount(ctx, tc.UID).
			Return(tc.Expected.Output.FollowCount, tc.Expected.Output.FollowerCount, tc.Expected.Error)
		usm.EXPECT().
			IsFriend(ctx, tc.UID, tc.CUID).
			Return(tc.Expected.Output.IsFollow, tc.Expected.Output.IsFollower)

		t.Run(result, func(t *testing.T) {
			target := NewUserApplication(urvm, usm)

			user, output, err := target.GetUserProfile(ctx, tc.UID, tc.CUID)
			if !reflect.DeepEqual(err, tc.Expected.Error) {
				t.Fatalf("want %#v, but %#v", tc.Expected.Error, err)
				return
			}

			if !reflect.DeepEqual(user, tc.Expected.User) {
				t.Fatalf("want %#v, but %#v", tc.Expected.User, user)
				return
			}

			if !reflect.DeepEqual(output, tc.Expected.Output) {
				t.Fatalf("want %#v, but %#v", tc.Expected.Output, output)
				return
			}
		})
	}
}

func TestUserApplication_RegisterFollow(t *testing.T) {
	current := time.Now()

	testCases := map[string]struct {
		UID      string
		CUID     string
		Expected struct {
			User   *user.User
			Output *output.UserProfile
			Error  error
		}
	}{
		"ok": {
			UID:  "00000000-0000-0000-0000-000000000000",
			CUID: "11111111-1111-1111-1111-111111111111",
			Expected: struct {
				User   *user.User
				Output *output.UserProfile
				Error  error
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
				Output: &output.UserProfile{
					IsFollow:      true,
					IsFollower:    false,
					FollowCount:   1,
					FollowerCount: 3,
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

		r := &user.Relationship{
			FollowID:   tc.CUID,
			FollowerID: tc.UID,
		}

		urvm := mock_validation.NewMockUserRequestValidation(ctrl)

		usm := mock_user.NewMockService(ctrl)
		usm.EXPECT().Show(ctx, tc.UID).Return(tc.Expected.User, tc.Expected.Error)
		usm.EXPECT().CreateRelationship(ctx, r).Return(tc.Expected.Error)
		usm.EXPECT().
			ListFriendCount(ctx, tc.UID).
			Return(tc.Expected.Output.FollowCount, tc.Expected.Output.FollowerCount, tc.Expected.Error)
		usm.EXPECT().
			IsFriend(ctx, tc.UID, tc.CUID).
			Return(tc.Expected.Output.IsFollow, tc.Expected.Output.IsFollower)

		t.Run(result, func(t *testing.T) {
			target := NewUserApplication(urvm, usm)

			user, output, err := target.RegisterFollow(ctx, tc.UID, tc.CUID)
			if !reflect.DeepEqual(err, tc.Expected.Error) {
				t.Fatalf("want %#v, but %#v", tc.Expected.Error, err)
				return
			}

			if !reflect.DeepEqual(user, tc.Expected.User) {
				t.Fatalf("want %#v, but %#v", tc.Expected.User, user)
				return
			}

			if !reflect.DeepEqual(output, tc.Expected.Output) {
				t.Fatalf("want %#v, but %#v", tc.Expected.Output, output)
				return
			}
		})
	}
}

func TestUserApplication_UnregisterFollow(t *testing.T) {
	current := time.Now()

	testCases := map[string]struct {
		UID      string
		CUID     string
		Expected struct {
			User   *user.User
			Output *output.UserProfile
			Error  error
		}
	}{
		"ok": {
			UID:  "00000000-0000-0000-0000-000000000000",
			CUID: "11111111-1111-1111-1111-111111111111",
			Expected: struct {
				User   *user.User
				Output *output.UserProfile
				Error  error
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
				Output: &output.UserProfile{
					IsFollow:      true,
					IsFollower:    false,
					FollowCount:   1,
					FollowerCount: 3,
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

		r := &user.Relationship{
			ID:         1,
			FollowID:   tc.CUID,
			FollowerID: tc.UID,
		}

		urvm := mock_validation.NewMockUserRequestValidation(ctrl)

		usm := mock_user.NewMockService(ctrl)
		usm.EXPECT().Show(ctx, tc.UID).Return(tc.Expected.User, tc.Expected.Error)
		usm.EXPECT().ShowRelationshipByUID(ctx, tc.UID, tc.CUID).Return(r, tc.Expected.Error)
		usm.EXPECT().DeleteRelationship(ctx, r.ID).Return(tc.Expected.Error)
		usm.EXPECT().
			ListFriendCount(ctx, tc.UID).
			Return(tc.Expected.Output.FollowCount, tc.Expected.Output.FollowerCount, tc.Expected.Error)
		usm.EXPECT().
			IsFriend(ctx, tc.UID, tc.CUID).
			Return(tc.Expected.Output.IsFollow, tc.Expected.Output.IsFollower)

		t.Run(result, func(t *testing.T) {
			target := NewUserApplication(urvm, usm)

			user, output, err := target.UnregisterFollow(ctx, tc.UID, tc.CUID)
			if !reflect.DeepEqual(err, tc.Expected.Error) {
				t.Fatalf("want %#v, but %#v", tc.Expected.Error, err)
				return
			}

			if !reflect.DeepEqual(user, tc.Expected.User) {
				t.Fatalf("want %#v, but %#v", tc.Expected.User, user)
				return
			}

			if !reflect.DeepEqual(output, tc.Expected.Output) {
				t.Fatalf("want %#v, but %#v", tc.Expected.Output, output)
				return
			}
		})
	}
}
