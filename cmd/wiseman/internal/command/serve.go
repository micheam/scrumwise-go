package command

import (
	"os"
	"path/filepath"

	"github.com/micheam/wiseman"
	"github.com/urfave/cli/v2"
)

var Serve = &cli.Command{
	Name: "serve", Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    "unixsocket",
			EnvVars: []string{"SCRUMWISE_UNIX_SOCKET"},
			Value:   filepath.Join(os.Getenv("HOME"), ".wiseman", "server.sock"),
		},
	},
	Action: func(c *cli.Context) error {
		socket := c.String("unixsocket")
		_ = os.Remove(socket)
		err := wiseman.HttpServer(socket, wiseman.HttpHandler())
		if err != nil {
			return err
		}
		return nil
	},
}
