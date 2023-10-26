package semester

import (
	"github.com/labstack/echo/v4"
	"github.com/teq-quocbang/course-register/usecase"
)

type Route struct {
	UseCase *usecase.UseCase
}

func Init(group *echo.Group, useCase *usecase.UseCase) {
	r := &Route{UseCase: useCase}

	group.POST("", r.CreateSemester)
	group.GET("", r.GetList)
	group.GET("/:id", r.Get)
}
