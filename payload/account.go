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

func (a *SignUpRequest) Validate() error {
	validate := validator.New()
	if err := validate.Struct(a); err != nil {
		return myerror.ErrAccountInvalidParam(err.Error())
	}
	return nil
}
