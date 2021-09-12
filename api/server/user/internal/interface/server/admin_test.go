package server

import (
	"context"
	"testing"

	"github.com/calmato/gran-book/api/server/user/internal/domain/exception"
	"github.com/calmato/gran-book/api/server/user/internal/domain/user"
	"github.com/calmato/gran-book/api/server/user/pkg/test"
	pb "github.com/calmato/gran-book/api/server/user/proto/service/user"
	"github.com/golang/mock/gomock"
	"google.golang.org/grpc/codes"
)

func TestAdminServer_ListAdmin(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	user1 := testUser("user01")
	user2 := testUser("user02")

	type args struct {
		req *pb.ListAdminRequest
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
				mocks.AdminRequestValidation.EXPECT().
					ListAdmin(gomock.Any()).
					Return(nil)
				mocks.UserApplication.EXPECT().
					ListAdmin(ctx, gomock.Any()).
					Return([]*user.User{user1, user2}, 2, nil)
			},
			args: args{
				req: &pb.ListAdminRequest{
					Limit:  100,
					Offset: 0,
				},
			},
			want: &test.TestResponse{
				Code:    codes.OK,
				Message: getAdminListResponse([]*user.User{user1, user2}, 100, 0, 2),
			},
		},
		{
			name: "failed: invalid argument",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.AdminRequestValidation.EXPECT().
					ListAdmin(gomock.Any()).
					Return(exception.InvalidRequestValidation.New(test.ErrMock))
			},
			args: args{
				req: &pb.ListAdminRequest{},
			},
			want: &test.TestResponse{
				Code:    codes.InvalidArgument,
				Message: nil,
			},
		},
		{
			name: "failed: internel error",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.AdminRequestValidation.EXPECT().
					ListAdmin(gomock.Any()).
					Return(nil)
				mocks.UserApplication.EXPECT().
					ListAdmin(ctx, gomock.Any()).
					Return(nil, 0, exception.ErrorInDatastore.New(test.ErrMock))
			},
			args: args{
				req: &pb.ListAdminRequest{
					Limit:  100,
					Offset: 0,
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
			target := NewAdminServer(mocks.AdminRequestValidation, mocks.UserApplication)

			res, err := target.ListAdmin(ctx, tt.args.req)
			test.TestGRPC(t, tt.want, res, err)
		})
	}
}

func TestAdminServer_GetAdmin(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	user1 := testUser("user01")

	type args struct {
		req *pb.GetAdminRequest
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
				mocks.AdminRequestValidation.EXPECT().
					GetAdmin(gomock.Any()).
					Return(nil)
				mocks.UserApplication.EXPECT().
					GetAdmin(ctx, "user01").
					Return(user1, nil)
			},
			args: args{
				req: &pb.GetAdminRequest{
					UserId: "user01",
				},
			},
			want: &test.TestResponse{
				Code:    codes.OK,
				Message: getAdminResponse(user1),
			},
		},
		{
			name: "failed: invalid argument",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.AdminRequestValidation.EXPECT().
					GetAdmin(gomock.Any()).
					Return(exception.InvalidRequestValidation.New(test.ErrMock))
			},
			args: args{
				req: &pb.GetAdminRequest{},
			},
			want: &test.TestResponse{
				Code:    codes.InvalidArgument,
				Message: nil,
			},
		},
		{
			name: "failed: internel error",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.AdminRequestValidation.EXPECT().
					GetAdmin(gomock.Any()).
					Return(nil)
				mocks.UserApplication.EXPECT().
					GetAdmin(ctx, "user01").
					Return(nil, exception.ErrorInDatastore.New(test.ErrMock))
			},
			args: args{
				req: &pb.GetAdminRequest{
					UserId: "user01",
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
			target := NewAdminServer(mocks.AdminRequestValidation, mocks.UserApplication)

			res, err := target.GetAdmin(ctx, tt.args.req)
			test.TestGRPC(t, tt.want, res, err)
		})
	}
}

func TestAdminServer_CreateAdmin(t *testing.T) {
	t.Parallel()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type args struct {
		req *pb.CreateAdminRequest
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
				mocks.AdminRequestValidation.EXPECT().
					CreateAdmin(gomock.Any()).
					Return(nil)
				mocks.UserApplication.EXPECT().
					Create(ctx, gomock.Any()).
					Return(nil)
			},
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
			want: &test.TestResponse{
				Code: codes.OK,
				Message: &pb.AdminResponse{
					Admin: &pb.Admin{
						Username:      "テストユーザー",
						Email:         "test-user@calmato.jp",
						PhoneNumber:   "000-0000-0000",
						Role:          pb.Role_ROLE_DEVELOPER,
						LastName:      "テスト",
						FirstName:     "ユーザー",
						LastNameKana:  "てすと",
						FirstNameKana: "ゆーざー",
						CreatedAt:     "",
						UpdatedAt:     "",
					},
				},
			},
		},
		{
			name: "failed: invalid argument",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.AdminRequestValidation.EXPECT().
					CreateAdmin(gomock.Any()).
					Return(exception.InvalidRequestValidation.New(test.ErrMock))
			},
			args: args{
				req: &pb.CreateAdminRequest{},
			},
			want: &test.TestResponse{
				Code:    codes.InvalidArgument,
				Message: nil,
			},
		},
		{
			name: "failed: internal error",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.AdminRequestValidation.EXPECT().
					CreateAdmin(gomock.Any()).
					Return(nil)
				mocks.UserApplication.EXPECT().
					Create(ctx, gomock.Any()).
					Return(exception.ErrorInDatastore.New(test.ErrMock))
			},
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
			target := NewAdminServer(mocks.AdminRequestValidation, mocks.UserApplication)

			res, err := target.CreateAdmin(ctx, tt.args.req)
			test.TestGRPC(t, tt.want, res, err)
		})
	}
}

func TestAdminServer_UpdateAdminContact(t *testing.T) {
	t.Parallel()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type args struct {
		req *pb.UpdateAdminContactRequest
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
				mocks.AdminRequestValidation.EXPECT().
					UpdateAdminContact(gomock.Any()).
					Return(nil)
				mocks.UserApplication.EXPECT().
					GetAdmin(ctx, "user01").
					Return(u, nil)
				mocks.UserApplication.EXPECT().
					Update(ctx, u).
					Return(nil)
			},
			args: args{
				req: &pb.UpdateAdminContactRequest{
					UserId:      "user01",
					Email:       "test-user@calmato.jp",
					PhoneNumber: "000-0000-0000",
				},
			},
			want: &test.TestResponse{
				Code: codes.OK,
				Message: &pb.AdminResponse{
					Admin: &pb.Admin{
						Email:       "test-user@calmato.jp",
						PhoneNumber: "000-0000-0000",
						CreatedAt:   "",
						UpdatedAt:   "",
					},
				},
			},
		},
		{
			name: "failed: invalid argument",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.AdminRequestValidation.EXPECT().
					UpdateAdminContact(gomock.Any()).
					Return(exception.InvalidRequestValidation.New(test.ErrMock))
			},
			args: args{
				req: &pb.UpdateAdminContactRequest{},
			},
			want: &test.TestResponse{
				Code:    codes.InvalidArgument,
				Message: nil,
			},
		},
		{
			name: "failed: not found",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.AdminRequestValidation.EXPECT().
					UpdateAdminContact(gomock.Any()).
					Return(nil)
				mocks.UserApplication.EXPECT().
					GetAdmin(ctx, "user01").
					Return(nil, exception.NotFound.New(test.ErrMock))
			},
			args: args{
				req: &pb.UpdateAdminContactRequest{
					UserId: "user01",
					Email:  "test-user@calmato.jp",
				},
			},
			want: &test.TestResponse{
				Code:    codes.NotFound,
				Message: nil,
			},
		},
		{
			name: "failed: internal error",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				u := &user.User{}
				mocks.AdminRequestValidation.EXPECT().
					UpdateAdminContact(gomock.Any()).
					Return(nil)
				mocks.UserApplication.EXPECT().
					GetAdmin(ctx, "user01").
					Return(u, nil)
				mocks.UserApplication.EXPECT().
					Update(ctx, u).
					Return(exception.ErrorInDatastore.New(test.ErrMock))
			},
			args: args{
				req: &pb.UpdateAdminContactRequest{
					UserId: "user01",
					Email:  "test-user@calmato.jp",
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
			target := NewAdminServer(mocks.AdminRequestValidation, mocks.UserApplication)

			res, err := target.UpdateAdminContact(ctx, tt.args.req)
			test.TestGRPC(t, tt.want, res, err)
		})
	}
}

func TestAdminServer_UpdateAdminPassword(t *testing.T) {
	t.Parallel()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type args struct {
		req *pb.UpdateAdminPasswordRequest
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
				mocks.AdminRequestValidation.EXPECT().
					UpdateAdminPassword(gomock.Any()).
					Return(nil)
				mocks.UserApplication.EXPECT().
					GetAdmin(ctx, "user01").
					Return(u, nil)
				mocks.UserApplication.EXPECT().
					UpdatePassword(ctx, u).
					Return(nil)
			},
			args: args{
				req: &pb.UpdateAdminPasswordRequest{
					UserId:               "user01",
					Password:             "12345678",
					PasswordConfirmation: "12345678",
				},
			},
			want: &test.TestResponse{
				Code: codes.OK,
				Message: &pb.AdminResponse{
					Admin: &pb.Admin{
						CreatedAt: "",
						UpdatedAt: "",
					},
				},
			},
		},
		{
			name: "failed: invalid argument",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.AdminRequestValidation.EXPECT().
					UpdateAdminPassword(gomock.Any()).
					Return(exception.InvalidRequestValidation.New(test.ErrMock))
			},
			args: args{
				req: &pb.UpdateAdminPasswordRequest{},
			},
			want: &test.TestResponse{
				Code:    codes.InvalidArgument,
				Message: nil,
			},
		},
		{
			name: "failed: not found",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.AdminRequestValidation.EXPECT().
					UpdateAdminPassword(gomock.Any()).
					Return(nil)
				mocks.UserApplication.EXPECT().
					GetAdmin(ctx, "user01").
					Return(nil, exception.NotFound.New(test.ErrMock))
			},
			args: args{
				req: &pb.UpdateAdminPasswordRequest{
					UserId: "user01",
				},
			},
			want: &test.TestResponse{
				Code:    codes.NotFound,
				Message: nil,
			},
		},
		{
			name: "failed: internal error",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				u := &user.User{}
				mocks.AdminRequestValidation.EXPECT().
					UpdateAdminPassword(gomock.Any()).
					Return(nil)
				mocks.UserApplication.EXPECT().
					GetAdmin(ctx, "user01").
					Return(u, nil)
				mocks.UserApplication.EXPECT().
					UpdatePassword(ctx, u).
					Return(exception.ErrorInDatastore.New(test.ErrMock))
			},
			args: args{
				req: &pb.UpdateAdminPasswordRequest{
					UserId:               "user01",
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
			target := NewAdminServer(mocks.AdminRequestValidation, mocks.UserApplication)

			res, err := target.UpdateAdminPassword(ctx, tt.args.req)
			test.TestGRPC(t, tt.want, res, err)
		})
	}
}

func TestAdminServer_UpdateAdminProfile(t *testing.T) {
	t.Parallel()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type args struct {
		req *pb.UpdateAdminProfileRequest
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
				mocks.AdminRequestValidation.EXPECT().
					UpdateAdminProfile(gomock.Any()).
					Return(nil)
				mocks.UserApplication.EXPECT().
					GetAdmin(ctx, "user01").
					Return(u, nil)
				mocks.UserApplication.EXPECT().
					Update(ctx, u).
					Return(nil)
			},
			args: args{
				req: &pb.UpdateAdminProfileRequest{
					UserId:        "user01",
					Username:      "テストユーザー",
					Role:          pb.Role_ROLE_DEVELOPER,
					LastName:      "テスト",
					FirstName:     "ユーザー",
					LastNameKana:  "てすと",
					FirstNameKana: "ゆーざー",
					ThumbnailUrl:  "https://go.dev/images/gophers/ladder.svg",
				},
			},
			want: &test.TestResponse{
				Code: codes.OK,
				Message: &pb.AdminResponse{
					Admin: &pb.Admin{
						Username:      "テストユーザー",
						Role:          pb.Role_ROLE_DEVELOPER,
						LastName:      "テスト",
						FirstName:     "ユーザー",
						LastNameKana:  "てすと",
						FirstNameKana: "ゆーざー",
						ThumbnailUrl:  "https://go.dev/images/gophers/ladder.svg",
						CreatedAt:     "",
						UpdatedAt:     "",
					},
				},
			},
		},
		{
			name: "failed: invalid argument",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.AdminRequestValidation.EXPECT().
					UpdateAdminProfile(gomock.Any()).
					Return(exception.InvalidRequestValidation.New(test.ErrMock))
			},
			args: args{
				req: &pb.UpdateAdminProfileRequest{},
			},
			want: &test.TestResponse{
				Code:    codes.InvalidArgument,
				Message: nil,
			},
		},
		{
			name: "failed: not found",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.AdminRequestValidation.EXPECT().
					UpdateAdminProfile(gomock.Any()).
					Return(nil)
				mocks.UserApplication.EXPECT().
					GetAdmin(ctx, "user01").
					Return(nil, exception.NotFound.New(test.ErrMock))
			},
			args: args{
				req: &pb.UpdateAdminProfileRequest{
					UserId: "user01",
				},
			},
			want: &test.TestResponse{
				Code:    codes.NotFound,
				Message: nil,
			},
		},
		{
			name: "failed: internal error",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				u := &user.User{}
				mocks.AdminRequestValidation.EXPECT().
					UpdateAdminProfile(gomock.Any()).
					Return(nil)
				mocks.UserApplication.EXPECT().
					GetAdmin(ctx, "user01").
					Return(u, nil)
				mocks.UserApplication.EXPECT().
					Update(ctx, u).
					Return(exception.ErrorInDatastore.New(test.ErrMock))
			},
			args: args{
				req: &pb.UpdateAdminProfileRequest{
					UserId:        "user01",
					Username:      "テストユーザー",
					Role:          pb.Role_ROLE_DEVELOPER,
					LastName:      "テスト",
					FirstName:     "ユーザー",
					LastNameKana:  "てすと",
					FirstNameKana: "ゆーざー",
					ThumbnailUrl:  "https://go.dev/images/gophers/ladder.svg",
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
			target := NewAdminServer(mocks.AdminRequestValidation, mocks.UserApplication)

			res, err := target.UpdateAdminProfile(ctx, tt.args.req)
			test.TestGRPC(t, tt.want, res, err)
		})
	}
}

func TestAdminServer_DeleteAdmin(t *testing.T) {
	t.Parallel()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type args struct {
		req *pb.DeleteAdminRequest
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
				mocks.AdminRequestValidation.EXPECT().
					DeleteAdmin(gomock.Any()).
					Return(nil)
				mocks.UserApplication.EXPECT().
					GetAdmin(ctx, "user01").
					Return(u, nil)
				mocks.UserApplication.EXPECT().
					DeleteAdmin(ctx, u).
					Return(nil)
			},
			args: args{
				req: &pb.DeleteAdminRequest{
					UserId: "user01",
				},
			},
			want: &test.TestResponse{
				Code:    codes.OK,
				Message: &pb.Empty{},
			},
		},
		{
			name: "failed: invalid argument",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.AdminRequestValidation.EXPECT().
					DeleteAdmin(gomock.Any()).
					Return(exception.InvalidRequestValidation.New(test.ErrMock))
			},
			args: args{
				req: &pb.DeleteAdminRequest{},
			},
			want: &test.TestResponse{
				Code:    codes.InvalidArgument,
				Message: nil,
			},
		},
		{
			name: "failed: not found",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.AdminRequestValidation.EXPECT().
					DeleteAdmin(gomock.Any()).
					Return(nil)
				mocks.UserApplication.EXPECT().
					GetAdmin(ctx, "user01").
					Return(nil, exception.NotFound.New(test.ErrMock))
			},
			args: args{
				req: &pb.DeleteAdminRequest{
					UserId: "user01",
				},
			},
			want: &test.TestResponse{
				Code:    codes.NotFound,
				Message: nil,
			},
		},
		{
			name: "failed: internal error",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				u := &user.User{}
				mocks.AdminRequestValidation.EXPECT().
					DeleteAdmin(gomock.Any()).
					Return(nil)
				mocks.UserApplication.EXPECT().
					GetAdmin(ctx, "user01").
					Return(u, nil)
				mocks.UserApplication.EXPECT().
					DeleteAdmin(ctx, u).
					Return(exception.ErrorInDatastore.New(test.ErrMock))
			},
			args: args{
				req: &pb.DeleteAdminRequest{
					UserId: "user01",
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
			target := NewAdminServer(mocks.AdminRequestValidation, mocks.UserApplication)

			res, err := target.DeleteAdmin(ctx, tt.args.req)
			test.TestGRPC(t, tt.want, res, err)
		})
	}
}
