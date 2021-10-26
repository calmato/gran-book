package repository

import "errors"

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
