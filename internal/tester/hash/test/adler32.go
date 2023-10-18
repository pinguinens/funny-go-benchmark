package test

import "hash/adler32"

type Adler32 struct{}

func (t *Adler32) Name() string {
	return "Adler 32"
}

func (t *Adler32) Exec(b []byte) (int, []byte, error) {
	return compute(adler32.New(), b)
}
