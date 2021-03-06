package v1

import (
	"context"

	"github.com/calmato/gran-book/api/server/user/internal/application"
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

// GetUserProfile - ユーザプロフィール取得
func (s *UserServer) GetUserProfile(
	ctx context.Context, req *pb.GetUserProfileRequest,
) (*pb.UserProfileResponse, error) {
	cu, err := s.AuthApplication.Authentication(ctx)
	if err != nil {
		return nil, errorHandling(err)
	}

	u, out, err := s.UserApplication.GetUserProfile(ctx, req.Id, cu.ID)
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

	u, out, err := s.UserApplication.RegisterFollow(ctx, req.Id, cu.ID)
	if err != nil {
		return nil, errorHandling(err)
	}

	res := getUserProfileResponse(u, out)
	return res, nil
}

// // UnregisterFollow - フォローの解除
// func (s *UserServer) UnregisterFollow(
// 	ctx context.Context, req *pb.UnregisterFollowRequest,
// ) (*pb.UserProfileResponse, error) {
// 	_, err := s.AuthApplication.Authentication(ctx)
// 	if err != nil {
// 		return nil, errorHandling(err)
// 	}

// 	u, out, err := s.UserApplication.UnregisterFollow(ctx, req.Id)
// 	if err != nil {
// 		return nil, errorHandling(err)
// 	}

// 	res := getUserProfileResponse(u, out)
// 	return res, nil
// }

func getUserProfileResponse(u *user.User, out *output.UserProfile) *pb.UserProfileResponse {
	return &pb.UserProfileResponse{
		Id:               u.ID,
		Username:         u.Username,
		ThumbnailUrl:     u.ThumbnailURL,
		SelfIntroduction: u.SelfIntroduction,
		IsFollow:         out.IsFollow,
		IsFollower:       out.IsFollower,
		FollowCount:      out.FollowCount,
		FollowerCount:    out.FollowerCount,
	}
}
