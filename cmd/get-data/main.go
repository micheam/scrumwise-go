package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/micheam/wiseman/scrumwise"
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
	app.Name = "get-data"
	app.Usage = "get-data from scruwise API"
	app.ArgsUsage = "[PROJECT_ID]"
	app.Version = version
	app.Authors = authors
	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:  "props",
			Value: "",
		},
		&cli.BoolFlag{
			Name:  "data-version",
			Value: false,
		},
	}
	app.Action = func(c *cli.Context) error {
		if c.Bool("data-version") {
			var dversion int64
			dversion, err := scrumwise.GetDataVersion(c.Context)
			if err != nil {
				return err
			}
			fmt.Printf("dataVersion: %d", dversion)
			return nil
		}
		if !c.Args().Present() {
			return errors.New("project-id is required")
		}
		projectID := c.Args().First()
		param := scrumwise.NewGetDataParam(projectID)

		if c.String("props") != "" {
			ss := strings.Split(c.String("props"), ",")
			param.AppendProps(ss...)
		}

		data, err := scrumwise.GetData(c.Context, *param)
		if err != nil {
			return err
		}
		b, err := json.Marshal(data) // NOTE: output format is fixed to JSON format, currently.
		if err != nil {
			return err
		}
		fmt.Print(string(b))
		return nil
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
