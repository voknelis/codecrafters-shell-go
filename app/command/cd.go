package command

import (
	"fmt"
	"os"
)

var COMMAND_CD = "cd"

type CD struct {
	path string
}

func (c CD) Exec() {
	info, err := os.Stat(c.path)
	if err != nil {
		fmt.Printf("cd: %s: No such file or directory\n", c.path)
		return
	}

	if info.IsDir() {
		cwd = c.path
	} else {
		fmt.Printf("cd: %s: Not a directory\n", c.path)
	}
}

func NewCD(path string) CD {
	return CD{path}
}
