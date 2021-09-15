package v1

import (
	"context"
	"net/http"
	"testing"

	"github.com/calmato/gran-book/api/gateway/native/internal/entity"
	request "github.com/calmato/gran-book/api/gateway/native/internal/request/v1"
	response "github.com/calmato/gran-book/api/gateway/native/internal/response/v1"
	"github.com/calmato/gran-book/api/gateway/native/pkg/test"
	"github.com/calmato/gran-book/api/gateway/native/proto/service/user"
	"github.com/golang/mock/gomock"
)

func TestAuth_Get(t *testing.T) {
	t.Parallel()

	user1 := testAuth("00000000-0000-0000-0000-000000000000")

	tests := []struct {
		name   string
		setup  func(ctx context.Context, t *testing.T, mocks *test.Mocks)
		expect *test.TestResponse
	}{
		{
			name: "success",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.AuthService.EXPECT().
					GetAuth(gomock.Any(), &user.Empty{}).
					Return(&user.AuthResponse{Auth: user1}, nil)
			},
			expect: &test.TestResponse{
				Code: http.StatusOK,
				Body: &response.AuthResponse{
					ID:               "00000000-0000-0000-0000-000000000000",
					Username:         "テストユーザー",
					Gender:           entity.GenderMan,
					Email:            "test-user@calmato.jp",
					PhoneNumber:      "000-0000-0000",
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
					CreatedAt:        test.TimeMock,
					UpdatedAt:        test.TimeMock,
				},
			},
		},
		{
			name: "failed to get auth",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.AuthService.EXPECT().
					GetAuth(gomock.Any(), &user.Empty{}).
					Return(nil, test.ErrMock)
			},
			expect: &test.TestResponse{
				Code: http.StatusInternalServerError,
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			ctx, res, mocks := test.NewHTTPClient(t, nil)
			tt.setup(ctx, t, mocks)

			handler := NewAuthHandler(mocks.AuthService)
			handler.Get(ctx)

			test.TestHTTP(t, tt.expect, res)
		})
	}
}

func TestAuth_Create(t *testing.T) {
	t.Parallel()

	user1 := testAuth("00000000-0000-0000-0000-000000000000")

	tests := []struct {
		name   string
		setup  func(ctx context.Context, t *testing.T, mocks *test.Mocks)
		req    *request.CreateAuthRequest
		expect *test.TestResponse
	}{
		{
			name: "success",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.AuthService.EXPECT().
					CreateAuth(gomock.Any(), &user.CreateAuthRequest{
						Username:             "テストユーザー",
						Email:                "test-user@calmato.jp",
						Password:             "12345678",
						PasswordConfirmation: "12345678",
					}).
					Return(&user.AuthResponse{Auth: user1}, nil)
			},
			req: &request.CreateAuthRequest{
				Username:             "テストユーザー",
				Email:                "test-user@calmato.jp",
				Password:             "12345678",
				PasswordConfirmation: "12345678",
			},
			expect: &test.TestResponse{
				Code: http.StatusOK,
				Body: &response.AuthResponse{
					ID:               "00000000-0000-0000-0000-000000000000",
					Username:         "テストユーザー",
					Gender:           entity.GenderMan,
					Email:            "test-user@calmato.jp",
					PhoneNumber:      "000-0000-0000",
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
					CreatedAt:        test.TimeMock,
					UpdatedAt:        test.TimeMock,
				},
			},
		},
		{
			name:  "failed to invalid argument",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {},
			req:   &request.CreateAuthRequest{},
			expect: &test.TestResponse{
				Code: http.StatusBadRequest,
			},
		},
		{
			name: "failed to create auth",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.AuthService.EXPECT().
					CreateAuth(gomock.Any(), &user.CreateAuthRequest{
						Username:             "テストユーザー",
						Email:                "test-user@calmato.jp",
						Password:             "12345678",
						PasswordConfirmation: "12345678",
					}).
					Return(nil, test.ErrMock)
			},
			req: &request.CreateAuthRequest{
				Username:             "テストユーザー",
				Email:                "test-user@calmato.jp",
				Password:             "12345678",
				PasswordConfirmation: "12345678",
			},
			expect: &test.TestResponse{
				Code: http.StatusInternalServerError,
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			ctx, res, mocks := test.NewHTTPClient(t, tt.req)
			tt.setup(ctx, t, mocks)

			handler := NewAuthHandler(mocks.AuthService)
			handler.Create(ctx)

			test.TestHTTP(t, tt.expect, res)
		})
	}
}

func TestAuth_UpdateProfile(t *testing.T) {
	t.Parallel()

	user1 := testAuth("00000000-0000-0000-0000-000000000000")

	tests := []struct {
		name   string
		setup  func(ctx context.Context, t *testing.T, mocks *test.Mocks)
		req    *request.UpdateAuthProfileRequest
		expect *test.TestResponse
	}{
		{
			name: "success",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.AuthService.EXPECT().
					UpdateAuthProfile(gomock.Any(), &user.UpdateAuthProfileRequest{
						Username:         "テストユーザー",
						Gender:           user.Gender_GENDER_MAN,
						ThumbnailUrl:     "https://go.dev/images/gophers/ladder.svg",
						SelfIntroduction: "テストコードです。",
					}).
					Return(&user.AuthResponse{Auth: user1}, nil)
			},
			req: &request.UpdateAuthProfileRequest{
				Username:         "テストユーザー",
				Gender:           entity.GenderMan,
				ThumbnailURL:     "https://go.dev/images/gophers/ladder.svg",
				SelfIntroduction: "テストコードです。",
			},
			expect: &test.TestResponse{
				Code: http.StatusOK,
				Body: &response.AuthResponse{
					ID:               "00000000-0000-0000-0000-000000000000",
					Username:         "テストユーザー",
					Gender:           entity.GenderMan,
					Email:            "test-user@calmato.jp",
					PhoneNumber:      "000-0000-0000",
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
					CreatedAt:        test.TimeMock,
					UpdatedAt:        test.TimeMock,
				},
			},
		},
		{
			name:  "failed to invalid argument",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {},
			req:   &request.UpdateAuthProfileRequest{},
			expect: &test.TestResponse{
				Code: http.StatusBadRequest,
			},
		},
		{
			name: "failed to update auth profile",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.AuthService.EXPECT().
					UpdateAuthProfile(gomock.Any(), &user.UpdateAuthProfileRequest{
						Username:         "テストユーザー",
						Gender:           user.Gender_GENDER_MAN,
						ThumbnailUrl:     "https://go.dev/images/gophers/ladder.svg",
						SelfIntroduction: "テストコードです。",
					}).
					Return(nil, test.ErrMock)
			},
			req: &request.UpdateAuthProfileRequest{
				Username:         "テストユーザー",
				Gender:           entity.GenderMan,
				ThumbnailURL:     "https://go.dev/images/gophers/ladder.svg",
				SelfIntroduction: "テストコードです。",
			},
			expect: &test.TestResponse{
				Code: http.StatusInternalServerError,
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			ctx, res, mocks := test.NewHTTPClient(t, tt.req)
			tt.setup(ctx, t, mocks)

			handler := NewAuthHandler(mocks.AuthService)
			handler.UpdateProfile(ctx)

			test.TestHTTP(t, tt.expect, res)
		})
	}
}

func TestAuth_UpdateAddress(t *testing.T) {
	t.Parallel()

	user1 := testAuth("00000000-0000-0000-0000-000000000000")

	tests := []struct {
		name   string
		setup  func(ctx context.Context, t *testing.T, mocks *test.Mocks)
		req    *request.UpdateAuthAddressRequest
		expect *test.TestResponse
	}{
		{
			name: "success",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.AuthService.EXPECT().
					UpdateAuthAddress(gomock.Any(), &user.UpdateAuthAddressRequest{
						LastName:      "テスト",
						FirstName:     "ユーザー",
						LastNameKana:  "てすと",
						FirstNameKana: "ゆーざー",
						PhoneNumber:   "000-0000-0000",
						PostalCode:    "000-0000",
						Prefecture:    "東京都",
						City:          "小金井市",
						AddressLine1:  "貫井北町4-1-1",
						AddressLine2:  "",
					}).
					Return(&user.AuthResponse{Auth: user1}, nil)
			},
			req: &request.UpdateAuthAddressRequest{
				LastName:      "テスト",
				FirstName:     "ユーザー",
				LastNameKana:  "てすと",
				FirstNameKana: "ゆーざー",
				PhoneNumber:   "000-0000-0000",
				PostalCode:    "000-0000",
				Prefecture:    "東京都",
				City:          "小金井市",
				AddressLine1:  "貫井北町4-1-1",
				AddressLine2:  "",
			},
			expect: &test.TestResponse{
				Code: http.StatusOK,
				Body: &response.AuthResponse{
					ID:               "00000000-0000-0000-0000-000000000000",
					Username:         "テストユーザー",
					Gender:           entity.GenderMan,
					Email:            "test-user@calmato.jp",
					PhoneNumber:      "000-0000-0000",
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
					CreatedAt:        test.TimeMock,
					UpdatedAt:        test.TimeMock,
				},
			},
		},
		{
			name:  "failed to invalid argument",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {},
			req:   &request.UpdateAuthAddressRequest{},
			expect: &test.TestResponse{
				Code: http.StatusBadRequest,
			},
		},
		{
			name: "failed to update auth address",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.AuthService.EXPECT().
					UpdateAuthAddress(gomock.Any(), &user.UpdateAuthAddressRequest{
						LastName:      "テスト",
						FirstName:     "ユーザー",
						LastNameKana:  "てすと",
						FirstNameKana: "ゆーざー",
						PhoneNumber:   "000-0000-0000",
						PostalCode:    "000-0000",
						Prefecture:    "東京都",
						City:          "小金井市",
						AddressLine1:  "貫井北町4-1-1",
						AddressLine2:  "",
					}).
					Return(nil, test.ErrMock)
			},
			req: &request.UpdateAuthAddressRequest{
				LastName:      "テスト",
				FirstName:     "ユーザー",
				LastNameKana:  "てすと",
				FirstNameKana: "ゆーざー",
				PhoneNumber:   "000-0000-0000",
				PostalCode:    "000-0000",
				Prefecture:    "東京都",
				City:          "小金井市",
				AddressLine1:  "貫井北町4-1-1",
				AddressLine2:  "",
			},
			expect: &test.TestResponse{
				Code: http.StatusInternalServerError,
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			ctx, res, mocks := test.NewHTTPClient(t, tt.req)
			tt.setup(ctx, t, mocks)

			handler := NewAuthHandler(mocks.AuthService)
			handler.UpdateAddress(ctx)

			test.TestHTTP(t, tt.expect, res)
		})
	}
}

func TestAuth_UpdateEmail(t *testing.T) {
	t.Parallel()

	user1 := testAuth("00000000-0000-0000-0000-000000000000")

	tests := []struct {
		name   string
		setup  func(ctx context.Context, t *testing.T, mocks *test.Mocks)
		req    *request.UpdateAuthEmailRequest
		expect *test.TestResponse
	}{
		{
			name: "success",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.AuthService.EXPECT().
					UpdateAuthEmail(gomock.Any(), &user.UpdateAuthEmailRequest{
						Email: "test-user@calmato.jp",
					}).
					Return(&user.AuthResponse{Auth: user1}, nil)
			},
			req: &request.UpdateAuthEmailRequest{
				Email: "test-user@calmato.jp",
			},
			expect: &test.TestResponse{
				Code: http.StatusOK,
				Body: &response.AuthResponse{
					ID:               "00000000-0000-0000-0000-000000000000",
					Username:         "テストユーザー",
					Gender:           entity.GenderMan,
					Email:            "test-user@calmato.jp",
					PhoneNumber:      "000-0000-0000",
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
					CreatedAt:        test.TimeMock,
					UpdatedAt:        test.TimeMock,
				},
			},
		},
		{
			name:  "failed to invalid argument",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {},
			req:   &request.UpdateAuthEmailRequest{},
			expect: &test.TestResponse{
				Code: http.StatusBadRequest,
			},
		},
		{
			name: "failed to update auth email",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.AuthService.EXPECT().
					UpdateAuthEmail(gomock.Any(), &user.UpdateAuthEmailRequest{
						Email: "test-user@calmato.jp",
					}).
					Return(nil, test.ErrMock)
			},
			req: &request.UpdateAuthEmailRequest{
				Email: "test-user@calmato.jp",
			},
			expect: &test.TestResponse{
				Code: http.StatusInternalServerError,
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			ctx, res, mocks := test.NewHTTPClient(t, tt.req)
			tt.setup(ctx, t, mocks)

			handler := NewAuthHandler(mocks.AuthService)
			handler.UpdateEmail(ctx)

			test.TestHTTP(t, tt.expect, res)
		})
	}
}

func TestAuth_UpdatePassword(t *testing.T) {
	t.Parallel()

	user1 := testAuth("00000000-0000-0000-0000-000000000000")

	tests := []struct {
		name   string
		setup  func(ctx context.Context, t *testing.T, mocks *test.Mocks)
		req    *request.UpdateAuthPasswordRequest
		expect *test.TestResponse
	}{
		{
			name: "success",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.AuthService.EXPECT().
					UpdateAuthPassword(gomock.Any(), &user.UpdateAuthPasswordRequest{
						Password:             "12345678",
						PasswordConfirmation: "12345678",
					}).
					Return(&user.AuthResponse{Auth: user1}, nil)
			},
			req: &request.UpdateAuthPasswordRequest{
				Password:             "12345678",
				PasswordConfirmation: "12345678",
			},
			expect: &test.TestResponse{
				Code: http.StatusOK,
				Body: &response.AuthResponse{
					ID:               "00000000-0000-0000-0000-000000000000",
					Username:         "テストユーザー",
					Gender:           entity.GenderMan,
					Email:            "test-user@calmato.jp",
					PhoneNumber:      "000-0000-0000",
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
					CreatedAt:        test.TimeMock,
					UpdatedAt:        test.TimeMock,
				},
			},
		},
		{
			name:  "failed to invalid argument",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {},
			req:   &request.UpdateAuthPasswordRequest{},
			expect: &test.TestResponse{
				Code: http.StatusBadRequest,
			},
		},
		{
			name: "failed to update auth password",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.AuthService.EXPECT().
					UpdateAuthPassword(gomock.Any(), &user.UpdateAuthPasswordRequest{
						Password:             "12345678",
						PasswordConfirmation: "12345678",
					}).
					Return(nil, test.ErrMock)
			},
			req: &request.UpdateAuthPasswordRequest{
				Password:             "12345678",
				PasswordConfirmation: "12345678",
			},
			expect: &test.TestResponse{
				Code: http.StatusInternalServerError,
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			ctx, res, mocks := test.NewHTTPClient(t, tt.req)
			tt.setup(ctx, t, mocks)

			handler := NewAuthHandler(mocks.AuthService)
			handler.UpdatePassword(ctx)

			test.TestHTTP(t, tt.expect, res)
		})
	}
}

func TestAuth_Delete(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		setup  func(ctx context.Context, t *testing.T, mocks *test.Mocks)
		req    *request.UpdateAuthPasswordRequest
		expect *test.TestResponse
	}{
		{
			name: "success",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.AuthService.EXPECT().
					DeleteAuth(gomock.Any(), &user.Empty{}).
					Return(&user.Empty{}, nil)
			},
			expect: &test.TestResponse{
				Code: http.StatusNoContent,
			},
		},
		{
			name: "failed to delete auth",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.AuthService.EXPECT().
					DeleteAuth(gomock.Any(), &user.Empty{}).
					Return(nil, test.ErrMock)
			},
			expect: &test.TestResponse{
				Code: http.StatusInternalServerError,
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			ctx, res, mocks := test.NewHTTPClient(t, tt.req)
			tt.setup(ctx, t, mocks)

			handler := NewAuthHandler(mocks.AuthService)
			handler.Delete(ctx)

			test.TestHTTP(t, tt.expect, res)
		})
	}
}

func TestAuth_RegisterDevice(t *testing.T) {
	t.Parallel()

	user1 := testAuth("00000000-0000-0000-0000-000000000000")

	tests := []struct {
		name   string
		setup  func(ctx context.Context, t *testing.T, mocks *test.Mocks)
		req    *request.RegisterAuthDeviceRequest
		expect *test.TestResponse
	}{
		{
			name: "success",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.AuthService.EXPECT().
					RegisterAuthDevice(gomock.Any(), &user.RegisterAuthDeviceRequest{
						InstanceId: "!Qaz2wsx3edc",
					}).
					Return(&user.AuthResponse{Auth: user1}, nil)
			},
			req: &request.RegisterAuthDeviceRequest{
				InstanceID: "!Qaz2wsx3edc",
			},
			expect: &test.TestResponse{
				Code: http.StatusOK,
				Body: &response.AuthResponse{
					ID:               "00000000-0000-0000-0000-000000000000",
					Username:         "テストユーザー",
					Gender:           entity.GenderMan,
					Email:            "test-user@calmato.jp",
					PhoneNumber:      "000-0000-0000",
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
					CreatedAt:        test.TimeMock,
					UpdatedAt:        test.TimeMock,
				},
			},
		},
		{
			name:  "failed to invalid argument",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {},
			req:   &request.RegisterAuthDeviceRequest{},
			expect: &test.TestResponse{
				Code: http.StatusBadRequest,
			},
		},
		{
			name: "failed to update auth password",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.AuthService.EXPECT().
					RegisterAuthDevice(gomock.Any(), &user.RegisterAuthDeviceRequest{
						InstanceId: "!Qaz2wsx3edc",
					}).
					Return(nil, test.ErrMock)
			},
			req: &request.RegisterAuthDeviceRequest{
				InstanceID: "!Qaz2wsx3edc",
			},
			expect: &test.TestResponse{
				Code: http.StatusInternalServerError,
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			ctx, res, mocks := test.NewHTTPClient(t, tt.req)
			tt.setup(ctx, t, mocks)

			handler := NewAuthHandler(mocks.AuthService)
			handler.RegisterDevice(ctx)

			test.TestHTTP(t, tt.expect, res)
		})
	}
}

func testAuth(id string) *user.Auth {
	return &user.Auth{
		Id:               id,
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
		CreatedAt:        test.TimeMock,
		UpdatedAt:        test.TimeMock,
	}
}
