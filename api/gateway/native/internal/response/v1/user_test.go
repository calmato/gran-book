package v1

import (
	"testing"

	"github.com/calmato/gran-book/api/gateway/native/internal/entity"
	"github.com/calmato/gran-book/api/gateway/native/proto/service/user"
	"github.com/stretchr/testify/assert"
)

func TestFollowListResponse(t *testing.T) {
	t.Parallel()
	type args struct {
		follows entity.Follows
		limit   int64
		offset  int64
		total   int64
	}
	tests := []struct {
		name   string
		args   args
		expect *FollowListResponse
	}{
		{
			name: "success",
			args: args{
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
				limit:  100,
				offset: 0,
				total:  2,
			},
			expect: &FollowListResponse{
				Users: []*followListUser{
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
				Limit:  100,
				Offset: 0,
				Total:  2,
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			actual := NewFollowListResponse(tt.args.follows, tt.args.limit, tt.args.offset, tt.args.total)
			assert.Equal(t, tt.expect, actual)
		})
	}
}

func TestFollowerListResponse(t *testing.T) {
	t.Parallel()
	type args struct {
		followers entity.Followers
		limit     int64
		offset    int64
		total     int64
	}
	tests := []struct {
		name   string
		args   args
		expect *FollowerListResponse
	}{
		{
			name: "success",
			args: args{
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
				limit:  100,
				offset: 0,
				total:  2,
			},
			expect: &FollowerListResponse{
				Users: []*followerListUser{
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
				Limit:  100,
				Offset: 0,
				Total:  2,
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			actual := NewFollowerListResponse(tt.args.followers, tt.args.limit, tt.args.offset, tt.args.total)
			assert.Equal(t, tt.expect, actual)
		})
	}
}

func TestUserProfileResponse(t *testing.T) {
	t.Parallel()
	type args struct {
		profile *entity.UserProfile
	}
	tests := []struct {
		name   string
		args   args
		expect *UserProfileResponse
	}{
		{
			name: "success",
			args: args{
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
			},
			expect: &UserProfileResponse{
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
				Products:         make([]*userProfileProduct, 0),
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			actual := NewUserProfileResponse(tt.args.profile)
			assert.Equal(t, tt.expect, actual)
		})
	}
}
