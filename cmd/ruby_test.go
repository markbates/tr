package cmd

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_RubyBuilder(t *testing.T) {
	r := require.New(t)
	oe := Exists
	Exists = func(path string) bool {
		return path == "test/foo_test.rb"
	}
	defer func() { Exists = oe }()
	cmd := RubyBuilder([]string{"foo.rb"})
	r.Equal("ruby -Itest -Ispec test/foo_test.rb", cmd.String())
}

func Test_RubyBuilder_WithBundler(t *testing.T) {
	r := require.New(t)
	os.Setenv("GEM_HOME", "/tmp")
	oe := Exists
	Exists = func(path string) bool {
		switch path {
		case "Gemfile", "test/foo_test.rb":
			return true
		}
		return false
	}
	defer func() { Exists = oe }()
	cmd := RubyBuilder([]string{"foo.rb"})
	r.Equal("/tmp/bin/bundle exec ruby -Itest -Ispec test/foo_test.rb", cmd.String())
}

func Test_RubyBuilder_FileChecker(t *testing.T) {
	r := require.New(t)
	paths := []string{}
	oe := Exists
	Exists = func(path string) bool {
		paths = append(paths, path)
		return false
	}
	defer func() { Exists = oe }()
	RubyBuilder([]string{"app/foo.rb"})
	r.Equal([]string{"Gemfile", "app/foo_test.rb", "app/foo_spec.rb", "test/foo_test.rb", "spec/foo_spec.rb"}, paths)
}
