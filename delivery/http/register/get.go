package register

// Get semester by id
// @Summary Get a semester
// @Description Get semester by id
// @Tags Example
// @Accept json
// @Produce json
// @Security AuthToken
// @Param id path int true "id"
// @Success 200 {object} presenter.RegisterResponseWrapper
// @Router /register/histories [get] .
// func (r *Route) GetHistories(c echo.Context) error {
// 	var (
// 		ctx  = &teq.CustomEchoContext{Context: c}
// 		req  = c.Param("id")
// 		resp *presenter.RegisterResponseWrapper
// 	)

// 	resp, err := r.UseCase.Register.GetListRegisteredHistories()
// 	if err != nil {
// 		return teq.Response.Error(c, err.(teqerror.TeqError))
// 	}

// 	return teq.Response.Success(c, resp)
// }
