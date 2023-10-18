package test

import (
	"crypto/sha512"
)

type SHA512 struct{}

func (t *SHA512) Name() string {
	return "SHA-512"
}

func (t *SHA512) Exec(b []byte) (int, []byte, error) {
	return compute(sha512.New(), b)
}
