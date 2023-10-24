package account

import (
	"context"

	"github.com/teq-quocbang/course-register/payload"
)

type IUseCase interface {
	CreateAccount(context.Context, *payload.CreateAccountRequest) (*payload.CreateAccountResponse, error)
}
