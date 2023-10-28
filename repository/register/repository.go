package register

import (
	"context"

	"github.com/teq-quocbang/course-register/model"
)

type Repository interface {
	CreateRegister(context.Context, *model.Register) error
	ListRegisterBySemester(ctx context.Context, semesterID string) ([]model.Register, error)
	GetRegister(context.Context, *model.Register) (*model.Register, error)
	// swap the state of the is_canceled field
	// false -> true and true -> false
	BatchUpdateSwapIsCanCeledStatus(context.Context, *model.Register) error
}
