package example

import (
	"github.com/teq-quocbang/course-register/config"
	"github.com/teq-quocbang/course-register/repository"
	"github.com/teq-quocbang/course-register/repository/example"
	mySES "github.com/teq-quocbang/course-register/util/ses"
)

type UseCase struct {
	ExampleRepo example.Repository

	SES mySES.ISES

	Config *config.Config
}

func New(repo *repository.Repository, ses mySES.ISES) IUseCase {
	return &UseCase{
		ExampleRepo: repo.Example,
		SES:         ses,
		Config:      config.GetConfig(),
	}
}
