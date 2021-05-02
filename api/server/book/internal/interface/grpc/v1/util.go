package v1

import (
	"github.com/calmato/gran-book/api/server/book/internal/domain/exception"
	"golang.org/x/xerrors"
)

func isCurrentUser(cuid string, uid string) error {
	if cuid != uid {
		err := xerrors.New("This ID belongs to the current user")
		return exception.Forbidden.New(err)
	}

	return nil
}
