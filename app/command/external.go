package command

import (
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
	cmd.Stdout = stdout
	cmd.Stderr = stderr

	cmd.Run()
	return nil
}

func NewExternalCommand(command string, args []string) (*External, error) {
	_, ok := IsExecutableCommand(command)
	if !ok {
		return nil, ErrUnknownExternalCommand
	}

	return &External{command, args}, nil
}
