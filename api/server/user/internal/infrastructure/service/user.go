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
	return s.userRepository.Authentication(ctx)
}

func (s *userService) List(ctx context.Context, q *domain.ListQuery) ([]*user.User, error) {
	return s.userRepository.List(ctx, q)
}

func (s *userService) ListFollow(ctx context.Context, q *domain.ListQuery) ([]*user.Follow, error) {
	return s.userRepository.ListFollow(ctx, q)
}

func (s *userService) ListFollower(ctx context.Context, q *domain.ListQuery) ([]*user.Follower, error) {
	return s.userRepository.ListFollower(ctx, q)
}

func (s *userService) ListCount(ctx context.Context, q *domain.ListQuery) (int64, error) {
	return s.userRepository.ListCount(ctx, q)
}

func (s *userService) ListFriendCount(ctx context.Context, uid string) (int64, int64, error) {
	followQuery := &domain.ListQuery{
		Conditions: []*domain.QueryCondition{
			{
				Field:    "follower_id",
				Operator: "==",
				Value:    uid,
			},
		},
	}

	followerQuery := &domain.ListQuery{
		Conditions: []*domain.QueryCondition{
			{
				Field:    "follow_id",
				Operator: "==",
				Value:    uid,
			},
		},
	}

	followCount, err := s.userRepository.ListRelationshipCount(ctx, followQuery)
	if err != nil {
		return 0, 0, err
	}

	followerCount, err := s.userRepository.ListRelationshipCount(ctx, followerQuery)
	if err != nil {
		return 0, 0, err
	}

	return followCount, followerCount, nil
}

func (s *userService) Show(ctx context.Context, uid string) (*user.User, error) {
	return s.userRepository.Show(ctx, uid)
}

func (s *userService) ShowRelationship(ctx context.Context, id int64) (*user.Relationship, error) {
	return s.userRepository.ShowRelationship(ctx, id)
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

func (s *userService) CreateRelationship(ctx context.Context, r *user.Relationship) error {
	err := s.userDomainValidation.Relationship(ctx, r)
	if err != nil {
		return err
	}

	current := time.Now()

	r.CreatedAt = current
	r.UpdatedAt = current

	return s.userRepository.CreateRelationship(ctx, r)
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

func (s *userService) DeleteRelationship(ctx context.Context, id int64) error {
	return s.userRepository.DeleteRelationship(ctx, id)
}

func (s *userService) UploadThumbnail(ctx context.Context, uid string, thumbnail []byte) (string, error) {
	return s.userUploader.Thumbnail(ctx, uid, thumbnail)
}

func (s *userService) IsFriend(ctx context.Context, friendID string, uid string) (bool, bool) {
	isFollow := false
	isFollower := false

	followID, _ := s.userRepository.GetRelationshipIDByUID(ctx, uid, friendID)
	followerID, _ := s.userRepository.GetRelationshipIDByUID(ctx, friendID, uid)

	if followID != 0 {
		isFollow = true
	}

	if followerID != 0 {
		isFollower = true
	}

	return isFollow, isFollower
}
