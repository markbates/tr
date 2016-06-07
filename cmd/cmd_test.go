package cmd

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Exists(t *testing.T) {
	r := require.New(t)
	r.True(Exists("root.go"))
	r.False(Exists("idontexist.go"))
}
