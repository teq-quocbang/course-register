package semester

import (
	"context"

	"github.com/teq-quocbang/course-register/util/myerror"
)

func (u *UseCase) Delete(ctx context.Context, id string) error {
	if id == "" {
		return myerror.ErrSemesterInvalidParam("id")
	}

	err := u.Semester.Delete(ctx, id)
	if err != nil {
		return myerror.ErrSemesterDelete(err)
	}

	return nil
}