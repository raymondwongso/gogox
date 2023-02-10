package errorx

import (
	"fmt"
	"strings"
)

// Error defines standard error of gogox
//
// Code is intended to be a unique identifier of each gogox error. You should use Code to determine your error handling.
// Message is intended to be a user-friendly error message.
// Details is intended to store additional information about the error, commonly used for 4xx or user induced error (such as invalid parameter).
// logMessage is intended to store log error message, typically contains more sensitive information that is needed by end user.
// DO NOT use logMessage for API response
type Error struct {
	Code    string     `json:"code"`
	Message string     `json:"message"`
	Details []*Details `json:"details"`

	logMessage string
	cause      error
	*stack
}

// Details defines error details of gogox's error
type Details struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

// Error implements error interface
func (e *Error) Error() string {
	return e.Message
}

// LogError returns message for debugging, concating cause(s)'s logMessage if any.
func (e *Error) LogError() string {
	if e.cause != nil {
		pe, ok := e.cause.(*Error)
		if !ok {
			return fmt.Sprintf("%s: %s", e.logMessage, e.cause.Error())
		}

		return fmt.Sprintf("%s: %s", e.logMessage, pe.LogError())
	}

	return e.logMessage
}

// New creates new gogox's error
// logMessage is automatically built from code and message
// stack is populated in each initialization
func New(code, message string) *Error {
	return &Error{
		Code:       code,
		Message:    message,
		logMessage: fmt.Sprintf("[%s] %s", code, message),
		stack:      callers(),
	}
}

// Newf creates new gogox's error but with message formatting.
func Newf(code, message string, args ...interface{}) *Error {
	msg := fmt.Sprintf(message, args...)
	return &Error{
		Code:       code,
		Message:    msg,
		logMessage: fmt.Sprintf("[%s] %s", code, msg),
		stack:      callers(),
	}
}

// NewWithLog creates new gogox's error with overriden logMessage
func NewWithLog(code, message, logMessage string) *Error {
	return &Error{
		Code:       code,
		Message:    message,
		logMessage: logMessage,
		stack:      callers(),
	}
}

// NewfWithLog creates new gogox's error with overriden logMessage and message formatting
func NewfWithLog(code, message, logMessage string, args ...interface{}) *Error {
	msg := fmt.Sprintf(message, args...)
	return &Error{
		Code:       code,
		Message:    msg,
		logMessage: logMessage,
		stack:      callers(),
	}
}

// Wrap create new gogox's error with cause
func Wrap(cause error, code, message string) *Error {
	return &Error{
		Code:       code,
		Message:    message,
		logMessage: fmt.Sprintf("[%s] %s", code, message),
		cause:      cause,
		stack:      callers(),
	}
}

// Wrapf create new gogox's error with cause and message formatting
func Wrapf(cause error, code, message string, args ...interface{}) *Error {
	msg := fmt.Sprintf(message, args...)
	return &Error{
		Code:       code,
		Message:    msg,
		logMessage: fmt.Sprintf("[%s] %s", code, msg),
		cause:      cause,
		stack:      callers(),
	}
}

// WrapWithLog create new gogox's error with cause and log message
func WrapWithLog(cause error, code, message, logMessage string) *Error {
	return &Error{
		Code:       code,
		Message:    message,
		logMessage: logMessage,
		cause:      cause,
		stack:      callers(),
	}
}

// WrapfWithLog create new gogox's error with cause, logMessage and message formatting
func WrapfWithLog(cause error, code, message, logMessage string, args ...interface{}) *Error {
	msg := fmt.Sprintf(message, args...)
	return &Error{
		Code:       code,
		Message:    msg,
		logMessage: logMessage,
		cause:      cause,
		stack:      callers(),
	}
}

// PrintStackTrace prints stack trace from error stack.
func (e *Error) PrintStackTrace() string {
	var b strings.Builder

	for _, st := range *e.stack {
		fr := Frame(st)
		fmt.Fprintf(&b, "%s\n\t%s:%d\n", fr.name(), fr.file(), fr.line())
	}

	return b.String()
}

// AddDetails add error details
func (e *Error) AddDetails(details ...*Details) {
	e.Details = append(e.Details, details...)
}

// Parse provides sugar syntax for parsing err to Error instance
func Parse(err error) (*Error, bool) {
	e, ok := err.(*Error)
	return e, ok
}

// ParseAndWrap provides sugar syntax for parsing err to Error instance
// if the parsing fail, automatically wrap the error into internal error with msg
// if the parsing successful, return the original error
func ParseAndWrap(err error, msg string) *Error {
	e, ok := err.(*Error)
	if !ok {
		return Wrap(err, CodeInternal, msg)
	}

	return e
}
