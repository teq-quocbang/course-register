package hashing

import (
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
