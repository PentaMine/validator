package controller

import (
	"github.com/labstack/echo/v4"
	"validator/dto"
	"validator/util"
)

type EvaluationController struct {
}

func (t EvaluationController) Init(e *echo.Echo) {
	eval := e.Group("/eval")

	{
		eval.POST("/py", t.validatePython)
	}

}

func (t EvaluationController) validatePython(ctx echo.Context) error {
	req := new(dto.ValidateCodeRequest)
	err := util.BindAndValidate(req, ctx)
	if err != nil {
		return err
	}

	//runner.EvaluatePython(req.Code, req.Location)

	return ctx.JSON(200, req)
}
