package errors

import "fmt"

// Error code consists of 4 digits
// [prefix][code], example 1056 (10 - prefix and 56 - code)
// 10 - internal errors
// 11 - request errors
// ...

// Interface assertion
// this code is necessary to check for compatibility
// with the Error type so that it would be
// more convenient to find differences
var _ error = (*Error)(nil)

// Error represent custom errors
// with the ability to add more information
// about the errors and its location
type Error struct {
	ID       uint16 // unique errors code
	Msg      string // message body
	Args     string // argument string with additional information
	CauseErr error  // reason errors
	CauseMsg string
	Trace    string // call stack
	Caller   string // file, line and name of the method in which the errors occurred
}

// Error print custom errors
func (e *Error) Error() string {
	mes := fmt.Sprintf("%s errId=%v", e.Msg, e.ID)
	if e.Args != "" {
		mes = fmt.Sprintf("%s args=%s", mes, e.Args)
	}

	if e.CauseMsg != "" {
		mes = fmt.Sprintf("%s causemes=%s", mes, e.CauseMsg)
	}

	return mes
}
