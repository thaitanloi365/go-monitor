package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/thaitanloi365/go-monitor/validation"
)

// Setup init all middlewares
func Setup(e *echo.Echo) {
	e.Use(RegisterCustomContext())
	e.Use(middleware.RequestID())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	e.Use(middleware.BodyLimit("20M"))

	e.Validator = validation.RegisterValidation()

	e.HTTPErrorHandler = CustomErrorHandler
}
