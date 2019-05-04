package main

import (
	"github.com/urfave/cli"
)

type (
	Commands struct {
		PathFile string
		CliApp   *cli.App
	}
)

func NewCommands() *Commands {
	return &Commands{}
}

func (cmd *Commands) Commands() *Commands {
	cmd.CliApp = cli.NewApp()

	cmd.CliApp.Name = "user center server"
	cmd.CliApp.Usage = "user center"
	cmd.CliApp.Author = "Mike Huang"
	cmd.CliApp.Email = "service@heywoof.com"
	cmd.CliApp.Version = "1.0.0"

	cmd.CliApp.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "config,c",
			Value:       string("config/local/config.yaml"),
			Destination: &cmd.PathFile,
			Usage:       "server config path",
		},
	}

	return cmd
}
