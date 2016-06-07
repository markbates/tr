package cmd

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_GoBuilder(t *testing.T) {
	r := require.New(t)
	os.Setenv("GO_ENV", "")
	cmd := GoBuilder([]string{"-v", "-race"})
	r.Equal("go test -v -race github.com/markbates/tt/cmd github.com/markbates/tt/cmd/models", cmd.String())
	r.Equal("test", os.Getenv("GO_ENV"))
}

func Test_GoBuilder_RunFlag(t *testing.T) {
	r := require.New(t)

	cmd := GoBuilder([]string{"-run", "Hello", "./foo"})
	r.Equal("go test -run Hello ./foo", cmd.String())
}
