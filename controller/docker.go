package controller

import (
	"context"

	"github.com/docker/docker/api/types"
	"github.com/labstack/echo/v4"
	"github.com/thaitanloi365/go-monitor/models"
)

// GetListDockerContainer Get list docker container
// @Tags Docker
// @Summary Get list docker container
// @Description Get list docker container
// @Accept  json
// @Produce  json
// @Header 200 {string} Bearer YOUR_TOKEN
// @Router /api/v1/docker/container/list [get]
func GetListDockerContainer(c echo.Context) error {
	var cc = c.(*models.CustomContext)

	var options = types.ContainerListOptions{}
	containers, err := cc.DockerClient.ContainerList(context.Background(), options)
	if err != nil {
		return err
	}

	return cc.Success(containers)

}
