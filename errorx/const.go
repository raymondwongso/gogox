package errorx

import (
	"google.golang.org/grpc/codes"
)

const (
	// CodeInternal defines common internal error code
	CodeInternal = "common.internal"
	// CodeNotFound defines common not found error code
	CodeNotFound = "common.not_found"
	// CodeUnauthorized defines common unauthorized error code
	CodeUnauthorized = "common.unauthorized"
	// CodeInvalidParameter defines common invalid parameter error code
	CodeInvalidParameter = "common.invalid_parameter"
	// CodeTimeout defines common timeout error code
	CodeTimeout = "common.timeout"
	// CodeAlreadyExists defines common entity already exists error code
	CodeAlreadyExists = "common.already_exist"
	// CodeForbidden defines common forbidden access error code
	CodeForbidden = "common.forbidden"
	// CodeUnimplemented defines common forbidden access error code
	CodeUnimplemented = "common.unimplemented"
)

var DefaultCodeMap = map[codes.Code]string{
	codes.Canceled:           CodeTimeout,
	codes.Unknown:            CodeInternal,
	codes.InvalidArgument:    CodeInvalidParameter,
	codes.DeadlineExceeded:   CodeTimeout,
	codes.NotFound:           CodeNotFound,
	codes.AlreadyExists:      CodeAlreadyExists,
	codes.PermissionDenied:   CodeForbidden,
	codes.ResourceExhausted:  CodeInternal,
	codes.FailedPrecondition: CodeInternal,
	codes.Aborted:            CodeInternal,
	codes.OutOfRange:         CodeInternal,
	codes.Unimplemented:      CodeUnimplemented,
	codes.Internal:           CodeInternal,
	codes.Unavailable:        CodeInternal,
	codes.DataLoss:           CodeInternal,
	codes.Unauthenticated:    CodeUnauthorized,
}

// ErrInternal returns generic gogox error with CodeInternal
func ErrInternal(msg string) *Error {
	return &Error{
		Code:    CodeInternal,
		Message: msg,
	}
}

// ErrNotFound returns generic gogox error with CodeNotFound
func ErrNotFound(msg string) *Error {
	return &Error{
		Code:    CodeNotFound,
		Message: msg,
	}
}

// ErrUnauthorized returns generic gogox error with CodeUnauthorized
func ErrUnauthorized(msg string) *Error {
	return &Error{
		Code:    CodeUnauthorized,
		Message: msg,
	}
}

// ErrInvalidParameter returns generic gogox error with CodeInvalidParameter
func ErrInvalidParameter(msg string) *Error {
	return &Error{
		Code:    CodeInvalidParameter,
		Message: msg,
	}
}
