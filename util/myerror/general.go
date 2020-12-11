package myerror

import (
	"net/http"
)

func ErrInvalidParam(err error) MyError {
	return MyError{
		Raw:       err,
		HTTPCode:  http.StatusBadRequest,
		ErrorCode: 100000,
		Message:   "Invalid params",
	}
}

func ErrGetByID(err error) MyError {
	return MyError{
		Raw:       err,
		HTTPCode:  http.StatusInternalServerError,
		ErrorCode: 100001,
		Message:   "GetByID error",
	}
}

func ErrInvalidInput(err error) MyError {
	return MyError{
		Raw:       err,
		HTTPCode:  http.StatusNotAcceptable,
		ErrorCode: 100002,
		Message:   "Invalid input",
	}
}
