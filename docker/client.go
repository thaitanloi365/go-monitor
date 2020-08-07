package docker

import "github.com/docker/docker/client"

// Client client
type Client struct {
	*client.Client
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
	}

	return instance
}

// GetInstance get instance
func GetInstance() *Client {
	return instance
}
