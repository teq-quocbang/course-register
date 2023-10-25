package account

import (
	"context"
	"log"

	"github.com/teq-quocbang/course-register/model"
	"github.com/teq-quocbang/course-register/payload"
	"github.com/teq-quocbang/course-register/presenter"
	"github.com/teq-quocbang/course-register/util/hashing"
	"github.com/teq-quocbang/course-register/util/myerror"
)

func (u *UseCase) SignUp(ctx context.Context, req *payload.SignUpRequest) (*presenter.AccountResponseWrapper, error) {
	// TODO: check permission

	// validate check
	if err := req.Validate(); err != nil {
		return nil, err
	}

	// create account
	hashPassword, err := hashing.ToHashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	createAccountRequest := &model.Account{
		Username:     req.Username,
		HashPassword: hashPassword,
		Email:        req.Email,
	}
	log.Println(createAccountRequest)
	ID, err := u.Account.CreateAccount(ctx, createAccountRequest)
	if err != nil {
		return nil, myerror.ErrAccountCreate(err)
	}

	return &presenter.AccountResponseWrapper{Account: &model.Account{
		ID: ID,
	}}, nil
}
