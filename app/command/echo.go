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
	var s string
	if len(args) != 0 {
		s = strings.Join(args, " ")
	}

	return NewEcho(s)
}
