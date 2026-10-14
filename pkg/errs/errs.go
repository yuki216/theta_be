package errs

import (
	"errors"
	"fmt"
)

var (
	ErrNotFound               = errors.New("resource not found")
	ErrInvalidInput           = errors.New("invalid input data")
	ErrConflict               = errors.New("resource already exists")
	ErrInvalidStateTransition = errors.New("invalid state transition")
	ErrUnauthorized           = errors.New("unauthorized access")
	ErrForbidden              = errors.New("forbidden operation")
	ErrInternal               = errors.New("internal server error")
)

type CodeError struct {
	Code    string
	Message string
	Err     error
}

func (e *CodeError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("[%s] %s: %v", e.Code, e.Message, e.Err)
	}
	return fmt.Sprintf("[%s] %s", e.Code, e.Message)
}

func (e *CodeError) Unwrap() error {
	return e.Err
}

func NewCodeError(code, message string, underlying error) error {
	return &CodeError{
		Code:    code,
		Message: message,
		Err:     underlying,
	}
}
