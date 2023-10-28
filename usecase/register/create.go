package register

import (
	"context"

	"github.com/teq-quocbang/course-register/model"
	"github.com/teq-quocbang/course-register/payload"
	"github.com/teq-quocbang/course-register/presenter"
	"github.com/teq-quocbang/course-register/util/contexts"
	"github.com/teq-quocbang/course-register/util/myerror"
)

func (u *UseCase) CreateSemester(ctx context.Context, req *payload.CreateSemesterRequest) (*presenter.SemesterResponseWrapper, error) {
	if err := req.Validate(); err != nil {
		return nil, myerror.ErrSemesterInvalidParam(err.Error())
	}

	// get user principle from context
	userPrinciple := contexts.GetUserPrincipleByContext(ctx)

	startTime, endTime, registerStartAt, registerExpiresAt, err := parseStringToTime(req.StartTime, req.EndTime, req.RegisterStartAt, req.RegisterExpiresAt)
	if err != nil {
		return nil, err
	}

	// validate time
	if err := timeValidate(*startTime, *endTime, *registerStartAt, *registerExpiresAt); err != nil {
		return nil, err
	}

	semester := &model.Semester{
		ID:                req.ID,
		MinCredits:        req.MinCredits,
		StartTime:         *startTime,
		EndTime:           *endTime,
		RegisterStartAt:   *registerStartAt,
		RegisterExpiresAt: *registerExpiresAt,
		CreatedBy:         &userPrinciple.User.ID,
	}
	if err := u.Semester.CreateSemester(ctx, semester); err != nil {
		return nil, myerror.ErrSemesterCreate(err)
	}

	return &presenter.SemesterResponseWrapper{
		Semester: *semester,
	}, nil
}
