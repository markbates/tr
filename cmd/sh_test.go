package cmd

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_TestSHBuilder(t *testing.T) {
	r := require.New(t)
	cmd := TestSHBuilder([]string{"foo"})
	r.Equal("./test.sh foo", cmd.String())
}
