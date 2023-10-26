package account

import (
	"github.com/labstack/echo/v4"
	"github.com/teq-quocbang/course-register/usecase"
)

type Route struct {
	UseCase *usecase.UseCase
}

func Init(group *echo.Group, useCase *usecase.UseCase) {
	r := &Route{UseCase: useCase}

	group.POST("/sign-up", r.SignUp)
	group.POST("/login", r.Login)
}
