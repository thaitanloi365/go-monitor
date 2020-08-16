package routes

import (
	"github.com/labstack/echo/v4"
	controllers "github.com/thaitanloi365/go-monitor/controller/admin"
)

// SetupAdminRoute setup admin's routes
func SetupAdminRoute(g *echo.Group) {
	g.POST("/login", controllers.Login)

	g.GET("/container/list", controllers.GetListDockerContainer)

	g.GET("/container/:id/stream_logs", controllers.StreamDockerContainerLogs)
	g.GET("/container/:id", controllers.GetDockerContainer)

	g.DELETE("/me/logout", controllers.Logout)
}
