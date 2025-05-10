package command

import (
	"bytes"
	"errors"
	"os"
	"os/exec"
)

var ErrUnknownExternalCommand = errors.New("unknown extrenal command")

type External struct {
	command string
	args    []string
}

func (e External) Exec(stdout, stderr Writer) error {
	cmd := exec.Command(e.command, e.args...)
	cmd.Stdin = os.Stdin

	var outBuf bytes.Buffer
	var errBuf bytes.Buffer

	cmd.Stdout = &outBuf
	cmd.Stderr = &errBuf

	cmd.Run()

	_, err := stdout.Write(outBuf.Bytes())
	if err != nil {
		return err
	}

	_, err = stderr.Write(errBuf.Bytes())
	return err
}

func NewExternalCommand(command string, args []string) (*External, error) {
	_, ok := IsExecutableCommand(command)
	if !ok {
		return nil, ErrUnknownExternalCommand
	}

	return &External{command, args}, nil
}
