package errorx

const (
	// CodeInternal defines common internal error code
	CodeInternal = "common.internal"
	// CodeNotFound defines common not found error code
	CodeNotFound = "common.not_found"
	// CodeUnauthorized defines common unauthorized error code
	CodeUnauthorized = "common.unauthorized"
	// CodeInvalidParameter defines common invalid parameter error code
	CodeInvalidParameter = "commmon.invalid_parameter"
)

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
