package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/thaitanloi365/go-monitor/controller"
)

// SetupAuthRoute setup root's routes
func SetupAuthRoute(g *echo.Group) {
	g.POST("/login", controller.Login)
}
