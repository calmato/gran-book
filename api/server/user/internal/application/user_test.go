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

func TestUserApplication_ListFollow(t *testing.T) {
	testCases := map[string]struct {
		UID      string
		CUID     string
		Input    *input.ListFollow
		Expected struct {
			Follows []*user.Follow
			Output  *output.ListQuery
			Error   error
		}
	}{
		"ok": {
			UID:  "00000000-0000-0000-0000-000000000000",
			CUID: "11111111-1111-1111-1111-111111111111",
			Input: &input.ListFollow{
				Limit:  100,
				Offset: 0,
			},
			Expected: struct {
				Follows []*user.Follow
				Output  *output.ListQuery
				Error   error
			}{
				Follows: []*user.Follow{
					{
						FollowID:         "00000000-0000-0000-0000-000000000000",
						FollowerID:       "11111111-1111-1111-1111-111111111111",
						Username:         "test-user",
						ThumbnailURL:     "",
						SelfIntroduction: "",
					},
				},
				Output: &output.ListQuery{
					Limit:  100,
					Offset: 0,
					Total:  1,
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

		var followCount int
		var followerCount int

		if tc.Expected.Output != nil {
			followCount = tc.Expected.Output.Total
		}

		urvm := mock_validation.NewMockUserRequestValidation(ctrl)
		urvm.EXPECT().ListFollow(tc.Input).Return(nil)

		usm := mock_user.NewMockService(ctrl)
		usm.EXPECT().ListFollow(ctx, gomock.Any(), tc.CUID).Return(tc.Expected.Follows, tc.Expected.Error)
		usm.EXPECT().
			ListFriendCount(ctx, tc.UID).
			Return(followCount, followerCount, tc.Expected.Error)

		t.Run(result, func(t *testing.T) {
			target := NewUserApplication(urvm, usm)

			follows, output, err := target.ListFollow(ctx, tc.Input, tc.UID, tc.CUID)
			if !reflect.DeepEqual(err, tc.Expected.Error) {
				t.Fatalf("want %#v, but %#v", tc.Expected.Error, err)
				return
			}

			if !reflect.DeepEqual(follows, tc.Expected.Follows) {
				t.Fatalf("want %#v, but %#v", tc.Expected.Follows, follows)
				return
			}

			if !reflect.DeepEqual(output, tc.Expected.Output) {
				t.Fatalf("want %#v, but %#v", tc.Expected.Output, output)
				return
			}
		})
	}
}

func TestUserApplication_ListFollower(t *testing.T) {
	testCases := map[string]struct {
		UID      string
		CUID     string
		Input    *input.ListFollower
		Expected struct {
			Followers []*user.Follower
			Output    *output.ListQuery
			Error     error
		}
	}{
		"ok": {
			UID:  "00000000-0000-0000-0000-000000000000",
			CUID: "11111111-1111-1111-1111-111111111111",
			Input: &input.ListFollower{
				Limit:  100,
				Offset: 0,
			},
			Expected: struct {
				Followers []*user.Follower
				Output    *output.ListQuery
				Error     error
			}{
				Followers: []*user.Follower{
					{
						FollowID:         "11111111-1111-1111-1111-111111111111",
						FollowerID:       "00000000-0000-0000-0000-000000000000",
						Username:         "test-user",
						ThumbnailURL:     "",
						SelfIntroduction: "",
					},
				},
				Output: &output.ListQuery{
					Limit:  100,
					Offset: 0,
					Total:  1,
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

		var followCount int
		var followerCount int

		if tc.Expected.Output != nil {
			followerCount = tc.Expected.Output.Total
		}

		urvm := mock_validation.NewMockUserRequestValidation(ctrl)
		urvm.EXPECT().ListFollower(tc.Input).Return(nil)

		usm := mock_user.NewMockService(ctrl)
		usm.EXPECT().ListFollower(ctx, gomock.Any(), tc.CUID).Return(tc.Expected.Followers, tc.Expected.Error)
		usm.EXPECT().
			ListFriendCount(ctx, tc.UID).
			Return(followCount, followerCount, tc.Expected.Error)

		t.Run(result, func(t *testing.T) {
			target := NewUserApplication(urvm, usm)

			followers, output, err := target.ListFollower(ctx, tc.Input, tc.UID, tc.CUID)
			if !reflect.DeepEqual(err, tc.Expected.Error) {
				t.Fatalf("want %#v, but %#v", tc.Expected.Error, err)
				return
			}

			if !reflect.DeepEqual(followers, tc.Expected.Followers) {
				t.Fatalf("want %#v, but %#v", tc.Expected.Followers, followers)
				return
			}

			if !reflect.DeepEqual(output, tc.Expected.Output) {
				t.Fatalf("want %#v, but %#v", tc.Expected.Output, output)
				return
			}
		})
	}
}

func TestUserApplication_Show(t *testing.T) {
	current := time.Now()

	testCases := map[string]struct {
		UID      string
		Expected struct {
			User  *user.User
			Error error
		}
	}{
		"ok": {
			UID: "00000000-0000-0000-0000-000000000000",
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
	}

	for result, tc := range testCases {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		urvm := mock_validation.NewMockUserRequestValidation(ctrl)

		usm := mock_user.NewMockService(ctrl)
		usm.EXPECT().Show(ctx, tc.UID).Return(tc.Expected.User, tc.Expected.Error)

		t.Run(result, func(t *testing.T) {
			target := NewUserApplication(urvm, usm)

			user, err := target.Show(ctx, tc.UID)
			if !reflect.DeepEqual(err, tc.Expected.Error) {
				t.Fatalf("want %#v, but %#v", tc.Expected.Error, err)
				return
			}

			if !reflect.DeepEqual(user, tc.Expected.User) {
				t.Fatalf("want %#v, but %#v", tc.Expected.User, user)
				return
			}
		})
	}
}

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
		usm.EXPECT().ValidationRelationship(ctx, r).Return(tc.Expected.Error)
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
