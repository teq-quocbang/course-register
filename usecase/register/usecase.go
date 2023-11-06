package register

import (
	"github.com/teq-quocbang/course-register/cache"
	"github.com/teq-quocbang/course-register/config"
	"github.com/teq-quocbang/course-register/repository"
	"github.com/teq-quocbang/course-register/repository/account"
	"github.com/teq-quocbang/course-register/repository/class"
	"github.com/teq-quocbang/course-register/repository/course"
	"github.com/teq-quocbang/course-register/repository/register"
	"github.com/teq-quocbang/course-register/repository/semester"
	mySES "github.com/teq-quocbang/course-register/util/ses"
)

type UseCase struct {
	Account  account.Repository
	Semester semester.Repository
	Class    class.Repository
	Course   course.Repository
	Register register.Repository

	SES mySES.ISES

	Cache cache.ICache

	Config *config.Config
}

func New(repo *repository.Repository, ses mySES.ISES, cache cache.ICache) IUseCase {
	return &UseCase{
		Account:  repo.Account,
		Semester: repo.Semester,
		Class:    repo.Class,
		Course:   repo.Course,
		Register: repo.Register,
		Cache:    cache,
		SES:      ses,
		Config:   config.GetConfig(),
	}
}
