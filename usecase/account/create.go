package account

import (
	"context"

	"github.com/teq-quocbang/course-register/model"
	"github.com/teq-quocbang/course-register/payload"
	"github.com/teq-quocbang/course-register/presenter"
	"github.com/teq-quocbang/course-register/util/hashing"
	"github.com/teq-quocbang/course-register/util/myerror"
	"github.com/teq-quocbang/course-register/util/token"
)

func (u *UseCase) SignUp(ctx context.Context, req *payload.SignUpRequest) (*presenter.AccountResponseWrapper, error) {
	// TODO: check permission
	// validate check
	if err := req.Validate(); err != nil {
		return nil, err
	}

	// check unique constraint
	account, err := u.Account.GetAccountByConstraint(ctx, &model.Account{
		Username: req.Username,
		Email:    req.Email,
	})
	if err != nil {
		return nil, myerror.ErrAccountGet(err)
	}

	// check whether constraint is existed
	if account != nil {
		return nil, myerror.ErrAccountConflictUniqueConstraint("Username or password was registered")
	}

	// create account
	hashPassword, err := hashing.ToHashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	createAccountRequest := &model.Account{
		Username:     req.Username,
		HashPassword: hashPassword,
		Email:        req.Email,
	}

	ID, err := u.Account.CreateAccount(ctx, createAccountRequest)
	if err != nil {
		return nil, myerror.ErrAccountCreate(err)
	}

	return &presenter.AccountResponseWrapper{Account: &model.Account{
		ID: ID,
	}}, nil
}

func (p *UseCase) Login(ctx context.Context, req *payload.LoginRequest) (*presenter.AccountLoginResponseWrapper, error) {
	// validate check
	if err := req.Validate(); err != nil {
		return nil, err
	}

	// get account
	account, err := p.Account.GetAccountByID(ctx, req.ID)
	if err != nil {
		return nil, myerror.ErrAccountGet(err)
	}

	// compare password
	if err := hashing.CompareHashPassword(req.Password, account.HashPassword); err != nil {
		return nil, err
	}

	jwt := token.JWT{
		SecretKey: p.Config.TokenSecretKey,
		User: token.UserInfo{
			Username: account.Username,
		},
		TokenLifeTime: p.Config.AccessTokenDuration,
	}
	// generate access token
	accessToken, _, err := jwt.GenerateToken()
	if err != nil {
		return nil, myerror.ErrAccountGenerateToken(err)
	}

	// generate refresh token
	jwt.TokenLifeTime = p.Config.RefreshTokenDuration
	refreshToken, _, err := jwt.GenerateToken()
	if err != nil {
		return nil, myerror.ErrAccountGenerateToken(err)
	}

	return &presenter.AccountLoginResponseWrapper{
		Data: presenter.LoginResponse{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		},
	}, nil
}
