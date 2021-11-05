package cmd

import (
	"github.com/midoks/go-p2p-server/internal/app"
	"github.com/urfave/cli"
)

var Service = cli.Command{
	Name:        "service",
	Usage:       "This command starts P2P services",
	Description: `Start Web P2P Server services`,
	Action:      RunService,
	Flags: []cli.Flag{
		stringFlag("config, c", "", "Custom configuration file path"),
	},
}

func RunService(c *cli.Context) error {
	app.Run()
	return nil
}
