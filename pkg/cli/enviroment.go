package cli

import (
	"os"
	"strconv"

	"github.com/spf13/pflag"
)

type EnvSettings struct {
	//  Debug indicates whether or not DEVSERVICES is running in Debug mode.
	Debug bool
}

func New() *EnvSettings {
	env := &EnvSettings{
		// debug: os.Getenv("DEVSERVICES_DEBUG")
	}
	env.Debug, _ = strconv.ParseBool(os.Getenv("DEVSERVICES_DEBUG"))

	return env
}

func (env *EnvSettings) AddFlags(fs *pflag.FlagSet) {
	fs.BoolVar(&env.Debug, "debug", env.Debug, "enable verbose output")
}

func (env *EnvSettings) EnvVars() map[string]string {
	return map[string]string{
		"DEVSERVICES_BIN":   os.Args[0],
		"DEVSERVICES_DEBUG": strconv.FormatBool(env.Debug),
	}
}
