package server

import "errors"

var (
	errInvalidUploadRequest = errors.New("server: position is duplicated")
)
