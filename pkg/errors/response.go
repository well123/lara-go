package errors

import "fmt"

type ResponseError struct {
	Code    int
	Message string
	Status  int
	ERR     error
}

func (r *ResponseError) Error() string {
	if r.ERR != nil {
		return r.ERR.Error()
	}
	return r.Message
}

func UnWrapResponse(err error) *ResponseError {
	if responseError, ok := err.(*ResponseError); ok {
		return responseError
	}
	return nil
}

func NewResponse(code, status int, msg string, args ...any) error {
	res := &ResponseError{
		Code:    code,
		Message: msg,
		Status:  status,
		ERR:     nil,
	}
	return res
}

func WarpResponse(err error, code, status int, message string, args ...any) *ResponseError {
	return &ResponseError{
		Code:    code,
		Message: fmt.Sprintf(message, args...),
		Status:  status,
		ERR:     err,
	}
}

func Wrap400Response(err error, message string, args ...any) *ResponseError {
	return WarpResponse(err, 0, 400, message, args...)
}

func Wrap500Response(err error, message string, args ...any) *ResponseError {
	return WarpResponse(err, 0, 500, message, args...)
}
