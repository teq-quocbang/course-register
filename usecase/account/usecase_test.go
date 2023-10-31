package account

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/teq-quocbang/course-register/repository/account"
)

type TestSuite struct {
	suite.Suite

	ctx     context.Context
	useCase *UseCase

	mockAccountRepo *account.MockRepository
}

func (suite *TestSuite) SetupTest() {
	suite.ctx = context.Background()

	suite.mockAccountRepo = &account.MockRepository{}

	suite.useCase = &UseCase{
		Account: suite.mockAccountRepo,
	}
}

func TestUseCaseAuth(t *testing.T) {
	t.Parallel()
	suite.Run(t, &TestSuite{})
}
