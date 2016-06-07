package cmd

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_NPMBuilder(t *testing.T) {
	r := require.New(t)
	cmd := NPMBuilder([]string{"foo"})
	r.Equal("npm test foo", cmd.String())
}
