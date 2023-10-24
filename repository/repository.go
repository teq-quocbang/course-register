package repository

import (
	"context"

	"gorm.io/gorm"

	"github.com/teq-quocbang/course-register/repository/account"
	"github.com/teq-quocbang/course-register/repository/example"
)

type Repository struct {
	Account account.Repository
	Example example.Repository
}

func New(getClient func(ctx context.Context) *gorm.DB) *Repository {
	return &Repository{
		Account: account.NewAccountPG(getClient),
		Example: example.NewPG(getClient),
	}
}
