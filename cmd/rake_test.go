package cmd

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_RakefileBuilder(t *testing.T) {
	r := require.New(t)
	cmd := RakefileBuilder([]string{"foo"})
	r.Equal("rake foo", cmd.String())
}

func Test_RakefileBuilder_WithBundler(t *testing.T) {
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
