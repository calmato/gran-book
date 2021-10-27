package entity

import (
	"testing"

	"github.com/calmato/gran-book/api/internal/gateway/entity"
	"github.com/calmato/gran-book/api/proto/user"
	"github.com/stretchr/testify/assert"
)

func TestFollows(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		follows entity.Follows
		expect  Follows
	}{
		{
			name: "success",
			follows: entity.Follows{
				{
					Follow: &user.Follow{
						Id:               "00000000-0000-0000-0000-000000000000",
						Username:         "テストユーザー01",
						ThumbnailUrl:     "https://go.dev//images/gophers/ladder.svg",
						SelfIntroduction: "テストコードです。",
						IsFollow:         false,
					},
				},
				{
					Follow: &user.Follow{
						Id:               "111111111-1111-1111-1111-111111111111",
						Username:         "テストユーザー02",
						ThumbnailUrl:     "https://go.dev//images/gophers/ladder.svg",
						SelfIntroduction: "テストコードです。",
						IsFollow:         true,
					},
				},
			},
			expect: Follows{
				{
					ID:               "00000000-0000-0000-0000-000000000000",
					Username:         "テストユーザー01",
					ThumbnailURL:     "https://go.dev//images/gophers/ladder.svg",
					SelfIntroduction: "テストコードです。",
					IsFollow:         false,
				},
				{
					ID:               "111111111-1111-1111-1111-111111111111",
					Username:         "テストユーザー02",
					ThumbnailURL:     "https://go.dev//images/gophers/ladder.svg",
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
			actual := NewFollows(tt.follows)
			assert.Equal(t, tt.expect, actual)
		})
	}
}

func TestFollowers(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name      string
		followers entity.Followers
		expect    Followers
	}{
		{
			name: "success",
			followers: entity.Followers{
				{
					Follower: &user.Follower{
						Id:               "00000000-0000-0000-0000-000000000000",
						Username:         "テストユーザー01",
						ThumbnailUrl:     "https://go.dev//images/gophers/ladder.svg",
						SelfIntroduction: "テストコードです。",
						IsFollow:         false,
					},
				},
				{
					Follower: &user.Follower{
						Id:               "111111111-1111-1111-1111-111111111111",
						Username:         "テストユーザー02",
						ThumbnailUrl:     "https://go.dev//images/gophers/ladder.svg",
						SelfIntroduction: "テストコードです。",
						IsFollow:         true,
					},
				},
			},
			expect: Followers{
				{
					ID:               "00000000-0000-0000-0000-000000000000",
					Username:         "テストユーザー01",
					ThumbnailURL:     "https://go.dev//images/gophers/ladder.svg",
					SelfIntroduction: "テストコードです。",
					IsFollow:         false,
				},
				{
					ID:               "111111111-1111-1111-1111-111111111111",
					Username:         "テストユーザー02",
					ThumbnailURL:     "https://go.dev//images/gophers/ladder.svg",
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
			actual := NewFollowers(tt.followers)
			assert.Equal(t, tt.expect, actual)
		})
	}
}

func TestUserProfile(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		profile *entity.UserProfile
		expect  *UserProfile
	}{
		{
			name: "success",
			profile: &entity.UserProfile{
				UserProfile: &user.UserProfile{
					Id:               "00000000-0000-0000-0000-000000000000",
					Username:         "テストユーザー",
					ThumbnailUrl:     "https://go.dev//images/gophers/ladder.svg",
					SelfIntroduction: "テストコードです。",
					IsFollow:         false,
					IsFollower:       true,
					FollowCount:      3,
					FollowerCount:    8,
				},
			},
			expect: &UserProfile{
				ID:               "00000000-0000-0000-0000-000000000000",
				Username:         "テストユーザー",
				ThumbnailURL:     "https://go.dev//images/gophers/ladder.svg",
				SelfIntroduction: "テストコードです。",
				IsFollow:         false,
				IsFollower:       true,
				FollowCount:      3,
				FollowerCount:    8,
				Rating:           0,
				ReviewCount:      0,
				Products:         make([]*UserProfileProduct, 0),
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			actual := NewUserProfile(tt.profile)
			assert.Equal(t, tt.expect, actual)
		})
	}
}
