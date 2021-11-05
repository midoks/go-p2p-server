package cmd

import (
	"github.com/urfave/cli"
)

var Service = cli.Command{
	Name:        "service",
	Usage:       "This command starts P2P services",
	Description: `Start Web P2P Server services`,
	Action:      runAllService,
	Flags: []cli.Flag{
		stringFlag("config, c", "", "Custom configuration file path"),
	},
}

func runAllService(c *cli.Context) error {

	// app.Start(conf.Web.HttpPort)
	return nil
}
