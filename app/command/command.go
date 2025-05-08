package command

import (
	"errors"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/codecrafters-io/shell-starter-go/app/tokenizer"
)

var ErrUnknownCommand = errors.New("unknown command")

type Command interface {
	Exec()
}

type CommandHandler func(args []string) Command

var commandRegistry = map[string]CommandHandler{}

func RegisterCommand(name string, handler CommandHandler) {
	commandRegistry[name] = handler
}

func IsBuiltinCommand(command string) bool {
	_, exists := commandRegistry[command]
	return exists
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
	commandAndArgs := tokenizer.Tokenize(input)
	if len(commandAndArgs) == 0 {
		return nil, ErrUnknownCommand
	}

	command := commandAndArgs[0]
	args := commandAndArgs[1:]

	commandHandler, exists := commandRegistry[command]
	if exists {
		return commandHandler(args), nil
	}

	externalCommand, err := NewExternalCommand(command, args)
	if err == nil {
		return externalCommand, nil
	}

	return nil, ErrUnknownCommand
}
