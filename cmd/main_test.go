package cmd

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Exists(t *testing.T) {
	r := require.New(t)
	r.True(Exists("root.go"))
	r.False(Exists("idontexist.go"))
	// r.True(false)
}

func Test_RunMakeFile(t *testing.T) {
	r := require.New(t)
	cmd := MakefileBuilder([]string{"foo"})
	r.Equal("make test foo", cmd.String())
}

func Test_RunTestSH(t *testing.T) {
	r := require.New(t)
	cmd := TestSHBuilder([]string{"foo"})
	r.Equal("./test.sh foo", cmd.String())
}

func Test_RunRakefile(t *testing.T) {
	r := require.New(t)
	cmd := RakefileBuilder([]string{"foo"})
	r.Equal("rake foo", cmd.String())
}

func Test_RunRakefile_WithBundler(t *testing.T) {
	r := require.New(t)
	os.Setenv("GEM_HOME", "/tmp")
	oe := Exists
	Exists = func(path string) bool {
		return path == "Gemfile"
	}
	defer func() { Exists = oe }()
	cmd := RakefileBuilder([]string{"foo"})
	r.Equal("/tmp/bin/bundle exec rake foo", cmd.String())
}

func Test_RunGoTests(t *testing.T) {
	r := require.New(t)
	os.Setenv("GO_ENV", "")
	cmd := GoBuilder([]string{"-v", "-race"})
	r.Equal("go test -v -race github.com/markbates/tt/cmd github.com/markbates/tt/cmd/models", cmd.String())
	r.Equal("test", os.Getenv("GO_ENV"))
}

func Test_RunGoTests_RunFlag(t *testing.T) {
	r := require.New(t)

	cmd := GoBuilder([]string{"-run", "Hello", "./foo"})
	r.Equal("go test -run Hello ./foo", cmd.String())
}
