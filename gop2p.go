package main

import (
	"fmt"
	"os"

	"github.com/midoks/go-p2p-server/internal/cmd"
	"github.com/midoks/go-p2p-server/internal/conf"
	"github.com/urfave/cli"
)

const Version = "0.0.1-dev"
const AppName = "gop2p"

func init() {
	conf.App.Version = Version
	conf.App.Name = AppName
}

func main() {

	app := cli.NewApp()
	app.Name = "gop2p"
	app.Version = Version
	app.Usage = "P2P Service"
	app.Commands = []cli.Command{
		cmd.Service,
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Println("Failed to start application: %v", err)
	}
}
