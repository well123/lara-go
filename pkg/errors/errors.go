package errors

import "github.com/pkg/errors"

var (
	WithStack = errors.WithStack
)

var (
	ErrInternalServer  = NewResponse(0, 500, "internal server error")
	ErrMethodNotFound  = NewResponse(0, 405, "method not allowed")
	ErrNotFound        = NewResponse(0, 404, "not found")
	ErrInternalToken   = NewResponse(9999, 401, "invalid signature")
	ErrNoPerm          = NewResponse(0, 401, "no permission")
	ErrTooManyRequests = NewResponse(0, 429, "too many requests")
)
