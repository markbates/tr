package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"os/exec"
	"syscall"

	"github.com/markbates/tt/cmd/models"
	"github.com/mattn/go-zglob"
)

func init() {
	err := models.Connect()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

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

	var bb bytes.Buffer
	w := io.MultiWriter(os.Stdout, &bb)

	cmd.Stdout = w
	cmd.Stderr = os.Stderr
	cmd.Start()
	err := cmd.Wait()
	if err != nil {
		h.Error = err.Error()
	}

	h.Results = bb.String()
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
