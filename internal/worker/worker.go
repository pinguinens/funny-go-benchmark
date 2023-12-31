package worker

import (
	"unsafe"

	"github.com/pinguinens/funny-go-benchmark/internal/buffer"

	"github.com/pinguinens/funny-go-benchmark/pkg/units"
)

type Worker struct {
	id     int
	buffer [][units.Kilobyte]byte
}

func New(id, bufferSize int) *Worker {
	return &Worker{
		id:     id,
		buffer: make([][units.Kilobyte]byte, 0, bufferSize),
	}
}

func (w *Worker) Run(b *buffer.Buffer) int {
	var (
		arr   [units.Kilobyte]byte
		wSize int
	)

	for i := 0; i < cap(w.buffer); i++ {
		for _, in := range arr {
			wSize += int(unsafe.Sizeof(in))
		}
		w.buffer = append(w.buffer, arr)
	}

	b.Input(w.buffer)
	return wSize
}

func (w *Worker) ID() int {
	return w.id
}

func (w *Worker) Buffer() *[][units.Kilobyte]byte {
	return &w.buffer
}

func (w *Worker) BufferCap() int {
	return cap(w.buffer)
}
