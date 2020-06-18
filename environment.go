package envgit

import (
	"fmt"
	"github.com/go-git/go-billy/v5"
	"github.com/go-git/go-billy/v5/memfs"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/transport"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
	"github.com/go-git/go-git/v5/storage/memory"
	"io/ioutil"
	"os"
	"strings"
)

// Environment hold the env. variables as array of strings 'key=value'
type Environment []string

// OsEnvironment returns you current environment variables
// we will merge with those from Git.
func OsEnvironment() Environment {
	return os.Environ()
}

// LoadGitEnvVars clone the repository into memory, read the file and
// returns you variables from this file.
//
// To prevent full slow clone, I'm cloning just first commit history
// in specific branch. Anyway could be very bad idea put the configuration
// into monorepo with thousands and hundreds of large files.
//
func LoadGitEnvVars(opts *Options) Environment {
	// clone only 1th commit and open the repo
	fs, err := openRepo(opts)
	if err != nil {
		fmt.Printf("cannot clone repository (check the --url option): %v", err)
		os.Exit(1)
	}

	// read the env variables from file
	f, err := fs.Open(opts.File)
	if err != nil {
		fmt.Errorf("cannot open file. Check the --file option %v", err)
		os.Exit(1)

	}
	defer f.Close()

	// read the env. variables
	envVars, err := readEnvVariables(f)
	if err != nil {
		fmt.Printf("cannot read env. variables from file %s: %v", opts.File, err)
		os.Exit(1)
	}

	return envVars
}


// function determine what auth. method to use. Currently
// I'm supporting just simple Basic auth or none. Basic auth
// is perfectly fine also for GitHub/GitLab private repos.
// All you need is just create OAuth token and use it as
// password.
func getAuthMethod(o *Options) transport.AuthMethod {
	if o.Username != "" || o.Password != "" {
		return &http.BasicAuth{
			Username: o.Username,
			Password: o.Password,
		}
	}
	return nil
}

func openRepo(opts *Options) (billy.Filesystem, error) {
	cloneOpts := &git.CloneOptions{
		URL:           opts.URL,
		Auth:          getAuthMethod(opts),
		ReferenceName: plumbing.ReferenceName("refs/heads/" + opts.Branch),
		SingleBranch:  true,
		Depth:         1,
	}

	fs := memfs.New()
	_, err := git.Clone(memory.NewStorage(), fs, cloneOpts)
	if err != nil {
		return nil, fmt.Errorf("cannot clone repository (check the --url option): %w", err)
	}

	return fs, nil
}

// function read the file on virtual Filesystem and parse variables
func readEnvVariables(f billy.File) ([]string, error) {
	data,err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}

	vars := strings.Split(string(data), "\n")
	return vars, nil
}