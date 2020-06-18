package envgit

import (
	"errors"
	"github.com/urfave/cli/v2"
	"os"
	"os/exec"
)

type Command struct {
	Name string
	Args []string
}

func GetCommand(c *cli.Context) *Command {
	args := c.Args().Slice()
	if len(args) == 0 {
		cli.ShowAppHelp(c)
		os.Exit(1)
	}

	cmd := &Command{
		Name: args[0],
		Args: args[1:],
	}

	return cmd
}

// Run the command in given environment
func (c Command) Execute(env Environment) error {
	cmd := exec.Command(c.Name, c.Args...)
	cmd.Env = env

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()

	// we need to forward the child's process exit code
	// and return it. We would like to return exactly same
	// code as child code
	var eerr *exec.ExitError
	if errors.As(err, &eerr) {
		return cli.Exit(nil, eerr.ExitCode())
	}

	return nil
}