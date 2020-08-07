package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/thaitanloi365/go-logging"
	"github.com/thaitanloi365/go-monitor/config"
	"github.com/thaitanloi365/go-monitor/docker"
	"github.com/thaitanloi365/go-monitor/models"
)

// RegisterCustomContext register custom context
func RegisterCustomContext() echo.MiddlewareFunc {

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			var cc = &models.CustomContext{
				Context:      c,
				DockerClient: docker.GetInstance(),
				Config:       config.GetInstance(),
				Logging:      logging.New(),
			}
			return next(cc)
		}
	}
}
