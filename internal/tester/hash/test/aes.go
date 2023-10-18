package test

import (
	"crypto/aes"
)

type AES128 struct{}

func (t *AES128) Name() string {
	return "AES-128"
}

func (t *AES128) Exec(b []byte) (int, []byte, error) {
	key := "thisis16bitlongp"
	c, err := aes.NewCipher([]byte(key))
	if err != nil {
		return 0, nil, err
	}

	l := len(b)
	out := make([]byte, l)
	c.Encrypt(out, b)

	return l, out, nil
}

type AES192 struct{}

func (t *AES192) Name() string {
	return "AES-192"
}

func (t *AES192) Exec(b []byte) (int, []byte, error) {
	key := "thisis24bitlongpassphras"
	c, err := aes.NewCipher([]byte(key))
	if err != nil {
		return 0, nil, err
	}

	l := len(b)
	out := make([]byte, l)
	c.Encrypt(out, b)

	return l, out, nil
}

type AES256 struct{}

func (t *AES256) Name() string {
	return "AES-256"
}

func (t *AES256) Exec(b []byte) (int, []byte, error) {
	key := "thisis32bitlongpassphraseimusing"
	c, err := aes.NewCipher([]byte(key))
	if err != nil {
		return 0, nil, err
	}

	l := len(b)
	out := make([]byte, l)
	c.Encrypt(out, b)

	return l, out, nil
}
