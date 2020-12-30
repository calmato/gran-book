package service

import (
	"context"
	"strings"
	"time"

	"github.com/calmato/gran-book/api/user/internal/domain/exception"
	"github.com/calmato/gran-book/api/user/internal/domain/user"
	"github.com/google/uuid"
	"golang.org/x/xerrors"
)

type userService struct {
	userRepository user.Repository
}

// NewUserService - UserServiceの生成
func NewUserService(ur user.Repository) user.Service {
	return &userService{
		userRepository: ur,
	}
}

func (s *userService) Create(ctx context.Context, u *user.User) error {
	// TODO: domain validation

	current := time.Now()

	u.ID = uuid.New().String()
	u.Email = strings.ToLower(u.Email)
	u.Gender = 0
	u.Role = 0
	u.CreatedAt = current
	u.UpdatedAt = current

	err := s.userRepository.Create(ctx, u)
	if err != nil {
		err = xerrors.Errorf("Failed to CreateUser for Repository: %w", err)
		return exception.ErrorInDatastore.New(err)
	}

	return nil
}
