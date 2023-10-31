package account

import (
	"github.com/stretchr/testify/mock"
	"github.com/teq-quocbang/course-register/model"
	"github.com/teq-quocbang/course-register/payload"
	"github.com/teq-quocbang/course-register/presenter"
	"github.com/teq-quocbang/course-register/util/hashing"
)

func (s *TestSuite) TestSignUp() {
	assertion := s.Assertions
	testUsername := "test_user"
	testPassword := "test_password"
	testEmail := "test@teqnological.asia"
	hashPassword, err := hashing.ToHashPassword(testPassword)
	assertion.NoError(err)

	// good case
	{
		// Arrange

		s.mockAccountRepo.EXPECT().GetAccountByConstraint(s.ctx, &model.Account{
			Username: testUsername,
			Email:    testEmail,
		}).ReturnArguments = mock.Arguments{
			nil, nil,
		}
		s.mockAccountRepo.EXPECT().CreateAccount(s.ctx, &model.Account{
			Username:     testUsername,
			Email:        testEmail,
			HashPassword: hashPassword,
		}).ReturnArguments = mock.Arguments{
			1, nil,
		}

		// Act
		reply, err := s.useCase.SignUp(s.ctx, &payload.SignUpRequest{
			Username: testUsername,
			Email:    testEmail,
			Password: testPassword,
		})

		// Assert
		assertion.NoError(err)
		expected := &presenter.AccountResponseWrapper{
			Account: &model.Account{
				ID: 1,
			},
		}
		assertion.Equal(expected.Account.ID, reply.Account.ID)
	}

	// bad case
	{
		// Arrange

		// Act

		// Assert
	}
}
