package command

import (
	"fmt"
	"os"
)

// pwd stands for "print working directory"
const COMMAND_PWD = "pwd"

// current working directory
var cwd string

type Pwd struct{}

func (Pwd) Exec() {
	var path string
	if cwd != "" {
		path = cwd
	} else {
		path, _ = os.Getwd()
	}

	fmt.Println(path)
}

func NewPwd() Pwd {
	return Pwd{}
}
