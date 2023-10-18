package test

import (
	"crypto/md5"
)

type MD5 struct{}

func (t *MD5) Name() string {
	return "MD5"
}

func (t *MD5) Exec(b []byte) (int, []byte, error) {
	return compute(md5.New(), b)
}
