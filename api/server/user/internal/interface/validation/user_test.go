package validation

import (
	"strconv"
	"testing"

	pb "github.com/calmato/gran-book/api/server/user/proto"
	"github.com/stretchr/testify/assert"
)

func TestUserRequestValidation_ListUser(t *testing.T) {
	t.Parallel()
	type args struct {
		req *pb.ListUserRequest
	}

	testCases := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "success",
			args: args{
				req: &pb.ListUserRequest{
					Limit:  200,
					Offset: 100,
				},
			},
			want: true,
		},
		{
			name: "validation error: Search.Field.min_len",
			args: args{
				req: &pb.ListUserRequest{
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
				req: &pb.ListUserRequest{
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
			name: "validation error: Order.OrderBy.in",
			args: args{
				req: &pb.ListUserRequest{
					Search: &pb.Search{
						Field: "username",
						Value: "テストユーザー",
					},
					Order: &pb.Order{
						Field:   "created_at",
						OrderBy: 2,
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
				req: &pb.ListUserRequest{
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
				req: &pb.ListUserRequest{
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
		req *pb.ListFollowRequest
	}

	testCases := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "success",
			args: args{
				req: &pb.ListFollowRequest{
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
				req: &pb.ListFollowRequest{
					UserId: "",
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
			name: "validation error: Order.OrderBy.in",
			args: args{
				req: &pb.ListFollowRequest{
					UserId: "12345678-1234-1234-123456789012",
					Order: &pb.Order{
						Field:   "created_at",
						OrderBy: 2,
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
				req: &pb.ListFollowRequest{
					UserId: "12345678-1234-1234-123456789012",
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
			name: "validation error: Offset.gte",
			args: args{
				req: &pb.ListFollowRequest{
					UserId: "12345678-1234-1234-123456789012",
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
		req *pb.ListFollowerRequest
	}

	testCases := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "success",
			args: args{
				req: &pb.ListFollowerRequest{
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
				req: &pb.ListFollowerRequest{
					UserId: "",
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
			name: "validation error: Order.OrderBy.in",
			args: args{
				req: &pb.ListFollowerRequest{
					UserId: "12345678-1234-1234-123456789012",
					Order: &pb.Order{
						Field:   "created_at",
						OrderBy: 2,
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
				req: &pb.ListFollowerRequest{
					UserId: "12345678-1234-1234-123456789012",
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
			name: "validation error: Offset.gte",
			args: args{
				req: &pb.ListFollowerRequest{
					UserId: "12345678-1234-1234-123456789012",
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
		req *pb.MultiGetUserRequest
	}

	testCases := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "success",
			args: args{
				req: &pb.MultiGetUserRequest{
					UserIds: []string{"12345678-1234-1234-123456789012"},
				},
			},
			want: true,
		},
		{
			name: "validation error: UserIds.max_items",
			args: args{
				req: &pb.MultiGetUserRequest{
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
				req: &pb.MultiGetUserRequest{
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
		req *pb.GetUserRequest
	}

	testCases := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "success",
			args: args{
				req: &pb.GetUserRequest{
					UserId: "12345678-1234-1234-123456789012",
				},
			},
			want: true,
		},
		{
			name: "validation error: UserId.min_len",
			args: args{
				req: &pb.GetUserRequest{
					UserId: "",
				},
			},
			want: false,
		},
	}

	for _, tt := range testCases {
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
		req *pb.GetUserProfileRequest
	}

	testCases := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "success",
			args: args{
				req: &pb.GetUserProfileRequest{
					UserId: "12345678-1234-1234-123456789012",
				},
			},
			want: true,
		},
		{
			name: "validation error: UserId.min_len",
			args: args{
				req: &pb.GetUserProfileRequest{
					UserId: "",
				},
			},
			want: false,
		},
	}

	for _, tt := range testCases {
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
		req *pb.FollowRequest
	}

	testCases := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "success",
			args: args{
				req: &pb.FollowRequest{
					UserId:     "12345678-1234-1234-123456789012",
					FollowerId: "87654321-4321-4321-210987654321",
				},
			},
			want: true,
		},
		{
			name: "validation error: UserId.min_len",
			args: args{
				req: &pb.FollowRequest{
					UserId:     "",
					FollowerId: "87654321-4321-4321-210987654321",
				},
			},
			want: false,
		},
		{
			name: "validation error: FollowerId.min_len",
			args: args{
				req: &pb.FollowRequest{
					UserId:     "12345678-1234-1234-123456789012",
					FollowerId: "",
				},
			},
			want: false,
		},
		{
			name: "validation error: FollowerId.unique",
			args: args{
				req: &pb.FollowRequest{
					UserId:     "12345678-1234-1234-123456789012",
					FollowerId: "12345678-1234-1234-123456789012",
				},
			},
			want: false,
		},
	}

	for _, tt := range testCases {
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
		req *pb.UnfollowRequest
	}

	testCases := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "success",
			args: args{
				req: &pb.UnfollowRequest{
					UserId:     "12345678-1234-1234-123456789012",
					FollowerId: "87654321-4321-4321-210987654321",
				},
			},
			want: true,
		},
		{
			name: "validation error: UserId.min_len",
			args: args{
				req: &pb.UnfollowRequest{
					UserId:     "",
					FollowerId: "87654321-4321-4321-210987654321",
				},
			},
			want: false,
		},
		{
			name: "validation error: FollowerId.min_len",
			args: args{
				req: &pb.UnfollowRequest{
					UserId:     "12345678-1234-1234-123456789012",
					FollowerId: "",
				},
			},
			want: false,
		},
		{
			name: "validation error: FollowerId.unique",
			args: args{
				req: &pb.UnfollowRequest{
					UserId:     "12345678-1234-1234-123456789012",
					FollowerId: "12345678-1234-1234-123456789012",
				},
			},
			want: false,
		},
	}

	for _, tt := range testCases {
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
