package main

import (
	"log"
	"os"

	"github.com/micheam/wiseman/cmd/wiseman/internal/command"

	"github.com/urfave/cli/v2"
)

var (
	version = "0.1.0"
	authors = []*cli.Author{
		{Name: "Michito Maeda", Email: "michito.maeda@gmail.com"},
	}
)

func main() {
	log.SetOutput(os.Stderr)
	app := cli.NewApp()
	app.Name = "wiseman"
	app.Version = version
	app.Authors = authors
	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:     "project",
			EnvVars:  []string{"SCRUMWISE_PROJECT"},
			Usage:    "scrumwise `PROJECT_ID` to be processed by this app.",
			Required: true,
		},
	}
	app.Commands = []*cli.Command{
		command.GetData,
		command.DataVersion,
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
