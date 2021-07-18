package command

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/micheam/wiseman/scrumwise"
	"github.com/urfave/cli/v2"
)

var GetData = &cli.Command{
	Name: "get-data",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    "prop",
			EnvVars: []string{"SCRUMWISE_PROP"},
			Value:   "Project.backlogItems",
		},
	},
	Action: func(c *cli.Context) error {
		projectID := c.String("project")
		param := scrumwise.NewGetDataParam(projectID)
		param.AppendProps(strings.Split(c.String("prop"), ",")...)

		data, err := scrumwise.GetData(c.Context, *param)
		if err != nil {
			return err
		}
		b, err := json.Marshal(data)
		if err != nil {
			return err
		}
		fmt.Print(string(b))
		return nil
	},
}
