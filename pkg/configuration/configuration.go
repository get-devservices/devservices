package configuration // import "github.com/get-devservices/devservices/pkg/configuration"

import "github.com/get-devservices/devservices/internal/devservices/docker"

type Configuration struct {
	DockerClient *docker.DockerClient
}

// Create object configuration
func New() *Configuration {
	return &Configuration{
		DockerClient: docker.New(),
	}
}
