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

func TestUserApplication_List(t *testing.T) {
	type args struct {
		input *input.ListUser
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
				input: &input.ListUser{
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

		arvm := mock_validation.NewMockUserRequestValidation(ctrl)
		arvm.EXPECT().ListUser(tc.args.input).Return(nil)

		usm := mock_user.NewMockService(ctrl)
		usm.EXPECT().List(ctx, gomock.Any()).Return(tc.want.users, nil)
		usm.EXPECT().ListCount(ctx, gomock.Any()).Return(tc.want.output.Total, tc.want.err)

		t.Run(name, func(t *testing.T) {
			target := NewUserApplication(arvm, usm)

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

func TestUserApplication_ListUserByIDs(t *testing.T) {
	type args struct {
		input *input.ListUserByUserIDs
	}
	type want struct {
		users []*user.User
		err   error
	}

	current := time.Now().Local()
	testCases := map[string]struct {
		args args
		want want
	}{
		"ok": {
			args: args{
				input: &input.ListUserByUserIDs{
					UserIDs: []string{
						"00000000-0000-0000-0000-000000000000",
					},
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
				err: nil,
			},
		},
	}

	for name, tc := range testCases {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		arvm := mock_validation.NewMockUserRequestValidation(ctrl)
		arvm.EXPECT().ListUserByUserIDs(tc.args.input).Return(nil)

		usm := mock_user.NewMockService(ctrl)
		usm.EXPECT().List(ctx, gomock.Any()).Return(tc.want.users, tc.want.err)

		t.Run(name, func(t *testing.T) {
			target := NewUserApplication(arvm, usm)

			users, err := target.ListByUserIDs(ctx, tc.args.input)
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

func TestUserApplication_ListFollow(t *testing.T) {
	type args struct {
		uid   string
		cuid  string
		input *input.ListFollow
	}
	type want struct {
		follows []*user.Follow
		output  *output.ListQuery
		err     error
	}

	testCases := map[string]struct {
		args args
		want want
	}{
		"ok": {
			args: args{
				uid:  "00000000-0000-0000-0000-000000000000",
				cuid: "11111111-1111-1111-1111-111111111111",
				input: &input.ListFollow{
					Limit:  100,
					Offset: 0,
				},
			},
			want: want{
				follows: []*user.Follow{
					{
						FollowID:         "00000000-0000-0000-0000-000000000000",
						FollowerID:       "11111111-1111-1111-1111-111111111111",
						Username:         "test-user",
						ThumbnailURL:     "",
						SelfIntroduction: "",
					},
				},
				output: &output.ListQuery{
					Limit:  100,
					Offset: 0,
					Total:  1,
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

		var followCount int
		var followerCount int

		if tc.want.output != nil {
			followCount = tc.want.output.Total
		}

		urvm := mock_validation.NewMockUserRequestValidation(ctrl)
		urvm.EXPECT().ListFollow(tc.args.input).Return(nil)

		usm := mock_user.NewMockService(ctrl)
		usm.EXPECT().ListFollow(ctx, gomock.Any(), tc.args.cuid).Return(tc.want.follows, tc.want.err)
		usm.EXPECT().ListFriendCount(ctx, tc.args.uid).Return(followCount, followerCount, tc.want.err)

		t.Run(name, func(t *testing.T) {
			target := NewUserApplication(urvm, usm)

			follows, output, err := target.ListFollow(ctx, tc.args.input, tc.args.uid, tc.args.cuid)
			if !reflect.DeepEqual(err, tc.want.err) {
				t.Fatalf("want %#v, but %#v", tc.want.err, err)
				return
			}

			if !reflect.DeepEqual(follows, tc.want.follows) {
				t.Fatalf("want %#v, but %#v", tc.want.follows, follows)
				return
			}

			if !reflect.DeepEqual(output, tc.want.output) {
				t.Fatalf("want %#v, but %#v", tc.want.output, output)
				return
			}
		})
	}
}

func TestUserApplication_ListFollower(t *testing.T) {
	type args struct {
		uid   string
		cuid  string
		input *input.ListFollower
	}
	type want struct {
		followers []*user.Follower
		output    *output.ListQuery
		err       error
	}

	testCases := map[string]struct {
		args args
		want want
	}{
		"ok": {
			args: args{
				uid:  "00000000-0000-0000-0000-000000000000",
				cuid: "11111111-1111-1111-1111-111111111111",
				input: &input.ListFollower{
					Limit:  100,
					Offset: 0,
				},
			},
			want: want{
				followers: []*user.Follower{
					{
						FollowID:         "11111111-1111-1111-1111-111111111111",
						FollowerID:       "00000000-0000-0000-0000-000000000000",
						Username:         "test-user",
						ThumbnailURL:     "",
						SelfIntroduction: "",
					},
				},
				output: &output.ListQuery{
					Limit:  100,
					Offset: 0,
					Total:  1,
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

		var followCount int
		var followerCount int

		if tc.want.output != nil {
			followerCount = tc.want.output.Total
		}

		urvm := mock_validation.NewMockUserRequestValidation(ctrl)
		urvm.EXPECT().ListFollower(tc.args.input).Return(nil)

		usm := mock_user.NewMockService(ctrl)
		usm.EXPECT().ListFollower(ctx, gomock.Any(), tc.args.cuid).Return(tc.want.followers, tc.want.err)
		usm.EXPECT().ListFriendCount(ctx, tc.args.uid).Return(followCount, followerCount, tc.want.err)

		t.Run(name, func(t *testing.T) {
			target := NewUserApplication(urvm, usm)

			followers, output, err := target.ListFollower(ctx, tc.args.input, tc.args.uid, tc.args.cuid)
			if !reflect.DeepEqual(err, tc.want.err) {
				t.Fatalf("want %#v, but %#v", tc.want.err, err)
				return
			}

			if !reflect.DeepEqual(followers, tc.want.followers) {
				t.Fatalf("want %#v, but %#v", tc.want.followers, followers)
				return
			}

			if !reflect.DeepEqual(output, tc.want.output) {
				t.Fatalf("want %#v, but %#v", tc.want.output, output)
				return
			}
		})
	}
}

func TestUserApplication_Search(t *testing.T) {
	type args struct {
		input *input.SearchUser
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
				input: &input.SearchUser{
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

		arvm := mock_validation.NewMockUserRequestValidation(ctrl)
		arvm.EXPECT().SearchUser(tc.args.input).Return(nil)

		usm := mock_user.NewMockService(ctrl)
		usm.EXPECT().List(ctx, gomock.Any()).Return(tc.want.users, tc.want.err)
		usm.EXPECT().ListCount(ctx, gomock.Any()).Return(tc.want.output.Total, tc.want.err)

		t.Run(name, func(t *testing.T) {
			target := NewUserApplication(arvm, usm)

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

func TestUserApplication_Show(t *testing.T) {
	type args struct {
		uid string
	}
	type want struct {
		user *user.User
		err  error
	}

	current := time.Now().Local()
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

		urvm := mock_validation.NewMockUserRequestValidation(ctrl)

		usm := mock_user.NewMockService(ctrl)
		usm.EXPECT().Show(ctx, tc.args.uid).Return(tc.want.user, tc.want.err)

		t.Run(name, func(t *testing.T) {
			target := NewUserApplication(urvm, usm)

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

func TestUserApplication_GetUserProfile(t *testing.T) {
	type args struct {
		uid  string
		cuid string
	}
	type want struct {
		user   *user.User
		output *output.UserProfile
		err    error
	}

	current := time.Now().Local()
	testCases := map[string]struct {
		args args
		want want
	}{
		"ok": {
			args: args{
				uid:  "00000000-0000-0000-0000-000000000000",
				cuid: "11111111-1111-1111-1111-111111111111",
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
				output: &output.UserProfile{
					IsFollow:      true,
					IsFollower:    false,
					FollowCount:   1,
					FollowerCount: 3,
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

		urvm := mock_validation.NewMockUserRequestValidation(ctrl)

		usm := mock_user.NewMockService(ctrl)
		usm.EXPECT().Show(ctx, tc.args.uid).Return(tc.want.user, nil)
		usm.EXPECT().
			ListFriendCount(ctx, tc.args.uid).
			Return(tc.want.output.FollowCount, tc.want.output.FollowerCount, tc.want.err)
		usm.EXPECT().
			IsFriend(ctx, tc.args.uid, tc.args.cuid).
			Return(tc.want.output.IsFollow, tc.want.output.IsFollower)

		t.Run(name, func(t *testing.T) {
			target := NewUserApplication(urvm, usm)

			user, output, err := target.GetUserProfile(ctx, tc.args.uid, tc.args.cuid)
			if !reflect.DeepEqual(err, tc.want.err) {
				t.Fatalf("want %#v, but %#v", tc.want.err, err)
				return
			}

			if !reflect.DeepEqual(user, tc.want.user) {
				t.Fatalf("want %#v, but %#v", tc.want.user, user)
				return
			}

			if !reflect.DeepEqual(output, tc.want.output) {
				t.Fatalf("want %#v, but %#v", tc.want.output, output)
				return
			}
		})
	}
}

func TestUserApplication_RegisterFollow(t *testing.T) {
	type args struct {
		uid  string
		cuid string
	}
	type want struct {
		user   *user.User
		output *output.UserProfile
		err    error
	}

	current := time.Now().Local()
	testCases := map[string]struct {
		args args
		want want
	}{
		"ok": {
			args: args{
				uid:  "00000000-0000-0000-0000-000000000000",
				cuid: "11111111-1111-1111-1111-111111111111",
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
				output: &output.UserProfile{
					IsFollow:      true,
					IsFollower:    false,
					FollowCount:   1,
					FollowerCount: 3,
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

		r := &user.Relationship{
			FollowID:   tc.args.cuid,
			FollowerID: tc.args.uid,
		}

		urvm := mock_validation.NewMockUserRequestValidation(ctrl)

		usm := mock_user.NewMockService(ctrl)
		usm.EXPECT().Show(ctx, tc.args.uid).Return(tc.want.user, nil)
		usm.EXPECT().ValidationRelationship(ctx, r).Return(nil)
		usm.EXPECT().CreateRelationship(ctx, r).Return(tc.want.err)
		usm.EXPECT().
			ListFriendCount(ctx, tc.args.uid).
			Return(tc.want.output.FollowCount, tc.want.output.FollowerCount, nil)
		usm.EXPECT().
			IsFriend(ctx, tc.args.uid, tc.args.cuid).
			Return(tc.want.output.IsFollow, tc.want.output.IsFollower)

		t.Run(name, func(t *testing.T) {
			target := NewUserApplication(urvm, usm)

			user, output, err := target.RegisterFollow(ctx, tc.args.uid, tc.args.cuid)
			if !reflect.DeepEqual(err, tc.want.err) {
				t.Fatalf("want %#v, but %#v", tc.want.err, err)
				return
			}

			if !reflect.DeepEqual(user, tc.want.user) {
				t.Fatalf("want %#v, but %#v", tc.want.user, user)
				return
			}

			if !reflect.DeepEqual(output, tc.want.output) {
				t.Fatalf("want %#v, but %#v", tc.want.output, output)
				return
			}
		})
	}
}

func TestUserApplication_UnregisterFollow(t *testing.T) {
	type args struct {
		uid  string
		cuid string
	}
	type want struct {
		user   *user.User
		output *output.UserProfile
		err    error
	}

	current := time.Now().Local()
	testCases := map[string]struct {
		args args
		want want
	}{
		"ok": {
			args: args{
				uid:  "00000000-0000-0000-0000-000000000000",
				cuid: "11111111-1111-1111-1111-111111111111",
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
				output: &output.UserProfile{
					IsFollow:      true,
					IsFollower:    false,
					FollowCount:   1,
					FollowerCount: 3,
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

		r := &user.Relationship{
			ID:         1,
			FollowID:   tc.args.cuid,
			FollowerID: tc.args.uid,
		}

		urvm := mock_validation.NewMockUserRequestValidation(ctrl)

		usm := mock_user.NewMockService(ctrl)
		usm.EXPECT().Show(ctx, tc.args.uid).Return(tc.want.user, tc.want.err)
		usm.EXPECT().ShowRelationshipByUID(ctx, tc.args.uid, tc.args.cuid).Return(r, tc.want.err)
		usm.EXPECT().DeleteRelationship(ctx, r.ID).Return(tc.want.err)
		usm.EXPECT().
			ListFriendCount(ctx, tc.args.uid).
			Return(tc.want.output.FollowCount, tc.want.output.FollowerCount, tc.want.err)
		usm.EXPECT().
			IsFriend(ctx, tc.args.uid, tc.args.cuid).
			Return(tc.want.output.IsFollow, tc.want.output.IsFollower)

		t.Run(name, func(t *testing.T) {
			target := NewUserApplication(urvm, usm)

			user, output, err := target.UnregisterFollow(ctx, tc.args.uid, tc.args.cuid)
			if !reflect.DeepEqual(err, tc.want.err) {
				t.Fatalf("want %#v, but %#v", tc.want.err, err)
				return
			}

			if !reflect.DeepEqual(user, tc.want.user) {
				t.Fatalf("want %#v, but %#v", tc.want.user, user)
				return
			}

			if !reflect.DeepEqual(output, tc.want.output) {
				t.Fatalf("want %#v, but %#v", tc.want.output, output)
				return
			}
		})
	}
}
