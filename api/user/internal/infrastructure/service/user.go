package service

import (
	"context"
	"time"

	"github.com/calmato/gran-book/api/user/internal/domain/user"
	"github.com/google/uuid"
)

type userService struct {
	userDomainValidation user.Validation
	userRepository       user.Repository
}

// NewUserService - UserServiceの生成
func NewUserService(udv user.Validation, ur user.Repository) user.Service {
	return &userService{
		userDomainValidation: udv,
		userRepository:       ur,
	}
}

func (s *userService) Authentication(ctx context.Context) (*user.User, error) {
	u, err := s.userRepository.Authentication(ctx)
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
