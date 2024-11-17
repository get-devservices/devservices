package main

import (
	"fmt"
	"io"

	"github.com/get-devservices/devservices/internal/devservices/version"
	"github.com/get-devservices/devservices/pkg/require"
	"github.com/spf13/cobra"
)

const versionDescription = `
Show the version for DevServices

This will print a representations the version of DevServices.
The output will look something like this:

version.BuildVersion{Version:"v0.1.0", GoVersion:"go1.23.2"}

- Version is the semantic version of the release.
- GoVersion is the version of Go that was used to compile DevService.
`

type versionOptions struct {
	short bool
}

func newVersionCmd(out io.Writer) *cobra.Command {
	options := &versionOptions{}

	cmd := &cobra.Command{
		Use:               "version",
		Short:             "print the client version information",
		Long:              versionDescription,
		Args:              require.NoArgs,
		ValidArgsFunction: require.NoMoreArgsCompFunc,
		RunE: func(_ *cobra.Command, _ []string) error {
			return options.run(out)
		},
	}

	return cmd
}

func (options *versionOptions) run(out io.Writer) error {
	fmt.Fprintln(out, buildVersion(options.short))
	return nil
}

func buildVersion(short bool) string {
	v := version.Get()
	if short {
		return v.Version
	}
	return fmt.Sprintf("%#v", v)
}
