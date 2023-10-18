package test

import (
	"hash/crc32"
)

type CRC32 struct{}

func (t *CRC32) Name() string {
	return "CRC32"
}

func (t *CRC32) Exec(b []byte) (int, []byte, error) {
	return compute(crc32.New(crc32.IEEETable), b)
}
