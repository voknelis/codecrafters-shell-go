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

func (c CD) Exec(stdout, stderr Writer) error {
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
		_, err := fmt.Fprintf(stderr, "cd: %s: No such file or directory\n", path)
		return err
	}

	if info.IsDir() {
		err := os.Chdir(path)
		if err != nil {
			_, err := fmt.Fprintf(stderr, "cd: failed to change directory: %s\n", err)
			return err
		}

		cwd = path
	} else {
		_, err := fmt.Fprintf(stderr, "cd: %s: Not a directory\n", path)
		if err != nil {
			return err
		}
	}

	return nil
}

func NewCD(path string) CD {
	return CD{path}
}

func NewCDWithArgs(args []string) CD {
	arg := strings.Join(args, " ")
	return NewCD(arg)
}

func init() {
	RegisterCommand(COMMAND_CD, func(args []string) CommandExec {
		return NewCDWithArgs(args)
	})
}
