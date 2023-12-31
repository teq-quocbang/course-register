package account

import (
	"bou.ke/monkey"
	"github.com/stretchr/testify/mock"

	"github.com/teq-quocbang/course-register/model"
	"github.com/teq-quocbang/course-register/payload"
	"github.com/teq-quocbang/course-register/presenter"
	"github.com/teq-quocbang/course-register/repository/account"
	"github.com/teq-quocbang/course-register/util/hashing"
	"github.com/teq-quocbang/course-register/util/myerror"
)

func (s *TestSuite) TestSignUp() {
	assertion := s.Assertions
	testUsername := "test_user"
	testPassword := "test_password"
	testEmail := "test@teqnological.asia"
	hashPassword, err := hashing.ToHashPassword(testPassword)
	assertion.NoError(err)

	hashing := monkey.Patch(hashing.ToHashPassword, func(string) ([]byte, error) {
		return hashPassword, nil
	})
	defer monkey.Unpatch(hashing)

	// good case
	{
		// Arrange
		mockRepo := account.NewMockRepository(s.T())
		mockRepo.EXPECT().GetAccountByConstraint(s.ctx, &model.Account{
			Username: testUsername,
			Email:    testEmail,
		}).ReturnArguments = mock.Arguments{
			nil, nil,
		}
		mockRepo.EXPECT().CreateAccount(s.ctx, &model.Account{
			Username:     testUsername,
			Email:        testEmail,
			HashPassword: hashPassword,
		}).ReturnArguments = mock.Arguments{
			uint(1), nil,
		}
		u := s.useCase(mockRepo)

		// Act
		reply, err := u.SignUp(s.ctx, &payload.SignUpRequest{
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
	{ // account already existed
		// Arrange
		mockRepo := account.NewMockRepository(s.T())
		mockRepo.EXPECT().GetAccountByConstraint(s.ctx, &model.Account{
			Username: testUsername,
			Email:    testEmail,
		}).ReturnArguments = mock.Arguments{
			&model.Account{
				ID:           1,
				Email:        testEmail,
				Username:     testUsername,
				HashPassword: hashPassword,
			}, nil,
		}
		u := s.useCase(mockRepo)

		// Act
		_, err := u.SignUp(s.ctx, &payload.SignUpRequest{
			Username: testUsername,
			Email:    testEmail,
			Password: testPassword,
		})

		// Assert
		assertion.Error(err)
		expected := myerror.ErrAccountConflictUniqueConstraint("Username or Email was registered")
		assertion.Equal(expected, err)
	}
}

func (s *TestSuite) TestLogin() {
	assertion := s.Assertions
	testUsername := "test_user"
	testPassword := "test_password"
	hashPassword, err := hashing.ToHashPassword(testPassword)
	assertion.NoError(err)

	// good case
	{
		// Arrange
		mockRepo := account.NewMockRepository(s.T())
		mockRepo.EXPECT().GetAccountByID(s.ctx, uint(1)).ReturnArguments = mock.Arguments{
			&model.Account{
				ID:           1,
				Username:     testUsername,
				HashPassword: hashPassword,
				IsVerified:   false,
			}, nil,
		}
		req := &payload.LoginRequest{
			ID:       1,
			Password: testPassword,
		}
		u := s.useCase(mockRepo)

		// Act
		reply, err := u.Login(s.ctx, req)

		// Assert
		assertion.NoError(err)
		assertion.NotNil(reply)
	}

	// bad case
	{ // missing id
		// Arrange
		mockRepo := account.NewMockRepository(s.T())
		req := &payload.LoginRequest{
			Password: testPassword,
		}
		u := s.useCase(mockRepo)

		// Act
		_, err := u.Login(s.ctx, req)

		// Assert
		assertion.Error(err)
		expected := myerror.ErrAccountInvalidParam("Key: 'LoginRequest.ID' Error:Field validation for 'ID' failed on the 'required' tag")
		assertion.Equal(expected, err)
	}
	{ // missing password
		// Arrange
		mockRepo := account.NewMockRepository(s.T())
		req := &payload.LoginRequest{
			ID: 1,
		}
		u := s.useCase(mockRepo)

		// Act
		_, err := u.Login(s.ctx, req)

		// Assert
		assertion.Error(err)
		expected := myerror.ErrAccountInvalidParam("Key: 'LoginRequest.Password' Error:Field validation for 'Password' failed on the 'required' tag")
		assertion.Equal(expected, err)
	}
}
