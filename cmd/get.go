package cmd

import (
	"fmt"

	"github.com/urfave/cli/v2"
	"github.com/vanillaiice/gover-js/load"
)

var getCmd = &cli.Command{
	Name:    "get",
	Aliases: []string{"e"},
	Usage:   "get the current version",
	Action: func(ctx *cli.Context) error {
		versionData, err := load.FromFile(ctx.Path("file"))
		if err != nil {
			return err
		}

		fmt.Printf("%s\n", versionData.Version)

		return nil
	},
}
