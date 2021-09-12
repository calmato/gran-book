package entity

import (
	"testing"

	"github.com/calmato/gran-book/api/gateway/native/pkg/test"
	pb "github.com/calmato/gran-book/api/gateway/native/proto/service/user"
	"github.com/stretchr/testify/assert"
)

func TestAuth(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name         string
		auth         *pb.Auth
		expect       *Auth
		expectGender Gender
	}{
		{
			name: "success",
			auth: &pb.Auth{
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
				CreatedAt:        test.TimeMock,
				UpdatedAt:        test.TimeMock,
			},
			expect: &Auth{
				Auth: &pb.Auth{
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
					CreatedAt:        test.TimeMock,
					UpdatedAt:        test.TimeMock,
				},
			},
			expectGender: GenderMan,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			actual := NewAuth(tt.auth)
			assert.Equal(t, tt.expect, actual)
			assert.Equal(t, tt.expectGender, actual.Gender())
		})
	}
}

func TestUser(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name   string
		user   *pb.User
		expect *User
	}{
		{
			name: "success",
			user: &pb.User{
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
				CreatedAt:        test.TimeMock,
				UpdatedAt:        test.TimeMock,
			},
			expect: &User{
				User: &pb.User{
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
					CreatedAt:        test.TimeMock,
					UpdatedAt:        test.TimeMock,
				},
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			actual := NewUser(tt.user)
			assert.Equal(t, tt.expect, actual)
		})
	}
}

func TestUsers(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name      string
		users     []*pb.User
		expect    Users
		expectMap map[string]*User
	}{
		{
			name: "success",
			users: []*pb.User{
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
					CreatedAt:        test.TimeMock,
					UpdatedAt:        test.TimeMock,
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
					CreatedAt:        test.TimeMock,
					UpdatedAt:        test.TimeMock,
				},
			},
			expect: Users{
				{
					User: &pb.User{
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
						CreatedAt:        test.TimeMock,
						UpdatedAt:        test.TimeMock,
					},
				},
				{
					User: &pb.User{
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
						CreatedAt:        test.TimeMock,
						UpdatedAt:        test.TimeMock,
					},
				},
			},
			expectMap: map[string]*User{
				"00000000-0000-0000-0000-000000000000": {
					User: &pb.User{
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
						CreatedAt:        test.TimeMock,
						UpdatedAt:        test.TimeMock,
					},
				},
				"11111111-1111-1111-1111-111111111111": {
					User: &pb.User{
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
						CreatedAt:        test.TimeMock,
						UpdatedAt:        test.TimeMock,
					},
				},
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			actual := NewUsers(tt.users)
			assert.Equal(t, tt.expect, actual)
			assert.Equal(t, tt.expectMap, actual.Map())
		})
	}
}

func TestUserProfile(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		profile *pb.UserProfile
		expect  *UserProfile
	}{
		{
			name: "success",
			profile: &pb.UserProfile{
				Id:               "00000000-0000-0000-0000-000000000000",
				Username:         "テストユーザー",
				ThumbnailUrl:     "https://go.dev//images/gophers/ladder.svg",
				SelfIntroduction: "テストコードです。",
				IsFollow:         false,
				IsFollower:       true,
				FollowCount:      3,
				FollowerCount:    8,
			},
			expect: &UserProfile{
				UserProfile: &pb.UserProfile{
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

func TestFollow(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name   string
		follow *pb.Follow
		expect *Follow
	}{
		{
			name: "success",
			follow: &pb.Follow{
				Id:               "00000000-0000-0000-0000-000000000000",
				Username:         "テストユーザー",
				ThumbnailUrl:     "https://go.dev//images/gophers/ladder.svg",
				SelfIntroduction: "テストコードです。",
				IsFollow:         false,
			},
			expect: &Follow{
				Follow: &pb.Follow{
					Id:               "00000000-0000-0000-0000-000000000000",
					Username:         "テストユーザー",
					ThumbnailUrl:     "https://go.dev//images/gophers/ladder.svg",
					SelfIntroduction: "テストコードです。",
					IsFollow:         false,
				},
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			actual := NewFollow(tt.follow)
			assert.Equal(t, tt.expect, actual)
		})
	}
}

func TestFollows(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		follows []*pb.Follow
		expect  Follows
	}{
		{
			name: "success",
			follows: []*pb.Follow{
				{
					Id:               "00000000-0000-0000-0000-000000000000",
					Username:         "テストユーザー01",
					ThumbnailUrl:     "https://go.dev//images/gophers/ladder.svg",
					SelfIntroduction: "テストコードです。",
					IsFollow:         false,
				},
				{
					Id:               "111111111-1111-1111-1111-111111111111",
					Username:         "テストユーザー02",
					ThumbnailUrl:     "https://go.dev//images/gophers/ladder.svg",
					SelfIntroduction: "テストコードです。",
					IsFollow:         true,
				},
			},
			expect: Follows{
				{
					Follow: &pb.Follow{
						Id:               "00000000-0000-0000-0000-000000000000",
						Username:         "テストユーザー01",
						ThumbnailUrl:     "https://go.dev//images/gophers/ladder.svg",
						SelfIntroduction: "テストコードです。",
						IsFollow:         false,
					},
				},
				{
					Follow: &pb.Follow{
						Id:               "111111111-1111-1111-1111-111111111111",
						Username:         "テストユーザー02",
						ThumbnailUrl:     "https://go.dev//images/gophers/ladder.svg",
						SelfIntroduction: "テストコードです。",
						IsFollow:         true,
					},
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

func TestFollower(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name     string
		follower *pb.Follower
		expect   *Follower
	}{
		{
			name: "success",
			follower: &pb.Follower{
				Id:               "00000000-0000-0000-0000-000000000000",
				Username:         "テストユーザー",
				ThumbnailUrl:     "https://go.dev//images/gophers/ladder.svg",
				SelfIntroduction: "テストコードです。",
				IsFollow:         false,
			},
			expect: &Follower{
				Follower: &pb.Follower{
					Id:               "00000000-0000-0000-0000-000000000000",
					Username:         "テストユーザー",
					ThumbnailUrl:     "https://go.dev//images/gophers/ladder.svg",
					SelfIntroduction: "テストコードです。",
					IsFollow:         false,
				},
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			actual := NewFollower(tt.follower)
			assert.Equal(t, tt.expect, actual)
		})
	}
}

func TestFollowers(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name      string
		followers []*pb.Follower
		expect    Followers
	}{
		{
			name: "success",
			followers: []*pb.Follower{
				{
					Id:               "00000000-0000-0000-0000-000000000000",
					Username:         "テストユーザー01",
					ThumbnailUrl:     "https://go.dev//images/gophers/ladder.svg",
					SelfIntroduction: "テストコードです。",
					IsFollow:         false,
				},
				{
					Id:               "111111111-1111-1111-1111-111111111111",
					Username:         "テストユーザー02",
					ThumbnailUrl:     "https://go.dev//images/gophers/ladder.svg",
					SelfIntroduction: "テストコードです。",
					IsFollow:         true,
				},
			},
			expect: Followers{
				{
					Follower: &pb.Follower{
						Id:               "00000000-0000-0000-0000-000000000000",
						Username:         "テストユーザー01",
						ThumbnailUrl:     "https://go.dev//images/gophers/ladder.svg",
						SelfIntroduction: "テストコードです。",
						IsFollow:         false,
					},
				},
				{
					Follower: &pb.Follower{
						Id:               "111111111-1111-1111-1111-111111111111",
						Username:         "テストユーザー02",
						ThumbnailUrl:     "https://go.dev//images/gophers/ladder.svg",
						SelfIntroduction: "テストコードです。",
						IsFollow:         true,
					},
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
