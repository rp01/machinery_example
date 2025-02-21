package main

import (
	"mchn/sender"
	"mchn/worker"
	"os"

	"github.com/urfave/cli"
)

var (
	app *cli.App
)

func init() {
	// Initialise a CLI app
	app = cli.NewApp()
	app.Name = "machinery"
	app.Usage = "machinery worker and send example tasks with machinery send"
	app.Version = "0.0.0"
}

func main() {
	// Set the CLI app commands
	app.Commands = []cli.Command{
		{
			Name:  "worker",
			Usage: "launch machinery worker",
			Action: func(c *cli.Context) error {
				if err := worker.Worker(); err != nil {
					return cli.NewExitError(err.Error(), 1)
				}
				return nil
			},
		},
		{
			Name:  "send",
			Usage: "send example tasks ",
			Action: func(c *cli.Context) error {
				if err := sender.Send(); err != nil {
					return cli.NewExitError(err.Error(), 1)
				}
				return nil
			},
		},
	}

	// Run the CLI app
	_ = app.Run(os.Args)
}
