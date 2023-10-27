package class

import (
	"context"

	"github.com/teq-quocbang/course-register/model"
)

type Repository interface {
	CreateClass(context.Context, *model.Class) error
	ListClassBySemester(ctx context.Context, semesterID string) ([]model.Class, error)
	GetClass(ctx context.Context, classID string) (model.Class, error)
	Update(context.Context, *model.Class) error
	Delete(context.Context, string) error
	BatchInCreMember(context.Context, string) error
	BatchDeCreMember(context.Context, string) error
}
