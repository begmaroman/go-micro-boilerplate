package rpc

import (
	"fmt"

	statuscode "google.golang.org/genproto/googleapis/rpc/code"

	status "github.com/begmaroman/go-micro-boilerplate/proto/status"
)

// ErrAbortedCode is the integer corresponding to the
// error-aborted status in the Google rpc/code library
var ErrAbortedCode = int32(statuscode.Code_ABORTED)

// ErrAbortedf returns a Google-style RPC status containing a
// an "aborted" error with the message constructed by formatting the
// given format string with the given varargs
func ErrAbortedf(format string, args ...interface{}) *status.Status {
	return Errf(ErrAbortedCode, format, args...)
}

// Errf returns a Google-style RPC status containing the given error
// code with the message constructed by formatting the given format
// string with the given varargs
func Errf(code int32, format string, args ...interface{}) *status.Status {
	return &status.Status{
		Code:    code,
		Message: fmt.Sprintf(format, args...),
	}
}
