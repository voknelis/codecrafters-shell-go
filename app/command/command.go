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
	rawParts := strings.Split(input, " ")
	parts := make([]string, 0)

	for _, part := range rawParts {
		trimmed := strings.TrimSpace(part)
		if trimmed != "" {
			parts = append(parts, trimmed)
		}
	}

	if len(parts) == 0 {
		return nil, ErrUnknownCommand
	}

	command := parts[0]
	args := parts[1:]

	switch {
	case strings.HasPrefix(command, COMMAND_EXIT):
		return NewExitWithArgs(args), nil
	default:
		return nil, ErrUnknownCommand
	}
}
