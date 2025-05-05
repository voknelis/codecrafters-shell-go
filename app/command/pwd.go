package command

import (
	"fmt"
	"os"
)

// pwd stands for "print working directory"
const COMMAND_PWD = "pwd"

type Pwd struct{}

func (Pwd) Exec() {
	path, _ := os.Getwd()
	fmt.Println(path)
}

func NewPwd() Pwd {
	return Pwd{}
}
