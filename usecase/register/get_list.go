package register

import (
	"context"

	"github.com/teq-quocbang/course-register/payload"
	"github.com/teq-quocbang/course-register/presenter"
	"github.com/teq-quocbang/course-register/util/contexts"
	"github.com/teq-quocbang/course-register/util/myerror"
)

func (u *UseCase) GetListBySemester(ctx context.Context, req *payload.ListSemesterInformationRequest) (*presenter.ListRegisterResponseWrapper, error) {
	if err := req.Validate(); err != nil {
		return nil, myerror.ErrRegisterInvalidParam(err.Error())
	}

	register, err := u.Register.GetListBySemesterID(ctx, req.SemesterID)
	if err != nil {
		return nil, myerror.ErrSemesterGet(err)
	}

	return &presenter.ListRegisterResponseWrapper{
		Register: register,
	}, nil
}

func (u *UseCase) GetListRegisteredHistories(ctx context.Context, req *payload.ListRegisteredHistories) (*presenter.ListRegisterResponseWrapper, error) {
	if err := req.Validate(); err != nil {
		return nil, myerror.ErrRegisterInvalidParam(err.Error())
	}

	userPrinciple := contexts.GetUserPrincipleByContext(ctx)

	registers, err := u.Register.GetListRegistered(ctx, userPrinciple.User.ID, req.SemesterID)
	if err != nil {
		return nil, myerror.ErrRegisterGet(err)
	}

	return &presenter.ListRegisterResponseWrapper{
		Register: registers,
	}, err
}
