package command

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var COMMAND_CD = "cd"

type CD struct {
	path string
}

func (c CD) Exec() {
	path := c.path
	// user directory is default directory
	if path == "" {
		path = "~"
	}

	isAbs := filepath.IsAbs(path)
	if !isAbs {
		// handle home directory path
		if strings.HasPrefix(path, "~") {
			userHomeDir, _ := os.UserHomeDir()

			if path == "~" {
				path = userHomeDir
			} else {
				path = strings.Replace(path, "~", userHomeDir, 1)
			}
		} else { // other relative path
			path = filepath.Join(cwd, path)
		}
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

func NewCDWithArgs(args []string) CD {
	arg := strings.Join(args, " ")
	return NewCD(arg)
}

func init() {
	RegisterCommand(COMMAND_CD, func(args []string) Command {
		return NewCDWithArgs(args)
	})
}
