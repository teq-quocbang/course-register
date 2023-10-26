package payload

import (
	"github.com/go-playground/validator/v10"
	"github.com/teq-quocbang/course-register/util/myerror"
)

type SignUpRequest struct {
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

type LoginRequest struct {
	ID       uint   `json:"id" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func (l *LoginRequest) Validate() error {
	validate := validator.New()
	if err := validate.Struct(l); err != nil {
		return myerror.ErrAccountInvalidParam(err.Error())
	}
	return nil
}

func (s *SignUpRequest) Validate() error {
	validate := validator.New()
	if err := validate.Struct(s); err != nil {
		return myerror.ErrAccountInvalidParam(err.Error())
	}
	return nil
}
