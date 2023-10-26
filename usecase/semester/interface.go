package semester

import (
	"context"

	"github.com/teq-quocbang/course-register/payload"
	"github.com/teq-quocbang/course-register/presenter"
)

type IUseCase interface {
	CreateSemester(context.Context, *payload.CreateSemesterRequest) (*presenter.SemesterResponseWrapper, error)
	GetList(context.Context, *payload.GetListSemesterRequest) (*presenter.ListSemesterResponseWrapper, error)
	GetByID(context.Context, string) (*presenter.SemesterResponseWrapper, error)
}
