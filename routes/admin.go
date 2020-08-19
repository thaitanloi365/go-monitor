package routes

import (
	"github.com/labstack/echo/v4"
	controllers "github.com/thaitanloi365/go-monitor/controller/admin"
)

// SetupAdminRoute setup admin's routes
func SetupAdminRoute(g *echo.Group) {
	// Auth
	g.POST("/login", controllers.Login)

	// Me
	g.DELETE("/me/logout", controllers.Logout)

	// Docker
	g.GET("/container/list", controllers.GetListDockerContainer)
	g.GET("/container/:id/stream_logs", controllers.StreamDockerContainerLogs)
	g.GET("/container/:id", controllers.GetDockerContainer)

	// Healthcheck
	g.GET("/job_healthcheck/list", controllers.GetListJobHealthCheck)
	g.DELETE("/job_healthcheck/:tag", controllers.RemoveJobHealthCheckByTag)
	g.GET("/job_healthcheck/:tag", controllers.GetJobHealthCheckByTag)
	g.POST("/job_healthcheck", controllers.AddJobHealthCheck)

	// Notifier
	g.GET("/notifier/:provider", controllers.GetNotifier)
	g.GET("/notifier/list", controllers.GetListNotifier)
	g.PUT("/notifier/:tag", controllers.UpdateNotifier)

}
