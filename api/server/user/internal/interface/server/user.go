package server

import (
	"context"

	"github.com/calmato/gran-book/api/server/user/internal/application"
	"github.com/calmato/gran-book/api/server/user/internal/domain/user"
	"github.com/calmato/gran-book/api/server/user/internal/interface/validation"
	"github.com/calmato/gran-book/api/server/user/pkg/database"
	pb "github.com/calmato/gran-book/api/server/user/proto/service/user"
)

type userServer struct {
	pb.UnimplementedUserServiceServer
	userRequestValidation validation.UserRequestValidation
	userApplication       application.UserApplication
}

func NewUserServer(urv validation.UserRequestValidation, ua application.UserApplication) pb.UserServiceServer {
	return &userServer{
		userRequestValidation: urv,
		userApplication:       ua,
	}
}

// ListUser - ユーザー一覧取得
func (s *userServer) ListUser(ctx context.Context, req *pb.ListUserRequest) (*pb.UserListResponse, error) {
	err := s.userRequestValidation.ListUser(req)
	if err != nil {
		return nil, toGRPCError(err)
	}

	limit := int(req.GetLimit())
	offset := int(req.GetOffset())
	q := &database.ListQuery{
		Limit:      limit,
		Offset:     offset,
		Conditions: []*database.ConditionQuery{},
	}

	if req.GetSearch() != nil {
		c := &database.ConditionQuery{
			Field:    req.GetSearch().GetField(),
			Operator: "LIKE",
			Value:    req.GetSearch().GetValue(),
		}

		q.Conditions = append(q.Conditions, c)
	}

	if req.GetOrder() != nil {
		o := &database.OrderQuery{
			Field:   req.GetOrder().GetField(),
			OrderBy: int(req.GetOrder().GetOrderBy()),
		}

		q.Order = o
	}

	us, total, err := s.userApplication.List(ctx, q)
	if err != nil {
		return nil, toGRPCError(err)
	}

	res := getUserListResponse(us, limit, offset, total)
	return res, nil
}

// ListFollow - フォロー一覧取得
func (s *userServer) ListFollow(ctx context.Context, req *pb.ListFollowRequest) (*pb.FollowListResponse, error) {
	// TODO: remove
	u, err := s.userApplication.Authentication(ctx)
	if err != nil {
		return nil, toGRPCError(err)
	}

	err = s.userRequestValidation.ListFollow(req)
	if err != nil {
		return nil, toGRPCError(err)
	}

	limit := int(req.GetLimit())
	offset := int(req.GetOffset())
	fs, total, err := s.userApplication.ListFollow(ctx, u.ID, req.GetUserId(), limit, offset)
	if err != nil {
		return nil, toGRPCError(err)
	}

	res := getFollowListResponse(fs, limit, offset, total)
	return res, nil
}

// ListFollower - フォロワー一覧取得
func (s *userServer) ListFollower(ctx context.Context, req *pb.ListFollowerRequest) (*pb.FollowerListResponse, error) {
	// TODO: remove
	u, err := s.userApplication.Authentication(ctx)
	if err != nil {
		return nil, toGRPCError(err)
	}

	err = s.userRequestValidation.ListFollower(req)
	if err != nil {
		return nil, toGRPCError(err)
	}

	limit := int(req.GetLimit())
	offset := int(req.GetOffset())
	fs, total, err := s.userApplication.ListFollower(ctx, u.ID, req.GetUserId(), limit, offset)
	if err != nil {
		return nil, toGRPCError(err)
	}

	res := getFollowerListResponse(fs, limit, offset, total)
	return res, nil
}

// MultiGetUser - ユーザー一覧取得 (ID指定)
func (s *userServer) MultiGetUser(ctx context.Context, req *pb.MultiGetUserRequest) (*pb.UserListResponse, error) {
	err := s.userRequestValidation.MultiGetUser(req)
	if err != nil {
		return nil, toGRPCError(err)
	}

	us, err := s.userApplication.MultiGet(ctx, req.GetUserIds())
	if err != nil {
		return nil, toGRPCError(err)
	}

	res := getUserListResponse(us, len(req.GetUserIds()), 0, len(us))
	return res, nil
}

// GetUser - ユーザー取得
func (s *userServer) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.UserResponse, error) {
	err := s.userRequestValidation.GetUser(req)
	if err != nil {
		return nil, toGRPCError(err)
	}

	u, err := s.userApplication.Get(ctx, req.GetUserId())
	if err != nil {
		return nil, toGRPCError(err)
	}

	res := getUserResponse(u)
	return res, nil
}

// GetUserProfile - プロフィール取得
func (s *userServer) GetUserProfile(
	ctx context.Context, req *pb.GetUserProfileRequest,
) (*pb.UserProfileResponse, error) {
	// TODO: remove
	cu, err := s.userApplication.Authentication(ctx)
	if err != nil {
		return nil, toGRPCError(err)
	}

	err = s.userRequestValidation.GetUserProfile(req)
	if err != nil {
		return nil, toGRPCError(err)
	}

	u, err := s.userApplication.GetUserProfile(ctx, cu.ID, req.GetUserId())
	if err != nil {
		return nil, toGRPCError(err)
	}

	res := getUserProfileResponse(u)
	return res, nil
}

// Follow - フォロー登録
func (s *userServer) Follow(ctx context.Context, req *pb.FollowRequest) (*pb.UserProfileResponse, error) {
	err := s.userRequestValidation.Follow(req)
	if err != nil {
		return nil, toGRPCError(err)
	}

	u, err := s.userApplication.Follow(ctx, req.GetUserId(), req.GetFollowerId())
	if err != nil {
		return nil, toGRPCError(err)
	}

	res := getUserProfileResponse(u)
	return res, nil
}

// Unfollow - フォロー解除
func (s *userServer) Unfollow(ctx context.Context, req *pb.UnfollowRequest) (*pb.UserProfileResponse, error) {
	err := s.userRequestValidation.Unfollow(req)
	if err != nil {
		return nil, toGRPCError(err)
	}

	u, err := s.userApplication.Unfollow(ctx, req.GetUserId(), req.GetFollowerId())
	if err != nil {
		return nil, toGRPCError(err)
	}

	res := getUserProfileResponse(u)
	return res, nil
}

func getUserListResponse(us user.Users, limit, offset, total int) *pb.UserListResponse {
	return &pb.UserListResponse{
		Users:  us.Proto(),
		Limit:  int64(limit),
		Offset: int64(offset),
		Total:  int64(total),
	}
}

func getFollowListResponse(fs user.Follows, limit, offset, total int) *pb.FollowListResponse {
	return &pb.FollowListResponse{
		Follows: fs.Proto(),
		Limit:   int64(limit),
		Offset:  int64(offset),
		Total:   int64(total),
	}
}

func getFollowerListResponse(fs user.Followers, limit, offset, total int) *pb.FollowerListResponse {
	return &pb.FollowerListResponse{
		Followers: fs.Proto(),
		Limit:     int64(limit),
		Offset:    int64(offset),
		Total:     int64(total),
	}
}

func getUserResponse(u *user.User) *pb.UserResponse {
	return &pb.UserResponse{
		User: u.Proto(),
	}
}

func getUserProfileResponse(u *user.User) *pb.UserProfileResponse {
	return &pb.UserProfileResponse{
		Profile: u.Profile(),
	}
}
