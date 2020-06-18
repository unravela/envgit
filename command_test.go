package envgit_test

import (
	"github.com/unravela/envgit"
	"runtime"
)

func ExampleCommand_Execute() {
	//determine if we're running win or linux
	script := "./testdata/script.sh"
	if runtime.GOOS == "windows" {
		script = "testdata\\script.bat"
	}

	// setup environment
	env := envgit.Environment{"TEST_VAR=hello world"}

	// execute 'script.sh' with this env.
	cmd := envgit.Command{
		Name: script,
		Args: []string{},
	}
	cmd.Execute(env)

	// Output:
	// var:hello world
}
