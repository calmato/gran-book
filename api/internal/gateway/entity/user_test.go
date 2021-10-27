package entity

import (
	"testing"

	"github.com/calmato/gran-book/api/pkg/datetime"
	"github.com/calmato/gran-book/api/pkg/test"
	"github.com/calmato/gran-book/api/proto/user"
	"github.com/stretchr/testify/assert"
)

func TestAuth(t *testing.T) {
	t.Parallel()
	now := datetime.FormatTime(test.Now())
	tests := []struct {
		name   string
		auth   *user.Auth
		expect *Auth
	}{
		{
			name: "success",
			auth: &user.Auth{
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
				CreatedAt:        now,
				UpdatedAt:        now,
			},
			expect: &Auth{
				Auth: &user.Auth{
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
					CreatedAt:        now,
					UpdatedAt:        now,
				},
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			actual := NewAuth(tt.auth)
			assert.Equal(t, tt.expect, actual)
		})
	}
}

func TestAuth_Gender(t *testing.T) {
	t.Parallel()
	now := datetime.FormatTime(test.Now())
	tests := []struct {
		name   string
		auth   *user.Auth
		expect Gender
	}{
		{
			name: "success",
			auth: &user.Auth{
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
				CreatedAt:        now,
				UpdatedAt:        now,
			},
			expect: GenderMan,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			actual := NewAuth(tt.auth)
			assert.Equal(t, tt.expect, actual.Gender())
		})
	}
}

func TestAuth_Role(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name   string
		auth   *Auth
		expect Role
	}{
		{
			name: "success",
			auth: &Auth{
				Auth: &user.Auth{
					Role: user.Role_ROLE_ADMIN,
				},
			},
			expect: RoleAdmin,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.expect, tt.auth.Role())
		})
	}
}

func TestUser(t *testing.T) {
	t.Parallel()
	now := datetime.FormatTime(test.Now())
	tests := []struct {
		name   string
		user   *user.User
		expect *User
	}{
		{
			name: "success",
			user: &user.User{
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
				CreatedAt:        now,
				UpdatedAt:        now,
			},
			expect: &User{
				User: &user.User{
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
					CreatedAt:        now,
					UpdatedAt:        now,
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
	now := datetime.FormatTime(test.Now())
	tests := []struct {
		name   string
		users  []*user.User
		expect Users
	}{
		{
			name: "success",
			users: []*user.User{
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
					CreatedAt:        now,
					UpdatedAt:        now,
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
					CreatedAt:        now,
					UpdatedAt:        now,
				},
			},
			expect: Users{
				{
					User: &user.User{
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
						CreatedAt:        now,
						UpdatedAt:        now,
					},
				},
				{
					User: &user.User{
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
						CreatedAt:        now,
						UpdatedAt:        now,
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
		})
	}
}

func TestUsers_Map(t *testing.T) {
	t.Parallel()
	now := datetime.FormatTime(test.Now())
	tests := []struct {
		name   string
		users  []*user.User
		expect map[string]*User
	}{
		{
			name: "success",
			users: []*user.User{
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
					CreatedAt:        now,
					UpdatedAt:        now,
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
					CreatedAt:        now,
					UpdatedAt:        now,
				},
			},
			expect: map[string]*User{
				"00000000-0000-0000-0000-000000000000": {
					User: &user.User{
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
						CreatedAt:        now,
						UpdatedAt:        now,
					},
				},
				"11111111-1111-1111-1111-111111111111": {
					User: &user.User{
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
						CreatedAt:        now,
						UpdatedAt:        now,
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
			assert.Equal(t, tt.expect, actual.Map())
		})
	}
}

func TestUsers_IsExists(t *testing.T) {
	t.Parallel()
	now := datetime.FormatTime(test.Now())
	tests := []struct {
		name    string
		users   []*user.User
		userIDs []string
		expect  bool
	}{
		{
			name: "success",
			users: []*user.User{
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
					CreatedAt:        now,
					UpdatedAt:        now,
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
					CreatedAt:        now,
					UpdatedAt:        now,
				},
			},
			userIDs: []string{
				"00000000-0000-0000-0000-000000000000",
				"11111111-1111-1111-1111-111111111111",
			},
			expect: true,
		},
		{
			name: "failed",
			users: []*user.User{
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
					CreatedAt:        now,
					UpdatedAt:        now,
				},
			},
			userIDs: []string{
				"00000000-0000-0000-0000-000000000000",
				"11111111-1111-1111-1111-111111111111",
			},
			expect: false,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			actual := NewUsers(tt.users)
			assert.Equal(t, tt.expect, actual.IsExists(tt.userIDs...))
		})
	}
}

func TestUserProfile(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		profile *user.UserProfile
		expect  *UserProfile
	}{
		{
			name: "success",
			profile: &user.UserProfile{
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
		follow *user.Follow
		expect *Follow
	}{
		{
			name: "success",
			follow: &user.Follow{
				Id:               "00000000-0000-0000-0000-000000000000",
				Username:         "テストユーザー",
				ThumbnailUrl:     "https://go.dev//images/gophers/ladder.svg",
				SelfIntroduction: "テストコードです。",
				IsFollow:         false,
			},
			expect: &Follow{
				Follow: &user.Follow{
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
		follows []*user.Follow
		expect  Follows
	}{
		{
			name: "success",
			follows: []*user.Follow{
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
		follower *user.Follower
		expect   *Follower
	}{
		{
			name: "success",
			follower: &user.Follower{
				Id:               "00000000-0000-0000-0000-000000000000",
				Username:         "テストユーザー",
				ThumbnailUrl:     "https://go.dev//images/gophers/ladder.svg",
				SelfIntroduction: "テストコードです。",
				IsFollow:         false,
			},
			expect: &Follower{
				Follower: &user.Follower{
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
		followers []*user.Follower
		expect    Followers
	}{
		{
			name: "success",
			followers: []*user.Follower{
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

func TestAdmin(t *testing.T) {
	t.Parallel()
	now := datetime.FormatTime(test.Now())
	tests := []struct {
		name   string
		admin  *user.Admin
		expect *Admin
	}{
		{
			name: "success",
			admin: &user.Admin{
				Id:               "00000000-0000-0000-0000-000000000000",
				Username:         "テストユーザー",
				Role:             user.Role_ROLE_ADMIN,
				Email:            "test-user01@calmato.jp",
				PhoneNumber:      "000-0000-0000",
				ThumbnailUrl:     "https://go.dev/images/gophers/ladder.svg",
				SelfIntroduction: "テストコードです。",
				LastName:         "テスト",
				FirstName:        "ユーザー",
				LastNameKana:     "てすと",
				FirstNameKana:    "ゆーざー",
				CreatedAt:        now,
				UpdatedAt:        now,
			},
			expect: &Admin{
				Admin: &user.Admin{
					Id:               "00000000-0000-0000-0000-000000000000",
					Username:         "テストユーザー",
					Role:             user.Role_ROLE_ADMIN,
					Email:            "test-user01@calmato.jp",
					PhoneNumber:      "000-0000-0000",
					ThumbnailUrl:     "https://go.dev/images/gophers/ladder.svg",
					SelfIntroduction: "テストコードです。",
					LastName:         "テスト",
					FirstName:        "ユーザー",
					LastNameKana:     "てすと",
					FirstNameKana:    "ゆーざー",
					CreatedAt:        now,
					UpdatedAt:        now,
				},
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			actual := NewAdmin(tt.admin)
			assert.Equal(t, tt.expect, actual)
		})
	}
}

func TestAdmin_Role(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name   string
		admin  *Admin
		expect Role
	}{
		{
			name: "success",
			admin: &Admin{
				Admin: &user.Admin{
					Role: user.Role_ROLE_ADMIN,
				},
			},
			expect: RoleAdmin,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.expect, tt.admin.Role())
		})
	}
}

func TestAdmins(t *testing.T) {
	t.Parallel()
	now := datetime.FormatTime(test.Now())
	tests := []struct {
		name   string
		admins []*user.Admin
		expect Admins
	}{
		{
			name: "success",
			admins: []*user.Admin{
				{
					Id:               "00000000-0000-0000-0000-000000000000",
					Username:         "テストユーザー",
					Role:             user.Role_ROLE_ADMIN,
					Email:            "test-user01@calmato.jp",
					PhoneNumber:      "000-0000-0000",
					ThumbnailUrl:     "https://go.dev/images/gophers/ladder.svg",
					SelfIntroduction: "テストコードです。",
					LastName:         "テスト",
					FirstName:        "ユーザー",
					LastNameKana:     "てすと",
					FirstNameKana:    "ゆーざー",
					CreatedAt:        now,
					UpdatedAt:        now,
				},
				{
					Id:               "11111111-1111-1111-1111-111111111111",
					Username:         "テストユーザー",
					Role:             user.Role_ROLE_ADMIN,
					Email:            "test-user02@calmato.jp",
					PhoneNumber:      "000-0000-0000",
					ThumbnailUrl:     "https://go.dev/images/gophers/ladder.svg",
					SelfIntroduction: "テストコードです。",
					LastName:         "テスト",
					FirstName:        "ユーザー",
					LastNameKana:     "てすと",
					FirstNameKana:    "ゆーざー",
					CreatedAt:        now,
					UpdatedAt:        now,
				},
			},
			expect: Admins{
				{
					Admin: &user.Admin{
						Id:               "00000000-0000-0000-0000-000000000000",
						Username:         "テストユーザー",
						Role:             user.Role_ROLE_ADMIN,
						Email:            "test-user01@calmato.jp",
						PhoneNumber:      "000-0000-0000",
						ThumbnailUrl:     "https://go.dev/images/gophers/ladder.svg",
						SelfIntroduction: "テストコードです。",
						LastName:         "テスト",
						FirstName:        "ユーザー",
						LastNameKana:     "てすと",
						FirstNameKana:    "ゆーざー",
						CreatedAt:        now,
						UpdatedAt:        now,
					},
				},
				{
					Admin: &user.Admin{
						Id:               "11111111-1111-1111-1111-111111111111",
						Username:         "テストユーザー",
						Role:             user.Role_ROLE_ADMIN,
						Email:            "test-user02@calmato.jp",
						PhoneNumber:      "000-0000-0000",
						ThumbnailUrl:     "https://go.dev/images/gophers/ladder.svg",
						SelfIntroduction: "テストコードです。",
						LastName:         "テスト",
						FirstName:        "ユーザー",
						LastNameKana:     "てすと",
						FirstNameKana:    "ゆーざー",
						CreatedAt:        now,
						UpdatedAt:        now,
					},
				},
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			actual := NewAdmins(tt.admins)
			assert.Equal(t, tt.expect, actual)
		})
	}
}
