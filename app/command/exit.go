package command

import (
	"os"
	"strconv"
)

const COMMAND_EXIT = "exit"

type Exit struct {
	code int
}

func (e Exit) Exec() {
	os.Exit(e.code)
}

func NewExit(code int) Exit {
	return Exit{code}
}

func NewExitWithArgs(args []string) Exit {
	code := 0

	if len(args) != 0 {
		exitCodeRaw := args[0]

		parsedCode, err := strconv.Atoi(exitCodeRaw)
		if err == nil {
			code = parsedCode
		}
	}

	return NewExit(code)
}

func init() {
	RegisterCommand(COMMAND_EXIT, func(args []string) Command {
		return NewExitWithArgs(args)
	})
}
