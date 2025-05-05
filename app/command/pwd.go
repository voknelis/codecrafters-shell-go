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

func (Pwd) Exec() {
	fmt.Println(cwd)
}

func NewPwd() Pwd {
	return Pwd{}
}
