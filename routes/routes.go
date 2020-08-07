package routes

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
	"github.com/thaitanloi365/go-monitor/config"
	"github.com/thaitanloi365/go-monitor/docs"
	"github.com/thaitanloi365/go-monitor/middleware"
)

// SetupRoutes Setup all routes for project
func SetupRoutes() {
	var e = echo.New()

	middleware.Setup(e)

	var apiGroup = e.Group("/api/v1")
	apiGroup.Use(echoMiddleware.Gzip())

	var docGroup = e.Group("/api/docs")
	docGroup.Use(middleware.IsBasicAuth())

	SetupAuthRoute(apiGroup.Group(""))

	SetupDockerRoute(apiGroup.Group("/docker"))

	SetupDocRoute(docGroup)

	gracefulShutdown(e)
}

func gracefulShutdown(e *echo.Echo) {
	go func() {
		var port = fmt.Sprintf(":%s", config.GetInstance().ServerPort)
		if err := e.Start(port); err != nil {
			e.Logger.Info("shutting down the server")
		}

		fmt.Printf("Server started at: %s\n", config.GetInstance().ServerBaseURL)
	}()

	go func() {
		setupDocConfig()

	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 10 seconds.
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}

func setupDocConfig() {
	var segments = strings.Split(config.GetInstance().ServerBaseURL, "://")
	docs.SwaggerInfo.Title = fmt.Sprintf("%s api", config.GetInstance().AppName)
	docs.SwaggerInfo.Version = config.GetInstance().Version
	docs.SwaggerInfo.Host = segments[1]
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	if strings.HasPrefix(segments[0], "https") {
		docs.SwaggerInfo.Schemes = []string{"https", "http"}
	}
}
