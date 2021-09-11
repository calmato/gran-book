package validation

import (
	"strings"
	"testing"

	pb "github.com/calmato/gran-book/api/server/user/proto/user"
	"github.com/stretchr/testify/assert"
)

func TestAuthRequestValidation_CreateAuth(t *testing.T) {
	t.Parallel()
	type args struct {
		req *pb.CreateAuthRequest
	}

	testCases := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "success",
			args: args{
				req: &pb.CreateAuthRequest{
					Username:             "テストユーザー",
					Email:                "test-user@calmato.jp",
					Password:             "12345678",
					PasswordConfirmation: "12345678",
				},
			},
			want: true,
		},
		{
			name: "validation error: Username.min_len",
			args: args{
				req: &pb.CreateAuthRequest{
					Username:             "",
					Email:                "test-user@calmato.jp",
					Password:             "12345678",
					PasswordConfirmation: "12345678",
				},
			},
			want: false,
		},
		{
			name: "validation error: Username.max_len",
			args: args{
				req: &pb.CreateAuthRequest{
					Username:             strings.Repeat("x", 33),
					Email:                "test-user@calmato.jp",
					Password:             "12345678",
					PasswordConfirmation: "12345678",
				},
			},
			want: false,
		},
		{
			name: "validation error: Email.min_len",
			args: args{
				req: &pb.CreateAuthRequest{
					Username:             "テストユーザー",
					Email:                "",
					Password:             "12345678",
					PasswordConfirmation: "12345678",
				},
			},
			want: false,
		},
		{
			name: "validation error: Email.max_len",
			args: args{
				req: &pb.CreateAuthRequest{
					Username:             "テストユーザー",
					Email:                strings.Repeat("x", 246) + "@calmato.jp",
					Password:             "12345678",
					PasswordConfirmation: "12345678",
				},
			},
			want: false,
		},
		{
			name: "validation error: Email.pattern",
			args: args{
				req: &pb.CreateAuthRequest{
					Username:             "テストユーザー",
					Email:                "test-user",
					Password:             "12345678",
					PasswordConfirmation: "12345678",
				},
			},
			want: false,
		},
		{
			name: "validation error: Password.min_len",
			args: args{
				req: &pb.CreateAuthRequest{
					Username:             "テストユーザー",
					Email:                "test-user@calmato.jp",
					Password:             "12345",
					PasswordConfirmation: "12345",
				},
			},
			want: false,
		},
		{
			name: "validation error: Password.max_len",
			args: args{
				req: &pb.CreateAuthRequest{
					Username:             "テストユーザー",
					Email:                "test-user@calmato.jp",
					Password:             strings.Repeat("x", 33),
					PasswordConfirmation: strings.Repeat("x", 33),
				},
			},
			want: false,
		},
		{
			name: "validation error: Password.pattern",
			args: args{
				req: &pb.CreateAuthRequest{
					Username:             "テストユーザー",
					Email:                "test-user@calmato.jp",
					Password:             "１２３４５６７８",
					PasswordConfirmation: "１２３４５６７８",
				},
			},
			want: false,
		},
		{
			name: "validation error: PassworcConfirmation.min_len",
			args: args{
				req: &pb.CreateAuthRequest{
					Username:             "テストユーザー",
					Email:                "test-user@calmato.jp",
					Password:             "12345678",
					PasswordConfirmation: "",
				},
			},
			want: false,
		},
		{
			name: "validation error: PasswordConfirmation.not_eq",
			args: args{
				req: &pb.CreateAuthRequest{
					Username:             "テストユーザー",
					Email:                "test-user@calmato.jp",
					Password:             "12345678",
					PasswordConfirmation: "123456789",
				},
			},
			want: false,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			target := NewAuthRequestValidation()

			got := target.CreateAuth(tt.args.req)
			switch tt.want {
			case true:
				assert.NoError(t, got)
			case false:
				assert.Error(t, got)
			}
		})
	}
}

func TestAuthRequestValidation_UpdateAuthEmail(t *testing.T) {
	t.Parallel()
	type args struct {
		req *pb.UpdateAuthEmailRequest
	}

	testCases := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "success",
			args: args{
				req: &pb.UpdateAuthEmailRequest{
					Email: "test-user@calmato.jp",
				},
			},
			want: true,
		},
		{
			name: "validation error: Email.min_len",
			args: args{
				req: &pb.UpdateAuthEmailRequest{
					Email: "",
				},
			},
			want: false,
		},
		{
			name: "validation error: Email.max_len",
			args: args{
				req: &pb.UpdateAuthEmailRequest{
					Email: strings.Repeat("x", 246) + "@calmato.jp",
				},
			},
			want: false,
		},
		{
			name: "validation error: Email.pattern",
			args: args{
				req: &pb.UpdateAuthEmailRequest{
					Email: "test-user",
				},
			},
			want: false,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			target := NewAuthRequestValidation()

			got := target.UpdateAuthEmail(tt.args.req)
			switch tt.want {
			case true:
				assert.NoError(t, got)
			case false:
				assert.Error(t, got)
			}
		})
	}
}

func TestAuthRequestValidation_UpdateAuthPassword(t *testing.T) {
	t.Parallel()
	type args struct {
		req *pb.UpdateAuthPasswordRequest
	}

	testCases := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "success",
			args: args{
				req: &pb.UpdateAuthPasswordRequest{
					Password:             "12345678",
					PasswordConfirmation: "12345678",
				},
			},
			want: true,
		},
		{
			name: "validation error: Password.min_len",
			args: args{
				req: &pb.UpdateAuthPasswordRequest{
					Password:             "12345",
					PasswordConfirmation: "12345",
				},
			},
			want: false,
		},
		{
			name: "validation error: Password.max_len",
			args: args{
				req: &pb.UpdateAuthPasswordRequest{
					Password:             strings.Repeat("x", 33),
					PasswordConfirmation: strings.Repeat("x", 33),
				},
			},
			want: false,
		},
		{
			name: "validation error: Password.pattern",
			args: args{
				req: &pb.UpdateAuthPasswordRequest{
					Password:             "１２３４５６７８",
					PasswordConfirmation: "１２３４５６７８",
				},
			},
			want: false,
		},
		{
			name: "validation error: PassworcConfirmation.min_len",
			args: args{
				req: &pb.UpdateAuthPasswordRequest{
					Password:             "12345678",
					PasswordConfirmation: "",
				},
			},
			want: false,
		},
		{
			name: "validation error: PasswordConfirmation.not_eq",
			args: args{
				req: &pb.UpdateAuthPasswordRequest{
					Password:             "12345678",
					PasswordConfirmation: "123456789",
				},
			},
			want: false,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			target := NewAuthRequestValidation()

			got := target.UpdateAuthPassword(tt.args.req)
			switch tt.want {
			case true:
				assert.NoError(t, got)
			case false:
				assert.Error(t, got)
			}
		})
	}
}

func TestAuthRequestValidation_UpdateAuthProfile(t *testing.T) {
	t.Parallel()
	type args struct {
		req *pb.UpdateAuthProfileRequest
	}

	testCases := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "success",
			args: args{
				req: &pb.UpdateAuthProfileRequest{
					Username:         "test-user",
					Gender:           pb.Gender_GENDER_MAN,
					ThumbnailUrl:     "https://go.dev/images/gophers/ladder.svg",
					SelfIntroduction: "自己紹介",
				},
			},
			want: true,
		},
		{
			name: "validation error: Username.min_len",
			args: args{
				req: &pb.UpdateAuthProfileRequest{
					Username:         "",
					Gender:           pb.Gender_GENDER_MAN,
					ThumbnailUrl:     "https://go.dev/images/gophers/ladder.svg",
					SelfIntroduction: "自己紹介",
				},
			},
			want: false,
		},
		{
			name: "validation error: Username.max_len",
			args: args{
				req: &pb.UpdateAuthProfileRequest{
					Username:         strings.Repeat("x", 33),
					Gender:           pb.Gender_GENDER_MAN,
					ThumbnailUrl:     "https://go.dev/images/gophers/ladder.svg",
					SelfIntroduction: "自己紹介",
				},
			},
			want: false,
		},
		{
			name: "validation error: Gender.defined_only",
			args: args{
				req: &pb.UpdateAuthProfileRequest{
					Username:         "test-user",
					Gender:           4,
					ThumbnailUrl:     "https://go.dev/images/gophers/ladder.svg",
					SelfIntroduction: "自己紹介",
				},
			},
			want: false,
		},
		{
			name: "validation error: SelfIntroduction.max_len",
			args: args{
				req: &pb.UpdateAuthProfileRequest{
					Username:         "test-user",
					Gender:           4,
					ThumbnailUrl:     "https://go.dev/images/gophers/ladder.svg",
					SelfIntroduction: strings.Repeat("x", 257),
				},
			},
			want: false,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			target := NewAuthRequestValidation()

			got := target.UpdateAuthProfile(tt.args.req)
			switch tt.want {
			case true:
				assert.NoError(t, got)
			case false:
				assert.Error(t, got)
			}
		})
	}
}

func TestAuthRequestValidation_UpdateAuthAddress(t *testing.T) {
	type args struct {
		req *pb.UpdateAuthAddressRequest
	}

	testCases := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "success",
			args: args{
				req: &pb.UpdateAuthAddressRequest{
					LastName:      "テスト",
					FirstName:     "ユーザ",
					LastNameKana:  "てすと",
					FirstNameKana: "ゆーざ",
					PhoneNumber:   "000-0000-0000",
					PostalCode:    "000-0000",
					Prefecture:    "東京都",
					City:          "小金井市",
					AddressLine1:  "貫井北町4-1-1",
					AddressLine2:  "",
				},
			},
			want: true,
		},
		{
			name: "validation error: LastName.min_len",
			args: args{
				req: &pb.UpdateAuthAddressRequest{
					LastName:      "",
					FirstName:     "ユーザ",
					LastNameKana:  "てすと",
					FirstNameKana: "ゆーざ",
					PhoneNumber:   "000-0000-0000",
					PostalCode:    "000-0000",
					Prefecture:    "東京都",
					City:          "小金井市",
					AddressLine1:  "貫井北町4-1-1",
					AddressLine2:  "",
				},
			},
			want: false,
		},
		{
			name: "validation error: LastName.max_len",
			args: args{
				req: &pb.UpdateAuthAddressRequest{
					LastName:      strings.Repeat("x", 17),
					FirstName:     "ユーザ",
					LastNameKana:  "てすと",
					FirstNameKana: "ゆーざ",
					PhoneNumber:   "000-0000-0000",
					PostalCode:    "000-0000",
					Prefecture:    "東京都",
					City:          "小金井市",
					AddressLine1:  "貫井北町4-1-1",
					AddressLine2:  "",
				},
			},
			want: false,
		},
		{
			name: "validation error: FirstName.min_len",
			args: args{
				req: &pb.UpdateAuthAddressRequest{
					LastName:      "テスト",
					FirstName:     "",
					LastNameKana:  "てすと",
					FirstNameKana: "ゆーざ",
					PhoneNumber:   "000-0000-0000",
					PostalCode:    "000-0000",
					Prefecture:    "東京都",
					City:          "小金井市",
					AddressLine1:  "貫井北町4-1-1",
					AddressLine2:  "",
				},
			},
			want: false,
		},
		{
			name: "validation error: FirstName.max_len",
			args: args{
				req: &pb.UpdateAuthAddressRequest{
					LastName:      "テスト",
					FirstName:     strings.Repeat("x", 17),
					LastNameKana:  "てすと",
					FirstNameKana: "ゆーざ",
					PhoneNumber:   "000-0000-0000",
					PostalCode:    "000-0000",
					Prefecture:    "東京都",
					City:          "小金井市",
					AddressLine1:  "貫井北町4-1-1",
					AddressLine2:  "",
				},
			},
			want: false,
		},
		{
			name: "validation error: LastNameKana.min_len",
			args: args{
				req: &pb.UpdateAuthAddressRequest{
					LastName:      "テスト",
					FirstName:     "ユーザ",
					LastNameKana:  "",
					FirstNameKana: "ゆーざ",
					PhoneNumber:   "000-0000-0000",
					PostalCode:    "000-0000",
					Prefecture:    "東京都",
					City:          "小金井市",
					AddressLine1:  "貫井北町4-1-1",
					AddressLine2:  "",
				},
			},
			want: false,
		},
		{
			name: "validation error: LastNameKana.max_len",
			args: args{
				req: &pb.UpdateAuthAddressRequest{
					LastName:      "テスト",
					FirstName:     "ユーザ",
					LastNameKana:  strings.Repeat("x", 33),
					FirstNameKana: "ゆーざ",
					PhoneNumber:   "000-0000-0000",
					PostalCode:    "000-0000",
					Prefecture:    "東京都",
					City:          "小金井市",
					AddressLine1:  "貫井北町4-1-1",
					AddressLine2:  "",
				},
			},
			want: false,
		},
		{
			name: "validation error: LastNameKana.pattern",
			args: args{
				req: &pb.UpdateAuthAddressRequest{
					LastName:      "テスト",
					FirstName:     "ユーザ",
					LastNameKana:  "テスト",
					FirstNameKana: "ゆーざ",
					PhoneNumber:   "000-0000-0000",
					PostalCode:    "000-0000",
					Prefecture:    "東京都",
					City:          "小金井市",
					AddressLine1:  "貫井北町4-1-1",
					AddressLine2:  "",
				},
			},
			want: false,
		},
		{
			name: "validation error: FirstNameKana.min_len",
			args: args{
				req: &pb.UpdateAuthAddressRequest{
					LastName:      "テスト",
					FirstName:     "ユーザ",
					LastNameKana:  "てすと",
					FirstNameKana: "",
					PhoneNumber:   "000-0000-0000",
					PostalCode:    "000-0000",
					Prefecture:    "東京都",
					City:          "小金井市",
					AddressLine1:  "貫井北町4-1-1",
					AddressLine2:  "",
				},
			},
			want: false,
		},
		{
			name: "validation error: FirstNameKana.max_len",
			args: args{
				req: &pb.UpdateAuthAddressRequest{
					LastName:      "テスト",
					FirstName:     "ユーザ",
					LastNameKana:  "てすと",
					FirstNameKana: strings.Repeat("x", 33),
					PhoneNumber:   "000-0000-0000",
					PostalCode:    "000-0000",
					Prefecture:    "東京都",
					City:          "小金井市",
					AddressLine1:  "貫井北町4-1-1",
					AddressLine2:  "",
				},
			},
			want: false,
		},
		{
			name: "validation error: FirstNameKana.pattern",
			args: args{
				req: &pb.UpdateAuthAddressRequest{
					LastName:      "テスト",
					FirstName:     "ユーザ",
					LastNameKana:  "てすと",
					FirstNameKana: "ユーザー",
					PhoneNumber:   "000-0000-0000",
					PostalCode:    "000-0000",
					Prefecture:    "東京都",
					City:          "小金井市",
					AddressLine1:  "貫井北町4-1-1",
					AddressLine2:  "",
				},
			},
			want: false,
		},
		{
			name: "validation error: PhoneNumber.min_len",
			args: args{
				req: &pb.UpdateAuthAddressRequest{
					LastName:      "テスト",
					FirstName:     "ユーザ",
					LastNameKana:  "てすと",
					FirstNameKana: "ゆーざー",
					PhoneNumber:   "",
					PostalCode:    "000-0000",
					Prefecture:    "東京都",
					City:          "小金井市",
					AddressLine1:  "貫井北町4-1-1",
					AddressLine2:  "",
				},
			},
			want: false,
		},
		{
			name: "validation error: PhoneNumber.max_len",
			args: args{
				req: &pb.UpdateAuthAddressRequest{
					LastName:      "テスト",
					FirstName:     "ユーザ",
					LastNameKana:  "てすと",
					FirstNameKana: "ゆーざー",
					PhoneNumber:   strings.Repeat("x", 17),
					PostalCode:    "000-0000",
					Prefecture:    "東京都",
					City:          "小金井市",
					AddressLine1:  "貫井北町4-1-1",
					AddressLine2:  "",
				},
			},
			want: false,
		},
		{
			name: "validation error: PostalCode.min_len",
			args: args{
				req: &pb.UpdateAuthAddressRequest{
					LastName:      "テスト",
					FirstName:     "ユーザ",
					LastNameKana:  "てすと",
					FirstNameKana: "ゆーざー",
					PhoneNumber:   "000-0000-0000",
					PostalCode:    "",
					Prefecture:    "東京都",
					City:          "小金井市",
					AddressLine1:  "貫井北町4-1-1",
					AddressLine2:  "",
				},
			},
			want: false,
		},
		{
			name: "validation error: PostalCode.max_len",
			args: args{
				req: &pb.UpdateAuthAddressRequest{
					LastName:      "テスト",
					FirstName:     "ユーザ",
					LastNameKana:  "てすと",
					FirstNameKana: "ゆーざー",
					PhoneNumber:   "000-0000-0000",
					PostalCode:    strings.Repeat("0", 17),
					Prefecture:    "東京都",
					City:          "小金井市",
					AddressLine1:  "貫井北町4-1-1",
					AddressLine2:  "",
				},
			},
			want: false,
		},
		{
			name: "validation error: Prefecture.min_len",
			args: args{
				req: &pb.UpdateAuthAddressRequest{
					LastName:      "テスト",
					FirstName:     "ユーザ",
					LastNameKana:  "てすと",
					FirstNameKana: "ゆーざー",
					PhoneNumber:   "000-0000-0000",
					PostalCode:    "000-0000",
					Prefecture:    "",
					City:          "小金井市",
					AddressLine1:  "貫井北町4-1-1",
					AddressLine2:  "",
				},
			},
			want: false,
		},
		{
			name: "validation error: Prefecture.max_len",
			args: args{
				req: &pb.UpdateAuthAddressRequest{
					LastName:      "テスト",
					FirstName:     "ユーザ",
					LastNameKana:  "てすと",
					FirstNameKana: "ゆーざー",
					PhoneNumber:   "000-0000-0000",
					PostalCode:    "000-0000",
					Prefecture:    strings.Repeat("x", 33),
					City:          "小金井市",
					AddressLine1:  "貫井北町4-1-1",
					AddressLine2:  "",
				},
			},
			want: false,
		},
		{
			name: "validation error: City.min_len",
			args: args{
				req: &pb.UpdateAuthAddressRequest{
					LastName:      "テスト",
					FirstName:     "ユーザ",
					LastNameKana:  "てすと",
					FirstNameKana: "ゆーざー",
					PhoneNumber:   "000-0000-0000",
					PostalCode:    "000-0000",
					Prefecture:    "東京都",
					City:          "",
					AddressLine1:  "貫井北町4-1-1",
					AddressLine2:  "",
				},
			},
			want: false,
		},
		{
			name: "validation error: City.max_len",
			args: args{
				req: &pb.UpdateAuthAddressRequest{
					LastName:      "テスト",
					FirstName:     "ユーザ",
					LastNameKana:  "てすと",
					FirstNameKana: "ゆーざー",
					PhoneNumber:   "000-0000-0000",
					PostalCode:    "000-0000",
					Prefecture:    "東京都",
					City:          strings.Repeat("x", 33),
					AddressLine1:  "貫井北町4-1-1",
					AddressLine2:  "",
				},
			},
			want: false,
		},
		{
			name: "validation error: AddressLine1.min_len",
			args: args{
				req: &pb.UpdateAuthAddressRequest{
					LastName:      "テスト",
					FirstName:     "ユーザ",
					LastNameKana:  "てすと",
					FirstNameKana: "ゆーざー",
					PhoneNumber:   "000-0000-0000",
					PostalCode:    "000-0000",
					Prefecture:    "東京都",
					City:          "小金井市",
					AddressLine1:  "",
					AddressLine2:  "",
				},
			},
			want: false,
		},
		{
			name: "validation error: AddressLine1.max_len",
			args: args{
				req: &pb.UpdateAuthAddressRequest{
					LastName:      "テスト",
					FirstName:     "ユーザ",
					LastNameKana:  "てすと",
					FirstNameKana: "ゆーざー",
					PhoneNumber:   "000-0000-0000",
					PostalCode:    "000-0000",
					Prefecture:    "東京都",
					City:          "小金井市",
					AddressLine1:  strings.Repeat("x", 65),
					AddressLine2:  "",
				},
			},
			want: false,
		},
		{
			name: "validation error: AddressLine2.max_len",
			args: args{
				req: &pb.UpdateAuthAddressRequest{
					LastName:      "テスト",
					FirstName:     "ユーザ",
					LastNameKana:  "てすと",
					FirstNameKana: "ゆーざー",
					PhoneNumber:   "000-0000-0000",
					PostalCode:    "000-0000",
					Prefecture:    "東京都",
					City:          "小金井市",
					AddressLine1:  "貫井北町4-1-1",
					AddressLine2:  strings.Repeat("x", 65),
				},
			},
			want: false,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			target := NewAuthRequestValidation()

			got := target.UpdateAuthAddress(tt.args.req)
			switch tt.want {
			case true:
				assert.NoError(t, got)
			case false:
				assert.Error(t, got)
			}
		})
	}
}

func TestAuthRequestValidation_UploadAuthThumbnail(t *testing.T) {
	type args struct {
		req *pb.UploadAuthThumbnailRequest
	}

	testCases := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "success",
			args: args{
				req: &pb.UploadAuthThumbnailRequest{
					Thumbnail: []byte{},
					Position:  0,
				},
			},
			want: true,
		},
		{
			name: "validation error: Position.lte",
			args: args{
				req: &pb.UploadAuthThumbnailRequest{
					Thumbnail: []byte{},
					Position:  -1,
				},
			},
			want: false,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			target := NewAuthRequestValidation()

			got := target.UploadAuthThumbnail(tt.args.req)
			switch tt.want {
			case true:
				assert.NoError(t, got)
			case false:
				assert.Error(t, got)
			}
		})
	}
}

func TestAuthRequestValidation_RegisterAuthDevice(t *testing.T) {
	type args struct {
		req *pb.RegisterAuthDeviceRequest
	}

	testCases := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "success",
			args: args{
				req: &pb.RegisterAuthDeviceRequest{
					InstanceId: "ExponentPushToken[!Qaz2wsx3edc4rfv5tgb6y]",
				},
			},
			want: true,
		},
		{
			name: "validation error: InstanceId.min_len",
			args: args{
				req: &pb.RegisterAuthDeviceRequest{
					InstanceId: "",
				},
			},
			want: false,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			target := NewAuthRequestValidation()

			got := target.RegisterAuthDevice(tt.args.req)
			switch tt.want {
			case true:
				assert.NoError(t, got)
			case false:
				assert.Error(t, got)
			}
		})
	}
}
