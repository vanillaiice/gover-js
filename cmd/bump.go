package cmd

import (
	"fmt"
	"log"

	"github.com/Masterminds/semver/v3"
	"github.com/urfave/cli/v2"
	"github.com/vanillaiice/gover-js/gen"
	"github.com/vanillaiice/gover-js/load"
)

// perm is the file permission used to create files.
const perm = 0644

// bumpCmd is the bump command.
var bumpCmd = &cli.Command{
	Name:    "bump",
	Usage:   "bump version",
	Aliases: []string{"b"},
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:    "major",
			Aliases: []string{"m"},
			Usage:   "bump major version",
		},
		&cli.BoolFlag{
			Name:    "minor",
			Aliases: []string{"n"},
			Usage:   "bump minor version",
		},
		&cli.BoolFlag{
			Name:    "patch",
			Aliases: []string{"p"},
			Usage:   "bump patch version",
		},
	},
	Action: func(ctx *cli.Context) (err error) {
		versionData, err := load.FromFile(ctx.Path("file"))
		if err != nil {
			return
		}

		version, err := semver.NewVersion(versionData.Version)
		if err != nil {
			return
		}

		if ctx.Bool("major") {
			*version = version.IncMajor()
		} else if ctx.Bool("minor") {
			*version = version.IncMinor()
		} else if ctx.Bool("patch") {
			*version = version.IncPatch()
		} else {
			return fmt.Errorf("no version bump specified")
		}

		versionData.Version = "v" + version.String()
		if err = gen.UpdatePackageVersion(ctx.Path("file"), version.String()); err != nil {
			return
		}

		if ctx.Bool("verbose") {
			log.Printf("bumped version to %s", versionData.Version)
		}

		return
	},
}
