package myerror

import "github.com/pkg/errors"

type MyError struct {
	Raw       error
	ErrorCode int
	HTTPCode  int
	Message   string
}

func (e MyError) Error() string {
	if e.Raw != nil {
		return errors.Wrap(e.Raw, e.Message).Error()
	}

	return e.Message
}

// NewError; ErrorCode = {1 digit of service}{2 digits of model}{3 digits of error} .
func NewError(err error, httpCode, errCode int, message string) MyError {
	return MyError{
		Raw:       err,
		ErrorCode: errCode,
		HTTPCode:  httpCode,
		Message:   message,
	}
}
