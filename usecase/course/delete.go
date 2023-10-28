package course

import (
	"context"

	"github.com/teq-quocbang/course-register/util/myerror"
)

func (u *UseCase) Delete(ctx context.Context, id string) error {
	if id == "" {
		return myerror.ErrClassInvalidParam("id")
	}

	_, err := u.Course.GetByID(ctx, id)
	if err != nil {
		return myerror.ErrCourseGet(err)
	}

	err = u.Course.Delete(ctx, id)
	if err != nil {
		return myerror.ErrSemesterDelete(err)
	}

	return nil
}
