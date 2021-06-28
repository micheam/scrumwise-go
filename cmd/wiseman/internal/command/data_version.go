package command

import (
	"fmt"

	"github.com/micheam/wiseman/scrumwise"
	"github.com/urfave/cli/v2"
)

var DataVersion = &cli.Command{
	Name: "data-version",
	Action: func(c *cli.Context) error {
		var v int64
		v, err := scrumwise.GetDataVersion(c.Context)
		if err != nil {
			return err
		}
		fmt.Printf("dataVersion: %d", v)
		return nil
	},
}
