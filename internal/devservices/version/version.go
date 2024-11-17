package version

import "runtime"

type BuildVersion struct {
	Version       string `json:"version,omitempty"`
	DockerVersion string `json:"dockerVersion,omitempty"`
	GoVersion     string `json:"goVersion,omitempty"`
}

const version = "0.1.0"

func Get(dockerVersion string) *BuildVersion {
	return &BuildVersion{
		Version:       version,
		DockerVersion: dockerVersion,
		GoVersion:     runtime.Version(),
	}
}
