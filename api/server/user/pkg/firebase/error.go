package firebase

import (
	"fmt"

	"firebase.google.com/go/v4/errorutils"
	ce "github.com/calmato/gran-book/api/server/user/pkg/errors"
)

func ToFirebaseError(err error) error {
	if err == nil {
		return nil
	}

	switch {
	case errorutils.IsInvalidArgument(err),
		errorutils.IsOutOfRange(err):
		return newFirebaseError(err.Error(), ce.ErrBadRequest)
	case errorutils.IsUnauthenticated(err):
		return newFirebaseError(err.Error(), ce.ErrUnauthorized)
	case errorutils.IsPermissionDenied(err):
		return newFirebaseError(err.Error(), ce.ErrForbidden)
	case errorutils.IsNotFound(err):
		return newFirebaseError(err.Error(), ce.ErrNotFound)
	case errorutils.IsConflict(err),
		errorutils.IsAlreadyExists(err):
		return newFirebaseError(err.Error(), ce.ErrConflict)
	case errorutils.IsFailedPrecondition(err):
		return newFirebaseError(err.Error(), ce.ErrPreconditionFailed)
	case errorutils.IsAborted(err),
		errorutils.IsCancelled(err),
		errorutils.IsDataLoss(err),
		errorutils.IsResourceExhausted(err),
		errorutils.IsInternal(err):
		return newFirebaseError(err.Error(), ce.ErrInternal)
	case errorutils.IsUnavailable(err):
		return newFirebaseError(err.Error(), ce.ErrServiceUnabailable)
	case errorutils.IsDeadlineExceeded(err):
		return newFirebaseError(err.Error(), ce.ErrGatewayTimeout)
	case errorutils.IsUnknown(err):
		return newFirebaseError(err.Error(), ce.ErrUnknown)
	default:
		return newFirebaseError(err.Error(), ce.ErrUnknown)
	}
}

func newFirebaseError(str string, err error) error {
	return fmt.Errorf("firebase: %s: %w", str, err)
}
