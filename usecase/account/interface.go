package account

import (
	"context"

	"github.com/teq-quocbang/course-register/payload"
	"github.com/teq-quocbang/course-register/presenter"
)

type IUseCase interface {
	SignUp(context.Context, *payload.SignUpRequest) (*presenter.AccountResponseWrapper, error)
	Login(context.Context, *payload.LoginRequest) (*presenter.AccountLoginResponseWrapper, error)
}
