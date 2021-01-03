package service

import (
	"context"
	"time"

	"github.com/calmato/gran-book/api/user/internal/domain/exception"
	"github.com/calmato/gran-book/api/user/internal/domain/user"
	"github.com/google/uuid"
	"golang.org/x/xerrors"
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

func (s *userService) Create(ctx context.Context, u *user.User) error {
	err := s.userDomainValidation.User(ctx, u)
	if err != nil {
		return err
	}

	current := time.Now()

	u.ID = uuid.New().String()
	u.CreatedAt = current
	u.UpdatedAt = current

	err = s.userRepository.Create(ctx, u)
	if err != nil {
		err = xerrors.Errorf("Failed to CreateUser for Repository: %w", err)
		return exception.ErrorInDatastore.New(err)
	}

	return nil
}
