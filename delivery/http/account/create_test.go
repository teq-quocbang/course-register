package account

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/teq-quocbang/course-register/fixture/database"
	"github.com/teq-quocbang/course-register/payload"
	"github.com/teq-quocbang/course-register/presenter"
	"github.com/teq-quocbang/course-register/repository"
	"github.com/teq-quocbang/course-register/usecase"
	"github.com/teq-quocbang/course-register/util/test"
)

func TestSignUp(t *testing.T) {
	assertion := assert.New(t)
	db := database.InitDatabase()
	defer db.TruncateTables()

	repo := repository.New(db.GetClient)
	r := Route{
		UseCase: usecase.New(repo, nil),
	}

	// good case
	{
		// Arrange
		db.TruncateTables() // clean database before test
		req := &payload.SignUpRequest{
			Username: "test_user_name",
			Password: "test_password",
			Email:    "test@teqnological.asia",
		}
		resp, ctx := setUpTestSignUp(req)
		defer resp.Result().Body.Close()

		// Act
		err := r.SignUp(ctx)

		// Assert
		assertion.NoError(err)
		assertion.Equal(200, resp.Code)
		actual, err := test.UnmarshalBody[*presenter.AccountResponseWrapper](resp.Body.Bytes())
		assertion.NoError(err)
		assertion.Equal(uint(1), actual.Account.ID)
	}

	// bad case
	{ // missing params
		// Arrange
		db.TruncateTables() // clean database before test
		resp, ctx := setUpTestSignUp(nil)
		defer resp.Result().Body.Close()

		// Act
		err := r.SignUp(ctx)

		// Assert
		assertion.NoError(err)
		assertion.Equal(http.StatusBadRequest, resp.Code)
	}
}

func TestLogin(t *testing.T) {
	assertion := assert.New(t)
	db := database.InitDatabase()
	defer db.TruncateTables()

	repo := repository.New(db.GetClient)
	r := Route{
		UseCase: usecase.New(repo, nil),
	}

	// good case
	{
		// Arrange
		db.TruncateTables() // clean database before test
		signUpRequest := payload.SignUpRequest{
			Username: "test_username",
			Password: "test_password",
			Email:    "test@teqnological.asia",
		}
		_, ctx := setUpTestSignUp(&signUpRequest)
		err := r.SignUp(ctx)
		assertion.NoError(err)

		req := &payload.LoginRequest{
			ID:       1,
			Password: "test_password",
		}
		resp, ctx := setUpTestLogin(req)

		// Act
		err = r.Login(ctx)

		// Assert
		assertion.NoError(err)
		assertion.Equal(200, resp.Code)
	}

	// bad case
	{ // 400
		// Arrange
		db.TruncateTables() // clean database before test
		resp, ctx := setUpTestLogin(nil)

		// Act
		err := r.Login(ctx)

		// Assert
		assertion.NoError(err)
		assertion.Equal(400, resp.Code)
	}
}
