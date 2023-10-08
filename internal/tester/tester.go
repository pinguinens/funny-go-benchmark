package tester

import (
	"github.com/pinguinens/funny-go-benchmark/internal/buffer"

	"github.com/pinguinens/funny-go-benchmark/pkg/log"
)

type ITester interface {
	Init(logger *log.Logger, buffer *buffer.Buffer, routines int)
	Run()
}
