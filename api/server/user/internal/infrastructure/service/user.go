package service

import (
	"context"
	"time"

	"github.com/calmato/gran-book/api/server/user/internal/domain"
	"github.com/calmato/gran-book/api/server/user/internal/domain/user"
	"github.com/google/uuid"
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

func (s *userService) ListFriendsCount(ctx context.Context, u *user.User) (int64, int64, error) {
	followsQuery := &domain.ListQuery{
		Conditions: []*domain.QueryCondition{
			{
				Field:    "follower_id",
				Operator: "==",
				Value:    u.ID,
			},
		},
	}

	followersQuery := &domain.ListQuery{
		Conditions: []*domain.QueryCondition{
			{
				Field:    "follow_id",
				Operator: "==",
				Value:    u.ID,
			},
		},
	}

	followsCount, err := s.userRepository.ListFollowersCount(ctx, followsQuery)
	if err != nil {
		return 0, 0, err
	}

	followersCount, err := s.userRepository.ListFollowersCount(ctx, followersQuery)
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

func (s *userService) UploadThumbnail(ctx context.Context, uid string, thumbnail []byte) (string, error) {
	return s.userUploader.Thumbnail(ctx, uid, thumbnail)
}
