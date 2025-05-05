package command

import (
	"fmt"
	"os"
	"path/filepath"
)

var COMMAND_CD = "cd"

type CD struct {
	path string
}

func (c CD) Exec() {
	path := c.path
	isAbs := filepath.IsAbs(path)

	if !isAbs {
		path = filepath.Join(cwd, path)
	}

	info, err := os.Stat(path)
	if err != nil {
		fmt.Printf("cd: %s: No such file or directory\n", path)
		return
	}

	if info.IsDir() {
		cwd = path
	} else {
		fmt.Printf("cd: %s: Not a directory\n", path)
	}
}

func NewCD(path string) CD {
	return CD{path}
}
