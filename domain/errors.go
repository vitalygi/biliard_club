package domain

import "fmt"

type Error struct {
	Code    int
	Message string
	Err     error
}

func (e *Error) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s: %v", e.Message, e.Err)
	}
	return e.Message
}

func NewError(code int, msg string, err error) *Error {
	return &Error{Code: code, Message: msg, Err: err}
}

var (
	ErrInternalServer = NewError(500, "internal server error", nil)
	ErrNotFound       = NewError(404, "item not found", nil)
	ErrConflict       = NewError(409, "already exists", nil)
	ErrBadRequest     = NewError(400, "invalid request", nil)
	ErrUnauthorized   = NewError(401, "unauthorized", nil)
)
