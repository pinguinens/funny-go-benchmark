package test

import (
	"hash"
)

func compute(h hash.Hash, b []byte) (int, []byte, error) {
	n, err := h.Write(b)
	if err != nil {
		return 0, nil, err
	}
	r := h.Sum(nil)

	return n, r, nil
}
