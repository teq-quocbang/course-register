package example

import (
	"context"

	"github.com/pkg/errors"
	"gorm.io/gorm"

	"github.com/teq-quocbang/course-register/payload"
	"github.com/teq-quocbang/course-register/presenter"
	"github.com/teq-quocbang/course-register/util/myerror"
)

func (u *UseCase) GetByID(ctx context.Context, req *payload.GetByIDRequest) (*presenter.ExampleResponseWrapper, error) {
	myExample, err := u.ExampleRepo.GetByID(ctx, req.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, myerror.ErrExampleNotFound()
		}

		return nil, myerror.ErrExampleGet(err)
	}

	return &presenter.ExampleResponseWrapper{Example: myExample}, nil
}
