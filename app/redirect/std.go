package redirect

import (
	"os"
)

type RedirectStd struct {
	target string
}

func (r *RedirectStd) Write(p []byte) (n int, err error) {
	f, err := os.OpenFile(r.target, os.O_RDWR|os.O_CREATE|os.O_SYNC, os.ModePerm)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	n, err = f.Write(p)
	return n, err
}

func NewRedirectStd(target string) *RedirectStd {
	return &RedirectStd{target}
}
