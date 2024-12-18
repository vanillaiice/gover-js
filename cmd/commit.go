package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"
	"text/template"

	"github.com/urfave/cli/v2"
	"github.com/vanillaiice/gover-js/load"
)

// commitCmdTemplateData is the template data for the commit command.
type commitCmdTemplateData struct {
	File    string
	Version string
}

// generateCommitCommand generates the commit command from the template.
func generateCommitCommand(tmpl, file, version string) (string, error) {
	template, err := template.New("tmpl").Parse(tmpl)
	if err != nil {
		return "", err
	}

	var b strings.Builder
	if err = template.Execute(&b, commitCmdTemplateData{
		File:    file,
		Version: version,
	}); err != nil {
		return "", err
	}

	return b.String(), nil
}

func splitCommand(command string) []string {
	re := regexp.MustCompile(`(?:[^\s'"]+|['"][^'"]*['"])`)
	matches := re.FindAllString(command, -1)
	for i, match := range matches {
		if len(match) > 1 && (match[0] == '"' || match[0] == '\'') && match[len(match)-1] == match[0] {
			matches[i] = match[1 : len(match)-1]
		}
	}
	return matches
}

// commitCmd is the commit command.
var commitCmd = &cli.Command{
	Name:    "commit",
	Aliases: []string{"c"},
	Usage:   "commit version",
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
			Usage:   "template for commit `COMMAND`",
			Value:   "git commit {{ .File }} -m \"chore: bump version to {{ .Version }}\"",
			EnvVars: []string{"COMMIT_COMMAND"},
		},
	},
	Action: func(ctx *cli.Context) (err error) {
		versionData, err := load.FromFile(ctx.Path("file"))
		if err != nil {
			return err
		}

		cmdString, err := generateCommitCommand(ctx.String("command"), ctx.Path("file"), versionData.Version)
		if err != nil {
			return err
		}
		cmdStringParts := splitCommand(cmdString)

		var cmd *exec.Cmd
		lenCmdStringParts := len(cmdStringParts)
		if lenCmdStringParts == 0 {
			return fmt.Errorf("invalid command: %s", cmdString)
		} else if lenCmdStringParts == 1 {
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
