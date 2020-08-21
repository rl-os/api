package errors

import (
	"fmt"
	"net/http"
	"strings"
)

// Interface assertion
// this code is necessary to check for compatibility
// with the Error type so that it would be
// more convenient to find differences
var _ error = (*Error)(nil)

// Error represent custom errors
// with the ability to add more information
// about the errors and its location
type Error struct {
	Id       string // unique errors code
	Code     int    // status code
	Msg      string // message body
	Args     string // argument string with additional information
	CauseErr error  // reason errors
	CauseMsg string
	Trace    string // call stack
	Caller   string // file, line and name of the method in which the errors occurred
}

type ResponseFormat struct {
	ErrorID          string `json:"error_id"`
	ErrorDescription string `json:"error_description"`
	Message          string `json:"message"`
}

func (e *Error) WithCause(err error, messages ...string) *Error {
	e.CauseErr = err
	e.CauseMsg = strings.Join(messages, "; ")
	return e
}

// Error print custom errors
func (e *Error) Error() string {
	mes := fmt.Sprintf("%s %s", e.Id, e.Msg)
	if e.Args != "" {
		mes = fmt.Sprintf("%s with %s", mes, e.Args)
	}

	if e.CauseMsg != "" {
		mes = fmt.Sprintf("%s causemes %s", mes, e.CauseMsg)
	}

	return mes
}

// responseFormat format an errors and return interface that must be show user
func (e *Error) ResponseFormat() ResponseFormat {
	return ResponseFormat{
		e.Id,
		e.Msg,
		http.StatusText(e.Code),
	}
}
