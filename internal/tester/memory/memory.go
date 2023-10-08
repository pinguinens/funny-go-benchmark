package memory

import (
	"sync"

	"github.com/pinguinens/funny-go-benchmark/internal/buffer"

	"github.com/pinguinens/funny-go-benchmark/pkg/log"
	"github.com/pinguinens/funny-go-benchmark/pkg/timer"
	"github.com/pinguinens/funny-go-benchmark/pkg/units"
)

type Tester struct {
	logger  *log.Logger
	buffer  *buffer.Buffer
	workers []*Worker
}

func New(logger *log.Logger, buffer *buffer.Buffer, routines int) *Tester {
	workers := make([]*Worker, routines)
	workerBuf := buffer.Size() / routines
	for i := 0; i < routines; i++ {
		workers[i] = &Worker{
			id:     i,
			buffer: make([][units.Kilobyte]byte, 0, workerBuf),
		}
	}

	return &Tester{
		logger:  logger,
		buffer:  buffer,
		workers: workers,
	}
}

func (t *Tester) Run() {
	t.logger.Info("Memory Test...")
	t.logger.Infof("CPU count: %v", len(t.workers))

	bs, unit := units.FormatPrettyBytes(t.buffer.Size() * units.Kilobyte)
	t.logger.Infof("target buffer size: %.2f %v", bs, unit)

	timer1 := timer.Timer{}
	timer1.Start()

	wg := sync.WaitGroup{}
	wg.Add(len(t.workers))
	results := make(chan *Result, len(t.workers))
	for _, w := range t.workers {
		go func(wrk *Worker) {
			defer wg.Done()

			bs, unit := units.FormatPrettyBytes(wrk.BufferCap() * units.Kilobyte)
			t.logger.Infof("[%v] target routine buffer size: %.2f %v", wrk.ID(), bs, unit)

			results <- wrk.Run(t.buffer)
		}(w)
	}

	wgB := sync.WaitGroup{}
	wgB.Add(1)
	var resultBufSize int
	go func() {
		defer wgB.Done()

		for wr := range results {
			bs, unit := units.FormatPrettyBytes(wr.BufferSize())
			t.logger.Infof("[%v] routine buffer size: %.2f %v", wr.WorkerID(), bs, unit)
			resultBufSize += wr.BufferSize()
		}

		bs, unit = units.FormatPrettyBytes(resultBufSize)
		t.logger.Infof("total buffer size: %.2f %v", bs, unit)
	}()

	wg.Wait()
	close(results)
	wgB.Wait()
	timer1.Stop()

	t.logger.Infof("test duration: %v", timer1.Result())
	t.logger.Info("Memory Test finished")
}
