package register

import (
	"context"

	"github.com/teq-quocbang/course-register/model"
	"github.com/teq-quocbang/course-register/payload"
	"github.com/teq-quocbang/course-register/presenter"
	"github.com/teq-quocbang/course-register/util/contexts"
	"github.com/teq-quocbang/course-register/util/myerror"
)

func (u *UseCase) UnRegister(ctx context.Context, req *payload.UnRegisterRequest) (*presenter.RegisterResponseWrapper, error) {
	if err := req.Validate(); err != nil {
		return nil, myerror.ErrRegisterInvalidParam(err.Error())
	}

	userPrinciple := contexts.GetUserPrincipleByContext(ctx)

	register, err := u.Register.Get(ctx, &model.Register{
		AccountID:  userPrinciple.User.ID,
		SemesterID: req.SemesterID,
		ClassID:    req.ClassID,
		CourseID:   req.CourseID,
	})
	if err != nil {
		return nil, myerror.ErrRegisterGet(err)
	}

	if register.IsCanceled {
		return &presenter.RegisterResponseWrapper{
			Register: *register,
		}, nil
	}

	register = &model.Register{
		AccountID:  userPrinciple.User.ID,
		SemesterID: req.SemesterID,
		ClassID:    req.ClassID,
		CourseID:   req.CourseID,
		IsCanceled: !register.IsCanceled,
	}

	if err := u.Register.BatchUpdateSwapIsCanCeledStatus(ctx, register); err != nil {
		return nil, myerror.ErrRegisterUpdate(err)
	}

	if err != nil {
		return nil, myerror.ErrRegisterGet(err)
	}

	return &presenter.RegisterResponseWrapper{
		Register: *register,
	}, nil
}