package register

import (
	"context"
	"errors"

	"github.com/teq-quocbang/course-register/model"
	"github.com/teq-quocbang/course-register/payload"
	"github.com/teq-quocbang/course-register/presenter"
	"github.com/teq-quocbang/course-register/util/contexts"
	"github.com/teq-quocbang/course-register/util/myerror"
	"gorm.io/gorm"
)

func (u *UseCase) Create(ctx context.Context, req *payload.CreateRegisterRequest) (*presenter.RegisterResponseWrapper, error) {
	if err := req.Validate(); err != nil {
		return nil, myerror.ErrSemesterInvalidParam(err.Error())
	}

	// get user principle from context
	userPrinciple := contexts.GetUserPrincipleByContext(ctx)

	// already existed
	register, err := u.Register.Get(ctx, &model.Register{
		AccountID:  userPrinciple.User.ID,
		SemesterID: req.SemesterID,
		ClassID:    req.ClassID,
		CourseID:   req.CourseID,
	})
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, myerror.ErrRegisterGet(err)
	}

	// do swap isCanceled if register already exited and isCanceled == true
	if register.IsCanceled {
		register.IsCanceled = !register.IsCanceled // swap status
		u.Register.BatchUpdateSwapIsCanCeledStatus(ctx, register)
		return &presenter.RegisterResponseWrapper{
			Register: *register,
		}, nil
	}

	// can not register with same course code
	firstCourseChar := string(req.CourseID[0])
	registers, err := u.Register.GetListByFirstCourseChar(ctx, firstCourseChar, userPrinciple.User.ID)
	if err != nil {
		return nil, myerror.ErrRegisterGet(err)
	}
	if len(registers) != 0 {
		return nil, myerror.ErrCanNotRegisterSameCourse("Can not register with same course in one semester")
	}

	register = &model.Register{
		AccountID:  userPrinciple.User.ID,
		SemesterID: req.SemesterID,
		ClassID:    req.ClassID,
		CourseID:   req.CourseID,
		CreatedBy:  userPrinciple.User.ID,
	}
	if err := u.Register.Create(ctx, register); err != nil {
		return nil, myerror.ErrRegisterCreate(err)
	}

	return &presenter.RegisterResponseWrapper{
		Register: *register,
	}, nil
}
