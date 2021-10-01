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

func TestUserServer_ListUser(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	user1 := testUser("user01")
	user2 := testUser("user02")

	type args struct {
		req *pb.ListUserRequest
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
				mocks.UserRequestValidation.EXPECT().
					ListUser(gomock.Any()).
					Return(nil)
				mocks.UserApplication.EXPECT().
					List(ctx, gomock.Any()).
					Return([]*user.User{user1, user2}, 2, nil)
			},
			args: args{
				req: &pb.ListUserRequest{
					Limit:  100,
					Offset: 0,
				},
			},
			want: &test.TestResponse{
				Code:    codes.OK,
				Message: getUserListResponse([]*user.User{user1, user2}, 100, 0, 2),
			},
		},
		{
			name: "failed: invalid argument",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.UserRequestValidation.EXPECT().
					ListUser(gomock.Any()).
					Return(exception.InvalidRequestValidation.New(test.ErrMock))
			},
			args: args{
				req: &pb.ListUserRequest{},
			},
			want: &test.TestResponse{
				Code:    codes.InvalidArgument,
				Message: nil,
			},
		},
		{
			name: "failed: internel error",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.UserRequestValidation.EXPECT().
					ListUser(gomock.Any()).
					Return(nil)
				mocks.UserApplication.EXPECT().
					List(ctx, gomock.Any()).
					Return(nil, 0, exception.ErrorInDatastore.New(test.ErrMock))
			},
			args: args{
				req: &pb.ListUserRequest{
					Limit:  100,
					Offset: 0,
					Search: &pb.Search{
						Field: "email",
						Value: "test@calmato.jp",
					},
					Order: &pb.Order{
						Field:   "id",
						OrderBy: pb.OrderBy_ORDER_BY_DESC,
					},
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
			target := NewUserServer(mocks.UserRequestValidation, mocks.UserApplication)

			res, err := target.ListUser(ctx, tt.args.req)
			test.TestGRPC(t, tt.want, res, err)
		})
	}
}

func TestUserServer_ListFollow(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	follow1 := testFollow("user01")
	follow2 := testFollow("user02")

	type args struct {
		req *pb.ListFollowRequest
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
				mocks.UserRequestValidation.EXPECT().
					ListFollow(gomock.Any()).
					Return(nil)
				mocks.UserApplication.EXPECT().
					Authentication(ctx).
					Return(&user.User{ID: "current-user"}, nil)
				mocks.UserApplication.EXPECT().
					ListFollow(ctx, "current-user", "12345678-1234-1234-123456789012", 100, 0).
					Return([]*user.Follow{follow1, follow2}, 2, nil)
			},
			args: args{
				req: &pb.ListFollowRequest{
					UserId: "12345678-1234-1234-123456789012",
					Limit:  100,
					Offset: 0,
				},
			},
			want: &test.TestResponse{
				Code:    codes.OK,
				Message: getFollowListResponse([]*user.Follow{follow1, follow2}, 100, 0, 2),
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
				req: &pb.ListFollowRequest{},
			},
			want: &test.TestResponse{
				Code:    codes.Unauthenticated,
				Message: nil,
			},
		},
		{
			name: "failed: invalid argument",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.UserRequestValidation.EXPECT().
					ListFollow(gomock.Any()).
					Return(exception.InvalidRequestValidation.New(test.ErrMock))
				mocks.UserApplication.EXPECT().
					Authentication(ctx).
					Return(&user.User{ID: "current-user"}, nil)
			},
			args: args{
				req: &pb.ListFollowRequest{},
			},
			want: &test.TestResponse{
				Code:    codes.InvalidArgument,
				Message: nil,
			},
		},
		{
			name: "failed: internel error",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.UserRequestValidation.EXPECT().
					ListFollow(gomock.Any()).
					Return(nil)
				mocks.UserApplication.EXPECT().
					Authentication(ctx).
					Return(&user.User{ID: "current-user"}, nil)
				mocks.UserApplication.EXPECT().
					ListFollow(ctx, "current-user", "12345678-1234-1234-123456789012", 100, 0).
					Return(nil, 0, exception.ErrorInDatastore.New(test.ErrMock))
			},
			args: args{
				req: &pb.ListFollowRequest{
					UserId: "12345678-1234-1234-123456789012",
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
			target := NewUserServer(mocks.UserRequestValidation, mocks.UserApplication)

			res, err := target.ListFollow(ctx, tt.args.req)
			test.TestGRPC(t, tt.want, res, err)
		})
	}
}

func TestUserServer_ListFollower(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	follower1 := testFollower("user01")
	follower2 := testFollower("user02")

	type args struct {
		req *pb.ListFollowerRequest
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
				mocks.UserRequestValidation.EXPECT().
					ListFollower(gomock.Any()).
					Return(nil)
				mocks.UserApplication.EXPECT().
					Authentication(ctx).
					Return(&user.User{ID: "current-user"}, nil)
				mocks.UserApplication.EXPECT().
					ListFollower(ctx, "current-user", "12345678-1234-1234-123456789012", 100, 0).
					Return([]*user.Follower{follower1, follower2}, 2, nil)
			},
			args: args{
				req: &pb.ListFollowerRequest{
					UserId: "12345678-1234-1234-123456789012",
					Limit:  100,
					Offset: 0,
				},
			},
			want: &test.TestResponse{
				Code:    codes.OK,
				Message: getFollowerListResponse([]*user.Follower{follower1, follower2}, 100, 0, 2),
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
				req: &pb.ListFollowerRequest{},
			},
			want: &test.TestResponse{
				Code:    codes.Unauthenticated,
				Message: nil,
			},
		},
		{
			name: "failed: invalid argument",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.UserRequestValidation.EXPECT().
					ListFollower(gomock.Any()).
					Return(exception.InvalidRequestValidation.New(test.ErrMock))
				mocks.UserApplication.EXPECT().
					Authentication(ctx).
					Return(&user.User{ID: "current-user"}, nil)
			},
			args: args{
				req: &pb.ListFollowerRequest{},
			},
			want: &test.TestResponse{
				Code:    codes.InvalidArgument,
				Message: nil,
			},
		},
		{
			name: "failed: internel error",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.UserRequestValidation.EXPECT().
					ListFollower(gomock.Any()).
					Return(nil)
				mocks.UserApplication.EXPECT().
					Authentication(ctx).
					Return(&user.User{ID: "current-user"}, nil)
				mocks.UserApplication.EXPECT().
					ListFollower(ctx, "current-user", "12345678-1234-1234-123456789012", 100, 0).
					Return(nil, 0, exception.ErrorInDatastore.New(test.ErrMock))
			},
			args: args{
				req: &pb.ListFollowerRequest{
					UserId: "12345678-1234-1234-123456789012",
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
			target := NewUserServer(mocks.UserRequestValidation, mocks.UserApplication)

			res, err := target.ListFollower(ctx, tt.args.req)
			test.TestGRPC(t, tt.want, res, err)
		})
	}
}

func TestUserServer_MultiGetUser(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	user1 := testUser("user01")

	type args struct {
		req *pb.MultiGetUserRequest
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
				mocks.UserRequestValidation.EXPECT().
					MultiGetUser(gomock.Any()).
					Return(nil)
				mocks.UserApplication.EXPECT().
					MultiGet(ctx, []string{"user01"}).
					Return([]*user.User{user1}, nil)
			},
			args: args{
				req: &pb.MultiGetUserRequest{
					UserIds: []string{"user01"},
				},
			},
			want: &test.TestResponse{
				Code:    codes.OK,
				Message: getUserListResponse([]*user.User{user1}, 1, 0, 1),
			},
		},
		{
			name: "failed: invalid argument",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.UserRequestValidation.EXPECT().
					MultiGetUser(gomock.Any()).
					Return(exception.InvalidRequestValidation.New(test.ErrMock))
			},
			args: args{
				req: &pb.MultiGetUserRequest{},
			},
			want: &test.TestResponse{
				Code:    codes.InvalidArgument,
				Message: nil,
			},
		},
		{
			name: "failed: internel error",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.UserRequestValidation.EXPECT().
					MultiGetUser(gomock.Any()).
					Return(nil)
				mocks.UserApplication.EXPECT().
					MultiGet(ctx, []string{"user01"}).
					Return(nil, exception.ErrorInDatastore.New(test.ErrMock))
			},
			args: args{
				req: &pb.MultiGetUserRequest{
					UserIds: []string{"user01"},
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
			target := NewUserServer(mocks.UserRequestValidation, mocks.UserApplication)

			res, err := target.MultiGetUser(ctx, tt.args.req)
			test.TestGRPC(t, tt.want, res, err)
		})
	}
}

func TestUserServer_GetUser(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	user1 := testUser("user01")

	type args struct {
		req *pb.GetUserRequest
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
				mocks.UserRequestValidation.EXPECT().
					GetUser(gomock.Any()).
					Return(nil)
				mocks.UserApplication.EXPECT().
					Get(ctx, "user01").
					Return(user1, nil)
			},
			args: args{
				req: &pb.GetUserRequest{
					UserId: "user01",
				},
			},
			want: &test.TestResponse{
				Code:    codes.OK,
				Message: getUserResponse(user1),
			},
		},
		{
			name: "failed: invalid argument",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.UserRequestValidation.EXPECT().
					GetUser(gomock.Any()).
					Return(exception.InvalidRequestValidation.New(test.ErrMock))
			},
			args: args{
				req: &pb.GetUserRequest{},
			},
			want: &test.TestResponse{
				Code:    codes.InvalidArgument,
				Message: nil,
			},
		},
		{
			name: "failed: internel error",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.UserRequestValidation.EXPECT().
					GetUser(gomock.Any()).
					Return(nil)
				mocks.UserApplication.EXPECT().
					Get(ctx, "user01").
					Return(nil, exception.ErrorInDatastore.New(test.ErrMock))
			},
			args: args{
				req: &pb.GetUserRequest{
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
			target := NewUserServer(mocks.UserRequestValidation, mocks.UserApplication)

			res, err := target.GetUser(ctx, tt.args.req)
			test.TestGRPC(t, tt.want, res, err)
		})
	}
}

func TestUserServer_GetUserProfile(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	user1 := testUser("user01")

	type args struct {
		req *pb.GetUserProfileRequest
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
				mocks.UserRequestValidation.EXPECT().
					GetUserProfile(gomock.Any()).
					Return(nil)
				mocks.UserApplication.EXPECT().
					Authentication(ctx).
					Return(&user.User{ID: "current-user"}, nil)
				mocks.UserApplication.EXPECT().
					GetUserProfile(ctx, "current-user", "user01").
					Return(user1, nil)
			},
			args: args{
				req: &pb.GetUserProfileRequest{
					UserId: "user01",
				},
			},
			want: &test.TestResponse{
				Code:    codes.OK,
				Message: getUserProfileResponse(user1),
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
				req: &pb.GetUserProfileRequest{},
			},
			want: &test.TestResponse{
				Code:    codes.Unauthenticated,
				Message: nil,
			},
		},
		{
			name: "failed: invalid argument",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.UserApplication.EXPECT().
					Authentication(ctx).
					Return(&user.User{ID: "current-user"}, nil)
				mocks.UserRequestValidation.EXPECT().
					GetUserProfile(gomock.Any()).
					Return(exception.InvalidRequestValidation.New(test.ErrMock))
			},
			args: args{
				req: &pb.GetUserProfileRequest{},
			},
			want: &test.TestResponse{
				Code:    codes.InvalidArgument,
				Message: nil,
			},
		},
		{
			name: "failed: internel error",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.UserRequestValidation.EXPECT().
					GetUserProfile(gomock.Any()).
					Return(nil)
				mocks.UserApplication.EXPECT().
					Authentication(ctx).
					Return(&user.User{ID: "current-user"}, nil)
				mocks.UserApplication.EXPECT().
					GetUserProfile(ctx, "current-user", "user01").
					Return(nil, exception.ErrorInDatastore.New(test.ErrMock))
			},
			args: args{
				req: &pb.GetUserProfileRequest{
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
			target := NewUserServer(mocks.UserRequestValidation, mocks.UserApplication)

			res, err := target.GetUserProfile(ctx, tt.args.req)
			test.TestGRPC(t, tt.want, res, err)
		})
	}
}

func TestUserServer_Follow(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	user1 := testUser("user01")

	type args struct {
		req *pb.FollowRequest
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
				mocks.UserRequestValidation.EXPECT().
					Follow(gomock.Any()).
					Return(nil)
				mocks.UserApplication.EXPECT().
					Follow(ctx, "12345678-1234-1234-123456789012", "user01").
					Return(user1, nil)
			},
			args: args{
				req: &pb.FollowRequest{
					UserId:     "12345678-1234-1234-123456789012",
					FollowerId: "user01",
				},
			},
			want: &test.TestResponse{
				Code:    codes.OK,
				Message: getUserProfileResponse(user1),
			},
		},
		{
			name: "failed: invalid argument",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.UserRequestValidation.EXPECT().
					Follow(gomock.Any()).
					Return(exception.InvalidRequestValidation.New(test.ErrMock))
			},
			args: args{
				req: &pb.FollowRequest{},
			},
			want: &test.TestResponse{
				Code:    codes.InvalidArgument,
				Message: nil,
			},
		},
		{
			name: "failed: internel error",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.UserRequestValidation.EXPECT().
					Follow(gomock.Any()).
					Return(nil)
				mocks.UserApplication.EXPECT().
					Follow(ctx, "12345678-1234-1234-123456789012", "user01").
					Return(nil, exception.ErrorInDatastore.New(test.ErrMock))
			},
			args: args{
				req: &pb.FollowRequest{
					UserId:     "12345678-1234-1234-123456789012",
					FollowerId: "user01",
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
			target := NewUserServer(mocks.UserRequestValidation, mocks.UserApplication)

			res, err := target.Follow(ctx, tt.args.req)
			test.TestGRPC(t, tt.want, res, err)
		})
	}
}

func TestUserServer_Unfollow(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	user1 := testUser("user01")

	type args struct {
		req *pb.UnfollowRequest
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
				mocks.UserRequestValidation.EXPECT().
					Unfollow(gomock.Any()).
					Return(nil)
				mocks.UserApplication.EXPECT().
					Unfollow(ctx, "12345678-1234-1234-123456789012", "user01").
					Return(user1, nil)
			},
			args: args{
				req: &pb.UnfollowRequest{
					UserId:     "12345678-1234-1234-123456789012",
					FollowerId: "user01",
				},
			},
			want: &test.TestResponse{
				Code:    codes.OK,
				Message: getUserProfileResponse(user1),
			},
		},
		{
			name: "failed: invalid argument",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.UserRequestValidation.EXPECT().
					Unfollow(gomock.Any()).
					Return(exception.InvalidRequestValidation.New(test.ErrMock))
			},
			args: args{
				req: &pb.UnfollowRequest{},
			},
			want: &test.TestResponse{
				Code:    codes.InvalidArgument,
				Message: nil,
			},
		},
		{
			name: "failed: internel error",
			setup: func(ctx context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.UserRequestValidation.EXPECT().
					Unfollow(gomock.Any()).
					Return(nil)
				mocks.UserApplication.EXPECT().
					Unfollow(ctx, "12345678-1234-1234-123456789012", "user01").
					Return(nil, exception.ErrorInDatastore.New(test.ErrMock))
			},
			args: args{
				req: &pb.UnfollowRequest{
					UserId:     "12345678-1234-1234-123456789012",
					FollowerId: "user01",
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
			target := NewUserServer(mocks.UserRequestValidation, mocks.UserApplication)

			res, err := target.Unfollow(ctx, tt.args.req)
			test.TestGRPC(t, tt.want, res, err)
		})
	}
}

func testUser(id string) *user.User {
	return &user.User{
		ID:               id,
		Username:         "テストユーザー",
		Gender:           user.MaleGender,
		Email:            "test-user@calmato.jp",
		PhoneNumber:      "000-0000-0000",
		Role:             user.UserRole,
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
		InstanceID:       "ExponentPushToken[!Qaz2wsx3edc4rfv5tgb6y]",
		CreatedAt:        test.TimeMock,
		UpdatedAt:        test.TimeMock,
	}
}

func testFollow(id string) *user.Follow {
	return &user.Follow{
		FollowID:         id,
		Username:         "テストユーザー",
		ThumbnailURL:     "https://go.dev/images/gophers/ladder.svg",
		SelfIntroduction: "テストコードです。",
		IsFollowing:      true,
		IsFollowed:       false,
		FollowCount:      100,
		FollowerCount:    64,
	}
}

func testFollower(id string) *user.Follower {
	return &user.Follower{
		FollowerID:       id,
		Username:         "テストユーザー",
		ThumbnailURL:     "https://go.dev/images/gophers/ladder.svg",
		SelfIntroduction: "テストコードです。",
		IsFollowing:      false,
		IsFollowed:       true,
		FollowCount:      64,
		FollowerCount:    100,
	}
}
