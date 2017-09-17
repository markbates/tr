package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"syscall"

	"github.com/markbates/tt/cmd/models"
	zglob "github.com/mattn/go-zglob"
)

var Exists = func(path string) bool {
	m, err := zglob.Glob(path)
	return err == nil && len(m) > 0
}

func Exit(err error) {
	os.Exit(exitStatus(err))
}

func Run(cmd *Cmd) {
	h := &models.History{
		CmdArgs: cmd.Args,
	}
	fmt.Println(cmd.String())

	// w := io.MultiWriter(os.Stdout, &bb)
	// w := io.MultiWriter(&bb, os.Stdout)

	cmd.Stdin = os.Stdin
	// cmd.Stdout = w
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		h.Error = err.Error()
	}

	// h.Results = string(b)
	h.ExitCode = exitStatus(err)
	err = h.Save()

	if err != nil {
		Exit(err)
	}
	os.Exit(h.ExitCode)
}

func exitStatus(err error) int {
	if err != nil {
		if exiterr, ok := err.(*exec.ExitError); ok {
			if status, ok := exiterr.Sys().(syscall.WaitStatus); ok {
				return status.ExitStatus()
			}
		}
		return 1
	}
	return 0
}

func printJSON(v interface{}) {
	b, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		fmt.Println(err)
		Exit(err)
	}
	os.Stdout.Write(b)
}
