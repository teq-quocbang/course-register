package repository

import (
	"context"

	"gorm.io/gorm"

	"github.com/teq-quocbang/course-register/repository/example"
)

type Repository struct {
	Example example.Repository
}

func New(getClient func(ctx context.Context) *gorm.DB) *Repository {
	return &Repository{
		Example: example.NewPG(getClient),
	}
}
