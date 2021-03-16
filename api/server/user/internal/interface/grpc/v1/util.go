package v1

import (
	"github.com/calmato/gran-book/api/server/user/internal/domain/exception"
	"github.com/calmato/gran-book/api/server/user/internal/domain/user"
	"golang.org/x/xerrors"
)

func authorization(u *user.User) error {
	switch u.Role {
	case user.AdminRole, user.DeveloperRole, user.OperatorRole:
		return nil
	default:
		err := xerrors.New("This account doesn't have administrator privileges")
		return exception.Forbidden.New(err)
	}
}

func hasAdminRole(u *user.User, uid string) error {
	if u == nil || u.ID == uid {
		err := xerrors.New("This ID belongs to the current user")
		return exception.Forbidden.New(err)
	}

	if u.Role != user.AdminRole {
		err := xerrors.New("This account doesn't have administrator privileges")
		return exception.Forbidden.New(err)
	}

	return nil
}
