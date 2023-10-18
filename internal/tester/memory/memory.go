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
	return &Tester{
		logger:  logger,
		buffer:  buffer,
		workers: makeWorkers(buffer.Size(), routines),
	}
}

func makeWorkers(bs, rcount int) []*Worker {
	workers := make([]*Worker, rcount)
	workerBuf := bs / rcount
	for i := 0; i < rcount; i++ {
		workers[i] = &Worker{
			id:     i,
			buffer: make([][units.Kilobyte]byte, 0, workerBuf),
		}
	}

	return workers
}

func (t *Tester) Run() {
	t.logger.Infoln("Memory Test...")
	t.logger.Infolnf("CPU count: %v", len(t.workers))

	bs, unit := units.FormatPrettyBytes(t.buffer.Size() * units.Kilobyte)
	t.logger.Infolnf("target buffer size: %.2f %v", bs, unit)

	timer1 := timer.Timer{}
	timer1.Start()

	wg := sync.WaitGroup{}
	wg.Add(len(t.workers))
	results := make(chan *Result, len(t.workers))
	for _, w := range t.workers {
		go func(wrk *Worker) {
			defer wg.Done()

			bs, unit := units.FormatPrettyBytes(wrk.BufferCap() * units.Kilobyte)
			t.logger.Infolnf("[%v] target routine buffer size: %.2f %v", wrk.ID(), bs, unit)

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
			t.logger.Infolnf("[%v] routine buffer size: %.2f %v", wr.WorkerID(), bs, unit)
			resultBufSize += wr.BufferSize()
		}

		bs, unit = units.FormatPrettyBytes(resultBufSize)
		t.logger.Infolnf("total buffer size: %.2f %v", bs, unit)
	}()

	wg.Wait()
	close(results)
	wgB.Wait()
	timer1.Stop()

	t.logger.Infolnf("test duration: %v", timer1.Result())
	t.logger.Infoln("Memory Test finished")
}
