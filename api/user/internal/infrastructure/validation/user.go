package validation

import (
	"context"

	"github.com/calmato/gran-book/api/user/internal/domain/exception"
	"github.com/calmato/gran-book/api/user/internal/domain/user"
	"golang.org/x/xerrors"
)

type userDomainValidation struct {
	userRepository user.Repository
}

// NewUserDomainValidation - User関連のドメインバリデータ
func NewUserDomainValidation(ur user.Repository) user.Validation {
	return &userDomainValidation{
		userRepository: ur,
	}
}

func (v *userDomainValidation) User(ctx context.Context, u *user.User) error {
	err := v.uniqueCheckEmail(ctx, u.ID, u.Email)
	if err != nil {
		ves := []*exception.ValidationError{
			{
				Field:   "email",
				Message: exception.CustomUniqueMessage,
			},
		}

		return exception.Conflict.New(err, ves...)
	}

	return nil
}

func (v *userDomainValidation) uniqueCheckEmail(ctx context.Context, id string, email string) error {
	if email == "" {
		return nil
	}

	uid, _ := v.userRepository.GetUIDByEmail(ctx, email)
	if uid == "" || id == uid {
		return nil
	}

	return xerrors.New("This email is already exists.")
}
