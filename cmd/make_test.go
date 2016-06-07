package cmd

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_MakefileBuilder(t *testing.T) {
	r := require.New(t)
	cmd := MakefileBuilder([]string{"foo"})
	r.Equal("make test foo", cmd.String())
}
