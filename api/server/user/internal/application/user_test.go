package application

import (
	"context"
	"testing"
	"time"

	"github.com/calmato/gran-book/api/server/user/internal/domain/exception"
	"github.com/calmato/gran-book/api/server/user/internal/domain/user"
	"github.com/calmato/gran-book/api/server/user/pkg/database"
	"github.com/calmato/gran-book/api/server/user/pkg/test"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"golang.org/x/xerrors"
)

func TestUserApplication_Authentication(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	user1 := testUser("user01")

	type want struct {
		user *user.User
		err  error
	}
	testCases := []struct {
		name  string
		setup func(context.Context, *testing.T, *test.Mocks)
		want  want
	}{
		{
			name: "success: already exist",
			setup: func(context context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.UserRepository.EXPECT().
					Authentication(ctx).
					Return("user01", nil)
				mocks.UserRepository.EXPECT().
					Get(ctx, "user01").
					Return(user1, nil)
			},
			want: want{
				user: user1,
				err:  nil,
			},
		},
		{
			name: "success: not found",
			setup: func(context context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.UserRepository.EXPECT().
					Authentication(ctx).
					Return("user01", nil)
				mocks.UserRepository.EXPECT().
					Get(ctx, "user01").
					Return(nil, exception.NotFound.New(test.ErrMock))
				mocks.UserRepository.EXPECT().
					CreateWithOAuth(ctx, gomock.Any()).
					Return(nil)
			},
			want: want{
				user: &user.User{
					ID:     "user01",
					Gender: user.UnkownGender,
					Role:   user.UserRole,
				},
			},
		},
		{
			name: "failed: unauthorized",
			setup: func(context context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.UserRepository.EXPECT().
					Authentication(ctx).
					Return("", exception.Unauthorized.New(test.ErrMock))
			},
			want: want{
				user: nil,
				err:  exception.Unauthorized.New(test.ErrMock),
			},
		},
		{
			name: "failed: internal error",
			setup: func(context context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.UserRepository.EXPECT().
					Authentication(ctx).
					Return("user01", nil)
				mocks.UserRepository.EXPECT().
					Get(ctx, "user01").
					Return(nil, exception.NotFound.New(test.ErrMock))
				mocks.UserRepository.EXPECT().
					CreateWithOAuth(ctx, gomock.Any()).
					Return(exception.ErrorInDatastore.New(test.ErrMock))
			},
			want: want{
				user: nil,
				err:  exception.ErrorInDatastore.New(test.ErrMock),
			},
		},
	}

	for _, tt := range testCases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			mocks := test.NewMocks(ctrl)
			tt.setup(ctx, t, mocks)
			target := NewUserApplication(
				mocks.UserDomainValidation,
				mocks.UserRepository,
				mocks.UserUploader,
			)

			u, err := target.Authentication(ctx)
			if tt.want.err != nil {
				require.Equal(t, tt.want.err.Error(), err.Error())
				require.Equal(t, tt.want.user, u)
				return
			}
			require.NoError(t, err)
			require.Equal(t, tt.want.user.ID, u.ID)
			require.Equal(t, tt.want.user.Gender, u.Gender)
			require.Equal(t, tt.want.user.Role, u.Role)
			require.NotZero(t, u.CreatedAt)
			require.NotZero(t, u.UpdatedAt)
		})
	}
}

func TestUserApplication_Authorization(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	user1 := testUser("user01")
	admin1 := testUser("admin01")
	admin1.Role = user.AdminRole

	type want struct {
		role int
		err  error
	}
	testCases := []struct {
		name  string
		setup func(context.Context, *testing.T, *test.Mocks)
		want  want
	}{
		{
			name: "success",
			setup: func(context context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.UserRepository.EXPECT().
					Authentication(ctx).
					Return("admin01", nil)
				mocks.UserRepository.EXPECT().
					Get(ctx, "admin01").
					Return(admin1, nil)
			},
			want: want{
				role: user.AdminRole,
				err:  nil,
			},
		},
		{
			name: "failed: unauthorized",
			setup: func(context context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.UserRepository.EXPECT().
					Authentication(ctx).
					Return("", exception.Unauthorized.New(test.ErrMock))
			},
			want: want{
				role: user.UserRole,
				err:  exception.Unauthorized.New(test.ErrMock),
			},
		},
		{
			name: "failed: not found",
			setup: func(context context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.UserRepository.EXPECT().
					Authentication(ctx).
					Return("user01", nil)
				mocks.UserRepository.EXPECT().
					Get(ctx, "user01").
					Return(nil, exception.NotFound.New(test.ErrMock))
			},
			want: want{
				role: user.UserRole,
				err:  exception.Forbidden.New(exception.NotFound.New(test.ErrMock)),
			},
		},
		{
			name: "failed: forbidden",
			setup: func(context context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.UserRepository.EXPECT().
					Authentication(ctx).
					Return("user01", nil)
				mocks.UserRepository.EXPECT().
					Get(ctx, "user01").
					Return(user1, nil)
			},
			want: want{
				role: user.UserRole,
				err:  exception.Forbidden.New(xerrors.New("This account doesn't have administrator privileges")),
			},
		},
	}

	for _, tt := range testCases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			mocks := test.NewMocks(ctrl)
			tt.setup(ctx, t, mocks)
			target := NewUserApplication(
				mocks.UserDomainValidation,
				mocks.UserRepository,
				mocks.UserUploader,
			)

			role, err := target.Authorization(ctx)
			if tt.want.err != nil {
				require.Equal(t, tt.want.err.Error(), err.Error())
				require.Equal(t, tt.want.role, role)
				return
			}
			require.NoError(t, err)
			require.Equal(t, tt.want.role, role)
		})
	}
}

func TestUserApplication_List(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	user1 := testUser("user01")
	user2 := testUser("user02")

	type args struct {
		query *database.ListQuery
	}
	type want struct {
		users []*user.User
		total int
		err   error
	}
	testCases := []struct {
		name  string
		setup func(context.Context, *testing.T, *test.Mocks)
		args  args
		want  want
	}{
		{
			name: "success",
			setup: func(context context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.UserRepository.EXPECT().
					List(ctx, gomock.Any()).
					Return([]*user.User{user1, user2}, nil)
				mocks.UserRepository.EXPECT().
					Count(ctx, gomock.Any()).
					Return(2, nil)
			},
			args: args{
				query: &database.ListQuery{},
			},
			want: want{
				users: []*user.User{user1, user2},
				total: 2,
				err:   nil,
			},
		},
		{
			name: "failed: internal error in list",
			setup: func(context context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.UserRepository.EXPECT().
					List(ctx, gomock.Any()).
					Return(nil, exception.ErrorInDatastore.New(test.ErrMock))
			},
			args: args{
				query: &database.ListQuery{},
			},
			want: want{
				users: nil,
				total: 0,
				err:   exception.ErrorInDatastore.New(test.ErrMock),
			},
		},
		{
			name: "failed: internal error in count",
			setup: func(context context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.UserRepository.EXPECT().
					List(ctx, gomock.Any()).
					Return([]*user.User{user1, user2}, nil)
				mocks.UserRepository.EXPECT().
					Count(ctx, gomock.Any()).
					Return(0, exception.ErrorInDatastore.New(test.ErrMock))
			},
			args: args{
				query: &database.ListQuery{},
			},
			want: want{
				users: nil,
				total: 0,
				err:   exception.ErrorInDatastore.New(test.ErrMock),
			},
		},
	}

	for _, tt := range testCases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			mocks := test.NewMocks(ctrl)
			tt.setup(ctx, t, mocks)
			target := NewUserApplication(
				mocks.UserDomainValidation,
				mocks.UserRepository,
				mocks.UserUploader,
			)

			us, total, err := target.List(ctx, tt.args.query)
			if tt.want.err != nil {
				require.Equal(t, tt.want.err.Error(), err.Error())
				require.Equal(t, tt.want.users, us)
				require.Equal(t, tt.want.total, total)
				return
			}
			require.NoError(t, err)
			require.Equal(t, tt.want.users, us)
			require.Equal(t, tt.want.total, total)
		})
	}
}

func TestUserApplication_ListAdmin(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	admin1 := testAdmin("admin01")
	admin2 := testAdmin("admin02")

	type args struct {
		query *database.ListQuery
	}
	type want struct {
		users []*user.User
		total int
		err   error
	}
	testCases := []struct {
		name  string
		setup func(context.Context, *testing.T, *test.Mocks)
		args  args
		want  want
	}{
		{
			name: "success",
			setup: func(context context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.UserRepository.EXPECT().
					List(ctx, gomock.Any()).
					Return([]*user.User{admin1, admin2}, nil)
				mocks.UserRepository.EXPECT().
					Count(ctx, gomock.Any()).
					Return(2, nil)
			},
			args: args{
				query: &database.ListQuery{},
			},
			want: want{
				users: []*user.User{admin1, admin2},
				total: 2,
				err:   nil,
			},
		},
		{
			name: "failed: internal error in list",
			setup: func(context context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.UserRepository.EXPECT().
					List(ctx, gomock.Any()).
					Return(nil, exception.ErrorInDatastore.New(test.ErrMock))
			},
			args: args{
				query: &database.ListQuery{},
			},
			want: want{
				users: nil,
				total: 0,
				err:   exception.ErrorInDatastore.New(test.ErrMock),
			},
		},
		{
			name: "failed: internal error in count",
			setup: func(context context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.UserRepository.EXPECT().
					List(ctx, gomock.Any()).
					Return([]*user.User{admin1, admin2}, nil)
				mocks.UserRepository.EXPECT().
					Count(ctx, gomock.Any()).
					Return(0, exception.ErrorInDatastore.New(test.ErrMock))
			},
			args: args{
				query: &database.ListQuery{},
			},
			want: want{
				users: nil,
				total: 0,
				err:   exception.ErrorInDatastore.New(test.ErrMock),
			},
		},
	}

	for _, tt := range testCases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			mocks := test.NewMocks(ctrl)
			tt.setup(ctx, t, mocks)
			target := NewUserApplication(
				mocks.UserDomainValidation,
				mocks.UserRepository,
				mocks.UserUploader,
			)

			us, total, err := target.ListAdmin(ctx, tt.args.query)
			if tt.want.err != nil {
				require.Equal(t, tt.want.err.Error(), err.Error())
				require.Equal(t, tt.want.users, us)
				require.Equal(t, tt.want.total, total)
				return
			}
			require.NoError(t, err)
			require.Equal(t, tt.want.users, us)
			require.Equal(t, tt.want.total, total)
		})
	}
}

func TestUserApplication_ListFollow(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	follow1 := testFollow("follow01")
	follow2 := testFollow("follow02")

	type args struct {
		currentUserID string
		targetUserID  string
		limit         int
		offset        int
	}
	type want struct {
		follows []*user.Follow
		total   int
		err     error
	}
	testCases := []struct {
		name  string
		setup func(context.Context, *testing.T, *test.Mocks)
		args  args
		want  want
	}{
		{
			name: "success",
			setup: func(context context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.UserRepository.EXPECT().
					ListFollow(ctx, gomock.Any()).
					Return([]*user.Follow{follow1, follow2}, nil)
				mocks.UserRepository.EXPECT().
					ListFollowID(ctx, "current-user", []string{"follow01", "follow02"}).
					Return([]string{"follow01"}, nil)
				mocks.UserRepository.EXPECT().
					ListFollowerID(ctx, "current-user", []string{"follow01", "follow02"}).
					Return([]string{"follow01"}, nil)
				mocks.UserRepository.EXPECT().
					CountRelationship(ctx, gomock.Any()).
					Return(2, nil)
			},
			args: args{
				currentUserID: "current-user",
				targetUserID:  "user01",
				limit:         100,
				offset:        0,
			},
			want: want{
				follows: []*user.Follow{follow1, follow2},
				total:   2,
				err:     nil,
			},
		},
		{
			name: "failed: internal error in list follow",
			setup: func(context context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.UserRepository.EXPECT().
					ListFollow(ctx, gomock.Any()).
					Return(nil, exception.ErrorInDatastore.New(test.ErrMock))
			},
			args: args{
				currentUserID: "current-user",
				targetUserID:  "user01",
				limit:         100,
				offset:        0,
			},
			want: want{
				follows: nil,
				total:   0,
				err:     exception.ErrorInDatastore.New(test.ErrMock),
			},
		},
		{
			name: "failed: internal error in list follow id",
			setup: func(context context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.UserRepository.EXPECT().
					ListFollow(ctx, gomock.Any()).
					Return([]*user.Follow{follow1, follow2}, nil)
				mocks.UserRepository.EXPECT().
					ListFollowID(ctx, "current-user", []string{"follow01", "follow02"}).
					Return(nil, exception.ErrorInDatastore.New(test.ErrMock))
			},
			args: args{
				currentUserID: "current-user",
				targetUserID:  "user01",
				limit:         100,
				offset:        0,
			},
			want: want{
				follows: nil,
				total:   0,
				err:     exception.ErrorInDatastore.New(test.ErrMock),
			},
		},
		{
			name: "failed: internal error in list follower id",
			setup: func(context context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.UserRepository.EXPECT().
					ListFollow(ctx, gomock.Any()).
					Return([]*user.Follow{follow1, follow2}, nil)
				mocks.UserRepository.EXPECT().
					ListFollowID(ctx, "current-user", []string{"follow01", "follow02"}).
					Return([]string{"follow01"}, nil)
				mocks.UserRepository.EXPECT().
					ListFollowerID(ctx, "current-user", []string{"follow01", "follow02"}).
					Return(nil, exception.ErrorInDatastore.New(test.ErrMock))
			},
			args: args{
				currentUserID: "current-user",
				targetUserID:  "user01",
				limit:         100,
				offset:        0,
			},
			want: want{
				follows: nil,
				total:   0,
				err:     exception.ErrorInDatastore.New(test.ErrMock),
			},
		},
		{
			name: "failed: internal error in count relationship",
			setup: func(context context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.UserRepository.EXPECT().
					ListFollow(ctx, gomock.Any()).
					Return([]*user.Follow{follow1, follow2}, nil)
				mocks.UserRepository.EXPECT().
					ListFollowID(ctx, "current-user", []string{"follow01", "follow02"}).
					Return([]string{"follow01"}, nil)
				mocks.UserRepository.EXPECT().
					ListFollowerID(ctx, "current-user", []string{"follow01", "follow02"}).
					Return([]string{"follow01"}, nil)
				mocks.UserRepository.EXPECT().
					CountRelationship(ctx, gomock.Any()).
					Return(0, exception.ErrorInDatastore.New(test.ErrMock))
			},
			args: args{
				currentUserID: "current-user",
				targetUserID:  "user01",
				limit:         100,
				offset:        0,
			},
			want: want{
				follows: nil,
				total:   0,
				err:     exception.ErrorInDatastore.New(test.ErrMock),
			},
		},
	}

	for _, tt := range testCases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			mocks := test.NewMocks(ctrl)
			tt.setup(ctx, t, mocks)
			target := NewUserApplication(
				mocks.UserDomainValidation,
				mocks.UserRepository,
				mocks.UserUploader,
			)

			fs, total, err := target.ListFollow(
				ctx, tt.args.currentUserID, tt.args.targetUserID, tt.args.limit, tt.args.offset,
			)
			if tt.want.err != nil {
				require.Equal(t, tt.want.err.Error(), err.Error())
				require.Equal(t, tt.want.follows, fs)
				require.Equal(t, tt.want.total, total)
				return
			}
			require.NoError(t, err)
			require.Equal(t, tt.want.total, total)
			for i := range tt.want.follows {
				require.NotNil(t, fs[i])
				require.Equal(t, tt.want.follows[i].FollowID, fs[i].FollowID)
			}
		})
	}
}

func TestUserApplication_ListFollower(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	follower1 := testFollower("follower01")
	follower2 := testFollower("follower02")

	type args struct {
		currentUserID string
		targetUserID  string
		limit         int
		offset        int
	}
	type want struct {
		followers []*user.Follower
		total     int
		err       error
	}
	testCases := []struct {
		name  string
		setup func(context.Context, *testing.T, *test.Mocks)
		args  args
		want  want
	}{
		{
			name: "success",
			setup: func(context context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.UserRepository.EXPECT().
					ListFollower(ctx, gomock.Any()).
					Return([]*user.Follower{follower1, follower2}, nil)
				mocks.UserRepository.EXPECT().
					ListFollowID(ctx, "current-user", []string{"follower01", "follower02"}).
					Return([]string{"follow01"}, nil)
				mocks.UserRepository.EXPECT().
					ListFollowerID(ctx, "current-user", []string{"follower01", "follower02"}).
					Return([]string{"follow01"}, nil)
				mocks.UserRepository.EXPECT().
					CountRelationship(ctx, gomock.Any()).
					Return(2, nil)
			},
			args: args{
				currentUserID: "current-user",
				targetUserID:  "user01",
				limit:         100,
				offset:        0,
			},
			want: want{
				followers: []*user.Follower{follower1, follower2},
				total:     2,
				err:       nil,
			},
		},
		{
			name: "failed: internal error in list follower",
			setup: func(context context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.UserRepository.EXPECT().
					ListFollower(ctx, gomock.Any()).
					Return(nil, exception.ErrorInDatastore.New(test.ErrMock))
			},
			args: args{
				currentUserID: "current-user",
				targetUserID:  "user01",
				limit:         100,
				offset:        0,
			},
			want: want{
				followers: nil,
				total:     0,
				err:       exception.ErrorInDatastore.New(test.ErrMock),
			},
		},
		{
			name: "failed: internal error in list follow id",
			setup: func(context context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.UserRepository.EXPECT().
					ListFollower(ctx, gomock.Any()).
					Return([]*user.Follower{follower1, follower2}, nil)
				mocks.UserRepository.EXPECT().
					ListFollowID(ctx, "current-user", []string{"follower01", "follower02"}).
					Return(nil, exception.ErrorInDatastore.New(test.ErrMock))
			},
			args: args{
				currentUserID: "current-user",
				targetUserID:  "user01",
				limit:         100,
				offset:        0,
			},
			want: want{
				followers: nil,
				total:     0,
				err:       exception.ErrorInDatastore.New(test.ErrMock),
			},
		},
		{
			name: "failed: internal error in list follower id",
			setup: func(context context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.UserRepository.EXPECT().
					ListFollower(ctx, gomock.Any()).
					Return([]*user.Follower{follower1, follower2}, nil)
				mocks.UserRepository.EXPECT().
					ListFollowID(ctx, "current-user", []string{"follower01", "follower02"}).
					Return([]string{"follow01"}, nil)
				mocks.UserRepository.EXPECT().
					ListFollowerID(ctx, "current-user", []string{"follower01", "follower02"}).
					Return(nil, exception.ErrorInDatastore.New(test.ErrMock))
			},
			args: args{
				currentUserID: "current-user",
				targetUserID:  "user01",
				limit:         100,
				offset:        0,
			},
			want: want{
				followers: nil,
				total:     0,
				err:       exception.ErrorInDatastore.New(test.ErrMock),
			},
		},
		{
			name: "failed: internal error in count relationship",
			setup: func(context context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.UserRepository.EXPECT().
					ListFollower(ctx, gomock.Any()).
					Return([]*user.Follower{follower1, follower2}, nil)
				mocks.UserRepository.EXPECT().
					ListFollowID(ctx, "current-user", []string{"follower01", "follower02"}).
					Return([]string{"follow01"}, nil)
				mocks.UserRepository.EXPECT().
					ListFollowerID(ctx, "current-user", []string{"follower01", "follower02"}).
					Return([]string{"follower01"}, nil)
				mocks.UserRepository.EXPECT().
					CountRelationship(ctx, gomock.Any()).
					Return(0, exception.ErrorInDatastore.New(test.ErrMock))
			},
			args: args{
				currentUserID: "current-user",
				targetUserID:  "user01",
				limit:         100,
				offset:        0,
			},
			want: want{
				followers: nil,
				total:     0,
				err:       exception.ErrorInDatastore.New(test.ErrMock),
			},
		},
	}

	for _, tt := range testCases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			mocks := test.NewMocks(ctrl)
			tt.setup(ctx, t, mocks)
			target := NewUserApplication(
				mocks.UserDomainValidation,
				mocks.UserRepository,
				mocks.UserUploader,
			)

			fs, total, err := target.ListFollower(
				ctx, tt.args.currentUserID, tt.args.targetUserID, tt.args.limit, tt.args.offset,
			)
			if tt.want.err != nil {
				require.Equal(t, tt.want.err.Error(), err.Error())
				require.Equal(t, tt.want.followers, fs)
				require.Equal(t, tt.want.total, total)
				return
			}
			require.NoError(t, err)
			require.Equal(t, tt.want.total, total)
			for i := range tt.want.followers {
				require.NotNil(t, fs[i])
				require.Equal(t, tt.want.followers[i].FollowerID, fs[i].FollowerID)
			}
		})
	}
}

func TestUserApplication_MultiGet(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	user1 := testUser("user01")
	user2 := testUser("user02")

	type args struct {
		userIDs []string
	}
	type want struct {
		users []*user.User
		err   error
	}
	testCases := []struct {
		name  string
		setup func(context.Context, *testing.T, *test.Mocks)
		args  args
		want  want
	}{
		{
			name: "success",
			setup: func(context context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.UserRepository.EXPECT().
					MultiGet(ctx, []string{"user01", "user02"}).
					Return([]*user.User{user1, user2}, nil)
			},
			args: args{
				userIDs: []string{"user01", "user02"},
			},
			want: want{
				users: []*user.User{user1, user2},
				err:   nil,
			},
		},
	}

	for _, tt := range testCases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			mocks := test.NewMocks(ctrl)
			tt.setup(ctx, t, mocks)
			target := NewUserApplication(
				mocks.UserDomainValidation,
				mocks.UserRepository,
				mocks.UserUploader,
			)

			us, err := target.MultiGet(ctx, tt.args.userIDs)
			if tt.want.err != nil {
				require.Equal(t, tt.want.err.Error(), err.Error())
				require.Equal(t, tt.want.users, us)
				return
			}
			require.NoError(t, err)
			require.Equal(t, tt.want.users, us)
		})
	}
}

func TestUserApplication_Get(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	user1 := testUser("user01")

	type args struct {
		userID string
	}
	type want struct {
		user *user.User
		err  error
	}
	testCases := []struct {
		name  string
		setup func(context.Context, *testing.T, *test.Mocks)
		args  args
		want  want
	}{
		{
			name: "success",
			setup: func(context context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.UserRepository.EXPECT().
					Get(ctx, "user01").
					Return(user1, nil)
			},
			args: args{
				userID: "user01",
			},
			want: want{
				user: user1,
				err:  nil,
			},
		},
	}

	for _, tt := range testCases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			mocks := test.NewMocks(ctrl)
			tt.setup(ctx, t, mocks)
			target := NewUserApplication(
				mocks.UserDomainValidation,
				mocks.UserRepository,
				mocks.UserUploader,
			)

			u, err := target.Get(ctx, tt.args.userID)
			if tt.want.err != nil {
				require.Equal(t, tt.want.err.Error(), err.Error())
				require.Equal(t, tt.want.user, u)
				return
			}
			require.NoError(t, err)
			require.Equal(t, tt.want.user, u)
		})
	}
}

func TestUserApplication_GetAdmin(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	admin1 := testAdmin("admin01")

	type args struct {
		userID string
	}
	type want struct {
		user *user.User
		err  error
	}
	testCases := []struct {
		name  string
		setup func(context.Context, *testing.T, *test.Mocks)
		args  args
		want  want
	}{
		{
			name: "success",
			setup: func(context context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.UserRepository.EXPECT().
					GetAdmin(ctx, "admin01").
					Return(admin1, nil)
			},
			args: args{
				userID: "admin01",
			},
			want: want{
				user: admin1,
				err:  nil,
			},
		},
	}

	for _, tt := range testCases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			mocks := test.NewMocks(ctrl)
			tt.setup(ctx, t, mocks)
			target := NewUserApplication(
				mocks.UserDomainValidation,
				mocks.UserRepository,
				mocks.UserUploader,
			)

			u, err := target.GetAdmin(ctx, tt.args.userID)
			if tt.want.err != nil {
				require.Equal(t, tt.want.err.Error(), err.Error())
				require.Equal(t, tt.want.user, u)
				return
			}
			require.NoError(t, err)
			require.Equal(t, tt.want.user, u)
		})
	}
}

func TestUserApplication_GetUserProfile(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	user1 := testUser("user01")

	type args struct {
		userID   string
		targetID string
	}
	type want struct {
		user          *user.User
		isFollowing   bool
		isFollowed    bool
		followCount   int
		followerCount int
		err           error
	}
	testCases := []struct {
		name  string
		setup func(context.Context, *testing.T, *test.Mocks)
		args  args
		want  want
	}{
		{
			name: "success",
			setup: func(context context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.UserRepository.EXPECT().
					Get(ctx, "user01").
					Return(user1, nil)
				mocks.UserRepository.EXPECT().
					ListFollowID(ctx, "user01").
					Return([]string{"user02", "current-user"}, nil)
				mocks.UserRepository.EXPECT().
					ListFollowerID(ctx, "user01").
					Return([]string{"user02"}, nil)
			},
			args: args{
				userID:   "current-user",
				targetID: "user01",
			},
			want: want{
				user:          user1,
				isFollowing:   true,
				isFollowed:    false,
				followCount:   2,
				followerCount: 1,
				err:           nil,
			},
		},
		{
			name: "failed: not found",
			setup: func(context context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.UserRepository.EXPECT().
					Get(ctx, "user01").
					Return(nil, exception.NotFound.New(test.ErrMock))
			},
			args: args{
				userID:   "current-user",
				targetID: "user01",
			},
			want: want{
				user:          nil,
				isFollowing:   false,
				isFollowed:    false,
				followCount:   0,
				followerCount: 0,
				err:           exception.NotFound.New(test.ErrMock),
			},
		},
		{
			name: "failed: internal error list follow id",
			setup: func(context context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.UserRepository.EXPECT().
					Get(ctx, "user01").
					Return(user1, nil)
				mocks.UserRepository.EXPECT().
					ListFollowID(ctx, "user01").
					Return(nil, exception.ErrorInDatastore.New(test.ErrMock))
			},
			args: args{
				userID:   "current-user",
				targetID: "user01",
			},
			want: want{
				user:          nil,
				isFollowing:   false,
				isFollowed:    false,
				followCount:   0,
				followerCount: 0,
				err:           exception.ErrorInDatastore.New(test.ErrMock),
			},
		},
		{
			name: "failed: internal error list follower id",
			setup: func(context context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.UserRepository.EXPECT().
					Get(ctx, "user01").
					Return(user1, nil)
				mocks.UserRepository.EXPECT().
					ListFollowID(ctx, "user01").
					Return([]string{"user02", "current-user"}, nil)
				mocks.UserRepository.EXPECT().
					ListFollowerID(ctx, "user01").
					Return(nil, exception.ErrorInDatastore.New(test.ErrMock))
			},
			args: args{
				userID:   "current-user",
				targetID: "user01",
			},
			want: want{
				user:          nil,
				isFollowing:   false,
				isFollowed:    false,
				followCount:   0,
				followerCount: 0,
				err:           exception.ErrorInDatastore.New(test.ErrMock),
			},
		},
	}

	for _, tt := range testCases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			mocks := test.NewMocks(ctrl)
			tt.setup(ctx, t, mocks)
			target := NewUserApplication(
				mocks.UserDomainValidation,
				mocks.UserRepository,
				mocks.UserUploader,
			)

			u,
				isFollowing,
				isFollowed,
				followCount,
				followerCount,
				err := target.GetUserProfile(ctx, tt.args.userID, tt.args.targetID)
			if tt.want.err != nil {
				require.Equal(t, tt.want.err.Error(), err.Error())
				require.Equal(t, tt.want.user, u)
				require.Equal(t, tt.want.isFollowing, isFollowing)
				require.Equal(t, tt.want.isFollowed, isFollowed)
				require.Equal(t, tt.want.followCount, followCount)
				require.Equal(t, tt.want.followerCount, followerCount)
				return
			}
			require.NoError(t, err)
			require.Equal(t, tt.want.user, u)
			require.Equal(t, tt.want.isFollowing, isFollowing)
			require.Equal(t, tt.want.isFollowed, isFollowed)
			require.Equal(t, tt.want.followCount, followCount)
			require.Equal(t, tt.want.followerCount, followerCount)
		})
	}
}

func TestUserApplication_Create(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type args struct {
		user *user.User
	}
	type want struct {
		err error
	}
	testCases := []struct {
		name  string
		setup func(context.Context, *testing.T, *test.Mocks)
		args  args
		want  want
	}{
		{
			name: "success",
			setup: func(context context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.UserDomainValidation.EXPECT().
					User(ctx, gomock.Any()).
					Return(nil)
				mocks.UserRepository.EXPECT().
					Create(ctx, gomock.Any()).
					Return(nil)
			},
			args: args{
				user: &user.User{
					Username: "test-user",
					Email:    "test-user@calmato.jp",
				},
			},
			want: want{
				err: nil,
			},
		},
		{
			name: "failed: invalid domain validation",
			setup: func(context context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.UserDomainValidation.EXPECT().
					User(ctx, gomock.Any()).
					Return(exception.InvalidDomainValidation.New(test.ErrMock))
			},
			args: args{
				user: &user.User{
					Username: "test-user",
					Email:    "test-user@calmato.jp",
				},
			},
			want: want{
				err: exception.InvalidDomainValidation.New(test.ErrMock),
			},
		},
	}

	for _, tt := range testCases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			mocks := test.NewMocks(ctrl)
			tt.setup(ctx, t, mocks)
			target := NewUserApplication(
				mocks.UserDomainValidation,
				mocks.UserRepository,
				mocks.UserUploader,
			)

			err := target.Create(ctx, tt.args.user)
			if tt.want.err != nil {
				require.Equal(t, tt.want.err.Error(), err.Error())
				return
			}
			require.NoError(t, err)
			require.NotZero(t, tt.args.user.CreatedAt)
			require.NotZero(t, tt.args.user.UpdatedAt)
		})
	}
}

func TestUserApplication_Update(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type args struct {
		user *user.User
	}
	type want struct {
		err error
	}
	testCases := []struct {
		name  string
		setup func(context.Context, *testing.T, *test.Mocks)
		args  args
		want  want
	}{
		{
			name: "success",
			setup: func(context context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.UserDomainValidation.EXPECT().
					User(ctx, gomock.Any()).
					Return(nil)
				mocks.UserRepository.EXPECT().
					Update(ctx, gomock.Any()).
					Return(nil)
			},
			args: args{
				user: &user.User{
					ID:       "test01",
					Username: "test-user",
					Email:    "test-user@calmato.jp",
				},
			},
			want: want{
				err: nil,
			},
		},
		{
			name: "failed: invalid domain validation",
			setup: func(context context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.UserDomainValidation.EXPECT().
					User(ctx, gomock.Any()).
					Return(exception.InvalidDomainValidation.New(test.ErrMock))
			},
			args: args{
				user: &user.User{
					ID:       "test01",
					Username: "test-user",
					Email:    "test-user@calmato.jp",
				},
			},
			want: want{
				err: exception.InvalidDomainValidation.New(test.ErrMock),
			},
		},
	}

	for _, tt := range testCases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			mocks := test.NewMocks(ctrl)
			tt.setup(ctx, t, mocks)
			target := NewUserApplication(
				mocks.UserDomainValidation,
				mocks.UserRepository,
				mocks.UserUploader,
			)

			err := target.Update(ctx, tt.args.user)
			if tt.want.err != nil {
				require.Equal(t, tt.want.err.Error(), err.Error())
				return
			}
			require.NoError(t, err)
			require.NotZero(t, tt.args.user.UpdatedAt)
		})
	}
}

func TestUserApplication_UpdatePassword(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type args struct {
		user *user.User
	}
	type want struct {
		err error
	}
	testCases := []struct {
		name  string
		setup func(context.Context, *testing.T, *test.Mocks)
		args  args
		want  want
	}{
		{
			name: "success",
			setup: func(context context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.UserDomainValidation.EXPECT().
					User(ctx, gomock.Any()).
					Return(nil)
				mocks.UserRepository.EXPECT().
					UpdatePassword(ctx, "user01", "12345678").
					Return(nil)
			},
			args: args{
				user: &user.User{
					ID:       "user01",
					Username: "test-user",
					Email:    "test-user@calmato.jp",
					Password: "12345678",
				},
			},
			want: want{
				err: nil,
			},
		},
		{
			name: "failed: invalid domain validation",
			setup: func(context context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.UserDomainValidation.EXPECT().
					User(ctx, gomock.Any()).
					Return(exception.InvalidDomainValidation.New(test.ErrMock))
			},
			args: args{
				user: &user.User{
					ID:       "user01",
					Username: "test-user",
					Email:    "test-user@calmato.jp",
					Password: "12345678",
				},
			},
			want: want{
				err: exception.InvalidDomainValidation.New(test.ErrMock),
			},
		},
	}

	for _, tt := range testCases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			mocks := test.NewMocks(ctrl)
			tt.setup(ctx, t, mocks)
			target := NewUserApplication(
				mocks.UserDomainValidation,
				mocks.UserRepository,
				mocks.UserUploader,
			)

			err := target.UpdatePassword(ctx, tt.args.user)
			if tt.want.err != nil {
				require.Equal(t, tt.want.err.Error(), err.Error())
				return
			}
			require.NoError(t, err)
		})
	}
}

func TestUserApplication_Delete(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type args struct {
		user *user.User
	}
	type want struct {
		err error
	}
	testCases := []struct {
		name  string
		setup func(context.Context, *testing.T, *test.Mocks)
		args  args
		want  want
	}{
		{
			name: "success",
			setup: func(context context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.UserRepository.EXPECT().
					Delete(ctx, gomock.Any()).
					Return(nil)
			},
			args: args{
				user: &user.User{
					ID: "test01",
				},
			},
			want: want{
				err: nil,
			},
		},
	}

	for _, tt := range testCases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			mocks := test.NewMocks(ctrl)
			tt.setup(ctx, t, mocks)
			target := NewUserApplication(
				mocks.UserDomainValidation,
				mocks.UserRepository,
				mocks.UserUploader,
			)

			err := target.Delete(ctx, tt.args.user)
			if tt.want.err != nil {
				require.Equal(t, tt.want.err.Error(), err.Error())
				return
			}
			require.NoError(t, err)
		})
	}
}

func TestUserApplication_DeleteAdmin(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type args struct {
		user *user.User
	}
	type want struct {
		err error
	}
	testCases := []struct {
		name  string
		setup func(context.Context, *testing.T, *test.Mocks)
		args  args
		want  want
	}{
		{
			name: "success",
			setup: func(context context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.UserDomainValidation.EXPECT().
					User(ctx, gomock.Any()).
					Return(nil)
				mocks.UserRepository.EXPECT().
					Update(ctx, gomock.Any()).
					Return(nil)
			},
			args: args{
				user: &user.User{
					ID:       "admin01",
					Username: "test-user",
					Email:    "test-user@calmato.jp",
					Role:     user.AdminRole,
				},
			},
			want: want{
				err: nil,
			},
		},
		{
			name: "failed: invalid domain validation",
			setup: func(context context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.UserDomainValidation.EXPECT().
					User(ctx, gomock.Any()).
					Return(exception.InvalidDomainValidation.New(test.ErrMock))
			},
			args: args{
				user: &user.User{
					ID:       "admin01",
					Username: "test-user",
					Email:    "test-user@calmato.jp",
					Role:     user.AdminRole,
				},
			},
			want: want{
				err: exception.InvalidDomainValidation.New(test.ErrMock),
			},
		},
	}

	for _, tt := range testCases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			mocks := test.NewMocks(ctrl)
			tt.setup(ctx, t, mocks)
			target := NewUserApplication(
				mocks.UserDomainValidation,
				mocks.UserRepository,
				mocks.UserUploader,
			)

			err := target.DeleteAdmin(ctx, tt.args.user)
			if tt.want.err != nil {
				require.Equal(t, tt.want.err.Error(), err.Error())
				return
			}
			require.NoError(t, err)
			require.Equal(t, user.UserRole, tt.args.user.Role)
			require.NotZero(t, tt.args.user.UpdatedAt)
		})
	}
}

func TestUserApplication_Follow(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type args struct {
		userID     string
		followerID string
	}
	type want struct {
		user          *user.User
		isFollowing   bool
		isFollowed    bool
		followCount   int
		followerCount int
		err           error
	}
	testCases := []struct {
		name  string
		setup func(context.Context, *testing.T, *test.Mocks)
		args  args
		want  want
	}{
		{
			name: "success",
			setup: func(context context.Context, t *testing.T, mocks *test.Mocks) {
				u := &user.User{ID: "user01"}
				mocks.UserDomainValidation.EXPECT().
					Relationship(ctx, gomock.Any()).
					Return(nil)
				mocks.UserRepository.EXPECT().
					Get(ctx, "user01").
					Return(u, nil)
				mocks.UserRepository.EXPECT().
					CreateRelationship(ctx, gomock.Any()).
					Return(nil)
				mocks.UserRepository.EXPECT().
					ListFollowID(ctx, "user01").
					Return([]string{"current-user"}, nil)
				mocks.UserRepository.EXPECT().
					ListFollowerID(ctx, "user01").
					Return([]string{}, nil)
			},
			args: args{
				userID:     "current-user",
				followerID: "user01",
			},
			want: want{
				user: &user.User{
					ID: "user01",
				},
				isFollowing:   true,
				isFollowed:    false,
				followCount:   1,
				followerCount: 0,
				err:           nil,
			},
		},
		{
			name: "failed: not found",
			setup: func(context context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.UserRepository.EXPECT().
					Get(ctx, "user01").
					Return(nil, exception.NotFound.New(test.ErrMock))
			},
			args: args{
				userID:     "current-user",
				followerID: "user01",
			},
			want: want{
				user:          nil,
				isFollowing:   false,
				isFollowed:    false,
				followCount:   0,
				followerCount: 0,
				err:           exception.NotFound.New(test.ErrMock),
			},
		},
		{
			name: "failed: invalid domain validation",
			setup: func(context context.Context, t *testing.T, mocks *test.Mocks) {
				u := &user.User{ID: "user01"}
				mocks.UserDomainValidation.EXPECT().
					Relationship(ctx, gomock.Any()).
					Return(exception.InvalidDomainValidation.New(test.ErrMock))
				mocks.UserRepository.EXPECT().
					Get(ctx, "user01").
					Return(u, nil)
			},
			args: args{
				userID:     "current-user",
				followerID: "user01",
			},
			want: want{
				user:          nil,
				isFollowing:   false,
				isFollowed:    false,
				followCount:   0,
				followerCount: 0,
				err:           exception.InvalidDomainValidation.New(test.ErrMock),
			},
		},
		{
			name: "failed: internal error",
			setup: func(context context.Context, t *testing.T, mocks *test.Mocks) {
				u := &user.User{ID: "user01"}
				mocks.UserDomainValidation.EXPECT().
					Relationship(ctx, gomock.Any()).
					Return(nil)
				mocks.UserRepository.EXPECT().
					Get(ctx, "user01").
					Return(u, nil)
				mocks.UserRepository.EXPECT().
					CreateRelationship(ctx, gomock.Any()).
					Return(exception.ErrorInDatastore.New(test.ErrMock))
			},
			args: args{
				userID:     "current-user",
				followerID: "user01",
			},
			want: want{
				user:          nil,
				isFollowing:   false,
				isFollowed:    false,
				followCount:   0,
				followerCount: 0,
				err:           exception.ErrorInDatastore.New(test.ErrMock),
			},
		},
	}

	for _, tt := range testCases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			mocks := test.NewMocks(ctrl)
			tt.setup(ctx, t, mocks)
			target := NewUserApplication(
				mocks.UserDomainValidation,
				mocks.UserRepository,
				mocks.UserUploader,
			)

			u,
				isFollowing,
				isFollowed,
				followCount,
				followerCount,
				err := target.Follow(ctx, tt.args.userID, tt.args.followerID)
			if tt.want.err != nil {
				require.Equal(t, tt.want.err.Error(), err.Error())
				return
			}
			require.NoError(t, err)
			require.Equal(t, tt.want.user, u)
			require.Equal(t, tt.want.isFollowing, isFollowing)
			require.Equal(t, tt.want.isFollowed, isFollowed)
			require.Equal(t, tt.want.followCount, followCount)
			require.Equal(t, tt.want.followerCount, followerCount)
		})
	}
}

func TestUserApplication_Unfollow(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type args struct {
		userID     string
		followerID string
	}
	type want struct {
		user          *user.User
		isFollowing   bool
		isFollowed    bool
		followCount   int
		followerCount int
		err           error
	}
	testCases := []struct {
		name  string
		setup func(context.Context, *testing.T, *test.Mocks)
		args  args
		want  want
	}{
		{
			name: "success",
			setup: func(context context.Context, t *testing.T, mocks *test.Mocks) {
				u := &user.User{ID: "user01"}
				mocks.UserDomainValidation.EXPECT().
					Relationship(ctx, gomock.Any()).
					Return(nil)
				mocks.UserRepository.EXPECT().
					Get(ctx, "user01").
					Return(u, nil)
				mocks.UserRepository.EXPECT().
					GetRelationshipIDByUserID(ctx, "current-user", "user01").
					Return(1, nil)
				mocks.UserRepository.EXPECT().
					DeleteRelationship(ctx, 1).
					Return(nil)
				mocks.UserRepository.EXPECT().
					ListFollowID(ctx, "user01").
					Return([]string{}, nil)
				mocks.UserRepository.EXPECT().
					ListFollowerID(ctx, "user01").
					Return([]string{}, nil)
			},
			args: args{
				userID:     "current-user",
				followerID: "user01",
			},
			want: want{
				user: &user.User{
					ID: "user01",
				},
				isFollowing:   false,
				isFollowed:    false,
				followCount:   0,
				followerCount: 0,
				err:           nil,
			},
		},
		{
			name: "failed: not found follower",
			setup: func(context context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.UserRepository.EXPECT().
					Get(ctx, "user01").
					Return(nil, exception.NotFound.New(test.ErrMock))
			},
			args: args{
				userID:     "current-user",
				followerID: "user01",
			},
			want: want{
				user:          nil,
				isFollowing:   false,
				isFollowed:    false,
				followCount:   0,
				followerCount: 0,
				err:           exception.NotFound.New(test.ErrMock),
			},
		},
		{
			name: "failed: not found relationship",
			setup: func(context context.Context, t *testing.T, mocks *test.Mocks) {
				u := &user.User{ID: "user01"}
				mocks.UserRepository.EXPECT().
					Get(ctx, "user01").
					Return(u, nil)
				mocks.UserRepository.EXPECT().
					GetRelationshipIDByUserID(ctx, "current-user", "user01").
					Return(0, exception.NotFound.New(test.ErrMock))
			},
			args: args{
				userID:     "current-user",
				followerID: "user01",
			},
			want: want{
				user:          nil,
				isFollowing:   false,
				isFollowed:    false,
				followCount:   0,
				followerCount: 0,
				err:           exception.NotFound.New(test.ErrMock),
			},
		},
		{
			name: "failed: internal error",
			setup: func(context context.Context, t *testing.T, mocks *test.Mocks) {
				u := &user.User{ID: "user01"}
				mocks.UserRepository.EXPECT().
					Get(ctx, "user01").
					Return(u, nil)
				mocks.UserRepository.EXPECT().
					GetRelationshipIDByUserID(ctx, "current-user", "user01").
					Return(1, nil)
				mocks.UserRepository.EXPECT().
					DeleteRelationship(ctx, 1).
					Return(exception.ErrorInDatastore.New(test.ErrMock))
			},
			args: args{
				userID:     "current-user",
				followerID: "user01",
			},
			want: want{
				user:          nil,
				isFollowing:   false,
				isFollowed:    false,
				followCount:   0,
				followerCount: 0,
				err:           exception.ErrorInDatastore.New(test.ErrMock),
			},
		},
	}

	for _, tt := range testCases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			mocks := test.NewMocks(ctrl)
			tt.setup(ctx, t, mocks)
			target := NewUserApplication(
				mocks.UserDomainValidation,
				mocks.UserRepository,
				mocks.UserUploader,
			)

			u,
				isFollowing,
				isFollowed,
				followCount,
				followerCount,
				err := target.Unfollow(ctx, tt.args.userID, tt.args.followerID)
			if tt.want.err != nil {
				require.Equal(t, tt.want.err.Error(), err.Error())
				return
			}
			require.NoError(t, err)
			require.Equal(t, tt.want.user, u)
			require.Equal(t, tt.want.isFollowing, isFollowing)
			require.Equal(t, tt.want.isFollowed, isFollowed)
			require.Equal(t, tt.want.followCount, followCount)
			require.Equal(t, tt.want.followerCount, followerCount)
		})
	}
}

func TestUserApplication_UploadThumbnail(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type args struct {
		userID    string
		thumbnail []byte
	}
	type want struct {
		thumbnailURL string
		err          error
	}
	testCases := []struct {
		name  string
		setup func(context.Context, *testing.T, *test.Mocks)
		args  args
		want  want
	}{
		{
			name: "success",
			setup: func(context context.Context, t *testing.T, mocks *test.Mocks) {
				mocks.UserUploader.EXPECT().
					Thumbnail(ctx, "user01", []byte{}).
					Return("https://go.dev/images/gophers/ladder.svg", nil)
			},
			args: args{
				userID:    "user01",
				thumbnail: []byte{},
			},
			want: want{
				thumbnailURL: "https://go.dev/images/gophers/ladder.svg",
				err:          nil,
			},
		},
	}

	for _, tt := range testCases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			mocks := test.NewMocks(ctrl)
			tt.setup(ctx, t, mocks)
			target := NewUserApplication(
				mocks.UserDomainValidation,
				mocks.UserRepository,
				mocks.UserUploader,
			)

			thumbnailURL, err := target.UploadThumbnail(ctx, tt.args.userID, tt.args.thumbnail)
			if tt.want.err != nil {
				require.Equal(t, tt.want.err.Error(), err.Error())
				require.Equal(t, tt.want.thumbnailURL, thumbnailURL)
				return
			}
			require.NoError(t, err)
			require.Equal(t, tt.want.thumbnailURL, thumbnailURL)
		})
	}
}

func TestUserApplication_HasAdminRole(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type args struct {
		role int
	}
	type want struct {
		err error
	}
	testCases := []struct {
		name  string
		setup func(context.Context, *testing.T, *test.Mocks)
		args  args
		want  want
	}{
		{
			name:  "success",
			setup: func(context context.Context, t *testing.T, mocks *test.Mocks) {},
			args: args{
				role: user.AdminRole,
			},
			want: want{
				err: nil,
			},
		},
		{
			name:  "failed: forbidden",
			setup: func(context context.Context, t *testing.T, mocks *test.Mocks) {},
			args: args{
				role: user.UserRole,
			},
			want: want{
				err: exception.Forbidden.New(xerrors.New("This account doesn't have administrator privileges")),
			},
		},
	}

	for _, tt := range testCases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			mocks := test.NewMocks(ctrl)
			tt.setup(ctx, t, mocks)
			target := NewUserApplication(
				mocks.UserDomainValidation,
				mocks.UserRepository,
				mocks.UserUploader,
			)

			err := target.HasAdminRole(tt.args.role)
			if tt.want.err != nil {
				require.Equal(t, tt.want.err.Error(), err.Error())
				return
			}
			require.NoError(t, err)
		})
	}
}

func testUser(id string) *user.User {
	current := time.Now().Local()

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
		CreatedAt:        current,
		UpdatedAt:        current,
	}
}

func testAdmin(id string) *user.User {
	current := time.Now().Local()

	return &user.User{
		ID:               id,
		Username:         "テストユーザー",
		Gender:           user.MaleGender,
		Email:            "test-user@calmato.jp",
		PhoneNumber:      "000-0000-0000",
		Role:             user.DeveloperRole,
		ThumbnailURL:     "https://go.dev/images/gophers/ladder.svg",
		SelfIntroduction: "テストコードです。",
		LastName:         "テスト",
		FirstName:        "ユーザー",
		LastNameKana:     "てすと",
		FirstNameKana:    "ゆーざー",
		PostalCode:       "",
		Prefecture:       "",
		City:             "",
		AddressLine1:     "",
		AddressLine2:     "",
		InstanceID:       "ExponentPushToken[!Qaz2wsx3edc4rfv5tgb6y]",
		CreatedAt:        current,
		UpdatedAt:        current,
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
