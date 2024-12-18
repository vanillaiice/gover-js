package cmd

import (
	"os"
	"os/exec"
	"strings"
	"text/template"

	"github.com/urfave/cli/v2"
	"github.com/vanillaiice/gover-js/load"
)

// tagCmdTemplateData	is the template data for the tag command.
type tagCmdTemplateData struct {
	Version string
}

// generateTagCommand generates the tag command from the template.
func generateTagCommand(tmpl, version string) (string, error) {
	template, err := template.New("tmpl").Parse(tmpl)
	if err != nil {
		return "", err
	}

	var b strings.Builder
	if err = template.Execute(&b, tagCmdTemplateData{
		Version: version,
	}); err != nil {
		return "", err
	}

	return b.String(), nil
}

// tagCmd is the tag command.
var tagCmd = &cli.Command{
	Name:    "tag",
	Aliases: []string{"t"},
	Usage:   "tag branch with the current version",
	Flags: []cli.Flag{
		&cli.PathFlag{
			Name:    "file",
			Aliases: []string{"f"},
			Usage:   "load version from `FILE`",
			Value:   "package.json",
			EnvVars: []string{"VERSION_FILE"},
		},
		&cli.StringFlag{
			Name:    "command",
			Aliases: []string{"c"},
			Usage:   "template for tag `COMMAND`",
			Value:   "git tag {{ .Version }}",
			EnvVars: []string{"TAG_COMMAND"},
		},
	},
	Action: func(ctx *cli.Context) (err error) {
		versionData, err := load.FromFile(ctx.Path("file"))
		if err != nil {
			return
		}

		cmdString, err := generateTagCommand(ctx.String("command"), versionData.Version)
		if err != nil {
			return
		}
		cmdStringParts := strings.Split(cmdString, " ")

		var cmd *exec.Cmd
		if len(cmdStringParts) == 1 {
			cmd = exec.Command(cmdStringParts[0])
		} else {
			cmd = exec.Command(cmdStringParts[0], cmdStringParts[1:]...)
		}

		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		if err = cmd.Run(); err != nil {
			return
		}

		return nil
	},
}
