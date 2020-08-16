package controllers

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"github.com/thaitanloi365/go-monitor/errs"
	"github.com/thaitanloi365/go-monitor/models"
)

// GetListDockerContainer Get list docker container
// @Tags Admin-Docker
// @Summary Get list docker container
// @Description Get list docker container
// @Accept  json
// @Produce  json
// @Header 200 {string} Bearer YOUR_TOKEN
// @Router /api/v1/admin/docker/container/list [get]
func GetListDockerContainer(c echo.Context) error {
	var cc = c.(*models.CustomContext)

	var options = types.ContainerListOptions{}
	containers, err := cc.DockerClient.ContainerList(context.Background(), options)
	if err != nil {
		return err
	}

	return cc.Success(containers)

}

// GetDockerContainer Get list docker container
// @Tags Docker
// @Summary Get list docker container
// @Description Get list docker container
// @Param id path string true "ID of container"
// @Accept  json
// @Produce  json
// @Header 200 {string} Bearer YOUR_TOKEN
// @Router /api/v1/docker/container/{id} [get]
func GetDockerContainer(c echo.Context) error {
	var cc = c.(*models.CustomContext)
	var id = cc.GetPathParamString("id")

	container, err := cc.DockerClient.FindContainer(id)
	if err != nil {
		return err
	}

	return cc.Success(container)

}

// StreamDockerContainerLogs Stream docker container logs
// @Tags Docker
// @Summary Stream docker container logs
// @Description Stream docker container logs
// @Param id path string true "ID of container"
// @Accept  json
// @Produce  json
// @Header 200 {string} Bearer YOUR_TOKEN
// @Router /api/v1/docker/container/{id}/stream_logs [get]
func StreamDockerContainerLogs(c echo.Context) error {
	var cc = c.(*models.CustomContext)
	var containerID = cc.GetPathParamString("id")
	var writer = c.Response().Writer

	var lastEventID = c.Request().Header.Get("Last-Event-ID")
	f, ok := writer.(http.Flusher)
	if !ok {
		return errs.ErrStreamingUnsupported
	}

	container, err := cc.DockerClient.FindContainer(containerID)
	if err != nil {
		return errs.ErrContainerNotFound
	}

	messages, errChanel := cc.DockerClient.StreamContainerLogs(c.Request().Context(), container.ID, cc.Config.DockerTailSize, lastEventID)

	writer.Header().Set("Content-Type", "text/event-stream")
	writer.Header().Set("Cache-Control", "no-cache")
	writer.Header().Set("Connection", "keep-alive")
	writer.Header().Set("X-Accel-Buffering", "no")

Loop:
	for {
		select {
		case message, ok := <-messages:
			if !ok {
				fmt.Fprintf(writer, "event: container-stopped\ndata: end of stream\n\n")
				err = errors.New("event: container-stopped\ndata: end of stream\n\n")
				break Loop
			}
			fmt.Fprintf(writer, "data: %s\n", message)
			if index := strings.IndexAny(message, " "); index != -1 {
				id := message[:index]
				if _, err := time.Parse(time.RFC3339Nano, id); err == nil {
					fmt.Fprintf(writer, "id: %s\n", id)
				}
			}
			fmt.Fprintf(writer, "\n")
			f.Flush()
		case e := <-errChanel:
			if e == io.EOF {
				cc.Logging.Debugf("Container stopped: %v", container.ID)
				fmt.Fprintf(writer, "event: container-stopped\ndata: end of stream\n\n")
				f.Flush()
			} else {
				cc.Logging.Debugf("Error while reading from log stream: %v", e)
				break Loop
			}
		}
	}

	return cc.Success(nil)
}
