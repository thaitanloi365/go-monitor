package docker

import (
	"github.com/docker/docker/client"
	"github.com/thaitanloi365/go-monitor/config"
)

// Client client
type Client struct {
	*client.Client
	config *config.Configuration
}

var instance *Client

// New init
func New() *Client {
	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}

	instance = &Client{
		cli,
		config.GetInstance(),
	}

	return instance
}

// GetInstance get instance
func GetInstance() *Client {
	return instance
}
