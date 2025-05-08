package command

import (
	"fmt"
	"strings"
)

const COMMAND_ECHO = "echo"

type Echo struct {
	s string
}

func (e Echo) Exec() {
	fmt.Println(e.s)
}

func NewEcho(s string) Echo {
	return Echo{s}
}

func NewEchoWithArgs(args []string) Echo {
	arg := strings.Join(args, " ")
	return NewEcho(arg)
}

func init() {
	RegisterCommand(COMMAND_ECHO, func(args []string) Command {
		return NewEchoWithArgs(args)
	})
}
