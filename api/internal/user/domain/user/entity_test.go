package user

import (
	"testing"
	"time"

	"github.com/calmato/gran-book/api/pkg/datetime"
	"github.com/calmato/gran-book/api/proto/user"
	"github.com/stretchr/testify/assert"
)

func TestUser(t *testing.T) {
	t.Parallel()
	now := time.Now().Local()
	tests := []struct {
		name          string
		user          *User
		expectProto   *user.User
		expectAuth    *user.Auth
		expectAdmin   *user.Admin
		expectProfile *user.UserProfile
	}{
		{
			name: "success",
			user: &User{
				ID:               "00000000-0000-0000-0000-000000000000",
				Username:         "テストユーザー",
				Gender:           MaleGender,
				Email:            "test-user@calmato.jp",
				PhoneNumber:      "000-0000-0000",
				Role:             UserRole,
				ThumbnailURL:     "https://go.dev/images/gophers/ladder.svg",
				SelfIntroduction: "テストコードです。",
				LastName:         "テスト",
				FirstName:        "ユーザー",
				LastNameKana:     "てすと",
				FirstNameKana:    "ゆーざー",
				PostalCode:       "000-0000",
				Prefecture:       "東京都",
				City:             "小金井市",
				AddressLine1:     "貫井北町4-1-1",
				AddressLine2:     "",
				InstanceID:       "ExponentPushToken[!Qaz2wsx3edc4rfv5tgb6y]",
				CreatedAt:        now,
				UpdatedAt:        now,
			},
			expectProto: &user.User{
				Id:               "00000000-0000-0000-0000-000000000000",
				Username:         "テストユーザー",
				Gender:           user.Gender_GENDER_MAN,
				Email:            "test-user@calmato.jp",
				PhoneNumber:      "000-0000-0000",
				ThumbnailUrl:     "https://go.dev/images/gophers/ladder.svg",
				SelfIntroduction: "テストコードです。",
				LastName:         "テスト",
				FirstName:        "ユーザー",
				LastNameKana:     "てすと",
				FirstNameKana:    "ゆーざー",
				CreatedAt:        datetime.FormatTime(now),
				UpdatedAt:        datetime.FormatTime(now),
			},
			expectAuth: &user.Auth{
				Id:               "00000000-0000-0000-0000-000000000000",
				Username:         "テストユーザー",
				Gender:           user.Gender_GENDER_MAN,
				Email:            "test-user@calmato.jp",
				PhoneNumber:      "000-0000-0000",
				Role:             user.Role_ROLE_USER,
				ThumbnailUrl:     "https://go.dev/images/gophers/ladder.svg",
				SelfIntroduction: "テストコードです。",
				LastName:         "テスト",
				FirstName:        "ユーザー",
				LastNameKana:     "てすと",
				FirstNameKana:    "ゆーざー",
				PostalCode:       "000-0000",
				Prefecture:       "東京都",
				City:             "小金井市",
				AddressLine1:     "貫井北町4-1-1",
				AddressLine2:     "",
				CreatedAt:        datetime.FormatTime(now),
				UpdatedAt:        datetime.FormatTime(now),
			},
			expectAdmin: &user.Admin{
				Id:               "00000000-0000-0000-0000-000000000000",
				Username:         "テストユーザー",
				Email:            "test-user@calmato.jp",
				PhoneNumber:      "000-0000-0000",
				Role:             user.Role_ROLE_USER,
				ThumbnailUrl:     "https://go.dev/images/gophers/ladder.svg",
				SelfIntroduction: "テストコードです。",
				LastName:         "テスト",
				FirstName:        "ユーザー",
				LastNameKana:     "てすと",
				FirstNameKana:    "ゆーざー",
				CreatedAt:        datetime.FormatTime(now),
				UpdatedAt:        datetime.FormatTime(now),
			},
			expectProfile: &user.UserProfile{
				Id:               "00000000-0000-0000-0000-000000000000",
				Username:         "テストユーザー",
				ThumbnailUrl:     "https://go.dev/images/gophers/ladder.svg",
				SelfIntroduction: "テストコードです。",
				IsFollow:         false,
				IsFollower:       false,
				FollowCount:      0,
				FollowerCount:    0,
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.expectProto, tt.user.Proto())
			assert.Equal(t, tt.expectAuth, tt.user.Auth())
			assert.Equal(t, tt.expectAdmin, tt.user.Admin())
			assert.Equal(t, tt.expectProfile, tt.user.Profile())
		})
	}
}

func TestUsers(t *testing.T) {
	t.Parallel()
	now := time.Now().Local()
	tests := []struct {
		name        string
		users       Users
		expectProto []*user.User
		expectAdmin []*user.Admin
	}{
		{
			name: "success",
			users: Users{
				{
					ID:               "00000000-0000-0000-0000-000000000000",
					Username:         "テストユーザー",
					Gender:           MaleGender,
					Email:            "test-user01@calmato.jp",
					PhoneNumber:      "000-0000-0000",
					Role:             UserRole,
					ThumbnailURL:     "https://go.dev/images/gophers/ladder.svg",
					SelfIntroduction: "テストコードです。",
					LastName:         "テスト",
					FirstName:        "ユーザー",
					LastNameKana:     "てすと",
					FirstNameKana:    "ゆーざー",
					PostalCode:       "000-0000",
					Prefecture:       "東京都",
					City:             "小金井市",
					AddressLine1:     "貫井北町4-1-1",
					AddressLine2:     "",
					InstanceID:       "ExponentPushToken[!Qaz2wsx3edc4rfv5tgb6y]",
					CreatedAt:        now,
					UpdatedAt:        now,
				},
				{
					ID:               "11111111-1111-1111-1111-111111111111",
					Username:         "テストユーザー",
					Gender:           MaleGender,
					Email:            "test-user02@calmato.jp",
					PhoneNumber:      "000-0000-0000",
					Role:             UserRole,
					ThumbnailURL:     "https://go.dev/images/gophers/ladder.svg",
					SelfIntroduction: "テストコードです。",
					LastName:         "テスト",
					FirstName:        "ユーザー",
					LastNameKana:     "てすと",
					FirstNameKana:    "ゆーざー",
					PostalCode:       "000-0000",
					Prefecture:       "東京都",
					City:             "小金井市",
					AddressLine1:     "貫井北町4-1-1",
					AddressLine2:     "",
					InstanceID:       "ExponentPushToken[!Qaz2wsx3edc4rfv5tgb6y]",
					CreatedAt:        now,
					UpdatedAt:        now,
				},
			},
			expectProto: []*user.User{
				{
					Id:               "00000000-0000-0000-0000-000000000000",
					Username:         "テストユーザー",
					Gender:           user.Gender_GENDER_MAN,
					Email:            "test-user01@calmato.jp",
					PhoneNumber:      "000-0000-0000",
					ThumbnailUrl:     "https://go.dev/images/gophers/ladder.svg",
					SelfIntroduction: "テストコードです。",
					LastName:         "テスト",
					FirstName:        "ユーザー",
					LastNameKana:     "てすと",
					FirstNameKana:    "ゆーざー",
					CreatedAt:        datetime.FormatTime(now),
					UpdatedAt:        datetime.FormatTime(now),
				},
				{
					Id:               "11111111-1111-1111-1111-111111111111",
					Username:         "テストユーザー",
					Gender:           user.Gender_GENDER_MAN,
					Email:            "test-user02@calmato.jp",
					PhoneNumber:      "000-0000-0000",
					ThumbnailUrl:     "https://go.dev/images/gophers/ladder.svg",
					SelfIntroduction: "テストコードです。",
					LastName:         "テスト",
					FirstName:        "ユーザー",
					LastNameKana:     "てすと",
					FirstNameKana:    "ゆーざー",
					CreatedAt:        datetime.FormatTime(now),
					UpdatedAt:        datetime.FormatTime(now),
				},
			},
			expectAdmin: []*user.Admin{
				{
					Id:               "00000000-0000-0000-0000-000000000000",
					Username:         "テストユーザー",
					Email:            "test-user01@calmato.jp",
					PhoneNumber:      "000-0000-0000",
					Role:             user.Role_ROLE_USER,
					ThumbnailUrl:     "https://go.dev/images/gophers/ladder.svg",
					SelfIntroduction: "テストコードです。",
					LastName:         "テスト",
					FirstName:        "ユーザー",
					LastNameKana:     "てすと",
					FirstNameKana:    "ゆーざー",
					CreatedAt:        datetime.FormatTime(now),
					UpdatedAt:        datetime.FormatTime(now),
				},
				{
					Id:               "11111111-1111-1111-1111-111111111111",
					Username:         "テストユーザー",
					Email:            "test-user02@calmato.jp",
					PhoneNumber:      "000-0000-0000",
					Role:             user.Role_ROLE_USER,
					ThumbnailUrl:     "https://go.dev/images/gophers/ladder.svg",
					SelfIntroduction: "テストコードです。",
					LastName:         "テスト",
					FirstName:        "ユーザー",
					LastNameKana:     "てすと",
					FirstNameKana:    "ゆーざー",
					CreatedAt:        datetime.FormatTime(now),
					UpdatedAt:        datetime.FormatTime(now),
				},
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.expectProto, tt.users.Proto())
			assert.Equal(t, tt.expectAdmin, tt.users.Admin())
		})
	}
}

func TestFollow(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name        string
		follow      *Follow
		expectProto *user.Follow
	}{
		{
			name: "success",
			follow: &Follow{
				FollowID:         "00000000-0000-0000-0000-000000000000",
				Username:         "テストユーザー",
				ThumbnailURL:     "https://go.dev/images/gophers/ladder.svg",
				SelfIntroduction: "テストコードです。",
				IsFollowing:      true,
				IsFollowed:       false,
				FollowCount:      100,
				FollowerCount:    64,
			},
			expectProto: &user.Follow{
				Id:               "00000000-0000-0000-0000-000000000000",
				Username:         "テストユーザー",
				ThumbnailUrl:     "https://go.dev/images/gophers/ladder.svg",
				SelfIntroduction: "テストコードです。",
				IsFollow:         true,
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.expectProto, tt.follow.Proto())
		})
	}
}

func TestFollows(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name        string
		follows     Follows
		expectProto []*user.Follow
	}{
		{
			name: "success",
			follows: Follows{
				{
					FollowID:         "00000000-0000-0000-0000-000000000000",
					Username:         "テストユーザー",
					ThumbnailURL:     "https://go.dev/images/gophers/ladder.svg",
					SelfIntroduction: "テストコードです。",
					IsFollowing:      true,
					IsFollowed:       false,
					FollowCount:      100,
					FollowerCount:    64,
				},
				{
					FollowID:         "11111111-1111-1111-1111-111111111111",
					Username:         "テストユーザー",
					ThumbnailURL:     "https://go.dev/images/gophers/ladder.svg",
					SelfIntroduction: "テストコードです。",
					IsFollowing:      true,
					IsFollowed:       false,
					FollowCount:      100,
					FollowerCount:    64,
				},
			},
			expectProto: []*user.Follow{
				{
					Id:               "00000000-0000-0000-0000-000000000000",
					Username:         "テストユーザー",
					ThumbnailUrl:     "https://go.dev/images/gophers/ladder.svg",
					SelfIntroduction: "テストコードです。",
					IsFollow:         true,
				},
				{
					Id:               "11111111-1111-1111-1111-111111111111",
					Username:         "テストユーザー",
					ThumbnailUrl:     "https://go.dev/images/gophers/ladder.svg",
					SelfIntroduction: "テストコードです。",
					IsFollow:         true,
				},
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.expectProto, tt.follows.Proto())
		})
	}
}

func TestFollower(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name        string
		follower    *Follower
		expectProto *user.Follower
	}{
		{
			name: "success",
			follower: &Follower{
				FollowerID:       "00000000-0000-0000-0000-000000000000",
				Username:         "テストユーザー",
				ThumbnailURL:     "https://go.dev/images/gophers/ladder.svg",
				SelfIntroduction: "テストコードです。",
				IsFollowing:      true,
				IsFollowed:       false,
				FollowCount:      100,
				FollowerCount:    64,
			},
			expectProto: &user.Follower{
				Id:               "00000000-0000-0000-0000-000000000000",
				Username:         "テストユーザー",
				ThumbnailUrl:     "https://go.dev/images/gophers/ladder.svg",
				SelfIntroduction: "テストコードです。",
				IsFollow:         true,
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.expectProto, tt.follower.Proto())
		})
	}
}

func TestFollowers(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name        string
		followers   Followers
		expectProto []*user.Follower
	}{
		{
			name: "success",
			followers: Followers{
				{
					FollowerID:       "00000000-0000-0000-0000-000000000000",
					Username:         "テストユーザー",
					ThumbnailURL:     "https://go.dev/images/gophers/ladder.svg",
					SelfIntroduction: "テストコードです。",
					IsFollowing:      true,
					IsFollowed:       false,
					FollowCount:      100,
					FollowerCount:    64,
				},
				{
					FollowerID:       "11111111-1111-1111-1111-111111111111",
					Username:         "テストユーザー",
					ThumbnailURL:     "https://go.dev/images/gophers/ladder.svg",
					SelfIntroduction: "テストコードです。",
					IsFollowing:      true,
					IsFollowed:       false,
					FollowCount:      100,
					FollowerCount:    64,
				},
			},
			expectProto: []*user.Follower{
				{
					Id:               "00000000-0000-0000-0000-000000000000",
					Username:         "テストユーザー",
					ThumbnailUrl:     "https://go.dev/images/gophers/ladder.svg",
					SelfIntroduction: "テストコードです。",
					IsFollow:         true,
				},
				{
					Id:               "11111111-1111-1111-1111-111111111111",
					Username:         "テストユーザー",
					ThumbnailUrl:     "https://go.dev/images/gophers/ladder.svg",
					SelfIntroduction: "テストコードです。",
					IsFollow:         true,
				},
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.expectProto, tt.followers.Proto())
		})
	}
}
