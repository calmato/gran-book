package service

import (
	"context"
	"time"

	"github.com/calmato/gran-book/api/server/user/internal/domain"
	"github.com/calmato/gran-book/api/server/user/internal/domain/exception"
	"github.com/calmato/gran-book/api/server/user/internal/domain/user"
	"github.com/google/uuid"
	"golang.org/x/xerrors"
)

type userService struct {
	userDomainValidation user.Validation
	userRepository       user.Repository
	userUploader         user.Uploader
}

// NewUserService - UserServiceの生成
func NewUserService(udv user.Validation, ur user.Repository, uu user.Uploader) user.Service {
	return &userService{
		userDomainValidation: udv,
		userRepository:       ur,
		userUploader:         uu,
	}
}

func (s *userService) Authentication(ctx context.Context) (string, error) {
	uid, err := s.userRepository.Authentication(ctx)
	if err != nil {
		return "", err
	}

	return uid, nil
}

func (s *userService) List(ctx context.Context, q *domain.ListQuery) ([]*user.User, int64, error) {
	us, err := s.userRepository.List(ctx, q)
	if err != nil {
		return nil, 0, err
	}

	count, err := s.userRepository.ListCount(ctx, q)
	if err != nil {
		return nil, 0, err
	}

	return us, count, nil
}

func (s *userService) ListFollow(ctx context.Context, q *domain.ListQuery) ([]*user.User, int64, error) {
	us, err := s.userRepository.ListFollow(ctx, q)
	if err != nil {
		return nil, 0, err
	}

	count, err := s.userRepository.ListFollowCount(ctx, q)
	if err != nil {
		return nil, 0, err
	}

	return us, count, nil
}

func (s *userService) ListFollower(ctx context.Context, q *domain.ListQuery) ([]*user.User, int64, error) {
	us, err := s.userRepository.ListFollower(ctx, q)
	if err != nil {
		return nil, 0, err
	}

	count, err := s.userRepository.ListFollowCount(ctx, q)
	if err != nil {
		return nil, 0, err
	}

	return us, count, nil
}

func (s *userService) ListFriendCount(ctx context.Context, u *user.User) (int64, int64, error) {
	if u == nil {
		err := xerrors.New("User is nil")
		return 0, 0, exception.NotFound.New(err)
	}

	followsQuery := &domain.ListQuery{
		Conditions: []*domain.QueryCondition{
			{
				Field:    "follow_id",
				Operator: "==",
				Value:    u.ID,
			},
		},
	}

	followersQuery := &domain.ListQuery{
		Conditions: []*domain.QueryCondition{
			{
				Field:    "follower_id",
				Operator: "==",
				Value:    u.ID,
			},
		},
	}

	followsCount, err := s.userRepository.ListFollowCount(ctx, followsQuery)
	if err != nil {
		return 0, 0, err
	}

	followersCount, err := s.userRepository.ListFollowCount(ctx, followersQuery)
	if err != nil {
		return 0, 0, err
	}

	return followsCount, followersCount, nil
}

func (s *userService) Show(ctx context.Context, uid string) (*user.User, error) {
	u, err := s.userRepository.Show(ctx, uid)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (s *userService) ShowFollow(ctx context.Context, id int64) (*user.Follow, error) {
	f, err := s.userRepository.ShowFollow(ctx, id)

	if err != nil {
		return nil, err
	}

	return f, nil
}

func (s *userService) Create(ctx context.Context, u *user.User) error {
	err := s.userDomainValidation.User(ctx, u)
	if err != nil {
		return err
	}

	current := time.Now()

	u.ID = uuid.New().String()
	u.CreatedAt = current
	u.UpdatedAt = current

	return s.userRepository.Create(ctx, u)
}

func (s *userService) CreateFollow(ctx context.Context, f *user.Follow) error {
	err := s.userDomainValidation.Follow(ctx, f)
	if err != nil {
		return err
	}

	current := time.Now()

	f.CreatedAt = current
	f.UpdatedAt = current

	return s.userRepository.CreateFollow(ctx, f)
}

func (s *userService) Update(ctx context.Context, u *user.User) error {
	err := s.userDomainValidation.User(ctx, u)
	if err != nil {
		return err
	}

	u.UpdatedAt = time.Now()

	return s.userRepository.Update(ctx, u)
}

func (s *userService) UpdatePassword(ctx context.Context, uid string, password string) error {
	return s.userRepository.UpdatePassword(ctx, uid, password)
}

func (s *userService) DeleteFollow(ctx context.Context, id int64) error {
	f, err := s.userRepository.ShowFollow(ctx, id)
	if err != nil {
		return err
	}

	if f == nil {
		err := xerrors.New("Follow is nil")
		return exception.NotFound.New(err)
	}

	return s.userRepository.DeleteFollow(ctx, f)
}

func (s *userService) UploadThumbnail(ctx context.Context, uid string, thumbnail []byte) (string, error) {
	return s.userUploader.Thumbnail(ctx, uid, thumbnail)
}

func (s *userService) IsFriend(ctx context.Context, u *user.User, cuid string) (bool, bool, error) {
	isFollow := false
	isFollower := false

	if u == nil {
		err := xerrors.New("User is nil")
		return false, false, exception.NotFound.New(err)
	}

	followsQuery := &domain.ListQuery{
		Limit: 1,
		Conditions: []*domain.QueryCondition{
			{
				Field:    "follow_id",
				Operator: "==",
				Value:    u.ID,
			},
			{
				Field:    "follower_id",
				Operator: "==",
				Value:    cuid,
			},
		},
	}

	followersQuery := &domain.ListQuery{
		Limit: 1,
		Conditions: []*domain.QueryCondition{
			{
				Field:    "follower_id",
				Operator: "==",
				Value:    u.ID,
			},
			{
				Field:    "follow_id",
				Operator: "==",
				Value:    cuid,
			},
		},
	}

	follows, err := s.userRepository.ListFollow(ctx, followsQuery)
	if err != nil {
		return false, false, err
	}

	followers, err := s.userRepository.ListFollower(ctx, followersQuery)
	if err != nil {
		return false, false, err
	}

	if len(follows) > 0 {
		isFollower = true
	}

	if len(followers) > 0 {
		isFollow = true
	}

	return isFollow, isFollower, nil
}
