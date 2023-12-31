package class

import (
	"git.teqnological.asia/teq-go/teq-pkg/teq"
	"git.teqnological.asia/teq-go/teq-pkg/teqerror"
	"github.com/labstack/echo/v4"
	"github.com/teq-quocbang/course-register/payload"
	"github.com/teq-quocbang/course-register/presenter"
)

// Create class
// @Summary Create class
// @Description create a class
// @Tags Class
// @Accept  json
// @Produce json
// @Security no
// @Param req body payload.CreateClassRequest
// @Success 200 {object} presenter.ClassResponseWrapper
// @Router /class [post] .
func (r *Route) CreateClass(c echo.Context) error {
	var (
		ctx  = &teq.CustomEchoContext{Context: c}
		req  = payload.CreateClassRequest{}
		resp *presenter.ClassResponseWrapper
	)

	if err := c.Bind(&req); err != nil {
		return teq.Response.Error(ctx, teqerror.ErrInvalidParams(err))
	}

	resp, err := r.UseCase.Class.CreateClass(ctx, &req)
	if err != nil {
		return teq.Response.Error(c, err.(teqerror.TeqError))
	}

	return teq.Response.Success(c, resp)
}
