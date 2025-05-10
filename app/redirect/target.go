package redirect

import (
	"fmt"
	"os"
	"path/filepath"
)

func CheckRedirectionTarget(target string) error {
	info, err := os.Stat(target)
	if err != nil {
		if !isTargetDirExists(target) {
			return fmt.Errorf("-shell: %s: No such file or directory", target)
		}

		return nil
	}

	if info.IsDir() {
		return fmt.Errorf("-shell: %s: Is a directory", target)
	}

	return nil
}

func isTargetDirExists(target string) bool {
	dir := filepath.Dir(target)
	_, err := os.Stat(dir)
	return err == nil
}
