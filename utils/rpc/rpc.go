package rpc

import (
	"errors"
	"fmt"

	statuscode "google.golang.org/genproto/googleapis/rpc/code"

	status "github.com/begmaroman/go-micro-boilerplate/proto/status"
)

// ErrNotFoundCode is the integer corresponding to the
// error-not-found status in the Google rpc/code library
var ErrNotFoundCode = int32(statuscode.Code_NOT_FOUND)

// ErrNotFound is returned when an entity can not be found or does not exist
var ErrNotFound = errors.New("the entity does not exist")

// ErrInvalidArgument is returned when the server encounters an error
var ErrInvalidArgument = errors.New("request parameter has an invalid value")

// ErrAlreadyExistsCode is the integer corresponding to the
// error-already-exists status in the Google rpc/code library
var ErrAlreadyExistsCode = int32(statuscode.Code_ALREADY_EXISTS)

// ErrAbortedCode is the integer corresponding to the
// error-aborted status in the Google rpc/code library
var ErrAbortedCode = int32(statuscode.Code_ABORTED)

// ErrAlreadyExists is returned when the entity attempting
// to be created already exists
var ErrAlreadyExists = errors.New("entity already exists")

// ErrInvalidArgumentCode is the integer corresponding to the
// error-invalid-argument status in the Google rpc/code library
var ErrInvalidArgumentCode = int32(statuscode.Code_INVALID_ARGUMENT)

// ErrInternal is returned when the server encounters an error
var ErrInternal = errors.New("server returned an error")

// ErrInternalCode is the integer corresponding to the error-internal
// status in the Google rpc/code library
var ErrInternalCode = int32(statuscode.Code_INTERNAL)

// ErrFailedPreconditionCode is returned when a request violates a
// business-logic-ish precondition as specified in cvspot-api/README.md
var ErrFailedPreconditionCode = int32(statuscode.Code_FAILED_PRECONDITION)

// ErrPermissionDeniedCode is returned when an unauthorized request is made
var ErrPermissionDeniedCode = int32(statuscode.Code_PERMISSION_DENIED)

// ErrUnauthenticatedCode is returned when an unauthorized request is made by an anonymous user
var ErrUnauthenticatedCode = int32(statuscode.Code_UNAUTHENTICATED)

// OKCode is returned when everything worked fine
var OKCode = int32(statuscode.Code_OK)

// OKf returns a Google-style RPC status containing the given successful
// code with the message constructed by formatting the given format
// string with the given varargs
func OKf(format string, args ...interface{}) *status.Status {
	return &status.Status{
		Code:    OKCode,
		Message: fmt.Sprintf(format, args...),
	}
}

// ErrNotFoundf returns a Google-style RPC status containing a
// not-found error with the message constructed by formatting the
// given format string with the given varargs
func ErrNotFoundf(format string, args ...interface{}) *status.Status {
	return Errf(ErrNotFoundCode, format, args...)
}

// ErrInvalidArgumentf returns a Google-style RPC status containing a
// an invalid argument error with the message constructed by formatting the
// given format string with the given varargs
func ErrInvalidArgumentf(format string, args ...interface{}) *status.Status {
	return Errf(ErrInvalidArgumentCode, format, args...)
}

// ErrAlreadyExistsf returns a Google-style RPC status containing a
// an "already exists" error with the message constructed by formatting the
// given format string with the given varargs
func ErrAlreadyExistsf(format string, args ...interface{}) *status.Status {
	return Errf(ErrAlreadyExistsCode, format, args...)
}

// ErrAbortedf returns a Google-style RPC status containing a
// an "aborted" error with the message constructed by formatting the
// given format string with the given varargs
func ErrAbortedf(format string, args ...interface{}) *status.Status {
	return Errf(ErrAbortedCode, format, args...)
}

// ErrPermissionDeniedf returns a Google-style RPC status containing a
// an "permission denied" error with the message constructed by formatting the
// given format string with the given varargs
func ErrPermissionDeniedf(format string, args ...interface{}) *status.Status {
	return Errf(ErrPermissionDeniedCode, format, args...)
}

// ErrFailedPreconditionCodef returns a Google-style RPC status containing a
// an "failed precondition" error with the message constructed by formatting the
// given format string with the given varargs
func ErrFailedPreconditionCodef(format string, args ...interface{}) *status.Status {
	return Errf(ErrFailedPreconditionCode, format, args...)
}

// ErrUnauthenticatedf returns a Google-style RPC status containing a
// an "unauthenticated" error with the message constructed by formatting the
// given format string with the given varargs
func ErrUnauthenticatedf(format string, args ...interface{}) *status.Status {
	return Errf(ErrUnauthenticatedCode, format, args...)
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

// DetailedErr provides the message via Error() and JSON-encoded details via Details()
type DetailedErr interface {
	Error() string
	Details() []byte
}

// ErrDetailed will populate the Details field of Status if the given error implements DetailedErr
func ErrDetailed(code int32, err error) *status.Status {
	s := &status.Status{
		Code:    code,
		Message: err.Error(),
	}

	if derr, ok := err.(DetailedErr); ok {
		s.Details = derr.Details()
	}

	return s
}
