package version

import "runtime"

type BuildVersion struct {
	Version   string `json:"version,omitempty"`
	GoVersion string `json:"goVersion,omitempty"`
}

const version = "0.1.0"

func Get() *BuildVersion {
	return &BuildVersion{
		Version:   version,
		GoVersion: runtime.Version(),
	}
}
