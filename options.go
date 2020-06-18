package envgit

import (
	"github.com/urfave/cli/v2"
	"os"
)

type Options struct {
	URL      string
	File     string
	Branch   string
	Username string
	Password string
}

var (
	Flags = []cli.Flag{
		&cli.StringFlag{
			Name:    "url",
			Value:   "",
			Usage:   "URL to repository where the env. file is",
			EnvVars: []string{"ENVGIT_URL"},
		},
		&cli.StringFlag{
			Name:    "file",
			Value:   "path to env. file in repository",
			EnvVars: []string{"ENVGIT_FILE"},
		},
		&cli.StringFlag{
			Name:    "branch",
			Value:   "main",
			Usage:   "determine what branch you want to use. The default is 'main'. Not 'master'! (blm)",
			EnvVars: []string{"ENVGIT_BRANCH"},
		},
		&cli.StringFlag{
			Name:    "username",
			Value:   "",
			Usage:   "If your repository is private, you have to pass the username and password",
			EnvVars: []string{"ENVGIT_USERNAME"},
		},
		&cli.StringFlag{
			Name:    "password",
			Value:   "",
			Usage:   "Password can be also OAuth token used by GitHub/GitLab",
			EnvVars: []string{"ENVGIT_PASSWORD"},
		},
	}
)

func ValidateAndGetOptions(c *cli.Context) *Options {
	// validate required options
	url := c.String("url")
	if url == "" {
		cli.ShowAppHelp(c)
		os.Exit(1)
	}

	file := c.String("file")
	if file == "" {
		cli.ShowAppHelp(c)
		os.Exit(1)
	}

	opts := &Options{
		URL:      url,
		File:     file,
		Branch:   c.String("branch"),
		Username: c.String("username"),
		Password: c.String("password"),
	}

	return opts
}
