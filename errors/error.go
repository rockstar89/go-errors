package errors

import (
	"errors"
)

type Error interface {
	Is(target error) bool
	Error() string
	Code() Code
}

type appError struct {
	code    Code
	message string
}

func New(code Code, message string) Error {
	return &appError{
		code:    code,
		message: message,
	}
}

func Is(e error, target error) bool {
	return errors.Is(e, target)
}

func (e *appError) Is(target error) bool {
	return errors.Is(e, target)
}

func (e *appError) Error() string {
	return e.message
}

func (e *appError) Code() Code {
	return e.code
}
