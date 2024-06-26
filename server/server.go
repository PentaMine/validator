package server

import (
	"fmt"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"validator/controller"
)

type Controller interface {
	Init(e *echo.Echo)
}

var controllers = []Controller{
	controller.ExecutionController{},
}

func StartServer(port int, host string) {
	e := echo.New()
	e.Validator = &JSONValidator{validator: validator.New()}
	initControllers(e)

	e.Logger.Fatal(e.Start(fmt.Sprintf("%s:%d", host, port)))
}

func initControllers(e *echo.Echo) {
	for _, c := range controllers {
		c.Init(e)
	}
}
