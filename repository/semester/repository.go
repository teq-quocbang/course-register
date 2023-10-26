package semester

import (
	"context"

	"github.com/teq-quocbang/course-register/model"
)

type Repository interface {
	CreateSemester(context.Context, *model.Semester) error
	ListSemesterByYear(ctx context.Context, year string) ([]model.Semester, error)
	GetSemester(ctx context.Context, semesterID string) (model.Semester, error)
}
