package util

import (
	"github.com/labstack/echo/v4"
)

func BindAndValidate(i interface{}, ctx echo.Context) error {
	err := ctx.Bind(i)
	if err != nil {
		return echo.NewHTTPError(400)
	}

	err = ctx.Validate(i)
	if err != nil {
		return echo.NewHTTPError(400)
	}

	return nil
}
