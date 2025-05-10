package command

import (
	"fmt"
	"os"
)

// pwd stands for "print working directory"
const COMMAND_PWD = "pwd"

// current working directory
var cwd string

func init() {
	cwd, _ = os.Getwd()
}

type Pwd struct{}

func (Pwd) Exec(stdout, stderr Writer) error {
	_, err := fmt.Fprintln(stdout, cwd)
	return err
}

func NewPwd() Pwd {
	return Pwd{}
}

func init() {
	RegisterCommand(COMMAND_PWD, func(args []string) CommandExec {
		return NewPwd()
	})
}
