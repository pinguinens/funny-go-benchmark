package test

import (
	"hash/fnv"
)

type FNV32 struct{}

func (t *FNV32) Name() string {
	return "FNV-1 32"
}

func (t *FNV32) Exec(b []byte) (int, []byte, error) {
	return compute(fnv.New32(), b)
}

type FNV32a struct{}

func (t *FNV32a) Name() string {
	return "FNV-1 32a"
}

func (t *FNV32a) Exec(b []byte) (int, []byte, error) {
	return compute(fnv.New32a(), b)
}

type FNV64 struct{}

func (t *FNV64) Name() string {
	return "FNV-1 64"
}

func (t *FNV64) Exec(b []byte) (int, []byte, error) {
	return compute(fnv.New64(), b)
}

type FNV64a struct{}

func (t *FNV64a) Name() string {
	return "FNV-1 64a"
}

func (t *FNV64a) Exec(b []byte) (int, []byte, error) {
	return compute(fnv.New64a(), b)
}

type FNV128 struct{}

func (t *FNV128) Name() string {
	return "FNV-1 128"
}

func (t *FNV128) Exec(b []byte) (int, []byte, error) {
	return compute(fnv.New128(), b)
}

type FNV128a struct{}

func (t *FNV128a) Name() string {
	return "FNV-1 128a"
}

func (t *FNV128a) Exec(b []byte) (int, []byte, error) {
	return compute(fnv.New128a(), b)
}
