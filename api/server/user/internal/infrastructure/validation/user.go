package validation

import (
	"context"

	"github.com/calmato/gran-book/api/server/user/internal/domain/exception"
	"github.com/calmato/gran-book/api/server/user/internal/domain/user"
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

func (v *userDomainValidation) Relationship(ctx context.Context, r *user.Relationship) error {
	err := v.uniqueCheckRelationship(ctx, r.ID, r.FollowID, r.FollowerID)
	if err != nil {
		ves := []*exception.ValidationError{
			{
				Field:   "followerId",
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

func (v *userDomainValidation) uniqueCheckRelationship(
	ctx context.Context, id int64, followID string, followerID string,
) error {
	if followID == "" || followerID == "" {
		return nil
	}

	rid, _ := v.userRepository.GetRelationshipIDByUID(ctx, followID, followerID)
	if rid == 0 || rid == id {
		return nil
	}

	return xerrors.New("This user is already following.")
}
