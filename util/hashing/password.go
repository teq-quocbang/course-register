package hashing

import (
	"net/http"

	"git.teqnological.asia/teq-go/teq-pkg/teqerror"
	"github.com/teq-quocbang/course-register/util/myerror"

	"golang.org/x/crypto/bcrypt"
)

func ToHashPassword(password string) ([]byte, error) {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, myerror.ErrAccountCreate(err)
	}
	return hashPassword, nil
}

func CompareHashPassword(password string, hashedPassword []byte) error {
	if err := bcrypt.CompareHashAndPassword(hashedPassword, []byte(password)); err != nil {
		return teqerror.TeqError{
			Raw:       nil,
			ErrorCode: "10010",
			HTTPCode:  http.StatusForbidden,
			Message:   "wrong password",
			IsSentry:  false,
		}
	}
	return nil
}
