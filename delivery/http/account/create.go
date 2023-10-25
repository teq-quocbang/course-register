package account

import (
	"git.teqnological.asia/teq-go/teq-pkg/teq"
	"git.teqnological.asia/teq-go/teq-pkg/teqerror"
	"github.com/labstack/echo/v4"
	"github.com/teq-quocbang/course-register/payload"
	"github.com/teq-quocbang/course-register/presenter"
)

// Sign up
// @Summary Sign up
// @Description create an account
// @Tags Account
// @Accept  json
// @Produce json
// @Security no
// @Param req body payload.SignUpRequest true "Account info"
// @Success 200 {object} presenter.AccountResponseWrapper
// @Router /user/sign-up [post] .
func (r *Route) SignUp(c echo.Context) error {
	var (
		ctx  = &teq.CustomEchoContext{Context: c}
		req  = payload.SignUpRequest{}
		resp *presenter.AccountResponseWrapper
	)

	if err := c.Bind(&req); err != nil {
		return teq.Response.Error(ctx, teqerror.ErrInvalidParams(err))
	}

	resp, err := r.UseCase.Account.SignUp(ctx, &req)
	if err != nil {
		return teq.Response.Error(c, err.(teqerror.TeqError))
	}

	return teq.Response.Success(c, resp)
}
