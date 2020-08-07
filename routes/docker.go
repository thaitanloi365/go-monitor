package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/thaitanloi365/go-monitor/controller"
)

// SetupDockerRoute setup root's routes
func SetupDockerRoute(g *echo.Group) {
	g.GET("/container/list", controller.GetListDockerContainer)
}
