package v1

import (
	"context"

	"github.com/calmato/gran-book/api/server/user/internal/application"
	"github.com/calmato/gran-book/api/server/user/internal/application/input"
	"github.com/calmato/gran-book/api/server/user/internal/application/output"
	"github.com/calmato/gran-book/api/server/user/internal/domain/user"
	pb "github.com/calmato/gran-book/api/server/user/proto"
)

// UserServer - Userインターフェースの構造体
type UserServer struct {
	pb.UnimplementedUserServiceServer
	AuthApplication application.AuthApplication
	UserApplication application.UserApplication
}

// ListFollow - フォロー一覧取得
func (s *UserServer) ListFollow(ctx context.Context, req *pb.ListFollowRequest) (*pb.FollowListResponse, error) {
	cu, err := s.AuthApplication.Authentication(ctx)
	if err != nil {
		return nil, errorHandling(err)
	}

	in := &input.ListFollow{
		Limit:  int(req.GetLimit()),
		Offset: int(req.GetOffset()),
	}

	fs, out, err := s.UserApplication.ListFollow(ctx, in, req.GetId(), cu.ID)
	if err != nil {
		return nil, errorHandling(err)
	}

	us := make([]*pb.FollowListResponse_User, len(fs))
	for i, f := range fs {
		u := &pb.FollowListResponse_User{
			Id:           f.FollowerID,
			Username:     f.Username,
			ThumbnailUrl: f.ThumbnailURL,
			IsFollow:     f.IsFollow,
		}

		us[i] = u
	}

	res := &pb.FollowListResponse{
		Users:  us,
		Limit:  int64(out.Limit),
		Offset: int64(out.Offset),
		Total:  int64(out.Total),
	}

	return res, nil
}

// ListFollower - フォロワー一覧取得
func (s *UserServer) ListFollower(ctx context.Context, req *pb.ListFollowerRequest) (*pb.FollowerListResponse, error) {
	cu, err := s.AuthApplication.Authentication(ctx)
	if err != nil {
		return nil, errorHandling(err)
	}

	in := &input.ListFollower{
		Limit:  int(req.GetLimit()),
		Offset: int(req.GetOffset()),
	}

	fs, out, err := s.UserApplication.ListFollower(ctx, in, req.GetId(), cu.ID)
	if err != nil {
		return nil, errorHandling(err)
	}

	us := make([]*pb.FollowerListResponse_User, len(fs))
	for i, f := range fs {
		u := &pb.FollowerListResponse_User{
			Id:           f.FollowID,
			Username:     f.Username,
			ThumbnailUrl: f.ThumbnailURL,
			IsFollow:     f.IsFollow,
		}

		us[i] = u
	}

	res := &pb.FollowerListResponse{
		Users:  us,
		Limit:  int64(out.Limit),
		Offset: int64(out.Offset),
		Total:  int64(out.Total),
	}

	return res, nil
}

// GetUserProfile - ユーザプロフィール取得
func (s *UserServer) GetUserProfile(
	ctx context.Context, req *pb.GetUserProfileRequest,
) (*pb.UserProfileResponse, error) {
	cu, err := s.AuthApplication.Authentication(ctx)
	if err != nil {
		return nil, errorHandling(err)
	}

	u, out, err := s.UserApplication.GetUserProfile(ctx, req.GetId(), cu.ID)
	if err != nil {
		return nil, errorHandling(err)
	}

	res := getUserProfileResponse(u, out)
	return res, nil
}

// RegisterFollow - フォローの追加
func (s *UserServer) RegisterFollow(
	ctx context.Context, req *pb.RegisterFollowRequest,
) (*pb.UserProfileResponse, error) {
	cu, err := s.AuthApplication.Authentication(ctx)
	if err != nil {
		return nil, errorHandling(err)
	}

	u, out, err := s.UserApplication.RegisterFollow(ctx, req.GetId(), cu.ID)
	if err != nil {
		return nil, errorHandling(err)
	}

	res := getUserProfileResponse(u, out)
	return res, nil
}

// UnregisterFollow - フォローの解除
func (s *UserServer) UnregisterFollow(
	ctx context.Context, req *pb.UnregisterFollowRequest,
) (*pb.UserProfileResponse, error) {
	cu, err := s.AuthApplication.Authentication(ctx)
	if err != nil {
		return nil, errorHandling(err)
	}

	u, out, err := s.UserApplication.UnregisterFollow(ctx, req.GetId(), cu.ID)
	if err != nil {
		return nil, errorHandling(err)
	}

	res := getUserProfileResponse(u, out)
	return res, nil
}

func getUserProfileResponse(u *user.User, out *output.UserProfile) *pb.UserProfileResponse {
	return &pb.UserProfileResponse{
		Id:               u.ID,
		Username:         u.Username,
		ThumbnailUrl:     u.ThumbnailURL,
		SelfIntroduction: u.SelfIntroduction,
		IsFollow:         out.IsFollow,
		IsFollower:       out.IsFollower,
		FollowCount:      int64(out.FollowCount),
		FollowerCount:    int64(out.FollowerCount),
	}
}
