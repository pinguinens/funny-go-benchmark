package buffer

import "github.com/pinguinens/funny-go-benchmark/pkg/units"

type Buffer struct {
	size   int
	buffer chan [][units.Kilobyte]byte
}

func New(size, count int) *Buffer {
	return &Buffer{
		size:   units.Kilobyte * size,
		buffer: make(chan [][units.Kilobyte]byte, count),
	}
}

func (b *Buffer) Size() int {
	return b.size
}

func (b *Buffer) Input(ch [][units.Kilobyte]byte) {
	b.buffer <- ch
}
