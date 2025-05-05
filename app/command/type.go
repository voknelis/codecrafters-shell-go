package command

import (
	"fmt"
)

const COMMAND_TYPE = "type"

type Type struct {
	command string
}

func (t Type) Exec() {
	if t.command == "" {
		return
	}

	if IsBuiltinCommand(t.command) {
		fmt.Println(t.command, "is a shell builtin")
	} else if cmdPath, ok := IsExecutableCommand(t.command); ok {
		fmt.Println(t.command, "is", cmdPath)
	} else {
		fmt.Println(t.command + ": not found")
	}
}

func NewType(command string) Type {
	return Type{command}
}
