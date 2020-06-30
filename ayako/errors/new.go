package errors

import (
	"fmt"
	"github.com/pkg/errors"
)

// New - create new custom errors
func New(code uint16, msg string, args ...interface{}) *Error {
	err := Error{
		Code:   code,
		Msg:    msg,
		Caller: caller(2),
		Args:   getArgsString(args...),
		Trace:  fmt.Sprintf("'%+v'", errors.New("")),
	}

	return &err
}

// WithCause - create new custom errors with cause
func WithCause(code uint16, msg string, causeErr error, args ...interface{}) *Error {
	err := Error{
		Code:     code,
		Msg:      msg,
		Caller:   caller(2),
		Args:     getArgsString(args...),
		Trace:    fmt.Sprintf("'%+v'", errors.New("")),
		CauseMsg: fmt.Sprintf("'%+v'", causeErr),
		CauseErr: causeErr,
	}

	return &err
}
