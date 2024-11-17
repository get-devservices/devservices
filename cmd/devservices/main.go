package main

import (
	"os"

	"github.com/get-devservices/devservices/pkg/cli"
	"github.com/get-devservices/devservices/pkg/configuration"
)

var enviroment = cli.New()

func main() {
	configuration := new(configuration.Configuration)
	cmd, err := newRootCmd(configuration, os.Stdout, os.Args[1:])
	if err != nil {
		panic(err)
	}

	// cobra.OnInitialize(func() { })

	if err := cmd.Execute(); err != nil {
		panic(err)
	}
}
