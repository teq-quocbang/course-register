package semester

import (
	"context"
	"fmt"
	"time"

	"github.com/teq-quocbang/course-register/model"
	"github.com/teq-quocbang/course-register/payload"
	"github.com/teq-quocbang/course-register/presenter"
	"github.com/teq-quocbang/course-register/util/contexts"
	"github.com/teq-quocbang/course-register/util/myerror"
	"github.com/teq-quocbang/course-register/util/times"
)

// read wiki for details https://en.wikipedia.org/wiki/Currying
func currying[t1 any, R1 time.Time, R2 error](f func(t1) (R1, R2)) func(t1) (R1, R2) {
	return func(t t1) (R1, R2) {
		return f(t)
	}
}

func (u *UseCase) CreateSemester(ctx context.Context, req *payload.CreateSemesterRequest) (*presenter.SemesterResponseWrapper, error) {
	if err := req.Validate(); err != nil {
		return nil, myerror.ErrSemesterInvalidParam(err.Error())
	}

	// get user principle from context
	userPrinciple := contexts.GetUserPrincipleByContext(ctx)

	// start currying
	parseTime := currying(times.StringToTime)

	// curried start time
	startTime, err := parseTime(req.StartTime)
	if err != nil {
		return nil, myerror.ErrSemesterInvalidParam(fmt.Sprintf("start time, error: %v", err))
	}
	// validate whether is less than current time
	if lessThan := isLessThanCurrentTime(startTime); lessThan {
		return nil, myerror.ErrSemesterInvalidParam("start time less than current time")
	}

	// curried end time
	endTime, err := parseTime(req.EndTime)
	if err != nil {
		return nil, myerror.ErrSemesterInvalidParam(fmt.Sprintf("end time, error: %v", err))
	}
	// validate whether is less than current time
	if lessThan := isLessThanCurrentTime(endTime); lessThan {
		return nil, myerror.ErrSemesterInvalidParam("end time less than current time")
	}

	// curried register start at
	registerStartAt, err := parseTime(req.RegisterStartAt)
	if err != nil {
		return nil, myerror.ErrSemesterInvalidParam(fmt.Sprintf("register start at, error: %v", err))
	}
	// validate whether is less than current time
	if lessThan := isLessThanCurrentTime(endTime); lessThan {
		return nil, myerror.ErrSemesterInvalidParam("register start at less than current time")
	}

	// curried register expiries at
	registerExpiresAt, err := parseTime(req.RegisterExpiresAt)
	if err != nil {
		return nil, myerror.ErrSemesterInvalidParam(fmt.Sprintf("register expires at, error: %v", err))
	}
	// validate whether is less than current time
	if lessThan := isLessThanCurrentTime(endTime); lessThan {
		return nil, myerror.ErrSemesterInvalidParam("register expires at less than current time")
	}

	semester := &model.Semester{
		ID:                req.ID,
		MinCredits:        req.MinCredits,
		StartTime:         startTime,
		EndTime:           endTime,
		RegisterStartAt:   registerStartAt,
		RegisterExpiresAt: registerExpiresAt,
		CreatedBy:         &userPrinciple.User.ID,
	}
	if err := u.Semester.CreateSemester(ctx, semester); err != nil {
		return nil, myerror.ErrSemesterCreate(err)
	}

	return &presenter.SemesterResponseWrapper{
		Semester: *semester,
	}, nil
}

func isLessThanCurrentTime(t time.Time) bool {
	return t.Before(time.Now())
}
