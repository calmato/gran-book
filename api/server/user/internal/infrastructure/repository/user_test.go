package repository

import (
	"context"
	"fmt"
	"testing"

	"github.com/calmato/gran-book/api/server/user/internal/domain/user"
	"github.com/calmato/gran-book/api/server/user/pkg/database"
	"github.com/calmato/gran-book/api/server/user/pkg/test"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUserRepository_List(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mocks, err := test.NewDBMock(ctrl)
	require.NoError(t, err)

	err = mocks.DeleteAll(mocks.UserDB, "users")
	require.NoError(t, err)

	users := make([]*user.User, 3)
	users[0] = testUser("00000000-0000-0000-0000-000000000000")
	users[1] = testUser("11111111-1111-1111-1111-111111111111")
	users[2] = testUser("22222222-2222-2222-2222-222222222222")

	err = mocks.UserDB.DB.Table("users").Create(&users).Error
	require.NoError(t, err)

	type args struct {
		query *database.ListQuery
	}
	type want struct {
		users []*user.User
		isErr bool
	}
	testCases := []struct {
		name string
		args args
		want want
	}{
		{
			name: "success",
			args: args{
				query: &database.ListQuery{
					Limit:      20,
					Offset:     0,
					Conditions: []*database.ConditionQuery{},
				},
			},
			want: want{
				users: users,
				isErr: false,
			},
		},
	}

	for _, tt := range testCases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			target := NewUserRepository(mocks.UserDB, nil)

			us, err := target.List(ctx, tt.args.query)
			assert.Equal(t, tt.want.isErr, err != nil, err)
			for i := range tt.want.users {
				assert.Contains(t, us, tt.want.users[i])
			}
		})
	}
}

func TestUserRepository_ListFollow(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mocks, err := test.NewDBMock(ctrl)
	require.NoError(t, err)

	err = mocks.DeleteAll(mocks.UserDB, "users")
	require.NoError(t, err)

	users := make([]*user.User, 3)
	users[0] = testUser("00000000-0000-0000-0000-000000000000")
	users[0].Username = "テストユーザー1"
	users[1] = testUser("11111111-1111-1111-1111-111111111111")
	users[1].Username = "テストユーザー2"
	users[2] = testUser("22222222-2222-2222-2222-222222222222")
	users[2].Username = "テストユーザー3"

	err = mocks.UserDB.DB.Table("users").Create(&users).Error
	require.NoError(t, err)

	relationships := make([]*user.Relationship, 2)
	relationships[0] = testRelationship(1, users[1].ID, users[0].ID)
	relationships[1] = testRelationship(2, users[2].ID, users[0].ID)

	err = mocks.UserDB.DB.Table("relationships").Create(&relationships).Error
	require.NoError(t, err)

	type args struct {
		query *database.ListQuery
	}
	type want struct {
		follows []*user.Follow
		isErr   bool
	}
	testCases := []struct {
		name string
		args args
		want want
	}{
		{
			name: "success",
			args: args{
				query: &database.ListQuery{
					Limit:  20,
					Offset: 0,
					Conditions: []*database.ConditionQuery{
						{
							Field:    "follower_id",
							Operator: "==",
							Value:    "00000000-0000-0000-0000-000000000000",
						},
					},
				},
			},
			want: want{
				follows: []*user.Follow{
					{
						FollowID:         "11111111-1111-1111-1111-111111111111",
						Username:         "テストユーザー2",
						ThumbnailURL:     "https://go.dev/images/gophers/ladder.svg",
						SelfIntroduction: "テストコードです。",
					},
					{
						FollowID:         "22222222-2222-2222-2222-222222222222",
						Username:         "テストユーザー3",
						ThumbnailURL:     "https://go.dev/images/gophers/ladder.svg",
						SelfIntroduction: "テストコードです。",
					},
				},
				isErr: false,
			},
		},
	}

	for _, tt := range testCases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			target := NewUserRepository(mocks.UserDB, nil)

			fs, err := target.ListFollow(ctx, tt.args.query)
			assert.Equal(t, tt.want.isErr, err != nil, err)
			for i := range tt.want.follows {
				assert.Contains(t, fs, tt.want.follows[i])
			}
		})
	}
}

func TestUserRepository_ListFollower(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mocks, err := test.NewDBMock(ctrl)
	require.NoError(t, err)

	err = mocks.DeleteAll(mocks.UserDB, "users")
	require.NoError(t, err)

	users := make([]*user.User, 3)
	users[0] = testUser("00000000-0000-0000-0000-000000000000")
	users[0].Username = "テストユーザー1"
	users[1] = testUser("11111111-1111-1111-1111-111111111111")
	users[1].Username = "テストユーザー2"
	users[2] = testUser("22222222-2222-2222-2222-222222222222")
	users[2].Username = "テストユーザー3"

	err = mocks.UserDB.DB.Table("users").Create(&users).Error
	require.NoError(t, err)

	relationships := make([]*user.Relationship, 2)
	relationships[0] = testRelationship(1, users[0].ID, users[1].ID)
	relationships[1] = testRelationship(2, users[0].ID, users[2].ID)

	err = mocks.UserDB.DB.Table("relationships").Create(&relationships).Error
	require.NoError(t, err)

	type args struct {
		query *database.ListQuery
	}
	type want struct {
		followers []*user.Follower
		isErr     bool
	}
	testCases := []struct {
		name string
		args args
		want want
	}{
		{
			name: "success",
			args: args{
				query: &database.ListQuery{
					Limit:  20,
					Offset: 0,
					Conditions: []*database.ConditionQuery{
						{
							Field:    "follow_id",
							Operator: "==",
							Value:    "00000000-0000-0000-0000-000000000000",
						},
					},
				},
			},
			want: want{
				followers: []*user.Follower{
					{
						FollowerID:       "11111111-1111-1111-1111-111111111111",
						Username:         "テストユーザー2",
						ThumbnailURL:     "https://go.dev/images/gophers/ladder.svg",
						SelfIntroduction: "テストコードです。",
					},
					{
						FollowerID:       "22222222-2222-2222-2222-222222222222",
						Username:         "テストユーザー3",
						ThumbnailURL:     "https://go.dev/images/gophers/ladder.svg",
						SelfIntroduction: "テストコードです。",
					},
				},
				isErr: false,
			},
		},
	}

	for _, tt := range testCases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			target := NewUserRepository(mocks.UserDB, nil)

			fs, err := target.ListFollower(ctx, tt.args.query)
			assert.Equal(t, tt.want.isErr, err != nil, err)
			for i := range tt.want.followers {
				assert.Contains(t, fs, tt.want.followers[i])
			}
		})
	}
}

func TestUserRepository_ListInstanceID(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mocks, err := test.NewDBMock(ctrl)
	require.NoError(t, err)

	err = mocks.DeleteAll(mocks.UserDB, "users")
	require.NoError(t, err)

	users := make([]*user.User, 3)
	users[0] = testUser("00000000-0000-0000-0000-000000000000")
	users[0].InstanceID = "ExponentPushToken[qwertyuiop]"
	users[1] = testUser("11111111-1111-1111-1111-111111111111")
	users[1].InstanceID = "ExponentPushToken[asdfghjkl]"
	users[2] = testUser("22222222-2222-2222-2222-222222222222")
	users[2].InstanceID = "ExponentPushToken[zxcvbnm]"

	err = mocks.UserDB.DB.Table("users").Create(&users).Error
	require.NoError(t, err)

	type args struct {
		query *database.ListQuery
	}
	type want struct {
		instanceIDs []string
		isErr       bool
	}
	testCases := []struct {
		name string
		args args
		want want
	}{
		{
			name: "success",
			args: args{
				query: &database.ListQuery{
					Limit:      20,
					Offset:     0,
					Conditions: []*database.ConditionQuery{},
				},
			},
			want: want{
				instanceIDs: []string{
					"ExponentPushToken[qwertyuiop]",
					"ExponentPushToken[asdfghjkl]",
					"ExponentPushToken[zxcvbnm]",
				},
				isErr: false,
			},
		},
	}

	for _, tt := range testCases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			target := NewUserRepository(mocks.UserDB, nil)

			us, err := target.ListInstanceID(ctx, tt.args.query)
			assert.Equal(t, tt.want.isErr, err != nil, err)
			for i := range tt.want.instanceIDs {
				assert.Contains(t, us, tt.want.instanceIDs[i])
			}
		})
	}
}

func TestUserRepository_ListFollowID(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mocks, err := test.NewDBMock(ctrl)
	require.NoError(t, err)

	err = mocks.DeleteAll(mocks.UserDB, "users")
	require.NoError(t, err)

	users := make([]*user.User, 3)
	users[0] = testUser("00000000-0000-0000-0000-000000000000")
	users[0].Username = "テストユーザー1"
	users[1] = testUser("11111111-1111-1111-1111-111111111111")
	users[1].Username = "テストユーザー2"
	users[2] = testUser("22222222-2222-2222-2222-222222222222")
	users[2].Username = "テストユーザー3"

	err = mocks.UserDB.DB.Table("users").Create(&users).Error
	require.NoError(t, err)

	relationships := make([]*user.Relationship, 2)
	relationships[0] = testRelationship(1, users[1].ID, users[0].ID)
	relationships[1] = testRelationship(2, users[2].ID, users[0].ID)

	err = mocks.UserDB.DB.Table("relationships").Create(&relationships).Error
	require.NoError(t, err)

	type args struct {
		followerID string
		followIDs  []string
	}
	type want struct {
		followIDs []string
		isErr     bool
	}
	testCases := []struct {
		name string
		args args
		want want
	}{
		{
			name: "success",
			args: args{
				followerID: "00000000-0000-0000-0000-000000000000",
				followIDs:  []string{},
			},
			want: want{
				followIDs: []string{
					"11111111-1111-1111-1111-111111111111",
					"22222222-2222-2222-2222-222222222222",
				},
				isErr: false,
			},
		},
		{
			name: "success: filter followIDs",
			args: args{
				followerID: "00000000-0000-0000-0000-000000000000",
				followIDs: []string{
					"11111111-1111-1111-1111-111111111111",
				},
			},
			want: want{
				followIDs: []string{
					"11111111-1111-1111-1111-111111111111",
				},
				isErr: false,
			},
		},
	}

	for _, tt := range testCases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			target := NewUserRepository(mocks.UserDB, nil)

			followIDs, err := target.ListFollowID(ctx, tt.args.followerID, tt.args.followIDs...)
			assert.Equal(t, tt.want.isErr, err != nil, err)
			for i := range tt.want.followIDs {
				assert.Contains(t, followIDs, tt.want.followIDs[i])
			}
		})
	}
}

func TestUserRepository_ListFollowerID(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mocks, err := test.NewDBMock(ctrl)
	require.NoError(t, err)

	err = mocks.DeleteAll(mocks.UserDB, "users")
	require.NoError(t, err)

	users := make([]*user.User, 3)
	users[0] = testUser("00000000-0000-0000-0000-000000000000")
	users[0].Username = "テストユーザー1"
	users[1] = testUser("11111111-1111-1111-1111-111111111111")
	users[1].Username = "テストユーザー2"
	users[2] = testUser("22222222-2222-2222-2222-222222222222")
	users[2].Username = "テストユーザー3"

	err = mocks.UserDB.DB.Table("users").Create(&users).Error
	require.NoError(t, err)

	relationships := make([]*user.Relationship, 2)
	relationships[0] = testRelationship(1, users[0].ID, users[1].ID)
	relationships[1] = testRelationship(2, users[0].ID, users[2].ID)

	err = mocks.UserDB.DB.Table("relationships").Create(&relationships).Error
	require.NoError(t, err)

	type args struct {
		followID    string
		followerIDs []string
	}
	type want struct {
		followerIDs []string
		isErr       bool
	}
	testCases := []struct {
		name string
		args args
		want want
	}{
		{
			name: "success",
			args: args{
				followID:    "00000000-0000-0000-0000-000000000000",
				followerIDs: []string{},
			},
			want: want{
				followerIDs: []string{
					"11111111-1111-1111-1111-111111111111",
					"22222222-2222-2222-2222-222222222222",
				},
				isErr: false,
			},
		},
		{
			name: "success: filter followerIDs",
			args: args{
				followID: "00000000-0000-0000-0000-000000000000",
				followerIDs: []string{
					"11111111-1111-1111-1111-111111111111",
				},
			},
			want: want{
				followerIDs: []string{
					"11111111-1111-1111-1111-111111111111",
				},
				isErr: false,
			},
		},
	}

	for _, tt := range testCases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			target := NewUserRepository(mocks.UserDB, nil)

			followerIDs, err := target.ListFollowerID(ctx, tt.args.followID, tt.args.followerIDs...)
			assert.Equal(t, tt.want.isErr, err != nil, err)
			for i := range tt.want.followerIDs {
				assert.Contains(t, followerIDs, tt.want.followerIDs[i])
			}
		})
	}
}

func TestUserRepository_Count(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mocks, err := test.NewDBMock(ctrl)
	require.NoError(t, err)

	err = mocks.DeleteAll(mocks.UserDB, "users")
	require.NoError(t, err)

	users := make([]*user.User, 3)
	users[0] = testUser("00000000-0000-0000-0000-000000000000")
	users[1] = testUser("11111111-1111-1111-1111-111111111111")
	users[2] = testUser("22222222-2222-2222-2222-222222222222")

	err = mocks.UserDB.DB.Table("users").Create(&users).Error
	require.NoError(t, err)

	type args struct {
		query *database.ListQuery
	}
	type want struct {
		count int
		isErr bool
	}
	testCases := []struct {
		name string
		args args
		want want
	}{
		{
			name: "success",
			args: args{
				query: &database.ListQuery{
					Limit:      20,
					Offset:     0,
					Conditions: []*database.ConditionQuery{},
				},
			},
			want: want{
				count: 3,
				isErr: false,
			},
		},
	}

	for _, tt := range testCases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			target := NewUserRepository(mocks.UserDB, nil)

			count, err := target.Count(ctx, tt.args.query)
			assert.Equal(t, tt.want.isErr, err != nil, err)
			assert.Equal(t, tt.want.count, count)
		})
	}
}

func TestUserRepository_CountRelationship(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mocks, err := test.NewDBMock(ctrl)
	require.NoError(t, err)

	err = mocks.DeleteAll(mocks.UserDB, "users")
	require.NoError(t, err)

	users := make([]*user.User, 3)
	users[0] = testUser("00000000-0000-0000-0000-000000000000")
	users[0].Username = "テストユーザー1"
	users[1] = testUser("11111111-1111-1111-1111-111111111111")
	users[1].Username = "テストユーザー2"
	users[2] = testUser("22222222-2222-2222-2222-222222222222")
	users[2].Username = "テストユーザー3"

	err = mocks.UserDB.DB.Table("users").Create(&users).Error
	require.NoError(t, err)

	relationships := make([]*user.Relationship, 2)
	relationships[0] = testRelationship(1, users[1].ID, users[0].ID)
	relationships[1] = testRelationship(2, users[2].ID, users[0].ID)

	err = mocks.UserDB.DB.Table("relationships").Create(&relationships).Error
	require.NoError(t, err)

	type args struct {
		query *database.ListQuery
	}
	type want struct {
		count int
		isErr bool
	}
	testCases := []struct {
		name string
		args args
		want want
	}{
		{
			name: "success",
			args: args{
				query: &database.ListQuery{
					Limit:  20,
					Offset: 0,
					Conditions: []*database.ConditionQuery{
						{
							Field:    "follower_id",
							Operator: "==",
							Value:    "00000000-0000-0000-0000-000000000000",
						},
					},
				},
			},
			want: want{
				count: 2,
				isErr: false,
			},
		},
	}

	for _, tt := range testCases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			target := NewUserRepository(mocks.UserDB, nil)

			count, err := target.CountRelationship(ctx, tt.args.query)
			assert.Equal(t, tt.want.isErr, err != nil, err)
			assert.Equal(t, tt.want.count, count)
		})
	}
}

func TestUserRepository_MultiGet(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mocks, err := test.NewDBMock(ctrl)
	require.NoError(t, err)

	err = mocks.DeleteAll(mocks.UserDB, "users")
	require.NoError(t, err)

	users := make([]*user.User, 3)
	users[0] = testUser("00000000-0000-0000-0000-000000000000")
	users[1] = testUser("11111111-1111-1111-1111-111111111111")
	users[2] = testUser("22222222-2222-2222-2222-222222222222")

	err = mocks.UserDB.DB.Table("users").Create(&users).Error
	require.NoError(t, err)

	type args struct {
		userIDs []string
	}
	type want struct {
		users []*user.User
		isErr bool
	}
	testCases := []struct {
		name string
		args args
		want want
	}{
		{
			name: "success",
			args: args{
				userIDs: []string{
					"00000000-0000-0000-0000-000000000000",
					"11111111-1111-1111-1111-111111111111",
					"22222222-2222-2222-2222-222222222222",
				},
			},
			want: want{
				users: users,
				isErr: false,
			},
		},
	}

	for _, tt := range testCases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			target := NewUserRepository(mocks.UserDB, nil)

			us, err := target.MultiGet(ctx, tt.args.userIDs)
			assert.Equal(t, tt.want.isErr, err != nil, err)
			for i := range tt.want.users {
				assert.Contains(t, us, tt.want.users[i])
			}
		})
	}
}

func TestUserRepository_Get(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mocks, err := test.NewDBMock(ctrl)
	require.NoError(t, err)

	err = mocks.DeleteAll(mocks.UserDB, "users")
	require.NoError(t, err)

	u := testUser("00000000-0000-0000-0000-000000000000")

	err = mocks.UserDB.DB.Table("users").Create(u).Error
	require.NoError(t, err)

	type args struct {
		userID string
	}
	type want struct {
		user  *user.User
		isErr bool
	}
	testCases := []struct {
		name string
		args args
		want want
	}{
		{
			name: "success",
			args: args{
				userID: "00000000-0000-0000-0000-000000000000",
			},
			want: want{
				user:  u,
				isErr: false,
			},
		},
		{
			name: "failed: not found",
			args: args{
				userID: "",
			},
			want: want{
				user:  nil,
				isErr: true,
			},
		},
	}

	for _, tt := range testCases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			target := NewUserRepository(mocks.UserDB, nil)

			u, err := target.Get(ctx, tt.args.userID)
			assert.Equal(t, tt.want.isErr, err != nil, err)
			assert.Equal(t, tt.want.user, u)
		})
	}
}

func TestUserRepository_GetAdmin(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mocks, err := test.NewDBMock(ctrl)
	require.NoError(t, err)

	err = mocks.DeleteAll(mocks.UserDB, "users")
	require.NoError(t, err)

	users := make([]*user.User, 2)
	users[0] = testUser("00000000-0000-0000-0000-000000000000")
	users[1] = testAdmin("11111111-1111-1111-1111-111111111111")

	err = mocks.UserDB.DB.Table("users").Create(&users).Error
	require.NoError(t, err)

	type args struct {
		userID string
	}
	type want struct {
		user  *user.User
		isErr bool
	}
	testCases := []struct {
		name string
		args args
		want want
	}{
		{
			name: "success",
			args: args{
				userID: "11111111-1111-1111-1111-111111111111",
			},
			want: want{
				user:  users[1],
				isErr: false,
			},
		},
		{
			name: "failed: not found",
			args: args{
				userID: "",
			},
			want: want{
				user:  nil,
				isErr: true,
			},
		},
		{
			name: "failed: not found when role is user role",
			args: args{
				userID: "00000000-0000-0000-0000-000000000000",
			},
			want: want{
				user:  nil,
				isErr: true,
			},
		},
	}

	for _, tt := range testCases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			target := NewUserRepository(mocks.UserDB, nil)

			u, err := target.GetAdmin(ctx, tt.args.userID)
			assert.Equal(t, tt.want.isErr, err != nil, err)
			assert.Equal(t, tt.want.user, u)
		})
	}
}

func TestUserRepository_GetRelationship(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mocks, err := test.NewDBMock(ctrl)
	require.NoError(t, err)

	err = mocks.DeleteAll(mocks.UserDB, "users")
	require.NoError(t, err)

	users := make([]*user.User, 2)
	users[0] = testUser("00000000-0000-0000-0000-000000000000")
	users[0].Username = "テストユーザー1"
	users[1] = testUser("11111111-1111-1111-1111-111111111111")
	users[1].Username = "テストユーザー2"

	err = mocks.UserDB.DB.Table("users").Create(&users).Error
	require.NoError(t, err)

	relationship := testRelationship(1, users[1].ID, users[0].ID)

	err = mocks.UserDB.DB.Table("relationships").Create(relationship).Error
	require.NoError(t, err)

	type args struct {
		followID   string
		followerID string
	}
	type want struct {
		relationship *user.Relationship
		isErr        bool
	}
	testCases := []struct {
		name string
		args args
		want want
	}{
		{
			name: "success",
			args: args{
				followID:   "11111111-1111-1111-1111-111111111111",
				followerID: "00000000-0000-0000-0000-000000000000",
			},
			want: want{
				relationship: relationship,
				isErr:        false,
			},
		},
	}

	for _, tt := range testCases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			target := NewUserRepository(mocks.UserDB, nil)

			r, err := target.GetRelationship(ctx, tt.args.followID, tt.args.followerID)
			assert.Equal(t, tt.want.isErr, err != nil, err)
			assert.Equal(t, tt.want.relationship, r)
		})
	}
}

func TestUserRepository_CreateRelationship(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mocks, err := test.NewDBMock(ctrl)
	require.NoError(t, err)

	err = mocks.DeleteAll(mocks.UserDB, "users")
	require.NoError(t, err)

	users := make([]*user.User, 2)
	users[0] = testUser("00000000-0000-0000-0000-000000000000")
	users[0].Username = "テストユーザー1"
	users[1] = testUser("11111111-1111-1111-1111-111111111111")
	users[1].Username = "テストユーザー2"

	err = mocks.UserDB.DB.Table("users").Create(&users).Error
	require.NoError(t, err)

	type args struct {
		relationship *user.Relationship
	}
	type want struct {
		isErr bool
	}
	testCases := []struct {
		name string
		args args
		want want
	}{
		{
			name: "success",
			args: args{
				relationship: &user.Relationship{
					FollowID:   "00000000-0000-0000-0000-000000000000",
					FollowerID: "11111111-1111-1111-1111-111111111111",
					CreatedAt:  test.TimeMock,
					UpdatedAt:  test.TimeMock,
				},
			},
			want: want{
				isErr: false,
			},
		},
		{
			name: "failed: internal error",
			args: args{
				relationship: &user.Relationship{},
			},
			want: want{
				isErr: true,
			},
		},
	}

	for _, tt := range testCases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			target := NewUserRepository(mocks.UserDB, nil)

			err := mocks.DeleteAll(mocks.UserDB, "relationships")
			require.NoError(t, err)

			err = target.CreateRelationship(ctx, tt.args.relationship)
			assert.Equal(t, tt.want.isErr, err != nil, err)
		})
	}
}

func TestUserRepository_DeleteRelationship(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mocks, err := test.NewDBMock(ctrl)
	require.NoError(t, err)

	err = mocks.DeleteAll(mocks.UserDB, "users")
	require.NoError(t, err)

	users := make([]*user.User, 2)
	users[0] = testUser("00000000-0000-0000-0000-000000000000")
	users[0].Username = "テストユーザー1"
	users[1] = testUser("11111111-1111-1111-1111-111111111111")
	users[1].Username = "テストユーザー2"

	err = mocks.UserDB.DB.Table("users").Create(&users).Error
	require.NoError(t, err)

	relationship := testRelationship(1, users[1].ID, users[0].ID)

	type args struct {
		relationshipID int
	}
	type want struct {
		isErr bool
	}
	testCases := []struct {
		name string
		args args
		want want
	}{
		{
			name: "success",
			args: args{
				relationshipID: 1,
			},
			want: want{
				isErr: false,
			},
		},
	}

	for _, tt := range testCases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			target := NewUserRepository(mocks.UserDB, nil)

			err := mocks.DeleteAll(mocks.UserDB, "relationships")
			require.NoError(t, err)

			err = mocks.UserDB.DB.Table("relationships").Create(relationship).Error
			require.NoError(t, err)

			err = target.DeleteRelationship(ctx, tt.args.relationshipID)
			assert.Equal(t, tt.want.isErr, err != nil, err)
		})
	}
}

func testUser(id string) *user.User {
	return &user.User{
		ID:               id,
		Username:         "テストユーザー",
		Gender:           user.MaleGender,
		Email:            fmt.Sprintf("%s@calmato.jp", id),
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

func testAdmin(id string) *user.User {
	return &user.User{
		ID:               id,
		Username:         "テストユーザー",
		Gender:           user.MaleGender,
		Email:            fmt.Sprintf("%s@calmato.jp", id),
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
		CreatedAt:        test.TimeMock,
		UpdatedAt:        test.TimeMock,
	}
}

func testRelationship(id int, followID, followerID string) *user.Relationship {
	return &user.Relationship{
		ID:         id,
		FollowID:   followID,
		FollowerID: followerID,
		CreatedAt:  test.TimeMock,
		UpdatedAt:  test.TimeMock,
	}
}
