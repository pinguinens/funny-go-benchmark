package runner

import (
	"github.com/pinguinens/funny-go-benchmark/internal/buffer"
	"github.com/pinguinens/funny-go-benchmark/internal/config"
	"github.com/pinguinens/funny-go-benchmark/internal/tester"

	"github.com/pinguinens/funny-go-benchmark/pkg/log"
)

type Runner struct {
	logger   *log.Logger
	buffer   *buffer.Buffer
	routines int
}

func New(c *config.App, l *log.Logger) *Runner {
	return &Runner{
		logger:   l,
		buffer:   buffer.New(c.Buffer.Multiplier, c.System.CoreCount),
		routines: c.System.CoreCount,
	}
}

func (s *Runner) Run(test tester.ITester) {
	s.logger.Info("== Funny Go Benchmark v.0.2 ==")

	test.Init(s.logger, s.buffer, s.routines)
	test.Run()
}
