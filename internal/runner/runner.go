package runner

import (
	"github.com/pinguinens/funny-go-benchmark/internal/buffer"
	"github.com/pinguinens/funny-go-benchmark/internal/config"
	"github.com/pinguinens/funny-go-benchmark/internal/tester/memory"

	"github.com/pinguinens/funny-go-benchmark/pkg/log"
)

type Service struct {
	logger   *log.Logger
	buffer   *buffer.Buffer
	routines int
}

func New(c *config.App, l *log.Logger) *Service {
	return &Service{
		logger:   l,
		buffer:   buffer.New(c.Buffer.Multiplier, c.System.CoreCount),
		routines: c.System.CoreCount,
	}
}

func (s *Service) Init() {
}

func (s *Service) Run() {
	s.logger.Info("== Funny Go Benchmark v.0.2 ==")

	memTest := memory.New(s.logger, s.buffer, s.routines)
	memTest.Run()
}
