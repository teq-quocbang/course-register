package course

import (
	"context"

	"github.com/teq-quocbang/course-register/model"
)

type Repository interface {
	CreateCourse(context.Context, *model.Course) error
	ListCourseBySemester(ctx context.Context, semesterID string) ([]model.Course, error)
	GetCourse(ctx context.Context, courseID string) (model.Course, error)
	Update(context.Context, *model.Course) error
	Delete(context.Context, string) error
}
