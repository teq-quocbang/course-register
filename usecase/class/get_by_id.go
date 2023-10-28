package class

import (
	"context"

	"github.com/teq-quocbang/course-register/presenter"
	"github.com/teq-quocbang/course-register/util/myerror"
)

func (u *UseCase) GetByID(ctx context.Context, id string) (*presenter.ClassResponseWrapper, error) {
	class, err := u.Class.GetByID(ctx, id)
	if err != nil {
		return nil, myerror.ErrClassGet(err)
	}

	return &presenter.ClassResponseWrapper{
		Class: class,
	}, nil
}
