package errs

import (
	"net/http"
)

// Error error struct
type Error struct {
	Code    int
	Message string
	Header  int
}

// Error implement error method
func (e *Error) Error() string {
	return e.Message
}

// WithMessage override msg
func (e *Error) WithMessage(msg string) *Error {
	e.Message = msg
	return e
}

// New create a custom error
func New(code int, message string, header ...int) *Error {
	var hc = http.StatusBadRequest
	if len(header) > 0 && header[0] < 1000 {
		hc = header[0]
	}
	return &Error{Code: code, Message: message, Header: hc}
}
