package runner

import (
	"github.com/pinguinens/funny-go-benchmark/internal/buffer"
	"github.com/pinguinens/funny-go-benchmark/internal/config"
	"github.com/pinguinens/funny-go-benchmark/internal/tester"
)

type Runner struct {
	buffer *buffer.Buffer
}

func New(c *config.App) *Runner {
	return &Runner{
		buffer: buffer.New(c.Buffer.Multiplier, c.System.CoreCount),
	}
}

func (r *Runner) Run(test tester.ITester) {
	test.Run()
}

func (r *Runner) Buffer() *buffer.Buffer {
	return r.buffer
}
