package c2api

import "errors"

var (
	ErrUnexpectedStatusCode = errors.New("unexpected status code")
	ErrKeyIsNil             = errors.New("key is nil")
	ErrTrustedKeyIsNil      = errors.New("trusted public key is not set")
	ErrNilAuthToken         = errors.New("api authentication token is nil")
)
