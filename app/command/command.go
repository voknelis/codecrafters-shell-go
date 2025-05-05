package command

import (
	"errors"
	"strings"
)

type Command interface {
	Exec()
}

var ErrUnknownCommand = errors.New("unknown command")

func NewCommand(input string) (Command, error) {
	commandAndArgs := strings.SplitN(input, " ", 2)
	if len(commandAndArgs) == 0 {
		return nil, ErrUnknownCommand
	}

	command := commandAndArgs[0]
	rawArgs := ""
	if len(commandAndArgs) > 1 {
		rawArgs = commandAndArgs[1]
	}

	args := make([]string, 0)
	argsParts := strings.Split(rawArgs, " ")

	for _, part := range argsParts {
		trimmed := strings.TrimSpace(part)
		if trimmed != "" {
			args = append(args, trimmed)
		}
	}

	switch {
	case strings.HasPrefix(command, COMMAND_EXIT):
		return NewExitWithArgs(args), nil
	case strings.HasPrefix(command, COMMAND_ECHO):
		return NewEcho(rawArgs), nil
	default:
		return nil, ErrUnknownCommand
	}
}
