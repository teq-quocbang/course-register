package course

import (
	"github.com/teq-quocbang/course-register/config"
	"github.com/teq-quocbang/course-register/repository"
	"github.com/teq-quocbang/course-register/repository/class"
	"github.com/teq-quocbang/course-register/repository/course"
	"github.com/teq-quocbang/course-register/repository/semester"
	mySES "github.com/teq-quocbang/course-register/util/ses"
)

type UseCase struct {
	Semester semester.Repository
	Class    class.Repository
	Course   course.Repository

	SES mySES.ISES

	Config *config.Config
}

func New(repo *repository.Repository, ses mySES.ISES) IUseCase {
	return &UseCase{
		Semester: repo.Semester,
		Class:    repo.Class,
		Course:   repo.Course,
		SES:      ses,
		Config:   config.GetConfig(),
	}
}
