package semester

import (
	"context"

	"github.com/teq-quocbang/course-register/payload"
	"github.com/teq-quocbang/course-register/presenter"
	"github.com/teq-quocbang/course-register/util/myerror"
)

func (u *UseCase) GetList(ctx context.Context, req *payload.GetListSemesterRequest) (*presenter.ListSemesterResponseWrapper, error) {
	semesters, err := u.Semester.ListSemesterByYear(ctx, req.Year)
	if err != nil {
		return nil, myerror.ErrSemesterGet(err)
	}

	return &presenter.ListSemesterResponseWrapper{
		Semester: semesters,
	}, nil
}
