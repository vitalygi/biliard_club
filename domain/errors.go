package domain

import "errors"

var (
	// ErrInternalServerError will throw if any the Internal Server Error happen
	ErrInternalServerError = errors.New("internal Server Error")
	// ErrNotFound will throw if the requested item is not exists
	ErrNotFound = errors.New("item is not found")
	// ErrConflict will throw if the current action already exists
	ErrConflict = errors.New("already exist")
	// ErrBadParamInput will throw if the given request-body or params is not valid
	ErrBadParamInput = errors.New("given Param is not valid")
	// ErrWrongCredentials will throw if the given email or password is wrong
	ErrWrongCredentials = errors.New("wrong email or password")
	// ErrUserExists will throw if the user already exists
	ErrUserExists = errors.New("user already exists")
	// ErrUserNotFound will throw if the user is not found
	ErrUserNotFound = errors.New("user not found")
	// ErrTableExists will throw if the table is not found
	ErrTableExists = errors.New("table already exists")
	// ErrTableNotFound will throw if the user is not found
	ErrTableNotFound = errors.New("table not found")
)
