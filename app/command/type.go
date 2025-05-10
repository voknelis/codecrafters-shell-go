package command

import (
	"fmt"
	"strings"
)

const COMMAND_TYPE = "type"

type Type struct {
	command string
}

func (t Type) Exec(stdout, stderr Writer) error {
	if t.command == "" {
		return nil
	}

	if IsBuiltinCommand(t.command) {
		_, err := fmt.Fprintln(stdout, t.command, "is a shell builtin")
		return err
	} else if cmdPath, ok := IsExecutableCommand(t.command); ok {
		_, err := fmt.Fprintln(stdout, t.command, "is", cmdPath)
		return err
	} else {
		_, err := fmt.Fprintln(stderr, t.command+": not found")
		return err
	}
}

func NewType(command string) Type {
	return Type{command}
}

func NewTypeWithArgs(args []string) Type {
	arg := strings.Join(args, " ")
	return NewType(arg)
}

func init() {
	RegisterCommand(COMMAND_TYPE, func(args []string) CommandExec {
		return NewTypeWithArgs(args)
	})
}
