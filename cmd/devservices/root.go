package main

import (
	"io"

	"github.com/get-devservices/devservices/pkg/configuration"
	"github.com/spf13/cobra"
)

var globalUsage = `The DevServices tool for managing infrastructure and repositories

Enviroment variables:

| Name | Descroption |
|------|-------------|
|      |             |


By default, the default directories depend on the Operating System. The defaults are listed below:

| Operating System | Cache Path                       | Configuration Path                    | Data Path                      |
|------------------|----------------------------------|---------------------------------------|--------------------------------|
| Linux            | $HOME/.cache/devservices         | $HOME/.config/devservices             | $HOME/.local/share/devservices |
| macOS            | $HOME/Library/Caches/devservices | $HOME/Library/Preferences/devservices | $HOME/Library/devservices      |
| Windows          | %TEMP%\devservices               | %APPDATA%\devservices                 | %APPDATA%\devservices          |
`

func newRootCmd(config *configuration.Configuration, out io.Writer, args []string) (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:          "devservices",
		Short:        "The DevServices tool for managing infrastructure and repositories.",
		Long:         globalUsage,
		SilenceUsage: true,
	}
	flags := cmd.PersistentFlags()
	enviroment.AddFlags(flags)
	flags.Parse(args)

	cmd.AddCommand(
		newVersionCmd(out, config),
	)

	return cmd, nil
}
