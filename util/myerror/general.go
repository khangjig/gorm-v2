package myerror

import (
	"net/http"
)

func ErrAuthentication(err error) MyError {
	return MyError{
		Raw:       err,
		HTTPCode:  http.StatusUnauthorized,
		ErrorCode: 10000,
		Message:   "Unauthorized!",
	}
}

func ErrInvalidParam(err error) MyError {
	return MyError{
		Raw:       err,
		HTTPCode:  http.StatusBadRequest,
		ErrorCode: 100001,
		Message:   "Invalid params",
	}
}

func ErrGetByID(err error) MyError {
	return MyError{
		Raw:       err,
		HTTPCode:  http.StatusInternalServerError,
		ErrorCode: 100002,
		Message:   "GetByID error",
	}
}

func ErrInvalidInput(err error) MyError {
	return MyError{
		Raw:       err,
		HTTPCode:  http.StatusNotAcceptable,
		ErrorCode: 100003,
		Message:   "Invalid input",
	}
}

func ErrDeleteByID(err error) MyError {
	return MyError{
		Raw:       err,
		HTTPCode:  http.StatusInternalServerError,
		ErrorCode: 100004,
		Message:   "ErrDeleteByID error",
	}
}
