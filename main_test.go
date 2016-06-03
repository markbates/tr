package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Exists(t *testing.T) {
	r := require.New(t)
	r.True(Exists("main.go"))
	r.False(Exists("idontexist.go"))
}

func Test_RunMakeFile(t *testing.T) {
	r := require.New(t)
	cmd := RunMakefile([]string{"foo"})
	r.Equal("make test foo", cmd.String())
}

func Test_RunTestSH(t *testing.T) {
	r := require.New(t)
	cmd := RunTestSH([]string{"foo"})
	r.Equal("./test.sh foo", cmd.String())
}

func Test_RunRakefile(t *testing.T) {
	r := require.New(t)
	cmd := RunRakefile([]string{"foo"})
	r.Equal("rake foo", cmd.String())
}

func Test_RunBundler(t *testing.T) {
	r := require.New(t)
	os.Setenv("GEM_HOME", "/tmp")
	cmd := RunBundler([]string{"foo"})
	r.Equal("/tmp/bin/bundle exec rake foo", cmd.String())
}

func Test_RunGoTests(t *testing.T) {
	r := require.New(t)
	os.Setenv("GO_ENV", "")
	oe := Exists
	defer func() { Exists = oe }()
	Exists = func(path string) bool {
		return !(path == "glide.lock")
	}

	cmd := RunGoTests([]string{"-v", "-race"})
	r.Equal("go test -v -race ./...", cmd.String())
	r.Equal("test", os.Getenv("GO_ENV"))
}

func Test_RunGoTests_Glide(t *testing.T) {
	r := require.New(t)
	os.Setenv("GO_ENV", "")
	oe := Exists
	defer func() { Exists = oe }()
	Exists = func(path string) bool {
		return path == "glide.lock"
	}

	cmd := RunGoTests([]string{"-v", "-race"})
	r.Equal("go test -v -race . ", cmd.String())
	r.Equal("test", os.Getenv("GO_ENV"))
}

func Test_RunGoTests_RunFlag(t *testing.T) {
	r := require.New(t)

	cmd := RunGoTests([]string{"-run", "Hello", "./foo"})
	r.Equal("go test -run Hello ./foo", cmd.String())
}
