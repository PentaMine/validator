package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"validator/dto"
	runner "validator/evaluator"
	"validator/util"
)

type ExecutionController struct {
}

func (t ExecutionController) Init(e *echo.Echo) {
	eval := e.Group("/exec")

	{
		eval.POST("/py", t.executePython)
		eval.POST("/js", t.executeJavascript)
	}

}

func (t ExecutionController) executePython(ctx echo.Context) error {
	req := new(dto.ExecuteCodeRequest)
	err := util.BindAndValidate(req, ctx)
	if err != nil {
		return err
	}

	exec := runner.NewPythonEvaluator()

	stdOut, err := exec.Execute(req.Code)

	if err != nil {
		log.Error(err)
		//return err
	}

	return ctx.JSON(200, dto.ExecuteCodeResponse{StdOut: stdOut})
}

func (t ExecutionController) executeJavascript(ctx echo.Context) error {
	req := new(dto.ExecuteCodeRequest)
	err := util.BindAndValidate(req, ctx)
	if err != nil {
		return err
	}

	exec := runner.NewJavaScriptEvaluator()

	stdOut, err := exec.Execute(req.Code)

	if err != nil {
		log.Error(err)
		//return err
	}

	return ctx.JSON(200, dto.ExecuteCodeResponse{StdOut: stdOut})
}
