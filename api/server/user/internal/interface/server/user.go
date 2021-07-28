package server

import (
	"context"

	"github.com/calmato/gran-book/api/server/user/internal/domain/user"
	"github.com/calmato/gran-book/api/server/user/internal/interface/validation"
	"github.com/calmato/gran-book/api/server/user/lib/datetime"
	pb "github.com/calmato/gran-book/api/server/user/proto"
)

// UserServer - Userサービスインターフェースの構造体
type UserServer struct {
	pb.UnimplementedUserServiceServer
	userRequestValidation validation.UserRequestValidation
}

// ListUser - ユーザー一覧取得
func (s *UserServer) ListUser(ctx context.Context, req *pb.ListUserRequest) (*pb.UserListResponse, error) {
	err := s.userRequestValidation.ListUser(req)
	if err != nil {
		return nil, errorHandling(err)
	}

	res := getUserListResponse(nil, 0, 0, 0)
	return res, nil
}

// ListFollow - フォロー一覧取得
func (s *UserServer) ListFollow(ctx context.Context, req *pb.ListFollowRequest) (*pb.FollowListResponse, error) {
	err := s.userRequestValidation.ListFollow(req)
	if err != nil {
		return nil, errorHandling(err)
	}

	res := getFollowListResponse(nil, 0, 0, 0)
	return res, nil
}

// ListFollower - フォロワー一覧取得
func (s *UserServer) ListFollower(ctx context.Context, req *pb.ListFollowerRequest) (*pb.FollowerListResponse, error) {
	err := s.userRequestValidation.ListFollower(req)
	if err != nil {
		return nil, errorHandling(err)
	}

	res := getFollowerListResponse(nil, 0, 0, 0)
	return res, nil
}

// MultiGetUser - ユーザー一覧取得 (ID指定)
func (s *UserServer) MultiGetUser(ctx context.Context, req *pb.MultiGetUserRequest) (*pb.UserMapResponse, error) {
	err := s.userRequestValidation.MultiGetUser(req)
	if err != nil {
		return nil, errorHandling(err)
	}

	res := getUserMapResponse(nil)
	return res, nil
}

// GetUser - ユーザー取得
func (s *UserServer) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.UserResponse, error) {
	err := s.userRequestValidation.GetUser(req)
	if err != nil {
		return nil, errorHandling(err)
	}

	res := getUserResponse(nil)
	return res, nil
}

// GetUserProfile - プロフィール取得
func (s *UserServer) GetUserProfile(ctx context.Context, req *pb.GetUserProfileRequest) (*pb.UserProfileResponse, error) {
	err := s.userRequestValidation.GetUserProfile(req)
	if err != nil {
		return nil, errorHandling(err)
	}

	res := getUserProfileResponse(nil)
	return res, nil
}

// RegisterFollow - フォロー登録
func (s *UserServer) RegisterFollow(ctx context.Context, req *pb.RegisterFollowRequest) (*pb.UserProfileResponse, error) {
	err := s.userRequestValidation.RegisterFollow(req)
	if err != nil {
		return nil, errorHandling(err)
	}

	res := getUserProfileResponse(nil)
	return res, nil
}

// UnregisterFollow - フォロー解除
func (s *UserServer) UnregisterFollow(ctx context.Context, req *pb.UnregisterFollowRequest) (*pb.UserProfileResponse, error) {
	err := s.userRequestValidation.UnregisterFollow(req)
	if err != nil {
		return nil, errorHandling(err)
	}

	res := getUserProfileResponse(nil)
	return res, nil
}

func getUserListResponse(us []*user.User, limit, offset, total int) *pb.UserListResponse {
	return &pb.UserListResponse{}
}

func getFollowListResponse(fs []*user.Follow, limit, offset, total int) *pb.FollowListResponse {
	return &pb.FollowListResponse{}
}

func getFollowerListResponse(fs []*user.Follower, limit, offset, total int) *pb.FollowerListResponse {
	return &pb.FollowerListResponse{}
}

func getUserMapResponse(us []*user.User) *pb.UserMapResponse {
	return &pb.UserMapResponse{}
}

func getUserResponse(u *user.User) *pb.UserResponse {
	return &pb.UserResponse{
		Id:               u.ID,
		Username:         u.Username,
		Gender:           u.Gender,
		Email:            u.Email,
		PhoneNumber:      u.PhoneNumber,
		Role:             u.Role,
		ThumbnailUrl:     u.ThumbnailURL,
		SelfIntroduction: u.SelfIntroduction,
		LastName:         u.LastName,
		FirstName:        u.FirstName,
		LastNameKana:     u.LastNameKana,
		FirstNameKana:    u.FirstNameKana,
		CreatedAt:        datetime.TimeToString(u.CreatedAt),
		UpdatedAt:        datetime.TimeToString(u.UpdatedAt),
	}
}

func getUserProfileResponse(u *user.User) *pb.UserProfileResponse {
	return &pb.UserProfileResponse{
		Id:               u.ID,
		Username:         u.Username,
		ThumbnailUrl:     u.ThumbnailURL,
		SelfIntroduction: u.SelfIntroduction,
		// IsFollow:         out.IsFollow,
		// IsFollower:       out.IsFollower,
		// FollowCount:      int64(out.FollowCount),
		// FollowerCount:    int64(out.FollowerCount),
	}
}
