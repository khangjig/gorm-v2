package test

import (
	"gorm-v2/model"
	"gorm-v2/util"
	"gorm-v2/util/myerror"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

func (r *Route) getAll(c echo.Context) error {
	var (
		ctx     = &util.CustomEchoContext{Context: c}
		req     = model.GetAllRequest{}
		myError = myerror.MyError{}
	)

	err := c.Bind(&req)
	if err != nil {
		return util.Response.Error(ctx, myerror.ErrInvalidInput(err))
	}

	data, err := r.useCase.User.GetAll(ctx, &req)
	if err != nil {
		_ = errors.As(err, &myError)

		return util.Response.Error(c, myError)
	}

	return util.Response.Success(c, data)
}

func (r *Route) getTestByID(c echo.Context) error {
	var (
		ctx     = &util.CustomEchoContext{Context: c}
		idStr   = c.Param("id")
		myError = myerror.MyError{}
	)

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return util.Response.Error(c, myerror.ErrInvalidParam(err))
	}

	user, err := r.useCase.User.GetByID(ctx, id)
	if err != nil {
		_ = errors.As(err, &myError)

		return util.Response.Error(c, myError)
	}

	return util.Response.Success(c, user)
}
