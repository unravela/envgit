package envgit

import (
	"github.com/go-git/go-billy/v5/memfs"
	"testing"
)

func Test_OpenRepo(t *testing.T) {
	// when we try to open existing repository
	fs,err := openRepo(&Options{
		URL: "https://github.com/unravela/gitwrk",
		Branch: "main",
	})

	// then we shouldn get the error
	if err != nil {
		t.FailNow()
	}

	// ... and we can reach the file
	f,err := fs.Open(".gitignore")
	if err != nil && f == nil {
		t.FailNow()
	}
	f.Close()
}

func Test_ReadVariables(t *testing.T) {
	// given filesystem with 'my.env' file
	fs := memfs.New()
	f, _ := fs.Create("my.env")
	f.Write([]byte("PARAM1=1\nPARAM2=\"2\""))
	f.Close()

	// when we read variables from the file
	f, _ = fs.Open("my.env")
	vars, err := readEnvVariables(f)

	// then we should have 2 variables loaded
	if err != nil || len(vars) != 2 {
		t.FailNow()
	}
}
