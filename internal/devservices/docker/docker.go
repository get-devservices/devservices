package docker // import "github.com/get-devservices/devservices/internal/devservices/docker"

import (
	"github.com/docker/docker/client"
)

type DockerClient struct {
	client *client.Client
}

func New() *DockerClient {
	client, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)
	}
	return &DockerClient{
		client: client,
	}
}

func (dockerClient *DockerClient) Version() string {
	return dockerClient.client.ClientVersion()
}
