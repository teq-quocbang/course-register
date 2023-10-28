package repository

import (
	"context"

	"gorm.io/gorm"

	"github.com/teq-quocbang/course-register/repository/account"
	"github.com/teq-quocbang/course-register/repository/class"
	"github.com/teq-quocbang/course-register/repository/course"
	"github.com/teq-quocbang/course-register/repository/example"
	"github.com/teq-quocbang/course-register/repository/register"
	"github.com/teq-quocbang/course-register/repository/semester"
)

type Repository struct {
	Account  account.Repository
	Semester semester.Repository
	Class    class.Repository
	Course   course.Repository
	Register register.Repository
	Example  example.Repository
}

func New(getClient func(ctx context.Context) *gorm.DB) *Repository {
	return &Repository{
		Account:  account.NewAccountPG(getClient),
		Semester: semester.NewSemesterPG(getClient),
		Course:   course.NewCoursePG(getClient),
		Class:    class.NewClassPG(getClient),
		Register: register.NewClassPG(getClient),
		Example:  example.NewPG(getClient),
	}
}
