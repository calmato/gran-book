package server

import (
	"context"

	"github.com/calmato/gran-book/api/server/user/internal/application"
	"github.com/calmato/gran-book/api/server/user/internal/domain/user"
	"github.com/calmato/gran-book/api/server/user/internal/interface/validation"
	"github.com/calmato/gran-book/api/server/user/pkg/database"
	"github.com/calmato/gran-book/api/server/user/pkg/datetime"
	pb "github.com/calmato/gran-book/api/server/user/proto"
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
		return nil, errorHandling(err)
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
		return nil, errorHandling(err)
	}

	res := getUserListResponse(us, limit, offset, total)
	return res, nil
}

// ListFollow - フォロー一覧取得
func (s *userServer) ListFollow(ctx context.Context, req *pb.ListFollowRequest) (*pb.FollowListResponse, error) {
	err := s.userRequestValidation.ListFollow(req)
	if err != nil {
		return nil, errorHandling(err)
	}

	limit := int(req.GetLimit())
	offset := int(req.GetOffset())
	fs, total, err := s.userApplication.ListFollow(ctx, req.GetUserId(), limit, offset)
	if err != nil {
		return nil, errorHandling(err)
	}

	res := getFollowListResponse(fs, limit, offset, total)
	return res, nil
}

// ListFollower - フォロワー一覧取得
func (s *userServer) ListFollower(ctx context.Context, req *pb.ListFollowerRequest) (*pb.FollowerListResponse, error) {
	err := s.userRequestValidation.ListFollower(req)
	if err != nil {
		return nil, errorHandling(err)
	}

	limit := int(req.GetLimit())
	offset := int(req.GetOffset())
	fs, total, err := s.userApplication.ListFollower(ctx, req.GetUserId(), limit, offset)
	if err != nil {
		return nil, errorHandling(err)
	}

	res := getFollowerListResponse(fs, limit, offset, total)
	return res, nil
}

// MultiGetUser - ユーザー一覧取得 (ID指定)
func (s *userServer) MultiGetUser(ctx context.Context, req *pb.MultiGetUserRequest) (*pb.UserMapResponse, error) {
	err := s.userRequestValidation.MultiGetUser(req)
	if err != nil {
		return nil, errorHandling(err)
	}

	us, err := s.userApplication.MultiGet(ctx, req.GetUserIds())
	if err != nil {
		return nil, errorHandling(err)
	}

	res := getUserMapResponse(us)
	return res, nil
}

// GetUser - ユーザー取得
func (s *userServer) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.UserResponse, error) {
	err := s.userRequestValidation.GetUser(req)
	if err != nil {
		return nil, errorHandling(err)
	}

	u, err := s.userApplication.Get(ctx, req.GetUserId())
	if err != nil {
		return nil, errorHandling(err)
	}

	res := getUserResponse(u)
	return res, nil
}

// GetUserProfile - プロフィール取得
func (s *userServer) GetUserProfile(ctx context.Context, req *pb.GetUserProfileRequest) (*pb.UserProfileResponse, error) {
	err := s.userRequestValidation.GetUserProfile(req)
	if err != nil {
		return nil, errorHandling(err)
	}

	u, isFollow, isFollower, followCount, followerCount, err := s.userApplication.GetUserProfile(ctx, req.GetUserId())
	if err != nil {
		return nil, errorHandling(err)
	}

	res := getUserProfileResponse(u, isFollow, isFollower, followCount, followerCount)
	return res, nil
}

// Follow - フォロー登録
func (s *userServer) Follow(ctx context.Context, req *pb.FollowRequest) (*pb.UserProfileResponse, error) {
	err := s.userRequestValidation.Follow(req)
	if err != nil {
		return nil, errorHandling(err)
	}

	u,
		isFollow,
		isFollower,
		followCount,
		followerCount,
		err := s.userApplication.Follow(ctx, req.GetUserId(), req.GetFollowerId())
	if err != nil {
		return nil, errorHandling(err)
	}

	res := getUserProfileResponse(u, isFollow, isFollower, followCount, followerCount)
	return res, nil
}

// Unfollow - フォロー解除
func (s *userServer) Unfollow(ctx context.Context, req *pb.UnfollowRequest) (*pb.UserProfileResponse, error) {
	err := s.userRequestValidation.Unfollow(req)
	if err != nil {
		return nil, errorHandling(err)
	}

	u,
		isFollow,
		isFollower,
		followCount,
		followerCount,
		err := s.userApplication.Unfollow(ctx, req.GetUserId(), req.GetFollowerId())
	if err != nil {
		return nil, errorHandling(err)
	}

	res := getUserProfileResponse(u, isFollow, isFollower, followCount, followerCount)
	return res, nil
}

func getUserListResponse(us []*user.User, limit, offset, total int) *pb.UserListResponse {
	users := make([]*pb.UserListResponse_User, len(us))
	for i, u := range us {
		user := &pb.UserListResponse_User{
			Id:               u.ID,
			Username:         u.Username,
			Gender:           pb.Gender(u.Gender),
			Email:            u.Email,
			PhoneNumber:      u.PhoneNumber,
			Role:             pb.Role(u.Role),
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

	return &pb.UserListResponse{
		Users:  users,
		Limit:  int64(limit),
		Offset: int64(offset),
		Total:  int64(total),
	}
}

func getFollowListResponse(fs []*user.Follow, limit, offset, total int) *pb.FollowListResponse {
	follows := make([]*pb.FollowListResponse_Follow, len(fs))
	for i, f := range fs {
		follow := &pb.FollowListResponse_Follow{
			Id:               f.FollowID,
			Username:         f.Username,
			ThumbnailUrl:     f.ThumbnailURL,
			SelfIntroduction: f.SelfIntroduction,
			IsFollow:         f.IsFollowing,
		}

		follows[i] = follow
	}

	return &pb.FollowListResponse{
		Follows: follows,
		Limit:   int64(limit),
		Offset:  int64(offset),
		Total:   int64(total),
	}
}

func getFollowerListResponse(fs []*user.Follower, limit, offset, total int) *pb.FollowerListResponse {
	followers := make([]*pb.FollowerListResponse_Follower, len(fs))
	for i, f := range fs {
		follower := &pb.FollowerListResponse_Follower{
			Id:               f.FollowerID,
			Username:         f.Username,
			ThumbnailUrl:     f.ThumbnailURL,
			SelfIntroduction: f.SelfIntroduction,
			IsFollow:         f.IsFollowing,
		}

		followers[i] = follower
	}

	return &pb.FollowerListResponse{
		Followers: followers,
		Limit:     int64(limit),
		Offset:    int64(offset),
		Total:     int64(total),
	}
}

func getUserMapResponse(us []*user.User) *pb.UserMapResponse {
	users := map[string]*pb.UserMapResponse_User{}
	for _, u := range us {
		user := &pb.UserMapResponse_User{
			Id:               u.ID,
			Username:         u.Username,
			Gender:           pb.Gender(u.Gender),
			Email:            u.Email,
			PhoneNumber:      u.PhoneNumber,
			Role:             pb.Role(u.Role),
			ThumbnailUrl:     u.ThumbnailURL,
			SelfIntroduction: u.SelfIntroduction,
			LastName:         u.LastName,
			FirstName:        u.FirstName,
			LastNameKana:     u.LastNameKana,
			FirstNameKana:    u.FirstNameKana,
			CreatedAt:        datetime.TimeToString(u.CreatedAt),
			UpdatedAt:        datetime.TimeToString(u.UpdatedAt),
		}

		users[u.ID] = user
	}

	return &pb.UserMapResponse{
		Users: users,
	}
}

func getUserResponse(u *user.User) *pb.UserResponse {
	return &pb.UserResponse{
		Id:               u.ID,
		Username:         u.Username,
		Gender:           pb.Gender(u.Gender),
		Email:            u.Email,
		PhoneNumber:      u.PhoneNumber,
		Role:             pb.Role(u.Role),
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

func getUserProfileResponse(
	u *user.User, isFollow, isFollower bool, followCount, followerCount int,
) *pb.UserProfileResponse {
	return &pb.UserProfileResponse{
		Id:               u.ID,
		Username:         u.Username,
		ThumbnailUrl:     u.ThumbnailURL,
		SelfIntroduction: u.SelfIntroduction,
		IsFollow:         isFollow,
		IsFollower:       isFollower,
		FollowCount:      int64(followCount),
		FollowerCount:    int64(followerCount),
	}
}
