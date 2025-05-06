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

func NewEchoWitArgs(args []string) Echo {
	s := strings.Join(args, " ")
	return NewEcho(s)
}
