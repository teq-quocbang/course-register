package course

import (
	"context"
	"time"

	"github.com/teq-quocbang/course-register/model"
	"github.com/teq-quocbang/course-register/payload"
	"github.com/teq-quocbang/course-register/presenter"
	"github.com/teq-quocbang/course-register/util/contexts"
	"github.com/teq-quocbang/course-register/util/myerror"
)

func (u *UseCase) Update(ctx context.Context, req *payload.UpdateCourseRequest) (*presenter.CourseResponseWrapper, error) {
	// validate check
	if err := req.Validate(); err != nil {
		return nil, myerror.ErrCourseInvalidParam(err.Error())
	}

	semester, err := u.Semester.GetSemester(ctx, req.SemesterID)
	if err != nil {
		return nil, myerror.ErrSemesterGet(err)
	}

	if semester.EndTime.Before(time.Now()) {
		return nil, myerror.ErrCourseInvalidParam("semester ended")
	}

	// update
	userPrinciple := contexts.GetUserPrincipleByContext(ctx)
	courseModel := &model.Course{
		SemesterID: req.SemesterID,
		UpdatedBy:  &userPrinciple.User.ID,
	}

	err = u.Course.Update(ctx, courseModel)
	if err != nil {
		return nil, myerror.ErrCourseUpdate(err)
	}

	return &presenter.CourseResponseWrapper{
		Course: *courseModel,
	}, nil
}
