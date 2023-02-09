package errorx

// TODO(raymondwongso): more common error
const (
	// CodeInternal defines common internal error code
	CodeInternal = "common.internal"
	// CodeNotFound defines common not found error code
	CodeNotFound = "common.not_found"
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
