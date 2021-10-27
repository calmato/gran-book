package exception

import "firebase.google.com/go/v4/errorutils"

func ToFirebaseError(err error) error {
	if err == nil {
		return nil
	}

	switch {
	case errorutils.IsInvalidArgument(err),
		errorutils.IsOutOfRange(err):
		return ErrInvalidArgument.New(err)
	case errorutils.IsUnauthenticated(err):
		return ErrUnauthorized.New(err)
	case errorutils.IsPermissionDenied(err):
		return ErrForbidden.New(err)
	case errorutils.IsNotFound(err):
		return ErrNotFound.New(err)
	case errorutils.IsConflict(err),
		errorutils.IsAlreadyExists(err):
		return ErrConflict.New(err)
	case errorutils.IsFailedPrecondition(err):
		return ErrPreconditionFailed.New(err)
	case errorutils.IsAborted(err),
		errorutils.IsCancelled(err),
		errorutils.IsDataLoss(err),
		errorutils.IsResourceExhausted(err),
		errorutils.IsInternal(err):
		return ErrInDatastore.New(err)
	case errorutils.IsUnavailable(err):
		return ErrNotImplemented.New(err)
	case errorutils.IsDeadlineExceeded(err):
		return ErrGatewayTimeout.New(err)
	case errorutils.IsUnknown(err):
		return ErrUnknown.New(err)
	default:
		return ErrUnknown.New(err)
	}
}
