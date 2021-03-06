package v1

import (
	"context"

	"github.com/calmato/gran-book/api/server/user/internal/application"
	"github.com/calmato/gran-book/api/server/user/internal/application/input"
	"github.com/calmato/gran-book/api/server/user/internal/application/output"
	"github.com/calmato/gran-book/api/server/user/internal/domain/user"
	"github.com/calmato/gran-book/api/server/user/lib/datetime"
	pb "github.com/calmato/gran-book/api/server/user/proto"
)

// UserServer - Userインターフェースの構造体
type UserServer struct {
	pb.UnimplementedUserServiceServer
	AuthApplication application.AuthApplication
	UserApplication application.UserApplication
}

// ListUser - ユーザー一覧取得
func (s *UserServer) ListUser(ctx context.Context, req *pb.ListUserRequest) (*pb.UserListResponse, error) {
	cu, err := s.AuthApplication.Authentication(ctx)
	if err != nil {
		return nil, errorHandling(err)
	}

	err = authorization(cu)
	if err != nil {
		return nil, errorHandling(err)
	}

	in := &input.ListUser{
		Limit:  int(req.GetLimit()),
		Offset: int(req.GetOffset()),
	}

	if o := req.GetOrder(); o != nil {
		in.By = o.GetBy()
		in.Direction = o.GetDirection()
	}

	us, out, err := s.UserApplication.List(ctx, in)
	if err != nil {
		return nil, errorHandling(err)
	}

	res := getUserListResponse(us, out)
	return res, nil
}

// ListUserByUserIds - ユーザー一覧取得 (ID指定)
func (s *UserServer) ListUserByUserIds(
	ctx context.Context, req *pb.ListUserByUserIdsRequest,
) (*pb.UserListResponse, error) {
	_, err := s.AuthApplication.Authentication(ctx)
	if err != nil {
		return nil, errorHandling(err)
	}

	in := &input.ListUserByUserIDs{
		UserIDs: req.GetUserIds(),
	}

	us, err := s.UserApplication.ListByUserIDs(ctx, in)
	if err != nil {
		return nil, errorHandling(err)
	}

	out := &output.ListQuery{
		Limit:  0,
		Offset: 0,
		Total:  len(us),
	}

	res := getUserListResponse(us, out)
	return res, nil
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

// SearchUser - ユーザー一覧取得
func (s *UserServer) SearchUser(ctx context.Context, req *pb.SearchUserRequest) (*pb.UserListResponse, error) {
	cu, err := s.AuthApplication.Authentication(ctx)
	if err != nil {
		return nil, errorHandling(err)
	}

	err = authorization(cu)
	if err != nil {
		return nil, errorHandling(err)
	}

	in := &input.SearchUser{
		Limit:  int(req.GetLimit()),
		Offset: int(req.GetOffset()),
	}

	if o := req.GetOrder(); o != nil {
		in.By = o.GetBy()
		in.Direction = o.GetDirection()
	}

	if s := req.GetSearch(); s != nil {
		in.Field = s.GetField()
		in.Value = s.GetValue()
	}

	us, out, err := s.UserApplication.Search(ctx, in)
	if err != nil {
		return nil, errorHandling(err)
	}

	res := getUserListResponse(us, out)
	return res, nil
}

func (s *UserServer) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.UserResponse, error) {
	_, err := s.AuthApplication.Authentication(ctx)
	if err != nil {
		return nil, errorHandling(err)
	}

	u, err := s.UserApplication.Show(ctx, req.GetId())
	if err != nil {
		return nil, errorHandling(err)
	}

	res := getUserResponse(u)
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

func getUserResponse(u *user.User) *pb.UserResponse {
	return &pb.UserResponse{
		Id:               u.ID,
		Username:         u.Username,
		Email:            u.Email,
		PhoneNumber:      u.PhoneNumber,
		Role:             int32(u.Role),
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

func getUserListResponse(us []*user.User, out *output.ListQuery) *pb.UserListResponse {
	users := make([]*pb.UserListResponse_User, len(us))
	for i, u := range us {
		user := &pb.UserListResponse_User{
			Id:               u.ID,
			Username:         u.Username,
			Email:            u.Email,
			PhoneNumber:      u.PhoneNumber,
			Role:             int32(u.Role),
			ThumbnailUrl:     u.ThumbnailURL,
			SelfIntroduction: u.SelfIntroduction,
			LastName:         u.LastName,
			FirstName:        u.FirstName,
			LastNameKana:     u.LastNameKana,
			FirstNameKana:    u.FirstNameKana,
			CreatedAt:        datetime.TimeToString(u.CreatedAt),
			UpdatedAt:        datetime.TimeToString(u.UpdatedAt),
		}

		users[i] = user
	}

	res := &pb.UserListResponse{
		Users:  users,
		Limit:  int64(out.Limit),
		Offset: int64(out.Offset),
		Total:  int64(out.Total),
	}

	if out.Order != nil {
		order := &pb.UserListResponse_Order{
			By:        out.Order.By,
			Direction: out.Order.Direction,
		}

		res.Order = order
	}

	return res
}
