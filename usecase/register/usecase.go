package register

import (
	"github.com/teq-quocbang/course-register/config"
	"github.com/teq-quocbang/course-register/repository"
	"github.com/teq-quocbang/course-register/repository/semester"
	mySES "github.com/teq-quocbang/course-register/util/ses"
)

type UseCase struct {
	Semester semester.Repository

	SES mySES.ISES

	Config *config.Config
}

func New(repo *repository.Repository, ses mySES.ISES) IUseCase {
	return &UseCase{
		Semester: repo.Semester,
		SES:      ses,
		Config:   config.GetConfig(),
	}
}
