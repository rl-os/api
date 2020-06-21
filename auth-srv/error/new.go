package error

import (
	"fmt"
	"github.com/pkg/errors"
)

// New - create new custom error
func New(id uint16, msg string, args ...interface{}) *Error {
	err := Error{
		ID:     id,
		Msg:    msg,
		Caller: caller(2),
		Args:   getArgsString(args...),
		Trace:  fmt.Sprintf("'%+v'", errors.New("")),
	}

	return &err
}

// WithCause - create new custom error with cause
func WithCause(id uint16, msg string, causeErr error, args ...interface{}) *Error {
	err := Error{
		ID:       id,
		Msg:      msg,
		Caller:   caller(2),
		Args:     getArgsString(args...),
		Trace:    fmt.Sprintf("'%+v'", errors.New("")),
		CauseMsg: fmt.Sprintf("'%+v'", causeErr),
		CauseErr: causeErr,
	}

	return &err
}
