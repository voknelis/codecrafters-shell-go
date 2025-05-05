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

func (e External) Exec() {
	cmd := exec.Command(e.command, e.args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	cmd.Run()
}

func NewExternalCommand(command string, args []string) (*External, error) {
	_, ok := IsExecutableCommand(command)
	if !ok {
		return nil, ErrUnknownExternalCommand
	}

	return &External{command, args}, nil
}
