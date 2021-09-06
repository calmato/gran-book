package user

import (
	"testing"
	"time"

	"github.com/calmato/gran-book/api/server/user/pkg/datetime"
	pb "github.com/calmato/gran-book/api/server/user/proto"
	"github.com/stretchr/testify/assert"
)

func TestUser(t *testing.T) {
	t.Parallel()
	now := time.Now().Local()
	tests := []struct {
		name          string
		user          *User
		expectProto   *pb.User
		expectAuth    *pb.Auth
		expectAdmin   *pb.Admin
		expectProfile *pb.UserProfile
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
			expectProto: &pb.User{
				Id:               "00000000-0000-0000-0000-000000000000",
				Username:         "テストユーザー",
				Gender:           pb.Gender_GENDER_MAN,
				Email:            "test-user@calmato.jp",
				PhoneNumber:      "000-0000-0000",
				ThumbnailUrl:     "https://go.dev/images/gophers/ladder.svg",
				SelfIntroduction: "テストコードです。",
				LastName:         "テスト",
				FirstName:        "ユーザー",
				LastNameKana:     "てすと",
				FirstNameKana:    "ゆーざー",
				CreatedAt:        datetime.TimeToString(now),
				UpdatedAt:        datetime.TimeToString(now),
			},
			expectAuth: &pb.Auth{
				Id:               "00000000-0000-0000-0000-000000000000",
				Username:         "テストユーザー",
				Gender:           pb.Gender_GENDER_MAN,
				Email:            "test-user@calmato.jp",
				PhoneNumber:      "000-0000-0000",
				Role:             pb.Role_ROLE_USER,
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
				CreatedAt:        datetime.TimeToString(now),
				UpdatedAt:        datetime.TimeToString(now),
			},
			expectAdmin: &pb.Admin{
				Id:               "00000000-0000-0000-0000-000000000000",
				Username:         "テストユーザー",
				Email:            "test-user@calmato.jp",
				PhoneNumber:      "000-0000-0000",
				Role:             pb.Role_ROLE_USER,
				ThumbnailUrl:     "https://go.dev/images/gophers/ladder.svg",
				SelfIntroduction: "テストコードです。",
				LastName:         "テスト",
				FirstName:        "ユーザー",
				LastNameKana:     "てすと",
				FirstNameKana:    "ゆーざー",
				CreatedAt:        datetime.TimeToString(now),
				UpdatedAt:        datetime.TimeToString(now),
			},
			expectProfile: &pb.UserProfile{
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
		expectProto []*pb.User
		expectAdmin []*pb.Admin
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
			expectProto: []*pb.User{
				{
					Id:               "00000000-0000-0000-0000-000000000000",
					Username:         "テストユーザー",
					Gender:           pb.Gender_GENDER_MAN,
					Email:            "test-user01@calmato.jp",
					PhoneNumber:      "000-0000-0000",
					ThumbnailUrl:     "https://go.dev/images/gophers/ladder.svg",
					SelfIntroduction: "テストコードです。",
					LastName:         "テスト",
					FirstName:        "ユーザー",
					LastNameKana:     "てすと",
					FirstNameKana:    "ゆーざー",
					CreatedAt:        datetime.TimeToString(now),
					UpdatedAt:        datetime.TimeToString(now),
				},
				{
					Id:               "11111111-1111-1111-1111-111111111111",
					Username:         "テストユーザー",
					Gender:           pb.Gender_GENDER_MAN,
					Email:            "test-user02@calmato.jp",
					PhoneNumber:      "000-0000-0000",
					ThumbnailUrl:     "https://go.dev/images/gophers/ladder.svg",
					SelfIntroduction: "テストコードです。",
					LastName:         "テスト",
					FirstName:        "ユーザー",
					LastNameKana:     "てすと",
					FirstNameKana:    "ゆーざー",
					CreatedAt:        datetime.TimeToString(now),
					UpdatedAt:        datetime.TimeToString(now),
				},
			},
			expectAdmin: []*pb.Admin{
				{
					Id:               "00000000-0000-0000-0000-000000000000",
					Username:         "テストユーザー",
					Email:            "test-user01@calmato.jp",
					PhoneNumber:      "000-0000-0000",
					Role:             pb.Role_ROLE_USER,
					ThumbnailUrl:     "https://go.dev/images/gophers/ladder.svg",
					SelfIntroduction: "テストコードです。",
					LastName:         "テスト",
					FirstName:        "ユーザー",
					LastNameKana:     "てすと",
					FirstNameKana:    "ゆーざー",
					CreatedAt:        datetime.TimeToString(now),
					UpdatedAt:        datetime.TimeToString(now),
				},
				{
					Id:               "11111111-1111-1111-1111-111111111111",
					Username:         "テストユーザー",
					Email:            "test-user02@calmato.jp",
					PhoneNumber:      "000-0000-0000",
					Role:             pb.Role_ROLE_USER,
					ThumbnailUrl:     "https://go.dev/images/gophers/ladder.svg",
					SelfIntroduction: "テストコードです。",
					LastName:         "テスト",
					FirstName:        "ユーザー",
					LastNameKana:     "てすと",
					FirstNameKana:    "ゆーざー",
					CreatedAt:        datetime.TimeToString(now),
					UpdatedAt:        datetime.TimeToString(now),
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
		expectProto *pb.Follow
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
			expectProto: &pb.Follow{
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
		expectProto []*pb.Follow
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
			expectProto: []*pb.Follow{
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
		expectProto *pb.Follower
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
			expectProto: &pb.Follower{
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
		expectProto []*pb.Follower
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
			expectProto: []*pb.Follower{
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
