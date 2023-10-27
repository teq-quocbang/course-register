package course

import (
	"context"

	"github.com/teq-quocbang/course-register/presenter"
	"github.com/teq-quocbang/course-register/util/myerror"
)

func (u *UseCase) GetByID(ctx context.Context, id string) (*presenter.CourseResponseWrapper, error) {
	course, err := u.Course.GetCourse(ctx, id)
	if err != nil {
		return nil, myerror.ErrCourseGet(err)
	}

	return &presenter.CourseResponseWrapper{
		Course: course,
	}, nil
}
