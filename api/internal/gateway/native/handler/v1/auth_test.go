package v1

import (
	"context"
	"net/http"
	"testing"

	gentity "github.com/calmato/gran-book/api/internal/gateway/entity"
	"github.com/calmato/gran-book/api/internal/gateway/native/entity"
	request "github.com/calmato/gran-book/api/internal/gateway/native/request/v1"
	response "github.com/calmato/gran-book/api/internal/gateway/native/response/v1"
	mock_user "github.com/calmato/gran-book/api/mock/proto/user"
	"github.com/calmato/gran-book/api/pkg/datetime"
	"github.com/calmato/gran-book/api/pkg/test"
	"github.com/calmato/gran-book/api/proto/user"
	"github.com/golang/mock/gomock"
)

func TestAuth_Get(t *testing.T) {
	t.Parallel()

	now := datetime.FormatTime(test.TimeMock)
	user1 := testAuth("00000000-0000-0000-0000-000000000000")

	tests := []struct {
		name   string
		setup  func(ctx context.Context, t *testing.T, mocks *test.Mocks, ctrl *gomock.Controller)
		expect *test.HTTPResponse
	}{
		{
			name: "success",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks, ctrl *gomock.Controller) {
				mocks.AuthService.EXPECT().
					GetAuth(gomock.Any(), &user.Empty{}).
					Return(&user.AuthResponse{Auth: user1}, nil)
			},
			expect: &test.HTTPResponse{
				Code: http.StatusOK,
				Body: &response.AuthResponse{
					Auth: &entity.Auth{
						ID:               "00000000-0000-0000-0000-000000000000",
						Username:         "テストユーザー",
						Gender:           gentity.GenderMan,
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
						CreatedAt:        now,
						UpdatedAt:        now,
					},
				},
			},
		},
		{
			name: "failed to get auth",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks, ctrl *gomock.Controller) {
				mocks.AuthService.EXPECT().
					GetAuth(gomock.Any(), &user.Empty{}).
					Return(nil, test.ErrMock)
			},
			expect: &test.HTTPResponse{
				Code: http.StatusInternalServerError,
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			path := "/v1/auth"
			req := test.NewHTTPRequest(t, http.MethodGet, path, nil)
			testHTTP(t, tt.setup, tt.expect, req)
		})
	}
}

func TestAuth_CreateAuth(t *testing.T) {
	t.Parallel()

	now := datetime.FormatTime(test.TimeMock)
	user1 := testAuth("00000000-0000-0000-0000-000000000000")

	tests := []struct {
		name   string
		setup  func(ctx context.Context, t *testing.T, mocks *test.Mocks, ctrl *gomock.Controller)
		req    *request.CreateAuthRequest
		expect *test.HTTPResponse
	}{
		{
			name: "success",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks, ctrl *gomock.Controller) {
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
			expect: &test.HTTPResponse{
				Code: http.StatusOK,
				Body: &response.AuthResponse{
					Auth: &entity.Auth{
						ID:               "00000000-0000-0000-0000-000000000000",
						Username:         "テストユーザー",
						Gender:           gentity.GenderMan,
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
						CreatedAt:        now,
						UpdatedAt:        now,
					},
				},
			},
		},
		{
			name:  "failed to invalid argument",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks, ctrl *gomock.Controller) {},
			req:   &request.CreateAuthRequest{},
			expect: &test.HTTPResponse{
				Code: http.StatusBadRequest,
			},
		},
		{
			name: "failed to create auth",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks, ctrl *gomock.Controller) {
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
			expect: &test.HTTPResponse{
				Code: http.StatusInternalServerError,
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			path := "/v1/auth"
			req := test.NewHTTPRequest(t, http.MethodPost, path, tt.req)
			testHTTP(t, tt.setup, tt.expect, req)
		})
	}
}

func TestAuth_UpdateAuthProfile(t *testing.T) {
	t.Parallel()

	now := datetime.FormatTime(test.TimeMock)
	user1 := testAuth("00000000-0000-0000-0000-000000000000")

	tests := []struct {
		name   string
		setup  func(ctx context.Context, t *testing.T, mocks *test.Mocks, ctrl *gomock.Controller)
		req    *request.UpdateAuthProfileRequest
		expect *test.HTTPResponse
	}{
		{
			name: "success",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks, ctrl *gomock.Controller) {
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
				Gender:           gentity.GenderMan,
				ThumbnailURL:     "https://go.dev/images/gophers/ladder.svg",
				SelfIntroduction: "テストコードです。",
			},
			expect: &test.HTTPResponse{
				Code: http.StatusOK,
				Body: &response.AuthResponse{
					Auth: &entity.Auth{
						ID:               "00000000-0000-0000-0000-000000000000",
						Username:         "テストユーザー",
						Gender:           gentity.GenderMan,
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
						CreatedAt:        now,
						UpdatedAt:        now,
					},
				},
			},
		},
		{
			name:  "failed to invalid argument",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks, ctrl *gomock.Controller) {},
			req:   &request.UpdateAuthProfileRequest{},
			expect: &test.HTTPResponse{
				Code: http.StatusBadRequest,
			},
		},
		{
			name: "failed to update auth profile",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks, ctrl *gomock.Controller) {
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
				Gender:           gentity.GenderMan,
				ThumbnailURL:     "https://go.dev/images/gophers/ladder.svg",
				SelfIntroduction: "テストコードです。",
			},
			expect: &test.HTTPResponse{
				Code: http.StatusInternalServerError,
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			path := "/v1/auth/profile"
			req := test.NewHTTPRequest(t, http.MethodPatch, path, tt.req)
			testHTTP(t, tt.setup, tt.expect, req)
		})
	}
}

func TestAuth_UpdateAuthAddress(t *testing.T) {
	t.Parallel()

	now := datetime.FormatTime(test.TimeMock)
	user1 := testAuth("00000000-0000-0000-0000-000000000000")

	tests := []struct {
		name   string
		setup  func(ctx context.Context, t *testing.T, mocks *test.Mocks, ctrl *gomock.Controller)
		req    *request.UpdateAuthAddressRequest
		expect *test.HTTPResponse
	}{
		{
			name: "success",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks, ctrl *gomock.Controller) {
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
			expect: &test.HTTPResponse{
				Code: http.StatusOK,
				Body: &response.AuthResponse{
					Auth: &entity.Auth{
						ID:               "00000000-0000-0000-0000-000000000000",
						Username:         "テストユーザー",
						Gender:           gentity.GenderMan,
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
						CreatedAt:        now,
						UpdatedAt:        now,
					},
				},
			},
		},
		{
			name:  "failed to invalid argument",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks, ctrl *gomock.Controller) {},
			req:   &request.UpdateAuthAddressRequest{},
			expect: &test.HTTPResponse{
				Code: http.StatusBadRequest,
			},
		},
		{
			name: "failed to update auth address",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks, ctrl *gomock.Controller) {
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
			expect: &test.HTTPResponse{
				Code: http.StatusInternalServerError,
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			path := "/v1/auth/address"
			req := test.NewHTTPRequest(t, http.MethodPatch, path, tt.req)
			testHTTP(t, tt.setup, tt.expect, req)
		})
	}
}

func TestAuth_UpdateEmail(t *testing.T) {
	t.Parallel()

	now := datetime.FormatTime(test.TimeMock)
	user1 := testAuth("00000000-0000-0000-0000-000000000000")

	tests := []struct {
		name   string
		setup  func(ctx context.Context, t *testing.T, mocks *test.Mocks, ctrl *gomock.Controller)
		req    *request.UpdateAuthEmailRequest
		expect *test.HTTPResponse
	}{
		{
			name: "success",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks, ctrl *gomock.Controller) {
				mocks.AuthService.EXPECT().
					UpdateAuthEmail(gomock.Any(), &user.UpdateAuthEmailRequest{
						Email: "test-user@calmato.jp",
					}).
					Return(&user.AuthResponse{Auth: user1}, nil)
			},
			req: &request.UpdateAuthEmailRequest{
				Email: "test-user@calmato.jp",
			},
			expect: &test.HTTPResponse{
				Code: http.StatusOK,
				Body: &response.AuthResponse{
					Auth: &entity.Auth{
						ID:               "00000000-0000-0000-0000-000000000000",
						Username:         "テストユーザー",
						Gender:           gentity.GenderMan,
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
						CreatedAt:        now,
						UpdatedAt:        now,
					},
				},
			},
		},
		{
			name:  "failed to invalid argument",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks, ctrl *gomock.Controller) {},
			req:   &request.UpdateAuthEmailRequest{},
			expect: &test.HTTPResponse{
				Code: http.StatusBadRequest,
			},
		},
		{
			name: "failed to update auth email",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks, ctrl *gomock.Controller) {
				mocks.AuthService.EXPECT().
					UpdateAuthEmail(gomock.Any(), &user.UpdateAuthEmailRequest{
						Email: "test-user@calmato.jp",
					}).
					Return(nil, test.ErrMock)
			},
			req: &request.UpdateAuthEmailRequest{
				Email: "test-user@calmato.jp",
			},
			expect: &test.HTTPResponse{
				Code: http.StatusInternalServerError,
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			path := "/v1/auth/email"
			req := test.NewHTTPRequest(t, http.MethodPatch, path, tt.req)
			testHTTP(t, tt.setup, tt.expect, req)
		})
	}
}

func TestAuth_UpdateAuthPassword(t *testing.T) {
	t.Parallel()

	now := datetime.FormatTime(test.TimeMock)
	user1 := testAuth("00000000-0000-0000-0000-000000000000")

	tests := []struct {
		name   string
		setup  func(ctx context.Context, t *testing.T, mocks *test.Mocks, ctrl *gomock.Controller)
		req    *request.UpdateAuthPasswordRequest
		expect *test.HTTPResponse
	}{
		{
			name: "success",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks, ctrl *gomock.Controller) {
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
			expect: &test.HTTPResponse{
				Code: http.StatusOK,
				Body: &response.AuthResponse{
					Auth: &entity.Auth{
						ID:               "00000000-0000-0000-0000-000000000000",
						Username:         "テストユーザー",
						Gender:           gentity.GenderMan,
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
						CreatedAt:        now,
						UpdatedAt:        now,
					},
				},
			},
		},
		{
			name:  "failed to invalid argument",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks, ctrl *gomock.Controller) {},
			req:   &request.UpdateAuthPasswordRequest{},
			expect: &test.HTTPResponse{
				Code: http.StatusBadRequest,
			},
		},
		{
			name: "failed to update auth password",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks, ctrl *gomock.Controller) {
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
			expect: &test.HTTPResponse{
				Code: http.StatusInternalServerError,
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			path := "/v1/auth/password"
			req := test.NewHTTPRequest(t, http.MethodPatch, path, tt.req)
			testHTTP(t, tt.setup, tt.expect, req)
		})
	}
}

func TestAuth_DeleteAuth(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		setup  func(ctx context.Context, t *testing.T, mocks *test.Mocks, ctrl *gomock.Controller)
		expect *test.HTTPResponse
	}{
		{
			name: "success",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks, ctrl *gomock.Controller) {
				mocks.AuthService.EXPECT().
					DeleteAuth(gomock.Any(), &user.Empty{}).
					Return(&user.Empty{}, nil)
			},
			expect: &test.HTTPResponse{
				Code: http.StatusNoContent,
			},
		},
		{
			name: "failed to delete auth",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks, ctrl *gomock.Controller) {
				mocks.AuthService.EXPECT().
					DeleteAuth(gomock.Any(), &user.Empty{}).
					Return(nil, test.ErrMock)
			},
			expect: &test.HTTPResponse{
				Code: http.StatusInternalServerError,
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			path := "/v1/auth"
			req := test.NewHTTPRequest(t, http.MethodDelete, path, nil)
			testHTTP(t, tt.setup, tt.expect, req)
		})
	}
}

func TestAuth_UploadAuthThumbnail(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		setup  func(ctx context.Context, t *testing.T, mocks *test.Mocks, ctrl *gomock.Controller)
		field  string
		expect *test.HTTPResponse
	}{
		{
			name: "success",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks, ctrl *gomock.Controller) {
				client := mock_user.NewMockAuthService_UploadAuthThumbnailClient(ctrl)
				mocks.AuthService.EXPECT().UploadAuthThumbnail(gomock.Any()).Return(client, nil)
				client.EXPECT().Send(gomock.Any()).AnyTimes().Return(nil)
				client.EXPECT().CloseAndRecv().Return(&user.AuthThumbnailResponse{
					ThumbnailUrl: "https://go.dev/images/gophers/ladder.svg",
				}, nil)
			},
			field: "thumbnail",
			expect: &test.HTTPResponse{
				Code: http.StatusOK,
				Body: &response.AuthThumbnailResponse{
					ThumbnailURL: "https://go.dev/images/gophers/ladder.svg",
				},
			},
		},
		{
			name:  "failed to invalid argument",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks, ctrl *gomock.Controller) {},
			expect: &test.HTTPResponse{
				Code: http.StatusBadRequest,
			},
		},
		{
			name: "failed to get thumbnail",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks, ctrl *gomock.Controller) {
				mocks.AuthService.EXPECT().UploadAuthThumbnail(gomock.Any()).Return(nil, test.ErrMock)
			},
			field: "thumbnail",
			expect: &test.HTTPResponse{
				Code: http.StatusInternalServerError,
			},
		},
		{
			name: "failed to create gRPC stream client",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks, ctrl *gomock.Controller) {
				mocks.AuthService.EXPECT().UploadAuthThumbnail(gomock.Any()).Return(nil, test.ErrMock)
			},
			field: "thumbnail",
			expect: &test.HTTPResponse{
				Code: http.StatusInternalServerError,
			},
		},
		{
			name: "failed to send stream request",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks, ctrl *gomock.Controller) {
				client := mock_user.NewMockAuthService_UploadAuthThumbnailClient(ctrl)
				mocks.AuthService.EXPECT().UploadAuthThumbnail(gomock.Any()).Return(client, nil)
				client.EXPECT().Send(gomock.Any()).Return(test.ErrMock)
			},
			field: "thumbnail",
			expect: &test.HTTPResponse{
				Code: http.StatusInternalServerError,
			},
		},
		{
			name: "failed to upload thumbnail",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks, ctrl *gomock.Controller) {
				client := mock_user.NewMockAuthService_UploadAuthThumbnailClient(ctrl)
				mocks.AuthService.EXPECT().UploadAuthThumbnail(gomock.Any()).Return(client, nil)
				client.EXPECT().Send(gomock.Any()).AnyTimes().Return(nil)
				client.EXPECT().CloseAndRecv().Return(nil, test.ErrMock)
			},
			field: "thumbnail",
			expect: &test.HTTPResponse{
				Code: http.StatusInternalServerError,
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			path := "/v1/auth/thumbnail"
			req := test.NewMultipartRequest(t, http.MethodPost, path, tt.field)
			testHTTP(t, tt.setup, tt.expect, req)
		})
	}
}

func TestAuth_RegisterAuthDevice(t *testing.T) {
	t.Parallel()

	now := datetime.FormatTime(test.TimeMock)
	user1 := testAuth("00000000-0000-0000-0000-000000000000")

	tests := []struct {
		name   string
		setup  func(ctx context.Context, t *testing.T, mocks *test.Mocks, ctrl *gomock.Controller)
		req    *request.RegisterAuthDeviceRequest
		expect *test.HTTPResponse
	}{
		{
			name: "success",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks, ctrl *gomock.Controller) {
				mocks.AuthService.EXPECT().
					RegisterAuthDevice(gomock.Any(), &user.RegisterAuthDeviceRequest{
						InstanceId: "!Qaz2wsx3edc",
					}).
					Return(&user.AuthResponse{Auth: user1}, nil)
			},
			req: &request.RegisterAuthDeviceRequest{
				InstanceID: "!Qaz2wsx3edc",
			},
			expect: &test.HTTPResponse{
				Code: http.StatusOK,
				Body: &response.AuthResponse{
					Auth: &entity.Auth{
						ID:               "00000000-0000-0000-0000-000000000000",
						Username:         "テストユーザー",
						Gender:           gentity.GenderMan,
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
						CreatedAt:        now,
						UpdatedAt:        now,
					},
				},
			},
		},
		{
			name:  "failed to invalid argument",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks, ctrl *gomock.Controller) {},
			req:   &request.RegisterAuthDeviceRequest{},
			expect: &test.HTTPResponse{
				Code: http.StatusBadRequest,
			},
		},
		{
			name: "failed to update auth password",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks, ctrl *gomock.Controller) {
				mocks.AuthService.EXPECT().
					RegisterAuthDevice(gomock.Any(), &user.RegisterAuthDeviceRequest{
						InstanceId: "!Qaz2wsx3edc",
					}).
					Return(nil, test.ErrMock)
			},
			req: &request.RegisterAuthDeviceRequest{
				InstanceID: "!Qaz2wsx3edc",
			},
			expect: &test.HTTPResponse{
				Code: http.StatusInternalServerError,
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			path := "/v1/auth/device"
			req := test.NewHTTPRequest(t, http.MethodPost, path, tt.req)
			testHTTP(t, tt.setup, tt.expect, req)
		})
	}
}

func testAuth(id string) *user.Auth {
	now := datetime.FormatTime(test.TimeMock)
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
		CreatedAt:        now,
		UpdatedAt:        now,
	}
}
