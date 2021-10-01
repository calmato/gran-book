package repository

import (
	"errors"

	"github.com/calmato/gran-book/api/server/user/internal/domain/exception"
	"github.com/calmato/gran-book/api/server/user/pkg/database"
	ce "github.com/calmato/gran-book/api/server/user/pkg/errors"
	"github.com/calmato/gran-book/api/server/user/pkg/firebase"
)

const (
	userTable         = "users"
	relationshipTable = "relationships"
)

var (
	errNotExistsUserInfo            = errors.New("repository: user info is not exists in firebase authentication")
	errNotExistsAuthorizationHeader = errors.New("repository: authorization header is not contain")
	errInvalidAuthorizationHeader   = errors.New("repository: authorization header is invalid")
	errNotVerifyToken               = errors.New("repository: id token is not verify")
)

func toDBError(err error) error {
	if err == nil {
		return nil
	}

	return newRepositoryError(database.ToDBError(err))
}

func toFirebaseError(err error) error {
	if err == nil {
		return nil
	}

	return newRepositoryError(firebase.ToFirebaseError(err))
}

func newRepositoryError(err error) error {
	switch {
	case errors.Is(err, ce.ErrBadRequest):
		return exception.InvalidDomainValidation.New(err)
	case errors.Is(err, ce.ErrUnauthorized),
		errors.Is(err, errNotExistsUserInfo),
		errors.Is(err, errNotExistsAuthorizationHeader),
		errors.Is(err, errInvalidAuthorizationHeader),
		errors.Is(err, errNotVerifyToken):
		return exception.Unauthorized.New(err)
	case errors.Is(err, ce.ErrForbidden),
		errors.Is(err, ce.ErrPreconditionFailed):
		return exception.Forbidden.New(err)
	case errors.Is(err, ce.ErrNotFound):
		return exception.NotFound.New(err)
	case errors.Is(err, ce.ErrConflict):
		return exception.Conflict.New(err)
	case errors.Is(err, ce.ErrTooManyRequests),
		errors.Is(err, ce.ErrInternal),
		errors.Is(err, ce.ErrServiceUnabailable),
		errors.Is(err, ce.ErrGatewayTimeout),
		errors.Is(err, ce.ErrUnknown):
		return exception.ErrorInDatastore.New(err)
	default:
		return exception.ErrorInDatastore.New(err)
	}
}
