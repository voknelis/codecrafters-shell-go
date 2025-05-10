package redirect

import (
	"os"
)

type RedirectStd struct {
	target string
	append bool
}

func (r *RedirectStd) Write(p []byte) (n int, err error) {
	fileFlags := os.O_RDWR | os.O_CREATE | os.O_SYNC
	if r.append {
		fileFlags = fileFlags | os.O_APPEND
	} else {
		fileFlags = fileFlags | os.O_TRUNC
	}

	f, err := os.OpenFile(r.target, fileFlags, os.ModePerm)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	n, err = f.Write(p)
	return n, err
}

func NewRedirectStd(target string, append bool) *RedirectStd {
	return &RedirectStd{target, append}
}
