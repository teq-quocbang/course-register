package account

import (
	"context"

	"github.com/teq-quocbang/course-register/model"
)

type Repository interface {
	CreateAccount(context.Context, *model.Account) (ID uint, err error)
	GetAccountByID(ctx context.Context, studentID uint) (*model.Account, error)
	GetByAccountConstraint(context.Context, *model.Account) (*model.Account, error)
	CreateVerifyAccount(context.Context, *model.AccountVerify) error
	GetVerifyAccountByID(ctx context.Context, studentID uint) (*model.AccountVerify, error)
}
