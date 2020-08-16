package docker

import (
	"context"
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
)

// FindContainer find container
func (client *Client) FindContainer(id string) (Container, error) {
	var container Container
	containers, err := client.ListContainers()
	if err != nil {
		return container, err
	}

	found := false
	for _, c := range containers {
		if c.ID == id {
			container = c
			found = true
			break
		}
	}
	if found == false {
		return container, fmt.Errorf("Unable to find container with id: %s", id)
	}

	return container, nil
}

// ListContainers list container
func (client *Client) ListContainers() ([]Container, error) {
	containerListOptions := types.ContainerListOptions{
		Filters: filters.Args{},
		All:     true,
	}
	list, err := client.ContainerList(context.Background(), containerListOptions)
	if err != nil {
		return nil, err
	}

	var containers = make([]Container, 0, len(list))
	for _, c := range list {
		container := Container{
			ID:      c.ID,
			Names:   c.Names,
			Name:    strings.TrimPrefix(c.Names[0], "/"),
			Image:   c.Image,
			ImageID: c.ImageID,
			Command: c.Command,
			Created: c.Created,
			State:   c.State,
			Status:  c.Status,
		}
		containers = append(containers, container)
	}

	sort.Slice(containers, func(i, j int) bool {
		return strings.ToLower(containers[i].Name) < strings.ToLower(containers[j].Name)
	})

	return containers, nil
}

// StreamContainerLogs stream container logs
func (client *Client) StreamContainerLogs(ctx context.Context, id string, tailSize int, since string) (<-chan string, <-chan error) {
	var options = types.ContainerLogsOptions{
		ShowStdout: true,
		ShowStderr: true,
		Follow:     true,
		Tail:       strconv.Itoa(tailSize),
		Timestamps: true,
		Since:      since,
	}
	reader, err := client.ContainerLogs(ctx, id, options)
	errChannel := make(chan error, 1)

	if err != nil {
		errChannel <- err
		close(errChannel)
		return nil, errChannel
	}

	messages := make(chan string)
	go func() {
		<-ctx.Done()
		reader.Close()
	}()

	containerJSON, _ := client.ContainerInspect(ctx, id)

	go func() {
		defer close(messages)
		defer close(errChannel)
		defer reader.Close()
		nextEntry := logReader(reader, containerJSON.Config.Tty)
		for {
			line, err := nextEntry()
			if err != nil {
				errChannel <- err
				break
			}
			select {
			case messages <- line:
			case <-ctx.Done():
			}
		}
	}()

	return messages, errChannel
}
