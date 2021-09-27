package validation

import (
	"strings"
	"testing"

	pb "github.com/calmato/gran-book/api/server/user/proto/service/user"
	"github.com/stretchr/testify/assert"
)

func TestAdminRequestValidation_ListAdmin(t *testing.T) {
	t.Parallel()
	type args struct {
		req *pb.ListAdminRequest
	}

	testCases := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "success",
			args: args{
				req: &pb.ListAdminRequest{
					Limit:  200,
					Offset: 100,
				},
			},
			want: true,
		},
		{
			name: "validation error: Search.Field.min_len",
			args: args{
				req: &pb.ListAdminRequest{
					Search: &pb.Search{
						Field: "",
						Value: "テストユーザー",
					},
					Order: &pb.Order{
						Field:   "created_at",
						OrderBy: pb.OrderBy_ORDER_BY_ASC,
					},
					Limit:  200,
					Offset: 100,
				},
			},
			want: false,
		},
		{
			name: "validation error: Search.Value.min_len",
			args: args{
				req: &pb.ListAdminRequest{
					Search: &pb.Search{
						Field: "username",
						Value: "",
					},
					Order: &pb.Order{
						Field:   "created_at",
						OrderBy: pb.OrderBy_ORDER_BY_ASC,
					},
					Limit:  200,
					Offset: 100,
				},
			},
			want: false,
		},
		{
			name: "validation error: Limit.lte",
			args: args{
				req: &pb.ListAdminRequest{
					Search: &pb.Search{
						Field: "username",
						Value: "テストユーザー",
					},
					Order: &pb.Order{
						Field:   "created_at",
						OrderBy: pb.OrderBy_ORDER_BY_ASC,
					},
					Limit:  201,
					Offset: 100,
				},
			},
			want: false,
		},
		{
			name: "validation error: Limit.gte",
			args: args{
				req: &pb.ListAdminRequest{
					Search: &pb.Search{
						Field: "username",
						Value: "テストユーザー",
					},
					Order: &pb.Order{
						Field:   "created_at",
						OrderBy: pb.OrderBy_ORDER_BY_ASC,
					},
					Limit:  200,
					Offset: -1,
				},
			},
			want: false,
		},
	}

	for _, tt := range testCases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			target := NewAdminRequestValidation()

			got := target.ListAdmin(tt.args.req)
			switch tt.want {
			case true:
				assert.NoError(t, got)
			case false:
				assert.Error(t, got)
			}
		})
	}
}

func TestAdminRequestValidation_GetAdmin(t *testing.T) {
	t.Parallel()
	type args struct {
		req *pb.GetAdminRequest
	}

	testCases := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "success",
			args: args{
				req: &pb.GetAdminRequest{
					UserId: "12345678-1234-1234-123456789012",
				},
			},
			want: true,
		},
		{
			name: "validation error: UserId.min_len",
			args: args{
				req: &pb.GetAdminRequest{
					UserId: "",
				},
			},
			want: false,
		},
	}

	for _, tt := range testCases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			target := NewAdminRequestValidation()

			got := target.GetAdmin(tt.args.req)
			switch tt.want {
			case true:
				assert.NoError(t, got)
			case false:
				assert.Error(t, got)
			}
		})
	}
}

func TestAdminRequestValidation_CreateAdmin(t *testing.T) {
	t.Parallel()
	type args struct {
		req *pb.CreateAdminRequest
	}

	testCases := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "success",
			args: args{
				req: &pb.CreateAdminRequest{
					Username:             "テストユーザー",
					Email:                "test-user@calmato.jp",
					PhoneNumber:          "000-0000-0000",
					Password:             "12345678",
					PasswordConfirmation: "12345678",
					Role:                 pb.Role_ROLE_DEVELOPER,
					LastName:             "テスト",
					FirstName:            "ユーザー",
					LastNameKana:         "てすと",
					FirstNameKana:        "ゆーざー",
				},
			},
			want: true,
		},
		{
			name: "validation error: Username.min_len",
			args: args{
				req: &pb.CreateAdminRequest{
					Username:             "",
					Email:                "test-user@calmato.jp",
					PhoneNumber:          "000-0000-0000",
					Password:             "12345678",
					PasswordConfirmation: "12345678",
					Role:                 pb.Role_ROLE_DEVELOPER,
					LastName:             "テスト",
					FirstName:            "ユーザー",
					LastNameKana:         "てすと",
					FirstNameKana:        "ゆーざー",
				},
			},
			want: false,
		},
		{
			name: "validation error: Username.max_len",
			args: args{
				req: &pb.CreateAdminRequest{
					Username:             strings.Repeat("x", 33),
					Email:                "test-user@calmato.jp",
					PhoneNumber:          "000-0000-0000",
					Password:             "12345678",
					PasswordConfirmation: "12345678",
					Role:                 pb.Role_ROLE_DEVELOPER,
					LastName:             "テスト",
					FirstName:            "ユーザー",
					LastNameKana:         "てすと",
					FirstNameKana:        "ゆーざー",
				},
			},
			want: false,
		},
		{
			name: "validation error: Email.min_len",
			args: args{
				req: &pb.CreateAdminRequest{
					Username:             "テストユーザー",
					Email:                "",
					PhoneNumber:          "000-0000-0000",
					Password:             "12345678",
					PasswordConfirmation: "12345678",
					Role:                 pb.Role_ROLE_DEVELOPER,
					LastName:             "テスト",
					FirstName:            "ユーザー",
					LastNameKana:         "てすと",
					FirstNameKana:        "ゆーざー",
				},
			},
			want: false,
		},
		{
			name: "validation error: Email.max_len",
			args: args{
				req: &pb.CreateAdminRequest{
					Username:             "テストユーザー",
					Email:                strings.Repeat("x", 246) + "@calmato.jp",
					PhoneNumber:          "000-0000-0000",
					Password:             "12345678",
					PasswordConfirmation: "12345678",
					Role:                 pb.Role_ROLE_DEVELOPER,
					LastName:             "テスト",
					FirstName:            "ユーザー",
					LastNameKana:         "てすと",
					FirstNameKana:        "ゆーざー",
				},
			},
			want: false,
		},
		{
			name: "validation error: Email.pattern",
			args: args{
				req: &pb.CreateAdminRequest{
					Username:             "テストユーザー",
					Email:                "test-user",
					PhoneNumber:          "000-0000-0000",
					Password:             "12345678",
					PasswordConfirmation: "12345678",
					Role:                 pb.Role_ROLE_DEVELOPER,
					LastName:             "テスト",
					FirstName:            "ユーザー",
					LastNameKana:         "てすと",
					FirstNameKana:        "ゆーざー",
				},
			},
			want: false,
		},
		{
			name: "validation error: PhoneNumber.min_len",
			args: args{
				req: &pb.CreateAdminRequest{
					Username:             "テストユーザー",
					Email:                "test-user@calmato.jp",
					PhoneNumber:          "",
					Password:             "12345678",
					PasswordConfirmation: "12345678",
					Role:                 pb.Role_ROLE_DEVELOPER,
					LastName:             "テスト",
					FirstName:            "ユーザー",
					LastNameKana:         "てすと",
					FirstNameKana:        "ゆーざー",
				},
			},
			want: false,
		},
		{
			name: "validation error: PhoneNumber.max_len",
			args: args{
				req: &pb.CreateAdminRequest{
					Username:             "テストユーザー",
					Email:                "test-user@calmato.jp",
					PhoneNumber:          strings.Repeat("x", 17),
					Password:             "12345678",
					PasswordConfirmation: "12345678",
					Role:                 pb.Role_ROLE_DEVELOPER,
					LastName:             "テスト",
					FirstName:            "ユーザー",
					LastNameKana:         "てすと",
					FirstNameKana:        "ゆーざー",
				},
			},
			want: false,
		},
		{
			name: "validation error: Password.min_len",
			args: args{
				req: &pb.CreateAdminRequest{
					Username:             "テストユーザー",
					Email:                "test-user@calmato.jp",
					PhoneNumber:          "000-0000-0000",
					Password:             "12345",
					PasswordConfirmation: "12345",
					Role:                 pb.Role_ROLE_DEVELOPER,
					LastName:             "テスト",
					FirstName:            "ユーザー",
					LastNameKana:         "てすと",
					FirstNameKana:        "ゆーざー",
				},
			},
			want: false,
		},
		{
			name: "validation error: Password.max_len",
			args: args{
				req: &pb.CreateAdminRequest{
					Username:             "テストユーザー",
					Email:                "test-user@calmato.jp",
					PhoneNumber:          "000-0000-0000",
					Password:             strings.Repeat("x", 33),
					PasswordConfirmation: strings.Repeat("x", 33),
					Role:                 pb.Role_ROLE_DEVELOPER,
					LastName:             "テスト",
					FirstName:            "ユーザー",
					LastNameKana:         "てすと",
					FirstNameKana:        "ゆーざー",
				},
			},
			want: false,
		},
		{
			name: "validation error: Password.max_len",
			args: args{
				req: &pb.CreateAdminRequest{
					Username:             "テストユーザー",
					Email:                "test-user@calmato.jp",
					PhoneNumber:          "000-0000-0000",
					Password:             "１２３４５６７８",
					PasswordConfirmation: "１２３４５６７８",
					Role:                 pb.Role_ROLE_DEVELOPER,
					LastName:             "テスト",
					FirstName:            "ユーザー",
					LastNameKana:         "てすと",
					FirstNameKana:        "ゆーざー",
				},
			},
			want: false,
		},
		{
			name: "validation error: PasswordConfirmation.min_len",
			args: args{
				req: &pb.CreateAdminRequest{
					Username:             "テストユーザー",
					Email:                "test-user@calmato.jp",
					PhoneNumber:          "000-0000-0000",
					Password:             "12345678",
					PasswordConfirmation: "",
					Role:                 pb.Role_ROLE_DEVELOPER,
					LastName:             "テスト",
					FirstName:            "ユーザー",
					LastNameKana:         "てすと",
					FirstNameKana:        "ゆーざー",
				},
			},
			want: false,
		},
		{
			name: "validation error: PasswordConfirmation.not_eq",
			args: args{
				req: &pb.CreateAdminRequest{
					Username:             "テストユーザー",
					Email:                "test-user@calmato.jp",
					PhoneNumber:          "000-0000-0000",
					Password:             "12345678",
					PasswordConfirmation: "123456789",
					Role:                 pb.Role_ROLE_DEVELOPER,
					LastName:             "テスト",
					FirstName:            "ユーザー",
					LastNameKana:         "てすと",
					FirstNameKana:        "ゆーざー",
				},
			},
			want: false,
		},
		{
			name: "validation error: Role.not_in",
			args: args{
				req: &pb.CreateAdminRequest{
					Username:             "テストユーザー",
					Email:                "test-user@calmato.jp",
					PhoneNumber:          "000-0000-0000",
					Password:             "12345678",
					PasswordConfirmation: "12345678",
					Role:                 pb.Role_ROLE_USER,
					LastName:             "テスト",
					FirstName:            "ユーザー",
					LastNameKana:         "てすと",
					FirstNameKana:        "ゆーざー",
				},
			},
			want: false,
		},
		{
			name: "validation error: Role.defined_only",
			args: args{
				req: &pb.CreateAdminRequest{
					Username:             "テストユーザー",
					Email:                "test-user@calmato.jp",
					PhoneNumber:          "000-0000-0000",
					Password:             "12345678",
					PasswordConfirmation: "12345678",
					Role:                 4,
					LastName:             "テスト",
					FirstName:            "ユーザー",
					LastNameKana:         "てすと",
					FirstNameKana:        "ゆーざー",
				},
			},
			want: false,
		},
		{
			name: "validation error: LastName.min_len",
			args: args{
				req: &pb.CreateAdminRequest{
					Username:             "テストユーザー",
					Email:                "test-user@calmato.jp",
					PhoneNumber:          "000-0000-0000",
					Password:             "12345678",
					PasswordConfirmation: "12345678",
					Role:                 pb.Role_ROLE_DEVELOPER,
					LastName:             "",
					FirstName:            "ユーザー",
					LastNameKana:         "てすと",
					FirstNameKana:        "ゆーざー",
				},
			},
			want: false,
		},
		{
			name: "validation error: LastName.max_len",
			args: args{
				req: &pb.CreateAdminRequest{
					Username:             "テストユーザー",
					Email:                "test-user@calmato.jp",
					PhoneNumber:          "000-0000-0000",
					Password:             "12345678",
					PasswordConfirmation: "12345678",
					Role:                 pb.Role_ROLE_DEVELOPER,
					LastName:             strings.Repeat("x", 17),
					FirstName:            "ユーザー",
					LastNameKana:         "てすと",
					FirstNameKana:        "ゆーざー",
				},
			},
			want: false,
		},
		{
			name: "validation error: FirstName.min_len",
			args: args{
				req: &pb.CreateAdminRequest{
					Username:             "テストユーザー",
					Email:                "test-user@calmato.jp",
					PhoneNumber:          "000-0000-0000",
					Password:             "12345678",
					PasswordConfirmation: "12345678",
					Role:                 pb.Role_ROLE_DEVELOPER,
					LastName:             "テスト",
					FirstName:            "",
					LastNameKana:         "てすと",
					FirstNameKana:        "ゆーざー",
				},
			},
			want: false,
		},
		{
			name: "validation error: FirstName.max_len",
			args: args{
				req: &pb.CreateAdminRequest{
					Username:             "テストユーザー",
					Email:                "test-user@calmato.jp",
					PhoneNumber:          "000-0000-0000",
					Password:             "12345678",
					PasswordConfirmation: "12345678",
					Role:                 pb.Role_ROLE_DEVELOPER,
					LastName:             "テスト",
					FirstName:            strings.Repeat("x", 17),
					LastNameKana:         "てすと",
					FirstNameKana:        "ゆーざー",
				},
			},
			want: false,
		},
		{
			name: "validation error: LastNameKana.min_len",
			args: args{
				req: &pb.CreateAdminRequest{
					Username:             "テストユーザー",
					Email:                "test-user@calmato.jp",
					PhoneNumber:          "000-0000-0000",
					Password:             "12345678",
					PasswordConfirmation: "12345678",
					Role:                 pb.Role_ROLE_DEVELOPER,
					LastName:             "テスト",
					FirstName:            "ユーザー",
					LastNameKana:         "",
					FirstNameKana:        "ゆーざー",
				},
			},
			want: false,
		},
		{
			name: "validation error: LastNameKana.max_len",
			args: args{
				req: &pb.CreateAdminRequest{
					Username:             "テストユーザー",
					Email:                "test-user@calmato.jp",
					PhoneNumber:          "000-0000-0000",
					Password:             "12345678",
					PasswordConfirmation: "12345678",
					Role:                 pb.Role_ROLE_DEVELOPER,
					LastName:             "テスト",
					FirstName:            "ユーザー",
					LastNameKana:         strings.Repeat("x", 33),
					FirstNameKana:        "ゆーざー",
				},
			},
			want: false,
		},
		{
			name: "validation error: LastNameKana.pattern",
			args: args{
				req: &pb.CreateAdminRequest{
					Username:             "テストユーザー",
					Email:                "test-user@calmato.jp",
					PhoneNumber:          "000-0000-0000",
					Password:             "12345678",
					PasswordConfirmation: "12345678",
					Role:                 pb.Role_ROLE_DEVELOPER,
					LastName:             "テスト",
					FirstName:            "ユーザー",
					LastNameKana:         "テスト",
					FirstNameKana:        "ゆーざー",
				},
			},
			want: false,
		},
		{
			name: "validation error: FirstNameKana.min_len",
			args: args{
				req: &pb.CreateAdminRequest{
					Username:             "テストユーザー",
					Email:                "test-user@calmato.jp",
					PhoneNumber:          "000-0000-0000",
					Password:             "12345678",
					PasswordConfirmation: "12345678",
					Role:                 pb.Role_ROLE_DEVELOPER,
					LastName:             "テスト",
					FirstName:            "ユーザー",
					LastNameKana:         "てすと",
					FirstNameKana:        "",
				},
			},
			want: false,
		},
		{
			name: "validation error: FirstNameKana.max_len",
			args: args{
				req: &pb.CreateAdminRequest{
					Username:             "テストユーザー",
					Email:                "test-user@calmato.jp",
					PhoneNumber:          "000-0000-0000",
					Password:             "12345678",
					PasswordConfirmation: "12345678",
					Role:                 pb.Role_ROLE_DEVELOPER,
					LastName:             "テスト",
					FirstName:            "ユーザー",
					LastNameKana:         "てすと",
					FirstNameKana:        strings.Repeat("x", 33),
				},
			},
			want: false,
		},
		{
			name: "validation error: FirstNameKana.pattern",
			args: args{
				req: &pb.CreateAdminRequest{
					Username:             "テストユーザー",
					Email:                "test-user@calmato.jp",
					PhoneNumber:          "000-0000-0000",
					Password:             "12345678",
					PasswordConfirmation: "12345678",
					Role:                 pb.Role_ROLE_DEVELOPER,
					LastName:             "テスト",
					FirstName:            "ユーザー",
					LastNameKana:         "てすと",
					FirstNameKana:        "ユーザー",
				},
			},
			want: false,
		},
	}

	for _, tt := range testCases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			target := NewAdminRequestValidation()

			got := target.CreateAdmin(tt.args.req)
			switch tt.want {
			case true:
				assert.NoError(t, got)
			case false:
				assert.Error(t, got)
			}
		})
	}
}

func TestAdminRequestValidation_UpdateAdminContact(t *testing.T) {
	t.Parallel()
	type args struct {
		req *pb.UpdateAdminContactRequest
	}

	testCases := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "success",
			args: args{
				req: &pb.UpdateAdminContactRequest{
					UserId:      "12345678-1234-1234-123456789012",
					Email:       "test-user@calmato.jp",
					PhoneNumber: "000-0000-0000",
				},
			},
			want: true,
		},
		{
			name: "validation error: UserId.min_len",
			args: args{
				req: &pb.UpdateAdminContactRequest{
					UserId:      "",
					Email:       "test-user@calmato.jp",
					PhoneNumber: "000-0000-0000",
				},
			},
			want: false,
		},
		{
			name: "validation error: Email.min_len",
			args: args{
				req: &pb.UpdateAdminContactRequest{
					UserId:      "12345678-1234-1234-123456789012",
					Email:       "",
					PhoneNumber: "000-0000-0000",
				},
			},
			want: false,
		},
		{
			name: "validation error: Email.max_len",
			args: args{
				req: &pb.UpdateAdminContactRequest{
					UserId:      "12345678-1234-1234-123456789012",
					Email:       strings.Repeat("x", 246) + "@calmato.jp",
					PhoneNumber: "000-0000-0000",
				},
			},
			want: false,
		},
		{
			name: "validation error: Email.pattern",
			args: args{
				req: &pb.UpdateAdminContactRequest{
					UserId:      "12345678-1234-1234-123456789012",
					Email:       "test-user",
					PhoneNumber: "000-0000-0000",
				},
			},
			want: false,
		},
		{
			name: "validation error: PhoneNumber.min_len",
			args: args{
				req: &pb.UpdateAdminContactRequest{
					UserId:      "12345678-1234-1234-123456789012",
					Email:       "test-user@calmato.jp",
					PhoneNumber: "",
				},
			},
			want: false,
		},
		{
			name: "validation error: PhoneNumber.max_len",
			args: args{
				req: &pb.UpdateAdminContactRequest{
					UserId:      "12345678-1234-1234-123456789012",
					Email:       "test-user@calmato.jp",
					PhoneNumber: strings.Repeat("x", 17),
				},
			},
			want: false,
		},
	}

	for _, tt := range testCases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			target := NewAdminRequestValidation()

			got := target.UpdateAdminContact(tt.args.req)
			switch tt.want {
			case true:
				assert.NoError(t, got)
			case false:
				assert.Error(t, got)
			}
		})
	}
}

func TestAdminRequestValidation_UpdateAdminPassword(t *testing.T) {
	t.Parallel()
	type args struct {
		req *pb.UpdateAdminPasswordRequest
	}

	testCases := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "success",
			args: args{
				req: &pb.UpdateAdminPasswordRequest{
					UserId:               "12345678-1234-1234-123456789012",
					Password:             "12345678",
					PasswordConfirmation: "12345678",
				},
			},
			want: true,
		},
		{
			name: "validation error: UserId.min_len",
			args: args{
				req: &pb.UpdateAdminPasswordRequest{
					UserId:               "",
					Password:             "12345678",
					PasswordConfirmation: "12345678",
				},
			},
			want: false,
		},
		{
			name: "validation error: Password.min_len",
			args: args{
				req: &pb.UpdateAdminPasswordRequest{
					UserId:               "12345678-1234-1234-123456789012",
					Password:             "12345",
					PasswordConfirmation: "12345",
				},
			},
			want: false,
		},
		{
			name: "validation error: Password.max_len",
			args: args{
				req: &pb.UpdateAdminPasswordRequest{
					UserId:               "12345678-1234-1234-123456789012",
					Password:             strings.Repeat("x", 33),
					PasswordConfirmation: strings.Repeat("x", 33),
				},
			},
			want: false,
		},
		{
			name: "validation error: Password.max_len",
			args: args{
				req: &pb.UpdateAdminPasswordRequest{
					UserId:               "12345678-1234-1234-123456789012",
					Password:             "１２３４５６７８",
					PasswordConfirmation: "１２３４５６７８",
				},
			},
			want: false,
		},
		{
			name: "validation error: PasswordConfirmation.min_len",
			args: args{
				req: &pb.UpdateAdminPasswordRequest{
					UserId:               "12345678-1234-1234-123456789012",
					Password:             "12345678",
					PasswordConfirmation: "",
				},
			},
			want: false,
		},
		{
			name: "validation error: PasswordConfirmation.not_eq",
			args: args{
				req: &pb.UpdateAdminPasswordRequest{
					UserId:               "12345678-1234-1234-123456789012",
					Password:             "12345678",
					PasswordConfirmation: "123456789",
				},
			},
			want: false,
		},
	}

	for _, tt := range testCases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			target := NewAdminRequestValidation()

			got := target.UpdateAdminPassword(tt.args.req)
			switch tt.want {
			case true:
				assert.NoError(t, got)
			case false:
				assert.Error(t, got)
			}
		})
	}
}

func TestAdminRequestValidation_UpdateAdminProfile(t *testing.T) {
	t.Parallel()
	type args struct {
		req *pb.UpdateAdminProfileRequest
	}

	testCases := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "success",
			args: args{
				req: &pb.UpdateAdminProfileRequest{
					UserId:        "12345678-1234-1234-123456789012",
					Username:      "テストユーザー",
					Role:          pb.Role_ROLE_DEVELOPER,
					LastName:      "テスト",
					FirstName:     "ユーザー",
					LastNameKana:  "てすと",
					FirstNameKana: "ゆーざー",
					ThumbnailUrl:  "https://go.dev/images/gophers/ladder.svg",
				},
			},
			want: true,
		},
		{
			name: "validation error: UserId.min_len",
			args: args{
				req: &pb.UpdateAdminProfileRequest{
					UserId:        "",
					Username:      "テストユーザー",
					Role:          pb.Role_ROLE_DEVELOPER,
					LastName:      "テスト",
					FirstName:     "ユーザー",
					LastNameKana:  "てすと",
					FirstNameKana: "ゆーざー",
					ThumbnailUrl:  "https://go.dev/images/gophers/ladder.svg",
				},
			},
			want: false,
		},
		{
			name: "validation error: Username.min_len",
			args: args{
				req: &pb.UpdateAdminProfileRequest{
					UserId:        "12345678-1234-1234-123456789012",
					Username:      "",
					Role:          pb.Role_ROLE_DEVELOPER,
					LastName:      "テスト",
					FirstName:     "ユーザー",
					LastNameKana:  "てすと",
					FirstNameKana: "ゆーざー",
					ThumbnailUrl:  "https://go.dev/images/gophers/ladder.svg",
				},
			},
			want: false,
		},
		{
			name: "validation error: Username.min_len",
			args: args{
				req: &pb.UpdateAdminProfileRequest{
					UserId:        "12345678-1234-1234-123456789012",
					Username:      strings.Repeat("x", 33),
					Role:          pb.Role_ROLE_DEVELOPER,
					LastName:      "テスト",
					FirstName:     "ユーザー",
					LastNameKana:  "てすと",
					FirstNameKana: "ゆーざー",
					ThumbnailUrl:  "https://go.dev/images/gophers/ladder.svg",
				},
			},
			want: false,
		},
		{
			name: "validation error: Role.not_in",
			args: args{
				req: &pb.UpdateAdminProfileRequest{
					UserId:        "12345678-1234-1234-123456789012",
					Username:      "テストユーザー",
					Role:          pb.Role_ROLE_USER,
					LastName:      "テスト",
					FirstName:     "ユーザー",
					LastNameKana:  "てすと",
					FirstNameKana: "ゆーざー",
					ThumbnailUrl:  "https://go.dev/images/gophers/ladder.svg",
				},
			},
			want: false,
		},
		{
			name: "validation error: Role.defined_only",
			args: args{
				req: &pb.UpdateAdminProfileRequest{
					UserId:        "12345678-1234-1234-123456789012",
					Username:      "テストユーザー",
					Role:          4,
					LastName:      "テスト",
					FirstName:     "ユーザー",
					LastNameKana:  "てすと",
					FirstNameKana: "ゆーざー",
					ThumbnailUrl:  "https://go.dev/images/gophers/ladder.svg",
				},
			},
			want: false,
		},
		{
			name: "validation error: LastName.min_len",
			args: args{
				req: &pb.UpdateAdminProfileRequest{
					UserId:        "12345678-1234-1234-123456789012",
					Username:      "テストユーザー",
					Role:          pb.Role_ROLE_DEVELOPER,
					LastName:      "",
					FirstName:     "ユーザー",
					LastNameKana:  "てすと",
					FirstNameKana: "ゆーざー",
					ThumbnailUrl:  "https://go.dev/images/gophers/ladder.svg",
				},
			},
			want: false,
		},
		{
			name: "validation error: LastName.max_len",
			args: args{
				req: &pb.UpdateAdminProfileRequest{
					UserId:        "12345678-1234-1234-123456789012",
					Username:      "テストユーザー",
					Role:          pb.Role_ROLE_DEVELOPER,
					LastName:      strings.Repeat("x", 17),
					FirstName:     "ユーザー",
					LastNameKana:  "てすと",
					FirstNameKana: "ゆーざー",
					ThumbnailUrl:  "https://go.dev/images/gophers/ladder.svg",
				},
			},
			want: false,
		},
		{
			name: "validation error: FirstName.min_len",
			args: args{
				req: &pb.UpdateAdminProfileRequest{
					UserId:        "12345678-1234-1234-123456789012",
					Username:      "テストユーザー",
					Role:          pb.Role_ROLE_DEVELOPER,
					LastName:      "テスト",
					FirstName:     "",
					LastNameKana:  "てすと",
					FirstNameKana: "ゆーざー",
					ThumbnailUrl:  "https://go.dev/images/gophers/ladder.svg",
				},
			},
			want: false,
		},
		{
			name: "validation error: FirstName.max_len",
			args: args{
				req: &pb.UpdateAdminProfileRequest{
					UserId:        "12345678-1234-1234-123456789012",
					Username:      "テストユーザー",
					Role:          pb.Role_ROLE_DEVELOPER,
					LastName:      "テスト",
					FirstName:     strings.Repeat("x", 17),
					LastNameKana:  "てすと",
					FirstNameKana: "ゆーざー",
					ThumbnailUrl:  "https://go.dev/images/gophers/ladder.svg",
				},
			},
			want: false,
		},
		{
			name: "validation error: LastNameKana.min_len",
			args: args{
				req: &pb.UpdateAdminProfileRequest{
					UserId:        "12345678-1234-1234-123456789012",
					Username:      "テストユーザー",
					Role:          pb.Role_ROLE_DEVELOPER,
					LastName:      "テスト",
					FirstName:     "ユーザー",
					LastNameKana:  "",
					FirstNameKana: "ゆーざー",
					ThumbnailUrl:  "https://go.dev/images/gophers/ladder.svg",
				},
			},
			want: false,
		},
		{
			name: "validation error: LastNameKana.max_len",
			args: args{
				req: &pb.UpdateAdminProfileRequest{
					UserId:        "12345678-1234-1234-123456789012",
					Username:      "テストユーザー",
					Role:          pb.Role_ROLE_DEVELOPER,
					LastName:      "テスト",
					FirstName:     "ユーザー",
					LastNameKana:  strings.Repeat("x", 33),
					FirstNameKana: "ゆーざー",
					ThumbnailUrl:  "https://go.dev/images/gophers/ladder.svg",
				},
			},
			want: false,
		},
		{
			name: "validation error: LastNameKana.pattern",
			args: args{
				req: &pb.UpdateAdminProfileRequest{
					UserId:        "12345678-1234-1234-123456789012",
					Username:      "テストユーザー",
					Role:          pb.Role_ROLE_DEVELOPER,
					LastName:      "テスト",
					FirstName:     "ユーザー",
					LastNameKana:  "テスト",
					FirstNameKana: "ゆーざー",
					ThumbnailUrl:  "https://go.dev/images/gophers/ladder.svg",
				},
			},
			want: false,
		},
		{
			name: "validation error: FirstNameKana.min_len",
			args: args{
				req: &pb.UpdateAdminProfileRequest{
					UserId:        "12345678-1234-1234-123456789012",
					Username:      "テストユーザー",
					Role:          pb.Role_ROLE_DEVELOPER,
					LastName:      "テスト",
					FirstName:     "ユーザー",
					LastNameKana:  "てすと",
					FirstNameKana: "",
					ThumbnailUrl:  "https://go.dev/images/gophers/ladder.svg",
				},
			},
			want: false,
		},
		{
			name: "validation error: FirstNameKana.max_len",
			args: args{
				req: &pb.UpdateAdminProfileRequest{
					UserId:        "12345678-1234-1234-123456789012",
					Username:      "テストユーザー",
					Role:          pb.Role_ROLE_DEVELOPER,
					LastName:      "テスト",
					FirstName:     "ユーザー",
					LastNameKana:  "てすと",
					FirstNameKana: strings.Repeat("x", 33),
					ThumbnailUrl:  "https://go.dev/images/gophers/ladder.svg",
				},
			},
			want: false,
		},
		{
			name: "validation error: FirstNameKana.pattern",
			args: args{
				req: &pb.UpdateAdminProfileRequest{
					UserId:        "12345678-1234-1234-123456789012",
					Username:      "テストユーザー",
					Role:          pb.Role_ROLE_DEVELOPER,
					LastName:      "テスト",
					FirstName:     "ユーザー",
					LastNameKana:  "てすと",
					FirstNameKana: "ユーザー",
					ThumbnailUrl:  "https://go.dev/images/gophers/ladder.svg",
				},
			},
			want: false,
		},
	}

	for _, tt := range testCases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			target := NewAdminRequestValidation()

			got := target.UpdateAdminProfile(tt.args.req)
			switch tt.want {
			case true:
				assert.NoError(t, got)
			case false:
				assert.Error(t, got)
			}
		})
	}
}

func TestAdminRequestValidation_UploadAdminThumbnail(t *testing.T) {
	t.Parallel()
	type args struct {
		req *pb.UploadAdminThumbnailRequest
	}

	testCases := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "success",
			args: args{
				req: &pb.UploadAdminThumbnailRequest{
					UserId:    "12345678-1234-1234-123456789012",
					Thumbnail: []byte{},
					Position:  0,
				},
			},
			want: true,
		},
		{
			name: "validation error: Position.gte",
			args: args{
				req: &pb.UploadAdminThumbnailRequest{
					UserId:    "12345678-1234-1234-123456789012",
					Thumbnail: []byte{},
					Position:  -1,
				},
			},
			want: false,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			target := NewAdminRequestValidation()

			got := target.UploadAdminThumbnail(tt.args.req)
			switch tt.want {
			case true:
				assert.NoError(t, got)
			case false:
				assert.Error(t, got)
			}
		})
	}
}

func TestAdminRequestValidation_DeleteAdmin(t *testing.T) {
	type args struct {
		req *pb.DeleteAdminRequest
	}

	testCases := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "success",
			args: args{
				req: &pb.DeleteAdminRequest{
					UserId: "12345678-1234-1234-123456789012",
				},
			},
			want: true,
		},
		{
			name: "validation error: UserId.min_len",
			args: args{
				req: &pb.DeleteAdminRequest{
					UserId: "",
				},
			},
			want: false,
		},
	}

	for _, tt := range testCases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			target := NewAdminRequestValidation()

			got := target.DeleteAdmin(tt.args.req)
			switch tt.want {
			case true:
				assert.NoError(t, got)
			case false:
				assert.Error(t, got)
			}
		})
	}
}
