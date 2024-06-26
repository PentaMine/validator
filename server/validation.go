package server

import (
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"net/http"
)

type JSONValidator struct {
	validator *validator.Validate
}

func (j *JSONValidator) Validate(i interface{}) error {
	if err := j.validator.Struct(i); err != nil {
		log.Error(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}
