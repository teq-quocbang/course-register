package grpc

import (
	"github.com/teq-quocbang/course-register/config"
	"github.com/teq-quocbang/course-register/repository"
	"github.com/teq-quocbang/course-register/repository/example"
)

type UseCase struct {
	ExampleRepo example.Repository

	Config *config.Config
}

func New(repo *repository.Repository) IUseCase {
	return &UseCase{
		ExampleRepo: repo.Example,
		Config:      config.GetConfig(),
	}
}
