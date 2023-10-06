package service

import (
	"sync"

	"github.com/pinguinens/funny-go-benchmark/internal/buffer"
	"github.com/pinguinens/funny-go-benchmark/internal/config"
	"github.com/pinguinens/funny-go-benchmark/internal/worker"

	"github.com/pinguinens/funny-go-benchmark/pkg/log"
	"github.com/pinguinens/funny-go-benchmark/pkg/timer"
	"github.com/pinguinens/funny-go-benchmark/pkg/units"
)

type Service struct {
	logger  *log.Logger
	buffer  *buffer.Buffer
	workers []*worker.Worker
}

func New(c *config.App, l *log.Logger) *Service {
	return &Service{
		logger:  l,
		buffer:  buffer.New(c.Buffer.Multiplier, c.System.CoreCount),
		workers: make([]*worker.Worker, c.System.CoreCount),
	}
}

func (s *Service) Init() {
	workerBuf := s.buffer.Size() / len(s.workers)
	for i := 0; i < len(s.workers); i++ {
		s.workers[i] = worker.New(i, workerBuf)
	}
}

func (s *Service) Run() {
	s.logger.Info("== Funny Go Benchmark v.0.1 ==")
	s.logger.Info("Memory Test...")
	s.logger.Infof("CPU count: %v", len(s.workers))

	bs, unit := units.FormatPrettyBytes(s.buffer.Size() * units.Kilobyte)
	s.logger.Infof("target buffer size: %.2f %v", bs, unit)

	timer1 := timer.Timer{}
	timer1.Start()

	wg := sync.WaitGroup{}
	wg.Add(len(s.workers))
	routineBufSize := make(chan int, len(s.workers))
	for _, w := range s.workers {
		go func(wrk *worker.Worker) {
			defer wg.Done()

			bs, unit := units.FormatPrettyBytes(wrk.BufferCap() * units.Kilobyte)
			s.logger.Infof("[%v] target routine buffer size: %.2f %v", wrk.ID(), bs, unit)

			routineBufSize <- wrk.Run(s.buffer)
		}(w)
	}

	wgB := sync.WaitGroup{}
	wgB.Add(1)
	var resultBufSize int
	go func() {
		defer wgB.Done()

		for rbs := range routineBufSize {
			bs, unit := units.FormatPrettyBytes(rbs)
			s.logger.Infof("routine buffer size: %.2f %v", bs, unit)
			resultBufSize += rbs
		}

		bs, unit = units.FormatPrettyBytes(resultBufSize)
		s.logger.Infof("total buffer size: %.2f %v", bs, unit)
	}()

	wg.Wait()
	close(routineBufSize)
	wgB.Wait()
	timer1.Stop()

	s.logger.Infof("test duration: %v", timer1.Result())
	s.logger.Info("Memory Test finished")
}
