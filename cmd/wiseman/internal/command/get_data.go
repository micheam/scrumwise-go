package command

import (
	"encoding/json"
	"fmt"

	"github.com/micheam/wiseman/scrumwise"
	"github.com/urfave/cli/v2"
)

var GetData = &cli.Command{
	Name: "get-data",
	Action: func(c *cli.Context) error {
		projectID := c.String("project")
		param := scrumwise.NewGetDataParam(projectID)
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
