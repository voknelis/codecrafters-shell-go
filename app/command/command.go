package command

import (
	"errors"
	"os"
	"path/filepath"
	"runtime"
	"slices"
	"strings"

	"github.com/codecrafters-io/shell-starter-go/app/tokenizer"
)

type Command interface {
	Exec()
}

var ErrUnknownCommand = errors.New("unknown command")

var builtinCommands = []string{
	COMMAND_EXIT,
	COMMAND_ECHO,
	COMMAND_TYPE,
	COMMAND_PWD,
	COMMAND_CD,
}

func IsBuiltinCommand(command string) bool {
	return slices.Contains(builtinCommands, command)
}

func IsExecutableCommand(command string) (string, bool) {
	pathEnv := os.Getenv("PATH")
	paths := strings.Split(pathEnv, string(os.PathListSeparator))

	extensions := []string{""}
	// On Windows, check PATHEXT for valid extensions
	if runtime.GOOS == "windows" {
		pathext := os.Getenv("PATHEXT")
		if pathext != "" {
			extensions = strings.Split(strings.ToLower(pathext), ";")
		}
	}

	for _, dir := range paths {
		for _, ext := range extensions {
			cmdPath := filepath.Join(dir, command)
			if ext != "" {
				cmdPath += ext
			}

			info, err := os.Stat(cmdPath)
			if err == nil && !info.IsDir() {
				return cmdPath, true
			}
		}
	}

	return "", false
}

func NewCommand(input string) (Command, error) {
	commandAndArgs := strings.SplitN(input, " ", 2)
	if len(commandAndArgs) == 0 {
		return nil, ErrUnknownCommand
	}

	command := strings.TrimSpace(commandAndArgs[0])
	rawArgs := ""
	if len(commandAndArgs) > 1 {
		rawArgs = commandAndArgs[1]
	}

	args := tokenizer.Tokenize(rawArgs)

	switch {
	case strings.HasPrefix(command, COMMAND_EXIT):
		return NewExitWithArgs(args), nil
	case strings.HasPrefix(command, COMMAND_ECHO):
		return NewEchoWitArgs(args), nil
	case strings.HasPrefix(command, COMMAND_TYPE):
		return NewType(rawArgs), nil
	case strings.HasPrefix(command, COMMAND_PWD):
		return NewPwd(), nil
	case strings.HasPrefix(command, COMMAND_PWD):
		return NewPwd(), nil
	case strings.HasPrefix(command, COMMAND_CD):
		return NewCD(rawArgs), nil
	default:
		cmd, err := NewExternalCommand(command, args)
		if err == nil {
			return cmd, nil
		}

		return nil, ErrUnknownCommand
	}
}
