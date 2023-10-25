package myerror

import (
	"fmt"
	"net/http"

	"git.teqnological.asia/teq-go/teq-pkg/teqerror"
)

func ErrAccountGet(err error) teqerror.TeqError {
	return teqerror.TeqError{
		Raw:       err,
		HTTPCode:  http.StatusInternalServerError,
		ErrorCode: "10000",
		Message:   "Failed to get account.",
		IsSentry:  true,
	}
}

func ErrAccountCreate(err error) teqerror.TeqError {
	return teqerror.TeqError{
		Raw:       err,
		HTTPCode:  http.StatusInternalServerError,
		ErrorCode: "10001",
		Message:   "Failed to create account.",
		IsSentry:  true,
	}
}

func ErrAccountUpdate(err error) teqerror.TeqError {
	return teqerror.TeqError{
		Raw:       err,
		HTTPCode:  http.StatusInternalServerError,
		ErrorCode: "10002",
		Message:   "Failed to update account.",
		IsSentry:  true,
	}
}

func ErrAccountDelete(err error) teqerror.TeqError {
	return teqerror.TeqError{
		Raw:       err,
		HTTPCode:  http.StatusInternalServerError,
		ErrorCode: "10003",
		Message:   "Failed to delete account.",
		IsSentry:  true,
	}
}

func ErrAccountNotFound() teqerror.TeqError {
	return teqerror.TeqError{
		Raw:       nil,
		HTTPCode:  http.StatusNotFound,
		ErrorCode: "10004",
		Message:   "Not found.",
		IsSentry:  false,
	}
}

func ErrAccountInvalidParam(param string) teqerror.TeqError {
	return teqerror.TeqError{
		Raw:       nil,
		HTTPCode:  http.StatusBadRequest,
		ErrorCode: "10005",
		Message:   fmt.Sprintf("Invalid parameter: `%s`.", param),
		IsSentry:  false,
	}
}
