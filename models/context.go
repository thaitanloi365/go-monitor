package models

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/thaitanloi365/go-logging"
	"github.com/thaitanloi365/go-monitor/config"
	"github.com/thaitanloi365/go-monitor/docker"
)

// CustomContext custom echo context
type CustomContext struct {
	echo.Context
	Config       *config.Configuration
	DockerClient *docker.Client
	Logging      *logging.Logging
}

// Success respond success
func (c CustomContext) Success(i interface{}) error {
	var code = http.StatusOK

	if v, ok := i.(string); ok {
		return c.JSON(code, map[string]interface{}{
			"message": v,
		})
	}

	return c.JSON(code, i)
}

// BindAndValidate bind and validate input
func (c CustomContext) BindAndValidate(i interface{}) error {
	var err = c.Bind(i)
	if err != nil {
		return err
	}

	err = c.Validate(i)
	if err != nil {
		return err
	}

	return nil
}
