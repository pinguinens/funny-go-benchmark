package test

import (
	"hash/crc64"
)

type CRC64 struct{}

func (t *CRC64) Name() string {
	return "CRC64"
}

func (t *CRC64) Exec(b []byte) (int, []byte, error) {
	return compute(crc64.New(crc64.MakeTable(crc64.ISO)), b)
}
