package validation

import (
	"strconv"
	"testing"

	"github.com/calmato/gran-book/api/proto/user"
	"github.com/stretchr/testify/assert"
)

func TestUserRequestValidation_ListUser(t *testing.T) {
	t.Parallel()
	type args struct {
		req *user.ListUserRequest
	}

	testCases := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "success",
			args: args{
				req: &user.ListUserRequest{
					Limit:  200,
					Offset: 100,
				},
			},
			want: true,
		},
		{
			name: "validation error: Limit.lte",
			args: args{
				req: &user.ListUserRequest{
					Search: &user.Search{
						Field: "username",
						Value: "テストユーザー",
					},
					Order: &user.Order{
						Field:   "created_at",
						OrderBy: user.OrderBy_ORDER_BY_ASC,
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
				req: &user.ListUserRequest{
					Search: &user.Search{
						Field: "username",
						Value: "テストユーザー",
					},
					Order: &user.Order{
						Field:   "created_at",
						OrderBy: user.OrderBy_ORDER_BY_ASC,
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
			target := NewUserRequestValidation()

			got := target.ListUser(tt.args.req)
			switch tt.want {
			case true:
				assert.NoError(t, got)
			case false:
				assert.Error(t, got)
			}
		})
	}
}

func TestUserRequestValidation_ListFollow(t *testing.T) {
	t.Parallel()
	type args struct {
		req *user.ListFollowRequest
	}

	testCases := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "success",
			args: args{
				req: &user.ListFollowRequest{
					UserId: "12345678-1234-1234-123456789012",
					Limit:  200,
					Offset: 100,
				},
			},
			want: true,
		},
		{
			name: "validation error: UserId.min_len",
			args: args{
				req: &user.ListFollowRequest{
					UserId: "",
					Order: &user.Order{
						Field:   "created_at",
						OrderBy: user.OrderBy_ORDER_BY_ASC,
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
				req: &user.ListFollowRequest{
					UserId: "12345678-1234-1234-123456789012",
					Order: &user.Order{
						Field:   "created_at",
						OrderBy: user.OrderBy_ORDER_BY_ASC,
					},
					Limit:  201,
					Offset: 100,
				},
			},
			want: false,
		},
		{
			name: "validation error: Offset.gte",
			args: args{
				req: &user.ListFollowRequest{
					UserId: "12345678-1234-1234-123456789012",
					Order: &user.Order{
						Field:   "created_at",
						OrderBy: user.OrderBy_ORDER_BY_ASC,
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
			target := NewUserRequestValidation()

			got := target.ListFollow(tt.args.req)
			switch tt.want {
			case true:
				assert.NoError(t, got)
			case false:
				assert.Error(t, got)
			}
		})
	}
}

func TestUserRequestValidation_ListFollower(t *testing.T) {
	t.Parallel()
	type args struct {
		req *user.ListFollowerRequest
	}

	testCases := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "success",
			args: args{
				req: &user.ListFollowerRequest{
					UserId: "12345678-1234-1234-123456789012",
					Limit:  200,
					Offset: 100,
				},
			},
			want: true,
		},
		{
			name: "validation error: UserId.min_len",
			args: args{
				req: &user.ListFollowerRequest{
					UserId: "",
					Order: &user.Order{
						Field:   "created_at",
						OrderBy: user.OrderBy_ORDER_BY_ASC,
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
				req: &user.ListFollowerRequest{
					UserId: "12345678-1234-1234-123456789012",
					Order: &user.Order{
						Field:   "created_at",
						OrderBy: user.OrderBy_ORDER_BY_ASC,
					},
					Limit:  201,
					Offset: 100,
				},
			},
			want: false,
		},
		{
			name: "validation error: Offset.gte",
			args: args{
				req: &user.ListFollowerRequest{
					UserId: "12345678-1234-1234-123456789012",
					Order: &user.Order{
						Field:   "created_at",
						OrderBy: user.OrderBy_ORDER_BY_ASC,
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
			target := NewUserRequestValidation()

			got := target.ListFollower(tt.args.req)
			switch tt.want {
			case true:
				assert.NoError(t, got)
			case false:
				assert.Error(t, got)
			}
		})
	}
}

func TestUserRequestValidation_MultiGetUser(t *testing.T) {
	t.Parallel()
	type args struct {
		req *user.MultiGetUserRequest
	}

	testCases := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "success",
			args: args{
				req: &user.MultiGetUserRequest{
					UserIds: []string{"12345678-1234-1234-123456789012"},
				},
			},
			want: true,
		},
		{
			name: "validation error: UserIds.max_items",
			args: args{
				req: &user.MultiGetUserRequest{
					UserIds: func() []string {
						userIDs := make([]string, 201)
						for i := 0; i < len(userIDs); i++ {
							userIDs[i] = strconv.Itoa(i)
						}
						return userIDs
					}(),
				},
			},
			want: false,
		},
		{
			name: "validation error: UserIds.unique",
			args: args{
				req: &user.MultiGetUserRequest{
					UserIds: []string{
						"12345678-1234-1234-123456789012",
						"12345678-1234-1234-123456789012",
					},
				},
			},
			want: false,
		},
	}

	for _, tt := range testCases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			target := NewUserRequestValidation()

			got := target.MultiGetUser(tt.args.req)
			switch tt.want {
			case true:
				assert.NoError(t, got)
			case false:
				assert.Error(t, got)
			}
		})
	}
}

func TestUserRequestValidation_GetUser(t *testing.T) {
	t.Parallel()
	type args struct {
		req *user.GetUserRequest
	}

	testCases := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "success",
			args: args{
				req: &user.GetUserRequest{
					UserId: "12345678-1234-1234-123456789012",
				},
			},
			want: true,
		},
		{
			name: "validation error: UserId.min_len",
			args: args{
				req: &user.GetUserRequest{
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
			target := NewUserRequestValidation()

			got := target.GetUser(tt.args.req)
			switch tt.want {
			case true:
				assert.NoError(t, got)
			case false:
				assert.Error(t, got)
			}
		})
	}
}

func TestUserRequestValidation_GetUserProfile(t *testing.T) {
	t.Parallel()
	type args struct {
		req *user.GetUserProfileRequest
	}

	testCases := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "success",
			args: args{
				req: &user.GetUserProfileRequest{
					UserId: "12345678-1234-1234-123456789012",
				},
			},
			want: true,
		},
		{
			name: "validation error: UserId.min_len",
			args: args{
				req: &user.GetUserProfileRequest{
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
			target := NewUserRequestValidation()

			got := target.GetUserProfile(tt.args.req)
			switch tt.want {
			case true:
				assert.NoError(t, got)
			case false:
				assert.Error(t, got)
			}
		})
	}
}

func TestUserRequestValidation_Follow(t *testing.T) {
	t.Parallel()
	type args struct {
		req *user.FollowRequest
	}

	testCases := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "success",
			args: args{
				req: &user.FollowRequest{
					UserId:     "12345678-1234-1234-123456789012",
					FollowerId: "87654321-4321-4321-210987654321",
				},
			},
			want: true,
		},
		{
			name: "validation error: UserId.min_len",
			args: args{
				req: &user.FollowRequest{
					UserId:     "",
					FollowerId: "87654321-4321-4321-210987654321",
				},
			},
			want: false,
		},
		{
			name: "validation error: FollowerId.min_len",
			args: args{
				req: &user.FollowRequest{
					UserId:     "12345678-1234-1234-123456789012",
					FollowerId: "",
				},
			},
			want: false,
		},
		{
			name: "validation error: FollowerId.unique",
			args: args{
				req: &user.FollowRequest{
					UserId:     "12345678-1234-1234-123456789012",
					FollowerId: "12345678-1234-1234-123456789012",
				},
			},
			want: false,
		},
	}

	for _, tt := range testCases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			target := NewUserRequestValidation()

			got := target.Follow(tt.args.req)
			switch tt.want {
			case true:
				assert.NoError(t, got)
			case false:
				assert.Error(t, got)
			}
		})
	}
}

func TestUserRequestValidation_Unfollow(t *testing.T) {
	t.Parallel()
	type args struct {
		req *user.UnfollowRequest
	}

	testCases := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "success",
			args: args{
				req: &user.UnfollowRequest{
					UserId:     "12345678-1234-1234-123456789012",
					FollowerId: "87654321-4321-4321-210987654321",
				},
			},
			want: true,
		},
		{
			name: "validation error: UserId.min_len",
			args: args{
				req: &user.UnfollowRequest{
					UserId:     "",
					FollowerId: "87654321-4321-4321-210987654321",
				},
			},
			want: false,
		},
		{
			name: "validation error: FollowerId.min_len",
			args: args{
				req: &user.UnfollowRequest{
					UserId:     "12345678-1234-1234-123456789012",
					FollowerId: "",
				},
			},
			want: false,
		},
		{
			name: "validation error: FollowerId.unique",
			args: args{
				req: &user.UnfollowRequest{
					UserId:     "12345678-1234-1234-123456789012",
					FollowerId: "12345678-1234-1234-123456789012",
				},
			},
			want: false,
		},
	}

	for _, tt := range testCases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			target := NewUserRequestValidation()

			got := target.Unfollow(tt.args.req)
			switch tt.want {
			case true:
				assert.NoError(t, got)
			case false:
				assert.Error(t, got)
			}
		})
	}
}
