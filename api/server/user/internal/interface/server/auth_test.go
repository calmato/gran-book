package server

import (
	"context"
	"testing"

	"github.com/calmato/gran-book/api/server/user/internal/domain/exception"
	"github.com/calmato/gran-book/api/server/user/internal/domain/user"
	"github.com/calmato/gran-book/api/server/user/pkg/test"
	pb "github.com/calmato/gran-book/api/server/user/proto"
	"github.com/golang/mock/gomock"
	"google.golang.org/grpc/codes"
)

func TestAuthServer_GetAuth(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	user1 := testUser("user01")

	testCases := []struct {
		name  string
		setup func(context.Context, *testing.T, *test.Mocks)
		want  *test.TestResponse
	}{
		{
			name: "success",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.UserApplication.EXPECT().
					Authentication(ctx).
					Return(user1, nil)
			},
			want: &test.TestResponse{
				Code:    codes.OK,
				Message: getAuthResponse(user1),
			},
		},
		{
			name: "failed: unauthorized",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.UserApplication.EXPECT().
					Authentication(ctx).
					Return(nil, exception.Unauthorized.New(test.ErrMock))
			},
			want: &test.TestResponse{
				Code:    codes.Unauthenticated,
				Message: nil,
			},
		},
	}

	for _, tt := range testCases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			mocks := test.NewMocks(ctrl)
			tt.setup(ctx, t, mocks)
			target := NewAuthServer(mocks.AuthRequestValidation, mocks.UserApplication)

			res, err := target.GetAuth(ctx, &pb.Empty{})
			test.TestGRPC(t, tt.want, res, err)
		})
	}
}

func TestAuthServer_CreateAuth(t *testing.T) {
	t.Parallel()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type args struct {
		req *pb.CreateAuthRequest
	}
	testCases := []struct {
		name  string
		setup func(context.Context, *testing.T, *test.Mocks)
		args  args
		want  *test.TestResponse
	}{
		{
			name: "success",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.AuthRequestValidation.EXPECT().
					CreateAuth(gomock.Any()).
					Return(nil)
				mocks.UserApplication.EXPECT().
					Create(ctx, gomock.Any()).
					Return(nil)
			},
			args: args{
				req: &pb.CreateAuthRequest{
					Username:             "テストユーザー",
					Email:                "test-user@calmato.jp",
					Password:             "12345678",
					PasswordConfirmation: "12345678",
				},
			},
			want: &test.TestResponse{
				Code: codes.OK,
				Message: &pb.AuthResponse{
					Username:  "テストユーザー",
					Email:     "test-user@calmato.jp",
					CreatedAt: "",
					UpdatedAt: "",
				},
			},
		},
		{
			name: "failed: invalid argument",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.AuthRequestValidation.EXPECT().
					CreateAuth(gomock.Any()).
					Return(exception.InvalidRequestValidation.New(test.ErrMock))
			},
			args: args{
				req: &pb.CreateAuthRequest{},
			},
			want: &test.TestResponse{
				Code:    codes.InvalidArgument,
				Message: nil,
			},
		},
		{
			name: "failed: internal error",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.AuthRequestValidation.EXPECT().
					CreateAuth(gomock.Any()).
					Return(nil)
				mocks.UserApplication.EXPECT().
					Create(ctx, gomock.Any()).
					Return(exception.ErrorInDatastore.New(test.ErrMock))
			},
			args: args{
				req: &pb.CreateAuthRequest{
					Username:             "テストユーザー",
					Email:                "test-user@calmato.jp",
					Password:             "12345678",
					PasswordConfirmation: "12345678",
				},
			},
			want: &test.TestResponse{
				Code:    codes.Internal,
				Message: nil,
			},
		},
	}

	for _, tt := range testCases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			mocks := test.NewMocks(ctrl)
			tt.setup(ctx, t, mocks)
			target := NewAuthServer(mocks.AuthRequestValidation, mocks.UserApplication)

			res, err := target.CreateAuth(ctx, tt.args.req)
			test.TestGRPC(t, tt.want, res, err)
		})
	}
}

func TestAuthServer_UpdateAuthEmail(t *testing.T) {
	t.Parallel()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type args struct {
		req *pb.UpdateAuthEmailRequest
	}
	testCases := []struct {
		name  string
		setup func(context.Context, *testing.T, *test.Mocks)
		args  args
		want  *test.TestResponse
	}{
		{
			name: "success",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				u := &user.User{}
				mocks.AuthRequestValidation.EXPECT().
					UpdateAuthEmail(gomock.Any()).
					Return(nil)
				mocks.UserApplication.EXPECT().
					Authentication(ctx).
					Return(u, nil)
				mocks.UserApplication.EXPECT().
					Update(ctx, u).
					Return(nil)
			},
			args: args{
				req: &pb.UpdateAuthEmailRequest{
					Email: "test-user@calmato.jp",
				},
			},
			want: &test.TestResponse{
				Code: codes.OK,
				Message: &pb.AuthResponse{
					Email:     "test-user@calmato.jp",
					CreatedAt: "",
					UpdatedAt: "",
				},
			},
		},
		{
			name: "failed: unauthenticated",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.UserApplication.EXPECT().
					Authentication(ctx).
					Return(nil, exception.Unauthorized.New(test.ErrMock))
			},
			args: args{
				req: &pb.UpdateAuthEmailRequest{},
			},
			want: &test.TestResponse{
				Code:    codes.Unauthenticated,
				Message: nil,
			},
		},
		{
			name: "failed: invalid argument",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				u := &user.User{}
				mocks.AuthRequestValidation.EXPECT().
					UpdateAuthEmail(gomock.Any()).
					Return(exception.InvalidRequestValidation.New(test.ErrMock))
				mocks.UserApplication.EXPECT().
					Authentication(ctx).
					Return(u, nil)
			},
			args: args{
				req: &pb.UpdateAuthEmailRequest{},
			},
			want: &test.TestResponse{
				Code:    codes.InvalidArgument,
				Message: nil,
			},
		},
		{
			name: "failed: internal error",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				u := &user.User{}
				mocks.AuthRequestValidation.EXPECT().
					UpdateAuthEmail(gomock.Any()).
					Return(nil)
				mocks.UserApplication.EXPECT().
					Authentication(ctx).
					Return(u, nil)
				mocks.UserApplication.EXPECT().
					Update(ctx, u).
					Return(exception.ErrorInDatastore.New(test.ErrMock))
			},
			args: args{
				req: &pb.UpdateAuthEmailRequest{
					Email: "test-user@calmato.jp",
				},
			},
			want: &test.TestResponse{
				Code:    codes.Internal,
				Message: nil,
			},
		},
	}

	for _, tt := range testCases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			mocks := test.NewMocks(ctrl)
			tt.setup(ctx, t, mocks)
			target := NewAuthServer(mocks.AuthRequestValidation, mocks.UserApplication)

			res, err := target.UpdateAuthEmail(ctx, tt.args.req)
			test.TestGRPC(t, tt.want, res, err)
		})
	}
}

func TestAuthServer_UpdateAuthPassword(t *testing.T) {
	t.Parallel()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type args struct {
		req *pb.UpdateAuthPasswordRequest
	}
	testCases := []struct {
		name  string
		setup func(context.Context, *testing.T, *test.Mocks)
		args  args
		want  *test.TestResponse
	}{
		{
			name: "success",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				u := &user.User{}
				mocks.AuthRequestValidation.EXPECT().
					UpdateAuthPassword(gomock.Any()).
					Return(nil)
				mocks.UserApplication.EXPECT().
					Authentication(ctx).
					Return(u, nil)
				mocks.UserApplication.EXPECT().
					UpdatePassword(ctx, u).
					Return(nil)
			},
			args: args{
				req: &pb.UpdateAuthPasswordRequest{
					Password:             "12345678",
					PasswordConfirmation: "12345678",
				},
			},
			want: &test.TestResponse{
				Code: codes.OK,
				Message: &pb.AuthResponse{
					CreatedAt: "",
					UpdatedAt: "",
				},
			},
		},
		{
			name: "failed: unauthenticated",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.UserApplication.EXPECT().
					Authentication(ctx).
					Return(nil, exception.Unauthorized.New(test.ErrMock))
			},
			args: args{
				req: &pb.UpdateAuthPasswordRequest{},
			},
			want: &test.TestResponse{
				Code:    codes.Unauthenticated,
				Message: nil,
			},
		},
		{
			name: "failed: invalid argument",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				u := &user.User{}
				mocks.AuthRequestValidation.EXPECT().
					UpdateAuthPassword(gomock.Any()).
					Return(exception.InvalidRequestValidation.New(test.ErrMock))
				mocks.UserApplication.EXPECT().
					Authentication(ctx).
					Return(u, nil)
			},
			args: args{
				req: &pb.UpdateAuthPasswordRequest{},
			},
			want: &test.TestResponse{
				Code:    codes.InvalidArgument,
				Message: nil,
			},
		},
		{
			name: "failed: internal error",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				u := &user.User{}
				mocks.AuthRequestValidation.EXPECT().
					UpdateAuthPassword(gomock.Any()).
					Return(nil)
				mocks.UserApplication.EXPECT().
					Authentication(ctx).
					Return(u, nil)
				mocks.UserApplication.EXPECT().
					UpdatePassword(ctx, u).
					Return(exception.ErrorInDatastore.New(test.ErrMock))
			},
			args: args{
				req: &pb.UpdateAuthPasswordRequest{
					Password:             "12345678",
					PasswordConfirmation: "12345678",
				},
			},
			want: &test.TestResponse{
				Code:    codes.Internal,
				Message: nil,
			},
		},
	}

	for _, tt := range testCases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			mocks := test.NewMocks(ctrl)
			tt.setup(ctx, t, mocks)
			target := NewAuthServer(mocks.AuthRequestValidation, mocks.UserApplication)

			res, err := target.UpdateAuthPassword(ctx, tt.args.req)
			test.TestGRPC(t, tt.want, res, err)
		})
	}
}

func TestAuthServer_UpdateAuthProfile(t *testing.T) {
	t.Parallel()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type args struct {
		req *pb.UpdateAuthProfileRequest
	}
	testCases := []struct {
		name  string
		setup func(context.Context, *testing.T, *test.Mocks)
		args  args
		want  *test.TestResponse
	}{
		{
			name: "success",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				u := &user.User{}
				mocks.AuthRequestValidation.EXPECT().
					UpdateAuthProfile(gomock.Any()).
					Return(nil)
				mocks.UserApplication.EXPECT().
					Authentication(ctx).
					Return(u, nil)
				mocks.UserApplication.EXPECT().
					Update(ctx, u).
					Return(nil)
			},
			args: args{
				req: &pb.UpdateAuthProfileRequest{
					Username:         "test-user",
					Gender:           pb.Gender_GENDER_MAN,
					ThumbnailUrl:     "https://go.dev/images/gophers/ladder.svg",
					SelfIntroduction: "自己紹介",
				},
			},
			want: &test.TestResponse{
				Code: codes.OK,
				Message: &pb.AuthResponse{
					Username:         "test-user",
					Gender:           pb.Gender_GENDER_MAN,
					ThumbnailUrl:     "https://go.dev/images/gophers/ladder.svg",
					SelfIntroduction: "自己紹介",
					CreatedAt:        "",
					UpdatedAt:        "",
				},
			},
		},
		{
			name: "failed: unauthenticated",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.UserApplication.EXPECT().
					Authentication(ctx).
					Return(nil, exception.Unauthorized.New(test.ErrMock))
			},
			args: args{
				req: &pb.UpdateAuthProfileRequest{},
			},
			want: &test.TestResponse{
				Code:    codes.Unauthenticated,
				Message: nil,
			},
		},
		{
			name: "failed: invalid argument",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				u := &user.User{}
				mocks.AuthRequestValidation.EXPECT().
					UpdateAuthProfile(gomock.Any()).
					Return(exception.InvalidRequestValidation.New(test.ErrMock))
				mocks.UserApplication.EXPECT().
					Authentication(ctx).
					Return(u, nil)
			},
			args: args{
				req: &pb.UpdateAuthProfileRequest{},
			},
			want: &test.TestResponse{
				Code:    codes.InvalidArgument,
				Message: nil,
			},
		},
		{
			name: "failed: internal error",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				u := &user.User{}
				mocks.AuthRequestValidation.EXPECT().
					UpdateAuthProfile(gomock.Any()).
					Return(nil)
				mocks.UserApplication.EXPECT().
					Authentication(ctx).
					Return(u, nil)
				mocks.UserApplication.EXPECT().
					Update(ctx, u).
					Return(exception.ErrorInDatastore.New(test.ErrMock))
			},
			args: args{
				req: &pb.UpdateAuthProfileRequest{
					Username:         "test-user",
					Gender:           pb.Gender_GENDER_MAN,
					ThumbnailUrl:     "https://go.dev/images/gophers/ladder.svg",
					SelfIntroduction: "自己紹介",
				},
			},
			want: &test.TestResponse{
				Code:    codes.Internal,
				Message: nil,
			},
		},
	}

	for _, tt := range testCases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			mocks := test.NewMocks(ctrl)
			tt.setup(ctx, t, mocks)
			target := NewAuthServer(mocks.AuthRequestValidation, mocks.UserApplication)

			res, err := target.UpdateAuthProfile(ctx, tt.args.req)
			test.TestGRPC(t, tt.want, res, err)
		})
	}
}

func TestAuthServer_UpdateAuthAddress(t *testing.T) {
	t.Parallel()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type args struct {
		req *pb.UpdateAuthAddressRequest
	}
	testCases := []struct {
		name  string
		setup func(context.Context, *testing.T, *test.Mocks)
		args  args
		want  *test.TestResponse
	}{
		{
			name: "success",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				u := &user.User{}
				mocks.AuthRequestValidation.EXPECT().
					UpdateAuthAddress(gomock.Any()).
					Return(nil)
				mocks.UserApplication.EXPECT().
					Authentication(ctx).
					Return(u, nil)
				mocks.UserApplication.EXPECT().
					Update(ctx, u).
					Return(nil)
			},
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
			want: &test.TestResponse{
				Code: codes.OK,
				Message: &pb.AuthResponse{
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
					CreatedAt:     "",
					UpdatedAt:     "",
				},
			},
		},
		{
			name: "failed: unauthenticated",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.UserApplication.EXPECT().
					Authentication(ctx).
					Return(nil, exception.Unauthorized.New(test.ErrMock))
			},
			args: args{
				req: &pb.UpdateAuthAddressRequest{},
			},
			want: &test.TestResponse{
				Code:    codes.Unauthenticated,
				Message: nil,
			},
		},
		{
			name: "failed: invalid argument",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				u := &user.User{}
				mocks.AuthRequestValidation.EXPECT().
					UpdateAuthAddress(gomock.Any()).
					Return(exception.InvalidRequestValidation.New(test.ErrMock))
				mocks.UserApplication.EXPECT().
					Authentication(ctx).
					Return(u, nil)
			},
			args: args{
				req: &pb.UpdateAuthAddressRequest{},
			},
			want: &test.TestResponse{
				Code:    codes.InvalidArgument,
				Message: nil,
			},
		},
		{
			name: "failed: internal error",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				u := &user.User{}
				mocks.AuthRequestValidation.EXPECT().
					UpdateAuthAddress(gomock.Any()).
					Return(nil)
				mocks.UserApplication.EXPECT().
					Authentication(ctx).
					Return(u, nil)
				mocks.UserApplication.EXPECT().
					Update(ctx, u).
					Return(exception.ErrorInDatastore.New(test.ErrMock))
			},
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
			want: &test.TestResponse{
				Code:    codes.Internal,
				Message: nil,
			},
		},
	}

	for _, tt := range testCases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			mocks := test.NewMocks(ctrl)
			tt.setup(ctx, t, mocks)
			target := NewAuthServer(mocks.AuthRequestValidation, mocks.UserApplication)

			res, err := target.UpdateAuthAddress(ctx, tt.args.req)
			test.TestGRPC(t, tt.want, res, err)
		})
	}
}

func TestAuthServer_DeleteAuth(t *testing.T) {
	t.Parallel()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	testCases := []struct {
		name  string
		setup func(context.Context, *testing.T, *test.Mocks)
		want  *test.TestResponse
	}{
		{
			name: "success",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				u := &user.User{}
				mocks.UserApplication.EXPECT().
					Authentication(ctx).
					Return(u, nil)
				mocks.UserApplication.EXPECT().
					Delete(ctx, u).
					Return(nil)
			},
			want: &test.TestResponse{
				Code:    codes.OK,
				Message: &pb.Empty{},
			},
		},
		{
			name: "failed: unauthenticated",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.UserApplication.EXPECT().
					Authentication(ctx).
					Return(nil, exception.Unauthorized.New(test.ErrMock))
			},
			want: &test.TestResponse{
				Code:    codes.Unauthenticated,
				Message: nil,
			},
		},
		{
			name: "failed: internal error",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				u := &user.User{}
				mocks.UserApplication.EXPECT().
					Authentication(ctx).
					Return(u, nil)
				mocks.UserApplication.EXPECT().
					Delete(ctx, u).
					Return(exception.ErrorInDatastore.New(test.ErrMock))
			},
			want: &test.TestResponse{
				Code:    codes.Internal,
				Message: nil,
			},
		},
	}

	for _, tt := range testCases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			mocks := test.NewMocks(ctrl)
			tt.setup(ctx, t, mocks)
			target := NewAuthServer(mocks.AuthRequestValidation, mocks.UserApplication)

			res, err := target.DeleteAuth(ctx, &pb.Empty{})
			test.TestGRPC(t, tt.want, res, err)
		})
	}
}

func TestAuthServer_RegisterAuthDevice(t *testing.T) {
	t.Parallel()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type args struct {
		req *pb.RegisterAuthDeviceRequest
	}
	testCases := []struct {
		name  string
		setup func(context.Context, *testing.T, *test.Mocks)
		args  args
		want  *test.TestResponse
	}{
		{
			name: "success",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				u := &user.User{}
				mocks.AuthRequestValidation.EXPECT().
					RegisterAuthDevice(gomock.Any()).
					Return(nil)
				mocks.UserApplication.EXPECT().
					Authentication(ctx).
					Return(u, nil)
				mocks.UserApplication.EXPECT().
					Update(ctx, u).
					Return(nil)
			},
			args: args{
				req: &pb.RegisterAuthDeviceRequest{
					InstanceId: "ExponentPushToken[!Qaz2wsx3edc4rfv5tgb6y]",
				},
			},
			want: &test.TestResponse{
				Code: codes.OK,
				Message: &pb.AuthResponse{
					CreatedAt: "",
					UpdatedAt: "",
				},
			},
		},
		{
			name: "failed: unauthenticated",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.UserApplication.EXPECT().
					Authentication(ctx).
					Return(nil, exception.Unauthorized.New(test.ErrMock))
			},
			args: args{
				req: &pb.RegisterAuthDeviceRequest{},
			},
			want: &test.TestResponse{
				Code:    codes.Unauthenticated,
				Message: nil,
			},
		},
		{
			name: "failed: invalid argument",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				u := &user.User{}
				mocks.AuthRequestValidation.EXPECT().
					RegisterAuthDevice(gomock.Any()).
					Return(exception.InvalidRequestValidation.New(test.ErrMock))
				mocks.UserApplication.EXPECT().
					Authentication(ctx).
					Return(u, nil)
			},
			args: args{
				req: &pb.RegisterAuthDeviceRequest{},
			},
			want: &test.TestResponse{
				Code:    codes.InvalidArgument,
				Message: nil,
			},
		},
		{
			name: "failed: internal error",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				u := &user.User{}
				mocks.AuthRequestValidation.EXPECT().
					RegisterAuthDevice(gomock.Any()).
					Return(nil)
				mocks.UserApplication.EXPECT().
					Authentication(ctx).
					Return(u, nil)
				mocks.UserApplication.EXPECT().
					Update(ctx, u).
					Return(exception.ErrorInDatastore.New(test.ErrMock))
			},
			args: args{
				req: &pb.RegisterAuthDeviceRequest{
					InstanceId: "ExponentPushToken[!Qaz2wsx3edc4rfv5tgb6y]",
				},
			},
			want: &test.TestResponse{
				Code:    codes.Internal,
				Message: nil,
			},
		},
	}

	for _, tt := range testCases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			mocks := test.NewMocks(ctrl)
			tt.setup(ctx, t, mocks)
			target := NewAuthServer(mocks.AuthRequestValidation, mocks.UserApplication)

			res, err := target.RegisterAuthDevice(ctx, tt.args.req)
			test.TestGRPC(t, tt.want, res, err)
		})
	}
}
