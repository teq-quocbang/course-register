package semester

import (
	"context"
	"time"

	"github.com/teq-quocbang/course-register/model"
	"github.com/teq-quocbang/course-register/payload"
	"github.com/teq-quocbang/course-register/presenter"
	"github.com/teq-quocbang/course-register/util/contexts"
	"github.com/teq-quocbang/course-register/util/myerror"
)

func (u *UseCase) Update(ctx context.Context, req *payload.UpdateSemesterRequest) (*presenter.SemesterResponseWrapper, error) {
	// validate check
	if err := req.Validate(); err != nil {
		return nil, myerror.ErrSemesterInvalidParam(err.Error())
	}

	// get semester
	semester, err := u.Semester.GetSemester(ctx, req.ID)
	if err != nil {
		return nil, myerror.ErrSemesterGet(err)
	}

	// check whether the semester is ended
	if semester.EndTime.Before(time.Now()) {
		return nil, myerror.ErrSemesterInvalidParam("the semester was ended")
	}

	start, end, registerStartAt, registerExpiresAt, err := parseStringToTime(req.StartTime, req.EndTime, req.RegisterStartAt, req.RegisterExpiresAt)
	if err != nil {
		return nil, err
	}

	// update
	userPrinciple := contexts.GetUserPrincipleByContext(ctx)
	u.Semester.Update(ctx, &model.Semester{
		ID:                req.ID,
		MinCredits:        req.MinCredits,
		StartTime:         *start,
		EndTime:           *end,
		RegisterStartAt:   *registerStartAt,
		RegisterExpiresAt: *registerExpiresAt,
		UpdatedBy:         &userPrinciple.User.ID,
	})

	return nil, nil
}
