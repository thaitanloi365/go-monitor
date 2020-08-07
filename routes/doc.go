package routes

import (
	// docs
	_ "github.com/thaitanloi365/go-monitor/docs"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// SetupDocRoute setup document route
func SetupDocRoute(g *echo.Group) {
	g.GET("/*", echoSwagger.WrapHandler)
}
