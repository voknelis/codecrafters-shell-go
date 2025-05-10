package command

import (
	"errors"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/codecrafters-io/shell-starter-go/app/parser"
	"github.com/codecrafters-io/shell-starter-go/app/redirect"
	"github.com/codecrafters-io/shell-starter-go/app/tokenizer"
)

var ErrUnknownCommand = errors.New("unknown command")

type Writer = io.Writer

type Command struct {
	Executable  CommandExec
	Redirection []redirect.Redirection
}

type CommandExec interface {
	Exec(stdout, stderr Writer) error
}

type CommandHandler func(args []string) CommandExec

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

func NewCommand(input string) (*Command, error) {
	commandAndArgs := tokenizer.Tokenize(input)
	if len(commandAndArgs) == 0 {
		return nil, ErrUnknownCommand
	}

	commandNode, err := parser.Parse(commandAndArgs)
	if err != nil {
		return nil, err
	}
	if commandNode.Executable == "" {
		return nil, ErrUnknownCommand
	}

	command := &Command{
		Redirection: commandNode.Redirection,
	}

	commandHandler, exists := commandRegistry[commandNode.Executable]
	if exists {
		command.Executable = commandHandler(commandNode.Arguments)
		return command, nil
	}

	externalCommand, err := NewExternalCommand(commandNode.Executable, commandNode.Arguments)
	if err == nil {
		command.Executable = externalCommand
		return command, nil
	}

	return nil, ErrUnknownCommand
}
