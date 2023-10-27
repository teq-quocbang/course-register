package course

import (
	"context"

	"github.com/teq-quocbang/course-register/payload"
	"github.com/teq-quocbang/course-register/presenter"
	"github.com/teq-quocbang/course-register/util/myerror"
)

func (u *UseCase) GetList(ctx context.Context, req *payload.ListCourseBySemesterRequest) (*presenter.ListCourseResponseWrapper, error) {
	if err := req.Validate(); err != nil {
		return nil, myerror.ErrCourseInvalidParam(err.Error())
	}

	course, err := u.Course.ListCourseBySemester(ctx, req.SemesterID)
	if err != nil {
		return nil, myerror.ErrClassGet(err)
	}

	return &presenter.ListCourseResponseWrapper{
		Course: course,
		Meta:   nil,
	}, nil
}
