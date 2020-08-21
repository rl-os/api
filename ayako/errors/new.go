package errors

import (
	"fmt"

	"github.com/pkg/errors"
)

// New - create new custom errors
func New(id string, code int, msg string, args ...interface{}) *Error {
	err := Error{
		Id:     id,
		Code:   code,
		Msg:    msg,
		Caller: caller(2),
		Args:   getArgsString(args...),
		Trace:  fmt.Sprintf("'%+v'", errors.New("")),
	}

	return &err
}

// WithCause - create new custom errors with cause
func WithCause(id string, code int, msg string, causeErr error, args ...interface{}) *Error {
	return New(id, code, msg, args).WithCause(causeErr)
}
