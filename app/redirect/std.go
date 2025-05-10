package redirect

import (
	"os"
)

type RedirectStd struct {
	target string
}

func (r *RedirectStd) Write(p []byte) (n int, err error) {
	f, err := os.OpenFile(r.target, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, os.ModePerm)
	if err != nil {
		return 0, err
	}

	n, err = f.Write(p)
	if err1 := f.Close(); err1 != nil && err == nil {
		err = err1
	}

	return n, err
}

func NewRedirectStd(target string) *RedirectStd {
	return &RedirectStd{target}
}
