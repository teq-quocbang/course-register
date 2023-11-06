package register

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/teq-quocbang/course-register/cache"
	"github.com/teq-quocbang/course-register/config"
	"github.com/teq-quocbang/course-register/repository/class"
	"github.com/teq-quocbang/course-register/repository/course"
	"github.com/teq-quocbang/course-register/repository/register"
	"github.com/teq-quocbang/course-register/repository/semester"
)

type TestSuite struct {
	suite.Suite

	ctx     context.Context
	useCase func(
		*register.MockRepository,
		*semester.MockRepository,
		*class.MockRepository,
		*course.MockRepository,
		*cache.MockICache) UseCase
}

func (suite *TestSuite) SetupTest() {
	suite.ctx = context.Background()

	suite.useCase = NewTestUseCase
}

func TestUseCaseAuth(t *testing.T) {
	t.Parallel()
	suite.Run(t, &TestSuite{})
}

func NewTestUseCase(
	register *register.MockRepository,
	semester *semester.MockRepository,
	class *class.MockRepository,
	course *course.MockRepository,
	cache *cache.MockICache) UseCase {
	return UseCase{
		Register: register,
		Class:    class,
		Course:   course,
		Semester: semester,
		Cache:    cache,
		Config:   config.GetConfig(),
	}
}
