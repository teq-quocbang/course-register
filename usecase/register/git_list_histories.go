package register

import (
	"context"
	"fmt"

	"github.com/teq-quocbang/course-register/payload"
	"github.com/teq-quocbang/course-register/presenter"
	"github.com/teq-quocbang/course-register/util/contexts"
	"github.com/teq-quocbang/course-register/util/myerror"
)

func (u *UseCase) GetListRegisteredHistories(ctx context.Context, req *payload.ListRegisteredHistories) (*presenter.ListRegisterResponseWrapper, error) {
	userPrinciple := contexts.GetUserPrincipleByContext(ctx)
	req.Format()

	var (
		order = make([]string, 0)
	)

	if req.OrderBy != "" {
		order = append(order, fmt.Sprintf("%s %s", req.OrderBy, req.SortBy))
	}

	registers, total, err := u.Register.GetListRegistered(ctx, userPrinciple.User.ID, req.SemesterID, order, req.Paginator)
	if err != nil {
		return nil, myerror.ErrRegisterGet(err)
	}

	response := &presenter.ListRegisterResponseWrapper{
		Register: make([]presenter.RegisterResponseCustom, len(registers)),
		Meta: map[string]interface{}{
			"page":  req.Paginator.Page,
			"limit": req.Paginator.Limit,
			"total": total,
		},
	}
	for i, r := range registers {
		semester, err := u.Semester.GetByID(ctx, r.SemesterID)
		if err != nil {
			return nil, myerror.ErrSemesterGet(err)
		}

		class, err := u.Class.GetByID(ctx, r.ClassID)
		if err != nil {
			return nil, myerror.ErrClassGet(err)
		}

		course, err := u.Course.GetByID(ctx, r.CourseID)
		if err != nil {
			return nil, myerror.ErrCourseGet(err)
		}

		response.Register[i] = presenter.RegisterResponseCustom{
			AccountID: r.AccountID,
			Semester:  semester,
			Class:     class,
			Course:    course,
		}
	}

	return response, err
}
