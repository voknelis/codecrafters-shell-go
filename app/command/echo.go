package command

import (
	"fmt"
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
