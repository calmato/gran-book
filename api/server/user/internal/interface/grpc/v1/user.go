package v1

import (
	"context"

	"github.com/calmato/gran-book/api/server/user/internal/application"
	"github.com/calmato/gran-book/api/server/user/internal/application/input"
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
	_, err := s.AuthApplication.Authentication(ctx)
	if err != nil {
		return nil, errorHandling(err)
	}

	in := &input.GetUserProfile{
		ID: req.Id,
	}

	u, out, err := s.UserApplication.GetProfile(ctx, in)
	if err != nil {
		return nil, errorHandling(err)
	}

	res := &pb.UserProfileResponse{
		Id:               u.ID,
		Username:         u.Username,
		ThumbnailUrl:     u.ThumbnailURL,
		SelfIntroduction: u.SelfIntroduction,
		IsFollow:         out.IsFollow,
		IsFollower:       out.IsFollower,
		FollowCount:      int32(out.FollowsTotal),
		FollowerCount:    int32(out.FollowersTotal),
	}

	return res, nil
}
