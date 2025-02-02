package main

import (
	"os"

	"github.com/get-devservices/devservices/internal/devservices/cli"
	"github.com/get-devservices/devservices/pkg/configuration"
)

var enviroment = cli.New()

func main() {
	configuration := configuration.New()
	cmd, err := newRootCmd(configuration, os.Stdout, os.Args[1:])
	if err != nil {
		panic(err)
	}

	// cobra.OnInitialize(func() { })

	if err := cmd.Execute(); err != nil {
		panic(err)
	}
}
