package main

import (
	"github.com/unravela/envgit"
	"github.com/urfave/cli/v2"
	"os"
)

func main() {
	app := &cli.App{
		Name:            "envgit",
		Usage:           "executing command with env.variables loaded from git repository",
		ArgsUsage:       "{command}",
		HideHelpCommand: true,
		Flags:           envgit.Flags,
		Action:          doMain,
	}

	app.Run(os.Args)
}

func doMain(c *cli.Context) error {
	opts := envgit.ValidateAndGetOptions(c)
	cmd := envgit.GetCommand(c)

	osEnv := envgit.OsEnvironment()
	gitEnv := envgit.LoadGitEnvVars(opts)
	env := append(osEnv, gitEnv...)

	return cmd.Execute(env)
}
